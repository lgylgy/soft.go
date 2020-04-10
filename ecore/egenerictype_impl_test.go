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

func discardEGenericType() {
	_ = assert.Equal
	_ = mock.Anything
	_ = testing.Coverage

	_ = time.Now()
}

func TestEGenericTypeEUpperBoundSet(t *testing.T) {
	obj := newEGenericTypeImpl()
	mockAdapter := &MockEAdapter{}
	mockAdapter.On("SetTarget", obj).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	obj.EAdapters().Add(mockAdapter)
	obj.SetEUpperBound(newEGenericTypeImpl())
	mockAdapter.AssertExpectations(t)
}

func TestEGenericTypeEUpperBoundEGet(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		assert.Equal(t, obj.GetEClassifier(), obj.EGetFromID(EGENERIC_TYPE__ECLASSIFIER, false, false))
	}
	{
		assert.Equal(t, obj.GetELowerBound(), obj.EGetFromID(EGENERIC_TYPE__ELOWER_BOUND, false, false))
	}
	{
		assert.Equal(t, obj.GetERawType(), obj.EGetFromID(EGENERIC_TYPE__ERAW_TYPE, false, false))
	}
	{
		assert.Equal(t, obj.GetETypeArguments(), obj.EGetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS, false, false))
	}
	{
		assert.Equal(t, obj.GetETypeParameter(), obj.EGetFromID(EGENERIC_TYPE__ETYPE_PARAMETER, false, false))
	}
	{
		assert.Equal(t, obj.GetEUpperBound(), obj.EGetFromID(EGENERIC_TYPE__EUPPER_BOUND, false, false))
	}
}

func TestEGenericTypeEUpperBoundEInvoke(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		assert.Panics(t, func() { obj.EInvokeFromID(EGENERIC_TYPE__IS_INSTANCE_EJAVAOBJECT, nil) })
	}
}

func TestEGenericTypeEUpperBoundEIsSet(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		assert.Equal(t, obj.eTypeArguments != nil && obj.eTypeArguments.Size() != 0, obj.EIsSetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEGenericTypeEUpperBoundEUnset(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		obj.EUnsetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS)
		obj.GetETypeArguments().Clear()
		assert.Equal(t, 0, obj.GetETypeArguments().Size())
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEGenericTypeEUpperBoundESet(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		elem := NewEmptyBasicEList()
		obj.ESetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS, elem)
		assert.Equal(t, 1, obj.GetETypeArguments().Size())
		assert.Equal(t, elem, obj.GetETypeArguments().Get(0))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEGenericTypeEUpperBoundEInverseRemove(t *testing.T) {
	{
	}
	{
		obj := newEGenericTypeImpl()
		mock := &MockEObject{}
		obj.GetETypeArguments().Add(mock)
		obj.EBasicInverseRemove(mock, EGENERIC_TYPE__ETYPE_ARGUMENTS, &MockENotificationChain{})
		assert.Equal(t, 0, obj.GetETypeArguments().Size())
	}
	{
	}
}

func TestEGenericTypeETypeArgumentsGetList(t *testing.T) {
	obj := newEGenericTypeImpl()
	assert.NotNil(t, obj.GetETypeArguments())
}

func TestEGenericTypeETypeArgumentsEGet(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		assert.Equal(t, obj.GetEClassifier(), obj.EGetFromID(EGENERIC_TYPE__ECLASSIFIER, false, false))
	}
	{
		assert.Equal(t, obj.GetELowerBound(), obj.EGetFromID(EGENERIC_TYPE__ELOWER_BOUND, false, false))
	}
	{
		assert.Equal(t, obj.GetERawType(), obj.EGetFromID(EGENERIC_TYPE__ERAW_TYPE, false, false))
	}
	{
		assert.Equal(t, obj.GetETypeArguments(), obj.EGetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS, false, false))
	}
	{
		assert.Equal(t, obj.GetETypeParameter(), obj.EGetFromID(EGENERIC_TYPE__ETYPE_PARAMETER, false, false))
	}
	{
		assert.Equal(t, obj.GetEUpperBound(), obj.EGetFromID(EGENERIC_TYPE__EUPPER_BOUND, false, false))
	}
}

func TestEGenericTypeETypeArgumentsEInvoke(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		assert.Panics(t, func() { obj.EInvokeFromID(EGENERIC_TYPE__IS_INSTANCE_EJAVAOBJECT, nil) })
	}
}

func TestEGenericTypeETypeArgumentsEIsSet(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		assert.Equal(t, obj.eTypeArguments != nil && obj.eTypeArguments.Size() != 0, obj.EIsSetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEGenericTypeETypeArgumentsEUnset(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		obj.EUnsetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS)
		obj.GetETypeArguments().Clear()
		assert.Equal(t, 0, obj.GetETypeArguments().Size())
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEGenericTypeETypeArgumentsESet(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		elem := NewEmptyBasicEList()
		obj.ESetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS, elem)
		assert.Equal(t, 1, obj.GetETypeArguments().Size())
		assert.Equal(t, elem, obj.GetETypeArguments().Get(0))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEGenericTypeETypeArgumentsEInverseRemove(t *testing.T) {
	{
	}
	{
		obj := newEGenericTypeImpl()
		mock := &MockEObject{}
		obj.GetETypeArguments().Add(mock)
		obj.EBasicInverseRemove(mock, EGENERIC_TYPE__ETYPE_ARGUMENTS, &MockENotificationChain{})
		assert.Equal(t, 0, obj.GetETypeArguments().Size())
	}
	{
	}
}

func TestEGenericTypeERawTypeEGet(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		assert.Equal(t, obj.GetEClassifier(), obj.EGetFromID(EGENERIC_TYPE__ECLASSIFIER, false, false))
	}
	{
		assert.Equal(t, obj.GetELowerBound(), obj.EGetFromID(EGENERIC_TYPE__ELOWER_BOUND, false, false))
	}
	{
		assert.Equal(t, obj.GetERawType(), obj.EGetFromID(EGENERIC_TYPE__ERAW_TYPE, false, false))
	}
	{
		assert.Equal(t, obj.GetETypeArguments(), obj.EGetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS, false, false))
	}
	{
		assert.Equal(t, obj.GetETypeParameter(), obj.EGetFromID(EGENERIC_TYPE__ETYPE_PARAMETER, false, false))
	}
	{
		assert.Equal(t, obj.GetEUpperBound(), obj.EGetFromID(EGENERIC_TYPE__EUPPER_BOUND, false, false))
	}
}

func TestEGenericTypeERawTypeEInvoke(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		assert.Panics(t, func() { obj.EInvokeFromID(EGENERIC_TYPE__IS_INSTANCE_EJAVAOBJECT, nil) })
	}
}

func TestEGenericTypeERawTypeEIsSet(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		assert.Equal(t, obj.eTypeArguments != nil && obj.eTypeArguments.Size() != 0, obj.EIsSetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEGenericTypeERawTypeEUnset(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		obj.EUnsetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS)
		obj.GetETypeArguments().Clear()
		assert.Equal(t, 0, obj.GetETypeArguments().Size())
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEGenericTypeERawTypeESet(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		elem := NewEmptyBasicEList()
		obj.ESetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS, elem)
		assert.Equal(t, 1, obj.GetETypeArguments().Size())
		assert.Equal(t, elem, obj.GetETypeArguments().Get(0))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEGenericTypeERawTypeEInverseRemove(t *testing.T) {
	{
	}
	{
		obj := newEGenericTypeImpl()
		mock := &MockEObject{}
		obj.GetETypeArguments().Add(mock)
		obj.EBasicInverseRemove(mock, EGENERIC_TYPE__ETYPE_ARGUMENTS, &MockENotificationChain{})
		assert.Equal(t, 0, obj.GetETypeArguments().Size())
	}
	{
	}
}

func TestEGenericTypeELowerBoundSet(t *testing.T) {
	obj := newEGenericTypeImpl()
	mockAdapter := &MockEAdapter{}
	mockAdapter.On("SetTarget", obj).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	obj.EAdapters().Add(mockAdapter)
	obj.SetELowerBound(newEGenericTypeImpl())
	mockAdapter.AssertExpectations(t)
}

func TestEGenericTypeELowerBoundEGet(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		assert.Equal(t, obj.GetEClassifier(), obj.EGetFromID(EGENERIC_TYPE__ECLASSIFIER, false, false))
	}
	{
		assert.Equal(t, obj.GetELowerBound(), obj.EGetFromID(EGENERIC_TYPE__ELOWER_BOUND, false, false))
	}
	{
		assert.Equal(t, obj.GetERawType(), obj.EGetFromID(EGENERIC_TYPE__ERAW_TYPE, false, false))
	}
	{
		assert.Equal(t, obj.GetETypeArguments(), obj.EGetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS, false, false))
	}
	{
		assert.Equal(t, obj.GetETypeParameter(), obj.EGetFromID(EGENERIC_TYPE__ETYPE_PARAMETER, false, false))
	}
	{
		assert.Equal(t, obj.GetEUpperBound(), obj.EGetFromID(EGENERIC_TYPE__EUPPER_BOUND, false, false))
	}
}

func TestEGenericTypeELowerBoundEInvoke(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		assert.Panics(t, func() { obj.EInvokeFromID(EGENERIC_TYPE__IS_INSTANCE_EJAVAOBJECT, nil) })
	}
}

func TestEGenericTypeELowerBoundEIsSet(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		assert.Equal(t, obj.eTypeArguments != nil && obj.eTypeArguments.Size() != 0, obj.EIsSetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEGenericTypeELowerBoundEUnset(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		obj.EUnsetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS)
		obj.GetETypeArguments().Clear()
		assert.Equal(t, 0, obj.GetETypeArguments().Size())
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEGenericTypeELowerBoundESet(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		elem := NewEmptyBasicEList()
		obj.ESetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS, elem)
		assert.Equal(t, 1, obj.GetETypeArguments().Size())
		assert.Equal(t, elem, obj.GetETypeArguments().Get(0))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEGenericTypeELowerBoundEInverseRemove(t *testing.T) {
	{
	}
	{
		obj := newEGenericTypeImpl()
		mock := &MockEObject{}
		obj.GetETypeArguments().Add(mock)
		obj.EBasicInverseRemove(mock, EGENERIC_TYPE__ETYPE_ARGUMENTS, &MockENotificationChain{})
		assert.Equal(t, 0, obj.GetETypeArguments().Size())
	}
	{
	}
}

func TestEGenericTypeETypeParameterGet(t *testing.T) {
	obj := newEGenericTypeImpl()
	obj.SetETypeParameter(newETypeParameterImpl())
	assert.Equal(t, newETypeParameterImpl(), obj.GetETypeParameter())
}

func TestEGenericTypeETypeParameterSet(t *testing.T) {
	obj := newEGenericTypeImpl()
	mockAdapter := &MockEAdapter{}
	mockAdapter.On("SetTarget", obj).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	obj.EAdapters().Add(mockAdapter)
	obj.SetETypeParameter(newETypeParameterImpl())
	mockAdapter.AssertExpectations(t)
}

func TestEGenericTypeETypeParameterEGet(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		assert.Equal(t, obj.GetEClassifier(), obj.EGetFromID(EGENERIC_TYPE__ECLASSIFIER, false, false))
	}
	{
		assert.Equal(t, obj.GetELowerBound(), obj.EGetFromID(EGENERIC_TYPE__ELOWER_BOUND, false, false))
	}
	{
		assert.Equal(t, obj.GetERawType(), obj.EGetFromID(EGENERIC_TYPE__ERAW_TYPE, false, false))
	}
	{
		assert.Equal(t, obj.GetETypeArguments(), obj.EGetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS, false, false))
	}
	{
		assert.Equal(t, obj.GetETypeParameter(), obj.EGetFromID(EGENERIC_TYPE__ETYPE_PARAMETER, false, false))
	}
	{
		assert.Equal(t, obj.GetEUpperBound(), obj.EGetFromID(EGENERIC_TYPE__EUPPER_BOUND, false, false))
	}
}

func TestEGenericTypeETypeParameterEInvoke(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		assert.Panics(t, func() { obj.EInvokeFromID(EGENERIC_TYPE__IS_INSTANCE_EJAVAOBJECT, nil) })
	}
}

func TestEGenericTypeETypeParameterEIsSet(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		assert.Equal(t, obj.eTypeArguments != nil && obj.eTypeArguments.Size() != 0, obj.EIsSetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEGenericTypeETypeParameterEUnset(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		obj.EUnsetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS)
		obj.GetETypeArguments().Clear()
		assert.Equal(t, 0, obj.GetETypeArguments().Size())
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEGenericTypeETypeParameterESet(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		elem := NewEmptyBasicEList()
		obj.ESetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS, elem)
		assert.Equal(t, 1, obj.GetETypeArguments().Size())
		assert.Equal(t, elem, obj.GetETypeArguments().Get(0))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEGenericTypeETypeParameterEInverseRemove(t *testing.T) {
	{
	}
	{
		obj := newEGenericTypeImpl()
		mock := &MockEObject{}
		obj.GetETypeArguments().Add(mock)
		obj.EBasicInverseRemove(mock, EGENERIC_TYPE__ETYPE_ARGUMENTS, &MockENotificationChain{})
		assert.Equal(t, 0, obj.GetETypeArguments().Size())
	}
	{
	}
}

func TestEGenericTypeEClassifierGet(t *testing.T) {
	obj := newEGenericTypeImpl()
	obj.SetEClassifier(newEClassifierImpl())
	assert.Equal(t, newEClassifierImpl(), obj.GetEClassifier())
}

func TestEGenericTypeEClassifierSet(t *testing.T) {
	obj := newEGenericTypeImpl()
	mockAdapter := &MockEAdapter{}
	mockAdapter.On("SetTarget", obj).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	obj.EAdapters().Add(mockAdapter)
	obj.SetEClassifier(newEClassifierImpl())
	mockAdapter.AssertExpectations(t)
}

func TestEGenericTypeEClassifierEGet(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		assert.Equal(t, obj.GetEClassifier(), obj.EGetFromID(EGENERIC_TYPE__ECLASSIFIER, false, false))
	}
	{
		assert.Equal(t, obj.GetELowerBound(), obj.EGetFromID(EGENERIC_TYPE__ELOWER_BOUND, false, false))
	}
	{
		assert.Equal(t, obj.GetERawType(), obj.EGetFromID(EGENERIC_TYPE__ERAW_TYPE, false, false))
	}
	{
		assert.Equal(t, obj.GetETypeArguments(), obj.EGetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS, false, false))
	}
	{
		assert.Equal(t, obj.GetETypeParameter(), obj.EGetFromID(EGENERIC_TYPE__ETYPE_PARAMETER, false, false))
	}
	{
		assert.Equal(t, obj.GetEUpperBound(), obj.EGetFromID(EGENERIC_TYPE__EUPPER_BOUND, false, false))
	}
}

func TestEGenericTypeEClassifierEInvoke(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		assert.Panics(t, func() { obj.EInvokeFromID(EGENERIC_TYPE__IS_INSTANCE_EJAVAOBJECT, nil) })
	}
}

func TestEGenericTypeEClassifierEIsSet(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		assert.Equal(t, obj.eTypeArguments != nil && obj.eTypeArguments.Size() != 0, obj.EIsSetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEGenericTypeEClassifierEUnset(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		obj.EUnsetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS)
		obj.GetETypeArguments().Clear()
		assert.Equal(t, 0, obj.GetETypeArguments().Size())
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEGenericTypeEClassifierESet(t *testing.T) {
	obj := newEGenericTypeImpl()
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		elem := NewEmptyBasicEList()
		obj.ESetFromID(EGENERIC_TYPE__ETYPE_ARGUMENTS, elem)
		assert.Equal(t, 1, obj.GetETypeArguments().Size())
		assert.Equal(t, elem, obj.GetETypeArguments().Get(0))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEGenericTypeEClassifierEInverseRemove(t *testing.T) {
	{
	}
	{
		obj := newEGenericTypeImpl()
		mock := &MockEObject{}
		obj.GetETypeArguments().Add(mock)
		obj.EBasicInverseRemove(mock, EGENERIC_TYPE__ETYPE_ARGUMENTS, &MockENotificationChain{})
		assert.Equal(t, 0, obj.GetETypeArguments().Size())
	}
	{
	}
}

func TestEGenericTypeIsInstanceOperation(t *testing.T) {
	obj := newEGenericTypeImpl()
	assert.Panics(t, func() { obj.IsInstance(nil) })
}
