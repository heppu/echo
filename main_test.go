package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEcho_SinglePartPath(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/hello", nil)
	echo(w, r)

	resp := readResponse(t, w)
	assertEqual(t, "hello", resp)
}

func TestEcho_MultiplePartPath(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/hello/world", nil)
	echo(w, r)

	resp := readResponse(t, w)
	assertEqual(t, "hello world", resp)
}

func TestEcho_EmptyPath(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	echo(w, r)

	resp := readResponse(t, w)
	assertEqual(t, "", resp)
}

func TestEcho_FormatUpperCase(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/hello?format=upper", nil)
	echo(w, r)

	resp := readResponse(t, w)
	assertEqual(t, "HELLO", resp)
}

func readResponse(t testing.TB, w *httptest.ResponseRecorder) string {
	resp, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatal(err)
	}
	return string(resp)
}

func assertEqual(t testing.TB, expected, got string) {
	if expected != got {
		t.Fatalf("\n\texpected: %#v\n\tgot:      %#v", expected, got)
	}
}
