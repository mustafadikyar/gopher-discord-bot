package main

import (
	"fmt"
	"golang-discord-bot/bot"
	"golang-discord-bot/config"
)

func main() {
	// config dosyasını oku
	err := config.ReadConfig()

	if err != nil {
		// Hata durumunda ekrana yazdır
		fmt.Println(err.Error())
		return
	}

	// Discord botunu başlat
	bot.Start()

	// Programın sonlanmasını engellemek için bir kanal bekletiliyor
	<-make(chan struct{})
	return
}
