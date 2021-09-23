package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"log"
	"tcp-chat/internal/app"
)

var (
	configPath string
)

func init() {
	// передаём указатель на строковую переменную, в которую будет записываться путь к файлу конфигрурации
	flag.StringVar(&configPath, "config-path", "configs/tcpchat.toml", "path to config file")
}

func main() {
	flag.Parse()

	logger := logrus.New()                     // инициализируем логгер
	config := app.NewConfig()                  // инициализируем конфигурацию
	_, err := toml.Decode(configPath, &config) // декодируем файл конфигурации в структуру
	if err != nil {
		log.Fatal(err.Error())
		//logger.Info("Config file not found. Chat's been run with default settings.")
	}
	// инициализируем и запускаем приложение
	chat := app.New(config, logger)
	chat.Start()
}
