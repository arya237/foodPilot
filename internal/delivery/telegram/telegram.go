package telegram

import (
	"errors"
	"fmt"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/services/auth"
	tele "gopkg.in/telebot.v3"
)

type Handler struct {
	bot  *tele.Bot
	auth auth.Auth
}

func AuthMiddleware(service auth.Auth) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			sender := c.Sender()

			if sender == nil {
				return next(c)
			}

			telegramID := fmt.Sprintf("%d", sender.ID)

			internalID, err := service.Login(models.TELEGRAM, telegramID)
			if err != nil {
				if !errors.Is(err, auth.ErrUserNotFound) {
					return c.Send(err)
				}
				internalID, err = service.SignUp(models.TELEGRAM, telegramID, &models.User{
					Username: sender.FirstName,
					Role:     models.RoleUser,
				})

				if err != nil {
					return c.Send(err)
				}

			}

			c.Set("id", internalID.Id)
			return next(c)
		}
	}
}
func Start(bot *tele.Bot, auth auth.Auth) error {

	if bot == nil {
		return errors.New("Bot is nil")
	}
	bot.Use(AuthMiddleware(auth))

	bot.Handle(tele.OnText, func(c tele.Context) error {
		id, ok := c.Get("id").(int)
		if !ok {
			return c.Send("no no")
		}
		return c.Send(fmt.Sprintf("id: %d \nmoew", id))
	})

	bot.Start()
	return nil
}
