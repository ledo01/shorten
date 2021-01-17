package redis

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/ledo01/shorten/shorten"
	"github.com/pkg/errors"
)

type redisRepository struct {
	client *redis.Client
}

func newRedisClient(redisURL string) (*redis.Client, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)
	_, err = client.Ping().Result()

	return client, err
}

func NewRedisRepository(redisURL string) (shorten.RedirectRepository, error) {
	repo := &redisRepository{}
	client, err := newRedisClient(redisURL)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewRedisRepository")
	}
	repo.client = client
	return repo, nil
}

func (r *redisRepository) generateKey(code string) string {
	return fmt.Sprintf("redirect:%s", code)
}

func (r *redisRepository) Find(code string) (*shorten.Redirect, error) {
	redirect := &shorten.Redirect{}
	key := r.generateKey(code)
	data, err := r.client.HGetAll(key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.Find")
	}

	if len(data) == 0 {
		return nil, errors.Wrap(shorten.ErrRedirectNotFound, "repository.Redirect.Find")
	}

	createAt, err := strconv.ParseInt(data["createdAt"], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.Find")
	}

	redirect.Code = data["code"]
	redirect.URL = data["url"]
	redirect.CreatedAt = createAt
	return redirect, nil
}

func (r *redisRepository) Store(redirect *shorten.Redirect) error {
	key := r.generateKey(redirect.Code)
	data := map[string]interface{}{
		"code":      redirect.Code,
		"url":       redirect.URL,
		"createdAt": redirect.CreatedAt,
	}

	if _, err := r.client.HMSet(key, data).Result(); err != nil {
		return errors.Wrap(err, "repository.Redirect.Store")
	}

	return nil
}
