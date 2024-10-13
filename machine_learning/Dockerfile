# Bước 1: Sử dụng Python base image từ Alpine Linux
FROM python:3.10-alpine AS base

# Cài đặt các dependencies cần thiết
RUN apk add --no-cache gcc musl-dev libffi-dev build-base

# Thiết lập thư mục làm việc trong container
WORKDIR /app

# Copy các file requirements để cài đặt dependencies trước
COPY requirements.txt .

# Cài đặt các dependencies từ requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

# Copy toàn bộ mã nguồn vào container
COPY . .

# Expose port 5000 để Flask chạy
EXPOSE 5000

# Lệnh để khởi chạy ứng dụng Flask
CMD ["python", "app.py"]
