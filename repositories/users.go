package repositories

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
)

type UsersRepository struct {
	Repo *gorm.Repository[models.User]
}

func NewUsersRepository(db *database.Database) *UsersRepository {
	return &UsersRepository{
		Repo: gorm.NewRepository[models.User](db, models.User{}),
	}
}

func (user *UsersRepository) Login(username string, password string) (models.User, error) {
	var record models.User
	err := user.Repo.DB.Where("username = ? AND password = ?", username, password).First(&record).Error
	if err != nil {
		return models.User{}, err
	}

	query := user.Repo.DB.Preload("Roles").Preload("Roles.Role")
	query = query.Preload("Roles.Avatar")
	query = query.Preload("Accounts")

	err = query.First(&record, record.ID).Error
	if err != nil {
		return models.User{}, err
	}

	return record, nil
}

func (user *UsersRepository) Register(username string, password string) (models.User, error) {
	token, err := user.GenerateToken(64)
	if err != nil {
		return models.User{}, err
	}

	_, err = user.GetByUsername(username)
	if err == nil {
		return models.User{}, fmt.Errorf("user already exists")
	}

	record := models.User{
		Username: username,
		Password: password,
		ApiToken: token,
	}

	newUser, err := user.Repo.Create(record)
	if err != nil {
		return record, err
	}

	return newUser, nil
}

func (user *UsersRepository) CreateInvite(username string) (models.User, error) {
	token, err := user.GenerateToken(64)
	if err != nil {
		return models.User{}, err
	}

	record := models.User{
		InviteToken: token,
		Username:    username,
	}

	userRecord, err := user.Repo.Create(record)
	if err != nil {
		return models.User{}, err
	}

	return userRecord, nil
}

func (user *UsersRepository) CreateAndSendRecoveryToken(email string) (models.User, error) {
	fmt.Println("CreateAndSendRecoveryToken ----", email)

	existingUser, err := user.GetByEmail(email)
	if err != nil {
		fmt.Println("EXISTING USER GET BY EMAIL ERR----", err)
		return models.User{}, err
	}
	fmt.Println("EXISTING USER GET BY EMAIL ----", existingUser)

	// generate pw recovery token
	recoveryToken, err := user.GenerateToken(20)
	if err != nil {
		return models.User{}, err
	}
	fmt.Println("GENERATED RECOVERY TOKEN ----", recoveryToken)

	fmt.Println("UPDATE USER WITH RECOVERY TOKEN ----", existingUser)
	// save token and token expiration date in user table

	//write function that will send email to user with reset link

	return existingUser, nil
}

func (user *UsersRepository) ConfirmInvite(username string, password string, token string) (models.User, error) {
	invitedUser, err := user.GetByInviteToken(token)
	if err != nil {
		return models.User{}, fmt.Errorf("invalid invite token")
	}

	existingUser, err := user.GetByUsername(username)
	if err == nil && existingUser.ID != invitedUser.ID {
		return models.User{}, fmt.Errorf("user already exists")
	}

	fmt.Println("username is not taken")

	apiToken, err := user.GenerateToken(64)
	if err != nil {
		return models.User{}, err
	}

	invitedUser.Username = username
	invitedUser.Password = password
	invitedUser.ApiToken = apiToken

	err = user.Repo.DB.Save(invitedUser).Error
	if err != nil {
		return models.User{}, err
	}

	return invitedUser, nil
}

func (user *UsersRepository) GetAll() ([]models.User, error) {
	var records []models.User

	query := user.Repo.DB
	query = query.Preload("Roles").Preload("Roles.Role")
	query = query.Preload("Roles.Avatar")
	query = query.Preload("Accounts")

	err := query.Find(&records).Error
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (user *UsersRepository) GetByUsername(username string) (models.User, error) {
	var record models.User
	err := user.Repo.DB.Where("username = ?", username).First(&record).Error
	if err != nil {
		return models.User{}, err
	}

	return record, nil
}

func (user *UsersRepository) GetByEmail(email string) (models.User, error) {
	var record models.User
	// do a "fake" search with dummy data in db to test functionality
	err := user.Repo.DB.Where("username = ?", email).First(&record).Error
	if err != nil {
		return models.User{}, fmt.Errorf("user with email %s doesn't exist", email)
	}
	return record, nil
}

func (user *UsersRepository) GetByResetToken(token string) (models.User, error) {
	var record models.User
	err := user.Repo.DB.Where("username = ?", token).First(&record).Error
	if err != nil {
		return models.User{}, fmt.Errorf("invalid reset token")
	}
	return record, nil
}

func (user *UsersRepository) GetByAPIToken(token string) (models.User, error) {
	var record models.User
	err := user.Repo.DB.Where("api_token = ?", token).First(&record).Error
	if err != nil {
		return models.User{}, err
	}
	return record, nil
}

func (user *UsersRepository) GetByInviteToken(token string) (models.User, error) {
	var record models.User
	err := user.Repo.DB.Where("invite_token = ?", token).First(&record).Error
	if err != nil {
		return models.User{}, err
	}
	return record, nil
}

func (user *UsersRepository) GetAccounts(userID uint) ([]models.Account, error) {
	var userRecord models.User
	query := user.Repo.DB.Preload("Accounts")

	err := query.Find(&userRecord, userID).Error
	if err != nil {
		return nil, err
	}
	return userRecord.Accounts, nil
}

func (user *UsersRepository) GetAccount(userID uint, accountID uint) (models.Account, error) {
	var record models.Account
	err := user.Repo.DB.Where("user_id = ? AND id = ?", userID, accountID).First(&record).Error
	if err != nil {
		return models.Account{}, err
	}
	return record, nil
}

func (user *UsersRepository) GetUserAvatar(userID uint) (models.Avatar, error) {
	query := user.Repo.DB
	query = query.Preload("Avatar").Preload("Avatar.Documents").Preload("Avatar.LLM")
	query = query.Preload("Avatar.ActiveAgents").Preload("Avatar.ActiveAgents.Agent")
	query = query.Preload("Avatar.ActiveTools").Preload("Avatar.ActiveTools.Tool")
	query = query.Preload("Avatar.ActivePlugins").Preload("Avatar.ActivePlugins.Plugin")
	query = query.Where("user_id = ? and role_id = ?", userID, 1)

	var record models.AvatarRole
	err := query.First(&record).Error
	if err != nil {
		return models.Avatar{}, err
	}

	return record.Avatar, nil
}

func (user *UsersRepository) GenerateToken(n int) (string, error) {
	token := make([]byte, n)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(token), nil
}

func (user *UsersRepository) Update(userData models.User) (models.User, error) {
	result := user.Repo.DB.Save(&userData)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	var updatedUser models.User
	query := user.Repo.DB.Preload("Roles").Preload("Roles.Role")
	query = query.Preload("Roles.Avatar")
	query = query.Preload("Accounts")
	query = query.First(&updatedUser, userData.ID)

	if query.Error != nil {
		return models.User{}, query.Error
	}

	return updatedUser, nil
}

func (user *UsersRepository) SaveAccount(accountData models.Account) (models.Account, error) {
	result := user.Repo.DB.Save(&accountData)
	if result.Error != nil {
		return models.Account{}, result.Error
	}

	return accountData, nil

}

func (user *UsersRepository) UpdatePassword(userID uint, password string) (models.User, error) {
	var record models.User
	err := user.Repo.DB.Where("id = ?", userID).First(&record).Error
	if err != nil {
		return models.User{}, err
	}

	record.Password = password

	_, err = user.Repo.Update(userID, record)
	if err != nil {
		return models.User{}, err
	}

	return record, nil
}
