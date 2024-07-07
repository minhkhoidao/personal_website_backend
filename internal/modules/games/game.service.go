package games

type GameService struct {
	r IGameRepository
}

func NewGameService(r IGameRepository) *GameService {
	return &GameService{
		r: r,
	}
}

func (s *GameService) Create(game *Game) error {
	return s.r.Create(game)
}

func (s *GameService) GetAll() ([]*Game, error) {
	return s.r.GetAll()
}
