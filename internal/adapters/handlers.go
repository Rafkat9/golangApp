package adapters

import (
	"fmt"
	"myappcorn/internal/games"

	"github.com/gin-gonic/gin"
)

type handler struct {
	gameService games.Service
}

func NewHandler(service games.Service) Handler {
	return &handler{gameService: service}
}

func (h *handler) Register(ctx *gin.Context) {

	router := gin.Default()

	router.GET("/games/:id", h.GetOneGamess)
	router.GET("/games", h.GetAllGamess)
}

func (h *handler) GetOneGamess(ctx *gin.Context) {
	// this data files one
	fmt.Println("grg")
	// return h.service.GameByName(ctx, name)
}
func (h *handler) GetAllGamess(ctx *gin.Context) {
	// this data files more
	fmt.Println("grg")
	// return h.service.GameAll(ctx, limit,offset)

}
