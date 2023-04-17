# mysportsfeeds-go
Go library to use MySportsFeeds API

## Example Use

```
import (
    msf "github.com/joelhill/mysportsfeeds-go"
)

authorization := "Basic asfafasdfasdfasdfasasdfsadfasdfsd"
config := msf.NewConfig(authorization)
client := msf.NewService(config)
options := client.NewSeasonalGamesOptions()
games, statusCode, err := client.SeasonalGames(options)
if err != nil {
    log.Println("something went wrong")
}
```