package Parser

import (
	"SmallGython/Objects"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
*
* @author Liu Weiyi
* @date 2020/4/25 7:13 下午
 */
const (
	info   = "********** Python Research **********"
	prompt = ">>> "
)

var m_LocatEnvironment = Objects.PyDict_Create()

func Execute() {
	fmt.Println(info)
	fmt.Print(prompt)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		command := input.Text()
		if len(command) == 0 {
			fmt.Print(prompt)
			continue
		} else if command == "exit" {
			return
		} else {
			ExecuteCommand(command)
		}
		fmt.Print(prompt)
	}
}

func ExecuteCommand(command string) {
	pos := 0
	if pos = strings.Index(command, "print "); pos != -1 {
		ExecutePrint(command[6:])
	} else if pos = strings.Index(command, " = "); pos != -1 {
		target := command[:pos]
		source := command[pos+3:]
		ExecuteAdd(target, source)
	}
}

func ExecuteAdd(target, source string) {
	if num, err := strconv.Atoi(source); err == nil {
		intValue := Objects.PyInt_Create(num)
		key := Objects.PyStr_Create(target)
		Objects.PyDict_SetItem(m_LocatEnvironment, key, intValue)
	} else if strings.Index(source, "\"") != -1 {
		strValue := Objects.PyStr_Create(source[1 : len(source)-1])
		key := Objects.PyStr_Create(target)
		Objects.PyDict_SetItem(m_LocatEnvironment, key, strValue)
	} else if strings.Index(source, "+") != -1 {
		source = strings.ReplaceAll(source, " ", "")
		pos := strings.Index(source, "+")
		var leftObj, rightObj Objects.PyObject

		if leftNum, err := strconv.Atoi(source[:pos]); err == nil {
			leftObj = Objects.PyInt_Create(leftNum)
		} else {
			leftObj = GetObjectBySymbol(source[:pos])
		}

		if rightNum, err := strconv.Atoi(source[pos+1:]); err == nil {
			rightObj = Objects.PyInt_Create(rightNum)
		} else {
			rightObj = GetObjectBySymbol(source[pos+1:])
		}

		if rightObj != nil && leftObj != nil &&
			(rightObj.(Objects.PyTypeGet).GetType() == leftObj.(Objects.PyTypeGet).GetType()) {
			resultValue := rightObj.(Objects.PyTypeGet).GetType().AddFun(leftObj, rightObj)
			key := Objects.PyStr_Create(target)
			Objects.PyDict_SetItem(m_LocatEnvironment, key, resultValue)
		}
	}
}

func ExecutePrint(symbol string) {
	obj := GetObjectBySymbol(symbol)
	obj.(Objects.PyTypeGet).GetType().PrintFun(obj)
}

func GetObjectBySymbol(symbol string) Objects.PyObject {
	key := Objects.PyStr_Create(symbol)
	value := Objects.PyDict_GetItem(m_LocatEnvironment, key)
	if value == nil {
		fmt.Println("[Error] :", symbol, "is not defined!")
	}
	return value
}
