package board

//estrutura de cada célula

type CellState int

const (
	Empty CellState = iota
	Ship
	Hit
	Miss
)

type Cell struct {
	Row   int
	Col   int
	State CellState
}
