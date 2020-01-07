package main

import (
	"fmt"
	"log"
	"net/http"
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
		http.Get("https://api.telegram.org/bot1030683957:AAF-Yo-CbCm1M6Zi5gf8zaNyMQf462PXSd0/sendMessage?chat_id=-355434429&text=Moetnaarbed")
		b.Send(m.Sender, fmt.Sprintf("You entered: %s %s %s %t %d %s %s %d", m.Payload, m.Sender.FirstName, m.Sender.Username, m.Sender.IsBot, m.Chat.ID, m.Chat.Type, m.Chat.Title, m.Sender.ID))
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
		Text:   "Doenja",
	}

	inlineBtn_Poelcherija := tb.InlineButton{
		Unique: "Poelcherija",
		Text:   "Poelcherija",
	}

	inlineBtn_Razoemichin := tb.InlineButton{
		Unique: "Razoemichin",
		Text:   "Razoemichin",
	}

	inlineBtn_Porfiri := tb.InlineButton{
		Unique: "Porfiri",
		Text:   "Porfiri",
	}

	inlineBtn_Alyona := tb.InlineButton{
		Unique: "Alyona",
		Text:   "Alyona",
	}

	inlineBtn_Lizaveta := tb.InlineButton{
		Unique: "Lizaveta",
		Text:   "Lizaveta",
	}
	b.Handle(&inlineBtn_Avdotja, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		b.Send(c.Sender, "Avdotja Romanovna Raskalnikov, we tend to call here Doenja, is my sister.")
	})
	b.Handle(&inlineBtn_Poelcherija, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		b.Send(c.Sender, "Poelcherija Aleksandrovna Raskalnikov is my beloved mother.")
	})
	b.Handle(&inlineBtn_Razoemichin, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		b.Send(c.Sender, "I know Dmitri Prokofjitsj Razoemichin from my studies.")
	})

	b.Handle(&inlineBtn_Alyona, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		b.Send(c.Sender, "Alyona Ivanovna is a 60-years old devil pawnbroker. She domineers and emotionally abuses her sister Lizaveta.")
	})
	b.Handle(&inlineBtn_Lizaveta, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		b.Send(c.Sender, "Lizaveta Ivanovna is 35-years old, unmarried, over six feet tall and mentally challanged. She is often described as slow, soft, shy, timid and submisive.")
	})

	inlineCharacterSisterKeys := [][]tb.InlineButton{
		[]tb.InlineButton{inlineBtn_Alyona, inlineBtn_Lizaveta},
	}

	_ = tb.SendOptions{
		//&tb.Message{Text: testert},
		//&tb.ReplyMarkup{InlineKeyboard: inlineCharacterSisterKeys},
		//DisableWebPagePreview: false,
	}

	b.Handle(&inlineBtn_Porfiri, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		b.Send(c.Sender, "Pjotr Petrovitsj Porfiri is the police officer in charge of investigating the murder of Aljona en Lizaveta Ivanovna.", inlineCharacterSisterKeys)
	})

	inlineCharacterKeys := [][]tb.InlineButton{
		[]tb.InlineButton{inlineBtn_Avdotja, inlineBtn_Poelcherija},
		[]tb.InlineButton{inlineBtn_Razoemichin, inlineBtn_Porfiri},
	}

	b.Handle("/characters", func(m *tb.Message) {
		b.Send(
			m.Sender,
			"Which friend do you want to know more about?",
			&tb.ReplyMarkup{InlineKeyboard: inlineCharacterKeys})
	})

	b.Start()
}
