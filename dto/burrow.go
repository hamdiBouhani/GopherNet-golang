package dto

import "github.com/hamdiBouhani/GopherNet-golang/storage/model"

type BurrowDto struct {
	Name     string  `json:"name"`
	Depth    float64 `json:"depth"`
	Wide     float64 `json:"wide"`
	Occupied bool    `json:"occupied"`
	Age      int     `json:"age"`
}

func (b *BurrowDto) ParseToModel() *model.Burrow {
	return &model.Burrow{
		Name:     b.Name,
		Depth:    b.Depth,
		Wide:     b.Wide,
		Occupied: b.Occupied,
		Age:      b.Age,
	}
}
