package redis

import (
	"fmt"
	"github.com/apache/skywalking-go/plugins/core/operator"
	"github.com/apache/skywalking-go/plugins/core/tracing"
	"github.com/gomodule/redigo/redis"
	"strings"
)

type DoInterceptor struct {
}

func (i *DoInterceptor) BeforeInvoke(invocation operator.Invocation) error {
	peer := "unknown"
	switch caller := invocation.CallerInstance().(type) {
	case *nativeActiveConn:
		if caller != nil && caller.pc != nil && caller.pc.c != nil {
			if co, ok := caller.pc.c.(operator.EnhancedInstance); ok && co != nil {
				if dbInfo, ok := co.GetSkyWalkingDynamicField().(*DatabaseInfo); ok && dbInfo != nil {
					peer = dbInfo.Peer()
				}
			}
		}
	case *nativeErrConn:
	default:
		return fmt.Errorf("unknown caller instance")
	}
	op := invocation.Args()[0].(string)
	s, err := tracing.CreateExitSpan(op, peer, func(k, v string) error {
		return nil
	}, tracing.WithComponent(7),
		tracing.WithLayer(tracing.SpanLayerCache),
		tracing.WithTag(tracing.TagCacheType, "redis"))

	if err != nil {
		return err
	}
	statement := op
	for _, arg := range invocation.Args()[1].([]interface{}) {
		statement += fmt.Sprintf(" %v", arg)
	}
	s.Tag("statement", statement)
	invocation.SetContext(s)
	return nil
}

func (i *DoInterceptor) AfterInvoke(invocation operator.Invocation, results ...interface{}) error {
	if invocation.GetContext() == nil {
		return nil
	}
	span := invocation.GetContext().(tracing.Span)
	span.Tag("result", fmt.Sprintf("%v", results[0]))
	if err, ok := results[1].(error); ok && err != nil {
		if err == redis.ErrNil || strings.Contains(err.Error(), "context canceled") {
			span.Log(err.Error())
		} else {
			span.Error(err.Error())
		}
	}
	span.End()
	return nil
}
