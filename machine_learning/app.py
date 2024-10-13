from flask import Flask, request, jsonify
import pandas as pd
import pickle
from sklearn.model_selection import train_test_split
from sklearn.svm import SVC
from statsmodels.tsa.arima.model import ARIMA

app = Flask(__name__)

# Bước 1: Tải dữ liệu và tiền xử lý
file_path = './machine_learning/Processed_Trash_Fill_Data.csv'
df = pd.read_csv(file_path)

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

# Lưu cả hai mô hình (tạm thời không lưu vào file để đơn giản hóa)
# with open('svm_trash_model.pkl', 'wb') as f:
#     pickle.dump(svm_model, f)

# with open('arima_model.pkl', 'wb') as f:
#     pickle.dump(arima_model_fit, f)

# Bước 4: API cho mô hình ARIMA
@app.route('/arima/predict', methods=['POST'])
def arima_predict():
    data = request.get_json()
    steps = int(data.get('steps', 1))  # Số bước dự đoán
    # Dự đoán số bước tiếp theo
    forecast = arima_model_fit.forecast(steps=steps)
    forecast_list = forecast.tolist()
    return jsonify(forecast=forecast_list)

# Bước 5: API cho mô hình SVM
@app.route('/svm/classify', methods=['POST'])
def svm_classify():
    data = request.get_json()
    weight = float(data['weight'])
    filled_level = float(data['filled_level'])

    # Dự đoán trạng thái thùng rác
    input_data = [[weight, filled_level]]
    prediction = svm_model.predict(input_data)
    
    return jsonify({'label': int(prediction[0])})

if __name__ == '__main__':
    app.run(port=5000, debug=True)
