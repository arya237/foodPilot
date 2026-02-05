package bot

import tele "gopkg.in/telebot.v3"

func onStart(c tele.Context) error {
	keyboard := &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	btnAutoReserve := keyboard.Text(btnAutoReserve)
	btnAboutUs := keyboard.Text(btnAboutUs)

	keyboard.Reply(
		keyboard.Row(btnAutoReserve),
		keyboard.Row(btnAboutUs),
	)
	return c.Send("به فود پایلوت خوش آمدید", keyboard)
}

func aboutUs(c tele.Context) error {
	return c.Send("ما خیلی خفنیم")
}

func autoRserve(c tele.Context) error {
	return c.Send("این فیچر در حال توسعه هستش... میاد انشالله")
}
