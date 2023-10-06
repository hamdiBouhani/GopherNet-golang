package pg

import "github.com/hamdiBouhani/GopherNet-golang/storage/model"

func (svc *DBConn) CreateBurrow(burrow *model.Burrow) error {
	return svc.Db.Create(burrow).Error
}

func (svc *DBConn) CreateManyBurrow(burrowList []*model.Burrow) error {
	return svc.Db.Create(burrowList).Error
}

func (svc *DBConn) IndexBurrow() ([]*model.Burrow, error) {
	var burrows []*model.Burrow
	err := svc.Db.Find(&burrows).Error
	if err != nil {
		return nil, err
	}
	return burrows, nil
}

func (svc *DBConn) SaveBurrow(burrow *model.Burrow) error {
	return svc.Db.Save(burrow).Error
}

func (svc *DBConn) ShowBurrow(id int64) (*model.Burrow, error) {
	var burrow model.Burrow
	err := svc.Db.First(&burrow, id).Error
	if err != nil {
		return nil, err
	}
	return &burrow, nil
}

func (svc *DBConn) UpdateBurrowAttributes(attributes map[string]interface{}) error {
	return svc.Db.Model(&model.Burrow{}).Updates(attributes).Error
}
