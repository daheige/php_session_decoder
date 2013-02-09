package php_session_decoder

import (
	"testing"
)


//TODO: write bool false parse test
// negative int parse test
// negative float parse test
// string parse test
// string with quotes parse test
// string with unicode parse
// string with $/@ and etc 
func TestDecoderFabrica(t *testing.T) {
	decoder := NewPhpDecoder("")
	if decoder == nil {
		t.Error("Can not create decoder object\n")
	}
}

const BOOLEAN_VALUE_ENCODED = "login_ok|b:1;"
func TestDecodeBooleanValueWithoutName(t *testing.T) {
	decoder := NewPhpDecoder(BOOLEAN_VALUE_ENCODED_WITHOUT_NAME)
	if result, err := decoder.DecodeValue(); err != nil {
		t.Errorf("Can not decode boolens value %#v \n", err)
	} else {
		if v, ok := (result).(bool); !ok {
			t.Errorf("Boolean value was not decoded \n")
		} else if v != true {
			t.Errorf("Boolean value was incorrectly decoded \n")
		}
	}
}

const BOOLEAN_VALUE_ENCODED_WITHOUT_NAME = "b:1;"
func TestDecodeBooleanValue(t *testing.T) {
	decoder := NewPhpDecoder(BOOLEAN_VALUE_ENCODED)
	if result, err := decoder.Decode(); err != nil {
		t.Errorf("Can not decode boolens value %#v \n", err)
	} else {
		if v, ok := (*result)["login_ok"]; !ok {
			t.Errorf("Boolean value was not decoded \n")
		} else if v != true {
			t.Errorf("Boolean value was incorrectly decoded \n")
		}
	}
}

const INT_VALUE_ENCODED = "inteiro|i:34;"
func TestDecodeIntValue(t *testing.T) {
	decoder := NewPhpDecoder(INT_VALUE_ENCODED)
	if result, err := decoder.Decode(); err != nil {
		t.Errorf("Can not decode int value %#v \n", err)
	} else {
		if v, ok := (*result)["inteiro"]; !ok {
			t.Errorf("Int value was not decoded \n")
		} else if v != 34 {
			t.Errorf("Int value was decoded incorrectly: %v\n", v)
		}
	}
}

const BOOLEAN_AND_INT_ENCODED = "login_ok|b:1;inteiro|i:34;"
func TestDecodeBooleanAndIntValue(t *testing.T) {
	decoder := NewPhpDecoder(BOOLEAN_AND_INT_ENCODED)
	if result, err := decoder.Decode(); err != nil {
		t.Errorf("Can not decode int value %#v \n", err)
	} else {
		if v, ok := (*result)["inteiro"]; !ok {
			t.Errorf("Int value was not decoded \n")
		} else if v != 34 {
			t.Errorf("Int value was decoded incorrectly: %v\n", v)
		}
	}
}

const FLOAT_VALUE_ENCODED = "float_test|d:34.467999999900002;"
func TestDecodeFloatValue(t *testing.T) {
	decoder := NewPhpDecoder(FLOAT_VALUE_ENCODED)
	if result, err := decoder.Decode(); err != nil {
		t.Errorf("Can not decode float value %#v \n", err)
	} else {
		if v, ok := (*result)["float_test"]; !ok {
			t.Errorf("Float value was not decoded \n")
		} else if v != 34.467999999900002 {
			t.Errorf("Float value was decoded incorrectly: %v\n", v)
		}
	}
}

const STRING_VALUE_ENCODED = "name|s:9:\"some text\";"
func TestDecodeStringValue(t *testing.T) {
	decoder := NewPhpDecoder(STRING_VALUE_ENCODED)
	if result, err := decoder.Decode(); err != nil {
		t.Errorf("Can not decode string value %#v \n", err)
	} else {
		if v, ok := (*result)["name"]; !ok {
			t.Errorf("String value was not decoded \n")
		} else if v != "some text" {
			t.Errorf("String value was decoded incorrectly: %v\n", v)
		}
	}
}

const ARRAY_VALUE_ENCODED = "arr|a:2:{s:4:\"test\";b:1;i:0;i:5;}" 
func TestDecodeArrayValue(t *testing.T) {
	decoder := NewPhpDecoder(ARRAY_VALUE_ENCODED)
	if result, err := decoder.Decode(); err != nil {
		t.Errorf("Can not decode array value %#v \n", err)
	} else {
		if v, ok := (*result)["arr"]; !ok {
			t.Errorf("Array value was not decoded \n")
		} else if arrType := reflect.TypeOf(v); arrType.Name() != "map" {
			t.Errorf("Array value was decoded incorrectly: %v\n", arrType.Name())
		} 
                //TODO: check array fields
	}
}

