package auth

import (
	"github.com/gofiber/fiber/v2"
)

func RoutesHandler(router fiber.Router) {
	r := Routes{}
	router.Post("/create_user", r.createUser).Name("AuthCreateUser")
	router.Post("/send_email_otp", r.sendEmailOtp).Name("AuthSendEmailOtp")
	router.Post("/verify_email_otp", r.verifyEmailOtp).Name("AuthVerifyEmailOtp")
	router.Post("/login", r.login).Name("AuthLogin")
	router.Post("/resent_otp", r.resendOtp).Name("AuthResetOtp")
	router.Post("/logout", r.logout).Name("AuthLogout")
	router.Get("/get_auth_credentials", r.getAuthCredentials).Name("AuthGetCredentials")
}

type Routes struct {
	service iService
}

func NewRoutes() iRoutes {
	return &Routes{}
}

func (*Routes) sendEmailOtp(c *fiber.Ctx) error {

	return c.SendString("Borrow Checker")
}
func (*Routes) verifyEmailOtp(c *fiber.Ctx) error {
	return c.SendString("Borrow Checker")
}
func (r *Routes) createUser(c *fiber.Ctx) error {
	body := new(createUserDto)

	if err := c.BodyParser(body); err != nil {
		return c.SendStatus(402)
	}

	result, err := r.service.CreateUser(*body)

	if err != nil {
		c.SendStatus(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Cannot create user",
		})
	}

	return c.JSON(result)
}

func (*Routes) login(c *fiber.Ctx) error {
	return c.SendString("Borrow Checker")
}
func (*Routes) resendOtp(c *fiber.Ctx) error {
	return c.SendString("Borrow Checker")
}
func (*Routes) logout(c *fiber.Ctx) error {
	return c.SendString("Borrow Checker")
}
func (*Routes) getAuthCredentials(c *fiber.Ctx) error {
	return c.SendString("Borrow Checker")
}
