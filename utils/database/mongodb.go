package database

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/hs622/ecommerce-cart/constants"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func QParamsConvertIntoFindOptions(findOptions *options.FindOptionsBuilder, url *url.URL) error {

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

	// Handle field selection
	if fields := query.Get("select"); fields != "" {
		projection := bson.D{}
		for _, field := range strings.Split(fields, ",") {
			if trimmed := strings.TrimSpace(field); trimmed != "" {
				projection = append(projection, bson.E{Key: trimmed, Value: 1})
			}
		}
		if len(projection) > 0 {
			findOptions.SetProjection(projection)
		}
	}

	return nil
}

func QParamsConvertIntoFilters(filters *bson.D, url *url.URL) error {

	filters = &bson.D{}

	return nil
}
