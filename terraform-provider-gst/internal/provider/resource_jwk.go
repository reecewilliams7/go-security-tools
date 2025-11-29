package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/reecewilliams7/go-security-tools/jwk"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &JWKResource{}
var _ resource.ResourceWithImportState = &JWKResource{}

func NewJWKResource() resource.Resource {
	return &JWKResource{}
}

// JWKResource defines the resource implementation.
type JWKResource struct{}

// JWKResourceModel describes the resource data model.
type JWKResourceModel struct {
	ID              types.String `tfsdk:"id"`
	Algorithm       types.String `tfsdk:"algorithm"`
	KeySize         types.Int64  `tfsdk:"key_size"`
	CurveType       types.String `tfsdk:"curve_type"`
	JWKString       types.String `tfsdk:"jwk_string"`
	JWKPublicString types.String `tfsdk:"jwk_public_string"`
	Base64JWK       types.String `tfsdk:"base64_jwk"`
	PEMPublicKey    types.String `tfsdk:"pem_public_key"`
	PEMPrivateKey   types.String `tfsdk:"pem_private_key"`
}

func (r *JWKResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_jwk"
}

func (r *JWKResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Generates a JSON Web Key (JWK) for cryptographic operations. Supports RSA and ECDSA algorithms.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The unique identifier for the JWK (kid - Key ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"algorithm": schema.StringAttribute{
				MarkdownDescription: "The algorithm type for the JWK. Supported values: 'rsa', 'ecdsa'. Default: 'rsa'",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("rsa"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"key_size": schema.Int64Attribute{
				MarkdownDescription: "The key size in bits for RSA keys. Default: 2048. Common values: 2048, 3072, 4096",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(2048),
				PlanModifiers: []planmodifier.Int64{
					RequiresReplaceInt64(),
				},
			},
			"curve_type": schema.StringAttribute{
				MarkdownDescription: "The elliptic curve type for ECDSA keys. Supported values: 'P256', 'P384', 'P521'. Default: 'P256'",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("P256"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"jwk_string": schema.StringAttribute{
				Computed:            true,
				Sensitive:           true,
				MarkdownDescription: "The complete JWK as a JSON string (including private key)",
			},
			"jwk_public_string": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The public portion of the JWK as a JSON string",
			},
			"base64_jwk": schema.StringAttribute{
				Computed:            true,
				Sensitive:           true,
				MarkdownDescription: "The Base64-encoded JWK",
			},
			"pem_public_key": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The public key in PEM format",
			},
			"pem_private_key": schema.StringAttribute{
				Computed:            true,
				Sensitive:           true,
				MarkdownDescription: "The private key in PEM format",
			},
		},
	}
}

func (r *JWKResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Provider configuration would be accessed here if needed
}

func (r *JWKResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data JWKResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Generate the JWK based on the algorithm
	var creator jwk.JWKCreator
	algorithm := data.Algorithm.ValueString()

	switch algorithm {
	case "rsa":
		keySize := int(data.KeySize.ValueInt64())
		creator = jwk.NewRSAJSONWebKeyCreator(keySize)
	case "ecdsa":
		curveType := data.CurveType.ValueString()
		creator = jwk.NewECDSAJWKCreator(curveType)
	default:
		resp.Diagnostics.AddError(
			"Invalid Algorithm",
			fmt.Sprintf("Unsupported algorithm: %s. Supported values are 'rsa' and 'ecdsa'", algorithm),
		)
		return
	}

	// Create the JWK
	output, err := creator.Create()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create JWK, got error: %s", err))
		return
	}

	// Extract the Key ID (kid) from the JWK
	keyID, ok := output.JWK.Get("kid")
	if !ok {
		resp.Diagnostics.AddError("Client Error", "Unable to extract Key ID from JWK")
		return
	}

	// Map the output to the resource model
	data.ID = types.StringValue(keyID.(string))
	data.JWKString = types.StringValue(output.JWKString)
	data.JWKPublicString = types.StringValue(output.JWKPublicString)
	data.Base64JWK = types.StringValue(output.Base64JWK)
	data.PEMPublicKey = types.StringValue(output.PEMPublicKey)
	data.PEMPrivateKey = types.StringValue(output.PEMPrivateKey)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *JWKResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data JWKResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// JWKs are immutable once created, so we don't need to refresh from an API
	// The state already contains all the information we need

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *JWKResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data JWKResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// JWKs should not be updated, they should be replaced
	// This is enforced by the RequiresReplace plan modifiers
	resp.Diagnostics.AddError(
		"Update Not Supported",
		"JWK resources cannot be updated. They must be replaced.",
	)
}

func (r *JWKResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data JWKResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// JWKs are local resources with no remote state to clean up
	// The deletion is handled by removing it from the Terraform state
}

func (r *JWKResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
