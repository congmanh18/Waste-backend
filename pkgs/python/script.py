# import pandas as pd
# import joblib
# import sys

# def main():
#     # Lấy đối số truyền vào từ Go (nếu cần, ở đây ví dụ tên)
#     if len(sys.argv) > 1:
#         name = sys.argv[1]
#         print(f"Hello {name}, ......")

#     model = joblib.load('./pkgs/python/model.pkl')

#     # Đọc dữ liệu mới (thay đổi tên tệp và định dạng nếu cần)
#     new_data = pd.read_csv('./pkgs/python/data.csv')

#     # Tiền xử lý dữ liệu mới
#     new_data['Timestamp'] = pd.to_datetime(new_data['Timestamp'])
#     new_data.sort_values(by='Timestamp', inplace=True)

#     # Tính toán sự thay đổi giữa các mẫu liên tiếp
#     new_data['TimeSinceLast'] = new_data['Timestamp'].diff().dt.total_seconds().fillna(0)
#     new_data['FilledLevelChange'] = new_data['FilledLevel(%)'].diff().fillna(0)
#     new_data['RateOfChange'] = new_data['FilledLevelChange'] / new_data['TimeSinceLast']
#     new_data['RateOfChange'].replace([float('inf'), -float('inf')], 0, inplace=True)

#     # Thêm biến TimeSinceStart
#     new_data['TimeSinceStart'] = (new_data['Timestamp'] - new_data['Timestamp'].iloc[0]).dt.total_seconds()

#     # Chọn các biến đầu vào
#     X_new = new_data[['Weight(kg)', 'AirQuality(ppm)', 'WaterLevel(cm)', 'TimeSinceStart']]

#     # Dự đoán với mô hình
#     y_pred = model.predict(X_new)

#     # Thêm dự đoán vào dữ liệu mới
#     new_data['PredictedRateOfChange'] = y_pred

#     # Lưu kết quả dự đoán vào tệp (nếu cần)
#     new_data.to_csv('./pkgs/python/predictions.csv', index=False)

#     print("Save 'predictions.csv'")

# if __name__ == '__main__':
#     main()

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
    print(f"Du doan toc do thay doi: {y_pred[0]}")

if __name__ == '__main__':
    main()
