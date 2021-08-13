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
