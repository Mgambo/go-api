package internal_repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type BaseRepositoryInterface[T any] interface {
	Create(entity *T) error
	GetByID(id string) (*T, error)
	Update(entity *T) error
	Delete(id string) error
	FindAll(filter interface{}) ([]*T, error)
	Count(filter interface{}) (int64, error)
}

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) Create(entity *T) (*T, error) {
	tx := r.db.Create(entity)
	if tx.Error != nil {
		fmt.Errorf("failed to create entity: %w", tx.Error)
		return nil, nil
	}

	return entity, nil
}

func (r *BaseRepository[T]) GetByID(id string) (*T, error) {
	var entity T
	if err := r.db.First(&entity, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get entity: %w", err)
	}
	return &entity, nil
}

func (r *BaseRepository[T]) Update(entity *T) error {
	if err := r.db.Save(entity).Error; err != nil {
		return fmt.Errorf("failed to update entity: %w", err)
	}
	return nil
}

func (r *BaseRepository[T]) Delete(id string) error {
	var entity T
	if err := r.db.Delete(&entity, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete entity: %w", err)
	}
	return nil
}

func (r *BaseRepository[T]) FindAll(filter interface{}) ([]*T, error) {
	var entities []*T
	query := r.db
	fmt.Println(query)

	if filter != nil {
		query = query.Where(filter)
	}

	if err := query.Find(&entities).Error; err != nil {
		return nil, fmt.Errorf("failed to list entities: %w", err)
	}

	return entities, nil
}

func (r *BaseRepository[T]) Count(filter interface{}) (int64, error) {
	var count int64
	query := r.db.Model(new(T))

	if filter != nil {
		query = query.Where(filter)
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count entities: %w", err)
	}
	return count, nil
}
