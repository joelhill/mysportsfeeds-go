```
package main

import (
	"github.com/davecgh/go-spew/spew"
	sf "github.com/joelhill/mysportsfeeds-go"
)

func main() {
	config := sf.NewConfig("Basic changethisheresothatitworks")
	client := sf.NewService(config)
	dailyGameOptions := client.NewDailyGamesOptions()
	dg, statusCode, dgErr := client.DailyGames(dailyGameOptions)

	spew.Dump(dg)
	spew.Dump(statusCode)
	spew.Dump(dgErr)
}
```