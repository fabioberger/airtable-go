Airtable Go Client Library
-------------------------------

### Installation

```
go get github.com/fabioberger/airtable-go
```
Import Airtable-go into your project:

```go
import "github.com/fabioberger/airtable-go"
```

### Usage

Create an instance of the Airtable-go client:

```go
import (
	"os"
	"github.com/fabioberger/airtable-go"
)

airtableAPIKey := os.Getenv("AIRTABLE_API_KEY")
baseID := "apphllLCpWnySSF7q" // replace this with your baseID
shouldRetryIfRateLimited := true

client := airtable.New(airtableAPIKey, baseID, shouldRetryIfRateLimited)
```

Detailed documentation and examples on how to use this library are available on the <a href="">Airtable-go GoDoc page</a>. In addition to this, check out the unit and integration tests, they provide working examples of all the client methods available.
