package entity

import "time"

type PointLog struct {
	ID            int64      `json:"id" gorm:"type:int;primaryKey;autoIncrement:true"`
	UID           int64      `json:"uid" gorm:"type:int;not null"`
	LeaderboardID int64      `json:"leaderboard_id" gorm:"type:int;not null;default: 0"`
	Point         int64      `json:"point" gorm:"type:int;not null;default:0"`
	CreatedAt     *time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt     *time.Time `json:"updated_at" gorm:"type:timestamp;null"`
}

type PointLogRequest struct {
	UID           int64 `json:"uid"`
	LeaderboardID int64 `json:"leaderboard_id"`
	Point         int64 `json:"score"`
}

func (r PointLogRequest) ToPointLog() PointLog {
	var now = time.Now()
	var pointLog = PointLog{}
	pointLog.UID = r.UID
	pointLog.LeaderboardID = r.LeaderboardID
	pointLog.Point = r.Point
	pointLog.CreatedAt = &now
	pointLog.UpdatedAt = &now

	return pointLog
}

type PointLogResponse struct {
	UID            int64      `json:"uid"`
	Name           string     `json:"name"`
	ProfilePicture string     `json:"profile_picture"`
	Point          int64      `json:"point"`
	Ranking        int64      `json:"ranking"`
	CreatedAt      *time.Time `json:"-"`
}

func (r PointLog) ToPointLogResponse() PointLogResponse {
	var pointLogResponse = PointLogResponse{}
	pointLogResponse.UID = r.UID
	pointLogResponse.Ranking = r.LeaderboardID
	pointLogResponse.Point = r.Point
	pointLogResponse.CreatedAt = r.CreatedAt

	return pointLogResponse
}
