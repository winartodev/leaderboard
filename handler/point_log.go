package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/winartodev/leaderboard/controller"
	"github.com/winartodev/leaderboard/entity"
	"github.com/winartodev/leaderboard/helper"
	"gorm.io/gorm"
)

type PointLogHandler struct {
	PointLogController controller.PointLogController
}

func (h *PointLogHandler) AddPoint() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var ctx = r.Context()
		var data entity.PointLogRequest

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&data); err != nil {
			helper.FailedResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err := h.PointLogController.AddPointLog(ctx, data)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				helper.FailedResponse(w, http.StatusNotFound, err.Error())
				return
			}
			helper.FailedResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		helper.SuccessResponse(w, http.StatusCreated, "success add user point")
	}
}

func (h *PointLogHandler) GetAllPointByLeaderboardID() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var ctx = r.Context()

		leaderboardID, err := strconv.ParseInt(p.ByName("id"), 10, 64)
		if err != nil {
			helper.FailedResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		result, err := h.PointLogController.GetAllPointByLeaderboardID(ctx, leaderboardID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				helper.FailedResponse(w, http.StatusNotFound, err.Error())
				return
			}
			helper.FailedResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		helper.SuccessResponse(w, http.StatusOK, result)
	}
}

func (h *PointLogHandler) GetPointByUserID() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var ctx = r.Context()
		var data entity.PointLogRequest

		leaderboardID, err := strconv.ParseInt(p.ByName("id"), 10, 64)
		if err != nil {
			helper.FailedResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&data); err != nil {
			helper.FailedResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		data.LeaderboardID = leaderboardID

		result, err := h.PointLogController.GetPointLogByLeaderboardIDAndUserID(ctx, data)
		if err != nil {
			// if err == gorm.ErrRecordNotFound {
			// 	helper.FailedResponse(w, http.StatusNotFound, err.Error())
			// 	return
			// }
			helper.FailedResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		helper.SuccessResponse(w, http.StatusOK, result)
	}
}
