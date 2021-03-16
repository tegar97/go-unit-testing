package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang-unit-testing/repository"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {

	categoryRepository.Mock.On("fFindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}
