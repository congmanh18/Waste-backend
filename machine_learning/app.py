from flask import Flask, request, jsonify
from flask_cors import CORS
import joblib
from statsmodels.tsa.arima.model import ARIMA
import pandas as pd
import os

# Khởi tạo Flask app
app = Flask(__name__)

# Cấu hình CORS để cho phép mọi nguồn truy cập (nếu muốn chỉ cho phép một domain cụ thể, bạn có thể chỉ định ở đây)
CORS(app, resources={r"/*": {"origins": "*"}})

# Đường dẫn tới các tệp mô hình và dữ liệu
svm_model_path = './machine_learning/svm_model.pkl'
arima_model_path = './machine_learning/arima_model.pkl'
file_path = './machine_learning/Processed_Trash_Fill_Data.csv'

# Kiểm tra xem các tệp có tồn tại không
if not os.path.exists(svm_model_path) or not os.path.exists(arima_model_path) or not os.path.exists(file_path):
    raise FileNotFoundError("Tệp mô hình hoặc dữ liệu không tồn tại!")

# Tải mô hình SVM và ARIMA đã lưu
svm_model = joblib.load(svm_model_path)
arima_model_fit = joblib.load(arima_model_path)

# Tải dữ liệu CSV
df = pd.read_csv(file_path)

# API cho ARIMA dự đoán
@app.route('/arima/predict', methods=['POST'])
def arima_predict():
    try:
        data = request.get_json()
        steps = int(data.get('steps', 1))  # Số bước dự đoán
        current_fill = float(data.get('current_fill'))  # Nhận giá trị "RemainingFill(%)" hiện tại
        
        # Cập nhật chuỗi dữ liệu với giá trị "RemainingFill(%)" hiện tại
        new_data = df['RemainingFill(%)'].tolist()
        new_data.append(current_fill)  # Thêm giá trị hiện tại vào chuỗi

        # Tạo mô hình ARIMA mới dựa trên dữ liệu đã cập nhật
        arima_model_updated = ARIMA(new_data, order=(5, 1, 0))
        arima_model_fit_updated = arima_model_updated.fit()

        # Dự đoán số bước tiếp theo dựa trên mô hình đã cập nhật
        forecast = arima_model_fit_updated.forecast(steps=steps)
        forecast_list = forecast.tolist()

        return jsonify(forecast=forecast_list), 200  # Trả về HTTP 200 OK
    except Exception as e:
        return jsonify({'error': str(e)}), 500  # Trả về HTTP 500 Internal Server Error

# API cho SVM phân loại
@app.route('/svm/classify', methods=['POST'])
def svm_classify():
    try:
        # Lấy dữ liệu JSON từ yêu cầu
        data = request.get_json()

        # Ghi log dữ liệu nhận được
        print("Received data:", data)
        
        # Chuyển đổi dữ liệu thành số thực
        weight = float(data['weight'])
        remaining_fill = float(data['remaining_fill'])

        # Ghi log dữ liệu sau khi chuyển đổi
        print("Weight:", weight, "Remaining fill:", remaining_fill)

        # Dự đoán với mô hình SVM
        input_data = [[weight, remaining_fill]]
        prediction = svm_model.predict(input_data)

        # Ghi log kết quả dự đoán
        print("Prediction result:", prediction)

        # Trả về kết quả dưới dạng JSON (không chuyển đổi thành số nguyên)
        return jsonify({'label': prediction[0]}), 200  # Trả về nhãn dạng chuỗi
    except Exception as e:
        # Ghi log lỗi ra console
        print("Error occurred:", str(e))
        return jsonify({'error': str(e)}), 500  # Trả về HTTP 500 Internal Server Error


# Chạy ứng dụng Flask
if __name__ == '__main__':
    app.run(port=5000, debug=True)
