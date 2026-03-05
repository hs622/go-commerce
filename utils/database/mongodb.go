package database

import (
	"fmt"
	"net/url"
	"slices"
	"strconv"
	"strings"

	"github.com/hs622/ecommerce-cart/constants"
	"github.com/hs622/ecommerce-cart/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type UserFetchOptions struct {
	Url      *url.URL
	FullPath string
}

func FindOptionsParams(url *url.URL) (*options.FindOptionsBuilder, error) {

	findOptions := options.Find()
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
		return nil, err
	}
	if len(projection) > 0 {
		findOptions.SetProjection(projection)
	}

	return findOptions, nil
}

func FindOptionsFilters(filters *bson.D, url *url.URL) error {

	*filters = bson.D{}

	return nil
}

func FindOneOptionsParams(url *url.URL) (*options.FindOneOptionsBuilder, error) {

	findOneOpts := options.FindOne()

	if url.Query().Get("select") == "*" {
		return nil, nil
	}

	projective, err := createProjection(url)
	if err != nil {
		return nil, err
	}
	if len(projective) > 0 {
		findOneOpts.SetProjection(projective)
	}

	return findOneOpts, nil
}

func FindOneOptionsFilters(opts UserFetchOptions) (bson.D, error) {

	filters := bson.D{}
	crumbs := strings.Split(opts.FullPath, "/:")
	urlWithOutParams := crumbs[0]
	paramsKeysSlices := slices.Delete(crumbs, 0, 1)
	paramsValuesString, _ := strings.CutPrefix(opts.Url.Path, urlWithOutParams+"/")

	paramsValuesSlices := strings.Split(paramsValuesString, "/")

	fmt.Println(paramsKeysSlices, paramsValuesSlices, paramsValuesString)
	if len(paramsKeysSlices) == len(paramsValuesSlices) {
		for i, key := range paramsKeysSlices {
			filters = append(filters, bson.E{
				Key:   utils.ToSnakeCase(key),
				Value: paramsValuesSlices[i],
			})
		}
	} else {
		return nil, fmt.Errorf("Unequal params received.")
	}

	return filters, nil
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
