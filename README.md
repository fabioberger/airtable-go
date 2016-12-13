<a href="https://godoc.org/github.com/fabioberger/airtable-go" ><img src="http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square" /></a>

Airtable Go Client Library
-------------------------------

### Installation

Make sure you have Golang v1.6 or higher installed. If not, <a href="https://golang.org/dl/">install it now</a>.

Fetch airtable-go:

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
baseID := "apphllLCpWnySSF7q" // replace this with your airtable base's id
shouldRetryIfRateLimited := true

client := airtable.New(airtableAPIKey, baseID, shouldRetryIfRateLimited)
```
You can now call many methods on the client instance. These are documented in the <a href="https://godoc.org/github.com/fabioberger/airtable-go">Airtable-go GoDoc page</a>. In addition to this, check out the <a href="https://github.com/fabioberger/airtable-go/blob/master/client_test.go">stubbed tests</a> and <a href="https://github.com/fabioberger/airtable-go/blob/master/integration_tests/client_test.go">integration tests</a> included in this project, they contain working examples of all the client methods available.
