package utils

import gonanoid "github.com/matoous/go-nanoid/v2"

func GenerateSlug() (string, error) {
	slug, err := gonanoid.New()

	return slug, err
}