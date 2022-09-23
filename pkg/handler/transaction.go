package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getHistory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := fmt.Sprintf("error bad input id: %e", err)
		log.Println(msg)
		c.JSON(http.StatusBadGateway, msg) //TODO
	} else {
		transactions, err := h.services.Transaction.GetHistory(id)
		if err != nil {
			msg := fmt.Sprintf("error get history: %e", err)
			log.Println(msg)
			c.JSON(http.StatusBadGateway, msg) //TODO
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{
				"user id":    id,
				"Транзакции": transactions,
			})
		}
	}
}

func (h *Handler) transaction(c *gin.Context) {
	var input schema.Transaction

	if err := c.BindJSON(&input); err != nil {
		msg := fmt.Sprintf("error bad input transaction data: %e", err)
		log.Println(msg)
		c.JSON(http.StatusBadGateway, msg) //TODO
	} else {
		output, err := h.services.Transaction.Transaction(input)
		if err != nil {
			msg := fmt.Sprintf("error transaction: %e", err)
			log.Println(msg)
			c.JSON(http.StatusBadGateway, msg) //TODO
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Педевод успешно выполене",
				"user id": output.Id,
				"amount":  output.Amount,
			})
		}
	}
}