package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/go-chi/chi/v5"
)

type RoomController interface {
	FindById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type roomController struct {
	roomService app.RoomService
}

func NewRoomController(roomService app.RoomService) RoomController {
	return roomController{
		roomService: roomService,
	}
}

func (c roomController) FindById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid room id", http.StatusBadRequest)
		return
	}

	room, err := c.roomService.FindById(id)
	if err != nil {
		http.Error(w, "room not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(room)
}

func (c roomController) Create(w http.ResponseWriter, r *http.Request) {
	var room domain.Room

	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	createdRoom, err := c.roomService.Create(room)
	if err != nil {
		http.Error(w, "failed to create room", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdRoom)
}

func (c roomController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid room id", http.StatusBadRequest)
		return
	}

	var room domain.Room
	err = json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	room.Id = id

	updatedRoom, err := c.roomService.Update(room)
	if err != nil {
		http.Error(w, "failed to update room", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedRoom)
}

func (c roomController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid room id", http.StatusBadRequest)
		return
	}

	err = c.roomService.Delete(id)
	if err != nil {
		http.Error(w, "failed to delete room", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
