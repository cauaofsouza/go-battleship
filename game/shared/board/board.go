package board

//estrutura do tabuleiro

const (
	Rows = 7
	Cols = 8
)

type Board struct {
	Cells [][]Cell
	X     float64 // posição na tela
	Y     float64
	Size  float64 // tamanho total
}

func NewBoard(x, y, size float64) *Board {
	cells := make([][]Cell, Rows)
	for i := 0; i < Rows; i++ {
		cells[i] = make([]Cell, Cols)
		for j := 0; j < Cols; j++ {
			cells[i][j] = Cell{
				Row:   i,
				Col:   j,
				State: Empty,
			}
		}
	}

	return &Board{
		Cells: cells,
		X:     x,
		Y:     y,
		Size:  size,
	}
}
