#!/bin/bash
set -e

# Define variables
APP_NAME="avmvps"
APP_PATH="/opt/$APP_NAME"
SYSTEMD_FILE="/etc/systemd/system/$APP_NAME.service"
BINARY_URL="https://min-awm.github.io/awmvps/app" # URL to download the app

echo "=== Installing AWMVPS ==="

# 1. Download the binary file
echo "Downloading binary file from $BINARY_URL..."
sudo mkdir -p $APP_PATH

if command -v wget > /dev/null 2>&1; then
    sudo wget -q $BINARY_URL -O $APP_PATH/$APP_NAME
elif command -v curl > /dev/null 2>&1; then
    sudo curl -s $BINARY_URL -o $APP_PATH/$APP_NAME
else
    echo "wget or curl is required to proceed." >&2
    exit 1
fi

sudo chmod +x $APP_PATH/$APP_NAME

# 2. Create the systemd service file
echo "Creating systemd service file..."
sudo bash -c "cat > $SYSTEMD_FILE" <<EOL
[Unit]
Description=$APP_NAME VPS Manager
After=network.target

[Service]
ExecStart=$APP_PATH/$APP_NAME
WorkingDirectory=$APP_PATH
Restart=always
User=$(whoami)

[Install]
WantedBy=multi-user.target
EOL

# 3. Reload systemd daemon
echo "Reloading systemd daemon and starting the service..."
sudo systemctl daemon-reload
sudo systemctl enable $APP_NAME
sudo systemctl start $APP_NAME

# 4. Check the service status
echo "Checking the service status..."
sudo systemctl status $APP_NAME

echo "=== Installation and application setup complete ==="
