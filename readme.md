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
	emptyCallback := NewCallback("empty", MaxCallbackPriority, emptyHandler)
	// int回调
    intCallback1 := NewCallback("int", MaxCallbackPriority, intHandler1)
	intCallback2 := NewCallback("int", MaxCallbackPriority-1, intHandler2)
	// 注册
    RegisterCallback(emptyCallback)
	RegisterCallback(intCallback1)
	RegisterCallback(intCallback2)
	// 发布数据
    Publish(EmptyTopic, nil)
	Publish(IntTopic, 3)
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

