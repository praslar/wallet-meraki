package repo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"wallet/internal/model"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Model(&model.User{}).Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserByID(id string) (*model.User, error) {
	var user model.User
	if err := r.db.Model(&model.User{}).Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) DeleteUser(userID uuid.UUID) error {
	user := model.User{}
	if err := r.db.Model(&user).Where("id = ?", userID).Take(&user).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) GetUser(userID uuid.UUID) (*model.User, error) {
	var user model.User
	if err := r.db.Preload("Role").Preload("Wallets").First(&user, userID).Error; err != nil {
		return nil, err
	}

	walletCount := len(user.Wallets)
	user.WalletCount = walletCount

	return &user, nil
}

func (r *UserRepo) GetRoleID(name string) (uuid.UUID, error) {
	var role model.Role
	err := r.db.Where("name = ?", name).First(&role).Error
	if err != nil {
		return uuid.Nil, err
	}
	return role.ID, nil
}

func (r *UserRepo) UpdateUserRole(userID uuid.UUID, role string) error {
	user := model.User{}
	if err := r.db.Model(&user).Where("id = ?", userID).Take(&user).Error; err != nil {
		return err
	}

	roleID, err := r.GetRoleID(role)
	if err != nil {
		return err
	}

	user.RoleID = roleID

	if err := r.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
