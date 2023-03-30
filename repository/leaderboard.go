package repository

import (
	"context"

	"github.com/winartodev/leaderboard/entity"
	"gorm.io/gorm"
)

type LeaderboardRepository struct {
	DB *gorm.DB
}

func NewLeaderboardRepository(repository LeaderboardRepository) LeaderboardRepository {
	return LeaderboardRepository{
		DB: repository.DB,
	}
}

func (r *LeaderboardRepository) CreateLeaderboardDB(ctx context.Context, data entity.Leaderboard) (result int64, err error) {
	tx := r.DB.Create(&data)
	if tx.Error != nil {
		return result, tx.Error
	}

	return data.ID, err
}

func (r *LeaderboardRepository) GetAllLeaderboardDB(ctx context.Context) (result []entity.Leaderboard, err error) {
	err = r.DB.Order("id DESC").Find(&result).Error
	if err != nil {
		return result, err
	}

	return result, err
}

func (r *LeaderboardRepository) GetLeaderboardByIDDB(ctx context.Context, id int64) (result entity.Leaderboard, err error) {
	err = r.DB.Where(`id = ?`, id).First(&result).Error
	if err != nil {
		return result, err
	}

	return result, err
}

func (r *LeaderboardRepository) UpdateLeaderboard(ctx context.Context, id int64, data entity.Leaderboard) (err error) {
	err = r.DB.Model(entity.Leaderboard{}).
		Where(`id = ?`, id).
		Updates(entity.Leaderboard{
			Slug:      data.Slug,
			IsActive:  data.IsActive,
			UpdatedAt: data.UpdatedAt,
		}).Error
	if err != nil {
		return err
	}

	return err
}

func (r *LeaderboardRepository) DeleteLeaderboardByIDDB(ctx context.Context, id int64) (err error) {
	err = r.DB.Where(`id = ?`, id).Delete(entity.Leaderboard{}).Error
	if err != nil {
		return err
	}

	return err
}
