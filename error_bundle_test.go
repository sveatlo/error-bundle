package error_bundle

import (
	"errors"
	"reflect"
	"testing"
)

func Test_errorsBundle_Evaluate(t *testing.T) {
	tests := []struct {
		name          string
		bundle        Bundle
		wantErr       bool
		wantErrString string
	}{
		{
			name:    "nil",
			bundle:  NewErrorsBundle(),
			wantErr: false,
		},
		{
			name:          "single error",
			bundle:        NewErrorsBundle(errors.New("error1")),
			wantErr:       true,
			wantErrString: "error1",
		},
		{
			name:          "multiple errors",
			bundle:        NewErrorsBundle(errors.New("error1"), errors.New("error2")),
			wantErr:       true,
			wantErrString: "error1 █ error2",
		},
		{
			name:    "empty with nil append",
			bundle:  NewErrorsBundle().Append(nil),
			wantErr: false,
		},
		{
			name:          "empty with errorous append",
			bundle:        NewErrorsBundle().Append(errors.New("error1")),
			wantErr:       true,
			wantErrString: "error1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.bundle.Evaluate()
			if (err != nil) != tt.wantErr {
				t.Errorf("errorsBundle.Evaluate() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got := tt.bundle.Error(); !reflect.DeepEqual(got, tt.wantErrString) {
				t.Errorf("NewErrorsBundle() = %v, want %v", got, tt.wantErrString)
			}

		})
	}
}
