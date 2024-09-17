package handler

import (
	"context"
	"smart-waste/pkgs/res"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (w WasteBinHandler) HandlerReadWasteBin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		wateBinEntity, err := w.ReadWasteBinUsecase.ReadWasteBinByID(ctx, c.Params("id"))
		if err != nil {
			res := res.NewRes(fiber.StatusNotFound, "Unable to load wastebin information", false, nil)
			res.SetError(err)
			return res.Send(c)
		}


		// fmt.Println(predictTimeUntilFull((wateBinEntity.FilledLevel), predictedRateOfChange))

		res := res.NewRes(fiber.StatusOK, "WasteBin Information: ", true, wateBinEntity)
		return res.Send(c)
	}
}

// func predictTimeUntilFull(filledLevel, predictedRateOfChange float64) string {
// 	// Bước 1: Tính phần trăm còn lại để thùng rác đầy
// 	percentRemaining := 100.0 - filledLevel

// 	// Bước 2: Nếu tốc độ thay đổi nhỏ hoặc bằng 0, không thể dự đoán
// 	if predictedRateOfChange <= 0 {
// 		return "Tốc độ thay đổi nhỏ hơn 0"
// 	}

// 	// Bước 3: Tính số giây còn lại để thùng rác đầy
// 	timeRemainingSeconds := percentRemaining / predictedRateOfChange

// 	// Bước 4: Chuyển đổi thời gian từ giây sang giờ, phút, giây
// 	hours := int(math.Floor(timeRemainingSeconds / 3600))
// 	minutes := int(math.Floor(math.Mod(timeRemainingSeconds, 3600) / 60))
// 	seconds := int(math.Mod(timeRemainingSeconds, 60))

// 	return fmt.Sprintf("Thùng rác sẽ đầy sau %d giờ, %d phút và %d giây", hours, minutes, seconds)
// }
