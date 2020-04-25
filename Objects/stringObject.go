package Objects

import (
	"fmt"
	"hash/fnv"
)

/**
*
* @author Liu Weiyi
* @date 2020/4/25 3:09 下午
 */

type PyStringObject struct {
	PyObject_HEAD
	length    int
	hashValue int32
	value     string
}

type PyStringType tagPyTypeObject

var PyString_Type = PyStringType{
	name:          "string",
	PyObject_HEAD: PyObject_HEAD_INIT((*tagPyTypeObject)(&PyType_Type)),
	PyTypeFunc:    &PyStringType{},
}

func PyStr_Create(value string) PyObject {
	obj := PyStringObject{}
	obj.refCount = 1
	obj.TagPyTypeObject = (*tagPyTypeObject)(&PyString_Type)
	obj.value = value
	obj.length = len(value)
	obj.hashValue = -1
	return &obj
}

func (stringType PyStringType) PrintFun(obj PyObject) {
	fmt.Println(obj.(*PyStringObject).value)
}

func (stringType PyStringType) HashFun(obj PyObject) int32 {
	strObj := obj.(*PyStringObject)
	if int(strObj.hashValue) != -1 {
		return strObj.hashValue
	}
	h := fnv.New32a()
	h.Write([]byte(strObj.value))
	return int32(h.Sum32())
}

func (stringType PyStringType) AddFun(left, right PyObject) PyObject {
	result := PyStr_Create(left.(*PyStringObject).value + right.(*PyStringObject).value)
	return result
}

func (obj PyStringType) GetType() *tagPyTypeObject {
	return obj.TagPyTypeObject
}
