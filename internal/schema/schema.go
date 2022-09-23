package schema

// для поплнения, списания, проверки баланса
type Balance struct {
	Id     string `json:"id" db:"id" binding:"required"`         // id пользователся
	Amount int    `json:"amount" db:"amount" binding:"required"` // баланс пользователя в рублях
}

// для перевода и проверки истории
type Transaction struct {
	ID        int    `json:"id" db:"id"`                                  // id транзакции
	Sender    string `json:"sender" db:"sender" binding:"required"`       // отправитель
	Recipient string `json:"recipient" db:"recipient" binding:"required"` // получатель
	Amount    int    `json:"amount" db:"amount" binding:"required"`       // сумма перевода в рублях
	Operation string `json:"operation" db:"operation"`                    // операция
	Date      string `json:"date" db:"date_"`                             // дата
}
