from flask import Flask, request, jsonify
import pickle

app = Flask(__name__)

# Load ARIMA model
with open('./machine_learning/arima_model.pkl', 'rb') as f:
    arima_model = pickle.load(f)

@app.route('/predict', methods=['POST'])
def predict():
    data = request.get_json()
    steps = int(data.get('steps', 1))  # Số bước dự đoán

    # Dự đoán số bước tiếp theo
    forecast = arima_model.forecast(steps=steps)
    forecast_list = forecast.tolist()

    return jsonify(forecast=forecast_list)

if __name__ == '__main__':
    app.run(debug=True)
