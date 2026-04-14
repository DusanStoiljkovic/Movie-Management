package repository

import (
	"context"
	"movie-management/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) GetUserByID(ctx context.Context, userID int) (*models.User, error) {
	var user models.User

	if err := repo.db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User

	err := repo.db.WithContext(ctx).Preload("FavouriteGenres").Find(&users).Error

	return users, err
}

func (repo *UserRepository) AssignGenresToUser(ctx context.Context, user *models.User, genres []models.Genre) (*models.User, error) {

	err := repo.db.Model(&user).Association("FavouriteGenres").Replace(genres)
	if err != nil {
		return nil, err
	}

	return user, nil
}
