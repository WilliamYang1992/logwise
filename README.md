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

import "github.com/WilliamYang1992/logwise"

func main() {
	logwise.SetLogLevel(logwise.InfoLevel)
	l := logwise.Default()
	l.Prefix = "Wow"
	l.Infoln("some text")
}
```
It will output like this: `Wow [INFO] 2019/05/07 21:53:36 main.go:9:
some text`

#### Roadmap
1. Support colored output.

#### License
MIT License