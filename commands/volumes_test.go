/*
Copyright 2018 The Doctl Authors All rights reserved.
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

package commands

import (
	"testing"

	"git.mammoth.com.au/github/bl-cli"
	"git.mammoth.com.au/github/bl-cli/do"
	godo "git.mammoth.com.au/github/go-binarylane"
	"github.com/stretchr/testify/assert"
)

var (
	testVolume = do.Volume{
		Volume: &godo.Volume{
			ID:            "00000000-0000-4000-8000-000000000000",
			SizeGigaBytes: 100,
			Name:          "test-volume",
			Description:   "test description",
			Region:        &godo.Region{Slug: "atlantis"},
		},
	}
	testVolumeList = []do.Volume{
		testVolume,
	}
)

func TestVolumeCommand(t *testing.T) {
	cmd := Volume()
	assert.NotNil(t, cmd)
	assertCommandNames(t, cmd, "create", "delete", "get", "list", "snapshot")
}

func TestVolumesGet(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		tm.volumes.EXPECT().Get("test-volume").Return(&testVolume, nil)

		config.Args = append(config.Args, "test-volume")

		err := RunVolumeGet(config)
		assert.NoError(t, err)
	})
}

func TestVolumesList(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		tm.volumes.EXPECT().List().Return(testVolumeList, nil)

		err := RunVolumeList(config)
		assert.NoError(t, err)
	})
}

func TestVolumesListID(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		tm.volumes.EXPECT().List().Return(testVolumeList, nil)

		config.Args = append(config.Args, testVolume.ID)

		err := RunVolumeList(config)
		assert.NoError(t, err)
	})
}

func TestVolumesListName(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		tm.volumes.EXPECT().List().Return(testVolumeList, nil)

		config.Args = append(config.Args, "test-volume")

		err := RunVolumeList(config)
		assert.NoError(t, err)
	})
}

func TestVolumeCreate(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		tcr := godo.VolumeCreateRequest{
			Name:          "test-volume",
			SizeGigaBytes: 100,
			Region:        "atlantis",
			Description:   "test description",
			Tags:          []string{"one", "two"},
		}
		tm.volumes.EXPECT().CreateVolume(&tcr).Return(&testVolume, nil)

		config.Args = append(config.Args, "test-volume")

		config.Doit.Set(config.NS, blcli.ArgVolumeRegion, "atlantis")
		config.Doit.Set(config.NS, blcli.ArgVolumeSize, "100GiB")
		config.Doit.Set(config.NS, blcli.ArgVolumeDesc, "test description")
		config.Doit.Set(config.NS, blcli.ArgTag, []string{"one", "two"})

		err := RunVolumeCreate(config)
		assert.NoError(t, err)
	})
}

func TestVolumeCreateFromSnapshot(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		tcr := godo.VolumeCreateRequest{
			Name:          "test-volume",
			SizeGigaBytes: 100,
			SnapshotID:    "ed6414f7-7873-4dd2-90cf-f4f354c293e6",
			Description:   "test description",
			Tags:          []string{"one", "two"},
		}
		tm.volumes.EXPECT().CreateVolume(&tcr).Return(&testVolume, nil)

		config.Args = append(config.Args, "test-volume")

		config.Doit.Set(config.NS, blcli.ArgVolumeSnapshot, "ed6414f7-7873-4dd2-90cf-f4f354c293e6")
		config.Doit.Set(config.NS, blcli.ArgVolumeSize, "100GiB")
		config.Doit.Set(config.NS, blcli.ArgVolumeDesc, "test description")
		config.Doit.Set(config.NS, blcli.ArgTag, []string{"one", "two"})

		err := RunVolumeCreate(config)
		assert.NoError(t, err)
	})
}

func TestVolumesDelete(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		tm.volumes.EXPECT().DeleteVolume("test-volume").Return(nil)

		config.Args = append(config.Args, "test-volume")

		config.Doit.Set(config.NS, blcli.ArgForce, true)

		err := RunVolumeDelete(config)
		assert.NoError(t, err)
	})
}

func TestVolumesSnapshot(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		tcr := godo.SnapshotCreateRequest{
			VolumeID:    testVolume.ID,
			Name:        "test-volume-snapshot",
			Description: "test description",
			Tags:        []string{"one", "two"},
		}
		tm.volumes.EXPECT().CreateSnapshot(&tcr).Return(nil, nil)

		config.Args = append(config.Args, testVolume.ID)
		config.Doit.Set(config.NS, blcli.ArgSnapshotName, "test-volume-snapshot")
		config.Doit.Set(config.NS, blcli.ArgSnapshotDesc, "test description")
		config.Doit.Set(config.NS, blcli.ArgTag, []string{"one", "two"})

		err := RunVolumeSnapshot(config)
		assert.NoError(t, err)
	})
}
