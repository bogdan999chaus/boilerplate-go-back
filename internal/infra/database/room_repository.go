package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const RoomsTableName = "rooms"

type room struct {
	Id             uint64     `db:"id,omitempty"`
	OrganizationId uint64     `db:"organization_id"`
	Name           string     `db:"name"`
	Description    *string    `db:"description"`
	CreatedDate    time.Time  `db:"created_date"`
	UpdatedDate    time.Time  `db:"updated_date"`
	DeletedDate    *time.Time `db:"deleted_date"`
}

type RoomRepository interface {
	FindById(id uint64) (*domain.Room, error)
	FindByOrgId(oId uint64) ([]domain.Room, error)
	Create(rm domain.Room) (*domain.Room, error)
	Update(rm domain.Room) (*domain.Room, error)
	Delete(id uint64) error
}

type roomRepository struct {
	coll db.Collection
	sess db.Session
}

func NewRoomRepository(session db.Session) RoomRepository {
	return roomRepository{
		coll: session.Collection(RoomsTableName),
		sess: session,
	}
}

func (r roomRepository) FindByOrgId(oId uint64) ([]domain.Room, error) {
	var rooms []room

	err := r.coll.
		Find(db.Cond{
			"organization_id": oId,
			"deleted_date":    nil,
		}).All(&rooms)
	if err != nil {
		return nil, err
	}

	rms := r.mapModelToDomainCollection(rooms)
	return rms, nil

}

func (r roomRepository) FindById(id uint64) (*domain.Room, error) {
	var rm room

	err := r.coll.Find(db.Cond{
		"id":           id,
		"deleted_date": nil,
	}).One(&rm)
	if err != nil {
		return nil, err
	}

	result := r.mapModelToDomain(rm)
	return &result, nil
}

func (r roomRepository) Create(rm domain.Room) (*domain.Room, error) {
	now := time.Now()

	model := r.mapDomainToModel(rm)
	model.CreatedDate = now
	model.UpdatedDate = now

	_, err := r.coll.Insert(model)
	if err != nil {
		return nil, err
	}

	var created room
	err = r.coll.Find(db.Cond{
		"organization_id": model.OrganizationId,
		"name":            model.Name,
		"deleted_date":    nil,
	}).OrderBy("-id").One(&created)
	if err != nil {
		return nil, err
	}

	result := r.mapModelToDomain(created)
	return &result, nil
}

func (r roomRepository) Update(rm domain.Room) (*domain.Room, error) {
	existingRoom, err := r.FindById(rm.Id)
	if err != nil {
		return nil, err
	}

	model := r.mapDomainToModel(rm)
	model.CreatedDate = existingRoom.CreatedDate
	model.UpdatedDate = time.Now()
	model.DeletedDate = existingRoom.DeletedDate

	err = r.coll.Find(db.Cond{
		"id":           rm.Id,
		"deleted_date": nil,
	}).Update(model)
	if err != nil {
		return nil, err
	}

	return r.FindById(rm.Id)
}

func (r roomRepository) Delete(id uint64) error {
	now := time.Now()

	return r.coll.Find(db.Cond{
		"id":           id,
		"deleted_date": nil,
	}).Update(map[string]interface{}{
		"deleted_date": now,
		"updated_date": now,
	})
}

func (r roomRepository) mapDomainToModel(rm domain.Room) room {
	return room{
		Id:             rm.Id,
		OrganizationId: rm.OrganizationId,
		Name:           rm.Name,
		Description:    rm.Description,
		CreatedDate:    rm.CreatedDate,
		UpdatedDate:    rm.UpdatedDate,
		DeletedDate:    rm.DeletedDate,
	}

}

func (r roomRepository) mapModelToDomain(rm room) domain.Room {
	return domain.Room{
		Id:             rm.Id,
		OrganizationId: rm.OrganizationId,
		Name:           rm.Name,
		Description:    rm.Description,
		CreatedDate:    rm.CreatedDate,
		UpdatedDate:    rm.UpdatedDate,
		DeletedDate:    rm.DeletedDate,
	}

}

func (r roomRepository) mapModelToDomainCollection(rooms []room) []domain.Room {
	rms := make([]domain.Room, len(rooms))
	for i := range rooms {
		rms[i] = r.mapModelToDomain(rooms[i])
	}
	return rms
}
