package handler

import (
	"github.com/Olexander753/microservice-for-working-with-user-balance/pkg/service"
	"github.com/gin-gonic/gin"
)

// структура Handler, включает в себя services
type Handler struct {
	services *service.Service
}

// констуктор для структуры Handler, принимает services и возвращает указатель на экземпляр структуры Handler
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// функция инициализации роутера
func (h *Handler) InutRoutes() *gin.Engine {
	router := gin.New()

	// прописывание пути для каждого используюемого метода протокола HTTP
	api := router.Group("/api")
	{
		api.POST("/", h.replenishment)              //пополнение средств
		api.GET("/:id", h.getBalance)               //проверка баланса по ID пользователя
		api.GET("/history/:id/", h.getHistory)      //получение истории транзакций по ID пользователя
		api.GET("/history/:id/:sort", h.getHistory) //получение истории транзакций по ID пользователя c сортрировкой
		api.PUT("/write-off", h.writeOff)           //списание средст
		api.PUT("/transaction", h.transaction)      //перевод средст
	}
	return router
}
