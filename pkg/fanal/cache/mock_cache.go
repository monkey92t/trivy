// Code generated by mockery v1.0.0. DO NOT EDIT.

package cache

import (
	mock "github.com/stretchr/testify/mock"

	types "github.com/zhanglimao/trivy/pkg/fanal/types"
)

// MockCache is an autogenerated mock type for the Cache type
type MockCache struct {
	mock.Mock
}

type CacheClearReturns struct {
	Err error
}

type CacheClearExpectation struct {
	Returns CacheClearReturns
}

func (_m *MockCache) ApplyClearExpectation(e CacheClearExpectation) {
	var args []interface{}
	_m.On("Clear", args...).Return(e.Returns.Err)
}

func (_m *MockCache) ApplyClearExpectations(expectations []CacheClearExpectation) {
	for _, e := range expectations {
		_m.ApplyClearExpectation(e)
	}
}

// Clear provides a mock function with given fields:
func (_m *MockCache) Clear() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type CacheCloseReturns struct {
	Err error
}

type CacheCloseExpectation struct {
	Returns CacheCloseReturns
}

func (_m *MockCache) ApplyCloseExpectation(e CacheCloseExpectation) {
	var args []interface{}
	_m.On("Close", args...).Return(e.Returns.Err)
}

func (_m *MockCache) ApplyCloseExpectations(expectations []CacheCloseExpectation) {
	for _, e := range expectations {
		_m.ApplyCloseExpectation(e)
	}
}

// Close provides a mock function with given fields:
func (_m *MockCache) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type CacheDeleteBlobArgs struct {
	BlobID         string
	BlobIDAnything bool
}

type CacheDeleteBlobReturns struct {
	_a0 error
}

type CacheDeleteBlobExpectation struct {
	Args    CacheDeleteBlobArgs
	Returns CacheDeleteBlobReturns
}

func (_m *MockCache) ApplyDeleteBlobExpectation(e CacheDeleteBlobExpectation) {
	var args []interface{}
	if e.Args.BlobIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.BlobID)
	}
	_m.On("DeleteBlob", args...).Return(e.Returns._a0)
}

func (_m *MockCache) ApplyDeleteBlobExpectations(expectations []CacheDeleteBlobExpectation) {
	for _, e := range expectations {
		_m.ApplyDeleteBlobExpectation(e)
	}
}

// DeleteBlob provides a mock function with given fields: blobID
func (_m *MockCache) DeleteBlob(blobID string) error {
	ret := _m.Called(blobID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(blobID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type CacheGetArtifactArgs struct {
	ArtifactID         string
	ArtifactIDAnything bool
}

type CacheGetArtifactReturns struct {
	ArtifactInfo types.ArtifactInfo
	Err          error
}

type CacheGetArtifactExpectation struct {
	Args    CacheGetArtifactArgs
	Returns CacheGetArtifactReturns
}

func (_m *MockCache) ApplyGetArtifactExpectation(e CacheGetArtifactExpectation) {
	var args []interface{}
	if e.Args.ArtifactIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.ArtifactID)
	}
	_m.On("GetArtifact", args...).Return(e.Returns.ArtifactInfo, e.Returns.Err)
}

func (_m *MockCache) ApplyGetArtifactExpectations(expectations []CacheGetArtifactExpectation) {
	for _, e := range expectations {
		_m.ApplyGetArtifactExpectation(e)
	}
}

// GetArtifact provides a mock function with given fields: artifactID
func (_m *MockCache) GetArtifact(artifactID string) (types.ArtifactInfo, error) {
	ret := _m.Called(artifactID)

	var r0 types.ArtifactInfo
	if rf, ok := ret.Get(0).(func(string) types.ArtifactInfo); ok {
		r0 = rf(artifactID)
	} else {
		r0 = ret.Get(0).(types.ArtifactInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(artifactID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type CacheGetBlobArgs struct {
	BlobID         string
	BlobIDAnything bool
}

type CacheGetBlobReturns struct {
	BlobInfo types.BlobInfo
	Err      error
}

type CacheGetBlobExpectation struct {
	Args    CacheGetBlobArgs
	Returns CacheGetBlobReturns
}

func (_m *MockCache) ApplyGetBlobExpectation(e CacheGetBlobExpectation) {
	var args []interface{}
	if e.Args.BlobIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.BlobID)
	}
	_m.On("GetBlob", args...).Return(e.Returns.BlobInfo, e.Returns.Err)
}

func (_m *MockCache) ApplyGetBlobExpectations(expectations []CacheGetBlobExpectation) {
	for _, e := range expectations {
		_m.ApplyGetBlobExpectation(e)
	}
}

// GetBlob provides a mock function with given fields: blobID
func (_m *MockCache) GetBlob(blobID string) (types.BlobInfo, error) {
	ret := _m.Called(blobID)

	var r0 types.BlobInfo
	if rf, ok := ret.Get(0).(func(string) types.BlobInfo); ok {
		r0 = rf(blobID)
	} else {
		r0 = ret.Get(0).(types.BlobInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(blobID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type CacheMissingBlobsArgs struct {
	ArtifactID         string
	ArtifactIDAnything bool
	BlobIDs            []string
	BlobIDsAnything    bool
}

type CacheMissingBlobsReturns struct {
	MissingArtifact bool
	MissingBlobIDs  []string
	Err             error
}

type CacheMissingBlobsExpectation struct {
	Args    CacheMissingBlobsArgs
	Returns CacheMissingBlobsReturns
}

func (_m *MockCache) ApplyMissingBlobsExpectation(e CacheMissingBlobsExpectation) {
	var args []interface{}
	if e.Args.ArtifactIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.ArtifactID)
	}
	if e.Args.BlobIDsAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.BlobIDs)
	}
	_m.On("MissingBlobs", args...).Return(e.Returns.MissingArtifact, e.Returns.MissingBlobIDs, e.Returns.Err)
}

func (_m *MockCache) ApplyMissingBlobsExpectations(expectations []CacheMissingBlobsExpectation) {
	for _, e := range expectations {
		_m.ApplyMissingBlobsExpectation(e)
	}
}

// MissingBlobs provides a mock function with given fields: artifactID, blobIDs
func (_m *MockCache) MissingBlobs(artifactID string, blobIDs []string) (bool, []string, error) {
	ret := _m.Called(artifactID, blobIDs)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, []string) bool); ok {
		r0 = rf(artifactID, blobIDs)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 []string
	if rf, ok := ret.Get(1).(func(string, []string) []string); ok {
		r1 = rf(artifactID, blobIDs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, []string) error); ok {
		r2 = rf(artifactID, blobIDs)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type CachePutArtifactArgs struct {
	ArtifactID           string
	ArtifactIDAnything   bool
	ArtifactInfo         types.ArtifactInfo
	ArtifactInfoAnything bool
}

type CachePutArtifactReturns struct {
	Err error
}

type CachePutArtifactExpectation struct {
	Args    CachePutArtifactArgs
	Returns CachePutArtifactReturns
}

func (_m *MockCache) ApplyPutArtifactExpectation(e CachePutArtifactExpectation) {
	var args []interface{}
	if e.Args.ArtifactIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.ArtifactID)
	}
	if e.Args.ArtifactInfoAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.ArtifactInfo)
	}
	_m.On("PutArtifact", args...).Return(e.Returns.Err)
}

func (_m *MockCache) ApplyPutArtifactExpectations(expectations []CachePutArtifactExpectation) {
	for _, e := range expectations {
		_m.ApplyPutArtifactExpectation(e)
	}
}

// PutArtifact provides a mock function with given fields: artifactID, artifactInfo
func (_m *MockCache) PutArtifact(artifactID string, artifactInfo types.ArtifactInfo) error {
	ret := _m.Called(artifactID, artifactInfo)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, types.ArtifactInfo) error); ok {
		r0 = rf(artifactID, artifactInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type CachePutBlobArgs struct {
	BlobID           string
	BlobIDAnything   bool
	BlobInfo         types.BlobInfo
	BlobInfoAnything bool
}

type CachePutBlobReturns struct {
	Err error
}

type CachePutBlobExpectation struct {
	Args    CachePutBlobArgs
	Returns CachePutBlobReturns
}

func (_m *MockCache) ApplyPutBlobExpectation(e CachePutBlobExpectation) {
	var args []interface{}
	if e.Args.BlobIDAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.BlobID)
	}
	if e.Args.BlobInfoAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.BlobInfo)
	}
	_m.On("PutBlob", args...).Return(e.Returns.Err)
}

func (_m *MockCache) ApplyPutBlobExpectations(expectations []CachePutBlobExpectation) {
	for _, e := range expectations {
		_m.ApplyPutBlobExpectation(e)
	}
}

// PutBlob provides a mock function with given fields: blobID, blobInfo
func (_m *MockCache) PutBlob(blobID string, blobInfo types.BlobInfo) error {
	ret := _m.Called(blobID, blobInfo)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, types.BlobInfo) error); ok {
		r0 = rf(blobID, blobInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
