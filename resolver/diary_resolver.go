package resolver

import (
	"context"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/hatena/go-Intern-Diary/model"
	"github.com/hatena/go-Intern-Diary/service"
)

type diaryResolver struct {
	diary *model.Diary
	app   service.DiaryApp
}

func (d *diaryResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(fmt.Sprint(d.diary.ID))
}

func (d *diaryResolver) User(ctx context.Context) (*userResolver, error) {
	user, err := d.app.FindUserByID(d.diary.UserID)
	if err != nil {
		return nil, err
	}
	return &userResolver{user, d.app}, nil
}

func (d *diaryResolver) Name(ctx context.Context) string {
	return d.diary.Name
}
