package services

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServices_Home(t *testing.T) {
	tests := []struct {
		name           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "#1",
			in:             httptest.NewRequest(http.MethodGet, "/", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "{\"message\":\"invalid request method\"}\n",
		},
		{
			name:           "#2",
			in:             httptest.NewRequest(http.MethodPost, "/", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "{\"message\":\"error reading request body\"}\n",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			s := New(nil)
			s.Home(test.out, test.in)
			if test.out.Code != test.expectedStatus {
				t.Logf("expected %d, got %d\n", test.expectedStatus, test.out.Code)
				t.Fail()
			}
			body := test.out.Body.String()
			if body != test.expectedBody {
				t.Logf("expected %s, got %s\n", test.expectedBody, body)
				t.Fail()
			}
		})
	}
}
