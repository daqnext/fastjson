package fj

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	jsonparser "github.com/daqnext/jsonparser"
)

type FastJson struct {
	content     []byte
	cacheString map[string]string
	cacheInt64  map[string]int64
	cacheBool   map[string]bool
	cacheFloat  map[string]float64
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

func GetInt(data []byte, keys ...string) (val int64, err error) {
	return jsonparser.GetInt(data, keys...)
}

func (fj *FastJson) GetInt(keys ...string) (val int64, err error) {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}
	if val, ok := fj.cacheInt64[longkey]; ok {
		return val, nil
	}
	return GetInt(fj.content, keys...)
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

func GetFloat(data []byte, keys ...string) (val float64, err error) {
	return jsonparser.GetFloat(data, keys...)
}
func (fj *FastJson) GetFloat(keys ...string) (val float64, err error) {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}
	if val, ok := fj.cacheFloat[longkey]; ok {
		return val, nil
	}
	return GetFloat(fj.content, keys...)
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
	return fj.Set([]byte(composedValue), keys...)
}

func SetInt(data []byte, val int64, keys ...string) (value []byte, err error) {
	return Set(data, []byte(strconv.FormatInt(val, 10)), keys...)
}

func SetIntArray(data []byte, vals []int64, keys ...string) (value []byte, err error) {

	composedValue := "["
	for _, val := range vals {
		composedValue = composedValue + strconv.FormatInt(val, 10) + ","
	}
	composedValue = strings.Trim(composedValue, ",")
	composedValue = composedValue + "]"
	return Set(data, []byte(composedValue), keys...)
}

func (fj *FastJson) SetInt(val int64, keys ...string) error {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}
	fj.cacheInt64[longkey] = val
	return fj.Set([]byte(strconv.FormatInt(val, 10)), keys...)
}

func (fj *FastJson) SetIntArray(vals []int64, keys ...string) error {
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

func SetFloat(data []byte, val float64, keys ...string) (value []byte, err error) {
	return Set(data, []byte(strconv.FormatFloat(val, 'E', -1, 64)), keys...)
}

func SetFloatArray(data []byte, vals []float64, keys ...string) (value []byte, err error) {

	composedValue := "["
	for _, val := range vals {
		composedValue = composedValue + strconv.FormatFloat(val, 'f', -1, 64) + ","
	}
	composedValue = strings.Trim(composedValue, ",")
	composedValue = composedValue + "]"

	return Set(data, []byte(composedValue), keys...)
}

func (fj *FastJson) SetFloat(val float64, keys ...string) error {
	longkey := ""
	for i := 0; i < len(keys); i++ {
		longkey = longkey + keys[i]
	}
	fj.cacheFloat[longkey] = val
	return fj.Set([]byte(strconv.FormatFloat(val, 'f', -1, 64)), keys...)
}

func (fj *FastJson) SetFloatArray(vals []float64, keys ...string) error {
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
	delete(fj.cacheFloat, longkey)

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
	return &FastJson{jdata, make(map[string]string), make(map[string]int64), map[string]bool{}, make(map[string]float64)}, nil
}

func NewFromString(strcontent string) (*FastJson, error) {
	return &FastJson{[]byte(strcontent), make(map[string]string), make(map[string]int64), map[string]bool{}, make(map[string]float64)}, nil
}

func NewFromBytes(content []byte) (*FastJson, error) {
	return &FastJson{content, make(map[string]string), make(map[string]int64), map[string]bool{}, make(map[string]float64)}, nil
}
