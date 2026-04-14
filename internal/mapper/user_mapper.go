package mapper

import (
	dto "movie-management/internal/dto/user"
	"movie-management/internal/models"
)

func ToUserModelFromRegister(req *dto.RequestRegisterUser, hashedPassword string) *models.User {
	return &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "user",
	}
}

func ToRegisterResponse(user *models.User) *dto.ResponseRegisterUser {
	return &dto.ResponseRegisterUser{
		ID:        int(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func ToLoginResponse(user *models.User) *dto.ResponseLoginUser {
	return &dto.ResponseLoginUser{
		ID:    int(user.ID),
		Email: user.Email,
		Role:  user.Role,
	}
}

func ToUserResponse(req *models.User, genres []models.Genre) *dto.UserResponse {
	var genreNames []string

	for _, genre := range genres {
		genreNames = append(genreNames, genre.Name)
	}

	return &dto.UserResponse{
		ID:              req.ID,
		Name:            req.Name,
		Email:           req.Email,
		Role:            req.Role,
		FavouriteGenres: genreNames,
		CreatedAt:       req.CreatedAt,
	}
}
