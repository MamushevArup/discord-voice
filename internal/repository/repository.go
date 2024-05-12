package repository

import (
	"github.com/MamushevArup/ds-voice/config"
	"github.com/MamushevArup/ds-voice/internal/repository/mongod"
	"github.com/MamushevArup/ds-voice/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Mongo *mongod.Repo
}

func New(db *mongo.Database, lg *logger.Logger, cfg *config.Config) *Repository {
	return &Repository{
		Mongo: mongod.New(db, lg, cfg),
	}
}
