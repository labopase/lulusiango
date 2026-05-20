package middlewares

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v5"
)

func TestRecover(t *testing.T) {
	tests := []struct {
		name          string
		config        RecoverConfig
		handler       echo.HandlerFunc
		expectPanic   bool
		wantErrType   interface{}
		wantErrorMsg  string
		wantStackPart string
	}{
		{
			name:   "should recover from panic and return wrapped PanicStackError",
			config: RecoverConfig{},
			handler: func(c *echo.Context) error {
				panic(errors.New("database connection lost"))
			},
			expectPanic:   false,
			wantErrType:   &PanicStackError{},
			wantErrorMsg:  "[PANIC RECOVER] database connection lost",
			wantStackPart: "recover.go",
		},
		{
			name: "should bypass recover when skipper returns true",
			config: RecoverConfig{
				Skipper: func(c *echo.Context) bool {
					return true
				},
			},
			handler: func(c *echo.Context) error {
				panic("skipped panic")
			},
			expectPanic: true,
		},
		{
			name: "should recover and wrap pure string panics",
			config: RecoverConfig{},
			handler: func(c *echo.Context) error {
				panic("runtime critical panic")
			},
			expectPanic:  false,
			wantErrType:  &PanicStackError{},
			wantErrorMsg: "[PANIC RECOVER] runtime critical panic",
		},
		{
			name: "should recover without stack trace if DisablePrintStack is set to true",
			config: RecoverConfig{
				DisablePrintStack: true,
			},
			handler: func(c *echo.Context) error {
				panic(errors.New("no stack error"))
			},
			expectPanic:  false,
			wantErrType:  errors.New(""),
			wantErrorMsg: "no stack error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			mw := Recover(tt.config)
			handler := mw(tt.handler)

			if tt.expectPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Error("expected panic to bubble up, but it was recovered")
					}
				}()
				_ = handler(c)
				return
			}

			err := handler(c)
			if err == nil {
				t.Fatal("expected an error to be returned, got nil")
			}

			switch tt.wantErrType.(type) {
			case *PanicStackError:
				pe, ok := err.(*PanicStackError)
				if !ok {
					t.Fatalf("expected error of type *PanicStackError, got %T", err)
				}
				if !strings.Contains(pe.Error(), tt.wantErrorMsg) {
					t.Errorf("expected error message to contain %q, got: %s", tt.wantErrorMsg, pe.Error())
				}
				if tt.wantStackPart != "" && !strings.Contains(string(pe.Stack), tt.wantStackPart) {
					t.Errorf("expected stack trace to contain %q, got: %s", tt.wantStackPart, string(pe.Stack))
				}
			default:
				if _, ok := err.(*PanicStackError); ok {
					t.Fatalf("expected error not to be of type *PanicStackError")
				}
				if !strings.Contains(err.Error(), tt.wantErrorMsg) {
					t.Errorf("expected error message to contain %q, got: %s", tt.wantErrorMsg, err.Error())
				}
			}
		})
	}
}
