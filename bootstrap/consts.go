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

package bootstrap

import (
	"github.com/polarismesh/polaris-sidecar/log"
	"os"
	"strconv"
)

const (
	EnvSidecarBind                     = "SIDECAR_BIND"
	EnvSidecarPort                     = "SIDECAR_PORT"
	EnvSidecarNamespace                = "SIDECAR_NAMESPACE"
	EnvSidecarRecurseEnable            = "SIDECAR_RECURSE_ENABLE"
	EnvSidecarRecurseTimeout           = "SIDECAR_RECURSE_TIMEOUT"
	EnvSidecarLogRotateOutputPath      = "SIDECAR_LOG_ROTATE_OUTPUT_PATH"
	EnvSidecarLogErrorRotateOutputPath = "SIDECAR_LOG_ERROR_ROTATE_OUTPUT_PATH"
	EnvSidecarLogRotationMaxSize       = "SIDECAR_LOG_ROTATION_MAX_SIZE"
	EnvSidecarLogRotationMaxBackups    = "SIDECAR_LOG_ROTATION_MAX_BACKUPS"
	EnvSidecarLogRotationMaxAge        = "SIDECAR_LOG_ROTATION_MAX_AGE"
	EnvSidecarLogLevel                 = "SIDECAR_LOG_LEVEL"
	EnvSidecarDnsTtl                   = "SIDECAR_DNS_TTL"
	EnvSidecarDnsEnable                = "SIDECAR_DNS_ENABLE"
	EnvSidecarDnsSuffix                = "SIDECAR_DNS_SUFFIX"
	EnvSidecarDnsRouteLabels           = "SIDECAR_DNS_ROUTE_LABELS"
	EnvSidecarMeshTtl                  = "SIDECAR_MESH_TTL"
	EnvSidecarMeshEnable               = "SIDECAR_MESH_ENABLE"
	EnvSidecarMeshReloadInterval       = "SIDECAR_MESH_RELOAD_INTERVAL"
	EnvSidecarMeshAnswerIp             = "SIDECAR_MESH_ANSWER_IP"
	EnvSidecarMtlsEnable               = "SIDECAR_MTLS_ENABLE"
	EnvSidecarMtlsCAServer             = "SIDECAR_MTLS_CA_SERVER"
)

func getEnvStringValue(envName string, defaultValue string) string {
	envValue := os.Getenv(envName)
	if len(envValue) > 0 {
		return envValue
	}
	return defaultValue
}

func getEnvIntValue(envName string, defaultValue int) int {
	envValue := os.Getenv(envName)
	if len(envValue) > 0 {
		intValue, err := strconv.Atoi(envValue)
		if nil != err {
			log.Errorf("[agent] fail to parse env %s, value %s to int, err: %v", envName, envValue, err)
			return defaultValue
		}
		return intValue
	}
	return defaultValue
}

func getEnvBoolValue(envName string, defaultValue bool) bool {
	envValue := os.Getenv(envName)
	if len(envValue) > 0 {
		boolValue, err := strconv.ParseBool(envValue)
		if nil != err {
			log.Errorf("[agent] fail to parse env %s, value %s to bool, err: %v", envName, envValue, err)
			return defaultValue
		}
		return boolValue
	}
	return defaultValue
}
