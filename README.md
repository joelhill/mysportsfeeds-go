# go-mysportsfeeds
Go library to use MySportsFeeds API

## Example Use

```
import (
    "context"

    msf "github.com/joelhill/go-mysportsfeeds"
)

ctx := context.Context
config := msf.NewConfig()
config.Authorization = "Basic asfafasdfasdfasdfasasdfsadfasdfsd"
client := msf.NewService(config)
options := msf.DefaultSeasonalGamesOptions()
games := client.SeasonalGames(ctx, options)
```