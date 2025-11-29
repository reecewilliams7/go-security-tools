package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// RequiresReplaceInt64 returns a plan modifier that requires replacement when the value changes
func RequiresReplaceInt64() planmodifier.Int64 {
	return int64planmodifier.RequiresReplace()
}
