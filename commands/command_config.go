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
	"fmt"
	"io"

	"github.com/binarylane/bl-cli"
	"github.com/binarylane/bl-cli/bl"
	"github.com/binarylane/bl-cli/commands/displayers"
	"github.com/spf13/viper"
)

// CmdConfig is a command configuration.
type CmdConfig struct {
	NS   string
	Doit blcli.Config
	Out  io.Writer
	Args []string

	initServices          func(*CmdConfig) error
	getContextAccessToken func() string
	setContextAccessToken func(string)

	// services
	Keys              func() bl.KeysService
	Sizes             func() bl.SizesService
	Regions           func() bl.RegionsService
	Images            func() bl.ImagesService
	ImageActions      func() bl.ImageActionsService
	LoadBalancers     func() bl.LoadBalancersService
	FloatingIPs       func() bl.FloatingIPsService
	FloatingIPActions func() bl.FloatingIPActionsService
	Servers           func() bl.ServersService
	ServerActions     func() bl.ServerActionsService
	Domains           func() bl.DomainsService
	Actions           func() bl.ActionsService
	Account           func() bl.AccountService
	Balance           func() bl.BalanceService
	BillingHistory    func() bl.BillingHistoryService
	Invoices          func() bl.InvoicesService
	Tags              func() bl.TagsService
	Snapshots         func() bl.SnapshotsService
	Firewalls         func() bl.FirewallsService
	Projects          func() bl.ProjectsService
	VPCs              func() bl.VPCsService
}

// NewCmdConfig creates an instance of a CmdConfig.
func NewCmdConfig(ns string, dc blcli.Config, out io.Writer, args []string, initGodo bool) (*CmdConfig, error) {

	cmdConfig := &CmdConfig{
		NS:   ns,
		Doit: dc,
		Out:  out,
		Args: args,

		initServices: func(c *CmdConfig) error {
			accessToken := c.getContextAccessToken()
			godoClient, err := c.Doit.GetGodoClient(Trace, accessToken)
			if err != nil {
				return fmt.Errorf("Unable to initialize BinaryLane API client: %s", err)
			}

			c.Keys = func() bl.KeysService { return bl.NewKeysService(godoClient) }
			c.Sizes = func() bl.SizesService { return bl.NewSizesService(godoClient) }
			c.Regions = func() bl.RegionsService { return bl.NewRegionsService(godoClient) }
			c.Images = func() bl.ImagesService { return bl.NewImagesService(godoClient) }
			c.ImageActions = func() bl.ImageActionsService { return bl.NewImageActionsService(godoClient) }
			c.FloatingIPs = func() bl.FloatingIPsService { return bl.NewFloatingIPsService(godoClient) }
			c.FloatingIPActions = func() bl.FloatingIPActionsService { return bl.NewFloatingIPActionsService(godoClient) }
			c.Servers = func() bl.ServersService { return bl.NewServersService(godoClient) }
			c.ServerActions = func() bl.ServerActionsService { return bl.NewServerActionsService(godoClient) }
			c.Domains = func() bl.DomainsService { return bl.NewDomainsService(godoClient) }
			c.Actions = func() bl.ActionsService { return bl.NewActionsService(godoClient) }
			c.Account = func() bl.AccountService { return bl.NewAccountService(godoClient) }
			c.Balance = func() bl.BalanceService { return bl.NewBalanceService(godoClient) }
			c.BillingHistory = func() bl.BillingHistoryService { return bl.NewBillingHistoryService(godoClient) }
			c.Invoices = func() bl.InvoicesService { return bl.NewInvoicesService(godoClient) }
			c.Tags = func() bl.TagsService { return bl.NewTagsService(godoClient) }
			c.Snapshots = func() bl.SnapshotsService { return bl.NewSnapshotsService(godoClient) }
			c.LoadBalancers = func() bl.LoadBalancersService { return bl.NewLoadBalancersService(godoClient) }
			c.Firewalls = func() bl.FirewallsService { return bl.NewFirewallsService(godoClient) }
			c.Projects = func() bl.ProjectsService { return bl.NewProjectsService(godoClient) }
			c.VPCs = func() bl.VPCsService { return bl.NewVPCsService(godoClient) }

			return nil
		},

		getContextAccessToken: func() string {
			context := Context
			if context == "" {
				context = viper.GetString("context")
			}
			token := ""

			switch context {
			case blcli.ArgDefaultContext:
				token = viper.GetString(blcli.ArgAccessToken)
			default:
				contexts := viper.GetStringMapString("auth-contexts")

				token = contexts[context]
			}

			return token
		},

		setContextAccessToken: func(token string) {
			context := Context
			if context == "" {
				context = viper.GetString("context")
			}

			switch context {
			case blcli.ArgDefaultContext:
				viper.Set(blcli.ArgAccessToken, token)
			default:
				contexts := viper.GetStringMapString("auth-contexts")
				contexts[context] = token

				viper.Set("auth-contexts", contexts)
			}
		},
	}

	if initGodo {
		if err := cmdConfig.initServices(cmdConfig); err != nil {
			return nil, err
		}
	}

	return cmdConfig, nil
}

// CmdRunner runs a command and passes in a cmdConfig.
type CmdRunner func(*CmdConfig) error

// Display displays the output from a command.
func (c *CmdConfig) Display(d displayers.Displayable) error {
	dc := &displayers.Displayer{
		Item: d,
		Out:  c.Out,
	}

	columnList, err := c.Doit.GetString(c.NS, blcli.ArgFormat)
	if err != nil {
		return err
	}

	withHeaders, err := c.Doit.GetBool(c.NS, blcli.ArgNoHeader)
	if err != nil {
		return err
	}

	dc.NoHeaders = withHeaders
	dc.ColumnList = columnList
	dc.OutputType = Output

	return dc.Display()
}
