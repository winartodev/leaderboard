package entity

import (
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
)

func TestLeaderboardRequest_ToLeaderboard(t *testing.T) {
	type fields struct {
		Slug     string
		IsActive bool
	}

	now := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	tests := []struct {
		name   string
		fields fields
		want   Leaderboard
		mock   func()
	}{
		{
			name: "success",
			fields: fields{
				Slug:     "test-leaderboard-123",
				IsActive: false,
			},
			want: Leaderboard{
				Slug:      "test-leaderboard-123",
				IsActive:  false,
				CreatedAt: &now,
				UpdatedAt: &now,
			},
			mock: func() {
				monkey.Patch(time.Now, func() time.Time {
					return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
				})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			defer tt.mock()
			r := &LeaderboardRequest{
				Slug:     tt.fields.Slug,
				IsActive: tt.fields.IsActive,
			}
			if got := r.ToLeaderboard(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LeaderboardRequest.ToLeaderboard() = %v, want %v", got, tt.want)
			}
		})
	}
}
