package skewer

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2022-03-01/compute" //nolint:staticcheck
)

// ResourceClient is the required Azure client interface used to populate skewer's data.
type ResourceClient interface {
	ListComplete(ctx context.Context, filter, includeExtendedLocations string) (compute.ResourceSkusResultIterator, error)
}

// ResourceProviderClient is a convenience interface for uses cases
// specific to Azure resource providers.
type ResourceProviderClient interface {
	List(ctx context.Context, filter, includeExtendedLocations string) (compute.ResourceSkusResultPage, error)
}

// client defines the internal interface required by the skewer Cache.
// TODO(ace): implement a lazy iterator with caching (and a cursor?)
type client interface {
	List(ctx context.Context, filter, includeExtendedLocations string) ([]compute.ResourceSku, error)
}
