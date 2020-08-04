package json

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

// TestUnmarshal .
func TestUnmarshal(t *testing.T) {
	jsonString := `{
		"string": "string",
		"integer": 1,
		"float": 2.3,
		"true": true,
		"false": false,
		"null": null,
		"array": [
			"string", 
			1, 
			2.3, 
			true, 
			false, 
			null, 
			[
				"string", 
				1, 
				2.3, 
				true, 
				false, 
				null,
				[
					"string", 
					1, 
					2.3, 
					true, 
					false, 
					null
				],
				{
					"string": "string",
					"integer": 1,
					"float": 2.3,
					"true": true,
					"false": false,
					"null": null
				}
			],
			{
				"string": "string",
				"integer": 1,
				"float": 2.3,
				"true": true,
				"false": false,
				"null": null,
				"array": [
					"string", 
					1, 
					2.3, 
					true, 
					false, 
					null
				],
				"object": {
					"string": "string",
					"integer": 1,
					"float": 2.3,
					"true": true,
					"false": false,
					"null": null
				} 
			}
		],
		"object": {
			"string": "string",
			"integer": 1,
			"float": 2.3,
			"true": true,
			"false": false,
			"null": null,
			"array": [
				"string", 
				1, 
				2.3, 
				true, 
				false, 
				null
			],
			"object": {
				"string": "string",
				"integer": 1,
				"float": 2.3,
				"true": true,
				"false": false,
				"null": null
			} 
		} 
	}`
	var (
		any           Any
		err           error
		objectMapping map[string]Any
	)
	if any, err = Unmarshal([]byte(jsonString)); err != nil {
		t.Error(err)
		return
	}
	if objectMapping, err = any.ObjectValue(); err != nil {
		t.Error(err)
		return
	}
	if err = assertString(objectMapping["string"], "string"); err != nil {
		t.Error(err)
		return
	}
	if err = assertInt(objectMapping["integer"]); err != nil {
		t.Error(err)
		return
	}
	if err = assertFloat(objectMapping["float"]); err != nil {
		t.Error(err)
		return
	}
	if err = assertBool(objectMapping["true"], true); err != nil {
		t.Error(err)
		return
	}
	if err = assertBool(objectMapping["false"], false); err != nil {
		t.Error(err)
		return
	}
	if err = assertNull(objectMapping["null"]); err != nil {
		t.Error(err)
		return
	}
}

func assertNull(src Any) error {
	if !src.Null() {
		return errors.New("src is not null")
	}
	if str, err := src.StringValue(); err != nil {
		return err
	} else if str != "" {
		return errors.New(`null node should return "" by ::StringValue`)
	}
	if str, err := src.AsText(); err != nil {
		return err
	} else if str != "" {
		return errors.New(`null node should return "" by ::AsText`)
	}
	if i, err := src.IntValue(); err != nil {
		return err
	} else if i != 0 {
		return errors.New(`null node should return 0 by ::IntValue`)
	}
	if i, err := src.AsInt(); err != nil {
		return err
	} else if i != 0 {
		return errors.New(`null node should return 0 by ::AsInt`)
	}
	if f, err := src.FloatValue(); err != nil {
		return err
	} else if f != 0 {
		return errors.New(`null node should return 0.0 by ::FloatValue`)
	}
	if f, err := src.AsFloat(); err != nil {
		return err
	} else if f != 0 {
		return errors.New(`null node should return 0.0 by ::AsFloat`)
	}
	if b, err := src.BoolValue(); err != nil {
		return err
	} else if b {
		return errors.New(`null node should return false by ::BoolValue`)
	}
	if b, err := src.AsBool(); err != nil {
		return err
	} else if b {
		return errors.New(`null node should return false by ::AsBool`)
	}
	if _, err := src.ArrayValue(); err == nil {
		return errors.New(`null node should return err by ::ArrayValue`)
	}
	if _, err := src.ObjectValue(); err == nil {
		return errors.New(`null node should return err by ::ObjectValue`)
	}
	return nil
}

func assertString(src Any, value string) error {
	if src.Null() {
		return errors.New("src is null")
	}
	if str, err := src.StringValue(); err != nil {
		return err
	} else if str != value {
		return fmt.Errorf("expecting %v, but is %v", value, str)
	}
	if str, err := src.AsText(); err != nil {
		return err
	} else if str != value {
		return fmt.Errorf("expecting %v, but is %v", value, str)
	}
	if _, err := src.IntValue(); err == nil {
		return errors.New(`string node should return error by ::IntValue`)
	}
	iValue, parseIntError := strconv.ParseInt(value, 10, 64)
	if i, err := src.AsInt(); err != nil {
		if parseIntError == nil {
			return fmt.Errorf("%v could parsed as int", value)
		}
	} else if parseIntError != nil {
		return fmt.Errorf("%v couldn't parsed as int", value)
	} else if i != iValue {
		return fmt.Errorf(`expecting parsed value %v, but is %v`, iValue, i)
	}
	if _, err := src.FloatValue(); err == nil {
		return errors.New(`string node should return error by ::FloatValue`)
	}
	fValue, parseFloatError := strconv.ParseFloat(value, 64)
	if f, err := src.AsFloat(); err != nil {
		if parseFloatError == nil {
			return fmt.Errorf("%v could parsed as float", value)
		}
	} else if parseFloatError != nil {
		return fmt.Errorf("%v couldn't parsed as float", value)
	} else if f != fValue {
		return fmt.Errorf(`expecting parsed value %v, but is %v`, fValue, f)
	}
	if _, err := src.BoolValue(); err == nil {
		return errors.New(`string node should return error by ::BoolValue`)
	}
	if b, err := src.AsBool(); err != nil {
		return err
	} else if b {
		return errors.New(`null node should return false by ::AsBool`)
	}
	if _, err := src.ArrayValue(); err == nil {
		return errors.New(`null node should return err by ::ArrayValue`)
	}
	if _, err := src.ObjectValue(); err == nil {
		return errors.New(`null node should return err by ::ObjectValue`)
	}
	return nil
}

func assertInt(src Any) error {
	if src.Null() {
		return errors.New("src is null")
	}
	if _, err := src.IntValue(); err != nil {
		return err
	}
	return nil
}

func assertFloat(src Any) error {
	if src.Null() {
		return errors.New("src is null")
	}
	if _, err := src.FloatValue(); err != nil {
		return err
	}
	return nil
}

func assertBool(src Any, value bool) error {
	if src.Null() {
		return errors.New("src is null")
	}
	if b, err := src.BoolValue(); err != nil {
		return err
	} else if b != value {
		return fmt.Errorf("expect %v, but is %v", value, b)
	}
	return nil
}

func assertArray(src Any) error {
	if src.Null() {
		return errors.New("src is null")
	}
	if _, err := src.ArrayValue(); err != nil {
		return err
	}
	return nil
}

func assertObject(src Any) error {
	if src.Null() {
		return errors.New("src is null")
	}
	if _, err := src.ObjectValue(); err != nil {
		return err
	}
	return nil
}
