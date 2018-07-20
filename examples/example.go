package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"

	sf "github.com/joelhill/mysportsfeeds-go"
)

func main() {
	config := sf.NewConfig("Basic setthishere")

	client := sf.NewService(config)
	c := context.Background()
	dailyGameOptions := client.NewDailyGamesOptions()
	dailyGameOptions.Date = "20180715"
	sg, sgErr := client.DailyGames(c, dailyGameOptions)

	spew.Dump(sg)
	spew.Dump(sgErr)
}
