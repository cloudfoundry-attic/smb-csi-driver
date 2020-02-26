// Code generated by counterfeiter. DO NOT EDIT.
package smbcsidriverfakes

import (
	"sync"

	v1a "k8s.io/api/core/v1"
	v1b "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type FakeConfigMapInterface struct {
	CreateStub        func(*v1a.ConfigMap) (*v1a.ConfigMap, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 *v1a.ConfigMap
	}
	createReturns struct {
		result1 *v1a.ConfigMap
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *v1a.ConfigMap
		result2 error
	}
	DeleteStub        func(string, *v1b.DeleteOptions) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
		arg2 *v1b.DeleteOptions
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteCollectionStub        func(*v1b.DeleteOptions, v1b.ListOptions) error
	deleteCollectionMutex       sync.RWMutex
	deleteCollectionArgsForCall []struct {
		arg1 *v1b.DeleteOptions
		arg2 v1b.ListOptions
	}
	deleteCollectionReturns struct {
		result1 error
	}
	deleteCollectionReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(string, v1b.GetOptions) (*v1a.ConfigMap, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
		arg2 v1b.GetOptions
	}
	getReturns struct {
		result1 *v1a.ConfigMap
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *v1a.ConfigMap
		result2 error
	}
	ListStub        func(v1b.ListOptions) (*v1a.ConfigMapList, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 v1b.ListOptions
	}
	listReturns struct {
		result1 *v1a.ConfigMapList
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 *v1a.ConfigMapList
		result2 error
	}
	PatchStub        func(string, types.PatchType, []byte, ...string) (*v1a.ConfigMap, error)
	patchMutex       sync.RWMutex
	patchArgsForCall []struct {
		arg1 string
		arg2 types.PatchType
		arg3 []byte
		arg4 []string
	}
	patchReturns struct {
		result1 *v1a.ConfigMap
		result2 error
	}
	patchReturnsOnCall map[int]struct {
		result1 *v1a.ConfigMap
		result2 error
	}
	UpdateStub        func(*v1a.ConfigMap) (*v1a.ConfigMap, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 *v1a.ConfigMap
	}
	updateReturns struct {
		result1 *v1a.ConfigMap
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 *v1a.ConfigMap
		result2 error
	}
	WatchStub        func(v1b.ListOptions) (watch.Interface, error)
	watchMutex       sync.RWMutex
	watchArgsForCall []struct {
		arg1 v1b.ListOptions
	}
	watchReturns struct {
		result1 watch.Interface
		result2 error
	}
	watchReturnsOnCall map[int]struct {
		result1 watch.Interface
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfigMapInterface) Create(arg1 *v1a.ConfigMap) (*v1a.ConfigMap, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 *v1a.ConfigMap
	}{arg1})
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConfigMapInterface) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeConfigMapInterface) CreateCalls(stub func(*v1a.ConfigMap) (*v1a.ConfigMap, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeConfigMapInterface) CreateArgsForCall(i int) *v1a.ConfigMap {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfigMapInterface) CreateReturns(result1 *v1a.ConfigMap, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *v1a.ConfigMap
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigMapInterface) CreateReturnsOnCall(i int, result1 *v1a.ConfigMap, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *v1a.ConfigMap
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *v1a.ConfigMap
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigMapInterface) Delete(arg1 string, arg2 *v1b.DeleteOptions) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
		arg2 *v1b.DeleteOptions
	}{arg1, arg2})
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *FakeConfigMapInterface) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeConfigMapInterface) DeleteCalls(stub func(string, *v1b.DeleteOptions) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeConfigMapInterface) DeleteArgsForCall(i int) (string, *v1b.DeleteOptions) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeConfigMapInterface) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigMapInterface) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigMapInterface) DeleteCollection(arg1 *v1b.DeleteOptions, arg2 v1b.ListOptions) error {
	fake.deleteCollectionMutex.Lock()
	ret, specificReturn := fake.deleteCollectionReturnsOnCall[len(fake.deleteCollectionArgsForCall)]
	fake.deleteCollectionArgsForCall = append(fake.deleteCollectionArgsForCall, struct {
		arg1 *v1b.DeleteOptions
		arg2 v1b.ListOptions
	}{arg1, arg2})
	fake.recordInvocation("DeleteCollection", []interface{}{arg1, arg2})
	fake.deleteCollectionMutex.Unlock()
	if fake.DeleteCollectionStub != nil {
		return fake.DeleteCollectionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteCollectionReturns
	return fakeReturns.result1
}

func (fake *FakeConfigMapInterface) DeleteCollectionCallCount() int {
	fake.deleteCollectionMutex.RLock()
	defer fake.deleteCollectionMutex.RUnlock()
	return len(fake.deleteCollectionArgsForCall)
}

func (fake *FakeConfigMapInterface) DeleteCollectionCalls(stub func(*v1b.DeleteOptions, v1b.ListOptions) error) {
	fake.deleteCollectionMutex.Lock()
	defer fake.deleteCollectionMutex.Unlock()
	fake.DeleteCollectionStub = stub
}

func (fake *FakeConfigMapInterface) DeleteCollectionArgsForCall(i int) (*v1b.DeleteOptions, v1b.ListOptions) {
	fake.deleteCollectionMutex.RLock()
	defer fake.deleteCollectionMutex.RUnlock()
	argsForCall := fake.deleteCollectionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeConfigMapInterface) DeleteCollectionReturns(result1 error) {
	fake.deleteCollectionMutex.Lock()
	defer fake.deleteCollectionMutex.Unlock()
	fake.DeleteCollectionStub = nil
	fake.deleteCollectionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigMapInterface) DeleteCollectionReturnsOnCall(i int, result1 error) {
	fake.deleteCollectionMutex.Lock()
	defer fake.deleteCollectionMutex.Unlock()
	fake.DeleteCollectionStub = nil
	if fake.deleteCollectionReturnsOnCall == nil {
		fake.deleteCollectionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteCollectionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigMapInterface) Get(arg1 string, arg2 v1b.GetOptions) (*v1a.ConfigMap, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
		arg2 v1b.GetOptions
	}{arg1, arg2})
	fake.recordInvocation("Get", []interface{}{arg1, arg2})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConfigMapInterface) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeConfigMapInterface) GetCalls(stub func(string, v1b.GetOptions) (*v1a.ConfigMap, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeConfigMapInterface) GetArgsForCall(i int) (string, v1b.GetOptions) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeConfigMapInterface) GetReturns(result1 *v1a.ConfigMap, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *v1a.ConfigMap
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigMapInterface) GetReturnsOnCall(i int, result1 *v1a.ConfigMap, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *v1a.ConfigMap
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *v1a.ConfigMap
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigMapInterface) List(arg1 v1b.ListOptions) (*v1a.ConfigMapList, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 v1b.ListOptions
	}{arg1})
	fake.recordInvocation("List", []interface{}{arg1})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConfigMapInterface) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeConfigMapInterface) ListCalls(stub func(v1b.ListOptions) (*v1a.ConfigMapList, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeConfigMapInterface) ListArgsForCall(i int) v1b.ListOptions {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfigMapInterface) ListReturns(result1 *v1a.ConfigMapList, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 *v1a.ConfigMapList
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigMapInterface) ListReturnsOnCall(i int, result1 *v1a.ConfigMapList, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 *v1a.ConfigMapList
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 *v1a.ConfigMapList
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigMapInterface) Patch(arg1 string, arg2 types.PatchType, arg3 []byte, arg4 ...string) (*v1a.ConfigMap, error) {
	var arg3Copy []byte
	if arg3 != nil {
		arg3Copy = make([]byte, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.patchMutex.Lock()
	ret, specificReturn := fake.patchReturnsOnCall[len(fake.patchArgsForCall)]
	fake.patchArgsForCall = append(fake.patchArgsForCall, struct {
		arg1 string
		arg2 types.PatchType
		arg3 []byte
		arg4 []string
	}{arg1, arg2, arg3Copy, arg4})
	fake.recordInvocation("Patch", []interface{}{arg1, arg2, arg3Copy, arg4})
	fake.patchMutex.Unlock()
	if fake.PatchStub != nil {
		return fake.PatchStub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.patchReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConfigMapInterface) PatchCallCount() int {
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	return len(fake.patchArgsForCall)
}

func (fake *FakeConfigMapInterface) PatchCalls(stub func(string, types.PatchType, []byte, ...string) (*v1a.ConfigMap, error)) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = stub
}

func (fake *FakeConfigMapInterface) PatchArgsForCall(i int) (string, types.PatchType, []byte, []string) {
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	argsForCall := fake.patchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeConfigMapInterface) PatchReturns(result1 *v1a.ConfigMap, result2 error) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = nil
	fake.patchReturns = struct {
		result1 *v1a.ConfigMap
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigMapInterface) PatchReturnsOnCall(i int, result1 *v1a.ConfigMap, result2 error) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = nil
	if fake.patchReturnsOnCall == nil {
		fake.patchReturnsOnCall = make(map[int]struct {
			result1 *v1a.ConfigMap
			result2 error
		})
	}
	fake.patchReturnsOnCall[i] = struct {
		result1 *v1a.ConfigMap
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigMapInterface) Update(arg1 *v1a.ConfigMap) (*v1a.ConfigMap, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 *v1a.ConfigMap
	}{arg1})
	fake.recordInvocation("Update", []interface{}{arg1})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConfigMapInterface) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeConfigMapInterface) UpdateCalls(stub func(*v1a.ConfigMap) (*v1a.ConfigMap, error)) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeConfigMapInterface) UpdateArgsForCall(i int) *v1a.ConfigMap {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfigMapInterface) UpdateReturns(result1 *v1a.ConfigMap, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 *v1a.ConfigMap
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigMapInterface) UpdateReturnsOnCall(i int, result1 *v1a.ConfigMap, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 *v1a.ConfigMap
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 *v1a.ConfigMap
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigMapInterface) Watch(arg1 v1b.ListOptions) (watch.Interface, error) {
	fake.watchMutex.Lock()
	ret, specificReturn := fake.watchReturnsOnCall[len(fake.watchArgsForCall)]
	fake.watchArgsForCall = append(fake.watchArgsForCall, struct {
		arg1 v1b.ListOptions
	}{arg1})
	fake.recordInvocation("Watch", []interface{}{arg1})
	fake.watchMutex.Unlock()
	if fake.WatchStub != nil {
		return fake.WatchStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.watchReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConfigMapInterface) WatchCallCount() int {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	return len(fake.watchArgsForCall)
}

func (fake *FakeConfigMapInterface) WatchCalls(stub func(v1b.ListOptions) (watch.Interface, error)) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = stub
}

func (fake *FakeConfigMapInterface) WatchArgsForCall(i int) v1b.ListOptions {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	argsForCall := fake.watchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfigMapInterface) WatchReturns(result1 watch.Interface, result2 error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	fake.watchReturns = struct {
		result1 watch.Interface
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigMapInterface) WatchReturnsOnCall(i int, result1 watch.Interface, result2 error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	if fake.watchReturnsOnCall == nil {
		fake.watchReturnsOnCall = make(map[int]struct {
			result1 watch.Interface
			result2 error
		})
	}
	fake.watchReturnsOnCall[i] = struct {
		result1 watch.Interface
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigMapInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deleteCollectionMutex.RLock()
	defer fake.deleteCollectionMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConfigMapInterface) recordInvocation(key string, args []interface{}) {
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

var _ v1.ConfigMapInterface = new(FakeConfigMapInterface)