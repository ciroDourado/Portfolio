package middlewares

import (
	"log"
	"net/http"

	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	"github.com/ulule/limiter/v3/drivers/store/redis"
	libredis "github.com/redis/go-redis/v9"
)

func NewRateLimiter(client *libredis.Client, formattedRate string) func(http.Handler) http.Handler {
	rate, err := limiter.NewRateFromFormatted(formattedRate)
	if err != nil {
		log.Fatal(err)
	}

	store, err := redis.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix:   "limiter",
		MaxRetry: 3,
	})
	if err != nil {
		log.Fatal(err)
	}

	instance := limiter.New(store, rate)

	return stdlib.NewMiddleware(instance).Handler
}