// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package metric

// NewCounter creates a new counter metrics.
func NewCounter(name string, opt ...MeterOpt) *CounterRef {
	return &CounterRef{}
}

// NewGauge creates a new gauge metrics.
func NewGauge(name string, getter func() float64, opts ...MeterOpt) *GaugeRef {
	return &GaugeRef{}
}

// NewHistogram creates a new histogram metrics.
func NewHistogram(name string, steps []float64, opts ...MeterOpt) *HistogramRef {
	return &HistogramRef{}
}

// MeterOpt Defines common options apply for Meter.
// This is implemented in core/metrics package and converted in interceptors.
type MeterOpt interface {
}

// WithLabels Add labels for metric
func WithLabels(key, val string) MeterOpt {
	return nil
}
