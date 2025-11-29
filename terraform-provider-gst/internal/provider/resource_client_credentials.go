package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/reecewilliams7/go-security-tools/clientcredentials"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ClientCredentialsResource{}
var _ resource.ResourceWithImportState = &ClientCredentialsResource{}

func NewClientCredentialsResource() resource.Resource {
	return &ClientCredentialsResource{}
}

// ClientCredentialsResource defines the resource implementation.
type ClientCredentialsResource struct{}

// ClientCredentialsResourceModel describes the resource data model.
type ClientCredentialsResourceModel struct {
	ID           types.String `tfsdk:"id"`
	ClientID     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
}

func (r *ClientCredentialsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_client_credentials"
}

func (r *ClientCredentialsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Generates OAuth2 client credentials (client_id and client_secret) for authentication.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The unique identifier for the client credentials (same as client_id)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"client_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The generated client ID (UUID format)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"client_secret": schema.StringAttribute{
				Computed:            true,
				Sensitive:           true,
				MarkdownDescription: "The generated client secret (base64-encoded random string)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *ClientCredentialsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Provider configuration would be accessed here if needed
}

func (r *ClientCredentialsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClientCredentialsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create the client credentials using GUID for client ID and crypto random for secret
	clientIDCreator := clientcredentials.NewUUIDv7ClientIDCreator()
	clientSecretCreator := clientcredentials.NewCryptoRandClientSecretCreator()

	creator := clientcredentials.NewClientCredentialsCreator(clientIDCreator, clientSecretCreator)

	credentials, err := creator.CreateClientCredentials()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create client credentials, got error: %s", err))
		return
	}

	// Map the output to the resource model
	data.ID = types.StringValue(credentials.ClientID())
	data.ClientID = types.StringValue(credentials.ClientID())
	data.ClientSecret = types.StringValue(credentials.ClientSecret())

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClientCredentialsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClientCredentialsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Client credentials are immutable once created
	// The state already contains all the information we need

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClientCredentialsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ClientCredentialsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Client credentials should not be updated, they should be replaced
	resp.Diagnostics.AddError(
		"Update Not Supported",
		"Client credentials resources cannot be updated. They must be replaced.",
	)
}

func (r *ClientCredentialsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClientCredentialsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Client credentials are local resources with no remote state to clean up
	// The deletion is handled by removing it from the Terraform state
}

func (r *ClientCredentialsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
