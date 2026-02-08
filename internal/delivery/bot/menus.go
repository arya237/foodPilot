package bot

import tele "gopkg.in/telebot.v3"

var MENU_CACHE = map[bool]*tele.ReplyMarkup{}

var (
	btnAboutUs           = "درباره ما"
	btnAutoReserve       = "رزور خودکار"
	btnRestaurantSetting = "تنظیمات رزور خودکار"
)

func mainMenu(sammadConnection bool) *tele.ReplyMarkup {
	if menu, ok := MENU_CACHE[sammadConnection]; ok {
		return menu
	}

	keyboard := &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}
	btnAutoReserve := keyboard.Text(btnAutoReserve)
	if sammadConnection {
		btnAutoReserve = keyboard.Text(btnRestaurantSetting)
	}
	btnAboutUs := keyboard.Text(btnAboutUs)

	keyboard.Reply(
		keyboard.Row(btnAutoReserve),
		keyboard.Row(btnAboutUs),
	)

	MENU_CACHE[sammadConnection] = keyboard
	return MENU_CACHE[sammadConnection]
}
