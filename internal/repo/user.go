package repo

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math"
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
	if err := r.db.Model(&model.User{}).Where("email = ?", email).Preload("Role").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserByID(id string) (*model.User, error) {
	var user model.User
	if err := r.db.Model(&model.User{}).Preload("Role").
		Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) CheckEmailExist(newEmail string) bool {
	rs := model.User{}
	err := r.db.Model(&model.User{}).Where("email = ?", newEmail).First(&rs).Error
	if err != nil {
		// email chưa tôn tại, nên không tìm thấy
		return false
	}
	return true
}

func (r *UserRepo) GetRoleID(namerole string) (uuid.UUID, error) {
	var roleID model.Role
	err := r.db.Where("name = ?", namerole).First(&roleID).Error
	if err != nil {
		return uuid.Nil, err
	}

	return roleID.ID, nil
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

func (r *UserRepo) GetAllUsers(filterEmail string, sortOrder string, page int, limit int) ([]model.User, int, error) {
	var users []model.User

	// Create a query builder
	query := r.db.Model(&model.User{}).
		Preload("Role").
		Preload("Wallets").
		Offset((page - 1) * limit).
		Limit(limit)

	// Apply filtering if filterName is provided
	if filterEmail != "" {
		query = query.Where("email ILIKE ?", fmt.Sprintf("%%%s%%", filterEmail))
	}

	// Apply sorting if sortOrder is provided
	if sortOrder != "" {
		query = query.Order(fmt.Sprintf("email %s", sortOrder))
	}

	// Count the total number of users (without pagination)
	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// Retrieve the users
	if err := query.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	// Iterate over users and populate the wallet count
	for i := range users {
		user := &users[i]
		walletCount := len(user.Wallets)
		user.WalletCount = walletCount
	}

	// Calculate the total number of pages based on the limit
	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	return users, totalPages, nil
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
