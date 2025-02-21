/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package rootfs

import (
	"os"

	"github.com/talos-systems/talos/internal/app/machined/internal/phase"
	"github.com/talos-systems/talos/internal/app/machined/internal/platform"
	"github.com/talos-systems/talos/internal/app/machined/internal/runtime"
	"github.com/talos-systems/talos/pkg/userdata"
)

// VarDirectories represents the VarDirectories task.
type VarDirectories struct{}

// NewVarDirectoriesTask initializes and returns an VarDirectories task.
func NewVarDirectoriesTask() phase.Task {
	return &VarDirectories{}
}

// RuntimeFunc returns the runtime function.
func (task *VarDirectories) RuntimeFunc(mode runtime.Mode) phase.RuntimeFunc {
	return task.runtime
}

func (task *VarDirectories) runtime(platform platform.Platform, data *userdata.UserData) (err error) {
	for _, p := range []string{"/var/log", "/var/lib/kubelet", "/var/log/pods"} {
		if err = os.MkdirAll(p, 0700); err != nil {
			return err
		}
	}

	return nil
}
