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
	Diary(context.Context, struct{ DiaryID string }) (*diaryResolver, error)

	CreateDiary(context.Context, struct{ Name string }) (*diaryResolver, error)
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

func (r *resolver) Diary(ctx context.Context, args struct{ DiaryID string }) (*diaryResolver, error) {
	diaryID, err := strconv.ParseUint(args.DiaryID, 10, 64)
	if err != nil {
		return nil, err
	}
	diary, err := r.app.FindDiaryByID(diaryID)
	if err != nil {
		return nil, err
	}
	if diary == nil {
		return nil, errors.New("user not found")
	}
	return &diaryResolver{diary}, nil
}

func (r *resolver) CreateDiary(ctx context.Context, args struct{ Name string }) (*diaryResolver, error) {
	user, _ := currentUser(ctx)
	if user == nil {
		return nil, errors.New("user not found")
	}

	diary, err := r.app.CreateNewDiary(user.ID, args.Name)
	if err != nil {
		return nil, err
	}

	return &diaryResolver{diary}, nil
}
