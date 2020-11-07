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

package bl

import (
	"context"

	godo "git.mammoth.com.au/github/go-binarylane"
)

// ServerActionsService is an interface for interacting with BinaryLane's server action api.
type ServerActionsService interface {
	Shutdown(int) (*Action, error)
	ShutdownByTag(string) (Actions, error)
	PowerOff(int) (*Action, error)
	PowerOffByTag(string) (Actions, error)
	PowerOn(int) (*Action, error)
	PowerOnByTag(string) (Actions, error)
	PowerCycle(int) (*Action, error)
	PowerCycleByTag(string) (Actions, error)
	Reboot(int) (*Action, error)
	Restore(int, int) (*Action, error)
	Resize(int, string, bool) (*Action, error)
	Rename(int, string) (*Action, error)
	Snapshot(int, string) (*Action, error)
	SnapshotByTag(string, string) (Actions, error)
	EnableBackups(int) (*Action, error)
	EnableBackupsByTag(string) (Actions, error)
	DisableBackups(int) (*Action, error)
	DisableBackupsByTag(string) (Actions, error)
	PasswordReset(int) (*Action, error)
	RebuildByImageID(int, int) (*Action, error)
	RebuildByImageSlug(int, string) (*Action, error)
	ChangeKernel(int, int) (*Action, error)
	EnableIPv6(int) (*Action, error)
	EnableIPv6ByTag(string) (Actions, error)
	EnablePrivateNetworking(int) (*Action, error)
	EnablePrivateNetworkingByTag(string) (Actions, error)
	Get(int, int) (*Action, error)
	GetByURI(string) (*Action, error)
}

type serverActionsService struct {
	client *godo.Client
}

var _ ServerActionsService = &serverActionsService{}

// NewServerActionsService builds an instance of ServerActionsService.
func NewServerActionsService(godoClient *godo.Client) ServerActionsService {
	return &serverActionsService{
		client: godoClient,
	}
}

func (das *serverActionsService) handleActionResponse(a *godo.Action, err error) (*Action, error) {
	if err != nil {
		return nil, err
	}

	return &Action{Action: a}, nil
}

func (das *serverActionsService) handleTagActionResponse(a []godo.Action, err error) (Actions, error) {
	if err != nil {
		return nil, err
	}

	actions := make([]Action, 0, len(a))

	for _, action := range a {
		actions = append(actions, Action{Action: &action})
	}

	return actions, nil
}

func (das *serverActionsService) Shutdown(id int) (*Action, error) {
	a, _, err := das.client.ServerActions.Shutdown(context.TODO(), id)
	return das.handleActionResponse(a, err)
}

func (das *serverActionsService) ShutdownByTag(tag string) (Actions, error) {
	a, _, err := das.client.ServerActions.ShutdownByTag(context.TODO(), tag)
	return das.handleTagActionResponse(a, err)
}

func (das *serverActionsService) PowerOff(id int) (*Action, error) {
	a, _, err := das.client.ServerActions.PowerOff(context.TODO(), id)
	return das.handleActionResponse(a, err)
}

func (das *serverActionsService) PowerOffByTag(tag string) (Actions, error) {
	a, _, err := das.client.ServerActions.PowerOffByTag(context.TODO(), tag)
	return das.handleTagActionResponse(a, err)
}

func (das *serverActionsService) PowerOn(id int) (*Action, error) {
	a, _, err := das.client.ServerActions.PowerOn(context.TODO(), id)
	return das.handleActionResponse(a, err)
}

func (das *serverActionsService) PowerOnByTag(tag string) (Actions, error) {
	a, _, err := das.client.ServerActions.PowerOnByTag(context.TODO(), tag)
	return das.handleTagActionResponse(a, err)
}

func (das *serverActionsService) PowerCycle(id int) (*Action, error) {
	a, _, err := das.client.ServerActions.PowerCycle(context.TODO(), id)
	return das.handleActionResponse(a, err)
}

func (das *serverActionsService) PowerCycleByTag(tag string) (Actions, error) {
	a, _, err := das.client.ServerActions.PowerCycleByTag(context.TODO(), tag)
	return das.handleTagActionResponse(a, err)
}

func (das *serverActionsService) Reboot(id int) (*Action, error) {
	a, _, err := das.client.ServerActions.Reboot(context.TODO(), id)
	return das.handleActionResponse(a, err)
}

func (das *serverActionsService) Restore(id, imageID int) (*Action, error) {
	a, _, err := das.client.ServerActions.Restore(context.TODO(), id, imageID)
	return das.handleActionResponse(a, err)
}

func (das *serverActionsService) Resize(id int, sizeSlug string, resizeDisk bool) (*Action, error) {
	a, _, err := das.client.ServerActions.Resize(context.TODO(), id, sizeSlug, resizeDisk)
	return das.handleActionResponse(a, err)
}

func (das *serverActionsService) Rename(id int, name string) (*Action, error) {
	a, _, err := das.client.ServerActions.Rename(context.TODO(), id, name)
	return das.handleActionResponse(a, err)
}

func (das *serverActionsService) Snapshot(id int, name string) (*Action, error) {
	a, _, err := das.client.ServerActions.Snapshot(context.TODO(), id, name)
	return das.handleActionResponse(a, err)
}

func (das *serverActionsService) SnapshotByTag(tag string, name string) (Actions, error) {
	a, _, err := das.client.ServerActions.SnapshotByTag(context.TODO(), tag, name)
	return das.handleTagActionResponse(a, err)
}

func (das *serverActionsService) EnableBackups(id int) (*Action, error) {
	a, _, err := das.client.ServerActions.EnableBackups(context.TODO(), id)
	return das.handleActionResponse(a, err)
}

func (das *serverActionsService) EnableBackupsByTag(tag string) (Actions, error) {
	a, _, err := das.client.ServerActions.EnableBackupsByTag(context.TODO(), tag)
	return das.handleTagActionResponse(a, err)
}

func (das *serverActionsService) DisableBackups(id int) (*Action, error) {
	a, _, err := das.client.ServerActions.DisableBackups(context.TODO(), id)
	return das.handleActionResponse(a, err)
}

func (das *serverActionsService) DisableBackupsByTag(tag string) (Actions, error) {
	a, _, err := das.client.ServerActions.DisableBackupsByTag(context.TODO(), tag)
	return das.handleTagActionResponse(a, err)
}

func (das *serverActionsService) PasswordReset(id int) (*Action, error) {
	a, _, err := das.client.ServerActions.PasswordReset(context.TODO(), id)
	return das.handleActionResponse(a, err)
}

func (das *serverActionsService) RebuildByImageID(id, imageID int) (*Action, error) {
	a, _, err := das.client.ServerActions.RebuildByImageID(context.TODO(), id, imageID)
	return das.handleActionResponse(a, err)
}

func (das *serverActionsService) RebuildByImageSlug(id int, slug string) (*Action, error) {
	a, _, err := das.client.ServerActions.RebuildByImageSlug(context.TODO(), id, slug)
	return das.handleActionResponse(a, err)
}

func (das *serverActionsService) ChangeKernel(id, kernelID int) (*Action, error) {
	a, _, err := das.client.ServerActions.ChangeKernel(context.TODO(), id, kernelID)
	return das.handleActionResponse(a, err)
}

func (das *serverActionsService) EnableIPv6(id int) (*Action, error) {
	a, _, err := das.client.ServerActions.EnableIPv6(context.TODO(), id)
	return das.handleActionResponse(a, err)
}

func (das *serverActionsService) EnableIPv6ByTag(tag string) (Actions, error) {
	a, _, err := das.client.ServerActions.EnableIPv6ByTag(context.TODO(), tag)
	return das.handleTagActionResponse(a, err)
}

func (das *serverActionsService) EnablePrivateNetworking(id int) (*Action, error) {
	a, _, err := das.client.ServerActions.EnablePrivateNetworking(context.TODO(), id)
	return das.handleActionResponse(a, err)
}

func (das *serverActionsService) EnablePrivateNetworkingByTag(tag string) (Actions, error) {
	a, _, err := das.client.ServerActions.EnablePrivateNetworkingByTag(context.TODO(), tag)
	return das.handleTagActionResponse(a, err)
}

func (das *serverActionsService) Get(id int, actionID int) (*Action, error) {
	a, _, err := das.client.ServerActions.Get(context.TODO(), id, actionID)
	return das.handleActionResponse(a, err)
}

func (das *serverActionsService) GetByURI(uri string) (*Action, error) {
	a, _, err := das.client.ServerActions.GetByURI(context.TODO(), uri)
	return das.handleActionResponse(a, err)
}