package controller

import (
	"net/http"

	"github.com/Mobilizes/materi-be-alpro/modules/user/service"
	"github.com/Mobilizes/materi-be-alpro/modules/user/validation"
	"github.com/Mobilizes/materi-be-alpro/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

// @Summary Create User
// @Description Register a new user in the system
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequest true "User Registration Data"
// @Success 201 {object} utils.SwaggerSuccessResponse{data=dto.UserResponse} "User Created"
// @Failure 400 {object} utils.SwaggerErrorResponse "Invalid Request"
// @Failure 500 {object} utils.SwaggerErrorResponse "Internal Server Error"
// @Router /users [post]
func (ctrl *UserController) CreateUser(c *gin.Context) {
	req, err := validation.ValidateCreateUser(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := ctrl.service.CreateUser(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal membuat user")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "User berhasil dibuat", user)
}

// @Summary Get User By ID
// @Description Get details of a single user by ID
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} utils.SwaggerSuccessResponse{data=dto.UserResponse} "User Found"
// @Failure 404 {object} utils.SwaggerErrorResponse "User Not Found"
// @Router /users/{id} [get]
func (ctrl *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := ctrl.service.GetUserByID(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User tidak ditemukan")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User ditemukan", user)
}

// @Summary Get All Users
// @Description Get a list of all registered users
// @Tags Users
// @Produce json
// @Success 200 {object} utils.SwaggerSuccessResponse{data=[]dto.UserResponse} "Success"
// @Failure 404 {object} utils.SwaggerErrorResponse "Error"
// @Router /users [get]
func (ctrl *UserController) GetAllUsers(c *gin.Context) {

	user, err := ctrl.service.GetAllUsers()
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "error")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Success", user)
}
