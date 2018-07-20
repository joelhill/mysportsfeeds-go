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
	dailyGameOptions := client.NewCurrentSeasonOptions()
	sg, sgErr := client.CurrentSeason(c, dailyGameOptions)

	spew.Dump(sg)
	spew.Dump(sgErr)
}
