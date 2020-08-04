package json

import (
	"errors"
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
	if err = assertString(objectMapping["string"]); err != nil {
		t.Error(err)
		return
	}
}

func assertString(src Any) error {
	if src.Null() {
		return errors.New("src is null")
	}
	if _, err := src.StringValue(); err != nil {
		return err
	}
	return nil
}

func assertInt(src Any) error {
	if src.Null() {
		return errors.New("src is null")
	}
	if _, err := src.StringValue(); err != nil {
		return err
	}
	return nil
}

func assertFloat(src Any) error {
	if src.Null() {
		return errors.New("src is null")
	}
	if _, err := src.StringValue(); err != nil {
		return err
	}
	return nil
}

func assertBool(src Any, value bool) error {
	if src.Null() {
		return errors.New("src is null")
	}
	if _, err := src.StringValue(); err != nil {
		return err
	}
	return nil
}

func assertArray(src Any) error {
	if src.Null() {
		return errors.New("src is null")
	}
	if _, err := src.StringValue(); err != nil {
		return err
	}
	return nil
}

func assertObject(src Any) error {
	if src.Null() {
		return errors.New("src is null")
	}
	if _, err := src.StringValue(); err != nil {
		return err
	}
	return nil
}
