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

// eOperationImpl is the implementation of the model object 'EOperation'
type eOperationImpl struct {
	*eTypedElementExt
	eContainingClass EClass
	eExceptions      EList
	eParameters      EList
	operationID      int
}

// newEOperationImpl is the constructor of a eOperationImpl
func newEOperationImpl() *eOperationImpl {
	eOperation := new(eOperationImpl)
	eOperation.eTypedElementExt = newETypedElementExt()
	eOperation.SetInterfaces(eOperation)
	eOperation.operationID = -1

	return eOperation
}

type eOperationImplInitializers interface {
	initEExceptions() EList
	initEParameters() EList
}

func (eOperation *eOperationImpl) getInitializers() eOperationImplInitializers {
	return eOperation.AsEObject().(eOperationImplInitializers)
}

func (eOperation *eOperationImpl) EStaticClass() EClass {
	return GetPackage().GetEOperation()
}

// IsOverrideOf default implementation
func (eOperation *eOperationImpl) IsOverrideOf(EOperation) bool {
	panic("IsOverrideOf not implemented")
}

// GetEContainingClass get the value of eContainingClass
func (eOperation *eOperationImpl) GetEContainingClass() EClass {
	if eOperation.EContainerFeatureID() == EOPERATION__ECONTAINING_CLASS {
		return eOperation.EContainer().(EClass)
	}
	return nil
}

// GetEExceptions get the value of eExceptions
func (eOperation *eOperationImpl) GetEExceptions() EList {
	if eOperation.eExceptions == nil {
		eOperation.eExceptions = eOperation.getInitializers().initEExceptions()
	}
	return eOperation.eExceptions
}

// UnsetEExceptions unset the value of eExceptions
func (eOperation *eOperationImpl) UnsetEExceptions() {
	if eOperation.eExceptions != nil {
		eOperation.eExceptions.Clear()
	}
}

// GetEParameters get the value of eParameters
func (eOperation *eOperationImpl) GetEParameters() EList {
	if eOperation.eParameters == nil {
		eOperation.eParameters = eOperation.getInitializers().initEParameters()
	}
	return eOperation.eParameters
}

// GetOperationID get the value of operationID
func (eOperation *eOperationImpl) GetOperationID() int {
	return eOperation.operationID
}

// SetOperationID set the value of operationID
func (eOperation *eOperationImpl) SetOperationID(newOperationID int) {
	oldOperationID := eOperation.operationID
	eOperation.operationID = newOperationID
	if eOperation.ENotificationRequired() {
		eOperation.ENotify(NewNotificationByFeatureID(eOperation.AsEObject(), SET, EOPERATION__OPERATION_ID, oldOperationID, newOperationID, NO_INDEX))
	}
}

func (eOperation *eOperationImpl) initEExceptions() EList {
	return NewEObjectEList(eOperation.AsEObjectInternal(), EOPERATION__EEXCEPTIONS, -1, false, false, false, true, true)
}

func (eOperation *eOperationImpl) initEParameters() EList {
	return NewEObjectEList(eOperation.AsEObjectInternal(), EOPERATION__EPARAMETERS, EPARAMETER__EOPERATION, true, true, true, false, false)
}

func (eOperation *eOperationImpl) EGetFromID(featureID int, resolve, coreType bool) interface{} {
	switch featureID {
	case EOPERATION__ECONTAINING_CLASS:
		return eOperation.GetEContainingClass()
	case EOPERATION__EEXCEPTIONS:
		return eOperation.GetEExceptions()
	case EOPERATION__EPARAMETERS:
		return eOperation.GetEParameters()
	case EOPERATION__OPERATION_ID:
		return eOperation.GetOperationID()
	default:
		return eOperation.eTypedElementExt.EGetFromID(featureID, resolve, coreType)
	}
}

func (eOperation *eOperationImpl) ESetFromID(featureID int, newValue interface{}) {
	switch featureID {
	case EOPERATION__EEXCEPTIONS:
		e := newValue.(EList)
		eOperation.GetEExceptions().Clear()
		eOperation.GetEExceptions().AddAll(e)
	case EOPERATION__EPARAMETERS:
		e := newValue.(EList)
		eOperation.GetEParameters().Clear()
		eOperation.GetEParameters().AddAll(e)
	case EOPERATION__OPERATION_ID:
		o := newValue.(int)
		eOperation.SetOperationID(o)
	default:
		eOperation.eTypedElementExt.ESetFromID(featureID, newValue)
	}
}

func (eOperation *eOperationImpl) EUnsetFromID(featureID int) {
	switch featureID {
	case EOPERATION__EEXCEPTIONS:
		eOperation.UnsetEExceptions()
	case EOPERATION__EPARAMETERS:
		eOperation.GetEParameters().Clear()
	case EOPERATION__OPERATION_ID:
		eOperation.SetOperationID(-1)
	default:
		eOperation.eTypedElementExt.EUnsetFromID(featureID)
	}
}

func (eOperation *eOperationImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case EOPERATION__ECONTAINING_CLASS:
		return eOperation.GetEContainingClass() != nil
	case EOPERATION__EEXCEPTIONS:
		return eOperation.eExceptions != nil && eOperation.eExceptions.Size() != 0
	case EOPERATION__EPARAMETERS:
		return eOperation.eParameters != nil && eOperation.eParameters.Size() != 0
	case EOPERATION__OPERATION_ID:
		return eOperation.operationID != -1
	default:
		return eOperation.eTypedElementExt.EIsSetFromID(featureID)
	}
}

func (eOperation *eOperationImpl) EInvokeFromID(operationID int, arguments EList) interface{} {
	switch operationID {
	case EOPERATION__IS_OVERRIDE_OF_EOPERATION:
		return eOperation.IsOverrideOf(arguments.Get(0).(EOperation))
	default:
		return eOperation.eTypedElementExt.EInvokeFromID(operationID, arguments)
	}
}

func (eOperation *eOperationImpl) EBasicInverseAdd(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case EOPERATION__ECONTAINING_CLASS:
		msgs := notifications
		if eOperation.EContainer() != nil {
			msgs = eOperation.EBasicRemoveFromContainer(msgs)
		}
		return eOperation.EBasicSetContainer(otherEnd, EOPERATION__ECONTAINING_CLASS, msgs)
	case EOPERATION__EPARAMETERS:
		list := eOperation.GetEParameters().(ENotifyingList)
		return list.AddWithNotification(otherEnd, notifications)
	default:
		return eOperation.eTypedElementExt.EBasicInverseAdd(otherEnd, featureID, notifications)
	}
}

func (eOperation *eOperationImpl) EBasicInverseRemove(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case EOPERATION__ECONTAINING_CLASS:
		return eOperation.EBasicSetContainer(nil, EOPERATION__ECONTAINING_CLASS, notifications)
	case EOPERATION__EPARAMETERS:
		list := eOperation.GetEParameters().(ENotifyingList)
		return list.RemoveWithNotification(otherEnd, notifications)
	default:
		return eOperation.eTypedElementExt.EBasicInverseRemove(otherEnd, featureID, notifications)
	}
}
