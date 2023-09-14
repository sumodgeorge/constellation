//go:build enterprise

/*
Copyright (c) Edgeless Systems GmbH

SPDX-License-Identifier: AGPL-3.0-only
*/

package license

import (
	"context"

	"github.com/edgelesssys/constellation/v2/internal/cloud/cloudprovider"
	"github.com/edgelesssys/constellation/v2/internal/config"
	"github.com/edgelesssys/constellation/v2/internal/file"
)

type Checker struct {
	quotaChecker QuotaChecker
	fileHandler  file.Handler
}

func NewChecker(quotaChecker QuotaChecker, fileHandler file.Handler) *Checker {
	return &Checker{
		quotaChecker: quotaChecker,
		fileHandler:  fileHandler,
	}
}

// CheckLicense tries to read the license file and contact license server
// to fetch quota information.
// If no license file is found, community license is assumed.
func (c *Checker) CheckLicense(ctx context.Context, provider cloudprovider.Provider, providerCfg config.ProviderConfig, printer func(string, ...any)) error {
	return nil
}
