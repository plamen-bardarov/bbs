// Code generated by counterfeiter. DO NOT EDIT.
package helpersfakes

import (
	"database/sql"
	"sync"
	"time"

	"code.cloudfoundry.org/bbs/db/sqldb/helpers"
)

type FakeQueryableDB struct {
	BeginStub        func() (helpers.Tx, error)
	beginMutex       sync.RWMutex
	beginArgsForCall []struct {
	}
	beginReturns struct {
		result1 helpers.Tx
		result2 error
	}
	beginReturnsOnCall map[int]struct {
		result1 helpers.Tx
		result2 error
	}
	ExecStub        func(string, ...interface{}) (sql.Result, error)
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	execReturns struct {
		result1 sql.Result
		result2 error
	}
	execReturnsOnCall map[int]struct {
		result1 sql.Result
		result2 error
	}
	OpenConnectionsStub        func() int
	openConnectionsMutex       sync.RWMutex
	openConnectionsArgsForCall []struct {
	}
	openConnectionsReturns struct {
		result1 int
	}
	openConnectionsReturnsOnCall map[int]struct {
		result1 int
	}
	PrepareStub        func(string) (*sql.Stmt, error)
	prepareMutex       sync.RWMutex
	prepareArgsForCall []struct {
		arg1 string
	}
	prepareReturns struct {
		result1 *sql.Stmt
		result2 error
	}
	prepareReturnsOnCall map[int]struct {
		result1 *sql.Stmt
		result2 error
	}
	QueryStub        func(string, ...interface{}) (*sql.Rows, error)
	queryMutex       sync.RWMutex
	queryArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	queryReturns struct {
		result1 *sql.Rows
		result2 error
	}
	queryReturnsOnCall map[int]struct {
		result1 *sql.Rows
		result2 error
	}
	QueryRowStub        func(string, ...interface{}) helpers.RowScanner
	queryRowMutex       sync.RWMutex
	queryRowArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	queryRowReturns struct {
		result1 helpers.RowScanner
	}
	queryRowReturnsOnCall map[int]struct {
		result1 helpers.RowScanner
	}
	WaitCountStub        func() int64
	waitCountMutex       sync.RWMutex
	waitCountArgsForCall []struct {
	}
	waitCountReturns struct {
		result1 int64
	}
	waitCountReturnsOnCall map[int]struct {
		result1 int64
	}
	WaitDurationStub        func() time.Duration
	waitDurationMutex       sync.RWMutex
	waitDurationArgsForCall []struct {
	}
	waitDurationReturns struct {
		result1 time.Duration
	}
	waitDurationReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeQueryableDB) Begin() (helpers.Tx, error) {
	fake.beginMutex.Lock()
	ret, specificReturn := fake.beginReturnsOnCall[len(fake.beginArgsForCall)]
	fake.beginArgsForCall = append(fake.beginArgsForCall, struct {
	}{})
	fake.recordInvocation("Begin", []interface{}{})
	fake.beginMutex.Unlock()
	if fake.BeginStub != nil {
		return fake.BeginStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.beginReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeQueryableDB) BeginCallCount() int {
	fake.beginMutex.RLock()
	defer fake.beginMutex.RUnlock()
	return len(fake.beginArgsForCall)
}

func (fake *FakeQueryableDB) BeginCalls(stub func() (helpers.Tx, error)) {
	fake.beginMutex.Lock()
	defer fake.beginMutex.Unlock()
	fake.BeginStub = stub
}

func (fake *FakeQueryableDB) BeginReturns(result1 helpers.Tx, result2 error) {
	fake.beginMutex.Lock()
	defer fake.beginMutex.Unlock()
	fake.BeginStub = nil
	fake.beginReturns = struct {
		result1 helpers.Tx
		result2 error
	}{result1, result2}
}

func (fake *FakeQueryableDB) BeginReturnsOnCall(i int, result1 helpers.Tx, result2 error) {
	fake.beginMutex.Lock()
	defer fake.beginMutex.Unlock()
	fake.BeginStub = nil
	if fake.beginReturnsOnCall == nil {
		fake.beginReturnsOnCall = make(map[int]struct {
			result1 helpers.Tx
			result2 error
		})
	}
	fake.beginReturnsOnCall[i] = struct {
		result1 helpers.Tx
		result2 error
	}{result1, result2}
}

func (fake *FakeQueryableDB) Exec(arg1 string, arg2 ...interface{}) (sql.Result, error) {
	fake.execMutex.Lock()
	ret, specificReturn := fake.execReturnsOnCall[len(fake.execArgsForCall)]
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Exec", []interface{}{arg1, arg2})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.execReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeQueryableDB) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FakeQueryableDB) ExecCalls(stub func(string, ...interface{}) (sql.Result, error)) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = stub
}

func (fake *FakeQueryableDB) ExecArgsForCall(i int) (string, []interface{}) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	argsForCall := fake.execArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeQueryableDB) ExecReturns(result1 sql.Result, result2 error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeQueryableDB) ExecReturnsOnCall(i int, result1 sql.Result, result2 error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	if fake.execReturnsOnCall == nil {
		fake.execReturnsOnCall = make(map[int]struct {
			result1 sql.Result
			result2 error
		})
	}
	fake.execReturnsOnCall[i] = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeQueryableDB) OpenConnections() int {
	fake.openConnectionsMutex.Lock()
	ret, specificReturn := fake.openConnectionsReturnsOnCall[len(fake.openConnectionsArgsForCall)]
	fake.openConnectionsArgsForCall = append(fake.openConnectionsArgsForCall, struct {
	}{})
	fake.recordInvocation("OpenConnections", []interface{}{})
	fake.openConnectionsMutex.Unlock()
	if fake.OpenConnectionsStub != nil {
		return fake.OpenConnectionsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.openConnectionsReturns
	return fakeReturns.result1
}

func (fake *FakeQueryableDB) OpenConnectionsCallCount() int {
	fake.openConnectionsMutex.RLock()
	defer fake.openConnectionsMutex.RUnlock()
	return len(fake.openConnectionsArgsForCall)
}

func (fake *FakeQueryableDB) OpenConnectionsCalls(stub func() int) {
	fake.openConnectionsMutex.Lock()
	defer fake.openConnectionsMutex.Unlock()
	fake.OpenConnectionsStub = stub
}

func (fake *FakeQueryableDB) OpenConnectionsReturns(result1 int) {
	fake.openConnectionsMutex.Lock()
	defer fake.openConnectionsMutex.Unlock()
	fake.OpenConnectionsStub = nil
	fake.openConnectionsReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeQueryableDB) OpenConnectionsReturnsOnCall(i int, result1 int) {
	fake.openConnectionsMutex.Lock()
	defer fake.openConnectionsMutex.Unlock()
	fake.OpenConnectionsStub = nil
	if fake.openConnectionsReturnsOnCall == nil {
		fake.openConnectionsReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.openConnectionsReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeQueryableDB) Prepare(arg1 string) (*sql.Stmt, error) {
	fake.prepareMutex.Lock()
	ret, specificReturn := fake.prepareReturnsOnCall[len(fake.prepareArgsForCall)]
	fake.prepareArgsForCall = append(fake.prepareArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Prepare", []interface{}{arg1})
	fake.prepareMutex.Unlock()
	if fake.PrepareStub != nil {
		return fake.PrepareStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.prepareReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeQueryableDB) PrepareCallCount() int {
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	return len(fake.prepareArgsForCall)
}

func (fake *FakeQueryableDB) PrepareCalls(stub func(string) (*sql.Stmt, error)) {
	fake.prepareMutex.Lock()
	defer fake.prepareMutex.Unlock()
	fake.PrepareStub = stub
}

func (fake *FakeQueryableDB) PrepareArgsForCall(i int) string {
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	argsForCall := fake.prepareArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeQueryableDB) PrepareReturns(result1 *sql.Stmt, result2 error) {
	fake.prepareMutex.Lock()
	defer fake.prepareMutex.Unlock()
	fake.PrepareStub = nil
	fake.prepareReturns = struct {
		result1 *sql.Stmt
		result2 error
	}{result1, result2}
}

func (fake *FakeQueryableDB) PrepareReturnsOnCall(i int, result1 *sql.Stmt, result2 error) {
	fake.prepareMutex.Lock()
	defer fake.prepareMutex.Unlock()
	fake.PrepareStub = nil
	if fake.prepareReturnsOnCall == nil {
		fake.prepareReturnsOnCall = make(map[int]struct {
			result1 *sql.Stmt
			result2 error
		})
	}
	fake.prepareReturnsOnCall[i] = struct {
		result1 *sql.Stmt
		result2 error
	}{result1, result2}
}

func (fake *FakeQueryableDB) Query(arg1 string, arg2 ...interface{}) (*sql.Rows, error) {
	fake.queryMutex.Lock()
	ret, specificReturn := fake.queryReturnsOnCall[len(fake.queryArgsForCall)]
	fake.queryArgsForCall = append(fake.queryArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Query", []interface{}{arg1, arg2})
	fake.queryMutex.Unlock()
	if fake.QueryStub != nil {
		return fake.QueryStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeQueryableDB) QueryCallCount() int {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return len(fake.queryArgsForCall)
}

func (fake *FakeQueryableDB) QueryCalls(stub func(string, ...interface{}) (*sql.Rows, error)) {
	fake.queryMutex.Lock()
	defer fake.queryMutex.Unlock()
	fake.QueryStub = stub
}

func (fake *FakeQueryableDB) QueryArgsForCall(i int) (string, []interface{}) {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	argsForCall := fake.queryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeQueryableDB) QueryReturns(result1 *sql.Rows, result2 error) {
	fake.queryMutex.Lock()
	defer fake.queryMutex.Unlock()
	fake.QueryStub = nil
	fake.queryReturns = struct {
		result1 *sql.Rows
		result2 error
	}{result1, result2}
}

func (fake *FakeQueryableDB) QueryReturnsOnCall(i int, result1 *sql.Rows, result2 error) {
	fake.queryMutex.Lock()
	defer fake.queryMutex.Unlock()
	fake.QueryStub = nil
	if fake.queryReturnsOnCall == nil {
		fake.queryReturnsOnCall = make(map[int]struct {
			result1 *sql.Rows
			result2 error
		})
	}
	fake.queryReturnsOnCall[i] = struct {
		result1 *sql.Rows
		result2 error
	}{result1, result2}
}

func (fake *FakeQueryableDB) QueryRow(arg1 string, arg2 ...interface{}) helpers.RowScanner {
	fake.queryRowMutex.Lock()
	ret, specificReturn := fake.queryRowReturnsOnCall[len(fake.queryRowArgsForCall)]
	fake.queryRowArgsForCall = append(fake.queryRowArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("QueryRow", []interface{}{arg1, arg2})
	fake.queryRowMutex.Unlock()
	if fake.QueryRowStub != nil {
		return fake.QueryRowStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.queryRowReturns
	return fakeReturns.result1
}

func (fake *FakeQueryableDB) QueryRowCallCount() int {
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	return len(fake.queryRowArgsForCall)
}

func (fake *FakeQueryableDB) QueryRowCalls(stub func(string, ...interface{}) helpers.RowScanner) {
	fake.queryRowMutex.Lock()
	defer fake.queryRowMutex.Unlock()
	fake.QueryRowStub = stub
}

func (fake *FakeQueryableDB) QueryRowArgsForCall(i int) (string, []interface{}) {
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	argsForCall := fake.queryRowArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeQueryableDB) QueryRowReturns(result1 helpers.RowScanner) {
	fake.queryRowMutex.Lock()
	defer fake.queryRowMutex.Unlock()
	fake.QueryRowStub = nil
	fake.queryRowReturns = struct {
		result1 helpers.RowScanner
	}{result1}
}

func (fake *FakeQueryableDB) QueryRowReturnsOnCall(i int, result1 helpers.RowScanner) {
	fake.queryRowMutex.Lock()
	defer fake.queryRowMutex.Unlock()
	fake.QueryRowStub = nil
	if fake.queryRowReturnsOnCall == nil {
		fake.queryRowReturnsOnCall = make(map[int]struct {
			result1 helpers.RowScanner
		})
	}
	fake.queryRowReturnsOnCall[i] = struct {
		result1 helpers.RowScanner
	}{result1}
}

func (fake *FakeQueryableDB) WaitCount() int64 {
	fake.waitCountMutex.Lock()
	ret, specificReturn := fake.waitCountReturnsOnCall[len(fake.waitCountArgsForCall)]
	fake.waitCountArgsForCall = append(fake.waitCountArgsForCall, struct {
	}{})
	fake.recordInvocation("WaitCount", []interface{}{})
	fake.waitCountMutex.Unlock()
	if fake.WaitCountStub != nil {
		return fake.WaitCountStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.waitCountReturns
	return fakeReturns.result1
}

func (fake *FakeQueryableDB) WaitCountCallCount() int {
	fake.waitCountMutex.RLock()
	defer fake.waitCountMutex.RUnlock()
	return len(fake.waitCountArgsForCall)
}

func (fake *FakeQueryableDB) WaitCountCalls(stub func() int64) {
	fake.waitCountMutex.Lock()
	defer fake.waitCountMutex.Unlock()
	fake.WaitCountStub = stub
}

func (fake *FakeQueryableDB) WaitCountReturns(result1 int64) {
	fake.waitCountMutex.Lock()
	defer fake.waitCountMutex.Unlock()
	fake.WaitCountStub = nil
	fake.waitCountReturns = struct {
		result1 int64
	}{result1}
}

func (fake *FakeQueryableDB) WaitCountReturnsOnCall(i int, result1 int64) {
	fake.waitCountMutex.Lock()
	defer fake.waitCountMutex.Unlock()
	fake.WaitCountStub = nil
	if fake.waitCountReturnsOnCall == nil {
		fake.waitCountReturnsOnCall = make(map[int]struct {
			result1 int64
		})
	}
	fake.waitCountReturnsOnCall[i] = struct {
		result1 int64
	}{result1}
}

func (fake *FakeQueryableDB) WaitDuration() time.Duration {
	fake.waitDurationMutex.Lock()
	ret, specificReturn := fake.waitDurationReturnsOnCall[len(fake.waitDurationArgsForCall)]
	fake.waitDurationArgsForCall = append(fake.waitDurationArgsForCall, struct {
	}{})
	fake.recordInvocation("WaitDuration", []interface{}{})
	fake.waitDurationMutex.Unlock()
	if fake.WaitDurationStub != nil {
		return fake.WaitDurationStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.waitDurationReturns
	return fakeReturns.result1
}

func (fake *FakeQueryableDB) WaitDurationCallCount() int {
	fake.waitDurationMutex.RLock()
	defer fake.waitDurationMutex.RUnlock()
	return len(fake.waitDurationArgsForCall)
}

func (fake *FakeQueryableDB) WaitDurationCalls(stub func() time.Duration) {
	fake.waitDurationMutex.Lock()
	defer fake.waitDurationMutex.Unlock()
	fake.WaitDurationStub = stub
}

func (fake *FakeQueryableDB) WaitDurationReturns(result1 time.Duration) {
	fake.waitDurationMutex.Lock()
	defer fake.waitDurationMutex.Unlock()
	fake.WaitDurationStub = nil
	fake.waitDurationReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeQueryableDB) WaitDurationReturnsOnCall(i int, result1 time.Duration) {
	fake.waitDurationMutex.Lock()
	defer fake.waitDurationMutex.Unlock()
	fake.WaitDurationStub = nil
	if fake.waitDurationReturnsOnCall == nil {
		fake.waitDurationReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.waitDurationReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeQueryableDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.beginMutex.RLock()
	defer fake.beginMutex.RUnlock()
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	fake.openConnectionsMutex.RLock()
	defer fake.openConnectionsMutex.RUnlock()
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	fake.waitCountMutex.RLock()
	defer fake.waitCountMutex.RUnlock()
	fake.waitDurationMutex.RLock()
	defer fake.waitDurationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeQueryableDB) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ helpers.QueryableDB = new(FakeQueryableDB)
