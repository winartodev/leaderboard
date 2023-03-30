package repository

var (
	GetPointByLeaderboardIDQuery = `
		SELECT
			a.uid AS uid,
			a.total_point AS point,
			a.created_time as created_at,
			ROW_NUMBER() OVER(ORDER BY a.total_point DESC, a.created_time ASC) AS ranking
		FROM (
			SELECT
				pl.uid,
				SUM(pl.point) AS total_point,
				MAX(pl.created_at) AS created_time
			FROM
				leaderboard_engine.point_logs pl
			WHERE
				pl.leaderboard_id = ?
			GROUP BY
				pl.uid
			LIMIT 100
		) AS a
		ORDER BY
			a.total_point DESC,
			a.created_time ASC
	`

	GetPointByLeaderboardIDAndUserIDQuery = `
		SELECT
				*
		FROM (
			SELECT
				a.uid AS uid,
				a.total_point AS point,
				a.created_time as created_at,
				ROW_NUMBER() OVER(ORDER BY a.total_point DESC, a.created_time ASC) AS ranking
			FROM (
				SELECT
					pl.uid,
					SUM(pl.point) AS total_point,
					MAX(pl.created_at) AS created_time
				FROM
					leaderboard_engine.point_logs pl
				WHERE
					pl.leaderboard_id = ?
				GROUP BY
					pl.uid
			) AS a
		) AS b
		WHERE b.uid = ?
	`
)
