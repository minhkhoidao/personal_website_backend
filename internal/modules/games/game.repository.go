package games

import (
	"database/sql"
	"log"
)

type IGameRepository interface {
	GetAll() ([]*Game, error)
	Create(game *Game) error
}

type GameRepository struct {
	db *sql.DB
}

func NewGameRepository(db *sql.DB) *GameRepository {
	return &GameRepository{
		db: db,
	}
}

func (r *GameRepository) Create(game *Game) error {
	_, err := r.db.Exec("INSERT INTO game (name, favorite, type) VALUES ($1, $2, $3)", game.Name, game.Favorite, game.Type)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (r *GameRepository) GetAll() ([]*Game, error) {
	var games []*Game
	rows, err := r.db.Query("SELECT * FROM game")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close() // Ensure rows are closed

	for rows.Next() {
		var game Game
		err := rows.Scan(&game.ID, &game.Name, &game.Favorite, &game.Type)
		if err != nil {
			log.Println(err)
			return nil, err // Return on error
		}
		games = append(games, &game)
	}

	if err = rows.Err(); err != nil { // Check for errors after iterating
		log.Println(err)
		return nil, err
	}

	return games, nil
}
