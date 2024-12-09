package handler

import (
	"context"
	"smart-waste/pkgs/res"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func classifyTrashStatus(weight, remainingFill float64) string {
	var distanceLevel, weightLevel, binStatus string

	// Phân loại RemainingFill (Distance Level)
	if remainingFill <= 20 {
		distanceLevel = "Low"
	} else if remainingFill <= 80 {
		distanceLevel = "Medium"
	} else {
		distanceLevel = "High"
	}

	// Phân loại Weight (Weight Level)
	if weight <= 7 {
		weightLevel = "Low"
	} else if weight <= 14 {
		weightLevel = "Medium"
	} else {
		weightLevel = "High"
	}

	// Phân loại trạng thái thùng rác (Bin Status)
	if distanceLevel == "Low" {
		if weightLevel == "Low" || weightLevel == "Medium" {
			binStatus = "Unfilled"
		} else {
			binStatus = "Half-Filled"
		}
	} else if distanceLevel == "Medium" {
		if weightLevel == "Low" {
			binStatus = "Unfilled"
		} else if weightLevel == "Medium" {
			binStatus = "Half-Filled"
		} else {
			binStatus = "Filled"
		}
	} else { // High
		if weightLevel == "Low" {
			binStatus = "Half-Filled"
		} else {
			binStatus = "Filled"
		}
	}

	return binStatus
}

func (w WasteBinHandler) HandlerSVM() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		wateBinEntity, err := w.ReadWasteBinUsecase.ReadWasteBinByID(ctx, c.Params("id"))
		if err != nil {
			res := res.NewRes(fiber.StatusNotFound, "Unable to load wastebin information", false, nil)
			res.SetError(err)
			return res.Send(c)
		}

		// Dereference the pointer and convert string to float64
		weight, err := strconv.ParseFloat(*wateBinEntity.Weight, 64)
		if err != nil {
			res := res.NewRes(fiber.StatusBadRequest, "Invalid weight value", false, nil)
			return res.Send(c)
		}

		remainingFill, err := strconv.ParseFloat(*wateBinEntity.RemainingFill, 64)
		if err != nil {
			res := res.NewRes(fiber.StatusBadRequest, "Invalid remaining fill value", false, nil)
			return res.Send(c)
		}

		svm := classifyTrashStatus(weight, remainingFill)

		res := res.NewRes(fiber.StatusOK, "SVM: ", true, svm)
		return res.Send(c)
	}
}
