/*
Copyright 2016 The Kubernetes Authors.

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

package util

import (
	"testing"
)

func Test_ParseKubernetesVersion(t *testing.T) {
	grid := map[string]string{
		"1.3.7":         "1.3.7",
		"v1.4.0-beta.8": "1.4.0-beta.8",
		"1.5.0":         "1.5.0",
		"https://storage.googleapis.com/kubernetes-release-dev/ci/v1.4.0-alpha.2.677+ea69570f61af8e/": "1.4.0",
	}
	for version, expected := range grid {
		sv, err := ParseKubernetesVersion(version)
		if err != nil {
			t.Errorf("ParseKubernetesVersion error parsing %q: %v", version, err)
		}

		actual := sv.String()
		if actual != expected {
			t.Errorf("version mismatch: %q -> %q but expected %q", version, actual, expected)
		}
	}

}
