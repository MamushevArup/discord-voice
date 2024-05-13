package discord

import "fmt"

// StartBot main entry point for `frontend of the discord`
func (b *Bot) StartBot() error {

	err := b.session.Open()
	if err != nil {
		b.log.Errorf("error opening connection: %v", err)
		return fmt.Errorf("error opening connection: %w", err)
	}

	b.session.AddHandler(b.voiceI.HandleChannelCreate)
	b.session.AddHandler(b.command.HandleCommand)

	b.log.Info("Bot is now running. Press CTRL+C to exit.")

	return nil
}
