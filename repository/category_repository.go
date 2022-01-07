package repository

import "github.com/dimascapella/go-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
