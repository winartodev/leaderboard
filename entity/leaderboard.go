package entity

import "time"

type Leaderboard struct {
	ID        int64      `json:"id" gorm:"type:int;primaryKey;autoIncrement:true"`
	Slug      string     `json:"slug" gorm:"type:varchar(100);uniqueIndex:slug;not null;"`
	IsActive  bool       `json:"is_active" gorm:"type:boolean;not null;default:false"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"type:timestamp;not null"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"type:timestamp; null"`
	PointLog  []PointLog `json:"point_log,omitempty" gorm:"foreignKey:leaderboard_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type LeaderboardRequest struct {
	Slug     string `json:"slug"`
	IsActive bool   `json:"is_active"`
}

func (r *LeaderboardRequest) ToLeaderboard() Leaderboard {
	var now = time.Now()
	var leaderboard Leaderboard
	leaderboard.Slug = r.Slug
	leaderboard.IsActive = r.IsActive
	leaderboard.CreatedAt = &now
	leaderboard.UpdatedAt = &now

	return leaderboard
}
