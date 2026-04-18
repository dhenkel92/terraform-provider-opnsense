package dyndns

import (
	"context"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dyndns"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// accountResourceModel describes the resource data model.
type accountResourceModel struct {
	Enabled        types.Bool   `tfsdk:"enabled"`
	Service        types.String `tfsdk:"service"`
	Protocol       types.String `tfsdk:"protocol"`
	Server         types.String `tfsdk:"server"`
	Username       types.String `tfsdk:"username"`
	Password       types.String `tfsdk:"password"`
	ResourceId     types.String `tfsdk:"resource_id"`
	Hostnames      types.Set    `tfsdk:"hostnames"`
	Wildcard       types.Bool   `tfsdk:"wildcard"`
	Zone           types.String `tfsdk:"zone"`
	Checkip        types.String `tfsdk:"checkip"`
	DynIpv6Host    types.String `tfsdk:"dyn_ipv6_host"`
	CheckipTimeout types.String `tfsdk:"checkip_timeout"`
	ForceSsl       types.Bool   `tfsdk:"force_ssl"`
	Ttl            types.String `tfsdk:"ttl"`
	Interface      types.String `tfsdk:"interface"`
	Description    types.String `tfsdk:"description"`
	CurrentIp      types.String `tfsdk:"current_ip"`
	CurrentMtime   types.String `tfsdk:"current_mtime"`

	Id types.String `tfsdk:"id"`
}

func accountResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages dynamic DNS account configurations in OPNsense.",

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this dynamic DNS account. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"service": schema.StringAttribute{
				MarkdownDescription: "The dynamic DNS service type (e.g. `cloudflare`, `dyndns2`, `noip`, etc.).",
				Required:            true,
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: "The dynamic DNS protocol to use. Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"server": schema.StringAttribute{
				MarkdownDescription: "The server hostname or URL for the dynamic DNS service.",
				Optional:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "Username for authentication with the dynamic DNS service.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password for authentication with the dynamic DNS service.",
				Optional:            true,
				Sensitive:           true,
			},
			"resource_id": schema.StringAttribute{
				MarkdownDescription: "The resource identifier for the dynamic DNS service (e.g. zone ID for Cloudflare).",
				Optional:            true,
			},
			"hostnames": schema.SetAttribute{
				MarkdownDescription: "Set of hostnames to update. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"wildcard": schema.BoolAttribute{
				MarkdownDescription: "Enable wildcard DNS record. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"zone": schema.StringAttribute{
				MarkdownDescription: "The DNS zone for the dynamic DNS record.",
				Optional:            true,
			},
			"checkip": schema.StringAttribute{
				MarkdownDescription: "The service to use for checking the current public IP address. Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"dyn_ipv6_host": schema.StringAttribute{
				MarkdownDescription: "The IPv6 host for dynamic DNS.",
				Optional:            true,
			},
			"checkip_timeout": schema.StringAttribute{
				MarkdownDescription: "Timeout in seconds for the IP check service.",
				Optional:            true,
			},
			"force_ssl": schema.BoolAttribute{
				MarkdownDescription: "Force SSL for dynamic DNS updates. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ttl": schema.StringAttribute{
				MarkdownDescription: "The TTL value for the DNS record.",
				Optional:            true,
			},
			"interface": schema.StringAttribute{
				MarkdownDescription: "The interface to monitor for IP changes. Defaults to `\"\"`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description for this dynamic DNS account.",
				Optional:            true,
			},
			"current_ip": schema.StringAttribute{
				MarkdownDescription: "The current IP address reported by the dynamic DNS service.",
				Computed:            true,
			},
			"current_mtime": schema.StringAttribute{
				MarkdownDescription: "The last modification time of the current IP address.",
				Computed:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func accountDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Reads a dynamic DNS account configuration from OPNsense.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"enabled": dschema.BoolAttribute{
				MarkdownDescription: "Whether this dynamic DNS account is enabled.",
				Computed:            true,
			},
			"service": dschema.StringAttribute{
				MarkdownDescription: "The dynamic DNS service type.",
				Computed:            true,
			},
			"protocol": dschema.StringAttribute{
				MarkdownDescription: "The dynamic DNS protocol.",
				Computed:            true,
			},
			"server": dschema.StringAttribute{
				MarkdownDescription: "The server hostname or URL for the dynamic DNS service.",
				Computed:            true,
			},
			"username": dschema.StringAttribute{
				MarkdownDescription: "Username for authentication with the dynamic DNS service.",
				Computed:            true,
			},
			"password": dschema.StringAttribute{
				MarkdownDescription: "Password for authentication with the dynamic DNS service. Not returned by the API.",
				Computed:            true,
				Sensitive:           true,
			},
			"resource_id": dschema.StringAttribute{
				MarkdownDescription: "The resource identifier for the dynamic DNS service.",
				Computed:            true,
			},
			"hostnames": dschema.SetAttribute{
				MarkdownDescription: "Set of hostnames to update.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"wildcard": dschema.BoolAttribute{
				MarkdownDescription: "Whether wildcard DNS record is enabled.",
				Computed:            true,
			},
			"zone": dschema.StringAttribute{
				MarkdownDescription: "The DNS zone for the dynamic DNS record.",
				Computed:            true,
			},
			"checkip": dschema.StringAttribute{
				MarkdownDescription: "The service used for checking the current public IP address.",
				Computed:            true,
			},
			"dyn_ipv6_host": dschema.StringAttribute{
				MarkdownDescription: "The IPv6 host for dynamic DNS.",
				Computed:            true,
			},
			"checkip_timeout": dschema.StringAttribute{
				MarkdownDescription: "Timeout in seconds for the IP check service.",
				Computed:            true,
			},
			"force_ssl": dschema.BoolAttribute{
				MarkdownDescription: "Whether SSL is forced for dynamic DNS updates.",
				Computed:            true,
			},
			"ttl": dschema.StringAttribute{
				MarkdownDescription: "The TTL value for the DNS record.",
				Computed:            true,
			},
			"interface": dschema.StringAttribute{
				MarkdownDescription: "The interface monitored for IP changes.",
				Computed:            true,
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Description for this dynamic DNS account.",
				Computed:            true,
			},
			"current_ip": dschema.StringAttribute{
				MarkdownDescription: "The current IP address reported by the dynamic DNS service.",
				Computed:            true,
			},
			"current_mtime": dschema.StringAttribute{
				MarkdownDescription: "The last modification time of the current IP address.",
				Computed:            true,
			},
		},
	}
}

func convertAccountSchemaToStruct(d *accountResourceModel) (*dyndns.Account, error) {
	// Parse 'Hostnames'
	var hostnamesList []string
	d.Hostnames.ElementsAs(context.Background(), &hostnamesList, false)

	return &dyndns.Account{
		Enabled:        tools.BoolToString(d.Enabled.ValueBool()),
		Service:        api.SelectedMap(d.Service.ValueString()),
		Protocol:       api.SelectedMap(d.Protocol.ValueString()),
		Server:         d.Server.ValueString(),
		Username:       d.Username.ValueString(),
		Password:       d.Password.ValueString(),
		ResourceId:     d.ResourceId.ValueString(),
		Hostnames:      api.SelectedMapList(hostnamesList),
		Wildcard:       tools.BoolToString(d.Wildcard.ValueBool()),
		Zone:           d.Zone.ValueString(),
		Checkip:        api.SelectedMap(d.Checkip.ValueString()),
		DynIpv6Host:    d.DynIpv6Host.ValueString(),
		CheckipTimeout: d.CheckipTimeout.ValueString(),
		ForceSsl:       tools.BoolToString(d.ForceSsl.ValueBool()),
		Ttl:            d.Ttl.ValueString(),
		Interface:      api.SelectedMap(d.Interface.ValueString()),
		Description:    d.Description.ValueString(),
	}, nil
}

func convertAccountStructToSchema(d *dyndns.Account) (*accountResourceModel, error) {
	model := &accountResourceModel{
		Enabled:        types.BoolValue(tools.StringToBool(d.Enabled)),
		Service:        types.StringValue(d.Service.String()),
		Protocol:       types.StringValue(d.Protocol.String()),
		Server:         tools.StringOrNull(d.Server),
		Username:       tools.StringOrNull(d.Username),
		Password:       types.StringNull(),
		ResourceId:     tools.StringOrNull(d.ResourceId),
		Hostnames:      tools.StringSliceToSet(d.Hostnames),
		Wildcard:       types.BoolValue(tools.StringToBool(d.Wildcard)),
		Zone:           tools.StringOrNull(d.Zone),
		Checkip:        types.StringValue(d.Checkip.String()),
		DynIpv6Host:    tools.StringOrNull(d.DynIpv6Host),
		CheckipTimeout: tools.StringOrNull(d.CheckipTimeout),
		ForceSsl:       types.BoolValue(tools.StringToBool(d.ForceSsl)),
		Ttl:            tools.StringOrNull(d.Ttl),
		Interface:      types.StringValue(d.Interface.String()),
		Description:    tools.StringOrNull(d.Description),
		CurrentIp:      tools.StringOrNull(d.CurrentIp),
		CurrentMtime:   tools.StringOrNull(d.CurrentMtime),
	}

	return model, nil
}
