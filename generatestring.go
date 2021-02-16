package struct2str

import (
	"errors"
	"fmt"
	"github.com/luckydog8686/logs"
	"reflect"
	"sort"
)

func GenerateString(obj interface{}) (string, error) {
	if reflect.TypeOf(obj).Kind() == reflect.Ptr {
		if reflect.TypeOf(obj).Elem().Kind() == reflect.Struct {
			return GenerateStringFromStructPtr(obj,true)
		}
	}
	if reflect.TypeOf(obj).Kind() == reflect.Struct {
		return GenerateStringFromStructPtr(obj,false)
	}
	return "", errors.New("Unsupport Type")
}

func GenerateStringFromStructPtr(obj interface{},isPtr bool) (string, error) {
	var t reflect.Type
	var iValue reflect.Value
	if isPtr{
		t = reflect.TypeOf(obj).Elem()
		iValue = reflect.ValueOf(obj).Elem()
	}else {
		t = reflect.TypeOf(obj)
		iValue = reflect.ValueOf(obj)
	}


	numField := t.NumField()
	structKeys := make([]string, 0, numField)
	for i := 0; i < numField; i++ {
		structKeys = append(structKeys, t.Field(i).Name)
	}
	sort.Strings(structKeys)
	msg := ""
	for i := 0; i < numField; i++ {
		fieldName := structKeys[i]
		key := fieldName
		field, ok := t.FieldByName(fieldName)
		if !ok {
			return "", errors.New(fmt.Sprintf("illegal filed name in reflect:%s", fieldName))
		}
		jsonTagValue, ok := field.Tag.Lookup("json")
		if ok {
			key = jsonTagValue
		}

		if i == 0 {
			msg = fmt.Sprintf("%s=%v", key,iValue.FieldByName(fieldName).Interface())
		}else{
			msg = fmt.Sprintf("%s&%s=%v",msg, key,iValue.FieldByName(fieldName).Interface())
		}
		logs.Info(msg)
	}
	return msg, nil
}

func GenerateStringFromStruct(obj interface{}) (string, error) {

	return "", nil
}
