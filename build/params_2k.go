// +build debug 2k

package build

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	v0miner "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	v0power "github.com/filecoin-project/specs-actors/actors/builtin/power"
	v0verifreg "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"
)

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 0

const UpgradeSmokeHeight = -1

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

func init() {
	v0power.ConsensusMinerMinPower = big.NewInt(2048)
	v0miner.SupportedProofTypes = map[abi.RegisteredSealProof]struct{}{
		abi.RegisteredSealProof_StackedDrg2KiBV1: {},
	}
	v0verifreg.MinVerifiedDealSize = big.NewInt(256)

	BuildType |= Build2k
}

const BlockDelaySecs = uint64(30)

const PropagationDelaySecs = uint64(1)

// SlashablePowerDelay is the number of epochs after ElectionPeriodStart, after
// which the miner is slashed
//
// Epochs
const SlashablePowerDelay = 20

// Epochs
const InteractivePoRepConfidence = 6
