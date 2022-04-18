package main

import (
	"github.com/Jamshid7/success-bot/botops"
)

// @Title Telegram Server v1
// @description API to send message
// @Schemes http
func main() {
	m1 := botops.NewMessage("The First MEssage", "medium", 2)
	m2 := botops.NewMessage("Highest MEssage", "high", 3)

	database := botops.Messages{}

	database.Msgs = append(database.Msgs, *m1)
	database.Msgs = append(database.Msgs, *m2)

	botops.InitMessages(&database)

	botops.Start()
	botops.Router()
}
