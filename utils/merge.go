package utils

import (
	"errors"
	"reflect"
)

/* 功能：把从表单提交的数据（结构体）合并到原始数据（结构体，一般读取自数据库）
 * 参数：dataOriginal（原始数据，一般读取自数据库，此参数同时也是输出参数，输出成功时最终的合并结果）
 *             dataSubmit（从表单提交的数据）
 * 返回值：err（错误信息，值为nil时表示合并成功，否则合并失败）
 */
func MergeData(dataOriginal interface{}, dataSubmit interface{}) (err error) {
	typeOfSubmit := reflect.TypeOf(dataSubmit)
	valueOfSubmit := reflect.ValueOf(dataSubmit)
	valueOfOriginal := reflect.ValueOf(dataOriginal)

	// 数据类型判断
	if valueOfOriginal.Kind() != reflect.Ptr {
		err = errors.New("The original data isn't a pointer to struct")
		return
	}
	elemOriginal := valueOfOriginal.Elem()
	if elemOriginal.Kind() != reflect.Struct {
		err = errors.New("The original pointer doesn't point to struct")
		return
	}
	if typeOfSubmit.Kind() != reflect.Struct {
		err = errors.New("The sumbit data isn't struct")
		return
	}

	// 遍历节点，进行合并
	for i := 0; i < typeOfSubmit.NumField(); i++ {
		strName := typeOfSubmit.Field(i).Name
		filedOriginal := elemOriginal.FieldByName(strName)
		fieldSubmit := valueOfSubmit.FieldByName(strName)
		if filedOriginal.Kind() == reflect.Ptr {
			if fieldSubmit.Kind() == reflect.Ptr {
				elemField1 := filedOriginal.Elem()
				elemField2 := fieldSubmit.Elem()
				if elemField1.Kind() == reflect.Struct && elemField2.Kind() == reflect.Struct {
					MergeData(filedOriginal.Interface(), elemField2.Interface())
				}
			}
		} else {
			// logs.Debug("%#v...............%#v", strName, elemOriginal.FieldByName(strName))
			if filedOriginal.CanSet() && !isValueBlank(fieldSubmit) {
				filedOriginal.Set(fieldSubmit)
				// logs.Debug("%#v***************%#v", strName, elemOriginal.FieldByName(strName))
			}
		}
	}

	return
}

/* 功能：判断数据是否为空
 * 参数：value（要判断的数据）
 * 返回值：bool（true表示为空，否则不为空）
 */
func isValueBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}
