[![Build Status](https://travis-ci.org/shivakar/xxhash.svg?branch=master)](https://travis-ci.org/shivakar/xxhash) [![Coverage Status](https://coveralls.io/repos/github/shivakar/xxhash/badge.svg?branch=master)](https://coveralls.io/github/shivakar/xxhash?branch=master) [![GoDoc](https://godoc.org/github.com/shivakar/xxhash?status.svg)](https://godoc.org/github.com/shivakar/xxhash)
# xxhash
A pure Go port of xxHash algorithm

For more information about xxHash, see:

* https://github.com/Cyan4973/xxHash
* http://www.xxhash.com/

## Installation

```bash
go get github.com/shivakar/xxhash
```

## Usage

```
package main

import (
    "fmt"

    "github.com/shivakar/xxhash"
)

func main() {
    // Create a new instance of the hash engine with default seed
    h := xxhash.NewXXHash64()

    // Create a new instance of the hash engine with custom seed
    _ = xxhash.NewSeedXXHash64(uint64(10))

    // Write some data to the hash
    h.Write([]byte("Hello, World!!"))

    // Write some more data to the hash
    h.Write([]byte("How are you doing?"))

    // Get the current hash as a byte array
    b := h.Sum(nil)
    fmt.Println(b)

    // Get the current hash as an integer (uint64) (little-endian)
    fmt.Println(h.Uint64())

    // Get the current hash as a hexadecimal string (big-endian)
    fmt.Println(h.String())

    // Reset the hash
    h.Reset()

    // Output:
    // [70 182 137 152 187 180 209 136]
    // 5095411317493518728
    // 46b68998bbb4d188

}

```

## License

`xxhash` is licensed under a MIT license.

