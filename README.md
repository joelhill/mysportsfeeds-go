# go-mysportsfeeds
Go library to use MySportsFeeds API

## Example Use

```
import (
    "context"

    msf "github.com/joelhill/go-mysportsfeeds"
)

ctx := context.Context
authorization := "Basic asfafasdfasdfasdfasasdfsadfasdfsd"
config := msf.NewConfig(authorization)
client := msf.NewService(config)
options := client.NewSeasonalGamesOptions()
games := client.SeasonalGames(ctx, options)
```