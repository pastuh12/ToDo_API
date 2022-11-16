package main

import (
	"context"
	"testing"

	"github.com/todo_api/config"
	"github.com/todo_api/models"
	"github.com/todo_api/store"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	conf := config.Get()
	auth, _ := store.New(ctx, conf)

	testcases := []struct {
		in, want *models.User
	}{
		{
			&models.User{
				EmailAddress: "dfckjms@mail1.ru",
				Password:     "43821",
				Name:         "test",
			},
			&models.User{
				ID:           12,
				EmailAddress: "dfckjms@mail1.ru",
				Password:     "43821",
				Name:         "test",
			},
		},
		// {
		// 	&models.User{
		// 		EmailAddress: "dfb",
		// 		Password:     "",
		// 		Name:         "test",
		// 	},
		// 	nil,
		// },
		// {
		// 	nil,
		// 	nil,
		// },
	}
	for _, tc := range testcases {
		rev, _ := auth.Authtorization.CreateUser(ctx, tc.in)
		if rev != tc.want {
			t.Errorf("CreateUser: %q, want %q", rev, tc.want)
		}
	}
}
