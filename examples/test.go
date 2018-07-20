package examples

import (
	"context"

	"github.com/davecgh/go-spew/spew"

	sf "github.com/joelhill/mysportsfeeds-go"
)

func test() {
	config := sf.NewConfig("Basic amlnZ2lkeXVvOk1ZU1BPUlRTRkVFRFM=")

	client := sf.NewService(config)
	c := context.Background()
	seasonalGamesOptions := client.NewSeasonalGamesOptions()
	sg, sgErr := client.SeasonalGames(c, seasonalGamesOptions)
	spew.Dump(sg)
	spew.Dump(sgErr)
}
