package controllers

import (
	"fmt"
	"net/http"

	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
	"github.com/thanhpk/randstr"

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
	ctrl.Post("/{id}/accounts/save", ctrl.SaveAccount)
	ctrl.Post("/{id}/save/profile", ctrl.SaveProfile)

	ctrl.Put("/{id}/change/password", ctrl.ChangePassword)
	// ctrl.Put("/{id}/change/email", ctrl.ChangeEmail)

	ctrl.Get("/{id}/accounts", ctrl.GetAccounts)
	ctrl.Get("/{id}/accounts/{account_id}", ctrl.GetAccount)
	ctrl.Get("/{id}/avatars", ctrl.GetAvatar)
	ctrl.Get("/invited/{token}", ctrl.GetInvitedUser)
	ctrl.Get("/token", ctrl.GetToken)

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

// Custom routes
//
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
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	user, err := ctrl.User.ConfirmInvite(
		reqData.Username,
		reqData.Password,
		reqData.InviteToken,
	)

	if err != nil {
		fmt.Println(err)
		ctx.SetStatus(http.StatusInternalServerError)
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
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	user, err := ctrl.User.Login(reqData.Username, reqData.Password)
	if err != nil {
		ctx.SetStatus(http.StatusUnauthorized)
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
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	if len(reqData.Password) < 8 || reqData.Password == "" {
		ctx.SetStatus(http.StatusBadRequest)
		return
	}

	user, err := ctrl.User.UpdatePassword(userID, reqData.Password)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return
	}

	ctx.JsonResponse(http.StatusOK, user)
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
