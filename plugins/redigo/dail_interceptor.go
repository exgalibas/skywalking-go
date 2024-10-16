package redigo

import (
	"fmt"
	"github.com/apache/skywalking-go/plugins/core/operator"
	"reflect"
)

type DialInterceptor struct {
}

func (i *DialInterceptor) BeforeInvoke(invocation operator.Invocation) error {
	fmt.Println("dial invoke")
	return nil
}

func (i *DialInterceptor) AfterInvoke(invocation operator.Invocation, result ...interface{}) error {
	data := &DatabaseInfo{
		Network: invocation.Args()[0].(string),
		Addr:    invocation.Args()[1].(string),
	}
	fmt.Printf("%v\n", data)
	if caller, ok := result[0].(operator.EnhancedInstance); ok && caller != nil && !reflect.ValueOf(caller).IsNil() {
		fmt.Println("redis dial SetSkyWalkingDynamicField success")
		caller.SetSkyWalkingDynamicField(data)
	}
	return nil
}
