package auth

import (
	"context"
	"errors"
	"smart-waste/pkgs/res"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Secret key để xác minh JWT
var jwtSecretKey = []byte("your-secret-key")

// VerifyToken xác minh token JWT và trả về claims nếu hợp lệ
func VerifyToken(ctx context.Context, tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired JWT")
	}

	return token, nil
}

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		tokenString := c.Get("Authorization")

		if tokenString == "" {
			res := res.NewRes(fiber.StatusBadRequest, "Missing or malformed JWT", false, nil)
			err := errors.New("missing or malformed JWT")
			res.SetError(err)
			return res.Send(c)
		}

		token, err := VerifyToken(ctx, tokenString)
		if err != nil {
			res := res.NewRes(fiber.StatusUnauthorized, "Invalid or expired JWT", false, nil)
			res.SetError(err)
			return res.Send(c)
		}

		// Lấy claims từ token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			res := res.NewRes(fiber.StatusUnauthorized, "Invalid JWT claims", false, nil)
			err := errors.New("invalid JWT claims")
			res.SetError(err)
			return res.Send(c)
		}

		// Lưu thông tin vào context để sử dụng trong các middleware khác
		c.Locals("username", claims["username"])
		c.Locals("role", claims["role"])

		return c.Next() // Tiếp tục middleware hoặc handler tiếp theo.
	}
}
