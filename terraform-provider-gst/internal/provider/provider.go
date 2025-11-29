package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure GSTProvider satisfies various provider interfaces.
var _ provider.Provider = &GSTProvider{}

// GSTProvider defines the provider implementation.
type GSTProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// GSTProviderModel describes the provider data model.
type GSTProviderModel struct {
	// Provider configuration would go here if needed
}

func (p *GSTProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "gst"
	resp.Version = p.version
}

func (p *GSTProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The GST (Go Security Tools) provider allows you to generate security credentials such as JSON Web Keys (JWK) and OAuth2 client credentials.",
		Attributes:  map[string]schema.Attribute{},
	}
}

func (p *GSTProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data GSTProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (p *GSTProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewJWKResource,
		NewClientCredentialsResource,
	}
}

func (p *GSTProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &GSTProvider{
			version: version,
		}
	}
}
