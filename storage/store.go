package storage

type Store interface {
	CreateSession() (session string)
	StoreGrid(session string, grid []int) error
	LoadGrid(session string) (grid []int, err error)
}
