/*
Copyright 2019 kubeflow.org.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha2

import (
	"fmt"
	"strings"

	"github.com/kubeflow/kfserving/pkg/constants"
	v1 "k8s.io/api/core/v1"
)

var (
	TorchServeServingGPUSuffix          = "-gpu"
	InvalidTorchServeRuntimeIncludesGPU = "TorchServe RuntimeVersion is not GPU enabled but GPU resources are requested. "
	InvalidTorchServeRuntimeExcludesGPU = "TorchServe RuntimeVersion is GPU enabled but GPU resources are not requested. "
)

var _ Predictor = (*TorchServeSpec)(nil)

func (p *TorchServeSpec) GetStorageUri() string {
	return p.StorageURI
}

func (p *TorchServeSpec) GetResourceRequirements() *v1.ResourceRequirements {
	// return the ResourceRequirements value if set on the spec
	return &p.Resources
}

func (p *TorchServeSpec) GetContainer(modelName string, parallelism int, config *InferenceServicesConfig) *v1.Container {
	arguments := []string{
		"torchserve",
		"--start",
		fmt.Sprintf("%s=%s", "--model-store", constants.DefaultModelLocalMountPath+"/model-store"),
		fmt.Sprintf("%s=%s", "--ts-config", constants.DefaultModelLocalMountPath+"/config/config.properties"),
	}
	// if isGPUEnabled(p.Resources) {
	// 	arguments = append(arguments, fmt.Sprintf("%s=%s", constants.ArgumentWorkers, "1"))
	// } else if parallelism != 0 {
	// 	arguments = append(arguments, fmt.Sprintf("%s=%s", constants.ArgumentWorkers, strconv.Itoa(parallelism)))
	// }
	return &v1.Container{
		Image:     config.Predictors.TorchServe.ContainerImage + ":" + p.RuntimeVersion,
		Name:      constants.InferenceServiceContainerName,
		Resources: p.Resources,
		Args:      arguments,
	}
}

func (p *TorchServeSpec) ApplyDefaults(config *InferenceServicesConfig) {
	if p.RuntimeVersion == "" {
		if isGPUEnabled(p.Resources) {
			p.RuntimeVersion = config.Predictors.TorchServe.DefaultGpuImageVersion
		} else {
			p.RuntimeVersion = config.Predictors.TorchServe.DefaultImageVersion
		}
	}
	setResourceRequirementDefaults(&p.Resources)
}

func (p *TorchServeSpec) Validate(config *InferenceServicesConfig) error {

	if isGPUEnabled(p.Resources) && !strings.Contains(p.RuntimeVersion, TorchServeServingGPUSuffix) {
		return fmt.Errorf(InvalidTorchServeRuntimeIncludesGPU)
	}

	if !isGPUEnabled(p.Resources) && strings.Contains(p.RuntimeVersion, TorchServeServingGPUSuffix) {
		return fmt.Errorf(InvalidTorchServeRuntimeExcludesGPU)
	}
	return nil
}
