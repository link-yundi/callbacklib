# CallbackLib

### 安装

```sh
go get -u github.com/link-yundi/callbacklib
```

### 示例

```go
import (
    "github.com/link-yundi/callbacklib"
    log "github.com/link-yundi/ylog"
)


func main() {
	EmptyTopic := "empty"
	IntTopic := "int"
	// 空回调
	emptyCallback := callbacklib.NewCallback("empty", callbacklib.MaxCallbackPriority, emptyHandler)
	// int回调
    intCallback1 := callbacklib.NewCallback("int", callbacklib.MaxCallbackPriority, intHandler1)
	intCallback2 := callbacklib.NewCallback("int", callbacklib.MaxCallbackPriority-1, intHandler2)
	// 注册
    callbacklib.RegisterCallback(emptyCallback)
	callbacklib.RegisterCallback(intCallback1)
	callbacklib.RegisterCallback(intCallback2)
	// 发布数据
    callbacklib.Publish(EmptyTopic, nil)
	callbacklib.Publish(IntTopic, 3)
}

func emptyHandler(d any) {
	log.Info("我是空回调")
}

func intHandler1(d any) {
	log.Info("int 回调1", d.(int))
}

func intHandler2(d any) {
	log.Info("int 回调2", d.(int))
}

```

