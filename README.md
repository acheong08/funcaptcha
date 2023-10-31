# Arkose Fetch

Usage for OpenAI

```go
import (
	"fmt"

	"github.com/acheong08/funcaptcha"
)

func main() {
	version := 4 // 0 - Auth, 3 - 3.5, 4 - 4
	token, _ := funcaptcha.GetOpenAIToken(version, "", "")
	fmt.Println(token)
}
```

## API:
You can download the binary from releases or `go run cmd/api/main.go`
