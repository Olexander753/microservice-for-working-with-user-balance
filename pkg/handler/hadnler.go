package handler

import (
	"github.com/Olexander753/microservice-for-working-with-user-balance/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InutRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/", h.replenishment)         //пополнение средств
		api.GET("/:id", h.getBalance)          //проверка баланса по ID пользователя
		api.GET("/history/:id", h.getHistory)  //получение истории транзакций по ID пользователя
		api.PUT("/write-off", h.writeOff)      //списание средст
		api.PUT("/transaction", h.transaction) //перевод средст
	}
	return router
}
