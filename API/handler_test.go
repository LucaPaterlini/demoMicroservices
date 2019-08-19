package API

import (
	"bytes"
	"fmt"
	"github.com/go-test/deep"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func testHandler(t *testing.T, f func(http.ResponseWriter, *http.Request), testCases []handlerTest) {
	for _, tc := range testCases {
		description := fmt.Sprintf("Test:%s, request=%v", tc.description, tc.requestPath)
		// initialization set the last block
		var err error

		req, err := http.NewRequest("GET", tc.requestPath, nil)
		if err != nil {
			t.Error(description + "\nrequest initialization error")
		}
		// initialize response and router
		gotW := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc(tc.requestPathSignature, f)
		handlerTime := http.TimeoutHandler(router, tc.requestTimeout, "service not available try later")
		handlerTime.ServeHTTP(gotW, req)

		// checking the response wrapper
		if diffList := deep.Equal(tc.expectedW, gotW); len(diffList) > 0 {
			t.Errorf(description+"\nDiff    : %v\n", diffList)
		}
		// checking the body of the response
		var gotBodyBytes []byte
		if gotBodyBytes, err = ioutil.ReadAll(gotW.Body); err != nil {
			t.Errorf(description + "\n error while ready response body")
			continue
		}
		if !bytes.HasPrefix(gotBodyBytes, tc.expectedBodyBytes) {
			t.Errorf(description+"\nExpected: %s\nGot     : %s", tc.expectedBodyBytes, gotBodyBytes)
		}
	}
}

func TestUpdateRoutine(t *testing.T) {
	// test the execution, it does not have any check because
	// the inner function its already tested ,
	// the main purpose is to be sure its not giving any panic.
	UpdateRoutine(time.Second, time.Nanosecond)
}

func TestGetBlockHandler(t *testing.T) {
	testHandler(t, GetBlockHandler, testCasesGetBlockHandler)
	// wait to allow the go routine that do the GetBlockHandler to be executed inside the ticker loop

}

func TestGetTransactionHandler(t *testing.T) {
	testHandler(t, GetTransactionHandler, testCasesGetTransaction)
}
