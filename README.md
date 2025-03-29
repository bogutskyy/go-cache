Go Cache
================================

Go Cache is an in-memory cache to keep any type values.

See it in action:

## Example

```go
package main

import (
	"fmt"
	"github.com/bogutskyy/go-cache"
)

func main() {
	c := cache.New()

	c.Set("a", "a")

	if "a" == c.Get("a") {
		fmt.Println("the value is cached")
	}

	c.Delete("a")

	if nil == c.Get("a") {
		fmt.Println("the value is deleted from the cache")
	}
}
```