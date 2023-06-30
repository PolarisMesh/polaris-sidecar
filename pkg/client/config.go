/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package client

import (
	"time"

	"github.com/polarismesh/polaris-go/pkg/config"
)

type Config struct {
	Addresses          []string `yaml:"addresses"`
	Metrics            *Metrics
	LocationConfigImpl *config.LocationConfigImpl
}

type Metrics struct {
	// port listen for metric message
	Port int `yaml:"port"`
	// Type metrics data report type pull/push
	Type string `yaml:"type"`
	// IP if use pull, need open Prometheus HttpServer
	IP string `yaml:"-"`
	// Interval if use push, need set report interval metrics data to pushgateway
	Interval time.Duration `yaml:"interval"`
	// Address if use push, need set report metrics data to pushgateway server
	Address string `yaml:"address"`
}
