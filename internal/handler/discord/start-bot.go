package discord

import "fmt"

// StartBot main entry point for `frontend of the discord`
func (b *Bot) StartBot() error {

	err := b.session.Open()
	if err != nil {
		return fmt.Errorf("error opening connection: %w", err)
	}

	b.session.AddHandler(b.handleChannelCreate)

	return nil
}
