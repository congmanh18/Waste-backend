def predict_time_until_full(filled_level, predicted_rate_of_change):
    # Bước 1: Tính phần trăm còn lại để thùng rác đầy
    percent_remaining = 100.0 - filled_level

    # Bước 2: Nếu tốc độ thay đổi nhỏ hoặc bằng 0, không thể dự đoán
    if predicted_rate_of_change <= 0:
        return "Tốc độ thay đổi nhỏ hơn 0"

    # Bước 3: Tính số giây còn lại để thùng rác đầy
    time_remaining_seconds = percent_remaining / predicted_rate_of_change

    # Bước 4: Chuyển đổi thời gian từ giây sang giờ, phút, giây
    hours = int(time_remaining_seconds // 3600)
    minutes = int((time_remaining_seconds % 3600) // 60)
    seconds = int(time_remaining_seconds % 60)

    return f"Thùng rác sẽ đầy sau {hours} giờ, {minutes} phút và {seconds} giây"

# Ví dụ: FilledLevel hiện tại là 67%, PredictedRateOfChange là 0.0011481846868489198
filled_level = 5.0
predicted_rate_of_change = 0.51481846868489198


# In kết quả ra terminal
print(predict_time_until_full(filled_level, predicted_rate_of_change))
