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

type LeaderboardHandler struct {
	LeaderboardController controller.LeaderboardController
}

func (h *LeaderboardHandler) CreateLeaderboard() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var ctx = r.Context()
		var data entity.LeaderboardRequest

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&data); err != nil {
			helper.FailedResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		result, err := h.LeaderboardController.CreateLeaderboard(ctx, data)
		if err != nil {
			helper.FailedResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		helper.SuccessResponse(w, http.StatusCreated, result)
	}
}

func (h *LeaderboardHandler) GetAllLeaderboard() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var ctx = r.Context()

		result, err := h.LeaderboardController.GetAllLeaderboard(ctx)
		if err != nil {
			helper.FailedResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		helper.SuccessResponse(w, http.StatusOK, result)
	}
}

func (h *LeaderboardHandler) GetLeaderboard() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var ctx = r.Context()

		id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
		if err != nil {
			helper.FailedResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		result, err := h.LeaderboardController.GetLeaderboard(ctx, id)
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

func (h *LeaderboardHandler) UpdateLeaderboard() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var ctx = r.Context()
		var data entity.LeaderboardRequest

		id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
		if err != nil {
			helper.FailedResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&data); err != nil {
			helper.FailedResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err = h.LeaderboardController.UpdateLeaderboard(ctx, id, data)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				helper.FailedResponse(w, http.StatusNotFound, err.Error())
				return
			}
			helper.FailedResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		helper.SuccessResponse(w, http.StatusOK, "success update data")
	}
}

func (h *LeaderboardHandler) DeleteLeaderboard() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var ctx = r.Context()

		id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
		if err != nil {
			helper.FailedResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err = h.LeaderboardController.DeleteLeaderboard(ctx, id)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				helper.FailedResponse(w, http.StatusNotFound, err.Error())
				return
			}
			helper.FailedResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		helper.SuccessResponse(w, http.StatusOK, "success delete data")
	}
}
