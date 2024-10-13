import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
from statsmodels.tsa.arima.model import ARIMA
import pickle

# Load dataset
df = pd.read_csv('./machine_learning/Processed_Trash_Fill_Data.csv', parse_dates=['Timestamp'], index_col='Timestamp')
df = df.asfreq('30T')

# Kiểm tra dữ liệu
print(df.head())

# Chọn cột remaining_fill (%)
remaining_fill = df['RemainingFill(%)']

# Vẽ biểu đồ chuỗi thời gian
plt.plot(remaining_fill)
plt.title('Remaining Fill (%) Over Time')
plt.xlabel('Time')
plt.ylabel('Remaining Fill (%)')
plt.show()

# Khởi tạo mô hình ARIMA với tham số (p=5, d=1, q=0) (cần thử nghiệm thêm với các giá trị khác)
model = ARIMA(remaining_fill, order=(5, 1, 0))
model_fit = model.fit()

# Tóm tắt mô hình
print(model_fit.summary())

# Dự đoán dữ liệu trong tương lai (ví dụ: 10 bước tiếp theo)
forecast = model_fit.forecast(steps=10)

print(forecast)

# Lưu mô hình ARIMA vào file
with open('./machine_learning/arima_model.pkl', 'wb') as f:
    pickle.dump(model_fit, f)
