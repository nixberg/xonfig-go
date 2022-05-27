# xonfig-go v2

Load configuration from either `CONFIG` environment variable or `config.toml`.
Simple, strict.

## Example

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nixberg/xonfig-go/v2"
)

func main() {
	config := xonfig.MustLoad[struct {
		GinMode         string
		ListenAddress   string
		TrustedPlatform string

		Accounts gin.Accounts
	}]()

	...
}
```

Contents of `./config.toml`:

```toml
GinMode = "debug"
ListenAddress = "localhost:8080"
TrustedPlatform = ""

Accounts = {}
```

Contents of `CONFIG` environment variable:

```toml
GinMode = "release"
ListenAddress = "0.0.0.0:8080"
TrustedPlatform = "CF-Connecting-IP"

[Accounts]
admin = "8mwf9mrtbu2z2zhbec7qg6kc63"
```
