package convert

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/config"
	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/schema"
)

type Conv struct {
	Result float32 `json:"result"`
}

func Convert(balance schema.Balance, currency string) (schema.Balance, error) {
	cfg := config.GetConfig()

	url := fmt.Sprintf("https://api.apilayer.com/exchangerates_data/convert?to=%s&from=%s&amount=%v", currency, "RUB", balance.Amount)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", cfg.ConvertAPI.Apikey)

	if err != nil {
		log.Println(err)
		return balance, err
	}

	res, err := client.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		log.Println(err)
		return balance, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return balance, err
	}

	result := Conv{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Println(err)
		return balance, err
	}

	balance.Amount = result.Result
	balance.Currency = currency

	return balance, nil
}
