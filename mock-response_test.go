// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"crypto/md5"
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
		queryStringPart := strings.Replace(r.URL.RawQuery, "key=user&token=pass", "", -1)
		if queryStringPart != "" {
			parts[len(parts)-1] = fmt.Sprintf("%s-%x", parts[len(parts)-1], md5.Sum([]byte(queryStringPart)))
		}
		parts[len(parts)-1] = parts[len(parts)-1] + ".json"

		filename := filepath.Join(parts...)

		if _, err := os.Stat(filename); os.IsNotExist(err) {
			http.Error(rw, fmt.Sprintf("%s doesn't exist. Create it with the mock you'd like to use.\n Args were: %s", filename, queryStringPart), http.StatusNotFound)
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
