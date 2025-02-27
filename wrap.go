package skewer

import "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2022-03-01/compute" //nolint:staticcheck

// Wrap takes an array of compute resource skus and wraps them into an
// array of our richer type.
func Wrap(in []compute.ResourceSku) []SKU {
	out := make([]SKU, len(in))
	for index, value := range in {
		out[index] = SKU(value)
	}
	return out
}
