package logic

type Grid struct {
	Size int

	Data []int
}

func NewGrid(size int) *Grid {
	return &Grid{
		Size: size,
		Data: make([]int, size*size),
	}
}

func FromData(size int, data []int) *Grid {
	return &Grid{
		Size: size,
		Data: data,
	}
}

func (g *Grid) index(x, y int) int {
	return x + y*g.Size
}

func (g *Grid) Increment(x, y int) {
	for i := 0; i < g.Size; i++ {
		g.Data[g.index(x, i)] += 1
		if i != x { // prevent double increment on cross
			g.Data[g.index(i, y)] += 1
		}
	}
}

func (g *Grid) Set(x, y, value int) {
	g.Data[g.index(x, y)] = value
}

func (g *Grid) Get(x, y int) int {
	return g.Data[g.index(x, y)]
}

func (g *Grid) ClearPaths(paths [][]int) {
	for _, path := range paths {
		for _, i := range path {
			g.Data[i] = 0
		}
	}
}

func (g *Grid) SearchForSequence() (paths [][]int) {
	for i := 0; i < g.Size*g.Size; i++ {
		value := g.Data[i]
		if IsFibonacci(value) {
			path := g.searchRight(i)
			i += len(path) - 1
			if len(path) >= 5 {
				paths = append(paths, path)
			}

			path = g.searchLeft(i)
			if len(path) >= 5 {
				paths = append(paths, path)
			}

			path = g.searchDown(i)
			if len(path) >= 5 {
				paths = append(paths, path)
			}

			path = g.searchUp(i)
			if len(path) >= 5 {
				paths = append(paths, path)
			}
		}
	}
	return
}

func (g *Grid) searchRight(ix int) []int {
	path := []int{ix}
	x := ix % g.Size
	for i := 1; i < g.Size-x; i++ {
		value := g.Data[ix+i]
		if IsFibonacci(value) && (len(path) <= 1 || value == g.Data[path[i-1]]+g.Data[path[i-2]]) {
			path = append(path, ix+i)
		} else {
			break
		}
	}
	return path
}

func (g *Grid) searchLeft(ix int) []int {
	path := []int{ix}
	x := ix % g.Size
	for i := 1; i < x+1; i++ {
		value := g.Data[ix-i]
		if IsFibonacci(value) && (len(path) <= 1 || value == g.Data[path[i-1]]+g.Data[path[i-2]]) {
			path = append(path, ix-i)
		} else {
			break
		}
	}
	return path
}

func (g *Grid) searchDown(ix int) []int {
	path := []int{ix}
	y := ix / g.Size
	for i := 1; i < g.Size-y; i++ {
		value := g.Data[ix+i*g.Size]
		if IsFibonacci(value) && (len(path) <= 1 || value == g.Data[path[i-1]]+g.Data[path[i-2]]) {
			path = append(path, ix+i*g.Size)
		} else {
			break
		}
	}
	return path
}

func (g *Grid) searchUp(ix int) []int {
	path := []int{ix}
	y := ix / g.Size
	for i := 1; i < y+1; i++ {
		value := g.Data[ix-i*g.Size]
		if IsFibonacci(value) && (len(path) <= 1 || value == g.Data[path[i-1]]+g.Data[path[i-2]]) {
			path = append(path, ix-i*g.Size)
		} else {
			break
		}
	}
	return path
}
