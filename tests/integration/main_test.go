//go:build integration && !unit

package integration_tests

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	redisImageTitle    = "redis:6.2.6-alpine3.15"
	redisDockerTimeout = 30 * time.Second
	redisPoolSize      = 150

	shortTimeout  = 500 * time.Millisecond
	redisClientDB = 1
)

var (
	_dao *IntegrationTestDAO
)

func TestMain(m *testing.M) {
	redisEndpoint, err := prepareRedisContainer(redisDockerTimeout)
	if err != nil {
		failInit(fmt.Errorf("prepare redis container: %w", err))
	}

	newDAO, err := newIntegrationDAO(redisEndpoint)
	if err != nil {
		failInit(fmt.Errorf("create integration DAO: %w", err))
	}

	_dao = newDAO

	fmt.Println("Redis endpoint:", redisEndpoint)

	m.Run()
}

func failInit(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "init integration tests: %s\n", err.Error())
	os.Exit(1)
}

func prepareRedisContainer(timeout time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	redisContainer, err := testcontainers.Run(
		ctx,
		redisImageTitle,
		testcontainers.WithExposedPorts("6379/tcp"),
		testcontainers.WithWaitStrategy(wait.ForLog("Ready to accept connections")),
	)
	if err != nil {
		return "", fmt.Errorf("run container: %w", err)
	}

	redisEndpoint, err := redisContainer.Endpoint(ctx, "")
	if err != nil {
		return "", fmt.Errorf("container endpoint: %w", err)
	}

	return redisEndpoint, nil
}

// TestRedisSimpleScenario is just a dummy test to check basic functionality of Redis in a container.
func TestRedisSimpleScenario(t *testing.T) {
	t.Run("DB is clean after prepareRedisDB", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), shortTimeout)
		defer cancel()

		keys, _, err := _dao.redisClient.Scan(ctx, 0, "*", 1).Result()
		if err != nil {
			t.Errorf("set 1 <= 1: %s", err.Error())
		}
		if len(keys) > 0 {
			t.Errorf("found some key in a DB %d (namely %s)", redisClientDB, keys[0])
		}
	})

	t.Run("prepare data", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), shortTimeout)
		defer cancel()

		err := _dao.redisClient.Set(ctx, "1", "1", 0).Err()
		if err != nil {
			t.Errorf("set 1 <= 1: %s", err.Error())
		}
	})

	t.Run("read prepared data", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), shortTimeout)
		defer cancel()

		x, err := _dao.redisClient.Get(ctx, "1").Result()
		if err != nil {
			t.Errorf("get 1: %s", err.Error())
		}
		if x != "1" {
			t.Errorf("expected 1, got %s", x)
		}
	})
}

type IntegrationTestDAO struct {
	redisClient redis.UniversalClient

	// nextGenEndpoint might be used just for debugging purpose.
	redisEndpoint string
}

func newIntegrationDAO(
	redisEndpoint string,
) (*IntegrationTestDAO, error) {
	result := IntegrationTestDAO{
		redisClient: redis.NewClient(&redis.Options{
			Addr:           redisEndpoint,
			DB:             redisClientDB,
			PoolSize:       redisPoolSize,
			MaxActiveConns: redisPoolSize,
			MaxIdleConns:   redisPoolSize,
		}),
		redisEndpoint: redisEndpoint,
	}

	return &result, nil
}
