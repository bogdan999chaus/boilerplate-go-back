package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type MeasurementsController struct {
	measurementRepository database.MeasurementRepository
}

func NewMeasurementsController(
	measurementRepository database.MeasurementRepository,
) MeasurementsController {
	return MeasurementsController{
		measurementRepository: measurementRepository,
	}
}

func (c MeasurementsController) FindList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomId, err := strconv.ParseUint(r.PathValue("roomId"), 10, 64)
		if err != nil {
			BadRequest(w, err)
			return
		}

		deviceId, err := strconv.ParseUint(r.PathValue("deviceId"), 10, 64)
		if err != nil {
			BadRequest(w, err)
			return
		}

		var createdDateFrom *time.Time
		if v := r.URL.Query().Get("created-date-from"); v != "" {
			t, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				BadRequest(w, err)
				return
			}
			d := time.Unix(t, 0)
			createdDateFrom = &d
		}

		var createdDateTo *time.Time
		if v := r.URL.Query().Get("created-date-to"); v != "" {
			t, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				BadRequest(w, err)
				return
			}
			d := time.Unix(t, 0)
			createdDateTo = &d
		}

		measurements, err := c.measurementRepository.FindList(
			domain.Pagination{},
			database.MeasurementFilters{
				RoomId:          roomId,
				DeviceId:        deviceId,
				CreatedDateFrom: createdDateFrom,
				CreatedDateTo:   createdDateTo,
			},
		)
		if err != nil {
			log.Printf("MeasurementsController: %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, measurements)
	}
}
