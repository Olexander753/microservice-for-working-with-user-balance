package schema

// для поплнения, списания, проверки баланса
type Balance struct {
	Id     int `json:"id"`
	Amount int `json:"amount"`
}

// для перевода и проверки истории
type Transaction struct {
	Sender    int `json:"sender"`
	Recipient int `json:"recipient"`
	Amount    int `json:"amount"`
}
