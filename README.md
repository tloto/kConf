Getting Started
===============

## Installing

To start using kConf, install Go and run `go get`:

```sh
$ go get github.com/tloto/kConf
```

## Get a value

```go
package main

import "github.com/tloto/kConf"

func init(){
    kConf.SetFiePath("./test.json")
}

func main() {
	str := kConf.GetConfString("compilerOptions.module")
	fmt.println(str)
}
```

This will print:

```
test
```
