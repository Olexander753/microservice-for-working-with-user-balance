package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/gin-gonic/gin"
)

// type getResponse struct {

// }

func (h *Handler) replenishment(c *gin.Context) {
	var input schema.Balance

	if err := c.BindJSON(&input); err != nil {
		log.Fatal("error bad input replenishment data: ", err)
	}

	//TODO

	// id, err := h.services.TransactionItem.CreateTransaction(input)
	// if err != nil {
	// 	log.Fatal("error insert: ", err)
	// }

	message := fmt.Sprintf("баланс успешно пополнен на %v", input.Sum)

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": message,
	})
}

func (h *Handler) getBalance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panicln(err) //TODO
	}

	sum := 0 // TODO

	c.JSON(http.StatusOK, map[string]int{
		"id":      id,
		"balance": sum,
	})
}

func (h *Handler) getHistory(c *gin.Context) {
	_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panicln(err) //TODO
	}
	//TODO

	transactions := []schema.Transaction{}
	c.JSON(http.StatusOK, map[string][]schema.Transaction{
		"Транзакции": transactions,
	})
}

func (h *Handler) writeOff(c *gin.Context) {
	var input schema.Balance

	if err := c.BindJSON(&input); err != nil {
		log.Fatal("error bad input writeOff data: ", err)
	}

	//TODO

	// err := h.services.TransactionItem.UpdateTransaction(input.ID, input.Discription)

	// if err != nil {
	// 	c.JSON(http.StatusOK, "Ошибка")
	// } else {
	// 	c.JSON(http.StatusOK, "OK")
	// }

}

func (h *Handler) transaction(c *gin.Context) {
	var input schema.Transaction

	if err := c.BindJSON(&input); err != nil {
		log.Fatal("error bad input transaction data: ", err)
	}

	// err := h.services.TransactionItem.UpdateTransaction(input.ID, input.Discription)

	// if err != nil {
	// 	c.JSON(http.StatusOK, "Ошибка")
	// } else {
	// 	c.JSON(http.StatusOK, "OK")
	// }

}
