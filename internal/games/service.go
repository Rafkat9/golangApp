package games

import "context"

// use cases
type Service interface {
	GameByName(ctx context.Context, name string) *Game
	GameAll(ctx context.Context, limit, offset int) []*Game
	// CreateGame(ctx context.Context, dto *CreateGameDTO) *Game
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) GameByName(ctx context.Context, name string) *Game {
	return s.storage.GetOne(name)
}
func (s *service) GameAll(ctx context.Context, limit, offset int) []*Game {
	return s.storage.GetAll(limit, offset)
}
