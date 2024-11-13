// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package clickhouse

import (
	drive "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/apache/skywalking-go/plugins/core/operator"
	"gorm.io/driver/clickhouse"
	"strings"
)

type InstanceInterceptor struct {
}

func (i *InstanceInterceptor) BeforeInvoke(invocation operator.Invocation) error {
	return nil
}

func (i *InstanceInterceptor) AfterInvoke(invocation operator.Invocation, result ...interface{}) error {
	if res, ok := result[0].(*clickhouse.Dialector); ok && res != nil && res.Config != nil {
		var data *DatabaseInfo
		if res.Config.DSN != "" {
			data = i.buildDBInfo(res)
		} else if res.Config.Conn != nil {
			if conn, ok := res.Config.Conn.(operator.EnhancedInstance); ok && conn != nil {
				if addr, ok := conn.GetSkyWalkingDynamicField().([]string); ok && len(addr) > 0 {
					data = &DatabaseInfo{
						PeerAddress: strings.Join(addr, "/"),
					}
				}
			}
		}
		if data != nil {
			invocation.DefineReturnValues(&DialWrapper{
				Data:      data,
				Dialector: res,
			})
		}
	}
	return nil
}

func (i *InstanceInterceptor) buildDBInfo(dial *clickhouse.Dialector) *DatabaseInfo {
	cfg, err := drive.ParseDSN(dial.Config.DSN)
	if err != nil {
		// ignore the db info if parse dsn failed
		return nil
	}
	return &DatabaseInfo{PeerAddress: strings.Join(cfg.Addr, `/`)}
}

type DialWrapper struct {
	Data *DatabaseInfo
	*clickhouse.Dialector
}

func (dial *DialWrapper) DataInfo() interface{} {
	return dial.Data
}

type DatabaseInfo struct {
	PeerAddress string
}

func (d *DatabaseInfo) Type() string {
	return "clickhouse"
}

func (d *DatabaseInfo) ComponentID() int32 {
	return 120
}

func (d *DatabaseInfo) Peer() string {
	return d.PeerAddress
}
