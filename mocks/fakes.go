package mocks

import (
	"github.com/brianvoe/gofakeit/v5"
	"github.com/google/uuid"
	"github.com/hamdiBouhani/GopherNet-golang/storage/model"
)

func MockBurrow(Occupied bool) *model.Burrow {
	return &model.Burrow{
		UUID:     uuid.New(),
		Name:     gofakeit.Name(),
		Depth:    gofakeit.Float64Range(1, 5),
		Wide:     gofakeit.Float64Range(1, 5),
		Occupied: Occupied,
		Age:      gofakeit.RandomInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
	}
}
