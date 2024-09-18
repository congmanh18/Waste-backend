import pandas as pd
import joblib

# Tải mô hình từ tệp
model = joblib.load('linear_regression_model_rate_of_change.pkl')

# Đọc dữ liệu mới (thay đổi tên tệp và định dạng nếu cần)
new_data = pd.read_csv('data.csv')

# Tiền xử lý dữ liệu mới
new_data['Timestamp'] = pd.to_datetime(new_data['Timestamp'])
new_data.sort_values(by='Timestamp', inplace=True)

# Tính toán sự thay đổi giữa các mẫu liên tiếp
new_data['TimeSinceLast'] = new_data['Timestamp'].diff().dt.total_seconds().fillna(0)
new_data['FilledLevelChange'] = new_data['FilledLevel(%)'].diff().fillna(0)
new_data['RateOfChange'] = new_data['FilledLevelChange'] / new_data['TimeSinceLast']
new_data['RateOfChange'].replace([float('inf'), -float('inf')], 0, inplace=True)

# Thêm biến TimeSinceStart
new_data['TimeSinceStart'] = (new_data['Timestamp'] - new_data['Timestamp'].iloc[0]).dt.total_seconds()

# Chọn các biến đầu vào
X_new = new_data[['Weight(kg)', 'AirQuality(ppm)', 'WaterLevel(cm)', 'TimeSinceStart']]

# Dự đoán với mô hình
y_pred = model.predict(X_new)

# Thêm dự đoán vào dữ liệu mới
new_data['PredictedRateOfChange'] = y_pred

# Lưu kết quả dự đoán vào tệp (nếu cần)
new_data.to_csv('predictions.csv', index=False)

print("Dự đoán hoàn tất và đã lưu vào 'predictions.csv'")
