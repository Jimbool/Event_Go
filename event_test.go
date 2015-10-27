package Event_Go

import (
	"fmt"
	"testing"
	"time"
)

type EventParam struct {
	Name string
	Age  int
}

func NewEventParam(name string, count int) *EventParam {
	return &EventParam{
		Name: name,
		Age:  count,
	}
}

func (param EventParam) GetTypeName() string {
	return "EventParam"
}

var (
	LvUpEvent = New()
	f1Result  string
	f2Result  string
	f3Result  string
)

func f1(param EventParamType) {
	if param_new, ok := param.(*EventParam); ok {
		f1Result = fmt.Sprintf("f1:Name:%s,Age:%d,paranName:%s", param_new.Name, param_new.Age, param_new.GetTypeName())
	}
}

func f2(param EventParamType) {
	if param_new, ok := param.(*EventParam); ok {
		f2Result = fmt.Sprintf("f2:Name:%s,Age:%d,paranName:%s", param_new.Name, param_new.Age, param_new.GetTypeName())
	}
}

func f3(param EventParamType) {
	if param_new, ok := param.(*EventParam); ok {
		f3Result = fmt.Sprintf("f3:Name:%s,Age:%d,paranName:%s", param_new.Name, param_new.Age, param_new.GetTypeName())
	}
}

func init() {
	LvUpEvent.Register(f1)
	LvUpEvent.Register(f2)
	LvUpEvent.UnRegister(f2)
	LvUpEvent.Register(f3)
}

func Test(t *testing.T) {
	eventParam := NewEventParam("Jordan", 30)
	LvUpEvent.Trigger(eventParam)
	if f1Result != "f1:Name:Jordan,Age:30,paranName:EventParam" {
		t.Errorf("f1触发失败")
	}
	if f2Result != "" {
		t.Errorf("f2触发失败")
	}
	if f3Result != "f3:Name:Jordan,Age:30,paranName:EventParam" {
		t.Errorf("f3触发失败")
	}
	t.Logf("f1Result:%s\n", f1Result)
	t.Logf("f2Result:%s\n", f2Result)
	t.Logf("f3Result:%s\n", f3Result)
}

func TestAsync(t *testing.T) {
	eventParam := NewEventParam("Jordan", 30)
	LvUpEvent.TriggerAsync(eventParam)
	time.Sleep(5 * time.Second)
	if f1Result != "f1:Name:Jordan,Age:30,paranName:EventParam" {
		t.Errorf("f1触发失败")
	}
	if f2Result != "" {
		t.Errorf("f2触发失败")
	}
	if f3Result != "f3:Name:Jordan,Age:30,paranName:EventParam" {
		t.Errorf("f3触发失败")
	}
	t.Logf("f1Result:%s\n", f1Result)
	t.Logf("f2Result:%s\n", f2Result)
	t.Logf("f3Result:%s\n", f3Result)
}
