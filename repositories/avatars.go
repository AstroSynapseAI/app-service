package repositories

import (
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
)

type AvatarsRepository struct {
	Repo *gorm.Repository[models.Avatar]
}

func NewAvatarsRepository(db *database.Database) *AvatarsRepository {
	return &AvatarsRepository{
		Repo: gorm.NewRepository[models.Avatar](db, models.Avatar{}),
	}
}

func (avatar *AvatarsRepository) Create(userID uint, data models.Avatar) (models.Avatar, error) {
	avatarRecord, err := avatar.Repo.Create(data)
	if err != nil {
		return models.Avatar{}, err
	}

	var user models.User
	result := avatar.Repo.DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return models.Avatar{}, result.Error
	}

	var role models.Role
	result = avatar.Repo.DB.Where("permission = ?", "owner").First(&role)
	if result.Error != nil {
		return models.Avatar{}, result.Error
	}

	userRole := models.AvatarRole{
		User:   user,
		Role:   role,
		Avatar: avatarRecord,
	}

	result = avatar.Repo.DB.Create(&userRole)
	if result.Error != nil {
		return models.Avatar{}, result.Error
	}

	record, err := avatar.Fetch(avatarRecord.ID)
	if err != nil {
		return models.Avatar{}, err
	}

	return record, nil
}

func (avatar *AvatarsRepository) Update(ID uint, data models.Avatar) (models.Avatar, error) {
	avatarRecord, err := avatar.Repo.Update(ID, data)
	if err != nil {
		return models.Avatar{}, err
	}

	record, err := avatar.Fetch(avatarRecord.ID)
	if err != nil {
		return models.Avatar{}, err
	}

	return record, nil
}

func (avatar *AvatarsRepository) Fetch(ID uint) (models.Avatar, error) {
	query := avatar.Repo.DB
	query = query.Preload("LLM")
	query = query.Preload("Documents")
	query = query.Preload("Roles").Preload("Roles.Role").Preload("Roles.User")
	query = query.Preload("ActiveAgents").Preload("ActiveAgents.Agent").Preload("ActiveAgents.LLM")
	query = query.Preload("ActiveAgents.ActiveAgentTools").Preload("ActiveAgents.ActiveAgentTools.Tool")
	query = query.Preload("ActiveTools").Preload("ActiveTools.Tool")
	query = query.Preload("ActivePlugins").Preload("ActivePlugins.Plugin")
	query = query.Preload("ActiveLLMs").Preload("ActiveLLMs.LLM")

	var record models.Avatar
	result := query.First(&record, ID)
	if result.Error != nil {
		return models.Avatar{}, result.Error
	}

	return record, nil
}
