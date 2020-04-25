package Objects

/**
*
* @author Liu Weiyi
* @date 2020/4/25 1:06 下午
 */

type tagPyTypeObject struct {
	name string
	PyObject_HEAD
	PyTypeFunc
}

type PyType tagPyTypeObject

var PyType_Type = PyType{
	name:          "type",
	PyObject_HEAD: PyObject_HEAD_INIT(nil),
	PyTypeFunc:    nil,
}

type PyTypeFunc interface {
	PrintFun(object PyObject)
	AddFun(left, right PyObject) PyObject
	HashFun(object PyObject) int32
}
