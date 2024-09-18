import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_squared_error
import joblib

# LinearRegression
# Đọc dữ liệu
data = pd.read_csv('data.csv')

# Tiền xử lý dữ liệu
data['Timestamp'] = pd.to_datetime(data['Timestamp'])
data.sort_values(by='Timestamp', inplace=True)  # Sắp xếp dữ liệu theo thời gian


data['TimeSinceLast'] = data['Timestamp'].diff().dt.total_seconds().fillna(0)  # Thời gian kể từ mẫu trước
data['FilledLevelChange'] = data['FilledLevel(%)'].diff().fillna(0)  # Thay đổi % giữa các mẫu

# Tính toán tỷ lệ thay đổi
data['RateOfChange'] = data['FilledLevelChange'] / data['TimeSinceLast']
data['RateOfChange'].replace([float('inf'), -float('inf')], 0, inplace=True)  # Xử lý vô cực

# Thêm biến TimeSinceStart (cần tính toán)
data['TimeSinceStart'] = (data['Timestamp'] - data['Timestamp'].iloc[0]).dt.total_seconds()

# Loại bỏ các hàng chứa NaN trong biến đầu vào và đầu ra
data.dropna(subset=['RateOfChange'], inplace=True)

# Xác định biến đầu vào và đầu ra cho mô hình hồi quy
X = data[['Weight(kg)', 'AirQuality(ppm)', 'WaterLevel(cm)', 'TimeSinceStart']]
y = data['RateOfChange']

# Chia dữ liệu thành tập huấn luyện và tập kiểm tra
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

# Xây dựng và huấn luyện mô hình hồi quy tuyến tính
model = LinearRegression()
model.fit(X_train, y_train)

# Dự đoán và đánh giá mô hình
y_pred = model.predict(X_test)
print("Mean Squared Error:", mean_squared_error(y_test, y_pred))

# Lưu mô hình vào tệp
joblib.dump(model, 'linear_regression_model_rate_of_change.pkl')
