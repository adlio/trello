// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
)

func mockResponse(paths ...string) *httptest.Server {
	parts := []string{".", "testdata"}
	filename := filepath.Join(append(parts, paths...)...)

	mockData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write(mockData)
	}))
}

func mockDynamicPathResponse() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		// Build the path for the dynamic request
		parts := []string{".", "testdata"}
		parts = append(parts, strings.Split(strings.TrimPrefix(r.URL.Path, "/"), "/")...)
		parts[len(parts)-1] = parts[len(parts)-1] + ".json"
		filename := filepath.Join(parts...)

		if _, err := os.Stat(filename); os.IsNotExist(err) {
			http.Error(rw, fmt.Sprintf("%s doesn't exist. Create it with the mock you'd like to use.", filename), http.StatusNotFound)
			return
		}

		mockData, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		rw.Write(mockData)

	}))
}

func mockErrorResponse(code int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		http.Error(rw, "An error occurred", code)
	}))
}
