/*
 * Copyright The Spinnaker Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package validate_hal

import (
	"testing"

	"github.com/spinnaker/kleat/api/client/cloudprovider"
	"github.com/spinnaker/kleat/api/client/config"
)

func TestEmptyHalConfig(t *testing.T) {
	h := &config.Hal{}
	err := ValidateHalConfig(h)
	if err != nil {
		t.Errorf("Expected no validation failures, got %s", err)
	}
}

func TestEmptyProviders(t *testing.T) {
	h := &config.Hal{
		Providers: &config.Hal_Providers{},
	}
	err := ValidateHalConfig(h)
	if err != nil {
		t.Errorf("Expected no validation failures, got %s", err)
	}
}

func TestEmptyKubernetes(t *testing.T) {
	h := &config.Hal{
		Providers: &config.Hal_Providers{
			Kubernetes: &cloudprovider.KubernetesProvider{},
		},
	}
	err := ValidateHalConfig(h)
	if err != nil {
		t.Errorf("Expected no validation failures, got %s", err)
	}
}

func TestNoKubernetesAccounts(t *testing.T) {
	h := &config.Hal{
		Providers: &config.Hal_Providers{
			Kubernetes: &cloudprovider.KubernetesProvider{
				Enabled:  false,
				Accounts: []*cloudprovider.KubernetesAccount{},
			},
		},
	}
	err := ValidateHalConfig(h)
	if err != nil {
		t.Errorf("Expected no validation failures, got %s", err)
	}
}

func TestKuberntesAccountWithNoOmitKinds(t *testing.T) {
	h := &config.Hal{
		Providers: &config.Hal_Providers{
			Kubernetes: &cloudprovider.KubernetesProvider{
				Accounts: []*cloudprovider.KubernetesAccount{
					{
						Name:  "my-account",
						Kinds: []string{"deployment"},
					},
				},
			},
		},
	}
	err := ValidateHalConfig(h)
	if err != nil {
		t.Errorf("Expected no validation failures, got %s", err)
	}
}

func TestKuberntesAccountWithEmptyOmitKinds(t *testing.T) {
	h := &config.Hal{
		Providers: &config.Hal_Providers{
			Kubernetes: &cloudprovider.KubernetesProvider{
				Accounts: []*cloudprovider.KubernetesAccount{
					{
						Name:      "my-account",
						Kinds:     []string{"deployment"},
						OmitKinds: []string{},
					},
				},
			},
		},
	}
	err := ValidateHalConfig(h)
	if err != nil {
		t.Errorf("Expected no validation failures, got %s", err)
	}
}

func TestInvalidKubernetesAccount(t *testing.T) {
	h := &config.Hal{
		Providers: &config.Hal_Providers{
			Kubernetes: &cloudprovider.KubernetesProvider{
				Accounts: []*cloudprovider.KubernetesAccount{
					{
						Name:      "my-account",
						Kinds:     []string{"deployment"},
						OmitKinds: []string{"replicaSet"},
					},
				},
			},
		},
	}
	err := ValidateHalConfig(h)
	if err == nil {
		t.Error("Expected a validation failure, got none")
	}
}
