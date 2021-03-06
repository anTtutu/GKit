package egroup

import (
	"Songzhibin/GKit/goroutine"
	"context"
	"fmt"
	"net/http"
	"testing"
)

var _admin = NewLifeAdmin()

func TestLifeAdmin_Start(t *testing.T) {
	srv := &http.Server{
		Addr: ":8080",
	}
	_admin.Add(Member{
		Start: func(ctx context.Context) error {
			t.Log("http start")
			return goroutine.Delegate(ctx, -1, func(ctx context.Context) error {
				return srv.ListenAndServe()
			})
		},
		Shutdown: func(ctx context.Context) error {
			t.Log("http shutdown")
			return srv.Shutdown(context.Background())
		},
	})
	fmt.Println("error", _admin.Start())
	defer _admin.shutdown()
}
