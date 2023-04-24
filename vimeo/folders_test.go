package vimeo

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestFoldersService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/folders", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormURLValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"name": "Test"}]}`)
	})

	categories, _, err := client.Folders.List(OptPage(1), OptPerPage(2))
	if err != nil {
		t.Errorf("Folders.List returned unexpected error: %v", err)
	}

	want := []*Category{{Name: "Test"}}
	if !reflect.DeepEqual(categories, want) {
		t.Errorf("Folders.List returned %+v, want %+v", categories, want)
	}
}
