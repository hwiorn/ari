// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/CyCoreSystems/ari (interfaces: LiveRecording)

package mock

import (
	ari "github.com/CyCoreSystems/ari"
	gomock "github.com/golang/mock/gomock"
)

// Mock of LiveRecording interface
type MockLiveRecording struct {
	ctrl     *gomock.Controller
	recorder *_MockLiveRecordingRecorder
}

// Recorder for MockLiveRecording (not exported)
type _MockLiveRecordingRecorder struct {
	mock *MockLiveRecording
}

func NewMockLiveRecording(ctrl *gomock.Controller) *MockLiveRecording {
	mock := &MockLiveRecording{ctrl: ctrl}
	mock.recorder = &_MockLiveRecordingRecorder{mock}
	return mock
}

func (_m *MockLiveRecording) EXPECT() *_MockLiveRecordingRecorder {
	return _m.recorder
}

func (_m *MockLiveRecording) Data(_param0 string) (*ari.LiveRecordingData, error) {
	ret := _m.ctrl.Call(_m, "Data", _param0)
	ret0, _ := ret[0].(*ari.LiveRecordingData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLiveRecordingRecorder) Data(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Data", arg0)
}

func (_m *MockLiveRecording) Delete(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Delete", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLiveRecordingRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0)
}

func (_m *MockLiveRecording) Get(_param0 string) *ari.LiveRecordingHandle {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(*ari.LiveRecordingHandle)
	return ret0
}

func (_mr *_MockLiveRecordingRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockLiveRecording) Mute(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Mute", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLiveRecordingRecorder) Mute(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Mute", arg0)
}

func (_m *MockLiveRecording) Pause(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Pause", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLiveRecordingRecorder) Pause(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Pause", arg0)
}

func (_m *MockLiveRecording) Resume(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Resume", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLiveRecordingRecorder) Resume(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Resume", arg0)
}

func (_m *MockLiveRecording) Scrap(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Scrap", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLiveRecordingRecorder) Scrap(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Scrap", arg0)
}

func (_m *MockLiveRecording) Stop(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Stop", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLiveRecordingRecorder) Stop(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop", arg0)
}

func (_m *MockLiveRecording) Unmute(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Unmute", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLiveRecordingRecorder) Unmute(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Unmute", arg0)
}
