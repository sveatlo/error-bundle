package error_bundle

import (
	"encoding/json"
)

type Bundle interface {
	error

	Append(error) Bundle
	Evaluate() error
	Errors() []error
}

type errorsBundle struct {
	errors          []error `json:"errors"`
	StringFormatter ErrorsBundleFormatFunc
}

func NewErrorsBundle(errs ...error) Bundle {
	bundle := &errorsBundle{
		errors:          errs,
		StringFormatter: defaultFormatter,
	}

	return bundle
}

func (b *errorsBundle) Append(e error) Bundle {
	if b == nil {
		return nil
	}
	if b == e {
		return b
	}
	if b == nil {
		return NewErrorsBundle(e)
	}

	if bundle, ok := e.(*errorsBundle); ok {
		b.errors = append(b.errors, bundle.errors...)
		return b
	} else if bundle, ok := e.(errorsBundle); ok {
		b.errors = append(b.errors, bundle.errors...)
		return b
	}

	b.errors = append(b.errors, e)
	return b
}

func (b *errorsBundle) Evaluate() error {
	if b == nil || len(b.errors) == 0 {
		return nil
	}

	return b
}

func (b *errorsBundle) Errors() []error {
	if b == nil {
		return []error{}
	}

	return b.errors
}

func (b errorsBundle) Len() int {
	return len(b.errors)
}

func (b errorsBundle) Swap(i, j int) {
	b.errors[i], b.errors[j] = b.errors[j], b.errors[i]
}

func (b errorsBundle) Less(i, j int) bool {
	return b.errors[i].Error() < b.errors[j].Error()
}

func (b errorsBundle) Error() string {
	return b.StringFormatter(&b)
}

func (b *errorsBundle) MarshalJSON() ([]byte, error) {
	errStrings := []string{}
	for _, err := range b.errors {
		errStrings = append(errStrings, err.Error())
	}

	return json.Marshal(errStrings)
}

type ErrorsBundleFormatFunc func(Bundle) string

func defaultFormatter(b Bundle) string {
	res := ""
	delimiter := " █ "

	for _, e := range b.Errors() {
		if e == nil {
			continue
		}

		res += e.Error() + delimiter
	}

	if l := len(res); l > 0 {
		res = res[:l-len(delimiter)]
	}

	return res
}
