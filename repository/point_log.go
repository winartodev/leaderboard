package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/winartodev/leaderboard/commons"
	"github.com/winartodev/leaderboard/entity"
	"github.com/winartodev/leaderboard/helper"
	"gorm.io/gorm"
)

type PointLogRepository struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

func NewPointLogRepository(repository PointLogRepository) PointLogRepository {
	return PointLogRepository{
		DB:          repository.DB,
		RedisClient: repository.RedisClient,
	}
}

func (r *PointLogRepository) CreatePointLogDB(ctx context.Context, data entity.PointLog) (id int64, err error) {
	err = r.DB.Create(&data).Error
	if err != nil {
		return id, err
	}

	return data.ID, err
}

func (r *PointLogRepository) GetPointLogByLeaderboardIDDB(ctx context.Context, leaderboardID int64) (result []entity.PointLogResponse, err error) {
	rows, err := r.DB.Raw(GetPointByLeaderboardIDQuery, leaderboardID).Rows()
	if err != nil {
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		var row entity.PointLogResponse
		rows.Scan(&row.UID, &row.Point, &row.CreatedAt, &row.Ranking)

		result = append(result, row)
	}
	return result, err
}

func (r *PointLogRepository) GetPointLogByLeaderboardIDAndUserIDDB(ctx context.Context, data entity.PointLogRequest) (result entity.PointLogResponse, err error) {
	err = r.DB.Raw(GetPointByLeaderboardIDAndUserIDQuery, data.LeaderboardID, data.UID).Row().Scan(&result.UID, &result.Point, &result.CreatedAt, &result.Ranking)
	if err != nil && err != gorm.ErrRecordNotFound {
		return result, err
	}

	return result, err
}

func (r *PointLogRepository) GetPointLogByLeaderboardIDRedis(ctx context.Context, key string) (result []entity.PointLogResponse, err error) {
	scores, err := r.RedisClient.ZRevRangeWithScores(ctx, key, 0, 100).Result()
	if err != nil {
		return result, err
	}

	var rank int64 = 1
	for _, score := range scores {
		var data entity.PointLogResponse
		data.UID, err = strconv.ParseInt(fmt.Sprint(score.Member), 10, 64)
		if err != nil {
			return result, err
		}

		data.Point = helper.ExtractBitToInt64(int64(score.Score), 0, 32)
		data.Ranking = int64(rank)
		result = append(result, data)
		rank++
	}

	return result, err
}

func (r *PointLogRepository) GetPointLogByLeaderboardIDAndUserIDRedis(ctx context.Context, key string, data entity.PointLogRequest) (result entity.PointLogResponse, err error) {
	rank, err := r.RedisClient.ZRevRank(ctx, key, fmt.Sprint(data.UID)).Result()
	if err != nil {
		return result, err
	}

	score, err := r.RedisClient.ZScore(ctx, key, fmt.Sprint(data.UID)).Result()
	if err != nil {
		return result, err
	}

	result.UID = data.UID
	result.Ranking = rank + 1
	result.Point = helper.ExtractBitToInt64(int64(score), 0, 32)

	return result, nil
}

func (r *PointLogRepository) PushPointToRedis(ctx context.Context, key string, data entity.PointLogResponse) (err error) {
	timeNow := int64(commons.UnixEpochalypse - data.CreatedAt.Unix())
	point := int64(data.Point<<32) | timeNow

	err = r.RedisClient.ZAdd(ctx, key, redis.Z{
		Member: data.UID,
		Score:  float64(point),
	}).Err()
	if err != nil {
		return err
	}

	return err
}

func (r *PointLogRepository) IncrementPointRedis(ctx context.Context, key string, data entity.PointLog) (err error) {
	score, err := r.RedisClient.ZScore(ctx, key, fmt.Sprint(data.UID)).Result()
	if err != nil && err != redis.Nil {
		return err
	}

	var newPoint int64
	timeNow := int64(commons.UnixEpochalypse - data.CreatedAt.Unix())

	if score > 0 {
		oldPoint := helper.ExtractBitToInt64(int64(score), 0, 32)
		newPoint = int64((data.Point+oldPoint)<<32) | timeNow
	} else {
		newPoint = int64(data.Point<<32) | timeNow
	}

	err = r.RedisClient.ZAdd(ctx, key, redis.Z{
		Member: data.UID,
		Score:  float64(newPoint),
	}).Err()
	if err != nil {
		return err
	}

	return err
}
