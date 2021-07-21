# support/zlog

`support/zlog` 是一个生产可用的日志包，基于 `zap` 包封装。兼容zap的全部功能之外，添加了以下额外功能。
- 默认注入了requestID
- 提供了额外的TimeEncoder
- 内置生产环境，和开发环境的日志记录器工厂，使用者可以在不了解zap配置的情况下直接使用。


### 一个简单的示例

```go
package main

import (
    "git.cai-inc.com/support/zlog"
)

func main() {
	loggerProduction:=zlog.NewZLogger(false,false)
	loggerProduction.Info("info")
	loggerProduction.Debug("debug")

	loggerDevelop:=zlog.NewZLogger(true,false)
	loggerDevelop.Info("info")
	loggerDevelop.Debug("debug")
}
```

执行代码：

```bash
2021-07-21 14:07:29.7212	info	zlog/logger_test.go:18	info	{"RequestID": "91f9c739-fb7f-4536-a91f-3882b038381b"}
2021-07-21 14:07:29.7212	debug	zlog/logger_test.go:19	debug	{"RequestID": "91f9c739-fb7f-4536-a91f-3882b038381b"}
```


