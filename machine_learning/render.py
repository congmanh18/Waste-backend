from flask import Flask, request, jsonify
from flask_cors import CORS
import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.svm import SVC
from statsmodels.tsa.arima.model import ARIMA
import joblib

# Khởi tạo Flask app
app = Flask(__name__)
CORS(app)

# Bước 1: Tải dữ liệu và tiền xử lý
file_path = './machine_learning/Processed_Trash_Fill_Data.csv'
df = pd.read_csv(file_path)

# Kiểm tra cột 'RemainingFill(%)' có tồn tại hay không
if 'RemainingFill(%)' not in df.columns:
    raise ValueError("Cột 'RemainingFill(%)' không tồn tại trong dữ liệu CSV")

# Tạo nhãn phân loại cho SVM dựa trên mức độ đầy
def classify_fill_level(filled_level):
    if filled_level < 20:
        return 0  # Trống
    elif 20 <= filled_level < 80:
        return 1  # Gần đầy
    else:
        return 2  # Đầy

df['label'] = df['FilledLevel(%)'].apply(classify_fill_level)

# Bước 2: Huấn luyện mô hình SVM
X = df[['Weight(kg)', 'FilledLevel(%)']]  # Đầu vào: Trọng lượng và Mức độ đầy
y = df['label']  # Nhãn phân loại
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

# Huấn luyện SVM
svm_model = SVC(kernel='linear')
svm_model.fit(X_train, y_train)

# Bước 3: Huấn luyện mô hình ARIMA với RemainingFill(%)
arima_model = ARIMA(df['RemainingFill(%)'], order=(5, 1, 0))
arima_model_fit = arima_model.fit()

# Bước 4: Lưu các mô hình SVM và ARIMA
svm_model_filename = './machine_learning/svm_model.pkl'
arima_model_filename = './machine_learning/arima_model.pkl'

# Lưu mô hình SVM
joblib.dump(svm_model, svm_model_filename)

# Lưu mô hình ARIMA
joblib.dump(arima_model_fit, arima_model_filename)

print("Models saved successfully.")

# Bước 5: Tải mô hình đã lưu khi khởi động ứng dụng
svm_model = joblib.load(svm_model_filename)
arima_model_fit = joblib.load(arima_model_filename)

print("Models loaded successfully.")