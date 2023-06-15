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

func (r *UserRepo) CreateUser(newUser *model.User) error {
	result := r.db.Create(newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepo) CheckEmailExist(email string) bool {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return true
	}
	return false
}

func (r *UserRepo) GetUserByID(id string) (*model.User, error) {
	var user model.User
	if err := r.db.Model(&model.User{}).Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).Preload("Role").First(&user).Error
	if err != nil {
		return nil, err
	}
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

func (r *UserRepo) GetAllUsers(page int, limit int, sortField string, sortOrder string, filterEmail string) ([]model.User, int, error) {
	var users []model.User

	// Create a query builder
	query := r.db.Model(&model.User{}).
		Preload("Role").
		Preload("Wallets").
		Offset((page - 1) * limit).
		Limit(limit)

	// Apply filtering if filterEmail is provided
	if filterEmail != "" {
		query = query.Where("email ILIKE ?", fmt.Sprintf("%%%s%%", filterEmail))
	}
	// Apply sorting if sortField and sortOrder are provided
	if sortField != "" && sortOrder != "" {
		order := fmt.Sprintf("%s %s", sortField, sortOrder)
		query = query.Order(order)
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
