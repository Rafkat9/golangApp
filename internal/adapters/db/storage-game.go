package db

import "myappcorn/internal/games"

type db struct {
	// клиент к монго или обертка или еще что
	
}

func (db *db) GetOne(name string) *games.Game {
	return nil
}
func (db *db) GetAll(limit, offset int) []*games.Game {
	return nil

}
