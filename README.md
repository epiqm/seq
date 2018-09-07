# Seq

[![Godoc Reference][godoc-img]][godoc]

Package for encoding/decoding data sequences.

## Usage
``` $ go get github.com/epiqm/seq ```

### Example
```go
// main.go
package main

import (
        "github.com/epiqm/seq"
)

func main() {
        // Marshal object to JSON
        obj := map[string]interface{}{
                "username": "Max",
                "age":      27,
                "hometown": "Kiev",
        }

        json := seq.Marshal(obj)
        // {"age":27,"hometown":"Kiev","username":"Max"}

        // Unmarshal JSON to object
        type Author struct {
                Username string `json:"username"`
                Hometown string `json:"hometown"`
                Age      int    `json:"age"`
        }
        var newObj Author

        seq.Unmarshal(json, &newObj)
        // {Max Kiev 27}
}
```

[godoc]: http://godoc.org/github.com/ewwwwwqm/jsonresp
[godoc-img]: https://godoc.org/github.com/ewwwwwqm/jsonresp?status.svg
