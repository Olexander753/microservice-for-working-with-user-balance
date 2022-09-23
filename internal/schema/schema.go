package schema

// для поплнения, списания, проверки баланса
type Balance struct {
	Id     string `json:"id" db:"id" binding:"required"`
	Amount int    `json:"amount" db:"amount" binding:"required"`
}

// для перевода и проверки истории
type Transaction struct {
	ID        int    `json:"id" db:"id"`
	Sender    string `json:"sender" db:"sender" binding:"required"`
	Recipient string `json:"recipient" db:"recipient" binding:"required"`
	Amount    int    `json:"amount" db:"amount" binding:"required"`
	Operation string `json:"operation" db:"operation"`
}
