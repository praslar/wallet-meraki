package repo

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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

func (r *UserRepo) GetAllUser(orderBy string) ([]model.User, error) {
	rs := []model.User{}
	if err := r.db.Preload("Role").Order(orderBy).Find(&rs).Error; err != nil {
		return nil, err
	}
	return rs, nil
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

func (r *UserRepo) CreateToken(newToken *model.Token) error {
	result := r.db.Create(&newToken)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepo) SymbolUnique(symbol string) bool {
	var token *model.Token
	result := r.db.Model(&model.Token{}).Where("symbol = ?", symbol).First(&token)
	if result == nil {
		logrus.Infof("token bi duplicate. Khong tao được ")
		return false
		// tim ko co thi chua co token
	}
	logrus.Infof("Khong tim thấy token. Co thể tao ")
	return true
	// tim co thi co token
}

func (r *UserRepo) DeleteToken(newToken *model.Token) error {
	result := r.db.Model(&newToken).Where("address = ? ", newToken.Address).Delete("symbol", newToken.Symbol)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepo) ValidateTokenInUse(tokenaddress uuid.UUID) bool {
	var transaction *model.Transaction
	result := r.db.Model(&model.Transaction{}).Where("token_address = ?", tokenaddress).First(&transaction)
	if result == nil {
		logrus.Infof("Token InUsed. Không Thể Xoá ")
		return false
	}
	logrus.Infof("Không tìm thấy Token InUsed. Có thể Xoá ")
	return true
	// tim co thi co token
}
