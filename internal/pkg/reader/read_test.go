package reader_test

import (
	appErrors "pizza-delivery/internal/errors"
	"pizza-delivery/internal/pkg/reader"
	"testing"
)

func TestInvalidPathFailure(t *testing.T) {
	path := ""
	_, err := reader.Read(path)

	if err != appErrors.ErrorValidationEmptyInput {
		t.Error("expected an error based on an empty path param")
	}
}
