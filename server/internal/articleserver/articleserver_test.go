package articleserver

import (
	"context"
	"errors"
	"testing"

	"github.com/diSpector/incrowd.git/internal/models/domain"
	"github.com/diSpector/incrowd.git/internal/storage/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPrepareGetOneArticleHandlerResponseStatus(t *testing.T) {
	mockStorage := mocks.NewServerStorage(t)
	s := New(mockStorage)

	mockStorage.On("GetArticleById", mock.Anything, `test_not_found`).Return(nil, nil)
	mockStorage.On("GetArticleById", mock.Anything, `test_error`).Return(nil, errors.New("error db"))
	mockStorage.On("GetArticleById", mock.Anything, `test_ok`).Return(&domain.Article{Id: `test_ok`}, nil)

	ctx := context.Background()
	response, _ := s.prepareGetOneArticleHandlerResponse(ctx, `test_not_found`)
	assert.Equal(t, response.Status, STATUS_FAIL)

	response, _ = s.prepareGetOneArticleHandlerResponse(ctx, `test_error`)
	assert.Equal(t, response.Status, STATUS_ERROR)

	response, _ = s.prepareGetOneArticleHandlerResponse(ctx, `test_ok`)
	assert.Equal(t, response.Status, STATUS_SUCCESS)

	mockStorage.AssertExpectations(t)
}
