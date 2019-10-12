package resolver

import (
	"context"
	"errors"
	"strconv"

	"github.com/hatena/go-Intern-Diary/model"
	"github.com/hatena/go-Intern-Diary/service"
)

type Resolver interface {
	Visitor(context.Context) (*userResolver, error)
	User(context.Context, struct{ UserID string }) (*userResolver, error)
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
	return &userResolver{user}, err
}

func (r *resolver) User(ctx context.Context, args struct{ UserID string }) (*userResolver, error) {
	userID, err := strconv.ParseUint(args.UserID, 10, 64)
	if err != nil {
		return nil, err
	}
	user, err := r.app.FindUserByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return &userResolver{user}, nil
}
