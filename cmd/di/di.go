package di

import (
	"context"
	"fmt"
	"reflect"

	"github.com/sds-2/config"
	"github.com/sds-2/db"
	"github.com/sds-2/feature/item"
	"github.com/sds-2/feature/review"
	"github.com/sds-2/route"
)

func must[T any](t T, err error) T {
	if err != nil {
		typeName := reflect.TypeOf(t).String()
		err := fmt.Errorf("failed to initialize %s: %w", typeName, err)
		panic(err)
	}
	return t
}

func InitDI(ctx context.Context, cfg *config.Config) (r *route.Handler, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	// db
	dbPG := db.InitPostgreSQL(cfg)

	// repository
	itemRepository := item.NewitemRepository(dbPG)
	reviewRepository := review.NewReviewRepository(dbPG)

	// domain
	itemDomain := item.NewitemDomain(itemRepository)
	reviewDomain := review.NewReviewDomain(reviewRepository)

	// handler
	itemHandler := item.NewItemHandler(itemDomain)
	reviewHandler := review.NewReviewHandler(reviewDomain)

	r = route.NewHandler(itemHandler,reviewHandler)

	return r, nil
}
