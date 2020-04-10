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

func discardEPackage() {
	_ = assert.Equal
	_ = mock.Anything
	_ = testing.Coverage

	_ = time.Now()
}

func TestEPackageNsURIGet(t *testing.T) {
	obj := newEPackageImpl()
	obj.SetNsURI("Test String")
	assert.Equal(t, "Test String", obj.GetNsURI())
}

func TestEPackageNsURISet(t *testing.T) {
	obj := newEPackageImpl()
	mockAdapter := &MockEAdapter{}
	mockAdapter.On("SetTarget", obj).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	obj.EAdapters().Add(mockAdapter)
	obj.SetNsURI("Test String")
	mockAdapter.AssertExpectations(t)
}

func TestEPackageNsURIEGet(t *testing.T) {
	obj := newEPackageImpl()
	{
		assert.Equal(t, obj.GetEClassifiers(), obj.EGetFromID(EPACKAGE__ECLASSIFIERS, false, false))
	}
	{
		assert.Equal(t, obj.GetEFactoryInstance(), obj.EGetFromID(EPACKAGE__EFACTORY_INSTANCE, false, false))
	}
	{
		assert.Equal(t, obj.GetESubPackages(), obj.EGetFromID(EPACKAGE__ESUB_PACKAGES, false, false))
	}
	{
		assert.Equal(t, obj.GetESuperPackage(), obj.EGetFromID(EPACKAGE__ESUPER_PACKAGE, false, false))
	}
	{
		assert.Equal(t, obj.GetNsPrefix(), obj.EGetFromID(EPACKAGE__NS_PREFIX, false, false))
	}
	{
		assert.Equal(t, obj.GetNsURI(), obj.EGetFromID(EPACKAGE__NS_URI, false, false))
	}
}

func TestEPackageNsURIEInvoke(t *testing.T) {
	obj := newEPackageImpl()
	{
		assert.Panics(t, func() { obj.EInvokeFromID(EPACKAGE__GET_ECLASSIFIER_ESTRING, nil) })
	}
}

func TestEPackageNsURIEIsSet(t *testing.T) {
	obj := newEPackageImpl()
	{
		assert.Equal(t, obj.eClassifiers != nil && obj.eClassifiers.Size() != 0, obj.EIsSetFromID(EPACKAGE__ECLASSIFIERS))
	}
	{
		_ = obj
	}
	{
		assert.Equal(t, obj.eSubPackages != nil && obj.eSubPackages.Size() != 0, obj.EIsSetFromID(EPACKAGE__ESUB_PACKAGES))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEPackageNsURIEUnset(t *testing.T) {
	obj := newEPackageImpl()
	{
		obj.EUnsetFromID(EPACKAGE__ECLASSIFIERS)
		obj.GetEClassifiers().Clear()
		assert.Equal(t, 0, obj.GetEClassifiers().Size())
	}
	{
		_ = obj
	}
	{
		obj.EUnsetFromID(EPACKAGE__ESUB_PACKAGES)
		obj.GetESubPackages().Clear()
		assert.Equal(t, 0, obj.GetESubPackages().Size())
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEPackageNsURIESet(t *testing.T) {
	obj := newEPackageImpl()
	{
		elem := NewEmptyBasicEList()
		obj.ESetFromID(EPACKAGE__ECLASSIFIERS, elem)
		assert.Equal(t, 1, obj.GetEClassifiers().Size())
		assert.Equal(t, elem, obj.GetEClassifiers().Get(0))
	}
	{
		_ = obj
	}
	{
		elem := NewEmptyBasicEList()
		obj.ESetFromID(EPACKAGE__ESUB_PACKAGES, elem)
		assert.Equal(t, 1, obj.GetESubPackages().Size())
		assert.Equal(t, elem, obj.GetESubPackages().Get(0))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEPackageNsURIEInverseAdd(t *testing.T) {
	{
		obj := newEPackageImpl()
		obj.EBasicInverseAdd(&MockEObject{}, EPACKAGE__ECLASSIFIERS, &MockENotificationChain{})
		assert.Equal(t, 1, obj.GetEClassifiers().Size())
	}
	{
	}
	{
		obj := newEPackageImpl()
		obj.EBasicInverseAdd(&MockEObject{}, EPACKAGE__ESUB_PACKAGES, &MockENotificationChain{})
		assert.Equal(t, 1, obj.GetESubPackages().Size())
	}
	{
	}
}

func TestEPackageNsURIEInverseRemove(t *testing.T) {
	{
		obj := newEPackageImpl()
		mock := &MockEObject{}
		obj.GetEClassifiers().Add(mock)
		obj.EBasicInverseRemove(mock, EPACKAGE__ECLASSIFIERS, &MockENotificationChain{})
		assert.Equal(t, 0, obj.GetEClassifiers().Size())
	}
	{
	}
	{
		obj := newEPackageImpl()
		mock := &MockEObject{}
		obj.GetESubPackages().Add(mock)
		obj.EBasicInverseRemove(mock, EPACKAGE__ESUB_PACKAGES, &MockENotificationChain{})
		assert.Equal(t, 0, obj.GetESubPackages().Size())
	}
	{
	}
}

func TestEPackageNsPrefixGet(t *testing.T) {
	obj := newEPackageImpl()
	obj.SetNsPrefix("Test String")
	assert.Equal(t, "Test String", obj.GetNsPrefix())
}

func TestEPackageNsPrefixSet(t *testing.T) {
	obj := newEPackageImpl()
	mockAdapter := &MockEAdapter{}
	mockAdapter.On("SetTarget", obj).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	obj.EAdapters().Add(mockAdapter)
	obj.SetNsPrefix("Test String")
	mockAdapter.AssertExpectations(t)
}

func TestEPackageNsPrefixEGet(t *testing.T) {
	obj := newEPackageImpl()
	{
		assert.Equal(t, obj.GetEClassifiers(), obj.EGetFromID(EPACKAGE__ECLASSIFIERS, false, false))
	}
	{
		assert.Equal(t, obj.GetEFactoryInstance(), obj.EGetFromID(EPACKAGE__EFACTORY_INSTANCE, false, false))
	}
	{
		assert.Equal(t, obj.GetESubPackages(), obj.EGetFromID(EPACKAGE__ESUB_PACKAGES, false, false))
	}
	{
		assert.Equal(t, obj.GetESuperPackage(), obj.EGetFromID(EPACKAGE__ESUPER_PACKAGE, false, false))
	}
	{
		assert.Equal(t, obj.GetNsPrefix(), obj.EGetFromID(EPACKAGE__NS_PREFIX, false, false))
	}
	{
		assert.Equal(t, obj.GetNsURI(), obj.EGetFromID(EPACKAGE__NS_URI, false, false))
	}
}

func TestEPackageNsPrefixEInvoke(t *testing.T) {
	obj := newEPackageImpl()
	{
		assert.Panics(t, func() { obj.EInvokeFromID(EPACKAGE__GET_ECLASSIFIER_ESTRING, nil) })
	}
}

func TestEPackageNsPrefixEIsSet(t *testing.T) {
	obj := newEPackageImpl()
	{
		assert.Equal(t, obj.eClassifiers != nil && obj.eClassifiers.Size() != 0, obj.EIsSetFromID(EPACKAGE__ECLASSIFIERS))
	}
	{
		_ = obj
	}
	{
		assert.Equal(t, obj.eSubPackages != nil && obj.eSubPackages.Size() != 0, obj.EIsSetFromID(EPACKAGE__ESUB_PACKAGES))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEPackageNsPrefixEUnset(t *testing.T) {
	obj := newEPackageImpl()
	{
		obj.EUnsetFromID(EPACKAGE__ECLASSIFIERS)
		obj.GetEClassifiers().Clear()
		assert.Equal(t, 0, obj.GetEClassifiers().Size())
	}
	{
		_ = obj
	}
	{
		obj.EUnsetFromID(EPACKAGE__ESUB_PACKAGES)
		obj.GetESubPackages().Clear()
		assert.Equal(t, 0, obj.GetESubPackages().Size())
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEPackageNsPrefixESet(t *testing.T) {
	obj := newEPackageImpl()
	{
		elem := NewEmptyBasicEList()
		obj.ESetFromID(EPACKAGE__ECLASSIFIERS, elem)
		assert.Equal(t, 1, obj.GetEClassifiers().Size())
		assert.Equal(t, elem, obj.GetEClassifiers().Get(0))
	}
	{
		_ = obj
	}
	{
		elem := NewEmptyBasicEList()
		obj.ESetFromID(EPACKAGE__ESUB_PACKAGES, elem)
		assert.Equal(t, 1, obj.GetESubPackages().Size())
		assert.Equal(t, elem, obj.GetESubPackages().Get(0))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEPackageNsPrefixEInverseAdd(t *testing.T) {
	{
		obj := newEPackageImpl()
		obj.EBasicInverseAdd(&MockEObject{}, EPACKAGE__ECLASSIFIERS, &MockENotificationChain{})
		assert.Equal(t, 1, obj.GetEClassifiers().Size())
	}
	{
	}
	{
		obj := newEPackageImpl()
		obj.EBasicInverseAdd(&MockEObject{}, EPACKAGE__ESUB_PACKAGES, &MockENotificationChain{})
		assert.Equal(t, 1, obj.GetESubPackages().Size())
	}
	{
	}
}

func TestEPackageNsPrefixEInverseRemove(t *testing.T) {
	{
		obj := newEPackageImpl()
		mock := &MockEObject{}
		obj.GetEClassifiers().Add(mock)
		obj.EBasicInverseRemove(mock, EPACKAGE__ECLASSIFIERS, &MockENotificationChain{})
		assert.Equal(t, 0, obj.GetEClassifiers().Size())
	}
	{
	}
	{
		obj := newEPackageImpl()
		mock := &MockEObject{}
		obj.GetESubPackages().Add(mock)
		obj.EBasicInverseRemove(mock, EPACKAGE__ESUB_PACKAGES, &MockENotificationChain{})
		assert.Equal(t, 0, obj.GetESubPackages().Size())
	}
	{
	}
}

func TestEPackageEFactoryInstanceSet(t *testing.T) {
	obj := newEPackageImpl()
	mockAdapter := &MockEAdapter{}
	mockAdapter.On("SetTarget", obj).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	obj.EAdapters().Add(mockAdapter)
	obj.SetEFactoryInstance(newEFactoryImpl())
	mockAdapter.AssertExpectations(t)
}

func TestEPackageEFactoryInstanceEGet(t *testing.T) {
	obj := newEPackageImpl()
	{
		assert.Equal(t, obj.GetEClassifiers(), obj.EGetFromID(EPACKAGE__ECLASSIFIERS, false, false))
	}
	{
		assert.Equal(t, obj.GetEFactoryInstance(), obj.EGetFromID(EPACKAGE__EFACTORY_INSTANCE, false, false))
	}
	{
		assert.Equal(t, obj.GetESubPackages(), obj.EGetFromID(EPACKAGE__ESUB_PACKAGES, false, false))
	}
	{
		assert.Equal(t, obj.GetESuperPackage(), obj.EGetFromID(EPACKAGE__ESUPER_PACKAGE, false, false))
	}
	{
		assert.Equal(t, obj.GetNsPrefix(), obj.EGetFromID(EPACKAGE__NS_PREFIX, false, false))
	}
	{
		assert.Equal(t, obj.GetNsURI(), obj.EGetFromID(EPACKAGE__NS_URI, false, false))
	}
}

func TestEPackageEFactoryInstanceEInvoke(t *testing.T) {
	obj := newEPackageImpl()
	{
		assert.Panics(t, func() { obj.EInvokeFromID(EPACKAGE__GET_ECLASSIFIER_ESTRING, nil) })
	}
}

func TestEPackageEFactoryInstanceEIsSet(t *testing.T) {
	obj := newEPackageImpl()
	{
		assert.Equal(t, obj.eClassifiers != nil && obj.eClassifiers.Size() != 0, obj.EIsSetFromID(EPACKAGE__ECLASSIFIERS))
	}
	{
		_ = obj
	}
	{
		assert.Equal(t, obj.eSubPackages != nil && obj.eSubPackages.Size() != 0, obj.EIsSetFromID(EPACKAGE__ESUB_PACKAGES))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEPackageEFactoryInstanceEUnset(t *testing.T) {
	obj := newEPackageImpl()
	{
		obj.EUnsetFromID(EPACKAGE__ECLASSIFIERS)
		obj.GetEClassifiers().Clear()
		assert.Equal(t, 0, obj.GetEClassifiers().Size())
	}
	{
		_ = obj
	}
	{
		obj.EUnsetFromID(EPACKAGE__ESUB_PACKAGES)
		obj.GetESubPackages().Clear()
		assert.Equal(t, 0, obj.GetESubPackages().Size())
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEPackageEFactoryInstanceESet(t *testing.T) {
	obj := newEPackageImpl()
	{
		elem := NewEmptyBasicEList()
		obj.ESetFromID(EPACKAGE__ECLASSIFIERS, elem)
		assert.Equal(t, 1, obj.GetEClassifiers().Size())
		assert.Equal(t, elem, obj.GetEClassifiers().Get(0))
	}
	{
		_ = obj
	}
	{
		elem := NewEmptyBasicEList()
		obj.ESetFromID(EPACKAGE__ESUB_PACKAGES, elem)
		assert.Equal(t, 1, obj.GetESubPackages().Size())
		assert.Equal(t, elem, obj.GetESubPackages().Get(0))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEPackageEFactoryInstanceEInverseAdd(t *testing.T) {
	{
		obj := newEPackageImpl()
		obj.EBasicInverseAdd(&MockEObject{}, EPACKAGE__ECLASSIFIERS, &MockENotificationChain{})
		assert.Equal(t, 1, obj.GetEClassifiers().Size())
	}
	{
	}
	{
		obj := newEPackageImpl()
		obj.EBasicInverseAdd(&MockEObject{}, EPACKAGE__ESUB_PACKAGES, &MockENotificationChain{})
		assert.Equal(t, 1, obj.GetESubPackages().Size())
	}
	{
	}
}

func TestEPackageEFactoryInstanceEInverseRemove(t *testing.T) {
	{
		obj := newEPackageImpl()
		mock := &MockEObject{}
		obj.GetEClassifiers().Add(mock)
		obj.EBasicInverseRemove(mock, EPACKAGE__ECLASSIFIERS, &MockENotificationChain{})
		assert.Equal(t, 0, obj.GetEClassifiers().Size())
	}
	{
	}
	{
		obj := newEPackageImpl()
		mock := &MockEObject{}
		obj.GetESubPackages().Add(mock)
		obj.EBasicInverseRemove(mock, EPACKAGE__ESUB_PACKAGES, &MockENotificationChain{})
		assert.Equal(t, 0, obj.GetESubPackages().Size())
	}
	{
	}
}

func TestEPackageEClassifiersGetList(t *testing.T) {
	obj := newEPackageImpl()
	assert.NotNil(t, obj.GetEClassifiers())
}

func TestEPackageEClassifiersEGet(t *testing.T) {
	obj := newEPackageImpl()
	{
		assert.Equal(t, obj.GetEClassifiers(), obj.EGetFromID(EPACKAGE__ECLASSIFIERS, false, false))
	}
	{
		assert.Equal(t, obj.GetEFactoryInstance(), obj.EGetFromID(EPACKAGE__EFACTORY_INSTANCE, false, false))
	}
	{
		assert.Equal(t, obj.GetESubPackages(), obj.EGetFromID(EPACKAGE__ESUB_PACKAGES, false, false))
	}
	{
		assert.Equal(t, obj.GetESuperPackage(), obj.EGetFromID(EPACKAGE__ESUPER_PACKAGE, false, false))
	}
	{
		assert.Equal(t, obj.GetNsPrefix(), obj.EGetFromID(EPACKAGE__NS_PREFIX, false, false))
	}
	{
		assert.Equal(t, obj.GetNsURI(), obj.EGetFromID(EPACKAGE__NS_URI, false, false))
	}
}

func TestEPackageEClassifiersEInvoke(t *testing.T) {
	obj := newEPackageImpl()
	{
		assert.Panics(t, func() { obj.EInvokeFromID(EPACKAGE__GET_ECLASSIFIER_ESTRING, nil) })
	}
}

func TestEPackageEClassifiersEIsSet(t *testing.T) {
	obj := newEPackageImpl()
	{
		assert.Equal(t, obj.eClassifiers != nil && obj.eClassifiers.Size() != 0, obj.EIsSetFromID(EPACKAGE__ECLASSIFIERS))
	}
	{
		_ = obj
	}
	{
		assert.Equal(t, obj.eSubPackages != nil && obj.eSubPackages.Size() != 0, obj.EIsSetFromID(EPACKAGE__ESUB_PACKAGES))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEPackageEClassifiersEUnset(t *testing.T) {
	obj := newEPackageImpl()
	{
		obj.EUnsetFromID(EPACKAGE__ECLASSIFIERS)
		obj.GetEClassifiers().Clear()
		assert.Equal(t, 0, obj.GetEClassifiers().Size())
	}
	{
		_ = obj
	}
	{
		obj.EUnsetFromID(EPACKAGE__ESUB_PACKAGES)
		obj.GetESubPackages().Clear()
		assert.Equal(t, 0, obj.GetESubPackages().Size())
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEPackageEClassifiersESet(t *testing.T) {
	obj := newEPackageImpl()
	{
		elem := NewEmptyBasicEList()
		obj.ESetFromID(EPACKAGE__ECLASSIFIERS, elem)
		assert.Equal(t, 1, obj.GetEClassifiers().Size())
		assert.Equal(t, elem, obj.GetEClassifiers().Get(0))
	}
	{
		_ = obj
	}
	{
		elem := NewEmptyBasicEList()
		obj.ESetFromID(EPACKAGE__ESUB_PACKAGES, elem)
		assert.Equal(t, 1, obj.GetESubPackages().Size())
		assert.Equal(t, elem, obj.GetESubPackages().Get(0))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEPackageEClassifiersEInverseAdd(t *testing.T) {
	{
		obj := newEPackageImpl()
		obj.EBasicInverseAdd(&MockEObject{}, EPACKAGE__ECLASSIFIERS, &MockENotificationChain{})
		assert.Equal(t, 1, obj.GetEClassifiers().Size())
	}
	{
	}
	{
		obj := newEPackageImpl()
		obj.EBasicInverseAdd(&MockEObject{}, EPACKAGE__ESUB_PACKAGES, &MockENotificationChain{})
		assert.Equal(t, 1, obj.GetESubPackages().Size())
	}
	{
	}
}

func TestEPackageEClassifiersEInverseRemove(t *testing.T) {
	{
		obj := newEPackageImpl()
		mock := &MockEObject{}
		obj.GetEClassifiers().Add(mock)
		obj.EBasicInverseRemove(mock, EPACKAGE__ECLASSIFIERS, &MockENotificationChain{})
		assert.Equal(t, 0, obj.GetEClassifiers().Size())
	}
	{
	}
	{
		obj := newEPackageImpl()
		mock := &MockEObject{}
		obj.GetESubPackages().Add(mock)
		obj.EBasicInverseRemove(mock, EPACKAGE__ESUB_PACKAGES, &MockENotificationChain{})
		assert.Equal(t, 0, obj.GetESubPackages().Size())
	}
	{
	}
}

func TestEPackageESubPackagesGetList(t *testing.T) {
	obj := newEPackageImpl()
	assert.NotNil(t, obj.GetESubPackages())
}

func TestEPackageESubPackagesEGet(t *testing.T) {
	obj := newEPackageImpl()
	{
		assert.Equal(t, obj.GetEClassifiers(), obj.EGetFromID(EPACKAGE__ECLASSIFIERS, false, false))
	}
	{
		assert.Equal(t, obj.GetEFactoryInstance(), obj.EGetFromID(EPACKAGE__EFACTORY_INSTANCE, false, false))
	}
	{
		assert.Equal(t, obj.GetESubPackages(), obj.EGetFromID(EPACKAGE__ESUB_PACKAGES, false, false))
	}
	{
		assert.Equal(t, obj.GetESuperPackage(), obj.EGetFromID(EPACKAGE__ESUPER_PACKAGE, false, false))
	}
	{
		assert.Equal(t, obj.GetNsPrefix(), obj.EGetFromID(EPACKAGE__NS_PREFIX, false, false))
	}
	{
		assert.Equal(t, obj.GetNsURI(), obj.EGetFromID(EPACKAGE__NS_URI, false, false))
	}
}

func TestEPackageESubPackagesEInvoke(t *testing.T) {
	obj := newEPackageImpl()
	{
		assert.Panics(t, func() { obj.EInvokeFromID(EPACKAGE__GET_ECLASSIFIER_ESTRING, nil) })
	}
}

func TestEPackageESubPackagesEIsSet(t *testing.T) {
	obj := newEPackageImpl()
	{
		assert.Equal(t, obj.eClassifiers != nil && obj.eClassifiers.Size() != 0, obj.EIsSetFromID(EPACKAGE__ECLASSIFIERS))
	}
	{
		_ = obj
	}
	{
		assert.Equal(t, obj.eSubPackages != nil && obj.eSubPackages.Size() != 0, obj.EIsSetFromID(EPACKAGE__ESUB_PACKAGES))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEPackageESubPackagesEUnset(t *testing.T) {
	obj := newEPackageImpl()
	{
		obj.EUnsetFromID(EPACKAGE__ECLASSIFIERS)
		obj.GetEClassifiers().Clear()
		assert.Equal(t, 0, obj.GetEClassifiers().Size())
	}
	{
		_ = obj
	}
	{
		obj.EUnsetFromID(EPACKAGE__ESUB_PACKAGES)
		obj.GetESubPackages().Clear()
		assert.Equal(t, 0, obj.GetESubPackages().Size())
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEPackageESubPackagesESet(t *testing.T) {
	obj := newEPackageImpl()
	{
		elem := NewEmptyBasicEList()
		obj.ESetFromID(EPACKAGE__ECLASSIFIERS, elem)
		assert.Equal(t, 1, obj.GetEClassifiers().Size())
		assert.Equal(t, elem, obj.GetEClassifiers().Get(0))
	}
	{
		_ = obj
	}
	{
		elem := NewEmptyBasicEList()
		obj.ESetFromID(EPACKAGE__ESUB_PACKAGES, elem)
		assert.Equal(t, 1, obj.GetESubPackages().Size())
		assert.Equal(t, elem, obj.GetESubPackages().Get(0))
	}
	{
		_ = obj
	}
	{
		_ = obj
	}
}

func TestEPackageESubPackagesEInverseAdd(t *testing.T) {
	{
		obj := newEPackageImpl()
		obj.EBasicInverseAdd(&MockEObject{}, EPACKAGE__ECLASSIFIERS, &MockENotificationChain{})
		assert.Equal(t, 1, obj.GetEClassifiers().Size())
	}
	{
	}
	{
		obj := newEPackageImpl()
		obj.EBasicInverseAdd(&MockEObject{}, EPACKAGE__ESUB_PACKAGES, &MockENotificationChain{})
		assert.Equal(t, 1, obj.GetESubPackages().Size())
	}
	{
	}
}

func TestEPackageESubPackagesEInverseRemove(t *testing.T) {
	{
		obj := newEPackageImpl()
		mock := &MockEObject{}
		obj.GetEClassifiers().Add(mock)
		obj.EBasicInverseRemove(mock, EPACKAGE__ECLASSIFIERS, &MockENotificationChain{})
		assert.Equal(t, 0, obj.GetEClassifiers().Size())
	}
	{
	}
	{
		obj := newEPackageImpl()
		mock := &MockEObject{}
		obj.GetESubPackages().Add(mock)
		obj.EBasicInverseRemove(mock, EPACKAGE__ESUB_PACKAGES, &MockENotificationChain{})
		assert.Equal(t, 0, obj.GetESubPackages().Size())
	}
	{
	}
}

func TestEPackageGetEClassifierOperation(t *testing.T) {
	obj := newEPackageImpl()
	assert.Panics(t, func() { obj.GetEClassifier("Test String") })
}
