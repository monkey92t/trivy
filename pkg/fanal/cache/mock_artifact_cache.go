// Code generated by mockery v1.0.0. DO NOT EDIT.

package cache

import (
	mock "github.com/stretchr/testify/mock"

	types "github.com/zhanglimao/trivy/pkg/fanal/types"
)

// MockArtifactCache is an autogenerated mock type for the ArtifactCache type
type MockArtifactCache struct {
	mock.Mock
}

type ArtifactCacheDeleteBlobsArgs struct {
	BlobIDs         []string
	BlobIDsAnything bool
}

type ArtifactCacheDeleteBlobsReturns struct {
	_a0 error
}

type ArtifactCacheDeleteBlobsExpectation struct {
	Args    ArtifactCacheDeleteBlobsArgs
	Returns ArtifactCacheDeleteBlobsReturns
}

func (_m *MockArtifactCache) ApplyDeleteBlobsExpectation(e ArtifactCacheDeleteBlobsExpectation) {
	var args []interface{}
	if e.Args.BlobIDsAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.BlobIDs)
	}
	_m.On("DeleteBlobs", args...).Return(e.Returns._a0)
}

func (_m *MockArtifactCache) ApplyDeleteBlobsExpectations(expectations []ArtifactCacheDeleteBlobsExpectation) {
	for _, e := range expectations {
		_m.ApplyDeleteBlobsExpectation(e)
	}
}

// DeleteBlobs provides a mock function with given fields: blobIDs
func (_m *MockArtifactCache) DeleteBlobs(blobIDs []string) error {
	ret := _m.Called(blobIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(blobIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type ArtifactCacheMissingBlobsArgs struct {
	ArtifactID         string
	ArtifactIDAnything bool
	BlobIDs            []string
	BlobIDsAnything    bool
}

type ArtifactCacheMissingBlobsReturns struct {
	MissingArtifact bool
	MissingBlobIDs  []string
	Err             error
}

type ArtifactCacheMissingBlobsExpectation struct {
	Args    ArtifactCacheMissingBlobsArgs
	Returns ArtifactCacheMissingBlobsReturns
}

func (_m *MockArtifactCache) ApplyMissingBlobsExpectation(e ArtifactCacheMissingBlobsExpectation) {
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

func (_m *MockArtifactCache) ApplyMissingBlobsExpectations(expectations []ArtifactCacheMissingBlobsExpectation) {
	for _, e := range expectations {
		_m.ApplyMissingBlobsExpectation(e)
	}
}

// MissingBlobs provides a mock function with given fields: artifactID, blobIDs
func (_m *MockArtifactCache) MissingBlobs(artifactID string, blobIDs []string) (bool, []string, error) {
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

type ArtifactCachePutArtifactArgs struct {
	ArtifactID           string
	ArtifactIDAnything   bool
	ArtifactInfo         types.ArtifactInfo
	ArtifactInfoAnything bool
}

type ArtifactCachePutArtifactReturns struct {
	Err error
}

type ArtifactCachePutArtifactExpectation struct {
	Args    ArtifactCachePutArtifactArgs
	Returns ArtifactCachePutArtifactReturns
}

func (_m *MockArtifactCache) ApplyPutArtifactExpectation(e ArtifactCachePutArtifactExpectation) {
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

func (_m *MockArtifactCache) ApplyPutArtifactExpectations(expectations []ArtifactCachePutArtifactExpectation) {
	for _, e := range expectations {
		_m.ApplyPutArtifactExpectation(e)
	}
}

// PutArtifact provides a mock function with given fields: artifactID, artifactInfo
func (_m *MockArtifactCache) PutArtifact(artifactID string, artifactInfo types.ArtifactInfo) error {
	ret := _m.Called(artifactID, artifactInfo)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, types.ArtifactInfo) error); ok {
		r0 = rf(artifactID, artifactInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type ArtifactCachePutBlobArgs struct {
	BlobID           string
	BlobIDAnything   bool
	BlobInfo         types.BlobInfo
	BlobInfoAnything bool
}

type ArtifactCachePutBlobReturns struct {
	Err error
}

type ArtifactCachePutBlobExpectation struct {
	Args    ArtifactCachePutBlobArgs
	Returns ArtifactCachePutBlobReturns
}

func (_m *MockArtifactCache) ApplyPutBlobExpectation(e ArtifactCachePutBlobExpectation) {
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

func (_m *MockArtifactCache) ApplyPutBlobExpectations(expectations []ArtifactCachePutBlobExpectation) {
	for _, e := range expectations {
		_m.ApplyPutBlobExpectation(e)
	}
}

// PutBlob provides a mock function with given fields: blobID, blobInfo
func (_m *MockArtifactCache) PutBlob(blobID string, blobInfo types.BlobInfo) error {

	for i := range blobInfo.Misconfigurations {
		// suppress misconfiguration code block
		for j := range blobInfo.Misconfigurations[i].Failures {
			blobInfo.Misconfigurations[i].Failures[j].Code = types.Code{}
		}
		for j := range blobInfo.Misconfigurations[i].Successes {
			blobInfo.Misconfigurations[i].Successes[j].Code = types.Code{}
		}
		for j := range blobInfo.Misconfigurations[i].Warnings {
			blobInfo.Misconfigurations[i].Warnings[j].Code = types.Code{}
		}
		for j := range blobInfo.Misconfigurations[i].Exceptions {
			blobInfo.Misconfigurations[i].Exceptions[j].Code = types.Code{}
		}
	}

	ret := _m.Called(blobID, blobInfo)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, types.BlobInfo) error); ok {
		r0 = rf(blobID, blobInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
