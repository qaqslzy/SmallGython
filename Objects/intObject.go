package Objects

import (
	"fmt"
)

/**
*
* @author Liu Weiyi
* @date 2020/4/25 1:19 下午
 */

type PyIntObject struct {
	PyObject_HEAD
	value int
}

type PyIntType tagPyTypeObject

var PyInt_Type = PyIntType{
	name:          "int",
	PyObject_HEAD: PyObject_HEAD_INIT((*tagPyTypeObject)(&PyType_Type)),
	PyTypeFunc:    &PyIntType{},
}

func PyInt_Create(value int) PyObject {
	object := PyIntObject{}
	object.refCount = 1
	object.TagPyTypeObject = (*tagPyTypeObject)(&PyInt_Type)
	object.value = value
	return &object
}

func (PyType PyIntType) AddFun(left, right PyObject) PyObject {
	intLeft := left.(*PyIntObject)
	intRight := right.(*PyIntObject)
	return PyInt_Create(intLeft.value + intRight.value)
}

func (PyType PyIntType) PrintFun(object PyObject) {
	fmt.Println(object.(*PyIntObject).value)
}

func (PyType PyIntType) HashFun(object PyObject) int32 {
	return int32(object.(*PyIntObject).value)
}

func (obj PyIntObject) GetType() *tagPyTypeObject {
	return obj.TagPyTypeObject
}
