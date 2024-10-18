import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.svm import SVC
import joblib

# Bước 1: Tải dữ liệu và tiền xử lý
file_path = './machine_learning/Processed_Trash_Fill_Data.csv'
df = pd.read_csv(file_path)

# Kiểm tra cột 'RemainingFill(%)' và 'Weight(kg)' có tồn tại hay không
if 'RemainingFill(%)' not in df.columns or 'Weight(kg)' not in df.columns:
    raise ValueError("Các cột 'RemainingFill(%)' và 'Weight(kg)' không tồn tại trong dữ liệu CSV")

# Bước 2: Gán nhãn phân loại trạng thái thùng rác
def classify_trash_status(weight, remaining_fill):
    # Phân loại RemainingFill (Distance Level)
    if remaining_fill <= 20:
        distance_level = 'Low'
    elif 20 < remaining_fill <= 80:
        distance_level = 'Medium'
    else:
        distance_level = 'High'
    
    # Phân loại Weight (Weight Level)
    if weight <= 7:
        weight_level = 'Low'
    elif 7 < weight <= 14:
        weight_level = 'Medium'
    else:
        weight_level = 'High'
    
    # Phân loại trạng thái thùng rác (Bin Status)
    if distance_level == 'Low':
        if weight_level == 'Low' or weight_level == 'Medium':
            return 'Unfilled'
        else:
            return 'Half-Filled'
    elif distance_level == 'Medium':
        if weight_level == 'Low':
            return 'Unfilled'
        elif weight_level == 'Medium':
            return 'Half-Filled'
        else:
            return 'Filled'
    else:  # High (nhiều khoảng trống)
        if weight_level == 'Low':
            return 'Half-Filled'
        else:
            return 'Filled'

# Áp dụng hàm phân loại cho từng hàng dữ liệu
df['label'] = df.apply(lambda row: classify_trash_status(row['Weight(kg)'], row['RemainingFill(%)']), axis=1)

# Bước 3: Chuẩn bị dữ liệu huấn luyện
X = df[['Weight(kg)', 'RemainingFill(%)']]  # Đầu vào: Trọng lượng và Mức độ đầy
y = df['label']  # Nhãn phân loại (Trống, Gần đầy, Đầy)

# Chia dữ liệu thành tập huấn luyện và tập kiểm tra
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

# Bước 4: Huấn luyện mô hình SVM
svm_model = SVC(kernel='linear', random_state=42)
svm_model.fit(X_train, y_train)

# Bước 5: Lưu mô hình đã huấn luyện
svm_model_filename = './machine_learning/svm_model.pkl'
joblib.dump(svm_model, svm_model_filename)

print("Mô hình SVM đã được huấn luyện và lưu thành công.")
