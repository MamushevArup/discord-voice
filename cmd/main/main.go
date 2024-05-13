package main

import (
	"github.com/MamushevArup/ds-voice/adapters/avatar/aws"
	"github.com/MamushevArup/ds-voice/config"
	"github.com/MamushevArup/ds-voice/internal/handler"
	"github.com/MamushevArup/ds-voice/internal/usecase"
	"github.com/MamushevArup/ds-voice/pkg/logger"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file %v", err)
	}

	lg := logger.NewLogger()

	lg.Info("Init logger...")

	cfg, err := config.New()
	if err != nil {
		lg.Fatalf("%v", err)
	}

	lg.Info("Init config...")

	dsSession, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		lg.Fatalf("error creating Discord session: %v", err)
	}

	defer dsSession.Close() // nolint: errcheck

	s3, err := aws.New(cfg)
	if err != nil {
		lg.Fatalf("%v", err)
	}

	useCase := usecase.New(s3)

	hdl := handler.New(dsSession, useCase, lg)

	err = hdl.Start()
	if err != nil {
		lg.Fatalf("%v", err)
	}

	select {}
}
