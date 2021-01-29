# xonfig

Load configuration from environment variables. Simple, strict.

## Example

```Go
package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nixberg/xonfig-go"
)

func main() {
	var config struct {
		DatabasePath  string `env:"DATABASE_PATH"`
		GinMode       string `env:"GIN_MODE"`
		ListenAddress string `env:"LISTEN_ADDRESS"`
	}
	xonfig.MustLoad(&config)

	db, err := sql.Open("sqlite3", config.DatabasePath)

	gin.SetMode(config.GinMode)
	router := gin.Default()

	router.Run(config.ListenAddress)
}
``` 