package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/wduartebr/goexpert/apis/internal/dto"

	"github.com/wduartebr/goexpert/apis/internal/entity"

	"github.com/wduartebr/goexpert/apis/internal/infra/database"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

// GetJWT godoc
// @Summary 		Get a user JWT
// @Description 	Get a user JWT
// @Tags 			users
// @Accept 			json
// @Produce 		json
// @Param 			request 	body	dto.GetJWTInput	true	"user credentials"
// @Success			201		{object} 	dto.GetJWTOutput
// @Failure			404		{object}	Error
// @Failure			500		{object}	Error
// @Router			/user/gen_token		[post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	experesin := r.Context().Value("experesin").(int)
	var user dto.GetJWTInput

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, token, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(experesin)).Unix(),
	})

	accessToken := dto.GetJWTOutput{AccessToken: token}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)

}

// CreateUser godoc
// @Summary 		Create User
// @Description 	Create User
// @Tags 			users
// @Accept 			json
// @Produce 		json
// @Param 			request 	body	dto.CreateUserInput	true	"user request"
// @Success			201
// @Failure			500			{object}	Error
// @Router			/user		[post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.UserDB.CreateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
