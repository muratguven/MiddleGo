package repository

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *repository[T] {
	return &repository[T]{
		db: db,
	}
}

// Insert data func
func (r *repository[T]) Insert(entity *T, ctx context.Context) error {
	return r.db.WithContext(ctx).Create(&entity).Error
}

// Insert all data array
func (r *repository[T]) InsertMany(entity *[]T, ctx context.Context) error {
	return r.db.WithContext(ctx).Create(&entity).Error
}

// Get data using id
func (r *repository[T]) GetById(id uuid.UUID, ctx context.Context) (*T, error) {
	var entity T
	err := r.db.WithContext(ctx).Model(&entity).Where("Id=? AND IsDeleted=?", id, true).FirstOrInit(&entity).Error
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

// Get all data
func (r *repository[T]) GetAll(ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, err
	}

	return &entities, nil
}

func (r *repository[T]) Update(entity *T, ctx context.Context) error {
	return r.db.WithContext(ctx).Save(&entity).Error
}

func (r *repository[T]) UpdateAll(entities *[]T, ctx context.Context) error {
	return r.db.WithContext(ctx).Save(&entities).Error
}

func (r *repository[T]) Delete(id uuid.UUID, ctx context.Context) error {
	var entity T
	return r.db.WithContext(ctx).Where("Id=?", id).FirstOrInit(&entity).UpdateColumn("IsDeleted", false).Error
}

func (r *repository[T]) SkipTake(skip int, take int, ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Offset(skip).Limit(take).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *repository[T]) Count(ctx context.Context) int64 {
	var entity T
	var count int64
	r.db.WithContext(ctx).Model(&entity).Count(&count)
	return count
}

func (r *repository[T]) CountWhere(params *T, ctx context.Context) int64 {
	var entity T
	var count int64
	r.db.WithContext(ctx).Model(&entity).Where(&params).Count(&count)
	return count
}
