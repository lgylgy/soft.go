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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func discardENamedElement() {
	_ = assert.Equal
	_ = mock.Anything
	_ = testing.Coverage

	_ = time.Now()
}

func TestENamedElementNameGet(t *testing.T) {
	obj := newENamedElementImpl()
	obj.SetName("Test String")
	assert.Equal(t, "Test String", obj.GetName())
}

func TestENamedElementNameSet(t *testing.T) {
	mockAdapter := &MockEAdapter{}
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	obj := newENamedElementImpl()
	obj.EAdapters().Add(mockAdapter)
	obj.SetName("Test String")
	mockAdapter.AssertExpectations(t)
}

func TestENamedElementNameEGet(t *testing.T) {
	obj := newENamedElementImpl()
	{
		assert.Equal(t, obj.GetName(), obj.EGetFromID(ENAMED_ELEMENT__NAME, false, false))
	}
}

func TestENamedElementNameEIsSet(t *testing.T) {
	obj := newENamedElementImpl()
	{
		_ = obj
	}
}

func TestENamedElementNameEUnset(t *testing.T) {
	obj := newENamedElementImpl()
	{
		_ = obj
	}
}

func TestENamedElementNameESet(t *testing.T) {
	obj := newENamedElementImpl()
	{
		_ = obj
	}
}