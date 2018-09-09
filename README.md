# Seq

[![Godoc Reference][godoc-img]][godoc]

Package for encoding/decoding data sequences.

## Usage
``` $ go get github.com/epiqm/seq ```

### Example

Marshal object to JSON text and unmarshal from JSON text string to custom object.

```go
// main.go
package main

import (
        "fmt"

        "github.com/epiqm/seq"
)

func main() {
        // Get MD5 hash from text string
        hash := seq.Hash("text string")

        fmt.Println(hash)
        // a278c7ab35780d23f34c75dd23278b4b
}
```
#### Other examples

##### Get MD5 hash from text string

```go
hash := seq.Hash("text string")

fmt.Println(hash)
// a278c7ab35780d23f34c75dd23278b4b
```
##### Marshal object to JSON text

```go
obj := map[string]interface{}{
        "username": "Max",
        "age":      27,
        "hometown": "Kiev",
}
json := seq.Marshal(obj)

fmt.Println(json)
// Output:
// {"age":27,"hometown":"Kiev","username":"Max"}
```

##### Unmarshal JSON text to Go object

```go
// Author example object for unmarshal
type Author struct {
        Username string `json:"username"`
        Hometown string `json:"hometown"`
        Age      int    `json:"age"`
}

var newObj Author
seq.Unmarshal(json, &newObj)

fmt.Println(newObj.Username)
// Output:
// Max
```

##### Base64 encode

```go
key := ""
text := "A message for encoding."

s, err := seq.Encode(text, key)
if err != nil {
        fmt.Println(err)
        return
}

fmt.Println(s)
// Output:
// QSBtZXNzYWdlIGZvciBlbmNvZGluZy4=
```

##### Base64 decode

```go
key := ""
text := "QSBtZXNzYWdlIGZvciBlbmNvZGluZy4="

s, err := seq.Decode(text, key)
if err != nil {
        fmt.Println(err)
        return
}

fmt.Println(s)
// Output:
// A message for encoding.
```

##### Create a file

```go
err := seq.CreateFile("testing.txt")
if err != nil {
        fmt.Println(err)
        return
}
```

##### Write a file

```go
err := seq.WriteFile("testing.txt", "A testing text file.")
if err != nil {
        fmt.Println(err)
        return
}
```

##### Read a file

```go
text, err := seq.ReadFile("testing.txt")
if err != nil {
        fmt.Println(err)
        return
}

fmt.Println(text)
// Output:
// A testing text file.
```

##### Copy a file

```go
err := seq.CopyFile("testing.txt", "testing.txt.copy")
if err != nil {
        fmt.Println(err)
        return
}
```

##### Move file

```go
err := seq.MoveFile("testing.txt.copy", "testing.txt")
if err != nil {
        fmt.Println(err)
        return
}
```

##### Base64 encode file contents and resave

```go
err := seq.EncodeFile("testing.txt", "")
if err != nil {
        fmt.Println(err)
        return
}
```

##### Base64 decode contents of a file and resave

```go
err := seq.DecodeFile("testing.txt", "")
if err != nil {
        fmt.Println(err)
        return
}
```

##### Delete file

```go
err := seq.RmFile("testing.txt")
if err != nil {
        fmt.Println(err)
        return
}
```

##### Create a directory

```go
err := seq.CreateDir("newdir")
if err != nil {
        fmt.Println(err)
        return
}
```

##### List directory files

```go
files, err := seq.Ls("newdir")
if err != nil {
        fmt.Println(err)
        return
}
```

##### Remove directory (recursively, removes subfolders/files)

```go
err := seq.RmDir("newdir")
if err != nil {
        fmt.Println(err)
        return
}
```

## Credits

Copyright (c) 2018 Maxim R. All rights reserved.

For feedback or questions &lt;epiqmax@gmail.com&gt;


[godoc]: http://godoc.org/github.com/epiqm/seq
[godoc-img]: https://godoc.org/github.com/epiqm/seq?status.svg
