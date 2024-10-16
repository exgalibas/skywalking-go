package redigo

import (
	"errors"
	"fmt"
	"github.com/apache/skywalking-go/plugins/core/operator"
	"github.com/apache/skywalking-go/plugins/core/tracing"
	"github.com/gomodule/redigo/redis"
	"reflect"
	"strings"
)

type DoInterceptor struct {
}

func (i *DoInterceptor) BeforeInvoke(invocation operator.Invocation) error {
	if caller, ok := invocation.CallerInstance().(operator.EnhancedInstance); ok && caller != nil && !reflect.ValueOf(caller).IsNil() {
		if data, ok := caller.GetSkyWalkingDynamicField().(*DatabaseInfo); ok && data != nil {
			// 拼接command
			var op string
			if len(invocation.Args()) <= 0 {
				op = "unknown"
			} else if len(invocation.Args()) >= 2 {
				op = fmt.Sprintf("%v %v", invocation.Args()[0], invocation.Args()[1])
			} else {
				op = fmt.Sprintf("%v", invocation.Args()[0])
			}
			// 创建span
			s, err := tracing.CreateExitSpan(op, data.Peer(), func(k, v string) error {
				return nil
			}, tracing.WithComponent(5041),
				tracing.WithLayer(tracing.SpanLayerCache),
				tracing.WithTag(tracing.TagCacheType, "redis"))

			if err != nil {
				return err
			}
			invocation.SetContext(s)
		}
	}
	return nil
}

func (i *DoInterceptor) AfterInvoke(invocation operator.Invocation, results ...interface{}) error {
	if invocation.GetContext() == nil {
		return nil
	}
	span := invocation.GetContext().(tracing.Span)
	span.Tag("result", fmt.Sprintf("%v", results[0]))
	if err, ok := results[1].(error); ok && err != nil {
		if errors.Is(err, redis.ErrNil) || strings.Contains(err.Error(), "context canceled") {
			span.Log(err.Error())
		} else {
			span.Error(err.Error())
		}
	}
	span.End()
	return nil
}
