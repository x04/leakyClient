# LeakyClient

An HTTP client for Go that uses [uber-go/ratelimit](https://github.com/uber-go/ratelimit)
to implement a leaky-bucket rate-limited client.

## Usage

```go
package main

import (
    "fmt"
    
	"github.com/x04/leakyClient"
)

func main() {
    client := leakyClient.New(10) // 10 is the maximum requests per second
                                  // the client will handle

    resp, err := client.Get("https://httpbin.org/get")
    if err != nil {
        panic(err)
    }
    resp.Body.Close()
    fmt.Println(resp.StatusCode)
}
```
