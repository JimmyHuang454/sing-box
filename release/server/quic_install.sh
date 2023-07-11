#!/usr/bin/env bash

# bash <(curl https://raw.githubusercontent.com/JimmyHuang454/sing-box/dev-next/release/server/quic_install.sh)
quit(){
  echo ""
  echo "Failed:"
  echo $1
  echo ""
  exit 1
}

check_systemd() {
  if ! [ -x "$(command -v systemctl)" ]; then
    quit 'Missing "systemd"'
  fi
}
check_systemd

if [ -x "$(command -v yum)" ]; then
  InstallType="rpm"
elif [ -x "$(command -v apt)" ]; then
  InstallType="deb"
else
  quit 'Missing "yum" or "apt"'
fi

ARCH_NAME="sing-box-linux-amd64.$InstallType"

download_lastest_singbox_to() {
  DOWNLOAD_LINK="https://github.com/JimmyHuang454/sing-box/releases/latest/download/$ARCH_NAME"
  echo "Downloading from: $DOWNLOAD_LINK"
  if ! curl -R -L -H 'Cache-Control: no-cache' -o "$1"  "$DOWNLOAD_LINK"; then
    quit 'Download failed! Please check your network or try again.'
  fi
}


SAVE_PATH="./$ARCH_NAME"
CONFIG_PATH="/etc/sing-box/config.json"
PASSWORD="$RANDOM$RANDOM$RANDOM$RANDOM$RANDOM$RANDOM$RANDOM$RANDOM"
IV="$RANDOM$RANDOM$RANDOM$RANDOM$RANDOM$RANDOM$RANDOM$RANDOM"

sudo systemctl stop sing-box

download_lastest_singbox_to $SAVE_PATH

if InstallType=="deb"; then
  sudo apt remove sing-box
  if ! sudo apt install $SAVE_PATH; then
    quit "Failed to install by apt."
  fi
else
  sudo yum remove sing-box
  if ! sudo yum localinstall $SAVE_PATH; then
    quit "Failed to install by yum."
  fi
fi



if ! sed -i "s/123456/$PASSWORD/" $CONFIG_PATH; then
  quit "Failed to generate password and random number."
fi
sed -i "s/abcabc/$IV/" $CONFIG_PATH

if ! sudo systemctl start sing-box; then
  quit "Failed to execute systemctl"
fi

echo ""
echo "password: $PASSWORD"
echo "random:   $IV"
echo " "
