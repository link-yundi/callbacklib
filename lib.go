package callbacklib

import (
	"errors"
	log "github.com/link-yundi/ylog"
	"sort"
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
var mapTopic = map[string][]*callback{}

func has(topic string) bool {
	if _, ok := mapTopic[topic]; ok {
		return true
	}
	return false
}

// 排序callback
func sortCallback(topic string) {
	if has(topic) {
		sort.SliceStable(mapTopic[topic], func(i, j int) bool {
			callbackI, callbackJ := mapTopic[topic][i], mapTopic[topic][j]
			return callbackI.priority >= callbackJ.priority
		})
	}

}

func RegisterCallback(callback *callback) {
	if has(callback.topic) {
		msg := callback.topic + "回调已存在,再次注册将会按照级别依次触发"
		log.Warn(msg)
	}
	mapTopic[callback.topic] = append(mapTopic[callback.topic], callback)
	sortCallback(callback.topic)
}

func Publish(topic string, data any) {
	if has(topic) {
		cbList := mapTopic[topic]
		for _, cb := range cbList {
			cb.handler(data)
		}
	}
}
