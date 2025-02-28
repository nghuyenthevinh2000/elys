// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	ammtypes "github.com/elys-network/elys/x/amm/types"

	math "cosmossdk.io/math"

	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// AmmKeeper is an autogenerated mock type for the AmmKeeper type
type AmmKeeper struct {
	mock.Mock
}

type AmmKeeper_Expecter struct {
	mock *mock.Mock
}

func (_m *AmmKeeper) EXPECT() *AmmKeeper_Expecter {
	return &AmmKeeper_Expecter{mock: &_m.Mock}
}

// CalcInAmtGivenOut provides a mock function with given fields: ctx, poolId, oracle, snapshot, tokensOut, tokenInDenom, swapFee
func (_m *AmmKeeper) CalcInAmtGivenOut(ctx types.Context, poolId uint64, oracle ammtypes.OracleKeeper, snapshot *ammtypes.Pool, tokensOut types.Coins, tokenInDenom string, swapFee math.LegacyDec) (types.Coin, error) {
	ret := _m.Called(ctx, poolId, oracle, snapshot, tokensOut, tokenInDenom, swapFee)

	var r0 types.Coin
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, uint64, ammtypes.OracleKeeper, *ammtypes.Pool, types.Coins, string, math.LegacyDec) (types.Coin, error)); ok {
		return rf(ctx, poolId, oracle, snapshot, tokensOut, tokenInDenom, swapFee)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint64, ammtypes.OracleKeeper, *ammtypes.Pool, types.Coins, string, math.LegacyDec) types.Coin); ok {
		r0 = rf(ctx, poolId, oracle, snapshot, tokensOut, tokenInDenom, swapFee)
	} else {
		r0 = ret.Get(0).(types.Coin)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint64, ammtypes.OracleKeeper, *ammtypes.Pool, types.Coins, string, math.LegacyDec) error); ok {
		r1 = rf(ctx, poolId, oracle, snapshot, tokensOut, tokenInDenom, swapFee)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AmmKeeper_CalcInAmtGivenOut_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CalcInAmtGivenOut'
type AmmKeeper_CalcInAmtGivenOut_Call struct {
	*mock.Call
}

// CalcInAmtGivenOut is a helper method to define mock.On call
//   - ctx types.Context
//   - poolId uint64
//   - oracle ammtypes.OracleKeeper
//   - snapshot *ammtypes.Pool
//   - tokensOut types.Coins
//   - tokenInDenom string
//   - swapFee math.LegacyDec
func (_e *AmmKeeper_Expecter) CalcInAmtGivenOut(ctx interface{}, poolId interface{}, oracle interface{}, snapshot interface{}, tokensOut interface{}, tokenInDenom interface{}, swapFee interface{}) *AmmKeeper_CalcInAmtGivenOut_Call {
	return &AmmKeeper_CalcInAmtGivenOut_Call{Call: _e.mock.On("CalcInAmtGivenOut", ctx, poolId, oracle, snapshot, tokensOut, tokenInDenom, swapFee)}
}

func (_c *AmmKeeper_CalcInAmtGivenOut_Call) Run(run func(ctx types.Context, poolId uint64, oracle ammtypes.OracleKeeper, snapshot *ammtypes.Pool, tokensOut types.Coins, tokenInDenom string, swapFee math.LegacyDec)) *AmmKeeper_CalcInAmtGivenOut_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(uint64), args[2].(ammtypes.OracleKeeper), args[3].(*ammtypes.Pool), args[4].(types.Coins), args[5].(string), args[6].(math.LegacyDec))
	})
	return _c
}

func (_c *AmmKeeper_CalcInAmtGivenOut_Call) Return(tokenIn types.Coin, err error) *AmmKeeper_CalcInAmtGivenOut_Call {
	_c.Call.Return(tokenIn, err)
	return _c
}

func (_c *AmmKeeper_CalcInAmtGivenOut_Call) RunAndReturn(run func(types.Context, uint64, ammtypes.OracleKeeper, *ammtypes.Pool, types.Coins, string, math.LegacyDec) (types.Coin, error)) *AmmKeeper_CalcInAmtGivenOut_Call {
	_c.Call.Return(run)
	return _c
}

// CalcOutAmtGivenIn provides a mock function with given fields: ctx, poolId, oracle, snapshot, tokensIn, tokenOutDenom, swapFee
func (_m *AmmKeeper) CalcOutAmtGivenIn(ctx types.Context, poolId uint64, oracle ammtypes.OracleKeeper, snapshot *ammtypes.Pool, tokensIn types.Coins, tokenOutDenom string, swapFee math.LegacyDec) (types.Coin, error) {
	ret := _m.Called(ctx, poolId, oracle, snapshot, tokensIn, tokenOutDenom, swapFee)

	var r0 types.Coin
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, uint64, ammtypes.OracleKeeper, *ammtypes.Pool, types.Coins, string, math.LegacyDec) (types.Coin, error)); ok {
		return rf(ctx, poolId, oracle, snapshot, tokensIn, tokenOutDenom, swapFee)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint64, ammtypes.OracleKeeper, *ammtypes.Pool, types.Coins, string, math.LegacyDec) types.Coin); ok {
		r0 = rf(ctx, poolId, oracle, snapshot, tokensIn, tokenOutDenom, swapFee)
	} else {
		r0 = ret.Get(0).(types.Coin)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint64, ammtypes.OracleKeeper, *ammtypes.Pool, types.Coins, string, math.LegacyDec) error); ok {
		r1 = rf(ctx, poolId, oracle, snapshot, tokensIn, tokenOutDenom, swapFee)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AmmKeeper_CalcOutAmtGivenIn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CalcOutAmtGivenIn'
type AmmKeeper_CalcOutAmtGivenIn_Call struct {
	*mock.Call
}

// CalcOutAmtGivenIn is a helper method to define mock.On call
//   - ctx types.Context
//   - poolId uint64
//   - oracle ammtypes.OracleKeeper
//   - snapshot *ammtypes.Pool
//   - tokensIn types.Coins
//   - tokenOutDenom string
//   - swapFee math.LegacyDec
func (_e *AmmKeeper_Expecter) CalcOutAmtGivenIn(ctx interface{}, poolId interface{}, oracle interface{}, snapshot interface{}, tokensIn interface{}, tokenOutDenom interface{}, swapFee interface{}) *AmmKeeper_CalcOutAmtGivenIn_Call {
	return &AmmKeeper_CalcOutAmtGivenIn_Call{Call: _e.mock.On("CalcOutAmtGivenIn", ctx, poolId, oracle, snapshot, tokensIn, tokenOutDenom, swapFee)}
}

func (_c *AmmKeeper_CalcOutAmtGivenIn_Call) Run(run func(ctx types.Context, poolId uint64, oracle ammtypes.OracleKeeper, snapshot *ammtypes.Pool, tokensIn types.Coins, tokenOutDenom string, swapFee math.LegacyDec)) *AmmKeeper_CalcOutAmtGivenIn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(uint64), args[2].(ammtypes.OracleKeeper), args[3].(*ammtypes.Pool), args[4].(types.Coins), args[5].(string), args[6].(math.LegacyDec))
	})
	return _c
}

func (_c *AmmKeeper_CalcOutAmtGivenIn_Call) Return(_a0 types.Coin, _a1 error) *AmmKeeper_CalcOutAmtGivenIn_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AmmKeeper_CalcOutAmtGivenIn_Call) RunAndReturn(run func(types.Context, uint64, ammtypes.OracleKeeper, *ammtypes.Pool, types.Coins, string, math.LegacyDec) (types.Coin, error)) *AmmKeeper_CalcOutAmtGivenIn_Call {
	_c.Call.Return(run)
	return _c
}

// CalcSwapEstimationByDenom provides a mock function with given fields: ctx, amount, denomIn, denomOut, baseCurrency, discount, overrideSwapFee, decimals
func (_m *AmmKeeper) CalcSwapEstimationByDenom(ctx types.Context, amount types.Coin, denomIn string, denomOut string, baseCurrency string, discount math.LegacyDec, overrideSwapFee math.LegacyDec, decimals uint64) ([]*ammtypes.SwapAmountInRoute, []*ammtypes.SwapAmountOutRoute, types.Coin, math.LegacyDec, math.LegacyDec, math.LegacyDec, types.Coin, math.LegacyDec, math.LegacyDec, error) {
	ret := _m.Called(ctx, amount, denomIn, denomOut, baseCurrency, discount, overrideSwapFee, decimals)

	var r0 []*ammtypes.SwapAmountInRoute
	var r1 []*ammtypes.SwapAmountOutRoute
	var r2 types.Coin
	var r3 math.LegacyDec
	var r4 math.LegacyDec
	var r5 math.LegacyDec
	var r6 types.Coin
	var r7 math.LegacyDec
	var r8 math.LegacyDec
	var r9 error
	if rf, ok := ret.Get(0).(func(types.Context, types.Coin, string, string, string, math.LegacyDec, math.LegacyDec, uint64) ([]*ammtypes.SwapAmountInRoute, []*ammtypes.SwapAmountOutRoute, types.Coin, math.LegacyDec, math.LegacyDec, math.LegacyDec, types.Coin, math.LegacyDec, math.LegacyDec, error)); ok {
		return rf(ctx, amount, denomIn, denomOut, baseCurrency, discount, overrideSwapFee, decimals)
	}
	if rf, ok := ret.Get(0).(func(types.Context, types.Coin, string, string, string, math.LegacyDec, math.LegacyDec, uint64) []*ammtypes.SwapAmountInRoute); ok {
		r0 = rf(ctx, amount, denomIn, denomOut, baseCurrency, discount, overrideSwapFee, decimals)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ammtypes.SwapAmountInRoute)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, types.Coin, string, string, string, math.LegacyDec, math.LegacyDec, uint64) []*ammtypes.SwapAmountOutRoute); ok {
		r1 = rf(ctx, amount, denomIn, denomOut, baseCurrency, discount, overrideSwapFee, decimals)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*ammtypes.SwapAmountOutRoute)
		}
	}

	if rf, ok := ret.Get(2).(func(types.Context, types.Coin, string, string, string, math.LegacyDec, math.LegacyDec, uint64) types.Coin); ok {
		r2 = rf(ctx, amount, denomIn, denomOut, baseCurrency, discount, overrideSwapFee, decimals)
	} else {
		r2 = ret.Get(2).(types.Coin)
	}

	if rf, ok := ret.Get(3).(func(types.Context, types.Coin, string, string, string, math.LegacyDec, math.LegacyDec, uint64) math.LegacyDec); ok {
		r3 = rf(ctx, amount, denomIn, denomOut, baseCurrency, discount, overrideSwapFee, decimals)
	} else {
		r3 = ret.Get(3).(math.LegacyDec)
	}

	if rf, ok := ret.Get(4).(func(types.Context, types.Coin, string, string, string, math.LegacyDec, math.LegacyDec, uint64) math.LegacyDec); ok {
		r4 = rf(ctx, amount, denomIn, denomOut, baseCurrency, discount, overrideSwapFee, decimals)
	} else {
		r4 = ret.Get(4).(math.LegacyDec)
	}

	if rf, ok := ret.Get(5).(func(types.Context, types.Coin, string, string, string, math.LegacyDec, math.LegacyDec, uint64) math.LegacyDec); ok {
		r5 = rf(ctx, amount, denomIn, denomOut, baseCurrency, discount, overrideSwapFee, decimals)
	} else {
		r5 = ret.Get(5).(math.LegacyDec)
	}

	if rf, ok := ret.Get(6).(func(types.Context, types.Coin, string, string, string, math.LegacyDec, math.LegacyDec, uint64) types.Coin); ok {
		r6 = rf(ctx, amount, denomIn, denomOut, baseCurrency, discount, overrideSwapFee, decimals)
	} else {
		r6 = ret.Get(6).(types.Coin)
	}

	if rf, ok := ret.Get(7).(func(types.Context, types.Coin, string, string, string, math.LegacyDec, math.LegacyDec, uint64) math.LegacyDec); ok {
		r7 = rf(ctx, amount, denomIn, denomOut, baseCurrency, discount, overrideSwapFee, decimals)
	} else {
		r7 = ret.Get(7).(math.LegacyDec)
	}

	if rf, ok := ret.Get(8).(func(types.Context, types.Coin, string, string, string, math.LegacyDec, math.LegacyDec, uint64) math.LegacyDec); ok {
		r8 = rf(ctx, amount, denomIn, denomOut, baseCurrency, discount, overrideSwapFee, decimals)
	} else {
		r8 = ret.Get(8).(math.LegacyDec)
	}

	if rf, ok := ret.Get(9).(func(types.Context, types.Coin, string, string, string, math.LegacyDec, math.LegacyDec, uint64) error); ok {
		r9 = rf(ctx, amount, denomIn, denomOut, baseCurrency, discount, overrideSwapFee, decimals)
	} else {
		r9 = ret.Error(9)
	}

	return r0, r1, r2, r3, r4, r5, r6, r7, r8, r9
}

// AmmKeeper_CalcSwapEstimationByDenom_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CalcSwapEstimationByDenom'
type AmmKeeper_CalcSwapEstimationByDenom_Call struct {
	*mock.Call
}

// CalcSwapEstimationByDenom is a helper method to define mock.On call
//   - ctx types.Context
//   - amount types.Coin
//   - denomIn string
//   - denomOut string
//   - baseCurrency string
//   - discount math.LegacyDec
//   - overrideSwapFee math.LegacyDec
//   - decimals uint64
func (_e *AmmKeeper_Expecter) CalcSwapEstimationByDenom(ctx interface{}, amount interface{}, denomIn interface{}, denomOut interface{}, baseCurrency interface{}, discount interface{}, overrideSwapFee interface{}, decimals interface{}) *AmmKeeper_CalcSwapEstimationByDenom_Call {
	return &AmmKeeper_CalcSwapEstimationByDenom_Call{Call: _e.mock.On("CalcSwapEstimationByDenom", ctx, amount, denomIn, denomOut, baseCurrency, discount, overrideSwapFee, decimals)}
}

func (_c *AmmKeeper_CalcSwapEstimationByDenom_Call) Run(run func(ctx types.Context, amount types.Coin, denomIn string, denomOut string, baseCurrency string, discount math.LegacyDec, overrideSwapFee math.LegacyDec, decimals uint64)) *AmmKeeper_CalcSwapEstimationByDenom_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(types.Coin), args[2].(string), args[3].(string), args[4].(string), args[5].(math.LegacyDec), args[6].(math.LegacyDec), args[7].(uint64))
	})
	return _c
}

func (_c *AmmKeeper_CalcSwapEstimationByDenom_Call) Return(inRoute []*ammtypes.SwapAmountInRoute, outRoute []*ammtypes.SwapAmountOutRoute, outAmount types.Coin, spotPrice math.LegacyDec, swapFee math.LegacyDec, discountOut math.LegacyDec, availableLiquidity types.Coin, weightBonus math.LegacyDec, priceImpact math.LegacyDec, err error) *AmmKeeper_CalcSwapEstimationByDenom_Call {
	_c.Call.Return(inRoute, outRoute, outAmount, spotPrice, swapFee, discountOut, availableLiquidity, weightBonus, priceImpact, err)
	return _c
}

func (_c *AmmKeeper_CalcSwapEstimationByDenom_Call) RunAndReturn(run func(types.Context, types.Coin, string, string, string, math.LegacyDec, math.LegacyDec, uint64) ([]*ammtypes.SwapAmountInRoute, []*ammtypes.SwapAmountOutRoute, types.Coin, math.LegacyDec, math.LegacyDec, math.LegacyDec, types.Coin, math.LegacyDec, math.LegacyDec, error)) *AmmKeeper_CalcSwapEstimationByDenom_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllPool provides a mock function with given fields: _a0
func (_m *AmmKeeper) GetAllPool(_a0 types.Context) []ammtypes.Pool {
	ret := _m.Called(_a0)

	var r0 []ammtypes.Pool
	if rf, ok := ret.Get(0).(func(types.Context) []ammtypes.Pool); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ammtypes.Pool)
		}
	}

	return r0
}

// AmmKeeper_GetAllPool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllPool'
type AmmKeeper_GetAllPool_Call struct {
	*mock.Call
}

// GetAllPool is a helper method to define mock.On call
//   - _a0 types.Context
func (_e *AmmKeeper_Expecter) GetAllPool(_a0 interface{}) *AmmKeeper_GetAllPool_Call {
	return &AmmKeeper_GetAllPool_Call{Call: _e.mock.On("GetAllPool", _a0)}
}

func (_c *AmmKeeper_GetAllPool_Call) Run(run func(_a0 types.Context)) *AmmKeeper_GetAllPool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context))
	})
	return _c
}

func (_c *AmmKeeper_GetAllPool_Call) Return(_a0 []ammtypes.Pool) *AmmKeeper_GetAllPool_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AmmKeeper_GetAllPool_Call) RunAndReturn(run func(types.Context) []ammtypes.Pool) *AmmKeeper_GetAllPool_Call {
	_c.Call.Return(run)
	return _c
}

// GetPool provides a mock function with given fields: _a0, _a1
func (_m *AmmKeeper) GetPool(_a0 types.Context, _a1 uint64) (ammtypes.Pool, bool) {
	ret := _m.Called(_a0, _a1)

	var r0 ammtypes.Pool
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, uint64) (ammtypes.Pool, bool)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint64) ammtypes.Pool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(ammtypes.Pool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint64) bool); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// AmmKeeper_GetPool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPool'
type AmmKeeper_GetPool_Call struct {
	*mock.Call
}

// GetPool is a helper method to define mock.On call
//   - _a0 types.Context
//   - _a1 uint64
func (_e *AmmKeeper_Expecter) GetPool(_a0 interface{}, _a1 interface{}) *AmmKeeper_GetPool_Call {
	return &AmmKeeper_GetPool_Call{Call: _e.mock.On("GetPool", _a0, _a1)}
}

func (_c *AmmKeeper_GetPool_Call) Run(run func(_a0 types.Context, _a1 uint64)) *AmmKeeper_GetPool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(uint64))
	})
	return _c
}

func (_c *AmmKeeper_GetPool_Call) Return(_a0 ammtypes.Pool, _a1 bool) *AmmKeeper_GetPool_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AmmKeeper_GetPool_Call) RunAndReturn(run func(types.Context, uint64) (ammtypes.Pool, bool)) *AmmKeeper_GetPool_Call {
	_c.Call.Return(run)
	return _c
}

// GetBestPoolWithDenoms provides a mock function with given fields: ctx, denoms
func (_m *AmmKeeper) GetBestPoolWithDenoms(ctx types.Context, denoms []string) (ammtypes.Pool, bool) {
	ret := _m.Called(ctx, denoms)

	var r0 ammtypes.Pool
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, []string) (ammtypes.Pool, bool)); ok {
		return rf(ctx, denoms)
	}
	if rf, ok := ret.Get(0).(func(types.Context, []string) ammtypes.Pool); ok {
		r0 = rf(ctx, denoms)
	} else {
		r0 = ret.Get(0).(ammtypes.Pool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, []string) bool); ok {
		r1 = rf(ctx, denoms)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// AmmKeeper_GetBestPoolWithDenoms_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBestPoolWithDenoms'
type AmmKeeper_GetBestPoolWithDenoms_Call struct {
	*mock.Call
}

// GetBestPoolWithDenoms is a helper method to define mock.On call
//   - ctx types.Context
//   - denoms []string
func (_e *AmmKeeper_Expecter) GetBestPoolWithDenoms(ctx interface{}, denoms interface{}) *AmmKeeper_GetBestPoolWithDenoms_Call {
	return &AmmKeeper_GetBestPoolWithDenoms_Call{Call: _e.mock.On("GetBestPoolWithDenoms", ctx, denoms)}
}

func (_c *AmmKeeper_GetBestPoolWithDenoms_Call) Run(run func(ctx types.Context, denoms []string)) *AmmKeeper_GetBestPoolWithDenoms_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].([]string))
	})
	return _c
}

func (_c *AmmKeeper_GetBestPoolWithDenoms_Call) Return(poolId uint64, found bool) *AmmKeeper_GetBestPoolWithDenoms_Call {
	_c.Call.Return(poolId, found)
	return _c
}

func (_c *AmmKeeper_GetBestPoolWithDenoms_Call) RunAndReturn(run func(types.Context, []string) (uint64, bool)) *AmmKeeper_GetBestPoolWithDenoms_Call {
	_c.Call.Return(run)
	return _c
}

// GetPoolSnapshotOrSet provides a mock function with given fields: ctx, pool
func (_m *AmmKeeper) GetPoolSnapshotOrSet(ctx types.Context, pool ammtypes.Pool) ammtypes.Pool {
	ret := _m.Called(ctx, pool)

	var r0 ammtypes.Pool
	if rf, ok := ret.Get(0).(func(types.Context, ammtypes.Pool) ammtypes.Pool); ok {
		r0 = rf(ctx, pool)
	} else {
		r0 = ret.Get(0).(ammtypes.Pool)
	}

	return r0
}

// AmmKeeper_GetPoolSnapshotOrSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPoolSnapshotOrSet'
type AmmKeeper_GetPoolSnapshotOrSet_Call struct {
	*mock.Call
}

// GetPoolSnapshotOrSet is a helper method to define mock.On call
//   - ctx types.Context
//   - pool ammtypes.Pool
func (_e *AmmKeeper_Expecter) GetPoolSnapshotOrSet(ctx interface{}, pool interface{}) *AmmKeeper_GetPoolSnapshotOrSet_Call {
	return &AmmKeeper_GetPoolSnapshotOrSet_Call{Call: _e.mock.On("GetPoolSnapshotOrSet", ctx, pool)}
}

func (_c *AmmKeeper_GetPoolSnapshotOrSet_Call) Run(run func(ctx types.Context, pool ammtypes.Pool)) *AmmKeeper_GetPoolSnapshotOrSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(ammtypes.Pool))
	})
	return _c
}

func (_c *AmmKeeper_GetPoolSnapshotOrSet_Call) Return(val ammtypes.Pool) *AmmKeeper_GetPoolSnapshotOrSet_Call {
	_c.Call.Return(val)
	return _c
}

func (_c *AmmKeeper_GetPoolSnapshotOrSet_Call) RunAndReturn(run func(types.Context, ammtypes.Pool) ammtypes.Pool) *AmmKeeper_GetPoolSnapshotOrSet_Call {
	_c.Call.Return(run)
	return _c
}

// IterateLiquidityPools provides a mock function with given fields: _a0, _a1
func (_m *AmmKeeper) IterateLiquidityPools(_a0 types.Context, _a1 func(ammtypes.Pool) bool) {
	_m.Called(_a0, _a1)
}

// AmmKeeper_IterateLiquidityPools_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IterateLiquidityPools'
type AmmKeeper_IterateLiquidityPools_Call struct {
	*mock.Call
}

// IterateLiquidityPools is a helper method to define mock.On call
//   - _a0 types.Context
//   - _a1 func(ammtypes.Pool) bool
func (_e *AmmKeeper_Expecter) IterateLiquidityPools(_a0 interface{}, _a1 interface{}) *AmmKeeper_IterateLiquidityPools_Call {
	return &AmmKeeper_IterateLiquidityPools_Call{Call: _e.mock.On("IterateLiquidityPools", _a0, _a1)}
}

func (_c *AmmKeeper_IterateLiquidityPools_Call) Run(run func(_a0 types.Context, _a1 func(ammtypes.Pool) bool)) *AmmKeeper_IterateLiquidityPools_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(func(ammtypes.Pool) bool))
	})
	return _c
}

func (_c *AmmKeeper_IterateLiquidityPools_Call) Return() *AmmKeeper_IterateLiquidityPools_Call {
	_c.Call.Return()
	return _c
}

func (_c *AmmKeeper_IterateLiquidityPools_Call) RunAndReturn(run func(types.Context, func(ammtypes.Pool) bool)) *AmmKeeper_IterateLiquidityPools_Call {
	_c.Call.Return(run)
	return _c
}

// NewAmmKeeper creates a new instance of AmmKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAmmKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *AmmKeeper {
	mock := &AmmKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
