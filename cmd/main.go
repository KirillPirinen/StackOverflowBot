package main

import (
	"fmt"
	"log"
	n "shazam-go/pkg/overflow"
	"strings"
	"time"
	tele "gopkg.in/telebot.v3"
)

var (
	// Universal markup builders.
	menu = &tele.ReplyMarkup{ResizeKeyboard: true, OneTimeKeyboard: true}

	// Reply buttons.
	btnJS	= menu.Text("⚙ JS")
	btnCSS = menu.Text("⚙ CSS")
)

var userData = make(map[int64]string)

func main() {
	pref := tele.Settings{
		Token:  "5555627299:AAGD54B6OMsphWmtI5a6o0Mxcj6Vxt_v_qE",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tele.Context) error {

		menu.Reply(
			menu.Row(btnJS, btnCSS),
		)

		return c.Send(`
		AAA! Что делать?? код не работает!
		Отправь боту ошибку или вопрос. Для лучшего поиска можете выбрать ЯП.
		`, menu)
	})

	b.Handle(&btnJS, func(c tele.Context) error {
		userData[c.Sender().ID] = "javascript"
		return c.Send("Ищем по джаваскрипту") 
	})

	b.Handle(&btnCSS, func(c tele.Context) error {
		userData[c.Sender().ID] = "css"
		return c.Send("Ищем по стилям") 
	})

	b.Handle(tele.OnText, func(c tele.Context) error {	
		var (
			user = c.Sender()
			text = c.Text()
		)
		
		tag, ok := userData[user.ID]

		if !ok {
			tag = ""
		}

		ans := n.SearchThroughOverflow(text, tag)

		var res strings.Builder

		res.WriteString(fmt.Sprintf(`
		По запросу %v мы нашли <b>%v</b> вопросов:
		`, text, len(ans.Items)))

		for i, answer := range ans.Items {
			var line = fmt.Sprintf(`
			%v) <a href="%v">%v</a>
			`, i + 1, answer.Link, answer.Title)
			res.WriteString(line)

			if i == 10 {
				break
			}

		}

		return c.Send(res.String(), tele.ModeHTML)
	})

	b.Start()
}
