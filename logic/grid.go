package logic

type Grid struct {
	Width  int
	Height int

	Data []int
}

func NewGrid(width, height int) *Grid {
	return &Grid{
		Width:  width,
		Height: height,
		Data:   make([]int, width*height),
	}
}

func FromData(width int, data []int) *Grid {
	height := len(data) / width
	return &Grid{
		Width:  width,
		Height: height,
		Data:   data,
	}
}

func (g *Grid) index(x, y int) int {
	return x + y*g.Width
}

func (g *Grid) Increment(x, y int) {
	g.Data[g.index(x, y)] += 1
}

func (g *Grid) Set(x, y, value int) {
	g.Data[g.index(x, y)] = value
}

func (g *Grid) Get(x, y int) int {
	return g.Data[g.index(x, y)]
}
