package handler

import (
	"github.com/MamushevArup/ds-voice/internal/handler/discord"
	"github.com/MamushevArup/ds-voice/internal/usecase"
	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	useCase *usecase.UseCase
	bot     *discord.Bot
}

func New(ds *discordgo.Session) *Handler {
	return &Handler{bot: discord.NewBot(ds)}
}

func (h *Handler) Start() error {
	return h.bot.StartBot()
}
