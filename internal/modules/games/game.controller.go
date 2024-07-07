package games

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GameController struct {
	service GameService
}

func NewGameController(service GameService) *GameController {
	return &GameController{service: service}
}

// Ensure this method is correctly defined
func (gc *GameController) GetAllGames(c *gin.Context) {
	games, err := gc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching games"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"games": games,
	})
}
