# logwise
logwise is a smart log tool for **go** and it is convenient and easy to
use.

#### Feature
1. Easy to use.
2. Support log level filtering.
3. Nearly no performance loss.
4. Concurrent safe.

#### Usage
```go
package main

import (
	"github.com/WilliamYang1992/logwise"
	"github.com/WilliamYang1992/logwise/loglevel"
)

func main() {
	l := logwise.Default()
	l.SetLogLevel(loglevel.Info)
	l.SetPrefix("Wow")
	l.Infoln("some text")
}
```
It will output like this: `Wow [INFO] 2019/05/07 21:53:36 main.go:12:
some text`

#### Roadmap
1. Support colored output.

#### License
**MIT License**
