package bucket

import (
	"fmt"

	"github.com/ettle/strcase"
)

// toResourceName allows the creation of a snake_case based names for resources
// being created in this Stack, using the Name of the bucket as an
// enforced-prefix within the Stack. Resource names must also not only be unique
// per resource, but unique in the Stack as well, otherwise CDKTF will throw the
// following error:
//
//	panic: There is already a Construct with name 'assets' in App
func toResourceName(o *BucketOptions, r string) *string {
	name := fmt.Sprintf(
		"%s_%s",
		strcase.ToSnake(o.Name),
		strcase.ToSnake(r),
	)

	return &name
}

// name allows the creation of a snake_case name based on any string
// provided, helping to standardize the names of Stacks and Resources in this
// Stack.
func toName(r string) *string {
	name := strcase.ToSnake(r)
	return &name
}
