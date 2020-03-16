# Logger

# Usage
```go
package main

import "github.com/gosched/logger"

func main() {
	logger.Init("./log.txt")
	logger.Info.Println("hello")
	logger.Error.Println("error")
	logger.Close()
}
```