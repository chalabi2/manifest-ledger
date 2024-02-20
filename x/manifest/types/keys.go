package types

import "cosmossdk.io/collections"

// ParamsKey saves the current module params.
var ParamsKey = collections.NewPrefix(0)

const (
	ModuleName = "manifest"

	StoreKey = ModuleName

	QuerierRoute = ModuleName
)
