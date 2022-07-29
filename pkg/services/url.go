package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/luisfer00/url-shortener/internal/config"
	"github.com/luisfer00/url-shortener/pkg/models"
	"github.com/luisfer00/url-shortener/pkg/utils"
)

var client = config.GetRedisClient()

var (
	ErrEmptyFuncParams = errors.New("error: invoking function with empty params")
)

func GetURL(slug string) (*models.Url, error) {
	url := models.Url{}

	if slug == "" {
		return nil, ErrEmptyFuncParams
	}

	urlString, err := client.Get(context.Background(), fmt.Sprintf("url:%v", slug)).Result()
	if err != nil {
		return nil, err
	}

	url.Slug = slug
	url.Url = urlString
	
	return &url, err
}

func InsertURL(url string) (*models.Url, error) {
	if url == ""  {
		return nil, ErrEmptyFuncParams
	}

	slug, err := utils.GenerateSlug()
	if err != nil {
		return nil, err
	}

	err = client.Set(context.Background(), fmt.Sprintf("url:%v", slug), url, time.Hour).Err()
	if err != nil {
		return nil, err
	}
	
	return &models.Url{
		Slug: slug,
		Url: url,
	}, nil
}