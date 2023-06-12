package repo

import (
	"fmt"
<<<<<<< HEAD
=======
	"github.com/google/uuid"
>>>>>>> f42f72261765b586a57e931f5a776a40c861c8d0
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

<<<<<<< HEAD
func (r *UserRepo) CreateUser(user *model.User) error {
	return r.db.Create(&user).Error
=======
func (r *UserRepo) CreateUser(newUser *model.User) error {
	result := r.db.Create(newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
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

func (r *UserRepo) GetAllUser() ([]model.User, error) {
	rs := []model.User{}
	if err := r.db.Find(&rs).Error; err != nil {
		return nil, err
	}
	return rs, nil
}

func (r *UserRepo) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	fmt.Print(r.db.Name())
	if err := r.db.Model(&model.User{}).Where("email = ?", email).Preload("Role").First(&user).Error; err != nil {
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

func (r *UserRepo) GetRoleID(namerole string) (uuid.UUID, error) {
	var roleID model.Role
	err := r.db.Where("name = ?", namerole).First(&roleID).Error
	if err != nil {
		return uuid.Nil, err
	}

	return roleID.ID, nil
}
func (r *UserRepo) GetUserIDByEmail(email string) (uuid.UUID, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
>>>>>>> f42f72261765b586a57e931f5a776a40c861c8d0
}

func (r *UserRepo) GetAllUser() ([]model.User, error) {
	var rs []model.User
	if err := r.db.Find(&rs).Error; err != nil {
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
	var user *model.User
	fmt.Print(r.db.Name())
	if err := r.db.Model(&model.User{}).Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) CreateWallet(newWallet *model.Wallet) error {
	result := r.db.Create(&newWallet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepo) CreateToken(newToken *model.Token) error {
	result := r.db.Create(&newToken)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepo) UpdateToken(newToken *model.Token) error {
	result := r.db.Model(&newToken).Where("address = ?", newToken.Address).Update("symbol", newToken.Symbol)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepo) DeleteToken(newToken *model.Token) error {
	result := r.db.Model(&newToken).Where("address = ?", newToken.Address).Delete("symbol", newToken.Symbol)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//func (r *UserRepo) TransferTokenAd(token *model.Token, transaction *model.Transaction) error {
//	var senderWallet model.Transaction
//	var recipientWallet model.Transaction
//	err := r.db.Where("address = ?", transaction.SenderWalletAddress).First(&senderWallet).Error
//	if err != nil {
//		return fmt.Errorf("Sender wallet not found: %v. ", err)
//	}
//	err = r.db.Where("address = ?", transaction.ReceiverWalletAddress).First(&recipientWallet).Error
//	if err != nil {
//		return fmt.Errorf("Recipient wallet not found: %v. ", err)
//	}
//
//	// Save the new token and transaction to the database
//	if err := r.db.Create(&token).Error; err != nil {
//		return fmt.Errorf("Failed to save token: %v. ", err)
//	}
//	if err := r.db.Create(&transaction).Error; err != nil {
//		return fmt.Errorf("Failed to save transaction: %v. ", err)
//	}
//	// Update amount from sender to receiver
//	senderWallet.Amount -= transaction.Amount
//	recipientWallet.Amount += transaction.Amount
//	if err := r.db.Save(&senderWallet).Error; err != nil {
//		return fmt.Errorf("Failed to update sender wallet balance: %v. ", err)
//	}
//	if err := r.db.Save(&recipientWallet).Error; err != nil {
//		return fmt.Errorf("Failed to update recipient wallet balance: %v. ", err)
//	}
//	return nil
//}
