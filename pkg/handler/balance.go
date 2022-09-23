package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/gin-gonic/gin"
)

func (h *Handler) replenishment(c *gin.Context) {
	var input schema.Balance

	if err := c.BindJSON(&input); err != nil {
		msg := fmt.Sprintf("error bad input replenishment data: %e", err)
		log.Println(msg)
		c.JSON(http.StatusBadGateway, msg) //TODO
	} else {
		output, err := h.services.Balance.Replenishment(input)
		if err != nil {
			msg := fmt.Sprintf("error replenishment: %e", err)
			log.Println(msg)
			c.JSON(http.StatusBadGateway, msg) //TODO
		} else {
			c.JSON(http.StatusOK, fmt.Sprintf("Баланс %v успешно пополнен на сумма %v. Текущий баланс %v", input.Id, input.Amount, output.Amount))
		}
	}

}

func (h *Handler) getBalance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := fmt.Sprintf("error bad input id: %e", err)
		log.Println(msg)
		c.JSON(http.StatusBadGateway, msg) //TODO
	} else {
		output, err := h.services.Balance.GetBalance(id)
		if err != nil {
			msg := fmt.Sprintf("error get balance: %e", err)
			log.Println(msg)
			c.JSON(http.StatusBadGateway, msg) //TODO
		} else {
			c.JSON(http.StatusOK, map[string]int{
				"id":      id,
				"balance": output.Amount,
			})
		}
	}

}

func (h *Handler) writeOff(c *gin.Context) {
	var input schema.Balance

	if err := c.BindJSON(&input); err != nil {
		msg := fmt.Sprintf("error bad input writeOff data: %e", err)
		log.Println(msg)
		c.JSON(http.StatusBadGateway, msg) //TODO
	} else {
		output, err := h.services.Balance.WriteOff(input)
		if err != nil {
			msg := fmt.Sprintf("error get balance: %e", err)
			log.Println(msg)
			c.JSON(http.StatusBadGateway, msg) //TODO
		} else {
			message := fmt.Sprintf("Средства успешно списаны на сумму %v. Текущий баланс %v", input.Amount, output.Amount)
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": message,
			})
		}
	}
}
