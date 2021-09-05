# fastjson
### 10x faster then standard lib &amp; dynamic set json
### examples

```go

go get github.com/daqnext/fastjson

import 
(
    FastJson "github.com/daqnext/fastjson"
)
```

```go

//////////////////////////////////////////////////////////////////

fj, err := FastJson.NewFromString("{\"test\":123}")
if err == nil {
	result, _ := fj.GetInt("test")
	if result != 123 {
		t.Error("read error")
	}
}

//////////////////////////////////////////////////////////////////

fj, err := FastJson.NewFromString("{\"test\":123}")
if err == nil {
	result := fj.GetContentAsString()
	if result != "{\"test\":123}" {
		t.Error("read error")
	}
}

//////////////////////////////////////////////////////////////////

fj, err := FastJson.NewFromFile("../test.json")
if err == nil {
	result, _ := fj.GetString("company", "name")
	if result == "acm" {
		t.Error("read error")
	}

}

//////////////////////////////////////////////////////////////////

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

//////////////////////////////////////////////////////////////////


fj, err := FastJson.NewFromFile("../test.json")
if err == nil {
	fj.ArrayEach(func(value []byte, offset int, err error) {
		link, _ := FastJson.GetString(value, "url")
		if link != "link1" {
			t.Error("read nested error")
		}
	}, "person", "avatars")
}

//////////////////////////////////////////////////////////////////


fj, err := FastJson.NewFromFile("../test.json")
if err == nil {
	link, _ := fj.GetString("person", "avatars", "[0]", "url")
	if link != "link1" {
		t.Error("read  array error")
	}
}

//////////////////////////////////////////////////////////////////

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

//////////////////////////////////////////////////////////////////


fj, err := FastJson.NewFromFile("../test.json")
if err == nil {
	last, _ := fj.GetString("person", "name", "last")
	if last != "Bugaev" {
		t.Error("read  object item error")
	}
}

//////////////////////////////////////////////////////////////////

fj, err := FastJson.NewFromFile("../test.json")
if err == nil {
	fj.SetFloat(123.123, "person", "weight")
	fj.ClearFileAndOutput("../test.json")
}

fj2, err2 := FastJson.NewFromFile("../test.json")
if err2 == nil {
	r, _ := fj2.GetFloat("person", "weight")
	if r != 123.123 {
		t.Error("read or set float error")
	}
}

//////////////////////////////////////////////////////////////////

```