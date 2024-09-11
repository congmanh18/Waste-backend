package res

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Res struct {
	StatusCode   int         `json:"status_code"`
	Message      string      `json:"message"`
	Success      bool        `json:"success"`
	Data         interface{} `json:"data"`
	Debug        bool        `json:"debug"`
	Error        error       `json:"-"`
	ErrorDetails string      `json:"error_details,omitempty"`
}

func NewRes(statusCode int, message string, success bool, data interface{}) *Res {
	return &Res{
		StatusCode: statusCode,
		Message:    message,
		Success:    success,
		Data:       data,
	}
}

// Hàm gửi JSON response
func (r *Res) Send(c *fiber.Ctx) error {
	return c.Status(r.StatusCode).JSON(r)
}

// Hàm cập nhật thông tin lỗi
func (r *Res) SetError(err error) {
	log.Printf("error occurred: %v", err)
	r.Error = err
	r.ErrorDetails = err.Error()
	r.Success = false
}

// Hàm cập nhật thông tin thành công
func (r *Res) SetSuccess(data interface{}, message string) {
	r.Data = data
	r.Message = message
	r.Success = true
	r.Error = nil
	r.ErrorDetails = ""
}

// Hàm thêm thông tin debug
func (r *Res) AddDebugInfo(debug bool) {
	r.Debug = debug
}

// Hàm chuyển Res thành JSON
func (r *Res) ToJSON() (string, error) {
	jsonData, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
