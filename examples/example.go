package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"

	sf "github.com/joelhill/mysportsfeeds-go"
)

func main() {
	config := sf.NewConfig("Basic afsdfasdfasdfsadffdfasdfadffas")

	client := sf.NewService(config)
	c := context.Background()
	dailyGameOptions := client.NewDailyGamesOptions()
	dg, statusCode, dgErr := client.DailyGames(c, dailyGameOptions)

	spew.Dump(dg)
	spew.Dump(statusCode)
	spew.Dump(dgErr)
}
