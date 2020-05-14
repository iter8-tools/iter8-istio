/*
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

package experiment

import (
	iter8v1alpha2 "github.com/iter8-tools/iter8-controller/pkg/apis/iter8/v1alpha2"
)

func (r *ReconcileExperiment) getInterval(instance *iter8v1alpha2.Experiment) (time.Duration, error) {
	if instance.Duration.Interval != nil {
		return time.ParseDuration(*instance.Duration.Interval)
	}

	// Get defaults
	if r.experimentConfig.GetInterval() != nil {
		return time.ParseDuration(*r.experimentConfig.GetInterval())
	}

	return iter8v1alpha2.DefaultInterval, nil
}
