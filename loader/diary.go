package loader

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"github.com/hatena/go-Intern-Diary/model"
	"github.com/hatena/go-Intern-Diary/service"
)

const diaryLoaderKey = "diaryLoader"

type diaryIDKey struct {
	id uint64
}

func (key diaryIDKey) String() string {
	return fmt.Sprint(key.id)
}

func (key diaryIDKey) Raw() interface{} {
	return key.id
}

func newDiaryLoader(app service.DiaryApp) dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		results := make([]*dataloader.Result, len(keys))
		userIDs := make([]uint64, 0, len(keys))

		for _, key := range keys {
			switch key := key.(type) {
			case userIDKey:
				userIDs = append(userIDs, key.id)
			}
		}

		diariesByUserIDs, _ := app.ListDiariesByUserIDs(userIDs)

		for i, key := range keys {
			results[i] = &dataloader.Result{Data: nil, Error: nil}
			switch key := key.(type) {
			case userIDKey:
				results[i].Data = diariesByUserIDs[key.id]
			}
		}
		return results
	}
}

func LoadDiariesByUserID(ctx context.Context, id uint64) ([]*model.Diary, error) {
	ldr, err := getLoader(ctx, diaryLoaderKey)
	if err != nil {
		return nil, err
	}
	data, err := ldr.Load(ctx, userIDKey{id: id})()
	if err != nil {
		return nil, err
	}
	return data.([]*model.Diary), nil
}
