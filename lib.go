package callbacklib

import (
	"errors"
	log "github.com/link-yundi/ylog"
	"sort"
	"sync"
)

/**
------------------------------------------------
Created on 2022-10-28 09:00
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

/**
回调包：管理发布回调
*/

type (
	CallbackPriority uint8     // 回调等级，同一主题多个回调中，按照n~1的顺序依次触发
	Handler          func(any) // 带数据的回调
)

const (
	MaxCallbackPriority CallbackPriority = 99
	MinCallbackPriority CallbackPriority = 1
	ErrPriority                          = "错误的回调级别: 范围1~99"
)

type callback struct {
	topic    string
	priority CallbackPriority
	handler  Handler
}

func NewCallback(topic string, priority CallbackPriority, handler Handler) *callback {
	if priority > MaxCallbackPriority || priority < MinCallbackPriority {
		err := errors.New(ErrPriority)
		log.Error(err)
		return nil
	}
	return &callback{
		topic:    topic,
		priority: priority,
		handler:  handler,
	}
}

// ========================== 回调中枢 ==========================
type callbackBus struct {
	mapTopic *sync.Map
}

func NewCallbackBus() *callbackBus {
	return &callbackBus{
		mapTopic: &sync.Map{},
	}
}

func (cbb *callbackBus) callbacks(topic string) []*callback {
	if cbs, ok := cbb.mapTopic.Load(topic); ok {
		return cbs.([]*callback)
	}
	return make([]*callback, 0)
}

// 排序callback
func (cbb *callbackBus) sortCallbacks(topic string, cbs []*callback) {
	sort.SliceStable(cbs, func(i, j int) bool {
		callbackI, callbackJ := cbs[i], cbs[j]
		return callbackI.priority >= callbackJ.priority
	})
	cbb.mapTopic.Store(topic, cbs)
}

func RegisterCallback(bus *callbackBus, callback *callback) {
	cbs := bus.callbacks(callback.topic)
	cbs = append(cbs, callback)
	bus.sortCallbacks(callback.topic, cbs)
}

func Publish(bus *callbackBus, topic string, data any) {
	cbs := bus.callbacks(topic)
	for _, cb := range cbs {
		cb.handler(data)
	}
}

func Delete(bus *callbackBus, topic string) {
	bus.mapTopic.Delete(topic)
}
