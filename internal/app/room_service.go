package app

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type RoomService interface {
	FindById(id uint64) (*domain.Room, error)
	FindByOrgId(oId uint64) ([]domain.Room, error)
	Create(room domain.Room) (*domain.Room, error)
	Update(room domain.Room) (*domain.Room, error)
	Delete(id uint64) error
}

type roomService struct {
	roomRepo database.RoomRepository
}

func NewRoomService(roomRepo database.RoomRepository) RoomService {
	return roomService{
		roomRepo: roomRepo,
	}
}

func (s roomService) FindById(id uint64) (*domain.Room, error) {
	return s.roomRepo.FindById(id)
}

func (s roomService) FindByOrgId(oId uint64) ([]domain.Room, error) {
	return s.roomRepo.FindByOrgId(oId)
}

func (s roomService) Create(room domain.Room) (*domain.Room, error) {
	return s.roomRepo.Create(room)
}

func (s roomService) Update(room domain.Room) (*domain.Room, error) {
	return s.roomRepo.Update(room)
}

func (s roomService) Delete(id uint64) error {
	return s.roomRepo.Delete(id)
}
