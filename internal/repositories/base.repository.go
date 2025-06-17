package internal_repositories

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type BaseRepositoryInterface[T any] interface {
	Create(ctx context.Context, entity *T) error
	GetByID(ctx context.Context, id string) (*T, error)
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter interface{}) ([]*T, error)
	Count(ctx context.Context, filter interface{}) (int64, error)
}

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) Create(ctx context.Context, entity *T) (*T, error) {
	tx := r.db.WithContext(ctx).Create(entity)
	if tx.Error != nil {
		fmt.Errorf("failed to create entity: %w", tx.Error)
		return nil, nil
	}

	return entity, nil
}

func (r *BaseRepository[T]) GetByID(ctx context.Context, id string) (*T, error) {
	var entity T
	if err := r.db.WithContext(ctx).First(&entity, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get entity: %w", err)
	}
	return &entity, nil
}

func (r *BaseRepository[T]) Update(ctx context.Context, entity *T) error {
	if err := r.db.WithContext(ctx).Save(entity).Error; err != nil {
		return fmt.Errorf("failed to update entity: %w", err)
	}
	return nil
}

func (r *BaseRepository[T]) Delete(ctx context.Context, id string) error {
	var entity T
	if err := r.db.WithContext(ctx).Delete(&entity, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete entity: %w", err)
	}
	return nil
}

func (r *BaseRepository[T]) List(ctx context.Context, filter interface{}) ([]*T, error) {
	var entities []*T
	query := r.db.WithContext(ctx)

	if filter != nil {
		query = query.Where(filter)
	}

	if err := query.Find(&entities).Error; err != nil {
		return nil, fmt.Errorf("failed to list entities: %w", err)
	}
	return entities, nil
}

func (r *BaseRepository[T]) Count(ctx context.Context, filter interface{}) (int64, error) {
	var count int64
	query := r.db.WithContext(ctx).Model(new(T))

	if filter != nil {
		query = query.Where(filter)
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count entities: %w", err)
	}
	return count, nil
}
