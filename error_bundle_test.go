package error_bundle

import (
	"reflect"
	"testing"
)

func TestNewErrorsBundle(t *testing.T) {
	type args struct {
		errs []error
	}
	tests := []struct {
		name string
		args args
		want Bundle
	}{
		{
			name: "nil",
			args: args{
				errs: nil,
			},
			want: nil,
		},
		{
			name: "empty",
			args: args{
				errs: []error{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewErrorsBundle(tt.args.errs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewErrorsBundle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorsBundle_Append(t *testing.T) {
	type fields struct {
		errors          []error
		StringFormatter ErrorsBundleFormatFunc
	}
	type args struct {
		e error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Bundle
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &errorsBundle{
				errors:          tt.fields.errors,
				StringFormatter: tt.fields.StringFormatter,
			}
			if got := err.Append(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errorsBundle.Append() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorsBundle_Evaluate(t *testing.T) {
	type fields struct {
		errors          []error
		StringFormatter ErrorsBundleFormatFunc
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &errorsBundle{
				errors:          tt.fields.errors,
				StringFormatter: tt.fields.StringFormatter,
			}
			if err := err.Evaluate(); (err != nil) != tt.wantErr {
				t.Errorf("errorsBundle.Evaluate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_errorsBundle_Errors(t *testing.T) {
	type fields struct {
		errors          []error
		StringFormatter ErrorsBundleFormatFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   []error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &errorsBundle{
				errors:          tt.fields.errors,
				StringFormatter: tt.fields.StringFormatter,
			}
			if got := err.Errors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errorsBundle.Errors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorsBundle_Len(t *testing.T) {
	type fields struct {
		errors          []error
		StringFormatter ErrorsBundleFormatFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := errorsBundle{
				errors:          tt.fields.errors,
				StringFormatter: tt.fields.StringFormatter,
			}
			if got := err.Len(); got != tt.want {
				t.Errorf("errorsBundle.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorsBundle_Swap(t *testing.T) {
	type fields struct {
		errors          []error
		StringFormatter ErrorsBundleFormatFunc
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := errorsBundle{
				errors:          tt.fields.errors,
				StringFormatter: tt.fields.StringFormatter,
			}
			err.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_errorsBundle_Less(t *testing.T) {
	type fields struct {
		errors          []error
		StringFormatter ErrorsBundleFormatFunc
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := errorsBundle{
				errors:          tt.fields.errors,
				StringFormatter: tt.fields.StringFormatter,
			}
			if got := err.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("errorsBundle.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorsBundle_Error(t *testing.T) {
	type fields struct {
		errors          []error
		StringFormatter ErrorsBundleFormatFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := errorsBundle{
				errors:          tt.fields.errors,
				StringFormatter: tt.fields.StringFormatter,
			}
			if got := err.Error(); got != tt.want {
				t.Errorf("errorsBundle.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorsBundle_MarshalJSON(t *testing.T) {
	type fields struct {
		errors          []error
		StringFormatter ErrorsBundleFormatFunc
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &errorsBundle{
				errors:          tt.fields.errors,
				StringFormatter: tt.fields.StringFormatter,
			}
			got, err := err.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("errorsBundle.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errorsBundle.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_defaultFormatter(t *testing.T) {
	type args struct {
		err *errorsBundle
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defaultFormatter(tt.args.err); got != tt.want {
				t.Errorf("defaultFormatter() = %v, want %v", got, tt.want)
			}
		})
	}
}
