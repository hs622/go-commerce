package database

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/hs622/ecommerce-cart/constants"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func FindOptionsParams(findOptions *options.FindOptionsBuilder, url *url.URL) error {

	findOptions = options.Find()
	query := url.Query()

	// Handle limit
	if limitStr := query.Get("limit"); limitStr != "" {
		if limit, err := strconv.ParseInt(limitStr, 10, 64); err == nil && limit > 0 {
			findOptions.SetLimit(limit)
		}
	} else {
		findOptions.SetLimit(constants.PRODUCT_LIMIT)
	}

	// Handle skip
	if skipStr := query.Get("skip"); skipStr != "" {
		if skip, err := strconv.ParseInt(skipStr, 10, 64); err == nil && skip >= 0 {
			findOptions.SetSkip(skip)
		}
	} else {
		findOptions.SetSkip(constants.PRODUCT_OFFSET)
	}

	projection, err := createProjection(url)
	if err != nil {
		return err
	}
	if len(projection) > 0 {
		findOptions.SetProjection(projection)
	}

	return nil
}

func FindOptionsFilters(filters *bson.D, url *url.URL) error {

	*filters = bson.D{}

	return nil
}

func FindOneOptionsParams(findOneOptions *options.FindOneOptionsBuilder, url *url.URL) error {

	findOneOptions = options.FindOne()

	projective, err := createProjection(url)
	if err != nil {
		return err
	}
	if len(projective) > 0 {
		findOneOptions.SetProjection(projective)
	}

	return nil
}

func createProjection(url *url.URL) (bson.D, error) {

	// Handle field projection
	projection := bson.D{}
	query := url.Query()

	if fields := query.Get("select"); fields != "" {
		for _, field := range strings.Split(fields, ",") {
			if trimmed := strings.TrimSpace(field); trimmed != "" {
				projection = append(projection, bson.E{Key: trimmed, Value: 1})
			}
		}
	}

	return projection, nil
}
