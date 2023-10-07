package bot

import (
	"fmt"
	"golang-discord-bot/config"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

// Start fonksiyonu, botun başlatılması için kullanılır.
func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot çalışıyor!")
}

// messageHandler fonksiyonu, gelen mesajları işler.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Botun kendi gönderdiği mesajları işleme.
	if m.Author.ID == BotID {
		return
	}

	message := strings.ToLower(m.Content)

	// Eğer gelen mesaj "selam" veya "merhaba" ise cevap ver.
	if message == "selam" || message == "merhaba" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Merhaba gopher :)")
	}

	// Eğer gelen mesaj "iyi geceler" ise cevap ver.
	if message == "iyi geceler" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "İyi geceler gopher :p")
	}
}
