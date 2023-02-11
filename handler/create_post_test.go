package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/iwashi623/go_posts/entity"
	"github.com/iwashi623/go_posts/store"
	"github.com/iwashi623/go_posts/testutil"
)

func TestCreatePost(t *testing.T) {
	t.Parallel()
	type want struct {
		status  int
		rspFile string
	}

	tests := map[string]struct {
		reqFile string
		want    want
	}{
		"ok": {
			reqFile: "testdata/create_post/ok_req.json.golden",
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/create_post/ok_rsp.json.golden",
			},
		},
		"badRequest": {
			reqFile: "testdata/create_post/bad_req.json.golden",
			want: want{
				status:  http.StatusBadRequest,
				rspFile: "testdata/create_post/bad_rsp.json.golden",
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/posts",
				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)),
			)

			sut := CreatePost{
				Store: &store.PostStore{
					Posts: map[entity.PostID]*entity.Post{},
				},
				Validator: validator.New(),
			}
			sut.ServeHTTP(w, r)

			res := w.Result()
			testutil.AssertResponse(t,
				res, tt.want.status, testutil.LoadFile(t, tt.want.rspFile),
			)
		})
	}
}
