package resolver

import (
	"context"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/hatena/go-Intern-Diary/loader"
	"github.com/hatena/go-Intern-Diary/model"
)

type diaryResolver struct {
	diary *model.Diary
}

func (d *diaryResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(fmt.Sprint(d.diary.ID))
}

func (d *diaryResolver) User(ctx context.Context) (*userResolver, error) {
	user, err := loader.LoadUser(ctx, d.diary.UserID)
	if err != nil {
		return nil, err
	}
	return &userResolver{user}, nil
}

func (d *diaryResolver) Name(ctx context.Context) string {
	return d.diary.Name
}
