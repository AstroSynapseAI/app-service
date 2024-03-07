package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/thanhpk/randstr"
	"gopkg.in/yaml.v2"

	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
)

type UsersController struct {
	rest.Controller
	User    *repositories.UsersRepository
	Account *gorm.Repository[models.Account]
}

func NewUsersController(db *database.Database) *UsersController {
	return &UsersController{
		User:    repositories.NewUsersRepository(db),
		Account: gorm.NewRepository[models.Account](db, models.Account{}),
	}
}

func (ctrl *UsersController) Run() {
	ctrl.Post("/login", ctrl.Login)
	ctrl.Post("/register", ctrl.Register)
	ctrl.Post("/register/invite", ctrl.RegisterInvite)
	ctrl.Post("/invite", ctrl.CreateInvite)
	ctrl.Post("/password_recovery", ctrl.CreatePasswordRecovery)
	ctrl.Post("/{id}/accounts/save", ctrl.SaveAccount)
	ctrl.Post("/{id}/save/profile", ctrl.SaveProfile)

	ctrl.Put("/{id}/change/password", ctrl.ChangePassword)
	// ctrl.Put("/{id}/change/email", ctrl.ChangeEmail)

	ctrl.Get("/{id}/accounts", ctrl.GetAccounts)
	ctrl.Get("/{id}/accounts/{account_id}", ctrl.GetAccount)
	ctrl.Get("/{id}/avatars", ctrl.GetAvatar)
	ctrl.Get("/invited/{token}", ctrl.GetInvitedUser)
	ctrl.Get("/token", ctrl.GetToken)
	ctrl.Get("/password_recovery/{token}", ctrl.ValidatePasswordRecoveryToken)

}

// Default CRUD routes
func (ctrl *UsersController) ReadAll(ctx *rest.Context) {
	fmt.Println("UsersController.ReadAll")
	users, err := ctrl.User.GetAll()
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}
	ctx.JsonResponse(http.StatusOK, users)
}

func (ctrl *UsersController) Read(ctx *rest.Context) {
	fmt.Println("Fetching user")
	userID := ctx.GetID("id")
	user, err := ctrl.User.Repo.Read(userID)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}
	ctx.JsonResponse(http.StatusOK, user)
}

func (ctrl *UsersController) GetToken(ctx *rest.Context) {
	fmt.Println("UsersController.GetToken")
	sessionToken := randstr.String(16)

	var reponseJson struct {
		Token string `json:"token"`
	}

	reponseJson.Token = sessionToken

	ctx.JsonResponse(http.StatusOK, reponseJson)
}

func (ctrl *UsersController) GetInvitedUser(ctx *rest.Context) {
	fmt.Println("Fetching user")
	inviteToken := ctx.GetParam("token")
	if inviteToken == "" {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	user, err := ctrl.User.GetByInviteToken(inviteToken)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, user)
}

func (ctrl *UsersController) ValidatePasswordRecoveryToken(ctx *rest.Context) {
	fmt.Println("UsersController.ValidatePasswordRecoveryToken")

	recoveryToken := ctx.GetParam("token")
	if recoveryToken == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Invalid request body"})
		return
	}

	user, err := ctrl.User.GetByResetToken(recoveryToken)
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: err.Error()})
		return
	}

	if time.Since(user.PasswordResetTokenExpiry) >= 24*time.Hour {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Token expired"})
		return
	}

	ctx.JsonResponse(http.StatusOK, user)
}

// Custom routes

// password recovery
func (ctrl *UsersController) CreatePasswordRecovery(ctx *rest.Context) {
	fmt.Println("UsersController.CreatePasswordRecovery")
	var input struct {
		Email string `json:"email"`
	}

	err := ctx.JsonDecode(&input)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Invalid request body"})
		return
	}

	record, err := ctrl.User.CreateAndSendRecoveryToken(input.Email)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: err.Error()})
		return
	}

	var envSetup string

	if os.Getenv("ENVIRONMENT") == "LOCAL DEV" {
		envSetup = "http://localhost:5173/password_reset/"
	}

	if os.Getenv("ENVIRONMENT") == "HEROKU DEV" {
		envSetup = "https://dev.asai.astrosynapse.ai/password_reset/"
	}

	if os.Getenv("ENVIRONMENT") == "AWS DEV" {
		envSetup = "https://asai.astrosynapse.ai/password_reset/"
	}

	fromEmail, err := getSendgridEmail()
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Error sending email"})
		return
	}

	from := mail.NewEmail("ASAI", fromEmail)
	subject := "Password recovery"
	to := mail.NewEmail("Recipient Name", input.Email)
	plainTextContent := "Password reset link"
	resetLink := envSetup + record.PasswordResetToken

	htmlContent := "<p>Open the following link to reset your password:</p>"
	htmlContent += "<p><a href=\"" + resetLink + "\">" + resetLink + "</a></p>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	apiKey, err := getSendgridAPIKey()
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Error sending email"})
		return
	}

	client := sendgrid.NewSendClient(apiKey)

	// Send the email
	response, err := client.Send(message)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Error sending email"})
		return
	}

	ctx.JsonResponse(http.StatusOK, response)
}

func getSendgridAPIKey() (string, error) {

	var Config struct {
		SendgridAPIKey string `yaml:"sendgrid_api_key"`
	}

	keys, err := os.ReadFile("./app/keys.yaml")
	if err != nil {
		fmt.Println("Error reading keys.yaml:", err)
	}

	err = yaml.Unmarshal(keys, &Config)
	if err != nil {
		fmt.Println("Error unmarshalling keys.yaml:", err)
	}

	return Config.SendgridAPIKey, nil
}

func getSendgridEmail() (string, error) {

	var Config struct {
		SendgridEmail string `yaml:"sendgrid_email"`
	}

	keys, err := os.ReadFile("./app/keys.yaml")
	if err != nil {
		fmt.Println("Error reading keys.yaml:", err)
	}

	err = yaml.Unmarshal(keys, &Config)
	if err != nil {
		fmt.Println("Error unmarshalling keys.yaml:", err)
	}

	return Config.SendgridEmail, nil
}

// create user invite
func (ctrl *UsersController) CreateInvite(ctx *rest.Context) {
	fmt.Println("UsersController.CreateInvite")
	var input struct {
		Username string `json:"username"`
	}

	err := ctx.JsonDecode(&input)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	record, err := ctrl.User.CreateInvite(input.Username)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, record)
}

// register invited user
func (ctrl *UsersController) RegisterInvite(ctx *rest.Context) {
	fmt.Println("UsersController.RegisterInvite")
	var reqData struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		InviteToken string `json:"invite_token"`
	}

	err := ctx.JsonDecode(&reqData)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Username is required"})
		return
	}
	if reqData.Username == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Username is required"})
		return
	}
	if reqData.Password == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Password is required"})
		return
	}
	if reqData.InviteToken == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Invite token is required"})
		return
	}

	if len(reqData.Password) < 8 {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Password must be at least 8 characters long"})
		return
	}

	user, err := ctrl.User.ConfirmInvite(
		reqData.Username,
		reqData.Password,
		reqData.InviteToken,
	)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: err.Error()})
		return
	}

	ctx.JsonResponse(http.StatusOK, user)
}

func (ctrl *UsersController) Register(ctx *rest.Context) {
	fmt.Println("UsersController.Register")
	var reqData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := ctx.JsonDecode(&reqData)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	user, err := ctrl.User.Register(reqData.Username, reqData.Password)
	if err != nil {
		ctx.SetStatus(http.StatusUnauthorized)
		return
	}

	ctx.JsonResponse(http.StatusOK, user)
}

func (ctrl *UsersController) Login(ctx *rest.Context) {
	fmt.Println("UsersController.Login")
	var reqData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := ctx.JsonDecode(&reqData)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: err.Error()})
		return
	}

	if reqData.Username == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Username is required"})
		return
	}

	if reqData.Password == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Password is required"})
		return
	}

	user, err := ctrl.User.Login(reqData.Username, reqData.Password)
	if err != nil {
		ctx.JsonResponse(http.StatusUnauthorized, struct{ Error string }{Error: err.Error()})
		return
	}

	ctx.JsonResponse(http.StatusOK, user)
}

func (ctrl *UsersController) GetAccounts(ctx *rest.Context) {
	fmt.Println("UsersController.GetAccounts")

	userID := ctx.GetID()

	accounts, err := ctrl.User.GetAccounts(userID)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, accounts)
}

// write a function that will fetch account by email from database
func (ctrl *UsersController) GetAccountByEmail(ctx *rest.Context) {
	fmt.Println("UsersController.GetAccountByEmail")
	//userID := ctx.GetID()
	var reqData struct {
		Email string `json:"email"`
	}

	err := ctx.JsonDecode(&reqData)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	account, err := ctrl.User.GetAccountByEmail(reqData.Email)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, account)
}

func (ctrl *UsersController) GetAccount(ctx *rest.Context) {
	accountID := ctx.GetID("account_id")
	account, err := ctrl.User.GetAccount(ctx.GetID(), accountID)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, account)
}

func (ctrl *UsersController) GetAvatar(ctx *rest.Context) {
	fmt.Println("UsersController.GetAvatar")
	avatar, err := ctrl.User.GetUserAvatar(ctx.GetID())
	if err != nil {
		ctx.SetStatus(http.StatusNotFound)
		return
	}

	ctx.JsonResponse(http.StatusOK, avatar)
}

func (ctrl *UsersController) Update(ctx *rest.Context) {
	fmt.Println("UsersController.Update")
	userID := ctx.GetID()

	var record models.User
	record.ID = userID

	err := ctx.JsonDecode(&record)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	user, err := ctrl.User.Update(record)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, user)
}

func (ctrl *UsersController) SaveAccount(ctx *rest.Context) {
	fmt.Println("UsersController.SaveAccount")
	userID := ctx.GetID()

	var record models.Account
	err := ctx.JsonDecode(&record)
	if err != nil {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	record.UserID = userID
	account, err := ctrl.User.SaveAccount(record)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, account)
}

func (ctrl *UsersController) SaveProfile(ctx *rest.Context) {
	fmt.Println("UsersController.SaveProfile")
	userID := ctx.GetID()

	var reqData struct {
		AccountID uint   `json:"account_id,omitempty"`
		Username  string `json:"username"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
	}

	err := ctx.JsonDecode(&reqData)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Invalid request body"})
		return
	}

	if reqData.Username == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Username is required"})
		return
	}

	account := models.Account{
		UserID:    userID,
		FirstName: reqData.FirstName,
		LastName:  reqData.LastName,
	}

	account.ID = reqData.AccountID

	if account.UserID == 0 || account.FirstName == "" || account.LastName == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Account data is invalid"})
		return
	}

	_, err = ctrl.User.SaveAccount(account)
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Failed to save account"})
		return
	}

	user := models.User{
		Username: reqData.Username,
	}

	user.ID = userID

	if user.Username == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Username is required"})
		return
	}

	userRecord, err := ctrl.User.Update(user)
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Failed to update user"})
		return
	}

	ctx.JsonResponse(http.StatusOK, userRecord)
}

func (ctrl *UsersController) ChangePassword(ctx *rest.Context) {
	fmt.Println("UsersController.ChangePassword")
	userID := ctx.GetID()

	var reqData struct {
		Password string `json:"password"`
	}

	err := ctx.JsonDecode(&reqData)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Invalid request body"})
		return
	}
	if len(reqData.Password) < 8 || reqData.Password == "" {
		ctx.JsonResponse(http.StatusBadRequest, struct{ Error string }{Error: "Password is not long enough"})
		return
	}

	user, err := ctrl.User.UpdatePassword(userID, reqData.Password)
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Failed to update password"})
		return
	}

	usr, err := ctrl.User.RemovePasswordResetToken(user.ID)
	if err != nil {
		ctx.JsonResponse(http.StatusInternalServerError, struct{ Error string }{Error: "Internal error"})
		return
	}

	ctx.JsonResponse(http.StatusOK, usr)
}

// func (ctrl *UsersController) ChangeEmail(ctx *rest.Context) {
// 	fmt.Println("UsersController.ChangeEmail")
// 	userID := ctx.GetID()
// 	var reqData struct {
// 		AccountID uint   `json:"account_id"`
// 		Email     string `json:"email"`
// 	}

// 	err := ctx.JsonDecode(&reqData)
// 	if err != nil {
// 		ctx.SetStatus(http.StatusBadRequest)
// 		return
// 	}

// 	var account models.Account

// }
