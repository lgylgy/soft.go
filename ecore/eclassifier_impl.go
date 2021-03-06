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

import "reflect"

// eClassifierImpl is the implementation of the model object 'EClassifier'
type eClassifierImpl struct {
	*eNamedElementImpl
	classifierID  int
	defaultValue  interface{}
	ePackage      EPackage
	instanceClass reflect.Type
}

// newEClassifierImpl is the constructor of a eClassifierImpl
func newEClassifierImpl() *eClassifierImpl {
	eClassifier := new(eClassifierImpl)
	eClassifier.eNamedElementImpl = newENamedElementImpl()
	eClassifier.SetInterfaces(eClassifier)
	eClassifier.classifierID = -1
	eClassifier.instanceClass = nil

	return eClassifier
}

func (eClassifier *eClassifierImpl) EStaticClass() EClass {
	return GetPackage().GetEClassifierClass()
}

// IsInstance default implementation
func (eClassifier *eClassifierImpl) IsInstance(interface{}) bool {
	panic("IsInstance not implemented")
}

// GetClassifierID get the value of classifierID
func (eClassifier *eClassifierImpl) GetClassifierID() int {
	return eClassifier.classifierID
}

// SetClassifierID set the value of classifierID
func (eClassifier *eClassifierImpl) SetClassifierID(newClassifierID int) {
	oldClassifierID := eClassifier.classifierID
	eClassifier.classifierID = newClassifierID
	if eClassifier.ENotificationRequired() {
		eClassifier.ENotify(NewNotificationByFeatureID(eClassifier.AsEObject(), SET, ECLASSIFIER__CLASSIFIER_ID, oldClassifierID, newClassifierID, NO_INDEX))
	}
}

// GetDefaultValue get the value of defaultValue
func (eClassifier *eClassifierImpl) GetDefaultValue() interface{} {
	panic("GetDefaultValue not implemented")
}

// GetEPackage get the value of ePackage
func (eClassifier *eClassifierImpl) GetEPackage() EPackage {
	if eClassifier.EContainerFeatureID() == ECLASSIFIER__EPACKAGE {
		return eClassifier.EContainer().(EPackage)
	}
	return nil
}

// GetInstanceClass get the value of instanceClass
func (eClassifier *eClassifierImpl) GetInstanceClass() reflect.Type {
	return eClassifier.instanceClass
}

// SetInstanceClass set the value of instanceClass
func (eClassifier *eClassifierImpl) SetInstanceClass(newInstanceClass reflect.Type) {
	oldInstanceClass := eClassifier.instanceClass
	eClassifier.instanceClass = newInstanceClass
	if eClassifier.ENotificationRequired() {
		eClassifier.ENotify(NewNotificationByFeatureID(eClassifier.AsEObject(), SET, ECLASSIFIER__INSTANCE_CLASS, oldInstanceClass, newInstanceClass, NO_INDEX))
	}
}

func (eClassifier *eClassifierImpl) EGetFromID(featureID int, resolve, coreType bool) interface{} {
	switch featureID {
	case ECLASSIFIER__CLASSIFIER_ID:
		return eClassifier.GetClassifierID()
	case ECLASSIFIER__DEFAULT_VALUE:
		return eClassifier.GetDefaultValue()
	case ECLASSIFIER__EPACKAGE:
		return eClassifier.GetEPackage()
	case ECLASSIFIER__INSTANCE_CLASS:
		return eClassifier.GetInstanceClass()
	default:
		return eClassifier.eNamedElementImpl.EGetFromID(featureID, resolve, coreType)
	}
}

func (eClassifier *eClassifierImpl) ESetFromID(featureID int, newValue interface{}) {
	switch featureID {
	case ECLASSIFIER__CLASSIFIER_ID:
		c := newValue.(int)
		eClassifier.SetClassifierID(c)
	case ECLASSIFIER__INSTANCE_CLASS:
		i := newValue.(reflect.Type)
		eClassifier.SetInstanceClass(i)
	default:
		eClassifier.eNamedElementImpl.ESetFromID(featureID, newValue)
	}
}

func (eClassifier *eClassifierImpl) EUnsetFromID(featureID int) {
	switch featureID {
	case ECLASSIFIER__CLASSIFIER_ID:
		eClassifier.SetClassifierID(-1)
	case ECLASSIFIER__INSTANCE_CLASS:
		eClassifier.SetInstanceClass(nil)
	default:
		eClassifier.eNamedElementImpl.EUnsetFromID(featureID)
	}
}

func (eClassifier *eClassifierImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case ECLASSIFIER__CLASSIFIER_ID:
		return eClassifier.classifierID != -1
	case ECLASSIFIER__DEFAULT_VALUE:
		return eClassifier.GetDefaultValue() != ""
	case ECLASSIFIER__EPACKAGE:
		return eClassifier.GetEPackage() != nil
	case ECLASSIFIER__INSTANCE_CLASS:
		return eClassifier.instanceClass != nil
	default:
		return eClassifier.eNamedElementImpl.EIsSetFromID(featureID)
	}
}

func (eClassifier *eClassifierImpl) EInvokeFromID(operationID int, arguments EList) interface{} {
	switch operationID {
	case ECLASSIFIER__IS_INSTANCE_EJAVAOBJECT:
		return eClassifier.IsInstance(arguments.Get(0))
	default:
		return eClassifier.eNamedElementImpl.EInvokeFromID(operationID, arguments)
	}
}

func (eClassifier *eClassifierImpl) EBasicInverseAdd(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case ECLASSIFIER__EPACKAGE:
		msgs := notifications
		if eClassifier.EContainer() != nil {
			msgs = eClassifier.EBasicRemoveFromContainer(msgs)
		}
		return eClassifier.EBasicSetContainer(otherEnd, ECLASSIFIER__EPACKAGE, msgs)
	default:
		return eClassifier.eNamedElementImpl.EBasicInverseAdd(otherEnd, featureID, notifications)
	}
}

func (eClassifier *eClassifierImpl) EBasicInverseRemove(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case ECLASSIFIER__EPACKAGE:
		return eClassifier.EBasicSetContainer(nil, ECLASSIFIER__EPACKAGE, notifications)
	default:
		return eClassifier.eNamedElementImpl.EBasicInverseRemove(otherEnd, featureID, notifications)
	}
}
