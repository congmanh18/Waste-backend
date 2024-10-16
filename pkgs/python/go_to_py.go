package python

import (
	"log"
	"os/exec"
	"strings"
)

func PassDataGoToPy(weightInput, airQualityInput, waterLevelInput, timeSinceStartInput string) (string, error) {
	// Gọi Python script và truyền tham số vào
	cmd := exec.Command("python", "./pkgs/python/script.py", weightInput, airQualityInput, waterLevelInput, timeSinceStartInput)

	// Chạy lệnh và lấy output
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Lỗi khi chạy Python: %s\nOutput: %s\n", err, output)
	}

	cleanOutput := strings.TrimSpace(string(output))

	// dự đoán tỷ lệ thay đổi
	return cleanOutput, nil
}
