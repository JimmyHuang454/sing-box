#!/usr/bin/env bash
#
quit(){
  echo $1
  exit 1
}

if sudo yum --help; then
  InstallType="rpm"
elif sudo yum --help; then
  InstallType="deb"
else
  quit "unknow arch type."
fi

download_lastest_singbox_to() {
  DOWNLOAD_LINK="https://github.com/JimmyHuang454/sing-box/releases/latest/download/sing-box-linux-amd64.$InstallType"
  echo "Downloading from: $DOWNLOAD_LINK"
  if ! curl -R -L -H 'Cache-Control: no-cache' -o "$1"  "$DOWNLOAD_LINK"; then
    quit 'error: Download failed! Please check your network or try again.'
  fi
}


SAVE_PATH="./sing-box-isntall-arch"

sudo systemctl stop sing-box

download_lastest_singbox_to $SAVE_PATH

if [[ InstallType=="deb" ]]; then
  sudo apt install $SAVE_PATH
else
  sudo yum localinstall $SAVE_PATH
fi

CONFIG_PATH="/etc/sing-box/config.json"
PASSWORD="$RANDOM$RANDOM$RANDOM$RANDOM$RANDOM$RANDOM$RANDOM$RANDOM"
IV="$RANDOM$RANDOM$RANDOM$RANDOM$RANDOM$RANDOM$RANDOM$RANDOM"

sed -i "s/123456/$PASSWORD/" $CONFIG_PATH
sed -i "s/abcabc/$IV/" $CONFIG_PATH

sudo systemctl enable sing-box
sudo systemctl start sing-box

echo ""
echo "password: $PASSWORD"
echo "random:   $IV"
echo " "
