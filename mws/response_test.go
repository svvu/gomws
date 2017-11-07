package mws

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResponse_ResultParser(t *testing.T) {
	Convey("when create ResultParser fail", t, func() {
		resp := NewResponse(&http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString("")),
		})

		rp, err := resp.ResultParser()

		Convey("return nil ResultParser", func() {
			So(rp, ShouldBeNil)
		})

		Convey("return error", func() {
			So(err, ShouldNotBeNil)
		})
	})

	Convey("when create ResultParser success", t, func() {
		resp := NewResponse(&http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString("<foo>bar</foo>")),
		})

		rp, err := resp.ResultParser()

		Convey("return created ResultParser", func() {
			So(rp, ShouldNotBeNil)
		})

		Convey("return no error", func() {
			So(err, ShouldBeNil)
		})
	})
}

func TestResponse_WriteBodyTo(t *testing.T) {
	resp := NewResponse(&http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString("foo bar")),
	})
	out := bytes.NewBufferString("")

	resp.WriteBodyTo(out)

	Convey("Data in the response body be written to buffer", t, func() {
		So(out.String(), ShouldEqual, "foo bar")
	})
}

func TestParseResponseError(t *testing.T) {
	// mwsErrorBody := ioutil.NopCloser(bytes.NewBufferString("<Error><Message>foo</Message></Error>"))

	testCases := []struct {
		desc       string
		input      *http.Response
		outputDesc string
		output     error
	}{
		{
			desc:       "when status code is 200",
			outputDesc: "no error return",
			input:      &http.Response{StatusCode: 200},
			output:     nil,
		},
		{
			desc:       "when status code not match",
			outputDesc: "error with status return",
			input: &http.Response{
				StatusCode: 300,
				Status:     "foo",
				Body:       ioutil.NopCloser(bytes.NewBufferString("<a>b</a>")),
			},
			output: fmt.Errorf("Request not success. Reason: foo"),
		},
		{
			desc:       "when status code is 400 and has no mws error",
			outputDesc: "error with status return",
			input: &http.Response{
				StatusCode: 400,
				Status:     "foo",
				Body:       ioutil.NopCloser(bytes.NewBufferString("<a>b</a>")),
			},
			output: fmt.Errorf("Request not success. Reason: foo"),
		},
		{
			desc:       "when status code is 401 and has no mws error",
			outputDesc: "error with status return",
			input: &http.Response{
				StatusCode: 401,
				Status:     "foo",
				Body:       ioutil.NopCloser(bytes.NewBufferString("<a>b</a>")),
			},
			output: fmt.Errorf("Request not success. Reason: foo"),
		},
		{
			desc:       "when status code is 403 and has no mws error",
			outputDesc: "error with status return",
			input: &http.Response{
				StatusCode: 403,
				Status:     "foo",
				Body:       ioutil.NopCloser(bytes.NewBufferString("<a>b</a>")),
			},
			output: fmt.Errorf("Request not success. Reason: foo"),
		},
		{
			desc:       "when status code is 404 and has no mws error",
			outputDesc: "error with status return",
			input: &http.Response{
				StatusCode: 404,
				Status:     "foo",
				Body:       ioutil.NopCloser(bytes.NewBufferString("<a>b</a>")),
			},
			output: fmt.Errorf("Request not success. Reason: foo"),
		},
		{
			desc:       "when status code is 500 and has no mws error",
			outputDesc: "error with status return",
			input: &http.Response{
				StatusCode: 500,
				Status:     "foo",
				Body:       ioutil.NopCloser(bytes.NewBufferString("<a>b</a>")),
			},
			output: fmt.Errorf("Request not success. Reason: foo"),
		},
		{
			desc:       "when status code is 503 and has no mws error",
			outputDesc: "error with status return",
			input: &http.Response{
				StatusCode: 503,
				Status:     "foo",
				Body:       ioutil.NopCloser(bytes.NewBufferString("<a>b</a>")),
			},
			output: fmt.Errorf("Request not success. Reason: foo"),
		},
		{
			desc:       "when status code is 400 and has mws error",
			outputDesc: "mws errors return",
			input: &http.Response{
				StatusCode: 400,
				Status:     "bar",
				Body:       ioutil.NopCloser(bytes.NewBufferString("<Error><Message>foo</Message></Error>")),
			},
			output: fmt.Errorf("Request not success. Reason: bar: foo"),
		},
		{
			desc:       "when status code is 401 and has mws error",
			outputDesc: "mws errors return",
			input: &http.Response{
				StatusCode: 401,
				Status:     "bar",
				Body:       ioutil.NopCloser(bytes.NewBufferString("<Error><Message>foo</Message></Error>")),
			},
			output: fmt.Errorf("Request not success. Reason: bar: foo"),
		},
		{
			desc:       "when status code is 403 and has mws error",
			outputDesc: "mws errors return",
			input: &http.Response{
				StatusCode: 403,
				Status:     "bar",
				Body:       ioutil.NopCloser(bytes.NewBufferString("<Error><Message>foo</Message></Error>")),
			},
			output: fmt.Errorf("Request not success. Reason: bar: foo"),
		},
		{
			desc:       "when status code is 404 and has mws error",
			outputDesc: "mws errors return",
			input: &http.Response{
				StatusCode: 404,
				Status:     "bar",
				Body:       ioutil.NopCloser(bytes.NewBufferString("<Error><Message>foo</Message></Error>")),
			},
			output: fmt.Errorf("Request not success. Reason: bar: foo"),
		},
		{
			desc:       "when status code is 500 and has mws error",
			outputDesc: "mws errors return",
			input: &http.Response{
				StatusCode: 500,
				Status:     "bar",
				Body:       ioutil.NopCloser(bytes.NewBufferString("<Error><Message>foo</Message></Error>")),
			},
			output: fmt.Errorf("Request not success. Reason: bar: foo"),
		},
		{
			desc:       "when status code is 503 and has mws error",
			outputDesc: "mws errors return",
			input: &http.Response{
				StatusCode: 503,
				Status:     "bar",
				Body:       ioutil.NopCloser(bytes.NewBufferString("<Error><Message>foo</Message></Error>")),
			},
			output: fmt.Errorf("Request not success. Reason: bar: foo"),
		},
		{
			desc:       "when response has more than one mws error",
			outputDesc: "mws errors return",
			input: &http.Response{
				StatusCode: 503,
				Status:     "bar",
				Body: ioutil.NopCloser(bytes.NewBufferString(
					"<Response><Error><Message>foo1</Message></Error>" +
						"<Error><Message>foo2</Message></Error></Response>",
				)),
			},
			output: fmt.Errorf("Request not success. Reason: bar: foo1\nfoo2"),
		},
	}

	for _, testCase := range testCases {
		Convey(testCase.desc, t, func() {
			resp := &Response{Response: testCase.input}
			err := parseResponseError(resp)

			Convey(testCase.outputDesc, func() {
				if testCase.output != nil {
					So(err.Error(), ShouldEqual, testCase.output.Error())
				} else {
					So(err, ShouldBeNil)
				}
			})
		})
	}

}
