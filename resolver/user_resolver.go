package resolver

import (
	"context"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/hatena/go-Intern-Diary/model"
)

type userResolver struct {
	user *model.User
}

func (u *userResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(fmt.Sprint(u.user.ID))
}

func (u *userResolver) Name(ctx context.Context) string {
	return u.user.Name
}

// func (u *userResolver) Diaries(ctx context.Context) string {
// 	return u.user.Name
// }
