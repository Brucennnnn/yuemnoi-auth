package di

import (
	"context"
	"fmt"
	"reflect"

	"github.com/sds-2/config"
	"github.com/sds-2/db"
	"github.com/sds-2/feature/items"
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
	_ = db.InitPostgreSQL(cfg)

	// domain
	itemsDomain := items.NewItemsDomain()

	// handler
	itemsHandler := items.NewItemsHandler(itemsDomain)

	r = route.NewHandler(itemsHandler)

	return r, nil
}
