package main

import (
	"log"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL") // you must add it to your config vars
		token     = os.Getenv("TOKEN")      // you must add it to your config vars
	)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	//hello
	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "You entered "+m.Payload)
	})

	//buttons
	inlineBtn1 := tb.InlineButton{
		Unique: "moon",
		Text:   "Moon 🌚",
	}

	inlineBtn2 := tb.InlineButton{
		Unique: "sun",
		Text:   "Sun 🌞",
	}
	b.Handle(&inlineBtn1, func(c *tb.Callback) {
		// Required for proper work
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		// Send messages here
		b.Send(c.Sender, "Moon says 'Hi'!")
	})

	b.Handle(&inlineBtn2, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		b.Send(c.Sender, "Sun says 'Hi'!")
	})

	inlineKeys := [][]tb.InlineButton{
		[]tb.InlineButton{inlineBtn1, inlineBtn2},
	}

	b.Handle("/pick_time", func(m *tb.Message) {
		b.Send(
			m.Sender,
			"Day or night, you choose",
			&tb.ReplyMarkup{InlineKeyboard: inlineKeys})
	})

	inlineBtn_Avdotja := tb.InlineButton{
		Unique: "Doenja",
		Text:   "Avdotja Romanovna Raskalnikov",
	}

	inlineBtn_Poelcherija := tb.InlineButton{
		Unique: "Poelcherija",
		Text:   "Poelcherija Aleksandrovna Raskalnikov",
	}

	inlineBtn_Razoemichin := tb.InlineButton{
		Unique: "Razoemichin",
		Text:   "Dmitri Prokofjitsj Razoemichin",
	}

	inlineBtn_Porfiri := tb.InlineButton{
		Unique: "Porfiri",
		Text:   "Pjotr Petrovitsj Porfiri",
	}
	b.Handle(&inlineBtn_Avdotja, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		b.Send(c.Sender, "Avdotja Romanovna Raskalnikov, also called Doenja, is the sister of Rodion.")
	})
	b.Handle(&inlineBtn_Poelcherija, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		b.Send(c.Sender, "Poelcherija Aleksandrovna Raskalnikov is the mother of Rodion.")
	})
	b.Handle(&inlineBtn_Razoemichin, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		b.Send(c.Sender, "Dmitri Prokofjitsj Razoemichin is a study-friend of Raskalnikov.")
	})
	b.Handle(&inlineBtn_Porfiri, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		b.Send(c.Sender, "Pjotr Petrovitsj Porfiri is the police officer in charge of investigating the murder of Aljona en Lizaveta Ivanovna.")
	})

	inlineCharacterKeys := [][]tb.InlineButton{
		[]tb.InlineButton{inlineBtn_Avdotja, inlineBtn_Poelcherija, inlineBtn_Razoemichin, inlineBtn_Porfiri},
	}

	b.Handle("/characters", func(m *tb.Message) {
		b.Send(
			m.Sender,
			"Which friend do you want to know more about?",
			&tb.ReplyMarkup{InlineKeyboard: inlineCharacterKeys})
	})

	b.Start()
}
