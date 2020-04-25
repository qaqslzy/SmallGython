package Objects

/**
*
* @author Liu Weiyi
* @date 2020/4/25 5:49 下午
 */

type PyDictObject struct {
	PyObject_HEAD
	dict map[int32]PyObject
}

type PyDictType tagPyTypeObject

var PyDict_Type = PyDictType{
	name:          "dict",
	PyObject_HEAD: PyObject_HEAD_INIT((*tagPyTypeObject)(&PyType_Type)),
	PyTypeFunc:    nil,
}

func PyDict_Create() PyObject {
	obj := PyDictObject{
		PyObject_HEAD: PyObject_HEAD_INIT((*tagPyTypeObject)(&PyDict_Type)),
		dict:          make(map[int32]PyObject),
	}
	obj.refCount = 1
	return &obj
}

func PyDict_GetItem(target, key PyObject) PyObject {
	hashValue := key.(PyTypeGet).GetType().HashFun(key)
	dict := target.(*PyDictObject).dict
	v, ok := dict[hashValue]
	if ok {
		return v
	}
	return nil
}

func PyDict_SetItem(target, key, value PyObject) int {
	hashValue := key.(PyTypeGet).GetType().HashFun(key)
	dict := target.(*PyDictObject).dict
	dict[hashValue] = value
	return 0
}

func (obj PyDictType) GetType() *tagPyTypeObject {
	return obj.TagPyTypeObject
}
