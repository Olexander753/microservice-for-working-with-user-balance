package schema

// для поплнения, списания, проверки баланса
type Balance struct {
	Id       string  `json:"id" db:"id" binding:"required"`         // id пользователся
	Amount   float32 `json:"amount" db:"amount" binding:"required"` // баланс пользователя в рублях
	Currency string  `json:"currency"`
}

// для перевода и проверки истории
type Transaction struct {
	ID        int     `json:"id" db:"id"`                                  // id транзакции
	Sender    string  `json:"sender" db:"sender" binding:"required"`       // отправитель
	Recipient string  `json:"recipient" db:"recipient" binding:"required"` // получатель
	Balance   int     `json:"balance" db:"balance"`                        // баланс после операции
	Amount    float32 `json:"amount" db:"amount" binding:"required"`       // сумма поперации
	Operation string  `json:"operation" db:"operation"`                    // операция
	Date      string  `json:"date" db:"date_"`                             // дата
}
