package main

import (
	"fmt"
	"github.com/go-martini/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	recorder *httptest.ResponseRecorder
)

func RequestToRoot(method string, handler martini.Handler) {
	m := martini.Classic()
	m.Get("/", handler)
	recorder = httptest.NewRecorder()
	req, _ := http.NewRequest(method, "/", nil)
	m.ServeHTTP(recorder, req)
	fmt.Println(recorder)
}

func TestSampleAppGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SampleAppGo Suite")
}
