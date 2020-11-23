package commands

import (
	"testing"

	"git.mammoth.com.au/github/bl-cli"
	"git.mammoth.com.au/github/bl-cli/bl"
	godo "git.mammoth.com.au/github/go-binarylane"

	"github.com/stretchr/testify/assert"
)

var (
	testLoadBalancer = bl.LoadBalancer{
		LoadBalancer: &godo.LoadBalancer{
			Algorithm: "round_robin",
			Region: &godo.Region{
				Slug: "nyc1",
			},
			SizeSlug:       "lb-small",
			StickySessions: &godo.StickySessions{},
			HealthCheck:    &godo.HealthCheck{},
		}}

	testLoadBalancerList = bl.LoadBalancers{
		testLoadBalancer,
	}
)

func TestLoadBalancerCommand(t *testing.T) {
	cmd := LoadBalancer()
	assert.NotNil(t, cmd)
	assertCommandNames(t, cmd, "get", "list", "create", "update", "delete", "add-servers", "remove-servers", "add-forwarding-rules", "remove-forwarding-rules")
}

func TestLoadBalancerGet(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		lbID := "cde2c0d6-41e3-479e-ba60-ad971227232c"
		tm.loadBalancers.EXPECT().Get(lbID).Return(&testLoadBalancer, nil)

		config.Args = append(config.Args, lbID)

		err := RunLoadBalancerGet(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerGetNoID(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		err := RunLoadBalancerGet(config)
		assert.Error(t, err)
	})
}

func TestLoadBalancerList(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		tm.loadBalancers.EXPECT().List().Return(testLoadBalancerList, nil)

		err := RunLoadBalancerList(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerCreateWithInvalidServerIDsArgs(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		config.Doit.Set(config.NS, blcli.ArgServerIDs, []string{"bogus"})

		err := RunLoadBalancerCreate(config)
		assert.Error(t, err)
	})
}

func TestLoadBalancerCreateWithMalformedForwardingRulesArgs(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		config.Doit.Set(config.NS, blcli.ArgForwardingRules, "something,something")

		err := RunLoadBalancerCreate(config)
		assert.Error(t, err)
	})
}

func TestLoadBalancerCreate(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		vpcUUID := "00000000-0000-4000-8000-000000000000"
		r := godo.LoadBalancerRequest{
			Name:       "lb-name",
			Region:     "nyc1",
			SizeSlug:   "lb-small",
			ServerIDs: []int{1, 2},
			StickySessions: &godo.StickySessions{
				Type: "none",
			},
			HealthCheck: &godo.HealthCheck{
				Protocol:               "http",
				Port:                   80,
				CheckIntervalSeconds:   4,
				ResponseTimeoutSeconds: 23,
				HealthyThreshold:       5,
				UnhealthyThreshold:     10,
			},
			ForwardingRules: []godo.ForwardingRule{
				{
					EntryProtocol:  "tcp",
					EntryPort:      3306,
					TargetProtocol: "tcp",
					TargetPort:     3306,
					TlsPassthrough: true,
				},
			},
			VPCUUID: vpcUUID,
		}
		tm.loadBalancers.EXPECT().Create(&r).Return(&testLoadBalancer, nil)

		config.Doit.Set(config.NS, blcli.ArgRegionSlug, "nyc1")
		config.Doit.Set(config.NS, blcli.ArgSizeSlug, "lb-small")
		config.Doit.Set(config.NS, blcli.ArgLoadBalancerName, "lb-name")
		config.Doit.Set(config.NS, blcli.ArgVPCUUID, vpcUUID)
		config.Doit.Set(config.NS, blcli.ArgServerIDs, []string{"1", "2"})
		config.Doit.Set(config.NS, blcli.ArgStickySessions, "type:none")
		config.Doit.Set(config.NS, blcli.ArgHealthCheck, "protocol:http,port:80,check_interval_seconds:4,response_timeout_seconds:23,healthy_threshold:5,unhealthy_threshold:10")
		config.Doit.Set(config.NS, blcli.ArgForwardingRules, "entry_protocol:tcp,entry_port:3306,target_protocol:tcp,target_port:3306,tls_passthrough:true")

		err := RunLoadBalancerCreate(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerUpdate(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		lbID := "cde2c0d6-41e3-479e-ba60-ad971227232c"
		r := godo.LoadBalancerRequest{
			Name:       "lb-name",
			Region:     "nyc1",
			ServerIDs: []int{1, 2},
			StickySessions: &godo.StickySessions{
				Type:             "cookies",
				CookieName:       "DO-LB",
				CookieTtlSeconds: 5,
			},
			HealthCheck: &godo.HealthCheck{
				Protocol:               "http",
				Port:                   80,
				CheckIntervalSeconds:   4,
				ResponseTimeoutSeconds: 23,
				HealthyThreshold:       5,
				UnhealthyThreshold:     10,
			},
			ForwardingRules: []godo.ForwardingRule{
				{
					EntryProtocol:  "http",
					EntryPort:      80,
					TargetProtocol: "http",
					TargetPort:     80,
				},
			},
		}

		tm.loadBalancers.EXPECT().Update(lbID, &r).Return(&testLoadBalancer, nil)

		config.Args = append(config.Args, lbID)
		config.Doit.Set(config.NS, blcli.ArgRegionSlug, "nyc1")
		config.Doit.Set(config.NS, blcli.ArgSizeSlug, "")
		config.Doit.Set(config.NS, blcli.ArgLoadBalancerName, "lb-name")
		config.Doit.Set(config.NS, blcli.ArgServerIDs, []string{"1", "2"})
		config.Doit.Set(config.NS, blcli.ArgStickySessions, "type:cookies,cookie_name:DO-LB,cookie_ttl_seconds:5")
		config.Doit.Set(config.NS, blcli.ArgHealthCheck, "protocol:http,port:80,check_interval_seconds:4,response_timeout_seconds:23,healthy_threshold:5,unhealthy_threshold:10")
		config.Doit.Set(config.NS, blcli.ArgForwardingRules, "entry_protocol:http,entry_port:80,target_protocol:http,target_port:80")

		err := RunLoadBalancerUpdate(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerUpdateNoID(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		err := RunLoadBalancerUpdate(config)
		assert.Error(t, err)
	})
}

func TestLoadBalancerDelete(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		lbID := "cde2c0d6-41e3-479e-ba60-ad971227232c"
		tm.loadBalancers.EXPECT().Delete(lbID).Return(nil)

		config.Args = append(config.Args, lbID)
		config.Doit.Set(config.NS, blcli.ArgForce, true)

		err := RunLoadBalancerDelete(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerDeleteNoID(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		err := RunLoadBalancerDelete(config)
		assert.Error(t, err)
	})
}

func TestLoadBalancerAddServers(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		lbID := "cde2c0d6-41e3-479e-ba60-ad971227232c"
		tm.loadBalancers.EXPECT().AddServers(lbID, 1, 23).Return(nil)

		config.Args = append(config.Args, lbID)
		config.Doit.Set(config.NS, blcli.ArgServerIDs, []string{"1", "23"})

		err := RunLoadBalancerAddServers(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerAddServersNoID(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		err := RunLoadBalancerAddServers(config)
		assert.Error(t, err)
	})
}

func TestLoadBalancerRemoveServers(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		lbID := "cde2c0d6-41e3-479e-ba60-ad971227232c"
		tm.loadBalancers.EXPECT().RemoveServers(lbID, 321).Return(nil)

		config.Args = append(config.Args, lbID)
		config.Doit.Set(config.NS, blcli.ArgServerIDs, []string{"321"})

		err := RunLoadBalancerRemoveServers(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerRemoveServersNoID(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		err := RunLoadBalancerRemoveServers(config)
		assert.Error(t, err)
	})
}

func TestLoadBalancerAddForwardingRules(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		lbID := "cde2c0d6-41e3-479e-ba60-ad971227232c"
		forwardingRule := godo.ForwardingRule{
			EntryProtocol:  "http",
			EntryPort:      80,
			TargetProtocol: "http",
			TargetPort:     80,
		}
		tm.loadBalancers.EXPECT().AddForwardingRules(lbID, forwardingRule).Return(nil)

		config.Args = append(config.Args, lbID)
		config.Doit.Set(config.NS, blcli.ArgForwardingRules, "entry_protocol:http,entry_port:80,target_protocol:http,target_port:80")

		err := RunLoadBalancerAddForwardingRules(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerAddForwardingRulesNoID(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		err := RunLoadBalancerAddForwardingRules(config)
		assert.Error(t, err)
	})
}

func TestLoadBalancerRemoveForwardingRules(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		lbID := "cde2c0d6-41e3-479e-ba60-ad971227232c"
		forwardingRules := []godo.ForwardingRule{
			{
				EntryProtocol:  "http",
				EntryPort:      80,
				TargetProtocol: "http",
				TargetPort:     80,
			},
			{
				EntryProtocol:  "tcp",
				EntryPort:      3306,
				TargetProtocol: "tcp",
				TargetPort:     3306,
				TlsPassthrough: true,
			},
		}
		tm.loadBalancers.EXPECT().RemoveForwardingRules(lbID, forwardingRules[0], forwardingRules[1]).Return(nil)

		config.Args = append(config.Args, lbID)
		config.Doit.Set(config.NS, blcli.ArgForwardingRules, "entry_protocol:http,entry_port:80,target_protocol:http,target_port:80 entry_protocol:tcp,entry_port:3306,target_protocol:tcp,target_port:3306,tls_passthrough:true")

		err := RunLoadBalancerRemoveForwardingRules(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerRemoveForwardingRulesNoID(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		err := RunLoadBalancerRemoveForwardingRules(config)
		assert.Error(t, err)
	})
}
