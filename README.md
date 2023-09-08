[JLS](https://github.com/JimmyHuang454/JLS) 是一个 Fake TLS，对 TLS 最小修改，完全无缝替换安全层。

当检测到是有效用户，那么就会进行代理操作。非有效用户，就会将全部流量转发到指定地址。

无需用户配置证书，做到开箱即用，又不降低安全性。

## 文档

配置跟原版 sing-box 一致，有 TLS 的协议都支持使用 JLS，包括 Tuic 和 Hysteria。

https://sing-box.sagernet.org

## 快速上手 JLS

### Server

假设使用 Ubuntu 系统 x64 架构，一条代码快速安装（要求有 apt 或 yum）：

```bash
bash <(curl https://raw.githubusercontent.com/JimmyHuang454/sing-box/dev-next/release/server/quic_install.sh)
```

安装成功后，会得到一串密码和随机数，端口默认443。 记得防火墙放行 TCP 443 和 UDP 443。

类似这样，记得保存好。

```bash
....
password: 123124293489023745908237
random: 123124293489023745908237
.....
```

查看sing-box状态：

```bash
systemctl status sing-box
```

停止运行：

```bash
systemctl stop sing-box
```

### Client

假设使用 Windows 系统 x64 架构，到 [下载页面](https://github.com/JimmyHuang454/sing-box/releases/latest) 下载最新版的 `sing-box-windows-amd64.zip`：
解压后，修改配置文件的密码，类似这样的：

```json5
{
  "type": "tuic",
  "tag": "tuic",
  "server": "0.0.0.0", // 修改成你 VPS 的 IP 地址
  "server_port": 443, // 服务端默认为443
  "tls": {
    "enabled": true,
    "server_name": "www.apple.com", // 伪装域名
    "jls": {
      "enabled": true,
      "random": "123456", // 密码和随机数必须要填对
      "password": "123456"
    }
  },
  "network": "tcp",
  "password": "password", // 可以不修改，但要填对
},
```

最后运行：

```bash
./sing-box run -c "./config.json"
```

默认HTTP代理端口为8088，设置好本机代理即可使用

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
