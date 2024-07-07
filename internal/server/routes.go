package server

import (
	"net/http"
	"personal_backend/internal/modules/games"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	gameRepository := games.NewGameRepository(s.db.GetDB())
	gameServices := games.NewGameService(gameRepository)
	gameController := games.NewGameController(*gameServices)

	r.GET("/", gameController.GetAllGames)

	r.GET("/health", s.healthHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
