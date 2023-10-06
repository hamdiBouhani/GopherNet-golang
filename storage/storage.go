package storage

import "github.com/hamdiBouhani/GopherNet-golang/storage/model"

type Storage interface {
	CreateConnection() error
	Close() error
	Migrate() error
	Drop() error

	CreateBurrow(burrow *model.Burrow) error
	CreateManyBurrow(burrowList []*model.Burrow) error
	IndexBurrow() ([]*model.Burrow, error)
	SaveBurrow(burrow *model.Burrow) error
	ShowBurrow(id int64) (*model.Burrow, error)
	UpdateBurrowAttributes(attributes map[string]interface{}) error
}
