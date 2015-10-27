package Event_Go

import (
	"reflect"
)

// 定义事件对象
// 任何具体的事件都是它的一个具体对象，通过调用New方法来创建
type Event struct {
	funcList []func(IEventParamType)
}

func New() *Event {
	return &Event{
		funcList: make([]func(IEventParamType), 0),
	}
}

// 注册事件方法
// f：待注册的事件方法
func (listener *Event) Register(f func(IEventParamType)) {
	listener.funcList = append(listener.funcList, f)
}

// 取消方法的注册
// f：待取消注册的方法
func (listener *Event) UnRegister(f func(IEventParamType)) {
	for index, value := range listener.funcList {
		if reflect.ValueOf(value) == reflect.ValueOf(f) {
			// 如果是最后一项，则单独处理
			if index == len(listener.funcList)-1 {
				listener.funcList = listener.funcList[:index]
			} else {
				listener.funcList = append(listener.funcList[:index], listener.funcList[index+1:]...)
			}

			return
		}
	}
}

// 同步触发事件
// param：传递的参数
func (listener *Event) Trigger(param IEventParamType) {
	for _, value := range listener.funcList {
		value(param)
	}
}

// 异步触发事件
// param：传递的参数
func (listener *Event) TriggerAsync(param IEventParamType) {
	for _, value := range listener.funcList {
		go value(param)
	}
}
