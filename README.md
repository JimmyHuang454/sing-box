# sing-box

The universal proxy platform.

[![Packaging status](https://repology.org/badge/vertical-allrepos/sing-box.svg)](https://repology.org/project/sing-box/versions)

## Documentation

https://sing-box.sagernet.org

## 快速上手 JLS

### Server
假设使用 Ubuntu 系统 x64 架构，到 [下载页面](https://github.com/JimmyHuang454/sing-box/releases) 下载最新版的 Linux 安装包：
```bash
# 下载最新安装包
wget https://github.com/JimmyHuang454/sing-box/releases/latest/download/sing-box-linux-amd64.deb

# 安装
apt install ./sing-box-linux-amd64.deb

# 必须修改密码
vi /etc/sing-box/config.json

# 最后运行
systemctl start sing-box
```

### Client

假设使用 Windows 系统 x64 架构，到 [下载页面](https://github.com/JimmyHuang454/sing-box/releases) 下载最新版的 `sing-box-windows-amd64.zip`：
解压后，修改配置文件的密码，类似这样的：
```json
{
  "type": "trojan",
  "tag": "trojanOut",
  "server": "0.0.0.0", // 修改成 VPS 的 IP 地址
  "server_port": 443, // 服务端默认为443
  "tls": {
    "enabled": true,
    "server_name": "www.visa.cn",
    "jls": {
      "enabled": true,
      "random": "123456", // 密码和随机数必须要填对
      "password": "123456"
    }
  },
  "network": "tcp",
  "password": "abcd" // 可以不修改，但要填对
},
```

最后运行：
```bash
./sing-box run --config "./config.json"
```
默认HTTP代理端口为8080，设置好本机代理即可使用


## License

```
Copyright (C) 2022 by nekohasekai <contact-sagernet@sekai.icu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.

In addition, no derivative work may use the name or imply association
with this application without prior consent.
```
