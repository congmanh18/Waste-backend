import pandas as pd

# Đọc dữ liệu
data = pd.read_csv('data.csv')

# Tiền xử lý dữ liệu
data['Timestamp'] = pd.to_datetime(data['Timestamp'])
data.sort_values(by='Timestamp', inplace=True)  # Sắp xếp dữ liệu theo thời gian

# Tính toán sự thay đổi giữa các mẫu liên tiếp
data['TimeSinceLast'] = data['Timestamp'].diff().dt.total_seconds().fillna(0)  # Thời gian kể từ mẫu trước
data['FilledLevelChange'] = data['FilledLevel(%)'].diff().fillna(0)  # Thay đổi % giữa các mẫu

# Tính toán tỷ lệ thay đổi
data['RateOfChange'] = data['FilledLevelChange'] / data['TimeSinceLast']
data['RateOfChange'].replace([float('inf'), -float('inf')], 0, inplace=True)  # Xử lý vô cực
