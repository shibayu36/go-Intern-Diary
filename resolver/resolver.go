package resolver

import (
	"context"

	"github.com/hatena/go-Intern-Diary/model"
	"github.com/hatena/go-Intern-Diary/service"
)

type Resolver interface {
	Visitor(context.Context) (*userResolver, error)
	// ...
}

func newResolver(app service.DiaryApp) Resolver {
	return &resolver{app: app}
}

type resolver struct {
	app service.DiaryApp
}

func currentUser(ctx context.Context) *model.User {
	return ctx.Value("user").(*model.User)
}

func (r *resolver) Visitor(ctx context.Context) (*userResolver, error) {
	return &userResolver{currentUser(ctx)}, nil
}
