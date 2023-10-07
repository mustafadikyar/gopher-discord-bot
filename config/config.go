package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string // Discord botunun token'ı
	BotPrefix string // Bot'un komutları algılamak için kullanacağı prefix

	config *configStruct
)

// configStruct, config.json dosyasındaki yapılandırma verilerini temsil eder.
type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

// ReadConfig fonksiyonu, yapılandırma dosyasını okur ve verileri ilgili değişkenlere atar.
func ReadConfig() error {
	fmt.Println("Yapılandırma dosyası okunuyor...")

	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println((err.Error()))
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
