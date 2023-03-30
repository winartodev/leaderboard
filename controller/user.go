package controller

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/winartodev/leaderboard/entity"
	"github.com/winartodev/leaderboard/repository"
)

var (
	getUserByIDKey = `user:%v`
)

type UserController struct {
	UserRepository repository.UserRepository
}

func NewUserController(controller UserController) UserController {
	return UserController{
		UserRepository: controller.UserRepository,
	}
}

func (u *UserController) GetUserControllerByID(ctx context.Context, id int64) (result *entity.User, err error) {
	key := fmt.Sprintf(getUserByIDKey, id)
	result, err = u.UserRepository.GetUserByIDRedis(ctx, key)
	if err != nil && err != redis.Nil {
		return result, err
	}

	if err == redis.Nil {
		result, err = u.UserRepository.GetUserByIDDB(ctx, id)
		if err != nil {
			return result, err
		}

		dataStr, _ := json.Marshal(result)
		err = u.UserRepository.SetUserRedis(ctx, key, dataStr)
		if err != nil {
			return result, err
		}
	}

	return result, err
}
