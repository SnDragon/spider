package reflect

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// CreateSQL
// @Description: 通过反射获取对应类型名和字段，生成SQL
// @param v
// @return string
func CreateSQL(v interface{}) string {
	if reflect.TypeOf(v).Kind() != reflect.Struct &&
		reflect.TypeOf(v).Kind() != reflect.Ptr {
		return ""
	}
	sql := ""
	var reflectVal reflect.Value
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		sql = fmt.Sprintf("insert into %s values(", reflect.TypeOf(v).Elem().Name())
		reflectVal = reflect.ValueOf(v).Elem()
	} else {
		sql = fmt.Sprintf("insert into %s values(", reflect.TypeOf(v).Name())
		reflectVal = reflect.ValueOf(v)
	}
	for i := 0; i < reflectVal.NumField(); i++ {
		switch reflectVal.Field(i).Kind() {
		case reflect.Int:
			sql += strconv.Itoa(int(reflectVal.Field(i).Int())) + ","
		case reflect.String:
			sql += "\"" + reflectVal.Field(i).String() + "\","
		}
	}
	sql = strings.TrimRight(sql, ",") + ")"
	return sql
}
