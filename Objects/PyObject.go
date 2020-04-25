package Objects

/**
*
* @author Liu Weiyi
* @date 2020/4/25 1:00 下午
 */

type PyObject_HEAD struct {
	refCount        int
	TagPyTypeObject *tagPyTypeObject
}

func PyObject_HEAD_INIT(typePtr *tagPyTypeObject) (head PyObject_HEAD) {
	head.TagPyTypeObject = typePtr
	head.refCount = 0
	return
}

type PyObject interface{}

type PyTypeGet interface {
	GetType() *tagPyTypeObject
}

func (obj PyObject_HEAD) GetType() *tagPyTypeObject {
	return obj.TagPyTypeObject
}
