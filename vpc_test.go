package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestVPCServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"vpc": {
				"id": "net539626f0798d7",
				"date_created": "2017-08-25 12:23:45",
				"region": "ewr",
				"description": "test1",
				"v4_subnet": "10.99.0.0",
				"v4_subnet_mask": 24
			}
		}
		`

		fmt.Fprint(writer, response)
	})

	options := &VPCReq{
		Region:      "ewr",
		Description: "test1",
	}

	net, _, err := client.VPC.Create(ctx, options)

	if err != nil {
		t.Errorf("VPC.Create returned %+v, expected %+v", err, nil)
	}

	expected := &VPC{
		ID:           "net539626f0798d7",
		Region:       "ewr",
		Description:  "test1",
		V4Subnet:     "10.99.0.0",
		V4SubnetMask: 24,
		DateCreated:  "2017-08-25 12:23:45",
	}

	if !reflect.DeepEqual(net, expected) {
		t.Errorf("VPC.Create returned %+v, expected %+v", net, expected)
	}
}

func TestVPCServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/net539626f0798d7", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.VPC.Delete(ctx, "net539626f0798d7")

	if err != nil {
		t.Errorf("VPC.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestVPCServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"vpcs": [{
				"id": "net539626f0798d7",
				"date_created": "2017-08-25 12:23:45",
				"region": "ewr",
				"description": "test1",
				"v4_subnet": "10.99.0.0",
				"v4_subnet_mask": 24
			}]
		}
		`
		fmt.Fprint(writer, response)
	})

	vpcs, _, _, err := client.VPC.List(ctx, nil)

	if err != nil {
		t.Errorf("VPC.List returned error: %v", err)
	}

	expected := []VPC{
		{
			ID:           "net539626f0798d7",
			Region:       "ewr",
			Description:  "test1",
			V4Subnet:     "10.99.0.0",
			V4SubnetMask: 24,
			DateCreated:  "2017-08-25 12:23:45",
		},
	}

	if !reflect.DeepEqual(vpcs, expected) {
		t.Errorf("VPC.List returned %+v, expected %+v", vpcs, expected)
	}
}

func TestVPCServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/net539626f0798d7", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.VPC.Update(ctx, "net539626f0798d7", "update")

	if err != nil {
		t.Errorf("VPC.Update returned %+v, expected %+v", err, nil)
	}
}

func TestVPCServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/net539626f0798d7", func(writer http.ResponseWriter, request *http.Request) {
		req := `{"vpc": {"id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60","date_created": "2020-10-10T01:56:20+00:00","region": "ewr","description": "sample desc","v4_subnet": "10.99.0.0","v4_subnet_mask": 24}}`
		fmt.Fprint(writer, req)
	})

	vpc, _, err := client.VPC.Get(ctx, "net539626f0798d7")
	if err != nil {
		t.Errorf("VPC.Get returned %+v, expected %+v", err, nil)
	}

	expected := &VPC{
		ID:           "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
		Region:       "ewr",
		Description:  "sample desc",
		V4Subnet:     "10.99.0.0",
		V4SubnetMask: 24,
		DateCreated:  "2020-10-10T01:56:20+00:00",
	}

	if !reflect.DeepEqual(vpc, expected) {
		t.Errorf("VPC.Get returned %+v, expected %+v", vpc, expected)
	}
}

func TestVPCServiceHandler_NATGatewayList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/59d6c282-00a7-4a92-9a41-3adad396abcd/nat-gateway", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"meta": {
				"total": 1,
				"links": {
					"next": "",
					"prev": ""
				}
			},
			"nat_gateways": [
				{
					"id": "7af46919-f6b0-4f34-8523-1d03911eabcd",
					"vpc_id": "59d6c282-00a7-4a92-9a41-3adad396abcd",
					"date_created": "2025-09-29 10:38:57",
					"status": "active",
					"label": "nat-gateway-59d6c282-00a7-4a92-9a41-3adad396abcd",
					"tag": "test tag",
					"public_ips": [
						"149.1.2.3"
					],
					"public_ips_v6": [
						"2001:19f0:0006:01d3:5400:05ff:fec2:abcd"
					],
					"private_ips": [
						"10.1.128.1"
					],
					"billing": {
						"charges": 20.16,
						"monthly": 20.16
					}
				}
			]
		}
		`
		fmt.Fprint(writer, response)
	})

	natGateways, _, _, err := client.VPC.ListNATGateways(ctx, "59d6c282-00a7-4a92-9a41-3adad396abcd", nil)

	if err != nil {
		t.Errorf("VPC.List returned error: %v", err)
	}

	publicIPs := []string{
		"149.1.2.3",
	}

	publicIPsV6 := []string{
		"2001:19f0:0006:01d3:5400:05ff:fec2:abcd",
	}

	privateIPs := []string{
		"10.1.128.1",
	}

	billing := NATGatewayBilling{
		Charges: 20.16,
		Monthly: 20.16,
	}

	expected := []NATGateway{
		{
			ID:          "7af46919-f6b0-4f34-8523-1d03911eabcd",
			VPCID:       "59d6c282-00a7-4a92-9a41-3adad396abcd",
			DateCreated: "2025-09-29 10:38:57",
			Status:      "active",
			Label:       "nat-gateway-59d6c282-00a7-4a92-9a41-3adad396abcd",
			Tag:         "test tag",
			PublicIPs:   publicIPs,
			PublicIPsV6: publicIPsV6,
			PrivateIPs:  privateIPs,
			Billing:     billing,
		},
	}

	if !reflect.DeepEqual(natGateways, expected) {
		t.Errorf("VPC.ListNATGateways returned %+v, expected %+v", natGateways, expected)
	}
}

func TestVPCServiceHandler_CreateNATGateway(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/59d6c282-00a7-4a92-9a41-3adad396abcd/nat-gateway", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"nat_gateway": {
				"id": "7af46919-f6b0-4f34-8523-1d03911eabcd",
				"vpc_id": "59d6c282-00a7-4a92-9a41-3adad396abcd",
				"date_created": "2025-09-29 10:38:57",
				"status": "active",
				"label": "nat-gateway-59d6c282-00a7-4a92-9a41-3adad396abcd",
				"tag": "test tag",
				"public_ips": [
					"149.1.2.3"
				],
				"public_ips_v6": [
					"2001:19f0:0006:01d3:5400:05ff:fec2:abcd"
				],
				"private_ips": [
					"10.1.128.1"
				],
				"billing": {
					"charges": 20.16,
					"monthly": 20.16
				}
			}
		}
		`

		fmt.Fprint(writer, response)
	})

	options := &NATGatewayReq{
		Label: "nat-gateway-59d6c282-00a7-4a92-9a41-3adad396abcd",
		Tag:   "test tag",
	}

	natGateway, _, err := client.VPC.CreateNATGateway(ctx, "59d6c282-00a7-4a92-9a41-3adad396abcd", options)

	if err != nil {
		t.Errorf("VPC.CreateNATGateway returned %+v, expected %+v", err, nil)
	}

	publicIPs := []string{
		"149.1.2.3",
	}

	publicIPsV6 := []string{
		"2001:19f0:0006:01d3:5400:05ff:fec2:abcd",
	}

	privateIPs := []string{
		"10.1.128.1",
	}

	billing := NATGatewayBilling{
		Charges: 20.16,
		Monthly: 20.16,
	}

	expected := &NATGateway{
		ID:          "7af46919-f6b0-4f34-8523-1d03911eabcd",
		VPCID:       "59d6c282-00a7-4a92-9a41-3adad396abcd",
		DateCreated: "2025-09-29 10:38:57",
		Status:      "active",
		Label:       "nat-gateway-59d6c282-00a7-4a92-9a41-3adad396abcd",
		Tag:         "test tag",
		PublicIPs:   publicIPs,
		PublicIPsV6: publicIPsV6,
		PrivateIPs:  privateIPs,
		Billing:     billing,
	}

	if !reflect.DeepEqual(natGateway, expected) {
		t.Errorf("VPC.CreateNATGateway returned %+v, expected %+v", natGateway, expected)
	}
}

func TestVPCServiceHandler_GetNATGateway(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/59d6c282-00a7-4a92-9a41-3adad396abcd/nat-gateway/7af46919-f6b0-4f34-8523-1d03911eabcd", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"nat_gateway": {
				"id": "7af46919-f6b0-4f34-8523-1d03911eabcd",
				"vpc_id": "59d6c282-00a7-4a92-9a41-3adad396abcd",
				"date_created": "2025-09-29 10:38:57",
				"status": "active",
				"label": "nat-gateway-59d6c282-00a7-4a92-9a41-3adad396abcd",
				"tag": "test tag",
				"public_ips": [
					"149.1.2.3"
				],
				"public_ips_v6": [
					"2001:19f0:0006:01d3:5400:05ff:fec2:abcd"
				],
				"private_ips": [
					"10.1.128.1"
				],
				"billing": {
					"charges": 20.16,
					"monthly": 20.16
				}
			}
		}
		`

		fmt.Fprint(writer, response)
	})

	natGateway, _, err := client.VPC.GetNATGateway(ctx, "59d6c282-00a7-4a92-9a41-3adad396abcd", "7af46919-f6b0-4f34-8523-1d03911eabcd")
	if err != nil {
		t.Errorf("VPC.GetNATGateway returned %+v, expected %+v", err, nil)
	}

	publicIPs := []string{
		"149.1.2.3",
	}

	publicIPsV6 := []string{
		"2001:19f0:0006:01d3:5400:05ff:fec2:abcd",
	}

	privateIPs := []string{
		"10.1.128.1",
	}

	billing := NATGatewayBilling{
		Charges: 20.16,
		Monthly: 20.16,
	}

	expected := &NATGateway{
		ID:          "7af46919-f6b0-4f34-8523-1d03911eabcd",
		VPCID:       "59d6c282-00a7-4a92-9a41-3adad396abcd",
		DateCreated: "2025-09-29 10:38:57",
		Status:      "active",
		Label:       "nat-gateway-59d6c282-00a7-4a92-9a41-3adad396abcd",
		Tag:         "test tag",
		PublicIPs:   publicIPs,
		PublicIPsV6: publicIPsV6,
		PrivateIPs:  privateIPs,
		Billing:     billing,
	}

	if !reflect.DeepEqual(natGateway, expected) {
		t.Errorf("VPC.GetNATGateway returned %+v, expected %+v", natGateway, expected)
	}
}

func TestVPCServiceHandler_UpdateNATGateway(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/59d6c282-00a7-4a92-9a41-3adad396abcd/nat-gateway/7af46919-f6b0-4f34-8523-1d03911eabcd", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"nat_gateway": {
				"id": "7af46919-f6b0-4f34-8523-1d03911eabcd",
				"vpc_id": "59d6c282-00a7-4a92-9a41-3adad396abcd",
				"date_created": "2025-09-29 10:38:57",
				"status": "active",
				"label": "label updated",
				"tag": "test tag updated",
				"public_ips": [
					"149.1.2.3"
				],
				"public_ips_v6": [
					"2001:19f0:0006:01d3:5400:05ff:fec2:abcd"
				],
				"private_ips": [
					"10.1.128.1"
				],
				"billing": {
					"charges": 20.16,
					"monthly": 20.16
				}
			}
		}
		`

		fmt.Fprint(writer, response)
	})

	options := &NATGatewayReq{
		Label: "label updated",
		Tag:   "test tag updated",
	}

	natGateway, _, err := client.VPC.UpdateNATGateway(ctx, "59d6c282-00a7-4a92-9a41-3adad396abcd", "7af46919-f6b0-4f34-8523-1d03911eabcd", options)

	if err != nil {
		t.Errorf("VPC.UpdateNATGateway returned %+v, expected %+v", err, nil)
	}

	publicIPs := []string{
		"149.1.2.3",
	}

	publicIPsV6 := []string{
		"2001:19f0:0006:01d3:5400:05ff:fec2:abcd",
	}

	privateIPs := []string{
		"10.1.128.1",
	}

	billing := NATGatewayBilling{
		Charges: 20.16,
		Monthly: 20.16,
	}

	expected := &NATGateway{
		ID:          "7af46919-f6b0-4f34-8523-1d03911eabcd",
		VPCID:       "59d6c282-00a7-4a92-9a41-3adad396abcd",
		DateCreated: "2025-09-29 10:38:57",
		Status:      "active",
		Label:       "label updated",
		Tag:         "test tag updated",
		PublicIPs:   publicIPs,
		PublicIPsV6: publicIPsV6,
		PrivateIPs:  privateIPs,
		Billing:     billing,
	}

	if !reflect.DeepEqual(natGateway, expected) {
		t.Errorf("VPC.UpdateNATGateway returned %+v, expected %+v", natGateway, expected)
	}
}

func TestVPCServiceHandler_DeleteNATGateway(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/59d6c282-00a7-4a92-9a41-3adad396abcd/nat-gateway/7af46919-f6b0-4f34-8523-1d03911eabcd", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.VPC.DeleteNATGateway(ctx, "59d6c282-00a7-4a92-9a41-3adad396abcd", "7af46919-f6b0-4f34-8523-1d03911eabcd")

	if err != nil {
		t.Errorf("VPC.DeleteNATGateway returned %+v, expected %+v", err, nil)
	}
}

func TestVPCServiceHandler_NATGatewayListFirewallRules(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/59d6c282-00a7-4a92-9a41-3adad396abcd/nat-gateway/7af46919-f6b0-4f34-8523-1d03911eabcd/global/firewall-rules", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"meta": {
				"total": 1,
				"links": {
					"next": "",
					"prev": ""
				}
			},
			"firewall_rules": [
				{
					"id": "822043f4-b135-470d-89d0-58476498abcd",
					"action": "accept",
					"protocol": "tcp",
					"port": "655",
					"subnet": "1.2.3.4",
					"subnet_size": 24,
					"notes": "test rule"
				}
			]
		}
		`
		fmt.Fprint(writer, response)
	})

	firewallRules, _, _, err := client.VPC.ListNATGatewayFirewallRules(ctx, "59d6c282-00a7-4a92-9a41-3adad396abcd", "7af46919-f6b0-4f34-8523-1d03911eabcd", nil)

	if err != nil {
		t.Errorf("VPC.ListNATGatewayFirewallRules returned error: %v", err)
	}

	expected := []NATGatewayFirewallRule{
		{
			ID:         "822043f4-b135-470d-89d0-58476498abcd",
			Action:     "accept",
			Protocol:   "tcp",
			Port:       "655",
			Subnet:     "1.2.3.4",
			SubnetSize: 24,
			Notes:      "test rule",
		},
	}

	if !reflect.DeepEqual(firewallRules, expected) {
		t.Errorf("VPC.ListNATGatewayFirewallRules returned %+v, expected %+v", firewallRules, expected)
	}
}

func TestVPCServiceHandler_CreateNATGatewayFirewallRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/59d6c282-00a7-4a92-9a41-3adad396abcd/nat-gateway/7af46919-f6b0-4f34-8523-1d03911eabcd/global/firewall-rules", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"firewall_rule": {
				"id": "822043f4-b135-470d-89d0-58476498abcd",
				"action": "accept",
				"protocol": "tcp",
				"port": "655",
				"subnet": "1.2.3.4",
				"subnet_size": 24,
				"notes": "test rule"
			}
		}
		`

		fmt.Fprint(writer, response)
	})

	options := &NATGatewayFirewallRuleCreateReq{
		Protocol:   "tcp",
		Port:       "655",
		Subnet:     "1.2.3.4",
		SubnetSize: 24,
		Notes:      "test rule",
	}

	firewallRule, _, err := client.VPC.CreateNATGatewayFirewallRule(ctx, "59d6c282-00a7-4a92-9a41-3adad396abcd", "7af46919-f6b0-4f34-8523-1d03911eabcd", options)

	if err != nil {
		t.Errorf("VPC.CreateNATGatewayFirewallRule returned %+v, expected %+v", err, nil)
	}

	expected := &NATGatewayFirewallRule{
		ID:         "822043f4-b135-470d-89d0-58476498abcd",
		Action:     "accept",
		Protocol:   "tcp",
		Port:       "655",
		Subnet:     "1.2.3.4",
		SubnetSize: 24,
		Notes:      "test rule",
	}

	if !reflect.DeepEqual(firewallRule, expected) {
		t.Errorf("VPC.CreateNATGatewayFirewallRule returned %+v, expected %+v", firewallRule, expected)
	}
}

func TestVPCServiceHandler_GetNATGatewayFirewallRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/59d6c282-00a7-4a92-9a41-3adad396abcd/nat-gateway/7af46919-f6b0-4f34-8523-1d03911eabcd/global/firewall-rules/822043f4-b135-470d-89d0-58476498abcd", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"firewall_rule": {
				"id": "822043f4-b135-470d-89d0-58476498abcd",
				"action": "accept",
				"protocol": "tcp",
				"port": "655",
				"subnet": "1.2.3.4",
				"subnet_size": 24,
				"notes": "test rule"
			}
		}
		`

		fmt.Fprint(writer, response)
	})

	firewallRule, _, err := client.VPC.GetNATGatewayFirewallRule(ctx, "59d6c282-00a7-4a92-9a41-3adad396abcd", "7af46919-f6b0-4f34-8523-1d03911eabcd", "822043f4-b135-470d-89d0-58476498abcd")
	if err != nil {
		t.Errorf("VPC.GetNATGatewayFirewallRule returned %+v, expected %+v", err, nil)
	}

	expected := &NATGatewayFirewallRule{
		ID:         "822043f4-b135-470d-89d0-58476498abcd",
		Action:     "accept",
		Protocol:   "tcp",
		Port:       "655",
		Subnet:     "1.2.3.4",
		SubnetSize: 24,
		Notes:      "test rule",
	}

	if !reflect.DeepEqual(firewallRule, expected) {
		t.Errorf("VPC.GetNATGatewayFirewallRule returned %+v, expected %+v", firewallRule, expected)
	}
}

func TestVPCServiceHandler_UpdateNATGatewayFirewallRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/59d6c282-00a7-4a92-9a41-3adad396abcd/nat-gateway/7af46919-f6b0-4f34-8523-1d03911eabcd/global/firewall-rules/822043f4-b135-470d-89d0-58476498abcd", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"firewall_rule": {
				"id": "822043f4-b135-470d-89d0-58476498abcd",
				"action": "accept",
				"protocol": "tcp",
				"port": "655",
				"subnet": "1.2.3.4",
				"subnet_size": 24,
				"notes": "test rule updated"
			}
		}
		`

		fmt.Fprint(writer, response)
	})

	options := &NATGatewayFirewallRuleUpdateReq{
		Notes: "test rule updated",
	}

	firewallRule, _, err := client.VPC.UpdateNATGatewayFirewallRule(ctx, "59d6c282-00a7-4a92-9a41-3adad396abcd", "7af46919-f6b0-4f34-8523-1d03911eabcd", "822043f4-b135-470d-89d0-58476498abcd", options)

	if err != nil {
		t.Errorf("VPC.UpdateNATGatewayFirewallRule returned %+v, expected %+v", err, nil)
	}

	expected := &NATGatewayFirewallRule{
		ID:         "822043f4-b135-470d-89d0-58476498abcd",
		Action:     "accept",
		Protocol:   "tcp",
		Port:       "655",
		Subnet:     "1.2.3.4",
		SubnetSize: 24,
		Notes:      "test rule updated",
	}

	if !reflect.DeepEqual(firewallRule, expected) {
		t.Errorf("VPC.UpdateNATGatewayFirewallRule returned %+v, expected %+v", firewallRule, expected)
	}
}

func TestVPCServiceHandler_DeleteNATGatewayFirewallRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/59d6c282-00a7-4a92-9a41-3adad396abcd/nat-gateway/7af46919-f6b0-4f34-8523-1d03911eabcd/global/firewall-rules/822043f4-b135-470d-89d0-58476498abcd", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.VPC.DeleteNATGatewayFirewallRule(ctx, "59d6c282-00a7-4a92-9a41-3adad396abcd", "7af46919-f6b0-4f34-8523-1d03911eabcd", "822043f4-b135-470d-89d0-58476498abcd")

	if err != nil {
		t.Errorf("VPC.DeleteNATGatewayFirewallRule returned %+v, expected %+v", err, nil)
	}
}

func TestVPCServiceHandler_NATGatewayListPortForwardingRules(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/59d6c282-00a7-4a92-9a41-3adad396abcd/nat-gateway/7af46919-f6b0-4f34-8523-1d03911eabcd/global/port-forwarding-rules", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"meta": {
				"total": 1,
				"links": {
					"next": "",
					"prev": ""
				}
			},
			"port_forwarding_rules": [
				{
					"id": "e0116495-7657-4790-9801-e93157fcabcd",
					"name": "test rule",
					"protocol": "tcp",
					"external_port": 655,
					"internal_ip": "10.1.2.3",
					"internal_port": 123,
					"enabled": true,
					"description": "test",
					"created_at": "2025-10-16 15:09:26",
					"updated_at": "2025-10-16 15:09:26"
				}
			]
		}
		`
		fmt.Fprint(writer, response)
	})

	portForwardingRules, _, _, err := client.VPC.ListNATGatewayPortForwardingRules(ctx, "59d6c282-00a7-4a92-9a41-3adad396abcd", "7af46919-f6b0-4f34-8523-1d03911eabcd", nil)

	if err != nil {
		t.Errorf("VPC.ListNATGatewayPortForwardingRules returned error: %v", err)
	}

	expected := []NATGatewayPortForwardingRule{
		{
			ID:           "e0116495-7657-4790-9801-e93157fcabcd",
			Name:         "test rule",
			Protocol:     "tcp",
			ExternalPort: 655,
			InternalIP:   "10.1.2.3",
			InternalPort: 123,
			Enabled:      BoolToBoolPtr(true),
			Description:  "test",
			DateCreated:  "2025-10-16 15:09:26",
			DateUpdated:  "2025-10-16 15:09:26",
		},
	}

	if !reflect.DeepEqual(portForwardingRules, expected) {
		t.Errorf("VPC.ListNATGatewayPortForwardingRules returned %+v, expected %+v", portForwardingRules, expected)
	}
}

func TestVPCServiceHandler_CreateNATGatewayPortForwardingRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/59d6c282-00a7-4a92-9a41-3adad396abcd/nat-gateway/7af46919-f6b0-4f34-8523-1d03911eabcd/global/port-forwarding-rules", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"port_forwarding_rule": {
				"id": "e0116495-7657-4790-9801-e93157fcabcd",
				"name": "test rule",
				"protocol": "tcp",
				"external_port": 655,
				"internal_ip": "10.1.2.3",
				"internal_port": 123,
				"enabled": true,
				"description": "test",
				"created_at": "2025-10-16 15:09:26",
				"updated_at": "2025-10-16 15:09:26"
			}
		}
		`

		fmt.Fprint(writer, response)
	})

	options := &NATGatewayPortForwardingRuleReq{
		Name:         "test rule",
		Protocol:     "tcp",
		ExternalPort: 655,
		InternalIP:   "10.1.2.3",
		InternalPort: 123,
		Enabled:      BoolToBoolPtr(true),
		Description:  "test",
	}

	portForwardingRule, _, err := client.VPC.CreateNATGatewayPortForwardingRule(ctx, "59d6c282-00a7-4a92-9a41-3adad396abcd", "7af46919-f6b0-4f34-8523-1d03911eabcd", options)

	if err != nil {
		t.Errorf("VPC.CreateNATGatewayPortForwardingRule returned %+v, expected %+v", err, nil)
	}

	expected := &NATGatewayPortForwardingRule{
		ID:           "e0116495-7657-4790-9801-e93157fcabcd",
		Name:         "test rule",
		Protocol:     "tcp",
		ExternalPort: 655,
		InternalIP:   "10.1.2.3",
		InternalPort: 123,
		Enabled:      BoolToBoolPtr(true),
		Description:  "test",
		DateCreated:  "2025-10-16 15:09:26",
		DateUpdated:  "2025-10-16 15:09:26",
	}

	if !reflect.DeepEqual(portForwardingRule, expected) {
		t.Errorf("VPC.CreateNATGatewayPortForwardingRule returned %+v, expected %+v", portForwardingRule, expected)
	}
}

func TestVPCServiceHandler_GetNATGatewayPortForwardingRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/59d6c282-00a7-4a92-9a41-3adad396abcd/nat-gateway/7af46919-f6b0-4f34-8523-1d03911eabcd/global/port-forwarding-rules/e0116495-7657-4790-9801-e93157fcabcd", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"port_forwarding_rule": {
				"id": "e0116495-7657-4790-9801-e93157fcabcd",
				"name": "test rule",
				"protocol": "tcp",
				"external_port": 655,
				"internal_ip": "10.1.2.3",
				"internal_port": 123,
				"enabled": true,
				"description": "test",
				"created_at": "2025-10-16 15:09:26",
				"updated_at": "2025-10-16 15:09:26"
			}
		}
		`

		fmt.Fprint(writer, response)
	})

	portForwardingRule, _, err := client.VPC.GetNATGatewayPortForwardingRule(ctx, "59d6c282-00a7-4a92-9a41-3adad396abcd", "7af46919-f6b0-4f34-8523-1d03911eabcd", "e0116495-7657-4790-9801-e93157fcabcd")
	if err != nil {
		t.Errorf("VPC.GetNATGatewayPortForwardingRule returned %+v, expected %+v", err, nil)
	}

	expected := &NATGatewayPortForwardingRule{
		ID:           "e0116495-7657-4790-9801-e93157fcabcd",
		Name:         "test rule",
		Protocol:     "tcp",
		ExternalPort: 655,
		InternalIP:   "10.1.2.3",
		InternalPort: 123,
		Enabled:      BoolToBoolPtr(true),
		Description:  "test",
		DateCreated:  "2025-10-16 15:09:26",
		DateUpdated:  "2025-10-16 15:09:26",
	}

	if !reflect.DeepEqual(portForwardingRule, expected) {
		t.Errorf("VPC.GetNATGatewayPortForwardingRule returned %+v, expected %+v", portForwardingRule, expected)
	}
}

func TestVPCServiceHandler_UpdateNATGatewayPortForwardingRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/59d6c282-00a7-4a92-9a41-3adad396abcd/nat-gateway/7af46919-f6b0-4f34-8523-1d03911eabcd/global/port-forwarding-rules/e0116495-7657-4790-9801-e93157fcabcd", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"port_forwarding_rule": {
				"id": "e0116495-7657-4790-9801-e93157fcabcd",
				"name": "test rule updated",
				"protocol": "tcp",
				"external_port": 655,
				"internal_ip": "10.1.2.3",
				"internal_port": 123,
				"enabled": true,
				"description": "test updated",
				"created_at": "2025-10-16 15:09:26",
				"updated_at": "2025-10-16 15:09:26"
			}
		}
		`

		fmt.Fprint(writer, response)
	})

	options := &NATGatewayPortForwardingRuleReq{
		Name:        "test rule updated",
		Description: "test updated",
	}

	portForwardingRule, _, err := client.VPC.UpdateNATGatewayPortForwardingRule(ctx, "59d6c282-00a7-4a92-9a41-3adad396abcd", "7af46919-f6b0-4f34-8523-1d03911eabcd", "e0116495-7657-4790-9801-e93157fcabcd", options)

	if err != nil {
		t.Errorf("VPC.UpdateNATGatewayPortForwardingRule returned %+v, expected %+v", err, nil)
	}

	expected := &NATGatewayPortForwardingRule{
		ID:           "e0116495-7657-4790-9801-e93157fcabcd",
		Name:         "test rule updated",
		Protocol:     "tcp",
		ExternalPort: 655,
		InternalIP:   "10.1.2.3",
		InternalPort: 123,
		Enabled:      BoolToBoolPtr(true),
		Description:  "test updated",
		DateCreated:  "2025-10-16 15:09:26",
		DateUpdated:  "2025-10-16 15:09:26",
	}

	if !reflect.DeepEqual(portForwardingRule, expected) {
		t.Errorf("VPC.UpdateNATGatewayPortForwardingRule returned %+v, expected %+v", portForwardingRule, expected)
	}
}

func TestVPCServiceHandler_DeleteNATGatewayPortForwardingRule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpcs/59d6c282-00a7-4a92-9a41-3adad396abcd/nat-gateway/7af46919-f6b0-4f34-8523-1d03911eabcd/global/port-forwarding-rules/e0116495-7657-4790-9801-e93157fcabcd", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.VPC.DeleteNATGatewayPortForwardingRule(ctx, "59d6c282-00a7-4a92-9a41-3adad396abcd", "7af46919-f6b0-4f34-8523-1d03911eabcd", "e0116495-7657-4790-9801-e93157fcabcd")

	if err != nil {
		t.Errorf("VPC.DeleteNATGatewayPortForwardingRule returned %+v, expected %+v", err, nil)
	}
}
