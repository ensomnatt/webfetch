#!/bin/bash

set -e 

echo -e "installing webfetch...\n"

if [ -e "/opt/webfetch" ]; then
  echo -e "deleting app directory\n"
  sudo rm -rf /opt/webfetch
  echo -e "deleting systemd file\n"
  sudo rm /etc/systemd/system/webfetch.service 
fi

echo -e "creating app config directory in the config directory\n"
mkdir -p $HOME/.config/webfetch
echo -e "copying basic frontend to the app config directory\n"
cp -r web/* $HOME/.config/webfetch

echo -e "building project\n"
go build -o webfetch main.go
echo -e "making app directory\n"
sudo mkdir -p /opt/webfetch
echo -e "copying binary file to the app directory\n"
sudo cp -r webfetch /opt/webfetch/
echo -e "giving to the binary file execute rights\n"
sudo chmod +x /opt/webfetch/webfetch

echo -e "preparing unit file\n"
sed -i "s/^User=.*/User=$USER/" "webfetch.service"
echo -e "copying unit file\n"
sudo cp -r webfetch.service /etc/systemd/system/webfetch.service
echo -e "restarting systemd. use sudo systemctl enable webfetch, sudo systemctl start webfetch to start a program\n"
sudo systemctl daemon-reload

echo "install finished"
