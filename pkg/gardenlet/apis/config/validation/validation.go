// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validation

import (
	"fmt"

	gardencore "github.com/gardener/gardener/pkg/apis/core"
	corevalidation "github.com/gardener/gardener/pkg/apis/core/validation"
	"github.com/gardener/gardener/pkg/gardenlet/apis/config"

	apivalidation "k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidateGardenletConfiguration validates a GardenletConfiguration object.
func ValidateGardenletConfiguration(cfg *config.GardenletConfiguration, fldPath *field.Path, inTemplate bool) field.ErrorList {
	allErrs := field.ErrorList{}

	if cfg.Controllers != nil {
		if cfg.Controllers.BackupEntry != nil {
			allErrs = append(allErrs, ValidateBackupEntryControllerConfiguration(cfg.Controllers.BackupEntry, fldPath.Child("controllers", "backupEntry"))...)
		}
		if cfg.Controllers.Bastion != nil {
			allErrs = append(allErrs, ValidateBastionControllerConfiguration(cfg.Controllers.Bastion, fldPath.Child("controllers", "bastion"))...)
		}
		if cfg.Controllers.Shoot != nil {
			allErrs = append(allErrs, ValidateShootControllerConfiguration(cfg.Controllers.Shoot, fldPath.Child("controllers", "shoot"))...)
		}
		if cfg.Controllers.ManagedSeed != nil {
			allErrs = append(allErrs, ValidateManagedSeedControllerConfiguration(cfg.Controllers.ManagedSeed, fldPath.Child("controllers", "managedSeed"))...)
		}
	}

	if !inTemplate && (cfg.SeedConfig == nil && cfg.SeedSelector == nil) || (cfg.SeedConfig != nil && cfg.SeedSelector != nil) {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("seedConfig"), cfg, "either seed config or seed selector is required"))
	}

	if cfg.SeedConfig != nil {
		allErrs = append(allErrs, corevalidation.ValidateSeedTemplate(&cfg.SeedConfig.SeedTemplate, fldPath.Child("seedConfig"))...)
	}

	serverPath := fldPath.Child("server")
	if cfg.Server != nil {
		if len(cfg.Server.HTTPS.BindAddress) == 0 {
			allErrs = append(allErrs, field.Required(serverPath.Child("https", "bindAddress"), "bind address is required"))
		}
		if cfg.Server.HTTPS.Port == 0 {
			allErrs = append(allErrs, field.Required(serverPath.Child("https", "port"), "port is required"))
		}
	}

	resourcesPath := fldPath.Child("resources")
	if cfg.Resources != nil {
		for resourceName, quantity := range cfg.Resources.Capacity {
			if reservedQuantity, ok := cfg.Resources.Reserved[resourceName]; ok && reservedQuantity.Value() > quantity.Value() {
				allErrs = append(allErrs, field.Invalid(resourcesPath.Child("reserved", string(resourceName)), cfg.Resources.Reserved[resourceName], "reserved must be lower or equal to capacity"))
			}
		}
		for resourceName := range cfg.Resources.Reserved {
			if _, ok := cfg.Resources.Capacity[resourceName]; !ok {
				allErrs = append(allErrs, field.Invalid(resourcesPath.Child("reserved", string(resourceName)), cfg.Resources.Reserved[resourceName], "reserved without capacity"))
			}
		}
	}

	exposureClassHandlersPath := fldPath.Child("exposureClassHandlers")
	for i, handler := range cfg.ExposureClassHandlers {
		handlerPath := exposureClassHandlersPath.Index(i)
		if len(handler.Name) == 0 {
			allErrs = append(allErrs, field.Invalid(handlerPath.Child("name"), handler.Name, "handler name cannot be empty"))
		}
	}

	return allErrs
}

// ValidateGardenletConfigurationUpdate validates a GardenletConfiguration object before an update.
func ValidateGardenletConfigurationUpdate(newCfg, oldCfg *config.GardenletConfiguration, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if newCfg.SeedConfig != nil && oldCfg.SeedConfig != nil {
		allErrs = append(allErrs, corevalidation.ValidateSeedTemplateUpdate(&newCfg.SeedConfig.SeedTemplate, &oldCfg.SeedConfig.SeedTemplate, fldPath.Child("seedConfig"))...)
	}

	return allErrs
}

// ValidateShootControllerConfiguration validates the shoot controller configuration.
func ValidateShootControllerConfiguration(cfg *config.ShootControllerConfiguration, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if cfg.ConcurrentSyncs != nil {
		allErrs = append(allErrs, apivalidation.ValidateNonnegativeField(int64(*cfg.ConcurrentSyncs), fldPath.Child("concurrentSyncs"))...)
	}

	if cfg.ProgressReportPeriod != nil {
		allErrs = append(allErrs, apivalidation.ValidateNonnegativeField(int64(cfg.ProgressReportPeriod.Duration), fldPath.Child("progressReporterPeriod"))...)
	}

	if cfg.RetryDuration != nil {
		allErrs = append(allErrs, apivalidation.ValidateNonnegativeField(int64(cfg.RetryDuration.Duration), fldPath.Child("retryDuration"))...)
	}

	if cfg.SyncPeriod != nil {
		allErrs = append(allErrs, apivalidation.ValidateNonnegativeField(int64(cfg.SyncPeriod.Duration), fldPath.Child("syncPeriod"))...)
	}

	if cfg.DNSEntryTTLSeconds != nil {
		const (
			dnsEntryTTLSecondsMin = 30
			dnsEntryTTLSecondsMax = 600
		)

		if *cfg.DNSEntryTTLSeconds < dnsEntryTTLSecondsMin || *cfg.DNSEntryTTLSeconds > dnsEntryTTLSecondsMax {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("dnsEntryTTLSeconds"), *cfg.DNSEntryTTLSeconds, fmt.Sprintf("must be within [%d,%d]", dnsEntryTTLSecondsMin, dnsEntryTTLSecondsMax)))
		}
	}

	return allErrs
}

// ValidateManagedSeedControllerConfiguration validates the managed seed controller configuration.
func ValidateManagedSeedControllerConfiguration(cfg *config.ManagedSeedControllerConfiguration, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if cfg.ConcurrentSyncs != nil {
		allErrs = append(allErrs, apivalidation.ValidateNonnegativeField(int64(*cfg.ConcurrentSyncs), fldPath.Child("concurrentSyncs"))...)
	}
	if cfg.SyncPeriod != nil {
		allErrs = append(allErrs, apivalidation.ValidateNonnegativeField(int64(cfg.SyncPeriod.Duration), fldPath.Child("syncPeriod"))...)
	}
	if cfg.WaitSyncPeriod != nil {
		allErrs = append(allErrs, apivalidation.ValidateNonnegativeField(int64(cfg.WaitSyncPeriod.Duration), fldPath.Child("waitSyncPeriod"))...)
	}
	if cfg.SyncJitterPeriod != nil {
		allErrs = append(allErrs, apivalidation.ValidateNonnegativeField(int64(cfg.SyncJitterPeriod.Duration), fldPath.Child("syncJitterPeriod"))...)
	}

	return allErrs
}

var availableShootPurposes = sets.NewString(
	string(gardencore.ShootPurposeEvaluation),
	string(gardencore.ShootPurposeTesting),
	string(gardencore.ShootPurposeDevelopment),
	string(gardencore.ShootPurposeInfrastructure),
	string(gardencore.ShootPurposeProduction),
)

// ValidateBackupEntryControllerConfiguration validates the BackupEntry controller configuration.
func ValidateBackupEntryControllerConfiguration(cfg *config.BackupEntryControllerConfiguration, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if len(cfg.DeletionGracePeriodShootPurposes) > 0 && (cfg.DeletionGracePeriodHours == nil || *cfg.DeletionGracePeriodHours <= 0) {
		allErrs = append(allErrs, field.Forbidden(fldPath.Child("deletionGracePeriodShootPurposes"), "must specify grace period hours > 0 when specifying purposes"))
	}

	for i, purpose := range cfg.DeletionGracePeriodShootPurposes {
		if !availableShootPurposes.Has(string(purpose)) {
			allErrs = append(allErrs, field.NotSupported(fldPath.Child("deletionGracePeriodShootPurposes").Index(i), purpose, availableShootPurposes.List()))
		}
	}

	return allErrs
}

// ValidateBastionControllerConfiguration validates the bastion configuration.
func ValidateBastionControllerConfiguration(cfg *config.BastionControllerConfiguration, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if cfg.ConcurrentSyncs != nil {
		allErrs = append(allErrs, apivalidation.ValidateNonnegativeField(int64(*cfg.ConcurrentSyncs), fldPath.Child("concurrentSyncs"))...)
	}

	return allErrs
}
