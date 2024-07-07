package games

type Game struct {
	ID       *int    `json:"id"`
	Name     *string `json:"name"`
	Favorite *bool   `json:"favorite"`
	Type     *string `json:"type"`
}
