// *****************************************************************************
//
// This file is part of a MASA library or program.
// Refer to the included end-user license agreement for restrictions.
//
// Copyright (c) 2019 MASA Group
//
// *****************************************************************************

// *****************************************************************************
//
// Warning: This file was generated by soft.generator.go Generator
//
// *****************************************************************************

package ecore

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func discardEAttribute() {
	_ = assert.Equal
	_ = mock.Anything
	_ = testing.Coverage

}

func TestEAttributeIDGet(t *testing.T) {
	var newValue bool = true
	obj := newEAttributeImpl()
	obj.SetID(newValue)
	assert.Equal(t, newValue, obj.IsID())
}

func TestEAttributeIDSet(t *testing.T) {
	var newValue bool = true
	obj := newEAttributeImpl()
	mockAdapter := &MockEAdapter{}
	mockAdapter.On("SetTarget", obj).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	obj.EAdapters().Add(mockAdapter)
	obj.SetID(newValue)
	mockAdapter.AssertExpectations(t)
}
