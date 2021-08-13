package fj

import (
	"io/ioutil"
	"os"
	"strconv"

	jsonparser "github.com/buger/jsonparser"
)

type FastJson struct {
	content []byte
}

func (fj *FastJson) GetContent() []byte {
	return fj.content
}

func (fj *FastJson) ClearFileAndOutput(fileurl string) error {
	desFile, err := os.OpenFile(fileurl, os.O_RDWR, 0666)
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
	return GetString(fj.content, keys...)
}

func GetInt(data []byte, keys ...string) (val int64, err error) {
	return jsonparser.GetInt(data, keys...)
}

func (fj *FastJson) GetInt(keys ...string) (val int64, err error) {
	return GetInt(fj.content, keys...)
}

func GetBoolean(data []byte, keys ...string) (val bool, err error) {
	return jsonparser.GetBoolean(data, keys...)
}
func (fj *FastJson) GetBoolean(keys ...string) (val bool, err error) {
	return GetBoolean(fj.content, keys...)
}

func GetFloat(data []byte, keys ...string) (val float64, err error) {
	return jsonparser.GetFloat(data, keys...)
}
func (fj *FastJson) GetFloat(keys ...string) (val float64, err error) {
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

func (fj *FastJson) SetString(val string, keys ...string) error {
	return fj.Set([]byte(val), keys...)
}

func SetInt(data []byte, val int, keys ...string) (value []byte, err error) {
	return Set(data, []byte(strconv.Itoa(val)), keys...)
}

func (fj *FastJson) SetInt(val int, keys ...string) error {
	return fj.Set([]byte(strconv.Itoa(val)), keys...)
}

func SetBoolean(data []byte, val bool, keys ...string) (value []byte, err error) {
	return Set(data, []byte(strconv.FormatBool(val)), keys...)
}

func (fj *FastJson) SetBoolean(val bool, keys ...string) error {
	return fj.Set([]byte(strconv.FormatBool(val)), keys...)
}

func SetFloat(data []byte, val float64, keys ...string) (value []byte, err error) {
	return Set(data, []byte(strconv.FormatFloat(val, 'E', -1, 64)), keys...)
}

func (fj *FastJson) SetFloat(val float64, keys ...string) error {
	return fj.Set([]byte(strconv.FormatFloat(val, 'f', -1, 64)), keys...)
}

func Delete(data []byte, keys ...string) []byte {
	return jsonparser.Delete(data, keys...)
}
func (fj *FastJson) Delete(keys ...string) []byte {
	return Delete(fj.content, keys...)
}

func NewFromFile(filepath string) (*FastJson, error) {
	jdata, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return &FastJson{content: jdata}, nil
}

func NewFromString(strcontent string) (*FastJson, error) {
	return &FastJson{[]byte(strcontent)}, nil
}

func NewFromBytes(content []byte) (*FastJson, error) {
	return &FastJson{content}, nil
}
