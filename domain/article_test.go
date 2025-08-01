package domain_test

import (
	"onion-sample-go/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArticle_Success(t *testing.T) {
	article, err := domain.NewArticle(1, "Test Article")

	assert.NoError(t, err)
	assert.NotNil(t, article)
}
