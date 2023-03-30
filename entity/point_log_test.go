package entity

import (
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
)

func TestPointLogRequest_ToPointLog(t *testing.T) {
	type fields struct {
		UID           int64
		LeaderboardID int64
		Point         int64
	}

	now := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	tests := []struct {
		name   string
		fields fields
		want   PointLog
		mock   func()
	}{
		{
			name: "success",
			fields: fields{
				UID:           123,
				LeaderboardID: 123,
				Point:         1,
			},
			want: PointLog{
				UID:           123,
				LeaderboardID: 123,
				Point:         1,
				CreatedAt:     &now,
				UpdatedAt:     &now,
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
			r := PointLogRequest{
				UID:           tt.fields.UID,
				LeaderboardID: tt.fields.LeaderboardID,
				Point:         tt.fields.Point,
			}
			if got := r.ToPointLog(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PointLogRequest.ToPointLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
