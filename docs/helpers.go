package api_service

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

func ValidateRequest(r *http.Request, v *validator.Validate) error {
	if err := v.Struct(r.PostForm); err != nil {
		return err
	}
	return nil
}

func Paginate(ctx context.Context, db *gorm.DB, page, perPage int) (*gorm.DB, error) {
	if page < 1 {
		return nil, fmt.Errorf("page must be greater than 0")
	}
	if perPage < 1 {
		return nil, fmt.Errorf("per_page must be greater than 0")
	}

	var offset int
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * perPage
	}

	return db.Offset(offset).Limit(perPage), nil
}

func IntSliceQuery(args []string) ([]interface{}, error) {
	results := make([]interface{}, len(args))
	for i, arg := range args {
		val, err := strconv.Atoi(arg)
		if err != nil {
			return nil, err
		}
		results[i] = val
	}
	return results, nil
}

func GetUUID() string {
	return uuid.New().String()
}