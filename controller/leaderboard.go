package controller

import (
	"context"

	"github.com/winartodev/leaderboard/entity"
	"github.com/winartodev/leaderboard/repository"
)

type LeaderboardController struct {
	LeaderboardRepository repository.LeaderboardRepository
}

func NewLeaderboardController(controller LeaderboardController) LeaderboardController {
	return LeaderboardController{
		LeaderboardRepository: controller.LeaderboardRepository,
	}
}

func (c *LeaderboardController) CreateLeaderboard(ctx context.Context, data entity.LeaderboardRequest) (result int64, err error) {
	return c.LeaderboardRepository.CreateLeaderboardDB(ctx, data.ToLeaderboard())
}

func (c *LeaderboardController) GetAllLeaderboard(ctx context.Context) (result []entity.Leaderboard, err error) {
	return c.LeaderboardRepository.GetAllLeaderboardDB(ctx)
}

func (c *LeaderboardController) GetLeaderboard(ctx context.Context, id int64) (result entity.Leaderboard, err error) {
	return c.LeaderboardRepository.GetLeaderboardByIDDB(ctx, id)
}

func (c *LeaderboardController) UpdateLeaderboard(ctx context.Context, id int64, data entity.LeaderboardRequest) (err error) {
	_, err = c.LeaderboardRepository.GetLeaderboardByIDDB(ctx, id)
	if err != nil {
		return err
	}

	return c.LeaderboardRepository.UpdateLeaderboard(ctx, id, data.ToLeaderboard())
}

func (c *LeaderboardController) DeleteLeaderboard(ctx context.Context, id int64) (err error) {
	_, err = c.LeaderboardRepository.GetLeaderboardByIDDB(ctx, id)
	if err != nil {
		return err
	}

	return c.LeaderboardRepository.DeleteLeaderboardByIDDB(ctx, id)
}
