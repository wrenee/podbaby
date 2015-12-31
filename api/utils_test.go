package api

import (
	"github.com/danjac/podbaby/models"
	"net/http"
	"testing"
)

func mockGetVars(vars map[string]string) varsGetter {
	return func(r *http.Request) map[string]string {
		return vars
	}
}

func mockGetContext(ctx map[string]interface{}) contextGetter {
	return func(r *http.Request, key string) (interface{}, bool) {
		result, ok := ctx[key]
		return result, ok
	}
}

func TestGetUserIfOK(t *testing.T) {
	user := &models.User{}
	ctx := map[string]interface{}{
		userKey: user,
	}
	getContext = mockGetContext(ctx)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	result, ok := getUser(req)
	if result == nil {
		t.Fail()
	}

	if !ok {
		t.Fail()
	}
}

func TestGetUserIfNone(t *testing.T) {
	ctx := make(map[string]interface{})
	getContext = mockGetContext(ctx)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	result, ok := getUser(req)
	if result != nil {
		t.Fail()
	}

	if ok {
		t.Fail()
	}
}

func TestGetPageIfOK(t *testing.T) {

	url := "/latest/?page=10"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	page := getPage(req)
	if page != 10 {
		t.Fail()
	}
}

func TestGetPageIfNone(t *testing.T) {

	url := "/latest/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	page := getPage(req)
	if page != 1 {
		t.Fail()
	}
}

func TestGetPageIfNotInt(t *testing.T) {

	url := "/latest/?page=foo"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	page := getPage(req)
	if page != 1 {
		t.Fail()
	}
}

func TestGetInt64IfOk(t *testing.T) {

	vars := map[string]string{
		"id": "11",
	}
	getVars = mockGetVars(vars)

	url := "/channels/11/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	result, err := getInt64(req, "id")
	if err != nil {
		t.Fatal(err)
	}

	if result != 11 {
		t.Errorf("Invalid parameter:%d", result)
	}

}

func TestGetInt64IfEmpty(t *testing.T) {

	vars := make(map[string]string)
	getVars = mockGetVars(vars)

	url := "/channels/11/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = getInt64(req, "id")
	if err == nil {
		t.Fail()
	}

}

func TestGetInt64IfNotInt(t *testing.T) {
	vars := map[string]string{
		"id": "foo",
	}

	getVars = mockGetVars(vars)

	url := "/channels/foo/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = getInt64(req, "id")
	if err == nil {
		t.Fail()
	}

}