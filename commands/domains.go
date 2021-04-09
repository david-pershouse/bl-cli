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
	"errors"
	"fmt"
	"strconv"

	"github.com/binarylane/bl-cli"
	"github.com/binarylane/bl-cli/bl"
	"github.com/binarylane/bl-cli/commands/displayers"
	godo "github.com/binarylane/go-binarylane"
	"github.com/spf13/cobra"
)

// Domain creates the domain commands hierarchy.
func Domain() *Command {
	cmd := &Command{
		Command: &cobra.Command{
			Use:   "domain",
			Short: "Display commands that manage domains",
			Long:  "Use the subcommands of `bl compute domain` to manage domains you have purchased from a domain name registrar that you are managing through the BinaryLane DNS interface.",
		},
	}

	cmdDomainCreate := CmdBuilder(cmd, RunDomainCreate, "create <domain>", "Add a domain to your account", `Use this command to add a domain to your account.`, Writer,
		aliasOpt("c"), displayerType(&displayers.Domain{}))
	AddStringFlag(cmdDomainCreate, blcli.ArgIPAddress, "", "", "Creates an A record when an IPv4 address is provided")

	CmdBuilder(cmd, RunDomainList, "list", "List all domains on your account", `Use this command to retrieve a list of domains on your account.`, Writer,
		aliasOpt("ls"), displayerType(&displayers.Domain{}))

	CmdBuilder(cmd, RunDomainGet, "get <domain>", "Retrieve information about a domain", `Use this command to retrieve information about the specified domain on your account.`, Writer,
		aliasOpt("g"), displayerType(&displayers.Domain{}))

	cmdRunDomainDelete := CmdBuilder(cmd, RunDomainDelete, "delete <domain>", "Permanently delete a domain from your account", `Use this command to delete a domain from your account. This is irreversible.`, Writer, aliasOpt("d", "rm"))
	AddBoolFlag(cmdRunDomainDelete, blcli.ArgForce, blcli.ArgShortForce, false, "Delete domain without confirmation prompt")

	cmdRecord := &Command{
		Command: &cobra.Command{
			Use:   "records",
			Short: "Manage DNS records",
			Long:  "Use the subcommands of `bl compute domain records` to manage the DNS records for your domains.",
		},
	}
	cmd.AddCommand(cmdRecord)

	CmdBuilder(cmdRecord, RunRecordList, "list <domain>", "List the DNS records for a domain", `Use this command to list the DNS records for a domain.`, Writer,
		aliasOpt("ls"), displayerType(&displayers.DomainRecord{}))

	cmdRecordCreate := CmdBuilder(cmdRecord, RunRecordCreate, "create <domain>", "Create a DNS record", `Use this command to create DNS records for a domain.`, Writer,
		aliasOpt("c"), displayerType(&displayers.DomainRecord{}))
	AddStringFlag(cmdRecordCreate, blcli.ArgRecordType, "", "", "The type of DNS record")
	AddStringFlag(cmdRecordCreate, blcli.ArgRecordName, "", "", "The host name, alias, or service being defined by the record")
	AddStringFlag(cmdRecordCreate, blcli.ArgRecordData, "", "", "Record data; varies depending on record type")
	AddIntFlag(cmdRecordCreate, blcli.ArgRecordPriority, "", 0, "Record priority")
	AddIntFlag(cmdRecordCreate, blcli.ArgRecordPort, "", 0, "The port value for an SRV record")
	AddIntFlag(cmdRecordCreate, blcli.ArgRecordTTL, "", 1800, "The record's Time To Live value, in seconds")
	AddIntFlag(cmdRecordCreate, blcli.ArgRecordWeight, "", 0, "The weight value for an SRV record")
	AddIntFlag(cmdRecordCreate, blcli.ArgRecordFlags, "", 0, "An unsigned integer between 0-255 used for CAA records")
	AddStringFlag(cmdRecordCreate, blcli.ArgRecordTag, "", "", "The parameter tag for CAA records. Valid values are `issue`, `issuewild`, or `iodef`")

	cmdRunRecordDelete := CmdBuilder(cmdRecord, RunRecordDelete, "delete <domain> <record-id>...", "Delete a DNS record", `Use this command to delete DNS records for a domain.`, Writer,
		aliasOpt("d"))
	AddBoolFlag(cmdRunRecordDelete, blcli.ArgForce, blcli.ArgShortForce, false, "Delete record without confirmation prompt")

	cmdRecordUpdate := CmdBuilder(cmdRecord, RunRecordUpdate, "update <domain>", "Update a DNS record", `Use this command to update or change DNS records for a domain.`, Writer,
		aliasOpt("u"), displayerType(&displayers.DomainRecord{}))
	AddIntFlag(cmdRecordUpdate, blcli.ArgRecordID, "", 0, "Record ID")
	AddStringFlag(cmdRecordUpdate, blcli.ArgRecordType, "", "", "The type of DNS record")
	AddStringFlag(cmdRecordUpdate, blcli.ArgRecordName, "", "", "The host name, alias, or service being defined by the record")
	AddStringFlag(cmdRecordUpdate, blcli.ArgRecordData, "", "", "Record data; varies depending on record type")
	AddIntFlag(cmdRecordUpdate, blcli.ArgRecordPriority, "", 0, "Record priority")
	AddIntFlag(cmdRecordUpdate, blcli.ArgRecordPort, "", 0, "The port value for an SRV record")
	AddIntFlag(cmdRecordUpdate, blcli.ArgRecordTTL, "", 1800, "The record's Time To Live value, in seconds")
	AddIntFlag(cmdRecordUpdate, blcli.ArgRecordWeight, "", 0, "The weight value for an SRV record")
	AddIntFlag(cmdRecordUpdate, blcli.ArgRecordFlags, "", 0, "An unsigned integer between 0-255 used for CAA records")
	AddStringFlag(cmdRecordUpdate, blcli.ArgRecordTag, "", "", "The parameter tag for CAA records. Valid values are `issue`, `issuewild`, or `iodef`")

	return cmd
}

// RunDomainCreate runs domain create.
func RunDomainCreate(c *CmdConfig) error {
	err := ensureOneArg(c)
	if err != nil {
		return err
	}
	domainName := c.Args[0]

	ds := c.Domains()

	ipAddress, err := c.Doit.GetString(c.NS, "ip-address")
	if err != nil {
		return err
	}

	req := &godo.DomainCreateRequest{
		Name:      domainName,
		IPAddress: ipAddress,
	}

	d, err := ds.Create(req)
	if err != nil {
		return err
	}

	return c.Display(&displayers.Domain{Domains: bl.Domains{*d}})
}

// RunDomainList runs domain create.
func RunDomainList(c *CmdConfig) error {

	ds := c.Domains()

	domains, err := ds.List()
	if err != nil {
		return err
	}

	item := &displayers.Domain{Domains: domains}
	return c.Display(item)
}

// RunDomainGet retrieves a domain by name.
func RunDomainGet(c *CmdConfig) error {
	err := ensureOneArg(c)
	if err != nil {
		return err
	}
	id := c.Args[0]

	ds := c.Domains()

	if len(id) < 1 {
		return errors.New("Invalid domain name.")
	}

	d, err := ds.Get(id)
	if err != nil {
		return err
	}

	item := &displayers.Domain{Domains: bl.Domains{*d}}
	return c.Display(item)
}

// RunDomainDelete deletes a domain by name.
func RunDomainDelete(c *CmdConfig) error {
	err := ensureOneArg(c)
	if err != nil {
		return err
	}
	name := c.Args[0]

	force, err := c.Doit.GetBool(c.NS, blcli.ArgForce)
	if err != nil {
		return err
	}

	if force || AskForConfirmDelete("domain", 1) == nil {
		ds := c.Domains()

		if len(name) < 1 {
			return errors.New("Invalid domain name.")
		}

		err := ds.Delete(name)
		return err
	}

	return fmt.Errorf("Operation aborted.")
}

// RunRecordList list records for a domain.
func RunRecordList(c *CmdConfig) error {
	err := ensureOneArg(c)
	if err != nil {
		return err
	}
	name := c.Args[0]

	ds := c.Domains()

	if len(name) < 1 {
		return errors.New("Domain name is missing.")
	}

	list, err := ds.Records(name)
	if err != nil {
		return err
	}

	items := &displayers.DomainRecord{DomainRecords: list}
	return c.Display(items)

}

// RunRecordCreate creates a domain record.
func RunRecordCreate(c *CmdConfig) error {
	err := ensureOneArg(c)
	if err != nil {
		return err
	}
	name := c.Args[0]

	ds := c.Domains()

	rType, err := c.Doit.GetString(c.NS, blcli.ArgRecordType)
	if err != nil {
		return err
	}

	rName, err := c.Doit.GetString(c.NS, blcli.ArgRecordName)
	if err != nil {
		return err
	}

	rData, err := c.Doit.GetString(c.NS, blcli.ArgRecordData)
	if err != nil {
		return err
	}

	rPriority, err := c.Doit.GetInt(c.NS, blcli.ArgRecordPriority)
	if err != nil {
		return err
	}

	rPort, err := c.Doit.GetIntPtr(c.NS, blcli.ArgRecordPort)
	if err != nil {
		return err
	}

	rTTL, err := c.Doit.GetInt(c.NS, blcli.ArgRecordTTL)
	if err != nil {
		return err
	}

	rWeight, err := c.Doit.GetInt(c.NS, blcli.ArgRecordWeight)
	if err != nil {
		return err
	}

	rFlags, err := c.Doit.GetInt(c.NS, blcli.ArgRecordFlags)
	if err != nil {
		return err
	}

	rTag, err := c.Doit.GetString(c.NS, blcli.ArgRecordTag)
	if err != nil {
		return err
	}

	drcr := &bl.DomainRecordEditRequest{
		Type:     rType,
		Name:     rName,
		Data:     rData,
		Priority: rPriority,
		Port:     rPort,
		TTL:      rTTL,
		Weight:   rWeight,
		Flags:    rFlags,
		Tag:      rTag,
	}

	if len(drcr.Type) == 0 {
		return errors.New("Record request is missing type.")
	}

	r, err := ds.CreateRecord(name, drcr)
	if err != nil {
		return err
	}

	item := &displayers.DomainRecord{DomainRecords: bl.DomainRecords{*r}}
	return c.Display(item)

}

// RunRecordDelete deletes a domain record.
func RunRecordDelete(c *CmdConfig) error {
	if len(c.Args) < 2 {
		return blcli.NewMissingArgsErr(c.NS)
	}

	force, err := c.Doit.GetBool(c.NS, blcli.ArgForce)
	if err != nil {
		return err
	}

	domainName, ids := c.Args[0], c.Args[1:]
	if len(ids) < 1 {
		return blcli.NewMissingArgsErr(c.NS)
	}

	if force || AskForConfirmDelete("domain record", len(ids)) == nil {
		ds := c.Domains()

		for _, i := range ids {
			id, err := strconv.Atoi(i)
			if err != nil {
				return fmt.Errorf("Invalid record id %q", i)
			}

			err = ds.DeleteRecord(domainName, id)
			if err != nil {
				return err
			}
		}
	} else {
		return fmt.Errorf("Operation aborted.")
	}

	return nil

}

// RunRecordUpdate updates a domain record.
func RunRecordUpdate(c *CmdConfig) error {
	err := ensureOneArg(c)
	if err != nil {
		return err
	}
	domainName := c.Args[0]

	ds := c.Domains()

	recordID, err := c.Doit.GetInt(c.NS, blcli.ArgRecordID)
	if err != nil {
		return err
	}

	rType, err := c.Doit.GetString(c.NS, blcli.ArgRecordType)
	if err != nil {
		return err
	}

	rName, err := c.Doit.GetString(c.NS, blcli.ArgRecordName)
	if err != nil {
		return err
	}

	rData, err := c.Doit.GetString(c.NS, blcli.ArgRecordData)
	if err != nil {
		return err
	}

	rPriority, err := c.Doit.GetInt(c.NS, blcli.ArgRecordPriority)
	if err != nil {
		return err
	}

	rPort, err := c.Doit.GetIntPtr(c.NS, blcli.ArgRecordPort)
	if err != nil {
		return err
	}

	rTTL, err := c.Doit.GetInt(c.NS, blcli.ArgRecordTTL)
	if err != nil {
		return err
	}

	rWeight, err := c.Doit.GetInt(c.NS, blcli.ArgRecordWeight)
	if err != nil {
		return err
	}

	rFlags, err := c.Doit.GetInt(c.NS, blcli.ArgRecordFlags)
	if err != nil {
		return err
	}

	rTag, err := c.Doit.GetString(c.NS, blcli.ArgRecordTag)
	if err != nil {
		return err
	}

	drcr := &bl.DomainRecordEditRequest{
		Type:     rType,
		Name:     rName,
		Data:     rData,
		Priority: rPriority,
		Port:     rPort,
		TTL:      rTTL,
		Weight:   rWeight,
		Flags:    rFlags,
		Tag:      rTag,
	}

	r, err := ds.EditRecord(domainName, recordID, drcr)
	if err != nil {
		return err
	}

	item := &displayers.DomainRecord{DomainRecords: bl.DomainRecords{*r}}
	return c.Display(item)
}
