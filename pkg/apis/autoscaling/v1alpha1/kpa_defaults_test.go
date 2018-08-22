/*
Copyright 2018 The Knative Authors

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

package v1alpha1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPodAutoscalerDefaulting(t *testing.T) {
	tests := []struct {
		name string
		in   *PodAutoscaler
		want *PodAutoscaler
	}{{
		name: "empty",
		in:   &PodAutoscaler{},
		want: &PodAutoscaler{
			Spec: PodAutoscalerSpec{
				// In the context of a PodAutoscaler we initialize ServingState.
				ContainerConcurrency: 0,
				ServingState:         "Active",
			},
		},
	}, {
		name: "no overwrite",
		in: &PodAutoscaler{
			Spec: PodAutoscalerSpec{
				ContainerConcurrency: 1,
				ServingState:         "Reserve",
			},
		},
		want: &PodAutoscaler{
			Spec: PodAutoscalerSpec{
				ContainerConcurrency: 1,
				ServingState:         "Reserve",
			},
		},
	}, {
		name: "partially initialized",
		in: &PodAutoscaler{
			Spec: PodAutoscalerSpec{},
		},
		want: &PodAutoscaler{
			Spec: PodAutoscalerSpec{
				ContainerConcurrency: 0,
				ServingState:         "Active",
			},
		},
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.in
			got.SetDefaults()
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("SetDefaults (-want, +got) = %v", diff)
			}
		})
	}
}
