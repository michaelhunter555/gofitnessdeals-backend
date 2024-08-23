package amazonclient

import (
	"os"

	paapi5 "github.com/goark/pa-api"
)

// create default client for now
func Client() paapi5.Client {
	client := paapi5.DefaultClient(
		"buycheapsoccer-20",
		os.Getenv("AMAZON_ACCESS_KEY"),
		os.Getenv("AMAZON_SECRET_KEY"),
	)

	return client
}
