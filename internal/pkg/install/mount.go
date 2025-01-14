/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package install

import (
	"log"

	"github.com/pkg/errors"
	"github.com/talos-systems/talos/internal/pkg/blockdevice/probe"
	"github.com/talos-systems/talos/internal/pkg/constants"
	"github.com/talos-systems/talos/internal/pkg/mount"
	"github.com/talos-systems/talos/pkg/userdata"
	"golang.org/x/sys/unix"
)

// Mount discovers the appropriate partitions by label and mounts them up
// to the appropriate mountpoint.
// TODO: See if we can consolidate this with rootfs/mount
func Mount(data *userdata.UserData) (err error) {
	var mp *mount.Points
	if mp, err = mountpoints(data.Install.Boot.InstallDevice.Device); err != nil {
		return errors.Errorf("error initializing block devices: %v", err)
	}

	iter := mp.Iter()
	for iter.Next() {
		if err = mount.WithRetry(iter.Value(), mount.WithPrefix(constants.NewRoot)); err != nil {
			return errors.Errorf("error mounting partitions: %v", err)
		}
	}
	if iter.Err() != nil {
		return iter.Err()
	}

	return nil
}

// nolint: dupl
func mountpoints(devpath string) (mountpoints *mount.Points, err error) {
	mountpoints = mount.NewMountPoints()
	for _, name := range []string{constants.DataPartitionLabel, constants.BootPartitionLabel} {
		var target string
		switch name {
		case constants.DataPartitionLabel:
			target = constants.DataMountPoint
		case constants.BootPartitionLabel:
			target = constants.BootMountPoint
		}

		var dev *probe.ProbedBlockDevice
		if dev, err = probe.DevForFileSystemLabel(devpath, name); err != nil {
			if name == constants.BootPartitionLabel {
				// A bootloader is not always required.
				log.Println("WARNING: no ESP partition was found")
				continue
			}
			return nil, errors.Errorf("failed to find device with label %s: %v", name, err)
		}

		mountpoint := mount.NewMountPoint(dev.Path, target, dev.SuperBlock.Type(), unix.MS_NOATIME, "")

		mountpoints.Set(name, mountpoint)
	}

	return mountpoints, nil
}
