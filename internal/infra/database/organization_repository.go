package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const OrganizationTableName = "organizations"

type organization struct {
	Id          uint64     `db:"id,omitempty"`
	UserId      uint64     `db:"user_id"`
	Name        uint64     `db:"name"`
	Description *string    `db:"description"`
	City        string     `db:"city"`
	Adress      string     `db:"adress"`
	Lat         float64    `db:"lat"`
	Lot         float64    `db:"lot"`
	CreatedDate time.Time  `db:"created_date"`
	UpdatedDate time.Time  `db:"updated_date"`
	DeletedDate *time.Time `db:"deleted_date"`
}

type organizationRepository struct {
	coll db.Collection
	sess db.Session
}

type organizationRepository interface {
	Save(o domain.Organization) (domain.Organization, error)
}

func NewOrganizationRepository(session db.Session) organizationRepository {
	return organizationRepository{
		sess: session,
		coll: session.Collection(OrganizationTableName),
	}
}

func (r organizationRepository) Save(o domain.Organization) (domain.Organization, error) {
	org := r.mapDomainToModel(o)
	now := time.Now()
	org.CreatedDate = now
	org.UpdatedDate = now

	err := r.coll.InsertReturning(&org)
	if err != nil {
		return domain.Organization{}, err
	}

	o = r.mapModelToDomain(org)
	return o, nil

}

func (r organizationRepository) mapDomainToModel(o domain.Organization) organization {
	return organization{
		Id:          o.Id,
		UserId:      o.UserId,
		Name:        o.Name,
		Description: o.Description,
		City:        o.City,
		Adress:      o.Adress,
		Lat:         o.Lat,
		Lot:         o.Lot,
		CreatedDate: o.CreatedDate,
		UpdatedDate: o.UpdatedDate,
		DeletedDate: o.DeletedDate,
	}
}

func (r organizationRepository) mapModelToDomain(o organization) domain.Organization {
	return domain.Organization{
		Id:          o.Id,
		UserId:      o.UserId,
		Name:        o.Name,
		Description: o.Description,
		City:        o.City,
		Adress:      o.Adress,
		Lat:         o.Lat,
		Lot:         o.Lot,
		CreatedDate: o.CreatedDate,
		UpdatedDate: o.UpdatedDate,
		DeletedDate: o.DeletedDate,
	}
}
