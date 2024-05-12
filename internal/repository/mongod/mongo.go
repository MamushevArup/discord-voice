package mongod

import (
	"github.com/MamushevArup/ds-voice/config"
	"github.com/MamushevArup/ds-voice/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	db  *mongo.Database
	lg  *logger.Logger
	cfg *config.Config
}

func New(db *mongo.Database, lg *logger.Logger, cfg *config.Config) *Repo {
	return &Repo{
		db:  db,
		lg:  lg,
		cfg: cfg,
	}
}
