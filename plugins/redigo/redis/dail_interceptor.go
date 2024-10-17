package redis

import (
	"fmt"
	"github.com/apache/skywalking-go/plugins/core/operator"
)

type DialInterceptor struct {
}

func (i *DialInterceptor) BeforeInvoke(invocation operator.Invocation) error {
	return nil
}

func (i *DialInterceptor) AfterInvoke(invocation operator.Invocation, result ...interface{}) error {
	data := &DatabaseInfo{
		Network: invocation.Args()[0].(string),
		Addr:    invocation.Args()[1].(string),
	}
	if err, ok := result[1].(error); ok && err != nil {
		return err
	}
	if caller, ok := result[0].(operator.EnhancedInstance); ok && caller != nil {
		fmt.Println("redis dial SetSkyWalkingDynamicField success")
		caller.SetSkyWalkingDynamicField(data)
	}
	return nil
}
