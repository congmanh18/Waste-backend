import random
import pandas as pd
from datetime import datetime, timedelta

# Thiết lập các tham số
max_weight = 20.0  # kg
min_weight = 5.0   # kg
max_filled_level = 100  # %
min_filled_level = 0
max_air_quality = 55  # ppm
min_air_quality = 5   # ppm
max_water_level = 7  # cm
start_date = datetime(2024, 9, 1, 8, 0, 0)  # Thời gian bắt đầu

# Hàm kiểm tra thời gian trong khoảng
def is_in_time_range(timestamp, ranges):
    for start_time, end_time in ranges:
        if start_time <= timestamp.time() <= end_time:
            return True
    return False

# Hàm tạo dữ liệu
def generate_bin_data(num_rows):
    data = []
    
    # Các khoảng thời gian đặc biệt
    special_time_ranges = [
        (datetime.strptime('06:00:00', '%H:%M:%S').time(), datetime.strptime('09:00:00', '%H:%M:%S').time()),
        (datetime.strptime('11:00:00', '%H:%M:%S').time(), datetime.strptime('13:00:00', '%H:%M:%S').time()),
        (datetime.strptime('16:00:00', '%H:%M:%S').time(), datetime.strptime('18:00:00', '%H:%M:%S').time())
    ]
    
    for i in range(num_rows):
        timestamp = start_date + timedelta(hours=i*2)  # Tăng thời gian mỗi dòng dữ liệu 2 giờ
        
        if is_in_time_range(timestamp, special_time_ranges):
            weight = round(random.uniform(16.0, 20.0), 1)
            filled_level = round(random.uniform(60, 100), 1)
        else:
            weight = round(random.uniform(min_weight, max_weight), 1)
            filled_level = round((weight / max_weight) * 100, 1)
        
        air_quality = round(random.uniform(min_air_quality, max_air_quality), 1)
        water_level = round(random.uniform(0, max_water_level), 1)
        
        row = {
            'ID': 'BIN001',
            'Timestamp': timestamp.strftime('%Y-%m-%d %H:%M:%S'),
            'Weight(kg)': weight,
            'FilledLevel(%)': filled_level,
            'AirQuality(ppm)': air_quality,
            'WaterLevel(cm)': water_level,
        }
        data.append(row)
    
    return pd.DataFrame(data)

# Tạo dữ liệu với 20 dòng
num_rows = 732
df = generate_bin_data(num_rows)
print(df)

# Lưu vào file CSV nếu cần thiết
df.to_csv('bin_data.csv', index=False)
