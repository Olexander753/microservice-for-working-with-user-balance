package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postgres struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"postgres"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	ConvertAPI struct {
		Apikey string `yaml:"apikey"`
	} `yaml:"convertAPI"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	//выполниться ровно один раз, при следующих вызовах просто будет возвращать instance
	once.Do(func() {
		log.Println("Read configuration")
		instance = &Config{}
		err := cleanenv.ReadConfig("config.yaml", instance)
		if err != nil {
			dis, _ := cleanenv.GetDescription(instance, nil)
			log.Println(dis)
			log.Fatal(err)
		}
	})
	return instance
}
