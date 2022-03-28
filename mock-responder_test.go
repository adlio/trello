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
	"testing"
)

// MockResponder is a thin wrapper around the httptest.Server. It adds the
// ability to specify a file or directory of files to use as mock responses,
// and provides the AsertRequest method to add test assertions that requests
// are being made correctly. Just like with httptest.Server, the caller should
// defer a call to .Close() to shutdown the server when all requests complete.
// MockResponders should be created via the NewMockResponder() constructor.
//
type MockResponder interface {
	Close()
	URL() string
	AssertRequest(func(t *testing.T, r *http.Request))
}

type mockResponder struct {
	t *testing.T

	// server will be nil until .URL() is called the first time
	server *httptest.Server

	// requestAssertions is a list of functions which is called on
	// each HTTP request before finding and returning the mock response content.
	// They should be used to make assertions on the contents of the HTTP
	// request being made against this MockResponder
	requestAssertions []func(t *testing.T, r *http.Request)

	// mockPath holds the results of filepath.Join on the provided path parts.
	// The constructor verifies the existence of the path, so this will always
	// hold a valid path to either a mock file, or a directory of many mocks
	mockPath string

	// useDynamicPaths is set to true when mockPath is a directory. It triggers
	// code which determies the mock file from the path of incoming HTTP
	// requests
	useDynamicPaths bool
}

// NewMockResponder creates a new MockResponder instance around the provided
// test case. The mockPath is the relative filesystem path under ./testdata/
// where the mock response JSON can be found.
//
// If mockPath describes the path to a *file*, then that file
// will be used for ALL requests. If the path is a directory, then the mock
// response will be built dynamically from the path of the request (e.g.
// GET /subdir/folder/file will return the file at subdir/folder/file.json,
// assuming it exists). This latter mode is described as "dynamic paths". When
// requests arrive with querystring arguments, the dynamic path builder will
// compute an MD5 hash of the arguments and include that as a suffix of the
// mock file path.
//
// If no mockPath is provided, then the MockResponder will run in dynamic path
// mode from the root of the testdata/ directory.
//
// The caller is expected to defer a call .Close() after NewMockResponder().
//
func NewMockResponder(t *testing.T, mockPath ...string) MockResponder {
	r := &mockResponder{t: t}

	// Verify a valid path was provided
	r.mockPath = filepath.Join(append([]string{".", "testdata"}, mockPath...)...)
	fi, err := os.Stat(r.mockPath)
	if err != nil {
		log.Fatalf("invalid mock path %v: %s", mockPath, err)
	}

	// If the provided mockPath points to a directory, then
	// we'll figure out the ultimate path dynamically as requests occur.
	r.useDynamicPaths = fi.IsDir()

	return r
}

// AssertRequest adds a new function to be run on each HTTP request the mock
// responder recveives. Its intended use is to make test assertions on the
// content of the request
func (mr *mockResponder) AssertRequest(ra func(t *testing.T, r *http.Request)) {
	mr.requestAssertions = append(mr.requestAssertions, ra)
}

// Close wraps the *httptest.Server's Close method
func (mr *mockResponder) Close() {
	if mr.server != nil {
		mr.server.Close()
	}
}

// URL is equivalent to the *httptest.Server property of the same name, but it
// is responsible for *creating* the *httptest.Server. This function should
// be called after all customization (including calls to AssertRequest) is
// complete.
//
func (mr *mockResponder) URL() string {
	if mr.server != nil {
		mr.t.Error("URL() should only be called once, after completing configuration")
	}
	mr.server = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		for _, assertion := range mr.requestAssertions {
			assertion(mr.t, r)
		}
		mr.mockHandler(rw, r)
	}))
	return mr.server.URL
}

// mockHandler is the http.HandlerFunc for the httptest.Server inside the
// mockResponder. When the mockPath points to a single file, it simply returns
// that file in the HTTP response. Otherwise it dynamically determines the
// path of the mock file to use and returns that if the file is found...
// otherwise it responds with an error instructing the user where to put their
// mock file.
//
func (mr *mockResponder) mockHandler(rw http.ResponseWriter, r *http.Request) {
	var filename string
	if mr.useDynamicPaths {
		parts := []string{mr.mockPath}
		parts = append(parts, strings.Split(strings.TrimPrefix(r.URL.Path, "/"), "/")...)
		queryStringPart := strings.Replace(r.URL.RawQuery, "key=user&token=pass", "", -1)
		if queryStringPart != "" {
			parts[len(parts)-1] = fmt.Sprintf("%s-%x", parts[len(parts)-1], md5.Sum([]byte(queryStringPart)))
		}

		filename = filepath.Join(parts...)
		if !strings.HasSuffix(filename, ".json") {
			filename = filename + ".json"
		}
		if _, err := os.Stat(filename); err != nil {
			http.Error(rw, fmt.Sprintf("%s doesn't exist or couldn't be read. Create it with the mock you'd like to use.\n Args were: %s", filename, queryStringPart), http.StatusNotFound)
			return
		}
	} else {
		filename = mr.mockPath
	}

	mockData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	rw.Write(mockData)
}
