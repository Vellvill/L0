package model

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"io/ioutil"
	"testing"
)

func TestModel(path string, t *testing.T) (*Model, error) {
	var testModel Model

	data, err := ioutil.ReadFile(path)

	err = validation.Validate(data, is.JSON)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &testModel.Json)

	testModel.OrderUID = testModel.Json.OrderUID

	err = testModel.Validate()

	if err != nil {
		return nil, err
	}
	return &testModel, nil
}
