package handler

import (
	"github.com/MamushevArup/ds-voice/internal/handler/discord"
	"github.com/MamushevArup/ds-voice/internal/usecase"
	"github.com/MamushevArup/ds-voice/pkg/logger"
	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	bot *discord.Bot
}

func New(ds *discordgo.Session, useCase *usecase.UseCase, logger *logger.Logger) *Handler {
	return &Handler{bot: discord.NewBot(ds, useCase, logger)}
}

func (h *Handler) Start() error {
	return h.bot.StartBot()
}
