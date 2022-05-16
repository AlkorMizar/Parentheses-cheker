package service_test

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/AlkorMizar/Parentheses-cheker/service"
	"github.com/AlkorMizar/Parentheses-cheker/service/usecases"
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
				respAns:       http.StatusBadRequest,
				respResLenght: 0,
			},
		},
		"get request with not numbers": {
			request: reqForTest{
				methodType: http.MethodGet,
				url:        requestWithParams + "four",
			},
			resp: respForTest{
				respAns:       http.StatusBadRequest,
				respResLenght: 0,
			},
		},
		"get request without params": {
			request: reqForTest{
				methodType: http.MethodGet,
				url:        requestWithParams,
			},
			resp: respForTest{
				respAns:       http.StatusBadRequest,
				respResLenght: 0,
			},
		},
		"get request without name of parameter": {
			request: reqForTest{
				methodType: http.MethodGet,
				url:        service.ServiceRoute,
			},
			resp: respForTest{
				respAns:       http.StatusBadRequest,
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

			m := &mock{}
			h := service.NewHandlers(m)
			h.ServeHTTP(response, request)

			if response.Code != tc.resp.respAns {
				t.Errorf("got %d, want %d", response.Code, tc.resp.respAns)
			}
		})
	}
}

type mock struct {
}

func (m *mock) Generate(leng int) string {
	return "()[]{}"
}

func TestGenerateResult(t *testing.T) {
	tests := map[string]struct {
		lenIn  int
		lenOut int
	}{
		"correct length": {
			lenIn:  8,
			lenOut: 8,
		},
		"odd lengrh": {
			lenIn:  7,
			lenOut: 7,
		},
		"big length": {
			lenIn:  100,
			lenOut: 100,
		},
		"zero": {
			lenIn:  0,
			lenOut: 0,
		},
	}

	for name, tc := range tests {
		generator := usecases.NewBraces()

		t.Run(name, func(t *testing.T) {
			got := generator.Generate(tc.lenIn)
			re := regexp.MustCompile("[^(){}\\[\\]]+") //nolint:gosimple // this is the only way to create RegEx

			if len(got) != tc.lenOut || re.FindString(got) != "" {
				t.Errorf("got %s, want len %d", got, tc.lenOut)
			}
		})
	}
}
