package callbacklib

import (
	log "github.com/link-yundi/ylog"
	"testing"
)

/**
------------------------------------------------
Created on 2022-10-28 09:42
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

func TestCallback(t *testing.T) {
	EmptyTopic := "empty"
	IntTopic := "int"
	bus := NewCallbackBus()
	// 空回调
	emptyCallback := NewCallback("empty", MaxCallbackPriority, emptyHandler)
	intCallback1 := NewCallback("int", MaxCallbackPriority, intHandler1)
	intCallback2 := NewCallback("int", MaxCallbackPriority-1, intHandler2)
	RegisterCallback(bus, emptyCallback)
	RegisterCallback(bus, intCallback1)
	RegisterCallback(bus, intCallback2)
	Publish(bus, EmptyTopic, nil)
	Publish(bus, IntTopic, 3)
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
