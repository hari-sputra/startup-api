package user

import "gorm.io/gorm"

type UserRepository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(id int) (User, error)
	Update(user User) (User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindById(id int) (User, error) {
	var user User

	err := r.db.Where("id = ?", id).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
