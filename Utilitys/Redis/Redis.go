package Redis

import (
	"github.com/go-redis/redis"
)

type Client struct {
	client *redis.Client
}

func New() *Client {
	r := new(Client)
	t := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	r.client = t
	return r
}

func (c *Client) Ping() error {
	_, err := c.client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func (c *Client) Set(Key, values string) error {

	if err := c.client.Set(Key, values, 0).Err(); err != nil {
		return err
	}
	return nil
}
func (c *Client) Get(key string) (string, error) {
	Val, err := c.client.Get(key).Result()
	if err != nil {
		return "", err
	}

	return Val, nil
}
