package repositories

import (
	"database/sql"

	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"

	db "gorm.io/gorm"
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
	query = query.Preload("ActiveAgents").Preload("ActiveAgents.Agent")
	query = query.Preload("ActiveTools").Preload("ActiveTools.Tool")
	query = query.Preload("ActivePlugins").Preload("ActivePlugins.Plugin")

	var record models.Avatar
	result := query.First(&record, ID)
	if result.Error != nil {
		return models.Avatar{}, result.Error
	}

	return record, nil
}

func (avatar *AvatarsRepository) SaveActiveAgent(avatarData models.ActiveAgent) error {
	result := avatar.Repo.DB.Save(&avatarData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (avatar *AvatarsRepository) SetActiveAgent(avatarID uint, agentID uint) error {
	var activeAgent models.ActiveAgent

	result := avatar.Repo.DB.Where("avatar_id = ? AND agent_id = ?", avatarID, agentID).First(&activeAgent)
	if result.Error == db.ErrRecordNotFound {
		activeAgent.AgentID = sql.NullInt64{Int64: int64(agentID), Valid: true}
		activeAgent.AvatarID = sql.NullInt64{Int64: int64(avatarID), Valid: true}
	}

	activeAgent.IsActive = true
	result = avatar.Repo.DB.Save(&activeAgent)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (avatar *AvatarsRepository) GetActiveAgents(avatarID uint) []models.ActiveAgent {
	query := avatar.Repo.DB
	query = query.Preload("Agent").Preload("LLM")
	query = query.Where("avatar_id = ?", avatarID)

	var activeAgents []models.ActiveAgent

	result := query.Find(&activeAgents)
	if result.Error != nil {
		return []models.ActiveAgent{}
	}

	return activeAgents
}

func (avatar *AvatarsRepository) GetActiveAgent(avatarID uint, agentID uint) (models.ActiveAgent, error) {
	query := avatar.Repo.DB
	query = query.Preload("Agent").Preload("LLM")
	query = query.Where("avatar_id = ? AND agent_id = ?", avatarID, agentID)

	var activeAgent models.ActiveAgent
	result := query.First(&activeAgent)
	if result.Error != nil {
		return models.ActiveAgent{}, result.Error
	}

	return activeAgent, nil
}

func (avatar *AvatarsRepository) GetActiveTools(avatarID uint) []models.ActiveTool {
	query := avatar.Repo.DB
	query = query.Preload("Tool")
	query = query.Where("avatar_id = ?", avatarID)

	var activeTools []models.ActiveTool

	result := query.Find(&activeTools)
	if result.Error != nil {
		return []models.ActiveTool{}
	}

	return activeTools
}

func (avatar *AvatarsRepository) GetActivePlugins(avatarID uint) []models.ActivePlugin {
	query := avatar.Repo.DB
	query = query.Preload("Plugin")
	query = query.Where("avatar_id = ?", avatarID)

	var activePlugins []models.ActivePlugin
	result := query.Find(&activePlugins)
	if result.Error != nil {
		return []models.ActivePlugin{}
	}

	return activePlugins
}

func (avatar *AvatarsRepository) GetDocuments(avatarID uint) []models.Document {
	var docs []models.Document
	result := avatar.Repo.DB.Where("avatar_id = ?", avatarID).Find(&docs)
	if result.Error != nil {
		return []models.Document{}
	}

	return docs
}
