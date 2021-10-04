package retriever_test

import (
	appErrors "pizza-delivery/internal/errors"
	"pizza-delivery/internal/pkg/retriever"
	"testing"
)

func TestInvalidPathFailure(t *testing.T) {
	path := ""
	_, err := retriever.Read(path)

	if err != appErrors.ErrorValidationEmptyInput {
		t.Error("expected an error based on an empty path param")
	}
}
