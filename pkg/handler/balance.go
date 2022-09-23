package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
	"github.com/gin-gonic/gin"
)

// методе поплнения баланса
func (h *Handler) replenishment(c *gin.Context) {
	var input schema.Balance

	// считывание входных данных, если ошибка отправляем ее в ответ
	if err := c.BindJSON(&input); err != nil {
		msg := fmt.Sprintf("error bad input replenishment data: %v", err)
		log.Println(msg)
		// ответ со статусом 400 (некорректный запрос) и сообщением
		c.JSON(http.StatusBadRequest, msg)
	} else {
		// вызов метода записи в бд данных
		output, err := h.services.Balance.Replenishment(input)
		// если в ходе записи в бд произошла ошибка то отправляем ее в ответ
		if err != nil {
			msg := fmt.Sprintf("error replenishment: %e", err)
			log.Println(msg)
			// ответ со статусом 400 (некорректный запрос) и сообщением
			c.JSON(http.StatusBadRequest, msg)
		} else {
			// если все прошло хорошо, то в ответе отправляем сообщение, id пользователя и его текущий баланс
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": fmt.Sprintf("Баланс успешно пополнен на сумма %v.", input.Amount),
				"user id": output.Id,
				"amount":  fmt.Sprintf("%v RUB", output.Amount),
			})
		}
	}

}

func (h *Handler) getBalance(c *gin.Context) {
	// считывание входных данных
	id := c.Param("id")
	// вызыв метода чтения данных из бд
	output, err := h.services.Balance.GetBalance(id)
	// если в ходе чтения из бд произошла ошибка то отправляем ее в ответ
	if err != nil {
		msg := fmt.Sprintf("error get balance: %v", err)
		log.Println(msg)
		// ответ со статусом 400 (некорректный запрос) и сообщением
		c.JSON(http.StatusBadRequest, msg)
	} else {
		// если все прошло хорошо, то в ответе отправляем id пользователя и его текущий баланс
		c.JSON(http.StatusOK, map[string]interface{}{
			"id":      id,
			"balance": fmt.Sprintf("%v RUB", output.Amount),
		})
	}

}

func (h *Handler) writeOff(c *gin.Context) {
	var input schema.Balance

	// считывание входных данных, если ошибка отправляем ее в ответ
	if err := c.BindJSON(&input); err != nil {
		msg := fmt.Sprintf("error bad input writeOff data: %v", err)
		log.Println(msg)
		// ответ со статусом 400 (некорректный запрос) и сообщением
		c.JSON(http.StatusBadRequest, msg)
	} else {
		// вызов метода изменения и записи данных в бд
		output, err := h.services.Balance.WriteOff(input)
		// если в ходе изменения данных в бд произошла ошибка то отправляем ее в ответ
		if err != nil {
			msg := fmt.Sprintf("error get balance: %v", err)
			log.Println(msg)
			// ответ со статусом 400 (некорректный запрос) и сообщением
			c.JSON(http.StatusBadRequest, msg)
		} else {
			// если все прошло хорошо, то в ответе отправляем статус 200 (ОК), сообщение, id пользователя и его текущий баланс
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": fmt.Sprintf("Средства успешно списаны на сумму %v.", input.Amount),
				"user id": output.Id,
				"amount":  fmt.Sprintf("%v RUB", output.Amount),
			})
		}
	}
}
