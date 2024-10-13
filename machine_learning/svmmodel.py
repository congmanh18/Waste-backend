import pandas as pd

# Load the uploaded dataset to inspect its contents
file_path = './machine_learning/Processed_Trash_Fill_Data.csv'
df = pd.read_csv(file_path)

# Display the first few rows of the dataset
df.head()

from sklearn.model_selection import train_test_split
from sklearn.svm import SVC
from sklearn.metrics import classification_report
import numpy as np

# Tạo nhãn phân loại dựa trên mức độ đầy
# 0 = Trống (Mức độ đầy < 20%), 1 = Gần đầy (20% <= Mức độ đầy < 80%), 2 = Đầy (Mức độ đầy >= 80%)
def classify_fill_level(filled_level):
    if filled_level < 20:
        return 0  # Trống
    elif 20 <= filled_level < 80:
        return 1  # Gần đầy
    else:
        return 2  # Đầy

df['label'] = df['FilledLevel(%)'].apply(classify_fill_level)

# Lấy dữ liệu đầu vào và nhãn
X = df[['Weight(kg)', 'FilledLevel(%)']]  # Đầu vào: Trọng lượng và Mức độ đầy
y = df['label']  # Nhãn phân loại

# Chia dữ liệu thành tập huấn luyện và kiểm tra
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

# Huấn luyện mô hình SVM
svm_model = SVC(kernel='linear')
svm_model.fit(X_train, y_train)

# Dự đoán trên tập kiểm tra
y_pred = svm_model.predict(X_test)

# Đánh giá mô hình
report = classification_report(y_test, y_pred, target_names=['Trống', 'Gần đầy', 'Đầy'])

report

# Tạo dữ liệu thử nghiệm (giả sử có trọng lượng và mức độ đầy khác nhau)
test_data = pd.DataFrame({
    'Weight(kg)': [5.0, 10.0, 12.5],  # Trọng lượng của thùng rác
    'FilledLevel(%)': [10, 50, 85]    # Mức độ đầy (%)
})

# Dự đoán trạng thái thùng rác cho dữ liệu thử nghiệm
test_predictions = svm_model.predict(test_data)

# Hiển thị kết quả dự đoán
for i, prediction in enumerate(test_predictions):
    print(f"Dữ liệu thử nghiệm {i+1}: Trọng lượng {test_data.iloc[i, 0]} kg, Mức độ đầy {test_data.iloc[i, 1]}% => Nhãn dự đoán: {prediction}")

