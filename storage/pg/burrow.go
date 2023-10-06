package pg

import "github.com/hamdiBouhani/GopherNet-golang/storage/model"

func (svc *DBConn) CreateBurrow(burrow *model.Burrow) error {
	return svc.Db.Create(burrow).Error
}

func (svc *DBConn) CreateManyBurrow(burrowList []*model.Burrow) error {
	return svc.Db.Create(burrowList).Error
}
