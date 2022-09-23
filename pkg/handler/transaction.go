package handler

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getHistory(c *gin.Context) {
	// считывание входных данных
	id := c.Param("id")
	sort_ := c.Param("sort")
	// вызыв метода чтения данных из бд
	transactions, err := h.services.Transaction.GetHistory(id)
	// если в ходе чтения из бд произошла ошибка то отправляем ее в ответ
	if err != nil {
		msg := fmt.Sprintf("error get history: %v", err)
		log.Println(msg)
		// ответ со статусом 400 (некорректный запрос) и сообщением
		c.JSON(http.StatusBadRequest, msg)
	} else {
		// сортировка списка операций
		switch sort_ {
		case "date<":
			sort.Slice(transactions, func(i, j int) (less bool) {
				return transactions[i].Date < transactions[j].Date
			})
		case "date>":
			sort.Slice(transactions, func(i, j int) (less bool) {
				return transactions[i].Date > transactions[j].Date
			})
		case "amount<":
			sort.Slice(transactions, func(i, j int) (less bool) {
				return transactions[i].Amount < transactions[j].Amount
			})
		case "amount>":
			sort.Slice(transactions, func(i, j int) (less bool) {
				return transactions[i].Amount > transactions[j].Amount
			})
		case "balance<":
			sort.Slice(transactions, func(i, j int) (less bool) {
				return transactions[i].Balance < transactions[j].Balance
			})
		case "balance>":
			sort.Slice(transactions, func(i, j int) (less bool) {
				return transactions[i].Balance > transactions[j].Balance
			})
		}

		// если все прошло хорошо, то в ответе отправляем статус 200 (ОК) id пользователя и его текущий баланс
		c.JSON(http.StatusOK, map[string]interface{}{
			"user id":    id,
			"Транзакции": transactions,
		})
	}
}

func (h *Handler) transaction(c *gin.Context) {
	var input schema.Transaction

	// считывание входных данных, если ошибка отправляем ее в ответ
	if err := c.BindJSON(&input); err != nil {
		msg := fmt.Sprintf("error bad input transaction data: %v", err)
		log.Println(msg)
		// ответ со статусом 400 (некорректный запрос) и сообщением
		c.JSON(http.StatusBadRequest, msg)
	} else {
		// вызов метода изменения и записи данных в бд
		output, err := h.services.Transaction.Transaction(input)
		// если в ходе изменения данных в бд произошла ошибка то отправляем ее в ответ
		if err != nil {
			msg := fmt.Sprintf("error transaction: %s", err)
			log.Println(msg)
			// ответ со статусом 400 (некорректный запрос) и сообщением
			c.JSON(http.StatusBadRequest, msg)
		} else {
			// если все прошло хорошо, то в ответе отправляем статус 200 (ОК), сообщение, id пользователя и его текущий баланс
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Педевод успешно выполене",
				"user id": output.Id,
				"amount":  fmt.Sprintf("%v RUB", output.Amount),
			})
		}
	}
}
