package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/winartodev/leaderboard/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

func NewUserRepository(repository UserRepository) UserRepository {
	return UserRepository{
		DB:          repository.DB,
		RedisClient: repository.RedisClient,
	}
}

func (r *UserRepository) CreateUserDB() {}

func (r *UserRepository) GetUserByIDDB(ctx context.Context, id int64) (result *entity.User, err error) {
	err = r.DB.Where(`id = ?`, id).First(&result).Error
	if err != nil {
		return result, err
	}

	return result, err
}

func (r *UserRepository) GetUserByIDRedis(ctx context.Context, key string) (result *entity.User, err error) {
	resultStr, err := r.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return result, err
	}

	err = json.Unmarshal([]byte(resultStr), &result)
	if err != nil {
		return result, err
	}

	return result, err
}

func (r *UserRepository) SetUserRedis(ctx context.Context, key string, data []byte) (err error) {
	err = r.RedisClient.SetEx(ctx, key, string(data), 60*time.Hour).Err()
	if err != nil {
		return err
	}

	return err
}
