package reposytories_test

import (
	"L0/internal/model"
	"L0/internal/reposytories"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepositoryPostgres_AddModel(t *testing.T) {
	s, down := reposytories.TestStore(t, DSN, MigrationsPath)
	defer down("models")

	testModel, err := model.TestModel("test_models/model.json", t)

	assert.NoError(t, err)

	assert.NoError(t, s.AddModel(context.TODO(), testModel, testModel.OrderUID))
	assert.NotNil(t, testModel)
}

func TestRepositoryPostgres_UpdateHash(t *testing.T) {
	s, down := reposytories.TestStore(t, DSN, MigrationsPath)
	defer down("models")

	testModel1, err := model.TestModel("test_models/model.json", t)
	testModel2, err := model.TestModel("test_models/model2.json", t)
	testModel3, err := model.TestModel("test_models/model3.json", t)

	err = s.AddModel(context.TODO(), testModel1, testModel1.OrderUID)
	err = s.AddModel(context.TODO(), testModel2, testModel1.OrderUID)
	err = s.AddModel(context.TODO(), testModel3, testModel1.OrderUID)

	assert.NoError(t, s.UpdateHash(context.TODO()))

	testdata1, err := json.Marshal(testModel1.Json)
	testdata2, err := json.Marshal(testModel2.Json)
	testdata3, err := json.Marshal(testModel3.Json)

	hashdata1, err := s.FindInHash(testModel1.OrderUID)
	hashdata2, err := s.FindInHash(testModel2.OrderUID)
	hashdata3, err := s.FindInHash(testModel3.OrderUID)

	assert.NoError(t, err)
	assert.Equal(t, testdata1, hashdata1)
	assert.Equal(t, testdata2, hashdata2)
	assert.Equal(t, testdata3, hashdata3)

}
