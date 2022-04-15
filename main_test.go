package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	testCase := []struct {
		Name     string `json:"-"`
		Login    string `json:"login"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Want     int    `json:"-"`
	}{
		{
			Name:     "Test1",
			Login:    "NameName",
			Password: "Password123",
			Email:    "login@gmail.com",
			Want:     200,
		},
	}

	handler := http.HandlerFunc(SignUpHandler)

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			data, err := json.Marshal(tc)
			if err != nil {
				log.Printf("json error: %v", err.Error())
			}

			readerData := bytes.NewReader(data)
			recorder := httptest.NewRecorder()
			req, err := http.NewRequest("POST", "/signup", readerData)
			if err != nil {
				log.Printf("requst error: %v", err.Error())
			}

			handler.ServeHTTP(recorder, req)
			assert.Equal(t, tc.Want, recorder.Code)
		})
	}
}
