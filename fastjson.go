package fj

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	jsonparser "github.com/daqnext/jsonparser"
)

type FastJson struct {
	content      []byte
	cacheString  map[string]string
	cacheInt64   map[string]int64
	cacheInt     map[string]int
	cacheBool    map[string]bool
	cacheFloat64 map[string]float64

	cacheStringArray  map[string][]string
	cacheIntArray     map[string][]int
	cacheInt64Array   map[string][]int64
	cacheFloat64Array map[string][]float64
}

func (fj *FastJson) GetContent() []byte {
	return fj.content
}

func (fj *FastJson) GetContentAsString() string {
	return string(fj.content)
}

func (fj *FastJson) ClearFileAndOutput(fileurl string) error {
	desFile, err := os.OpenFile(fileurl, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer desFile.Close()
	desFile.Truncate(0)
	desFile.Seek(0, 0)
	desFile.Write(fj.content)
	desFile.Sync()
	return nil
}

func GetString(data []byte, keys ...string) (val string, err error) {
	return jsonparser.GetString(data, keys...)
}

func (fj *FastJson) GetString(keys ...string) (val string, err error) {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}
	if val, ok := fj.cacheString[longkey]; ok {
		return val, nil
	}
	return GetString(fj.content, keys...)
}

func GetStringArray(data []byte, keys ...string) ([]string, error) {

	var Result []string
	var ParseError error

	_, err := ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if ParseError != nil {
			return
		}
		if err != nil {
			Result = nil
			ParseError = errors.New("parse error")
			return
		}

		Result = append(Result, string(value))
	}, keys...)

	if err != nil {
		return nil, err
	}

	if ParseError != nil {
		return nil, ParseError
	}

	return Result, nil
}

func (fj *FastJson) GetStringArray(keys ...string) ([]string, error) {
	return GetStringArray(fj.content, keys...)
}

func GetInt64(data []byte, keys ...string) (val int64, err error) {
	return jsonparser.GetInt(data, keys...)
}

func (fj *FastJson) GetInt64(keys ...string) (val int64, err error) {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}
	if val, ok := fj.cacheInt64[longkey]; ok {
		return val, nil
	}
	return GetInt64(fj.content, keys...)
}

func GetInt64Array(data []byte, keys ...string) ([]int64, error) {

	var Result []int64
	var ParseError error

	_, err := ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if ParseError != nil {
			return
		}
		if err != nil {
			Result = nil
			ParseError = errors.New("parse error")
			return
		}
		int64v, errint := strconv.ParseInt(string(value), 10, 64)
		if errint != nil {
			Result = nil
			ParseError = errors.New("parse int64 error")
		}

		Result = append(Result, int64v)
	}, keys...)

	if err != nil {
		return nil, err
	}

	if ParseError != nil {
		return nil, ParseError
	}

	return Result, nil
}

func (fj *FastJson) GetInt64Array(keys ...string) ([]int64, error) {
	return GetInt64Array(fj.content, keys...)
}

func GetInt(data []byte, keys ...string) (val int, err error) {
	int64result, err := jsonparser.GetInt(data, keys...)
	if err != nil {
		return int(int64result), err
	}
	return int(int64result), nil
}

func (fj *FastJson) GetInt(keys ...string) (val int, err error) {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}
	if val, ok := fj.cacheInt[longkey]; ok {
		return val, nil
	}
	return GetInt(fj.content, keys...)
}

func GetIntArray(data []byte, keys ...string) ([]int, error) {

	var Result []int
	var ParseError error

	_, err := ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if ParseError != nil {
			return
		}
		if err != nil {
			Result = nil
			ParseError = errors.New("parse error")
			return
		}
		intv, errint := strconv.Atoi(string(value))
		if errint != nil {
			Result = nil
			ParseError = errors.New("parse int error")
		}

		Result = append(Result, intv)
	}, keys...)

	if err != nil {
		return nil, err
	}

	if ParseError != nil {
		return nil, ParseError
	}

	return Result, nil
}

func (fj *FastJson) GetIntArray(keys ...string) ([]int, error) {
	return GetIntArray(fj.content, keys...)
}

func GetBoolean(data []byte, keys ...string) (val bool, err error) {
	return jsonparser.GetBoolean(data, keys...)
}
func (fj *FastJson) GetBoolean(keys ...string) (val bool, err error) {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}
	if val, ok := fj.cacheBool[longkey]; ok {
		return val, nil
	}
	return GetBoolean(fj.content, keys...)
}

func GetFloat64(data []byte, keys ...string) (val float64, err error) {
	return jsonparser.GetFloat(data, keys...)
}
func (fj *FastJson) GetFloat64(keys ...string) (val float64, err error) {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}
	if val, ok := fj.cacheFloat64[longkey]; ok {
		return val, nil
	}
	return GetFloat64(fj.content, keys...)
}

func GetFloat64Array(data []byte, keys ...string) ([]float64, error) {

	var Result []float64
	var ParseError error

	_, err := ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if ParseError != nil {
			return
		}
		if err != nil {
			Result = nil
			ParseError = errors.New("parse error")
			return
		}
		float64v, errint := strconv.ParseFloat(string(value), 64)
		if errint != nil {
			Result = nil
			ParseError = errors.New("parse int64 error")
		}

		Result = append(Result, float64v)
	}, keys...)

	if err != nil {
		return nil, err
	}

	if ParseError != nil {
		return nil, ParseError
	}

	return Result, nil
}

func (fj *FastJson) GetFloat64Array(keys ...string) ([]float64, error) {
	return GetFloat64Array(fj.content, keys...)
}

func ArrayEach(data []byte, cb func(value []byte, dataType jsonparser.ValueType, offset int, err error), keys ...string) (offset int, err error) {
	return jsonparser.ArrayEach(data, cb, keys...)
}

func (fj *FastJson) ArrayEach(itemfunc func(value []byte, offset int, err error), keys ...string) (offset int, err error) {
	return ArrayEach(fj.content, func(value []byte, _ jsonparser.ValueType, offset int, err error) {
		itemfunc(value, offset, err)
	}, keys...)
}

func ObjectEach(data []byte, callback func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error, keys ...string) (err error) {
	return jsonparser.ObjectEach(data, callback, keys...)
}

func (fj *FastJson) ObjectEach(itemfunc func(key []byte, value []byte, offset int) error, keys ...string) (err error) {
	return ObjectEach(fj.content, func(key []byte, value []byte, _ jsonparser.ValueType, offset int) error {
		return itemfunc(key, value, offset)
	}, keys...)
}

func Set(data []byte, setValue []byte, keys ...string) (value []byte, err error) {
	return jsonparser.Set(data, setValue, keys...)
}

func (fj *FastJson) Set(setValue []byte, keys ...string) error {

	result, err := Set(fj.content, setValue, keys...)

	if err != nil {
		return err
	}
	fj.content = result
	return nil
}

func SetString(data []byte, val string, keys ...string) (value []byte, err error) {
	return Set(data, []byte(val), keys...)
}

func SetStringArray(data []byte, vals []string, keys ...string) (value []byte, err error) {

	composedValue := "["
	for _, val := range vals {
		composedValue = composedValue + strconv.Quote(val) + ","
	}
	composedValue = strings.Trim(composedValue, ",")
	composedValue = composedValue + "]"
	return Set(data, []byte(composedValue), keys...)
}

func (fj *FastJson) SetString(val string, keys ...string) error {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}
	fj.cacheString[longkey] = val
	return fj.Set([]byte(strconv.Quote(val)), keys...)
}

func (fj *FastJson) SetStringArray(vals []string, keys ...string) error {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}

	composedValue := "["
	for _, val := range vals {
		composedValue = composedValue + strconv.Quote(val) + ","
	}
	composedValue = strings.Trim(composedValue, ",")
	composedValue = composedValue + "]"

	fj.cacheStringArray[longkey] = vals
	return fj.Set([]byte(composedValue), keys...)
}

func SetInt64(data []byte, val int64, keys ...string) (value []byte, err error) {
	return Set(data, []byte(strconv.FormatInt(val, 10)), keys...)
}

func SetInt64Array(data []byte, vals []int64, keys ...string) (value []byte, err error) {

	composedValue := "["
	for _, val := range vals {
		composedValue = composedValue + strconv.FormatInt(val, 10) + ","
	}
	composedValue = strings.Trim(composedValue, ",")
	composedValue = composedValue + "]"
	return Set(data, []byte(composedValue), keys...)
}

func (fj *FastJson) SetInt64(val int64, keys ...string) error {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}
	fj.cacheInt64[longkey] = val
	return fj.Set([]byte(strconv.FormatInt(val, 10)), keys...)
}

func (fj *FastJson) SetInt64Array(vals []int64, keys ...string) error {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}

	composedValue := "["
	for _, val := range vals {
		composedValue = composedValue + strconv.FormatInt(val, 10) + ","
	}
	composedValue = strings.Trim(composedValue, ",")
	composedValue = composedValue + "]"
	fj.cacheInt64Array[longkey] = vals
	return fj.Set([]byte(composedValue), keys...)
}

func SetInt(data []byte, val int, keys ...string) (value []byte, err error) {
	return Set(data, []byte(strconv.Itoa(val)), keys...)
}

func SetIntArray(data []byte, vals []int, keys ...string) (value []byte, err error) {

	composedValue := "["
	for _, val := range vals {
		composedValue = composedValue + strconv.Itoa(val) + ","
	}
	composedValue = strings.Trim(composedValue, ",")
	composedValue = composedValue + "]"
	return Set(data, []byte(composedValue), keys...)
}

func (fj *FastJson) SetInt(val int, keys ...string) error {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}
	fj.cacheInt[longkey] = val
	return fj.Set([]byte(strconv.Itoa(val)), keys...)
}

func (fj *FastJson) SetIntArray(vals []int, keys ...string) error {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}

	composedValue := "["
	for _, val := range vals {
		composedValue = composedValue + strconv.Itoa(val) + ","
	}
	composedValue = strings.Trim(composedValue, ",")
	composedValue = composedValue + "]"
	fj.cacheIntArray[longkey] = vals
	return fj.Set([]byte(composedValue), keys...)
}

func SetBoolean(data []byte, val bool, keys ...string) (value []byte, err error) {
	return Set(data, []byte(strconv.FormatBool(val)), keys...)
}

func (fj *FastJson) SetBoolean(val bool, keys ...string) error {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}
	fj.cacheBool[longkey] = val
	return fj.Set([]byte(strconv.FormatBool(val)), keys...)
}

func SetFloat64(data []byte, val float64, keys ...string) (value []byte, err error) {
	return Set(data, []byte(strconv.FormatFloat(val, 'E', -1, 64)), keys...)
}

func SetFloat64Array(data []byte, vals []float64, keys ...string) (value []byte, err error) {

	composedValue := "["
	for _, val := range vals {
		composedValue = composedValue + strconv.FormatFloat(val, 'f', -1, 64) + ","
	}
	composedValue = strings.Trim(composedValue, ",")
	composedValue = composedValue + "]"

	return Set(data, []byte(composedValue), keys...)
}

func (fj *FastJson) SetFloat64(val float64, keys ...string) error {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}
	fj.cacheFloat64[longkey] = val
	return fj.Set([]byte(strconv.FormatFloat(val, 'f', -1, 64)), keys...)
}

func (fj *FastJson) SetFloat64Array(vals []float64, keys ...string) error {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}

	composedValue := "["
	for _, val := range vals {
		composedValue = composedValue + strconv.FormatFloat(val, 'f', -1, 64) + ","
	}
	composedValue = strings.Trim(composedValue, ",")
	composedValue = composedValue + "]"
	fj.cacheFloat64Array[longkey] = vals
	return fj.Set([]byte(composedValue), keys...)
}

func Delete(data []byte, keys ...string) []byte {
	return jsonparser.Delete(data, keys...)
}

func (fj *FastJson) Delete(keys ...string) {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}
	delete(fj.cacheString, longkey)
	delete(fj.cacheBool, longkey)
	delete(fj.cacheInt64, longkey)
	delete(fj.cacheInt, longkey)
	delete(fj.cacheFloat64, longkey)

	fj.content = jsonparser.Delete(fj.content, keys...)
}

func NewFromFile(filepath string) (*FastJson, error) {
	jdata, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	if len(jdata) == 0 {
		jdata = []byte("{}")
	}
	return &FastJson{jdata, make(map[string]string), make(map[string]int64), make(map[string]int), map[string]bool{},
		make(map[string]float64), make(map[string][]string), make(map[string][]int), make(map[string][]int64), make(map[string][]float64)}, nil
}

func NewFromString(strcontent string) *FastJson {
	return &FastJson{[]byte(strcontent), make(map[string]string), make(map[string]int64), make(map[string]int), map[string]bool{}, make(map[string]float64),
		make(map[string][]string), make(map[string][]int), make(map[string][]int64), make(map[string][]float64)}
}

func NewFromBytes(content []byte) *FastJson {
	return &FastJson{content, make(map[string]string), make(map[string]int64), make(map[string]int), map[string]bool{}, make(map[string]float64),
		make(map[string][]string), make(map[string][]int), make(map[string][]int64), make(map[string][]float64)}
}
