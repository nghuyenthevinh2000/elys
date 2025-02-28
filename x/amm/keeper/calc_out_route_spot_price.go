package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/amm/types"
)

// CalcOutRouteSpotPrice calculates the spot price of the given token and out route
func (k Keeper) CalcOutRouteSpotPrice(ctx sdk.Context, tokenOut sdk.Coin, routes []*types.SwapAmountOutRoute, discount sdk.Dec, overrideSwapFee sdk.Dec) (sdk.Dec, sdk.Coin, sdk.Dec, sdk.Dec, sdk.Coin, sdk.Dec, error) {
	if routes == nil || len(routes) == 0 {
		return sdk.ZeroDec(), sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroDec(), sdk.Coin{}, sdk.ZeroDec(), types.ErrEmptyRoutes
	}

	// Start with the initial token input
	tokensOut := sdk.Coins{tokenOut}

	// The final token in denom
	var tokenInDenom string

	// Initialize the total discounted swap fee
	totalDiscountedSwapFee := sdk.ZeroDec()

	// Track the total available liquidity in the pool for final token out denom
	var availableLiquidity sdk.Coin

	weightBonus := sdk.ZeroDec()

	for _, route := range routes {
		poolId := route.PoolId
		tokenInDenom = route.TokenInDenom

		pool, found := k.GetPool(ctx, poolId)
		if !found {
			return sdk.ZeroDec(), sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroDec(), sdk.Coin{}, sdk.ZeroDec(), types.ErrPoolNotFound
		}

		// Get Pool swap fee
		swapFee := pool.GetPoolParams().SwapFee

		// Override swap fee if applicable
		if overrideSwapFee.IsPositive() {
			swapFee = overrideSwapFee
		}

		// Apply discount
		swapFee = types.ApplyDiscount(swapFee, discount)

		// Calculate the total discounted swap fee
		totalDiscountedSwapFee = totalDiscountedSwapFee.Add(swapFee)

		// Estimate swap
		snapshot := k.GetPoolSnapshotOrSet(ctx, pool)
		cacheCtx, _ := ctx.CacheContext()
		swapResult, _, weightBalanceBonus, err := k.SwapInAmtGivenOut(cacheCtx, pool.PoolId, k.oracleKeeper, &snapshot, tokensOut, tokenInDenom, swapFee)
		if err != nil {
			return sdk.ZeroDec(), sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroDec(), sdk.Coin{}, sdk.ZeroDec(), err
		}

		if swapResult.IsZero() {
			return sdk.ZeroDec(), sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroDec(), sdk.Coin{}, sdk.ZeroDec(), types.ErrAmountTooLow
		}

		// Use the current swap result as the input for the next iteration
		tokensOut = sdk.Coins{swapResult}

		// Get the available liquidity for the final token in denom
		_, poolAsset, err := pool.GetPoolAssetAndIndex(tokenInDenom)
		if err != nil {
			return sdk.ZeroDec(), sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroDec(), sdk.Coin{}, sdk.ZeroDec(), err
		}
		availableLiquidity = poolAsset.Token
		weightBonus = weightBalanceBonus
	}

	// Ensure tokenIn.Amount is not zero to avoid division by zero
	if tokenOut.IsZero() {
		return sdk.ZeroDec(), sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroDec(), sdk.Coin{}, sdk.ZeroDec(), types.ErrAmountTooLow
	}

	// Calculate the spot price given the initial token in and the final token in
	spotPrice := sdk.NewDecFromInt(tokensOut[0].Amount).Quo(sdk.NewDecFromInt(tokenOut.Amount))

	// Calculate the token in amount
	tokenInAmt := sdk.NewDecFromInt(tokenOut.Amount).Mul(spotPrice)

	// invert the spot price
	spotPrice = sdk.OneDec().Quo(spotPrice)

	// Construct the token out coin
	tokenIn := sdk.NewCoin(tokenInDenom, tokenInAmt.TruncateInt())

	return spotPrice, tokenIn, totalDiscountedSwapFee, discount, availableLiquidity, weightBonus, nil
}
