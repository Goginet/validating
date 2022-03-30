package validating_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	v "github.com/webdeveloppro/validating"
)

func Test_JsonValidator_SuccessCase(t *testing.T) {
	// import "fmt"
	// import v "github.com/RussellLuo/validating"

	value := "{\"test_key\": 1}"

	validator := v.Json()

	errors := validator.Validate(v.Field{
		Name:     "meta",
		ValuePtr: &value,
	})

	require.Empty(t, errors)
}

func Test_JsonValidator_ErrorCase(t *testing.T) {
	// import "fmt"
	// import v "github.com/RussellLuo/validating"

	value := "{\"test_key\": }}}"

	validator := v.Json()

	errors := validator.Validate(v.Field{
		Name:     "meta",
		ValuePtr: &value,
	})

	require.NotEmpty(t, errors)
}
