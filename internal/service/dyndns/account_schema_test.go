package dyndns

import (
	"testing"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dyndns"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestConvertAccountSchemaToStruct(t *testing.T) {
	tests := []struct {
		name     string
		input    *accountResourceModel
		expected *dyndns.Account
	}{
		{
			name: "basic conversion",
			input: &accountResourceModel{
				Enabled:        types.BoolValue(true),
				Service:        types.StringValue("cloudflare"),
				Protocol:       types.StringValue("dyndns2"),
				Server:         types.StringValue("api.cloudflare.com"),
				Username:       types.StringValue("user@example.com"),
				Password:       types.StringValue("secret123"),
				ResourceId:     types.StringValue("zone-id-123"),
				Hostnames:      types.SetValueMust(types.StringType, []attr.Value{types.StringValue("example.com")}),
				Wildcard:       types.BoolValue(true),
				Zone:           types.StringValue("example.com"),
				Checkip:        types.StringValue("web_dyndns"),
				DynIpv6Host:    types.StringValue(""),
				CheckipTimeout: types.StringValue("10"),
				ForceSsl:       types.BoolValue(true),
				Ttl:            types.StringValue("300"),
				Interface:      types.StringValue("wan"),
				Description:    types.StringValue("Test DynDNS Account"),
				Id:             types.StringValue("uuid-123"),
			},
			expected: &dyndns.Account{
				Enabled:        "1",
				Service:        api.SelectedMap("cloudflare"),
				Protocol:       api.SelectedMap("dyndns2"),
				Server:         "api.cloudflare.com",
				Username:       "user@example.com",
				Password:       "secret123",
				ResourceId:     "zone-id-123",
				Hostnames:      api.SelectedMapList([]string{"example.com"}),
				Wildcard:       "1",
				Zone:           "example.com",
				Checkip:        api.SelectedMap("web_dyndns"),
				DynIpv6Host:    "",
				CheckipTimeout: "10",
				ForceSsl:       "1",
				Ttl:            "300",
				Interface:      api.SelectedMap("wan"),
				Description:    "Test DynDNS Account",
			},
		},
		{
			name: "disabled with empty optional fields",
			input: &accountResourceModel{
				Enabled:        types.BoolValue(false),
				Service:        types.StringValue("noip"),
				Protocol:       types.StringValue(""),
				Server:         types.StringNull(),
				Username:       types.StringNull(),
				Password:       types.StringNull(),
				ResourceId:     types.StringNull(),
				Hostnames:      types.SetValueMust(types.StringType, []attr.Value{}),
				Wildcard:       types.BoolValue(false),
				Zone:           types.StringNull(),
				Checkip:        types.StringValue(""),
				DynIpv6Host:    types.StringNull(),
				CheckipTimeout: types.StringNull(),
				ForceSsl:       types.BoolValue(false),
				Ttl:            types.StringNull(),
				Interface:      types.StringValue(""),
				Description:    types.StringNull(),
				Id:             types.StringValue("uuid-456"),
			},
			expected: &dyndns.Account{
				Enabled:        "0",
				Service:        api.SelectedMap("noip"),
				Protocol:       api.SelectedMap(""),
				Server:         "",
				Username:       "",
				Password:       "",
				ResourceId:     "",
				Hostnames:      api.SelectedMapList([]string{}),
				Wildcard:       "0",
				Zone:           "",
				Checkip:        api.SelectedMap(""),
				DynIpv6Host:    "",
				CheckipTimeout: "",
				ForceSsl:       "0",
				Ttl:            "",
				Interface:      api.SelectedMap(""),
				Description:    "",
			},
		},
		{
			name: "multiple hostnames",
			input: &accountResourceModel{
				Enabled:        types.BoolValue(true),
				Service:        types.StringValue("dyndns2"),
				Protocol:       types.StringValue("dyndns2"),
				Server:         types.StringValue("members.dyndns.org"),
				Username:       types.StringValue("myuser"),
				Password:       types.StringValue("mypass"),
				ResourceId:     types.StringValue(""),
				Hostnames: types.SetValueMust(types.StringType, []attr.Value{
					types.StringValue("host1.example.com"),
					types.StringValue("host2.example.com"),
				}),
				Wildcard:       types.BoolValue(false),
				Zone:           types.StringValue(""),
				Checkip:        types.StringValue("web_dyndns"),
				DynIpv6Host:    types.StringValue(""),
				CheckipTimeout: types.StringValue("30"),
				ForceSsl:       types.BoolValue(false),
				Ttl:            types.StringValue(""),
				Interface:      types.StringValue("lan"),
				Description:    types.StringValue(""),
				Id:             types.StringValue("uuid-789"),
			},
			expected: &dyndns.Account{
				Enabled:        "1",
				Service:        api.SelectedMap("dyndns2"),
				Protocol:       api.SelectedMap("dyndns2"),
				Server:         "members.dyndns.org",
				Username:       "myuser",
				Password:       "mypass",
				ResourceId:     "",
				Hostnames: api.SelectedMapList([]string{
					"host1.example.com",
					"host2.example.com",
				}),
				Wildcard:       "0",
				Zone:           "",
				Checkip:        api.SelectedMap("web_dyndns"),
				DynIpv6Host:    "",
				CheckipTimeout: "30",
				ForceSsl:       "0",
				Ttl:            "",
				Interface:      api.SelectedMap("lan"),
				Description:    "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := convertAccountSchemaToStruct(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected.Enabled, result.Enabled)
			assert.Equal(t, string(tt.expected.Service), string(result.Service))
			assert.Equal(t, string(tt.expected.Protocol), string(result.Protocol))
			assert.Equal(t, tt.expected.Server, result.Server)
			assert.Equal(t, tt.expected.Username, result.Username)
			assert.Equal(t, tt.expected.Password, result.Password)
			assert.Equal(t, tt.expected.ResourceId, result.ResourceId)
			assert.Equal(t, []string(tt.expected.Hostnames), []string(result.Hostnames))
			assert.Equal(t, tt.expected.Wildcard, result.Wildcard)
			assert.Equal(t, tt.expected.Zone, result.Zone)
			assert.Equal(t, string(tt.expected.Checkip), string(result.Checkip))
			assert.Equal(t, tt.expected.DynIpv6Host, result.DynIpv6Host)
			assert.Equal(t, tt.expected.CheckipTimeout, result.CheckipTimeout)
			assert.Equal(t, tt.expected.ForceSsl, result.ForceSsl)
			assert.Equal(t, tt.expected.Ttl, result.Ttl)
			assert.Equal(t, string(tt.expected.Interface), string(result.Interface))
			assert.Equal(t, tt.expected.Description, result.Description)
		})
	}
}

func TestConvertAccountStructToSchema(t *testing.T) {
	tests := []struct {
		name     string
		input    *dyndns.Account
		expected *accountResourceModel
	}{
		{
			name: "basic conversion",
			input: &dyndns.Account{
				Enabled:        "1",
				Service:        api.SelectedMap("cloudflare"),
				Protocol:       api.SelectedMap("dyndns2"),
				Server:         "api.cloudflare.com",
				Username:       "user@example.com",
				Password:       "",
				ResourceId:     "zone-id-123",
				Hostnames:      api.SelectedMapList([]string{"example.com"}),
				Wildcard:       "1",
				Zone:           "example.com",
				Checkip:        api.SelectedMap("web_dyndns"),
				DynIpv6Host:    "",
				CheckipTimeout: "10",
				ForceSsl:       "1",
				Ttl:            "300",
				Interface:      api.SelectedMap("wan"),
				Description:    "Test DynDNS Account",
				CurrentIp:      "1.2.3.4",
				CurrentMtime:   "2024-01-01 00:00:00",
			},
			expected: &accountResourceModel{
				Enabled:        types.BoolValue(true),
				Service:        types.StringValue("cloudflare"),
				Protocol:       types.StringValue("dyndns2"),
				Server:         types.StringValue("api.cloudflare.com"),
				Username:       types.StringValue("user@example.com"),
				Password:       types.StringNull(),
				ResourceId:     types.StringValue("zone-id-123"),
				Hostnames:      types.SetValueMust(types.StringType, []attr.Value{types.StringValue("example.com")}),
				Wildcard:       types.BoolValue(true),
				Zone:           types.StringValue("example.com"),
				Checkip:        types.StringValue("web_dyndns"),
				DynIpv6Host:    types.StringNull(),
				CheckipTimeout: types.StringValue("10"),
				ForceSsl:       types.BoolValue(true),
				Ttl:            types.StringValue("300"),
				Interface:      types.StringValue("wan"),
				Description:    types.StringValue("Test DynDNS Account"),
				CurrentIp:      types.StringValue("1.2.3.4"),
				CurrentMtime:   types.StringValue("2024-01-01 00:00:00"),
			},
		},
		{
			name: "disabled with empty fields",
			input: &dyndns.Account{
				Enabled:        "0",
				Service:        api.SelectedMap("noip"),
				Protocol:       api.SelectedMap(""),
				Server:         "",
				Username:       "",
				Password:       "",
				ResourceId:     "",
				Hostnames:      api.SelectedMapList([]string{}),
				Wildcard:       "0",
				Zone:           "",
				Checkip:        api.SelectedMap(""),
				DynIpv6Host:    "",
				CheckipTimeout: "",
				ForceSsl:       "0",
				Ttl:            "",
				Interface:      api.SelectedMap(""),
				Description:    "",
				CurrentIp:      "",
				CurrentMtime:   "",
			},
			expected: &accountResourceModel{
				Enabled:        types.BoolValue(false),
				Service:        types.StringValue("noip"),
				Protocol:       types.StringValue(""),
				Server:         types.StringNull(),
				Username:       types.StringNull(),
				Password:       types.StringNull(),
				ResourceId:     types.StringNull(),
				Hostnames:      types.SetValueMust(types.StringType, []attr.Value{}),
				Wildcard:       types.BoolValue(false),
				Zone:           types.StringNull(),
				Checkip:        types.StringValue(""),
				DynIpv6Host:    types.StringNull(),
				CheckipTimeout: types.StringNull(),
				ForceSsl:       types.BoolValue(false),
				Ttl:            types.StringNull(),
				Interface:      types.StringValue(""),
				Description:    types.StringNull(),
				CurrentIp:      types.StringNull(),
				CurrentMtime:   types.StringNull(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := convertAccountStructToSchema(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected.Enabled, result.Enabled)
			assert.Equal(t, tt.expected.Service, result.Service)
			assert.Equal(t, tt.expected.Protocol, result.Protocol)
			assert.Equal(t, tt.expected.Server, result.Server)
			assert.Equal(t, tt.expected.Username, result.Username)
			assert.Equal(t, tt.expected.Password, result.Password)
			assert.Equal(t, tt.expected.ResourceId, result.ResourceId)
			assert.Equal(t, tt.expected.Wildcard, result.Wildcard)
			assert.Equal(t, tt.expected.Zone, result.Zone)
			assert.Equal(t, tt.expected.Checkip, result.Checkip)
			assert.Equal(t, tt.expected.DynIpv6Host, result.DynIpv6Host)
			assert.Equal(t, tt.expected.CheckipTimeout, result.CheckipTimeout)
			assert.Equal(t, tt.expected.ForceSsl, result.ForceSsl)
			assert.Equal(t, tt.expected.Ttl, result.Ttl)
			assert.Equal(t, tt.expected.Interface, result.Interface)
			assert.Equal(t, tt.expected.Description, result.Description)
			assert.Equal(t, tt.expected.CurrentIp, result.CurrentIp)
			assert.Equal(t, tt.expected.CurrentMtime, result.CurrentMtime)
		})
	}
}
