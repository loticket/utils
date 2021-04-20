package maper

// @Title  struct 转 map
// @Description  struct 中的tag 转换成 map
// @Author youjixiaozhao
// @Update 2021年4月20日
func ToMap(in interface{}, dbTag string) map[string]interface{} {
	out := make(map[string]interface{})
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(dbTag); tagv != "" {
			// set key of map to value in struct field
			val := v.Field(i)
			zero := reflect.Zero(val.Type()).Interface()
			current := val.Interface()

			if reflect.DeepEqual(current, zero) {
				continue
			}
			out[tagv] = current
		}
	}

	return out
}

// @Title  struct 转 array[数组]
// @Description  struct 中的tag 转换成 array[数组]
// @Author youjixiaozhao
// @Update 2021年4月20日
func FieldNames(in interface{}, dbTag string) []string {
	out := make([]string, 0)
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(dbTag); tagv != "" {
			out = append(out, tagv)
		} else {
			out = append(out, fi.Name)
		}
	}

	return out
}

// @Title  struct 转 array[数组]
// @Description  struct 中的tag 转换成 array[数组] 数组中的字段带有``用于数据库查询
// @Author youjixiaozhao
// @Update 2021年4月20日
func RawFieldNames(in interface{}) []string {
	out := make([]string, 0)
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(dbTag); tagv != "" {
			out = append(out, fmt.Sprintf("`%s`", tagv))
		} else {
			out = append(out, fmt.Sprintf(`"%s"`, fi.Name))
		}
	}

	return out
}

// @Title  struct 转 array[数组] 根据tag获取到结构体中的tag 包含mark中的字段
// @Description  struct 中的tag 转换成 array[数组] 数组中的字段带有``用于数据库查询
// @Author youjixiaozhao
// @Update 2021年4月20日
func RawFieldNamesByMart(in interface{}, mark string, tag string) []string {
	out := make([]string, 0)
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			fmt.Println(tagv)
			var tagArr []string = strings.Split(tagv, ",")

			if stringx.Contains(tagArr, mark) {
				out = append(out, fmt.Sprintf("`%s`", tagArr[0]))
			}
		}
	}

	return out
}

// @Title  struct 转 map
// @Description  struct 中的tag 转换成 map 据tag获取到结构体中的tag 包含mark中的字段
// @Author youjixiaozhao
// @Update 2021年4月20日
func ToMapByMart(in interface{}, tag string, mark string) map[string]interface{} {
	out := make(map[string]interface{})
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			var tagArr []string = strings.Split(tagv, ",")
			if mark != "" && !stringx.Contains(tagArr, mark) {
				continue
			}
			val := v.Field(i)
			//zero := reflect.Zero(val.Type()).Interface()
			current := val.Interface()

			//if reflect.DeepEqual(current, zero) {
			//	continue
			//}
			out[tagArr[0]] = current
		}
	}

	return out
}
