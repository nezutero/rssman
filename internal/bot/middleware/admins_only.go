package middleware

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"main.go/internal/botkit"
)

func AdminsOnly(channelID int64, next botkit.ViewFunc) botkit.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		admins, err := bot.GetChatAdministrators(
			tgbotapi.ChatAdministratorsConfig{
				ChatConfig: tgbotapi.ChatConfig{
					ChatID: channelID,
				},
			},
		)

		if err != nil {
			return err
		}

		for _, admin := range admins {
			if admin.User.ID == update.SentFrom().ID {
				return next(ctx, bot, update)
			}
		}

		if _, err := bot.Send(tgbotapi.NewMessage(
			update.FromChat().ID,
			"You do not have permission to execute this command.",
		)); err != nil {
			return err
		}

		return nil
	}
}