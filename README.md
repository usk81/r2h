# R2H - Romaji to Hiragana

r2h is a golang package that converts ROMAJI to HIRAGANA.

# Install

```shell
go get github.com/usk81/r2h
```

## Example

```go
package main

import (
	"github.com/usk81/r2h"
)

func main() {
    // romaji
    result, isCompleted := r2h.Convert("konnichiha")
    // result: こんにちは
    // isCompleted: true

    // non-romaji
    result, isCompleted = r2h.Convert("github")
    // result: ぎてゅb
    // isCompleted: false

    // strict: romaji
    result, err := r2h.ConvertStrict("konnichiha")
    // result: こんにちは
    // err: nil

    // non-romaji
    result, err = r2h.ConvertStrict("github")
    // result: (empty)
    // err: b is not romaji
}
```