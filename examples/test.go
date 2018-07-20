package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"

	sf "github.com/joelhill/mysportsfeeds-go"
)

func main() {
	config := sf.NewConfig("Basic amlnZ2lkeXVvOk1ZU1BPUlRTRkVFRFM=")

	client := sf.NewService(config)
	c := context.Background()
	spew.Dump(client)
	seasonalGamesOptions := client.NewGameLineupOptions()
	seasonalGamesOptions.Game = "20180506-CHC-STL"
	sg, sgErr := client.GameLineup(c, seasonalGamesOptions)
	spew.Dump(sg)
	spew.Dump(sgErr)
}
