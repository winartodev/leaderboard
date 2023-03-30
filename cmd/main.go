package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/winartodev/leaderboard/configs"
	"github.com/winartodev/leaderboard/controller"
	"github.com/winartodev/leaderboard/handler"
	"github.com/winartodev/leaderboard/repository"
)

type srv struct {
	router             *httprouter.Router
	leaderboardHandler handler.LeaderboardHandler
	pointLogHandler    handler.PointLogHandler
}

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	gormDB, err := config.LoadDatabase()
	if err != nil {
		panic(err)
	}

	redisClient, err := config.LoadRedisClient()
	if err != nil {
		panic(err)
	}

	// initial repository
	leaderboardRepository := repository.NewLeaderboardRepository(repository.LeaderboardRepository{DB: gormDB})
	pointLogRepository := repository.NewPointLogRepository(repository.PointLogRepository{DB: gormDB, RedisClient: redisClient})
	userRepository := repository.NewUserRepository(repository.UserRepository{DB: gormDB, RedisClient: redisClient})

	// initial controller
	leaderboardController := controller.NewLeaderboardController(controller.LeaderboardController{LeaderboardRepository: leaderboardRepository})
	userController := controller.NewUserController(controller.UserController{UserRepository: userRepository})
	pointLogController := controller.NewPointLogController(controller.PointLogController{LeaderboardRepository: leaderboardRepository, PointLogRepository: pointLogRepository, UserController: userController})

	// initial handler
	leaderboardHandler := handler.LeaderboardHandler(handler.LeaderboardHandler{LeaderboardController: leaderboardController})
	pointLogHandler := handler.PointLogHandler(handler.PointLogHandler{PointLogController: pointLogController})

	srv := srv{
		router:             httprouter.New(),
		leaderboardHandler: leaderboardHandler,
		pointLogHandler:    pointLogHandler,
	}
	srv.routes()

	fmt.Printf("http listen and serve at :%v\n", config.App.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", config.App.Port), srv.router); err != nil {
		panic(err)
	}
}

func (s *srv) healthz() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("OK"))
	}
}

func (s *srv) routes() {
	s.router.GET("/healthcheck", s.healthz())

	// Leaderboard
	s.router.GET("/leaderboards", s.leaderboardHandler.GetAllLeaderboard())
	s.router.POST("/leaderboards", s.leaderboardHandler.CreateLeaderboard())
	s.router.GET("/leaderboards/:id", s.leaderboardHandler.GetLeaderboard())
	s.router.PUT("/leaderboards/:id", s.leaderboardHandler.UpdateLeaderboard())
	s.router.DELETE("/leaderboards/:id", s.leaderboardHandler.DeleteLeaderboard())

	// Point
	s.router.POST("/leaderboard/push-point", s.pointLogHandler.AddPoint())
	s.router.GET("/leaderboard/view-point/:id", s.pointLogHandler.GetAllPointByLeaderboardID())
	s.router.GET("/leaderboard/user-point/:id", s.pointLogHandler.GetPointByUserID())

	// user
	s.router.POST("/users", nil)
}
