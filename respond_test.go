package respond_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	is "github.com/cheekybits/is"
	respond "github.com/gmdmgithub/tdd-blueprint"
)

func TestAll(t *testing.T) {
	is := is.New(t)

	var tests = []struct {
		Body       string //body string
		Fn         func(w http.ResponseWriter, r *http.Request)
		StatusCode int
	}{
		{
			Fn: func(w http.ResponseWriter, r *http.Request) {
				respond.PerformRequest(w, r, http.StatusOK, "Happy answer!")
			},
			StatusCode: 200,
			Body:       `"Happy answer!"`,
		},
	}

	//best solution to test http is using httptest package
	w := httptest.NewRecorder()

	for _, test := range tests {
		test.Fn(w, nil)

		// asserssion
		is.Equal(test.StatusCode, w.Code)
		is.Equal(test.Body, w.Body.String())

	}
}
