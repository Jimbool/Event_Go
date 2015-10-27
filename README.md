# Event_Go
用Go语言写的事件项目，为了与第三方Fork过来的相区别，所以加了后缀

使用此项目的方法：
1、先下载此项目：go get github.com/Jordanzuo/Event_Go
2、对需要触发事件的代码，先声明一个Event对象。如：LvUpEvent = Event_Go.New()
3、如果事件需要传递参数，则需要实现IEventParamType接口的一个具体类型：有两种方式可以实现IEventParamType接口，1：实现GetTypeName方法；2：包含DefaultEventParamType属性
4、声明需要被触发的方法，方法的签名为：func(IEventParamType)
5、将需要触发的方法注册到事件对象中，如：LvUpEvent.Register(f1)
6、在需要触发事件时，调用Event对象的Trigger或TriggerAsync方法

详情可以参考：event_test.go文件
