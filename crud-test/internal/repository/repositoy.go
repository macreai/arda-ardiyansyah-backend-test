package repository

import "gorm.io/gorm"

type Repository[T any] struct{}

func (r *Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, entity *T) error {
	return db.Delete(entity).Error
}

func (r *Repository[T]) FindById(db *gorm.DB, id uint64) (*T, error) {
	var data *T
	err := db.Where("id = ?", id).First(&data).Error
	return data, err
}

func (r *Repository[T]) GetAll(db *gorm.DB) ([]*T, error) {
	var entities []*T
	result := db.Find(&entities)
	return entities, result.Error
}
