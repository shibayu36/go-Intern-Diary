package resolver

import (
	"context"
	"errors"

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

func currentUser(ctx context.Context) (*model.User, error) {
	user := ctx.Value("user").(*model.User)
	if user == nil {
		return nil, errors.New("visitor is not found")
	}
	return user, nil
}

func (r *resolver) Visitor(ctx context.Context) (*userResolver, error) {
	user, err := currentUser(ctx)
	return &userResolver{user, r.app}, err
}
