package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liftedinit/manifest-ledger/x/distribution/types"
)

func TestValidateGenesis(t *testing.T) {
	fp := types.InitialFeePool()
	require.Nil(t, fp.ValidateGenesis())

	fp2 := types.FeePool{CommunityPool: sdk.DecCoins{{Denom: "stake", Amount: math.LegacyNewDec(-1)}}}
	require.NotNil(t, fp2.ValidateGenesis())
}
