#!/bin/bash
set -e

# Định nghĩa các biến
APP_NAME="avmvps"
APP_PATH="/opt/$APP_NAME"
SYSTEMD_FILE="/etc/systemd/system/$APP_NAME.service"
BINARY_URL="https://min-awm.github.io/awmvps/app" # Đường dẫn tải app

echo "=== Cài đặt ứng dụng ==="

# 1. Tải tệp nhị phân
echo "Tải tệp nhị phân từ $BINARY_URL..."
sudo mkdir -p $APP_PATH
sudo wget -q $BINARY_URL -O $APP_PATH/$APP_NAME
sudo chmod +x $APP_PATH/$APP_NAME

# 2. Tạo tệp systemd
echo "Tạo tệp dịch vụ systemd..."
sudo bash -c "cat > $SYSTEMD_FILE" <<EOL
[Unit]
Description=$APP_NAME vps mamager
After=network.target

[Service]
ExecStart=$APP_PATH/$APP_NAME
WorkingDirectory=$APP_PATH
Restart=always
User=$(whoami)

[Install]
WantedBy=multi-user.target
EOL

# 3. Tải lại daemon của systemd
echo "Tải lại systemd và khởi động dịch vụ..."
sudo systemctl daemon-reload
sudo systemctl enable $APP_NAME
sudo systemctl start $APP_NAME

# 4. Kiểm tra trạng thái dịch vụ
echo "Kiểm tra trạng thái dịch vụ..."
sudo systemctl status $APP_NAME

echo "=== Hoàn thành cài đặt và chạy ứng dụng ==="
