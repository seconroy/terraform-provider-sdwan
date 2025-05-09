//go:build ignore
// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0


package provider

// Section below is generated&owned by "gen/generator.go". //template:begin provider
import (
	"context"
	"os"
	"strconv"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-sdwan"
)

// SdwanProvider defines the provider implementation.
type SdwanProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// SdwanProviderModel describes the provider data model.
type SdwanProviderModel struct {
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
	URL      types.String `tfsdk:"url"`
	Insecure types.Bool   `tfsdk:"insecure"`
	Retries  types.Int64  `tfsdk:"retries"`
}

// SdwanProviderData describes the data maintained by the provider.
type SdwanProviderData struct {
	Client *sdwan.Client
	UpdateMutex *sync.Mutex
}

// Metadata returns the provider type name.
func (p *SdwanProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "sdwan"
	resp.Version = p.version
}

func (p *SdwanProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				MarkdownDescription: "Username for the SD-WAN Manager account. This can also be set as the `SDWAN_USERNAME` environment variable.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password for the SD-WAN Manager account. This can also be set as the `SDWAN_PASSWORD` environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: "URL of the Cisco SD-WAN Manager device. This can also be set as the `SDWAN_URL` environment variable.",
				Optional:            true,
			},
			"insecure": schema.BoolAttribute{
				MarkdownDescription: "Allow insecure HTTPS client. This can also be set as the `SDWAN_INSECURE` environment variable. Defaults to `true`.",
				Optional:            true,
			},
			"retries": schema.Int64Attribute{
				MarkdownDescription: "Number of retries for REST API calls. This can also be set as the `SDWAN_RETRIES` environment variable. Defaults to `3`.",
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 9),
				},
			},
		},
	}
}

func (p *SdwanProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Set the default log level to debug
	slog.SetLogLoggerLevel(slog.LevelDebug)

	// Retrieve provider data from configuration
	var config SdwanProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// User must provide a username to the provider
	var username string
	if config.Username.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as username",
		)
		return
	}

	if config.Username.IsNull() {
		username = os.Getenv("SDWAN_USERNAME")
	} else {
		username = config.Username.ValueString()
	}

	if username == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find username",
			"Username cannot be an empty string",
		)
		return
	}

	// User must provide a password to the provider
	var password string
	if config.Password.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as password",
		)
		return
	}

	if config.Password.IsNull() {
		password = os.Getenv("SDWAN_PASSWORD")
	} else {
		password = config.Password.ValueString()
	}

	if password == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find password",
			"Password cannot be an empty string",
		)
		return
	}

	// User must provide a username to the provider
	var url string
	if config.URL.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as url",
		)
		return
	}

	if config.URL.IsNull() {
		url = os.Getenv("SDWAN_URL")
	} else {
		url = config.URL.ValueString()
	}

	if url == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find url",
			"URL cannot be an empty string",
		)
		return
	}

	var insecure bool
	if config.Insecure.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as insecure",
		)
		return
	}

	if config.Insecure.IsNull() {
		insecureStr := os.Getenv("SDWAN_INSECURE")
		if insecureStr == "" {
			insecure = true
		} else {
			insecure, _ = strconv.ParseBool(insecureStr)
		}
	} else {
		insecure = config.Insecure.ValueBool()
	}

	var retries int64
	if config.Retries.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as retries",
		)
		return
	}

	if config.Retries.IsNull() {
		retriesStr := os.Getenv("SDWAN_RETRIES")
		if retriesStr == "" {
			retries = 3
		} else {
			retries, _ = strconv.ParseInt(retriesStr, 0, 64)
		}
	} else {
		retries = config.Retries.ValueInt64()
	}

	// Create a new SDWAN client and set it to the provider client
	c, err := sdwan.NewClient(url, username, password, insecure, sdwan.MaxRetries(int(retries)))
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Unable to create sdwan client:\n\n"+err.Error(),
		)
		return
	}

	err = c.Authenticate()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Failed to Authenticate with SDWAN Manager:"+err.Error(),
		)
		return 
	}

	data := SdwanProviderData{Client: &c, UpdateMutex: &sync.Mutex{}}
	resp.DataSourceData = &data
	resp.ResourceData = &data
}

func (p *SdwanProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		{{- range .FeatureTemplates}}
		New{{camelCase .Name}}FeatureTemplateResource,
		{{- end}}
		{{- range .ProfileParcels}}
		New{{camelCase .Name}}ProfileParcelResource,
		{{- end}}
		{{- range .Generic}}
		{{- if not .NoResource}}
		{{- if not (contains .SkipTemplates "resource.go")}}
		New{{camelCase .Name}}Resource,
		{{- end}}
		{{- end}}
		{{- end}}
		NewAttachFeatureDeviceTemplateResource,
		NewActivateCentralizedPolicyResource,
	}
}

func (p *SdwanProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		{{- range .FeatureTemplates}}
		New{{camelCase .Name}}FeatureTemplateDataSource,
		{{- end}}
		{{- range .ProfileParcels}}
		New{{camelCase .Name}}ProfileParcelDataSource,
		{{- end}}
		{{- range .Generic}}
		{{- if not .NoDataSource}}
		{{- if not (contains .SkipTemplates "data_source.go")}}
		New{{camelCase .Name}}DataSource,
		{{- end}}
		{{- end}}
		{{- end}}
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &SdwanProvider{
			version: version,
		}
	}
}
// End of section. //template:end provider
