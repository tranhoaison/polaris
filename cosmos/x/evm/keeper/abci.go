// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlock runs on the Cosmos-SDK lifecycle EndBlock() during ABCI Finalize.
func (k *Keeper) EndBlock(ctx context.Context) error {
	// Verify that the EVM block was written.
	// TODO: Set/GetHead to set and get the canonical head.
	blockNum := uint64(sdk.UnwrapSDKContext(ctx).BlockHeight())
	block := k.chain.GetBlockByNumber(blockNum)
	if block == nil {
		return fmt.Errorf(
			"evm block %d failed to process", blockNum,
		)
	} else if block.NumberU64() != blockNum {
		return fmt.Errorf(
			"evm block [%d] does not match comet block [%d]", block.NumberU64(), blockNum,
		)
	}
	return nil
}

// SetLatestQueryContext runs on the Cosmos-SDK lifecycle PrepareCheckState() during ABCI Commit.
func (k *Keeper) SetLatestQueryContext(ctx context.Context) error {
	k.spf.SetLatestQueryContext(ctx)
	return nil
}
