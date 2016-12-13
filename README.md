<a href="https://godoc.org/github.com/fabioberger/airtable-go" ><img src="http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square" /></a>

Airtable Go Client Library
-------------------------------

### Installation

Make sure you have Golang v1.6 or higher installed. If not, <a href="https://golang.org/dl/">install it now</a>.

Fetch airtable-go:

```
go get github.com/fabioberger/airtable-go
```
Import airtable-go into your project:

```go
import "github.com/fabioberger/airtable-go"
```

### Usage

Create an instance of the airtable-go client:

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
You can now call methods on the client instance. All client methods are documented in the project's <a href="https://godoc.org/github.com/fabioberger/airtable-go">GoDoc page</a>. You can also check out the <a href="https://github.com/fabioberger/airtable-go/blob/master/client_test.go">stubbed</a> and <a href="https://github.com/fabioberger/airtable-go/blob/master/integration_tests/client_test.go">integration</a> tests included in this project for working examples of all the client methods and options.
