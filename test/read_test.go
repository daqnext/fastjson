package test

import (
	"testing"

	FastJson "github.com/daqnext/fastjson"
)

func TestFromFile(t *testing.T) {
	fj, err := FastJson.NewFromFile("../test.json")
	if err == nil {
		result, _ := fj.GetString("company", "name")
		if result == "acm" {
			t.Error("read error")
		}

	}
}

func TestFromString(t *testing.T) {
	fj, err := FastJson.NewFromString("{\"test\":123}")
	if err == nil {
		result, _ := fj.GetInt("test")
		if result != 123 {
			t.Error("read error")
		}
	}
}

func TestGetContentAsString(t *testing.T) {
	fj, err := FastJson.NewFromString("{\"test\":123}")
	if err == nil {
		result := fj.GetContentAsString()
		if result != "{\"test\":123}" {
			t.Error("read error")
		}
	}
}

func TestRead(t *testing.T) {
	fj, err := FastJson.NewFromFile("../test.json")
	if err == nil {

		result, _ := fj.GetString("company", "name")
		if result == "acm" {
			t.Error("read error")
		}

		_, err := fj.GetBoolean("company", "name")
		if err == nil {
			t.Error("get boolean error")
		}

		flnum, _ := fj.GetInt("person", "github", "followers")
		if flnum != 109 {
			t.Error("read int error")
		}
	}
}

func TestArray(t *testing.T) {
	fj, err := FastJson.NewFromFile("../test.json")
	if err == nil {
		fj.ArrayEach(func(value []byte, offset int, err error) {
			link, _ := FastJson.GetString(value, "url")
			if link != "link1" {
				t.Error("read nested error")
			}
		}, "person", "avatars")
	}
}

func TestArrayItem(t *testing.T) {
	fj, err := FastJson.NewFromFile("../test.json")
	if err == nil {
		link, _ := fj.GetString("person", "avatars", "[0]", "url")
		if link != "link1" {
			t.Error("read  array error")
		}
	}
}

func TestObject(t *testing.T) {
	fj, err := FastJson.NewFromFile("../test.json")
	if err == nil {
		fj.ObjectEach(func(key []byte, value []byte, offset int) error {
			if string(key) == "last" {
				if string(value) != "Bugaev" {
					t.Error("read object error")
				}
			}
			return nil
		}, "person", "name")
	}
}

func TestObjectItem(t *testing.T) {
	fj, err := FastJson.NewFromFile("../test.json")
	if err == nil {
		last, _ := fj.GetString("person", "name", "last")
		if last != "Bugaev" {
			t.Error("read  object item error")
		}
	}
}

func TestSetItemToFile(t *testing.T) {
	fj, err := FastJson.NewFromFile("../test.json")
	if err == nil {
		fj.SetFloat(123456.123, "person", "weight")
		err := fj.ClearFileAndOutput("../test.json")
		if err != nil {
			t.Error("something wrong", err)
		}
	}

	fj2, err2 := FastJson.NewFromFile("../test.json")
	if err2 == nil {
		r, _ := fj2.GetFloat("person", "weight")
		if r != 123456.123 {
			t.Error("read or set float error")
		}
	}

}
