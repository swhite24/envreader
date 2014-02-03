# envreader

**envreader** provides a small wrapper for `os.GetEnv` for retrieving a collection of environment variables into a map.

## Install
```bash
$ go get github.com/swhite24/envreader
```

## Example

```go
package main
import (
	"fmt"
	"github.com/swhite24/envreader"
)

func main () {
	myEnv := envreader.Read("PATH", "GOPATH")

	fmt.Printf("PATH: %s\n", myEnv["PATH"])
	fmt.Printf("GOPATH: %s\n", myEnv["GOPATH"])
}
```