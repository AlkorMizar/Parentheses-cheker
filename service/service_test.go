package service_test

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/AlkorMizar/Parentheses-cheker/service"
)

type respForTest struct {
	respAns       int
	respResLenght int
}
type reqForTest struct {
	methodType string
	url        string
}

func TestServiceResult(t *testing.T) {
	var requestWithParams = service.ServiceRoute + "?n="

	tests := map[string]struct {
		request reqForTest
		resp    respForTest
	}{
		"get request with correct length": {
			request: reqForTest{
				methodType: http.MethodGet,
				url:        requestWithParams + "8",
			},
			resp: respForTest{
				respAns:       http.StatusOK,
				respResLenght: 8,
			},
		},
		"get request with negative length": {
			request: reqForTest{
				methodType: http.MethodGet,
				url:        requestWithParams + "-10",
			},
			resp: respForTest{
				respAns:       http.StatusNotAcceptable,
				respResLenght: 0,
			},
		},
		"get request with not numbers": {
			request: reqForTest{
				methodType: http.MethodGet,
				url:        requestWithParams + "four",
			},
			resp: respForTest{
				respAns:       http.StatusNotAcceptable,
				respResLenght: 0,
			},
		},
		"get request without params": {
			request: reqForTest{
				methodType: http.MethodGet,
				url:        requestWithParams,
			},
			resp: respForTest{
				respAns:       http.StatusNotAcceptable,
				respResLenght: 0,
			},
		},
		"get request without name of parameter": {
			request: reqForTest{
				methodType: http.MethodGet,
				url:        service.ServiceRoute,
			},
			resp: respForTest{
				respAns:       http.StatusNotAcceptable,
				respResLenght: 0,
			},
		},
		"another type of request": {
			request: reqForTest{
				methodType: http.MethodPost,
				url:        requestWithParams + "8",
			},
			resp: respForTest{
				respAns:       http.StatusNotImplemented,
				respResLenght: 0,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			request, _ := http.NewRequest(tc.request.methodType, tc.request.url, http.NoBody)
			response := httptest.NewRecorder()

			service.GenrateHandler(response, request)

			if response.Code != tc.resp.respAns {
				t.Errorf("got %d, want %d", response.Code, tc.resp.respAns)
			}

			got := response.Body.String()
			re := regexp.MustCompile("[^(){}\\[\\]]+") //nolint:gosimple // this is the only way to create RegEx
			if len(got) != tc.resp.respResLenght || re.FindString(got) != "" {
				t.Errorf("got %s, want len %d", got, tc.resp.respResLenght)
			}
		})
	}
}
