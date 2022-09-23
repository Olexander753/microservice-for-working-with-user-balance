package schema

// для поплнения, списания, проверки баланса
type Balance struct {
	Id     int `json:"id" db:"id" binding:"required"`
	Amount int `json:"amount" db:"amount" binding:"required"`
}

// для перевода и проверки истории
type Transaction struct {
	Sender    int `json:"sender" db:"sender" binding:"required"`
	Recipient int `json:"recipient" db:"recipient" binding:"required"`
	Amount    int `json:"amount" db:"amount" binding:"required"`
}
