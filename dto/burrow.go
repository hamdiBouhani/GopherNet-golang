package dto

import (
	"github.com/google/uuid"
	"github.com/hamdiBouhani/GopherNet-golang/storage/model"
)

type BurrowDto struct {
	Name     string  `json:"name"`
	Depth    float64 `json:"depth"`
	Width    float64 `json:"width"`
	Occupied bool    `json:"occupied"`
	Age      int     `json:"age"`
}

func (b *BurrowDto) ParseToModel() *model.Burrow {
	return &model.Burrow{
		Name:     b.Name,
		Depth:    b.Depth,
		Width:    b.Width,
		Occupied: b.Occupied,
		Age:      b.Age,
		UUID:     uuid.New(),
	}
}
