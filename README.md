## How to Use?

```go
import "github.com/weiliang-ms/gotool"
```

### log

```go
package main

import (
	wacherLog "github.com/weiliang-ms/gotool/log"
)

func main()  {
	wacherLog.Logger = wacherLog.NewWLLogger("watcher")
	// 启动容器重启监听
	wacherLog.Logger.Info("开始监听...")
}
```