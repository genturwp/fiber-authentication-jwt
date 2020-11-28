package api

import (
	"fiberauthenticationjwt/entities"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

//UserManagementHandler api handler related to user management
func UserManagementHandler(handler *handler) {
	handler.app.Post("/user/register-user", handler.handleRegisterUser)
	handler.app.Post("/user/create-credential", handler.handleCreateCredential)
	handler.app.Post("/user/authenticate", handler.handleAuthenticateUser)
	handler.secured.Get("/user/info", handler.handleGetUserInfo)
}

func (handler *handler) handleRegisterUser(c *fiber.Ctx) error {
	ctx := c.Context()

	userProfile := new(entities.UserProfile)
	if err := c.BodyParser(userProfile); err != nil {
		errObj := &errObj{
			ErrorCode: fiber.StatusBadRequest,
			ErrorMsg:  err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(errObj)
	}

	_userProfile, err := handler.service.UserManagementService.CreateUserProfile(ctx, userProfile)
	if err != nil {
		errObj := &errObj{
			ErrorCode: fiber.StatusInternalServerError,
			ErrorMsg:  err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errObj)
	}

	return c.Status(fiber.StatusOK).JSON(_userProfile)
}

func (handler *handler) handleCreateCredential(c *fiber.Ctx) error {
	ctx := c.Context()
	credential := new(entities.Credential)
	if err := c.BodyParser(credential); err != nil {
		errObj := &errObj{
			ErrorCode: fiber.StatusBadRequest,
			ErrorMsg:  err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(errObj)
	}

	_credential, err := handler.service.UserManagementService.CreateUserCredential(ctx, credential)
	if err != nil {
		errObj := &errObj{
			ErrorCode: fiber.StatusInternalServerError,
			ErrorMsg:  err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errObj)
	}

	return c.Status(fiber.StatusOK).JSON(_credential)
}

func (handler *handler) handleAuthenticateUser(c *fiber.Ctx) error {
	ctx := c.Context()

	credential := new(entities.Credential)
	if err := c.BodyParser(credential); err != nil {
		errObj := &errObj{
			ErrorCode: fiber.StatusBadRequest,
			ErrorMsg:  err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(errObj)
	}

	authObj, err := handler.service.UserManagementService.AuthenticateUser(ctx, credential)
	if err != nil {
		errObj := &errObj{
			ErrorCode: fiber.StatusInternalServerError,
			ErrorMsg:  err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errObj)
	}
	return c.Status(fiber.StatusOK).JSON(authObj)
}

func (handler *handler) handleGetUserInfo(c *fiber.Ctx) error {
	ctx := c.Context()
	user := c.Locals("user").(*jwt.Token)
	jwtClaims := user.Claims.(*entities.JwtApplicationUserClaim)

	userProfile, err := handler.service.UserManagementService.GetUserInfo(ctx, jwtClaims.ProfileID)
	if err != nil {
		errObj := &errObj{
			ErrorCode: fiber.StatusInternalServerError,
			ErrorMsg:  err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errObj)
	}
	return c.Status(fiber.StatusOK).JSON(userProfile)

}
