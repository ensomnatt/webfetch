#!/bin/bash

APP_DIRECTORY=/opt/webfetch
SYSTEMD_DIRECTORY=/etc/systemd/system

echo "creating app directory"
sudo mkdir -p $APP_DIRECTORY

echo "building"
go build -o webfetch cmd/main.go

echo "copying binary file to the app directory"
sudo cp -r webfetch $APP_DIRECTORY

echo "preparing systemd file"
sed -i "s/^User=\$/User=$(whoami)/" webfetch.service

echo "copying systemd file to the systemd directory"
sudo cp -r webfetch.service $SYSTEMD_DIRECTORY

echo "done. run sudo systemctl enable --now webfetch to enable and start the programm"
