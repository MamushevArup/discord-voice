package main

import (
	"github.com/MamushevArup/ds-voice/cmd/event/bot"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("error loading .env file")
		return
	}

	s, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Println("error creating Discord session:", err)
		return
	}
	defer s.Close()

	b := bot.NewBot(s)

	err = b.StartBot()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}
