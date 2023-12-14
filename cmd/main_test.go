package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_readUserIP(t *testing.T) {
	type testcase struct {
		name string
		arg  *http.Request
		want string
	}
	cases := [...]testcase{
		{
			name: "X-Real-IP",
			arg: &http.Request{
				Header: map[string][]string{
					"X-Real-IP": {"100.100.1.1"},
				},
			},
			want: "100.100.1.1",
		},
		{
			name: "X-Forwarded-For",
			arg: &http.Request{
				Header: map[string][]string{
					"X-Forwarded-For": {"100.100.1.2"},
				},
			},
			want: "100.100.1.2",
		},
		{
			name: "remote addr",
			arg: &http.Request{
				RemoteAddr: "100.100.1.3",
			},
			want: "100.100.1.3",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, readUserIP(tt.arg))
		})
	}
}

func Test_formatResponse(t *testing.T) {
	t.Parallel()
	assert.Equal(t, []byte(`<html><body><p>Your IP address is 127.0.0.1</p></body></html>`), formatResponse("127.0.0.1"))
}
