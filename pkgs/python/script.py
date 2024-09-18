import joblib
import sys
import pandas as pd

def main():
    # Lấy dữ liệu đầu vào từ Go
    if len(sys.argv) < 5:
        print("Thieu du lieu dau vao")
        return

    weight = float(sys.argv[1])
    air_quality = float(sys.argv[2])
    water_level = float(sys.argv[3])
    time_since_start = float(sys.argv[4])

    # Tạo DataFrame từ dữ liệu đầu vào
    input_data = pd.DataFrame({
        'Weight(kg)': [weight],
        'AirQuality(ppm)': [air_quality],
        'WaterLevel(cm)': [water_level],
        'TimeSinceStart': [time_since_start]
    })

    # Tải mô hình từ tệp
    model = joblib.load('./pkgs/python/model.pkl')

    # Dự đoán với mô hình
    y_pred = model.predict(input_data)

    # Hiển thị kết quả dự đoán
    print(f"{y_pred[0]}")

if __name__ == '__main__':
    main()
