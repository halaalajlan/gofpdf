package gofpdf

type Table struct {
	Cols   int
	Rows   int
	Header []string
	Data   [][]string
	W      []float64
	wSum   float64
	left   float64
}

func (f *Fpdf) CreateTable(cols int, rows int, width []float64) {

	f.table.Cols = cols
	f.table.Rows = rows
	f.table.W = width
	wSum := 0.0
	for _, v := range width {
		wSum += v
	}
	f.table.wSum = wSum
	if f.defOrientation == "P" {
		f.table.left = (210 - wSum) / 2

	} else { //landscape
		f.table.left = (297 - wSum) / 2

	}
}
func (f *Fpdf) TableData(header []string, data [][]string) {
	f.table.Header = header
	f.table.Data = data
	f.SetX(f.table.left)
	for j, str := range header {
		f.CellFormat(f.table.W[j], 7, str, "1", 0, "C", true, 0, "")
	}
	marginCell := 2. // margin of top/bottom of cell

	_, pageh := f.GetPageSize()
	_, _, _, mbottom := f.GetMargins()

	for _, row := range data {
		curx, y := f.GetXY()
		x := curx

		height := 0.
		_, lineHt := f.GetFontSize()

		for i, txt := range row {
			lines := f.SplitLines([]byte(txt), f.table.W[i])
			h := float64(len(lines))*lineHt + marginCell*float64(len(lines))
			if h > height {
				height = h
			}
		}
		// add a new page if the height of the row doesn't fit on the page
		if f.GetY()+height > pageh-mbottom {
			f.AddPage()
			y = f.GetY()
		}
		for i, txt := range row {
			width := f.table.W[i]
			f.Rect(x, y, width, height, "")
			f.MultiCell(width, lineHt+marginCell, txt, "", "", false)
			x += width
			f.SetXY(x, y)
		}
		f.SetXY(curx, y+height)
	}

}
