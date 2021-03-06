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

	"github.com/stretchr/testify/assert"
)

func TestFactoryCreate(t *testing.T) {
	factory := newEcoreFactoryImpl()
	{
		mockEClass := &MockEClass{}
		mockEClass.On("GetClassifierID").Return(EANNOTATION)
		assert.NotNil(t, factory.Create(mockEClass))
	}
	{
		mockEClass := &MockEClass{}
		mockEClass.On("GetClassifierID").Return(EATTRIBUTE)
		assert.NotNil(t, factory.Create(mockEClass))
	}
	{
		mockEClass := &MockEClass{}
		mockEClass.On("GetClassifierID").Return(ECLASS)
		assert.NotNil(t, factory.Create(mockEClass))
	}
	{
		mockEClass := &MockEClass{}
		mockEClass.On("GetClassifierID").Return(EDATA_TYPE)
		assert.NotNil(t, factory.Create(mockEClass))
	}
	{
		mockEClass := &MockEClass{}
		mockEClass.On("GetClassifierID").Return(EENUM)
		assert.NotNil(t, factory.Create(mockEClass))
	}
	{
		mockEClass := &MockEClass{}
		mockEClass.On("GetClassifierID").Return(EENUM_LITERAL)
		assert.NotNil(t, factory.Create(mockEClass))
	}
	{
		mockEClass := &MockEClass{}
		mockEClass.On("GetClassifierID").Return(EFACTORY)
		assert.NotNil(t, factory.Create(mockEClass))
	}
	{
		mockEClass := &MockEClass{}
		mockEClass.On("GetClassifierID").Return(EGENERIC_TYPE)
		assert.NotNil(t, factory.Create(mockEClass))
	}
	{
		mockEClass := &MockEClass{}
		mockEClass.On("GetClassifierID").Return(EOBJECT)
		assert.NotNil(t, factory.Create(mockEClass))
	}
	{
		mockEClass := &MockEClass{}
		mockEClass.On("GetClassifierID").Return(EOPERATION)
		assert.NotNil(t, factory.Create(mockEClass))
	}
	{
		mockEClass := &MockEClass{}
		mockEClass.On("GetClassifierID").Return(EPACKAGE)
		assert.NotNil(t, factory.Create(mockEClass))
	}
	{
		mockEClass := &MockEClass{}
		mockEClass.On("GetClassifierID").Return(EPARAMETER)
		assert.NotNil(t, factory.Create(mockEClass))
	}
	{
		mockEClass := &MockEClass{}
		mockEClass.On("GetClassifierID").Return(EREFERENCE)
		assert.NotNil(t, factory.Create(mockEClass))
	}
	{
		mockEClass := &MockEClass{}
		mockEClass.On("GetClassifierID").Return(ESTRING_TO_STRING_MAP_ENTRY)
		assert.NotNil(t, factory.Create(mockEClass))
	}
	{
		mockEClass := &MockEClass{}
		mockEClass.On("GetClassifierID").Return(ETYPE_PARAMETER)
		assert.NotNil(t, factory.Create(mockEClass))
	}
}

func TestFactoryConvert(t *testing.T) {
	factory := newEcoreFactoryImpl()
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EBIG_DECIMAL)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EBIG_DECIMAL)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EBIG_INTEGER)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EBIG_INTEGER)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EBOOLEAN)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EBOOLEAN)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EBOOLEAN_OBJECT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EBOOLEAN_OBJECT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EBYTE)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EBYTE)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EBYTE_ARRAY)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EBYTE_ARRAY)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EBYTE_OBJECT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EBYTE_OBJECT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(ECHAR)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(ECHAR)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(ECHARACTER_OBJECT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(ECHARACTER_OBJECT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EDATE)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EDATE)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EDOUBLE)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EDOUBLE)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EDOUBLE_OBJECT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EDOUBLE_OBJECT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EFLOAT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EFLOAT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EFLOAT_OBJECT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EFLOAT_OBJECT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EINT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EINT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EINTEGER_OBJECT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EINTEGER_OBJECT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EJAVA_CLASS)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EJAVA_CLASS)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EJAVA_OBJECT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(EJAVA_OBJECT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(ELONG)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(ELONG)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(ELONG_OBJECT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(ELONG_OBJECT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(ESHORT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(ESHORT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(ESHORT_OBJECT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(ESHORT_OBJECT)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(ESTRING)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
	{
		mockEDataType := &MockEDataType{}
		mockEDataType.On("GetClassifierID").Return(ESTRING)
		assert.Panics(t, func() { factory.CreateFromString(mockEDataType, "") })
	}
}
