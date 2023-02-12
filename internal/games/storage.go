package games

type Storage interface {
	GetOne(name string) *Game
	GetAll(limit, offset int) []*Game
	// Create
	// Delete
}
