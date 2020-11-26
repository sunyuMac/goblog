package tests

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllPages(t *testing.T) {
	baseUrl := "http://localhost:3000"

	tests := []struct {
		method   string
		url      string
		expected int
	}{
		{"GET", "/", 200},
		{"GET", "/about", 200},
		{"GET", "/notfound", 404},
		{"GET", "/articles", 200},
		{"GET", "/articles/create", 200},
		{"GET", "/articles/3", 200},
		{"GET", "/articles/3/edit", 200},
		{"POST", "/articles/3", 200},
		{"POST", "/articles", 200},
		{"POST", "/articles/1/delete", 404},
	}

	for _, test := range tests {
		t.Logf("当前请求URL: %v \n", test.url)

		var (
			resp *http.Response
			err  error
		)

		switch test.method {
		case "POST":
			data := make(map[string][]string)
			resp, err = http.PostForm(baseUrl+test.url, data)
		default:
			resp, err = http.Get(baseUrl + test.url)
		}

		assert.NoError(t, err, "请求"+test.url+"时报错")
		assert.Equal(t, test.expected, resp.StatusCode, test.url+"应该返回状态码"+strconv.Itoa(test.expected))

	}

}
