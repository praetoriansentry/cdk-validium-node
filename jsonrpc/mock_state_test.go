// Code generated by mockery v2.13.1. DO NOT EDIT.

package jsonrpc

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	runtime "github.com/hermeznetwork/hermez-core/state/runtime"

	state "github.com/hermeznetwork/hermez-core/state"

	time "time"

	types "github.com/ethereum/go-ethereum/core/types"
)

// stateMock is an autogenerated mock type for the stateInterface type
type stateMock struct {
	mock.Mock
}

// DebugTransaction provides a mock function with given fields: ctx, transactionHash, tracer
func (_m *stateMock) DebugTransaction(ctx context.Context, transactionHash common.Hash, tracer string) (*runtime.ExecutionResult, error) {
	ret := _m.Called(ctx, transactionHash, tracer)

	var r0 *runtime.ExecutionResult
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, string) *runtime.ExecutionResult); ok {
		r0 = rf(ctx, transactionHash, tracer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*runtime.ExecutionResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, string) error); ok {
		r1 = rf(ctx, transactionHash, tracer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EstimateGas provides a mock function with given fields: transaction, senderAddress
func (_m *stateMock) EstimateGas(transaction *types.Transaction, senderAddress common.Address) (uint64, error) {
	ret := _m.Called(transaction, senderAddress)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*types.Transaction, common.Address) uint64); ok {
		r0 = rf(transaction, senderAddress)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.Transaction, common.Address) error); ok {
		r1 = rf(transaction, senderAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBalance provides a mock function with given fields: ctx, address, batchNumber, txBundleID
func (_m *stateMock) GetBalance(ctx context.Context, address common.Address, batchNumber uint64, txBundleID string) (*big.Int, error) {
	ret := _m.Called(ctx, address, batchNumber, txBundleID)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, uint64, string) *big.Int); ok {
		r0 = rf(ctx, address, batchNumber, txBundleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, uint64, string) error); ok {
		r1 = rf(ctx, address, batchNumber, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBatchByHash provides a mock function with given fields: ctx, hash, txBundleID
func (_m *stateMock) GetBatchByHash(ctx context.Context, hash common.Hash, txBundleID string) (*state.Batch, error) {
	ret := _m.Called(ctx, hash, txBundleID)

	var r0 *state.Batch
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, string) *state.Batch); ok {
		r0 = rf(ctx, hash, txBundleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, string) error); ok {
		r1 = rf(ctx, hash, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBatchByNumber provides a mock function with given fields: ctx, batchNumber, txBundleID
func (_m *stateMock) GetBatchByNumber(ctx context.Context, batchNumber uint64, txBundleID string) (*state.Batch, error) {
	ret := _m.Called(ctx, batchNumber, txBundleID)

	var r0 *state.Batch
	if rf, ok := ret.Get(0).(func(context.Context, uint64, string) *state.Batch); ok {
		r0 = rf(ctx, batchNumber, txBundleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, string) error); ok {
		r1 = rf(ctx, batchNumber, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBatchHashesSince provides a mock function with given fields: ctx, since, txBundleID
func (_m *stateMock) GetBatchHashesSince(ctx context.Context, since time.Time, txBundleID string) ([]common.Hash, error) {
	ret := _m.Called(ctx, since, txBundleID)

	var r0 []common.Hash
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, string) []common.Hash); ok {
		r0 = rf(ctx, since, txBundleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, time.Time, string) error); ok {
		r1 = rf(ctx, since, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBatchHeader provides a mock function with given fields: ctx, batchNumber, txBundleID
func (_m *stateMock) GetBatchHeader(ctx context.Context, batchNumber uint64, txBundleID string) (*types.Header, error) {
	ret := _m.Called(ctx, batchNumber, txBundleID)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(context.Context, uint64, string) *types.Header); ok {
		r0 = rf(ctx, batchNumber, txBundleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, string) error); ok {
		r1 = rf(ctx, batchNumber, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBatchTransactionCountByHash provides a mock function with given fields: ctx, hash, txBundleID
func (_m *stateMock) GetBatchTransactionCountByHash(ctx context.Context, hash common.Hash, txBundleID string) (uint64, error) {
	ret := _m.Called(ctx, hash, txBundleID)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, string) uint64); ok {
		r0 = rf(ctx, hash, txBundleID)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, string) error); ok {
		r1 = rf(ctx, hash, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBatchTransactionCountByNumber provides a mock function with given fields: ctx, batchNumber, txBundleID
func (_m *stateMock) GetBatchTransactionCountByNumber(ctx context.Context, batchNumber uint64, txBundleID string) (uint64, error) {
	ret := _m.Called(ctx, batchNumber, txBundleID)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, uint64, string) uint64); ok {
		r0 = rf(ctx, batchNumber, txBundleID)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, string) error); ok {
		r1 = rf(ctx, batchNumber, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCode provides a mock function with given fields: ctx, address, batchNumber, txBundleID
func (_m *stateMock) GetCode(ctx context.Context, address common.Address, batchNumber uint64, txBundleID string) ([]byte, error) {
	ret := _m.Called(ctx, address, batchNumber, txBundleID)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, uint64, string) []byte); ok {
		r0 = rf(ctx, address, batchNumber, txBundleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, uint64, string) error); ok {
		r1 = rf(ctx, address, batchNumber, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBatch provides a mock function with given fields: ctx, isVirtual, txBundleID
func (_m *stateMock) GetLastBatch(ctx context.Context, isVirtual bool, txBundleID string) (*state.Batch, error) {
	ret := _m.Called(ctx, isVirtual, txBundleID)

	var r0 *state.Batch
	if rf, ok := ret.Get(0).(func(context.Context, bool, string) *state.Batch); ok {
		r0 = rf(ctx, isVirtual, txBundleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, bool, string) error); ok {
		r1 = rf(ctx, isVirtual, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBatchNumber provides a mock function with given fields: ctx, txBundleID
func (_m *stateMock) GetLastBatchNumber(ctx context.Context, txBundleID string) (uint64, error) {
	ret := _m.Called(ctx, txBundleID)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, string) uint64); ok {
		r0 = rf(ctx, txBundleID)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastConsolidatedBatchNumber provides a mock function with given fields: ctx, txBundleID
func (_m *stateMock) GetLastConsolidatedBatchNumber(ctx context.Context, txBundleID string) (uint64, error) {
	ret := _m.Called(ctx, txBundleID)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, string) uint64); ok {
		r0 = rf(ctx, txBundleID)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLogs provides a mock function with given fields: ctx, fromBatch, toBatch, addresses, topics, batchHash, since, txBundleID
func (_m *stateMock) GetLogs(ctx context.Context, fromBatch uint64, toBatch uint64, addresses []common.Address, topics [][]common.Hash, batchHash *common.Hash, since *time.Time, txBundleID string) ([]*types.Log, error) {
	ret := _m.Called(ctx, fromBatch, toBatch, addresses, topics, batchHash, since, txBundleID)

	var r0 []*types.Log
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, []common.Address, [][]common.Hash, *common.Hash, *time.Time, string) []*types.Log); ok {
		r0 = rf(ctx, fromBatch, toBatch, addresses, topics, batchHash, since, txBundleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Log)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, []common.Address, [][]common.Hash, *common.Hash, *time.Time, string) error); ok {
		r1 = rf(ctx, fromBatch, toBatch, addresses, topics, batchHash, since, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNonce provides a mock function with given fields: ctx, address, batchNumber, txBundleID
func (_m *stateMock) GetNonce(ctx context.Context, address common.Address, batchNumber uint64, txBundleID string) (uint64, error) {
	ret := _m.Called(ctx, address, batchNumber, txBundleID)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, uint64, string) uint64); ok {
		r0 = rf(ctx, address, batchNumber, txBundleID)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, uint64, string) error); ok {
		r1 = rf(ctx, address, batchNumber, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageAt provides a mock function with given fields: ctx, address, position, batchNumber, txBundleID
func (_m *stateMock) GetStorageAt(ctx context.Context, address common.Address, position *big.Int, batchNumber uint64, txBundleID string) (*big.Int, error) {
	ret := _m.Called(ctx, address, position, batchNumber, txBundleID)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int, uint64, string) *big.Int); ok {
		r0 = rf(ctx, address, position, batchNumber, txBundleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int, uint64, string) error); ok {
		r1 = rf(ctx, address, position, batchNumber, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSyncingInfo provides a mock function with given fields: ctx, txBundleID
func (_m *stateMock) GetSyncingInfo(ctx context.Context, txBundleID string) (state.SyncingInfo, error) {
	ret := _m.Called(ctx, txBundleID)

	var r0 state.SyncingInfo
	if rf, ok := ret.Get(0).(func(context.Context, string) state.SyncingInfo); ok {
		r0 = rf(ctx, txBundleID)
	} else {
		r0 = ret.Get(0).(state.SyncingInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByBatchHashAndIndex provides a mock function with given fields: ctx, batchHash, index, txBundleID
func (_m *stateMock) GetTransactionByBatchHashAndIndex(ctx context.Context, batchHash common.Hash, index uint64, txBundleID string) (*types.Transaction, error) {
	ret := _m.Called(ctx, batchHash, index, txBundleID)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, uint64, string) *types.Transaction); ok {
		r0 = rf(ctx, batchHash, index, txBundleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, uint64, string) error); ok {
		r1 = rf(ctx, batchHash, index, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByBatchNumberAndIndex provides a mock function with given fields: ctx, batchNumber, index, txBundleID
func (_m *stateMock) GetTransactionByBatchNumberAndIndex(ctx context.Context, batchNumber uint64, index uint64, txBundleID string) (*types.Transaction, error) {
	ret := _m.Called(ctx, batchNumber, index, txBundleID)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, string) *types.Transaction); ok {
		r0 = rf(ctx, batchNumber, index, txBundleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, string) error); ok {
		r1 = rf(ctx, batchNumber, index, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByHash provides a mock function with given fields: ctx, transactionHash, txBundleID
func (_m *stateMock) GetTransactionByHash(ctx context.Context, transactionHash common.Hash, txBundleID string) (*types.Transaction, error) {
	ret := _m.Called(ctx, transactionHash, txBundleID)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, string) *types.Transaction); ok {
		r0 = rf(ctx, transactionHash, txBundleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, string) error); ok {
		r1 = rf(ctx, transactionHash, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionReceipt provides a mock function with given fields: ctx, transactionHash, txBundleID
func (_m *stateMock) GetTransactionReceipt(ctx context.Context, transactionHash common.Hash, txBundleID string) (*state.Receipt, error) {
	ret := _m.Called(ctx, transactionHash, txBundleID)

	var r0 *state.Receipt
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, string) *state.Receipt); ok {
		r0 = rf(ctx, transactionHash, txBundleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, string) error); ok {
		r1 = rf(ctx, transactionHash, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBatchProcessor provides a mock function with given fields: ctx, sequencerAddress, stateRoot, txBundleID
func (_m *stateMock) NewBatchProcessor(ctx context.Context, sequencerAddress common.Address, stateRoot []byte, txBundleID string) (BatchProcessorInterface, error) {
	ret := _m.Called(ctx, sequencerAddress, stateRoot, txBundleID)

	var r0 BatchProcessorInterface
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, []byte, string) BatchProcessorInterface); ok {
		r0 = rf(ctx, sequencerAddress, stateRoot, txBundleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(BatchProcessorInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, []byte, string) error); ok {
		r1 = rf(ctx, sequencerAddress, stateRoot, txBundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewStateMock interface {
	mock.TestingT
	Cleanup(func())
}

// newStateMock creates a new instance of stateMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newStateMock(t mockConstructorTestingTnewStateMock) *stateMock {
	mock := &stateMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}