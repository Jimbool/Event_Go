package Event_Go

// 定义事件参数类型的接口
// 具体的事件需要用到的参数需要实现此接口
// 并且注册到事件的方法需要使用此接口作为参数类型
type IEventParamType interface {
	GetTypeName() string
}

// 定义默认的事件参数类型
type DefaultEventParamType struct {
}

// 实现接口定义的方法
func (param DefaultEventParamType) GetTypeName() string {
	return "EventParam"
}
