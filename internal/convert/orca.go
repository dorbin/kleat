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

package convert

import "github.com/spinnaker/kleat/api/client/config"

// HalToOrca generates the orca config for the supplied config.Hal h.
func HalToOrca(h *config.Hal) *config.Orca {
	return &config.Orca{
		Default:           getDefaults(h),
		PipelineTemplates: getPipelineTemplates(h),
		Webhook:           h.GetWebhook(),
		Services:          getOrcaServices(h),
		Tasks:             getOrcaTasks(h),
	}
}

func getDefaults(h *config.Hal) *config.Orca_Defaults {
	if !h.GetProviders().GetAws().GetEnabled() {
		return nil
	}
	return &config.Orca_Defaults{
		Bake: &config.Orca_Defaults_BakeDefaults{
			Account: h.GetProviders().GetAws().GetPrimaryAccount(),
		},
	}
}

func getPipelineTemplates(h *config.Hal) *config.Orca_PipelineTemplates {
	// TODO: Consider whether this should actually be a BoolWrapper so that we
	// can distinguish an explicitly set 'false' from an unset value
	if !h.GetFeatures().GetPipelineTemplates() {
		return nil
	}
	return &config.Orca_PipelineTemplates{Enabled: h.GetFeatures().GetPipelineTemplates()}
}

func getOrcaServices(h *config.Hal) *config.Orca_Services {
	if h.GetCanary().GetEnabled() {
		return &config.Orca_Services{
			Kayenta: &config.ServiceSettings{
				Enabled: h.GetCanary().GetEnabled(),
			},
		}
	}
	return nil
}

func getOrcaTasks(h *config.Hal) *config.Orca_Tasks {
	if h.GetTimezone() == "" {
		return nil
	}

	return &config.Orca_Tasks{
		ExecutionWindow: &config.Orca_Tasks_ExecutionWindow{
			Timezone: h.GetTimezone(),
		},
	}
}