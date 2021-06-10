package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pyankovzhe/dictionary/internal/app/store/teststore"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestServer_Card(t *testing.T) {
	s := newServer(logrus.New(), teststore.New(), "fakeaddr")
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid with only original",
			payload: map[string]string{
				"original": "word",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "valid with translation",
			payload: map[string]string{
				"original":    "word",
				"translation": "слово",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "missing original",
			payload: map[string]string{
				"translation": "слово",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name:         "empty payload",
			payload:      map[string]string{},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/api/card", b)
			s.Handler.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
