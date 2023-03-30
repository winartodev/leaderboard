package controller

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/winartodev/leaderboard/commons"
	"github.com/winartodev/leaderboard/entity"
	"github.com/winartodev/leaderboard/repository"
)

var (
	getPointByleaderboardIDKey = `leaderboard:%v`
)

type PointLogController struct {
	LeaderboardRepository repository.LeaderboardRepository
	PointLogRepository    repository.PointLogRepository
	UserController        UserController
}

func NewPointLogController(controller PointLogController) PointLogController {
	return PointLogController{
		LeaderboardRepository: controller.LeaderboardRepository,
		PointLogRepository:    controller.PointLogRepository,
		UserController:        controller.UserController,
	}
}

func (c *PointLogController) AddPointLog(ctx context.Context, data entity.PointLogRequest) (err error) {
	leaderboard, err := c.LeaderboardRepository.GetLeaderboardByIDDB(ctx, data.LeaderboardID)
	if err != nil {
		return err
	}

	if !leaderboard.IsActive {
		return commons.LeaderboardInActiveErr
	}

	pointLog := data.ToPointLog()
	_, err = c.PointLogRepository.CreatePointLogDB(ctx, pointLog)
	if err != nil {
		return err
	}

	err = c.PointLogRepository.IncrementPointRedis(ctx, fmt.Sprintf(getPointByleaderboardIDKey, leaderboard.ID), pointLog)
	if err != nil {
		return err
	}

	return err
}

func (c *PointLogController) GetAllPointByLeaderboardID(ctx context.Context, leaderboardID int64) (result []entity.PointLogResponse, err error) {
	leaderboard, err := c.LeaderboardRepository.GetLeaderboardByIDDB(ctx, leaderboardID)
	if err != nil {
		return result, err
	}

	if !leaderboard.IsActive {
		return result, commons.LeaderboardInActiveErr
	}

	var pointLogs []entity.PointLogResponse
	pointLogs, err = c.PointLogRepository.GetPointLogByLeaderboardIDRedis(ctx, fmt.Sprintf(getPointByleaderboardIDKey, leaderboardID))
	if err != nil {
		return result, err
	}

	if pointLogs == nil {
		pointLogs, err = c.PointLogRepository.GetPointLogByLeaderboardIDDB(ctx, leaderboardID)
		if err != nil {
			return result, err
		}

		for _, pointLog := range pointLogs {
			err = c.PointLogRepository.PushPointToRedis(ctx, fmt.Sprintf(getPointByleaderboardIDKey, leaderboardID), pointLog)
			if err != nil {
				return result, err
			}
		}
	}

	for _, score := range pointLogs {
		user, err := c.UserController.GetUserControllerByID(ctx, score.UID)
		if err != nil {
			return result, err
		}

		score.Name = user.Name
		score.ProfilePicture = user.ProfilePicture
		result = append(result, score)
	}

	return result, err
}

func (c *PointLogController) GetPointLogByLeaderboardIDAndUserID(ctx context.Context, data entity.PointLogRequest) (result entity.PointLogResponse, err error) {
	key := fmt.Sprintf(getPointByleaderboardIDKey, data.LeaderboardID)

	leaderboard, err := c.LeaderboardRepository.GetLeaderboardByIDDB(ctx, data.LeaderboardID)
	if err != nil {
		return result, err
	}

	if !leaderboard.IsActive {
		return result, commons.LeaderboardInActiveErr
	}

	user, err := c.UserController.GetUserControllerByID(ctx, data.UID)
	if err != nil {
		return result, err
	}

	result.Name = user.Name
	result.ProfilePicture = user.ProfilePicture

	point, err := c.PointLogRepository.GetPointLogByLeaderboardIDAndUserIDRedis(ctx, key, data)
	if err != nil && err != redis.Nil {
		return result, err
	}

	if err == redis.Nil {
		point, err = c.PointLogRepository.GetPointLogByLeaderboardIDAndUserIDDB(ctx, data)
		if err != nil && err != sql.ErrNoRows {
			return result, nil
		}
	}

	result.UID = point.UID
	result.Ranking = point.Ranking
	result.Point = point.Point

	return result, nil
}
