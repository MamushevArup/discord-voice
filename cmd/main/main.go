package main

import (
	"context"
	"fmt"
	"github.com/MamushevArup/ds-voice/adapters/avatar/aws"
	"github.com/MamushevArup/ds-voice/config"
	"github.com/MamushevArup/ds-voice/internal/handler"
	"github.com/MamushevArup/ds-voice/internal/repository"
	"github.com/MamushevArup/ds-voice/internal/usecase"
	"github.com/MamushevArup/ds-voice/pkg/logger"
	"github.com/MamushevArup/ds-voice/pkg/mongodb"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

var configPath = "config/config.yml"

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lg := logger.NewLogger()

	lg.Info("Init logger...")

	cfg, err := config.New(configPath)
	if err != nil {
		lg.Fatalf("%v", err)
	}
	lg.Info("Init config...")

	client, err := mongoClient(ctx, cfg)
	if err != nil {
		lg.Fatalf("%v", err)
	}

	lg.Infof("Init mongo client... listening on port %s", cfg.Mongo.Port)

	dsSession, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		lg.Fatalf("error creating Discord session: %v", err)
	}
	defer dsSession.Close() // nolint: errcheck

	lg.Info("Bot is now running. Press CTRL+C to exit.")

	_ = repository.New(client, lg, cfg)

	_, err = aws.New(cfg)
	if err != nil {
		lg.Fatalf("%v", err)
	}

	usecase.New()

	hdl := handler.New(dsSession)

	err = hdl.Start()
	if err != nil {
		lg.Fatalf("%v", err)
	}

	select {}
}

func mongoClient(ctx context.Context, cfg *config.Config) (*mongo.Database, error) {
	url := fmt.Sprintf("mongodb://%s:%s", cfg.Mongo.Host, cfg.Mongo.Port)
	mclient, err := mongodb.NewClient(ctx, url, cfg.Mongo.Database)
	if err != nil {
		return nil, fmt.Errorf("failed to create mongo client: %w", err)
	}

	return mclient, nil
}
