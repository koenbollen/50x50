package storage

// Store is the interface to store the data of a 50x50 grid by
// a session identifier.
type Store interface {
	CreateSession() (session string)
	StoreGrid(session string, grid []int) error
	LoadGrid(session string) (grid []int, err error)
}
