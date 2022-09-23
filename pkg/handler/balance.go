package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/gin-gonic/gin"
)

func (h *Handler) replenishment(c *gin.Context) {
	var input schema.Balance

	if err := c.BindJSON(&input); err != nil {
		msg := fmt.Sprintf("error bad input replenishment data: %v", err)
		log.Println(msg)
		c.JSON(http.StatusBadGateway, msg) //TODO
	} else {
		output, err := h.services.Balance.Replenishment(input)
		if err != nil {
			msg := fmt.Sprintf("error replenishment: %e", err)
			log.Println(msg)
			c.JSON(http.StatusBadGateway, msg) //TODO
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": fmt.Sprintf("Баланс успешно пополнен на сумма %v.", input.Amount),
				"user id": output.Id,
				"amount":  output.Amount,
			})
		}
	}

}

func (h *Handler) getBalance(c *gin.Context) {
	id := c.Param("id")
	output, err := h.services.Balance.GetBalance(id)
	if err != nil {
		msg := fmt.Sprintf("error get balance: %v", err)
		log.Println(msg)
		c.JSON(http.StatusBadGateway, msg) //TODO
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"id":      id,
			"balance": output.Amount,
		})
	}

}

func (h *Handler) writeOff(c *gin.Context) {
	var input schema.Balance

	if err := c.BindJSON(&input); err != nil {
		msg := fmt.Sprintf("error bad input writeOff data: %v", err)
		log.Println(msg)
		c.JSON(http.StatusBadGateway, msg) //TODO
	} else {
		output, err := h.services.Balance.WriteOff(input)
		if err != nil {
			msg := fmt.Sprintf("error get balance: %v", err)
			log.Println(msg)
			c.JSON(http.StatusBadGateway, msg) //TODO
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": fmt.Sprintf("Средства успешно списаны на сумму %v.", input.Amount),
				"user id": output.Id,
				"amount":  output.Amount,
			})
		}
	}
}
