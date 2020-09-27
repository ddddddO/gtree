package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_indexHandler(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		path       string
		wantStatus int
	}{
		{
			name:       "成功",
			method:     http.MethodGet,
			path:       "http://shost/",
			wantStatus: http.StatusOK,
		},
		{
			name:       "パスが間違い",
			method:     http.MethodGet,
			path:       "http://shost/aaaa",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "メソッドがGET以外(POST)",
			method:     http.MethodPost,
			path:       "http://shost/",
			wantStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(tt.method, tt.path, nil)
			res := httptest.NewRecorder()

			indexHandler(res, req)

			if res.Code != tt.wantStatus {
				t.Errorf("want %d, got %d", tt.wantStatus, res.Code)
			}
		})
	}
}

// FIXME: redis
// func Test_shortenedurlsHandler(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodPost, "http://shost/surls/", nil)
// 	v := url.Values{}
// 	v.Set("url", "https://golang.org/pkg/bytes/")
// 	req.PostForm = v
// 	res := httptest.NewRecorder()

// 	shortenedurlsHandler(res, req)

// 	if res.Code != http.StatusCreated {
// 		t.Errorf("want %d, got %d", http.StatusCreated, res.Code)
// 	}
// }
