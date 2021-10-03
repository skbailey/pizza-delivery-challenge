package parser_test

import (
	"pizza-delivery/internal/pkg/parser"
	"reflect"
	"testing"
)

func TestParseDefault(t *testing.T) {
	input := ""
	parsedInput := parser.Parse(input)
	expectedInput := make([]string, 0)

	if len(parsedInput) != len(expectedInput) {
		t.Errorf("expected an array of %d items but instead found %d items", len(expectedInput), len(parsedInput))
	}
}

func TestParseWithNonEmptyString(t *testing.T) {
	input := "<>>v"
	parsedInput := parser.Parse(input)
	expectedInput := []string{"<", ">", ">", "v"}

	if !reflect.DeepEqual(parsedInput, expectedInput) {
		t.Error("unexpected parsed output")
	}
}
