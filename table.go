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
}
