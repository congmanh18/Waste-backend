package python

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Thu thập dữ liệu từ bàn phím
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập cân nặng (kg): ")
	weightInput, _ := reader.ReadString('\n')
	weightInput = strings.TrimSpace(weightInput)

	fmt.Print("Nhập chất lượng không khí (ppm): ")
	airQualityInput, _ := reader.ReadString('\n')
	airQualityInput = strings.TrimSpace(airQualityInput)

	fmt.Print("Nhập mực nước (cm): ")
	waterLevelInput, _ := reader.ReadString('\n')
	waterLevelInput = strings.TrimSpace(waterLevelInput)

	fmt.Print("Nhập thời gian kể từ lúc bắt đầu (giây): ")
	timeSinceStartInput, _ := reader.ReadString('\n')
	timeSinceStartInput = strings.TrimSpace(timeSinceStartInput)

	// Gọi Python script và truyền tham số vào
	cmd := exec.Command("python", "./pkgs/python/script.py", weightInput, airQualityInput, waterLevelInput, timeSinceStartInput)

	// Chạy lệnh và lấy output
	output, err := cmd.CombinedOutput() // Lấy cả stdout và stderr
	if err != nil {
		log.Fatalf("Lỗi khi chạy Python: %s\nOutput: %s\n", err, output)
	}

	// In kết quả trả về từ Python
	fmt.Println(string(output))
}

func PassDataGoToPy(weightInput, airQualityInput, waterLevelInput, timeSinceStartInput string) (string, error) {
	// Gọi Python script và truyền tham số vào
	cmd := exec.Command("python", "./pkgs/python/script.py", weightInput, airQualityInput, waterLevelInput, timeSinceStartInput)

	// Chạy lệnh và lấy output
	output, err := cmd.CombinedOutput() // Lấy cả stdout và stderr
	if err != nil {
		log.Fatalf("Lỗi khi chạy Python: %s\nOutput: %s\n", err, output)
	}

	return string(output), nil
}
