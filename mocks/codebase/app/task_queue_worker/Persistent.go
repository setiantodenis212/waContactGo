// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	taskqueueworker "github.com/golangid/candi/codebase/app/task_queue_worker"
	mock "github.com/stretchr/testify/mock"
)

// Persistent is an autogenerated mock type for the Persistent type
type Persistent struct {
	mock.Mock
}

// AggregateAllTaskJob provides a mock function with given fields: ctx, filter
func (_m *Persistent) AggregateAllTaskJob(ctx context.Context, filter taskqueueworker.Filter) []taskqueueworker.TaskResolver {
	ret := _m.Called(ctx, filter)

	var r0 []taskqueueworker.TaskResolver
	if rf, ok := ret.Get(0).(func(context.Context, taskqueueworker.Filter) []taskqueueworker.TaskResolver); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]taskqueueworker.TaskResolver)
		}
	}

	return r0
}

// CleanJob provides a mock function with given fields: ctx, taskName
func (_m *Persistent) CleanJob(ctx context.Context, taskName string) {
	_m.Called(ctx, taskName)
}

// CountAllJob provides a mock function with given fields: ctx, filter
func (_m *Persistent) CountAllJob(ctx context.Context, filter taskqueueworker.Filter) int {
	ret := _m.Called(ctx, filter)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, taskqueueworker.Filter) int); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// FindAllJob provides a mock function with given fields: ctx, filter
func (_m *Persistent) FindAllJob(ctx context.Context, filter taskqueueworker.Filter) []taskqueueworker.Job {
	ret := _m.Called(ctx, filter)

	var r0 []taskqueueworker.Job
	if rf, ok := ret.Get(0).(func(context.Context, taskqueueworker.Filter) []taskqueueworker.Job); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]taskqueueworker.Job)
		}
	}

	return r0
}

// FindJobByID provides a mock function with given fields: ctx, id
func (_m *Persistent) FindJobByID(ctx context.Context, id string) (*taskqueueworker.Job, error) {
	ret := _m.Called(ctx, id)

	var r0 *taskqueueworker.Job
	if rf, ok := ret.Get(0).(func(context.Context, string) *taskqueueworker.Job); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taskqueueworker.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveJob provides a mock function with given fields: ctx, job
func (_m *Persistent) SaveJob(ctx context.Context, job *taskqueueworker.Job) {
	_m.Called(ctx, job)
}

// UpdateAllStatus provides a mock function with given fields: ctx, taskName, currentStatus, updatedStatus
func (_m *Persistent) UpdateAllStatus(ctx context.Context, taskName string, currentStatus []taskqueueworker.JobStatusEnum, updatedStatus taskqueueworker.JobStatusEnum) {
	_m.Called(ctx, taskName, currentStatus, updatedStatus)
}