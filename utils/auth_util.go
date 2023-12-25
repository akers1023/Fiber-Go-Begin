package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// Params?
func CheckUserType(context *fiber.Ctx, role string) (err error) {
	userType := context.Params("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}

// Kiểm tra ID user và ID có trùng với tài nguyên hay không và chặn truy cập tài nguyên nếu là USER
func MatchUserTypeToUid(context *fiber.Ctx, userId string) (err error) {
	userType := context.Params("user_type")
	uid := context.Params("uid")
	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	// Kiểm tra loại User có phải là ADMIN hay không?
	err = CheckUserType(context, userType)
	return err
}
