package dataCollection

import (
	"fmt"
	"github.com/go-test/deep"
	"testing"
)

func TestGetBlock(t *testing.T) {
	for _, tc := range testCasesGetBlock {
		description := fmt.Sprintf("Test:%s, GetBlock(%d,%d), ",
			tc.description, tc.blockNumber, tc.requestTimeout)

		gotStatus, gotHeader, gotBody, gotErr := GetBlock(tc.blockNumber, tc.requestTimeout)

		switch {
		case tc.expectedError != nil && gotErr == nil:
			t.Error(description + "expected error")
		case tc.expectedError == nil && gotErr != nil:
			t.Errorf(description+"unexpected error \n%s", gotErr.Error())
		default:
			gotResponse := expectedRetrieveTransaction{Status: gotStatus, Header: gotHeader, Body: gotBody}
			// removing delete because cannot be tested against live api
			delete(gotHeader, "Date")
			if diffList := deep.Equal(tc.expected, gotResponse); len(diffList) > 0 {
				t.Errorf(description+"\nDiff    : %v\n", diffList)
			}
		}
	}
}

func TestGetTransaction(t *testing.T) {
	for _, tc := range testCasesGetTransaction {
		description := fmt.Sprintf("Test:%s, GetTransaction(%d,%d,%d), ",
			tc.description, tc.blockNumber, tc.index, tc.requestTimeout)

		gotStatus, gotHeader, gotBody, gotErr := GetTransaction(tc.blockNumber, tc.index, tc.requestTimeout)

		switch {
		case tc.expectedError != nil && gotErr == nil:
			t.Error(description + "expected error")
		case tc.expectedError == nil && gotErr != nil:
			t.Errorf(description+"unexpected error \n%s", gotErr.Error())

		default:
			gotResponse := expectedRetrieveTransaction{Status: gotStatus, Header: gotHeader, Body: gotBody}
			// removing delete because cannot be tested against live apiz
			delete(gotHeader, "Date")
			if diffList := deep.Equal(tc.expected, gotResponse); len(diffList) > 0 {
				t.Errorf(description+"\nDiff    : %v\n", diffList)
			}
		}
	}
}

func TestGetLastBlockNumber(t *testing.T) {
	for _, tc := range testCasesGetLastBlock {
		description := fmt.Sprintf("Test:%s, GetLastBlockNumber(%d), ",
			tc.description, tc.requestTimeout)

		gotLastBlock, gotErr := GetLastBlockNumber(tc.requestTimeout)

		switch {
		case tc.expectedError != nil && gotErr == nil:
			t.Error(description + "expected error")
		case tc.expectedError == nil && gotErr != nil:
			t.Errorf(description+"unexpected error \n%s", gotErr.Error())
		default:
			if tc.expectedBlock > gotLastBlock {
				t.Errorf("Expected block bigger than %d\n Got:", tc.expectedBlock)
			}
		}
	}
}