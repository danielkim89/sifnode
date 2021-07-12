package v42

import (
	v039ethbridge "github.com/Sifchain/sifnode/x/ethbridge/legacy/v39"
	"github.com/Sifchain/sifnode/x/ethbridge/types"
)

func Migrate(state v039ethbridge.GenesisState) *types.GenesisState {
	return &types.GenesisState{
		NativeTokenReceiveAccount: state.NativeTokenReceiverAccount.String(),
		PeggyTokens:               state.PeggyTokens,
	}
}
