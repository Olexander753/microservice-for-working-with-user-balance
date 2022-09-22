package schema

// для поплнения, списания, проверки баланса
type Balance struct {
	Id  int `json:"id"`
	Sum int `json:"sum"`
}

// для перевода и проверки истории
type Transaction struct {
	Sender    int `json:"sender"`
	Recipient int `json:"recipient"`
	Sum       int `json:"sum"`
}
