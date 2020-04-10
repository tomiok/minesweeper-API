package minesweepersvc

type DB struct {
	games map[string]*Game
	users map[string]*User
}

func New() *DB {
	return &DB{
		games: make(map[string]*Game),
		users: make(map[string]*User),
	}
}


