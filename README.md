# sing-box

The universal proxy platform.

[![Packaging status](https://repology.org/badge/vertical-allrepos/sing-box.svg)](https://repology.org/project/sing-box/versions)

## Documentation

https://sing-box.sagernet.org

## 快速上手 JLS

### Server
假设使用 Ubuntu 系统 x64 架构，一条代码快速安装（要求有 apt 或 yum）：
```bash
bash <(curl https://raw.githubusercontent.com/JimmyHuang454/sing-box/dev-next/release/server/quic_install.sh)
```

安装成功后，会得到一串密码和随机数，端口默认443。

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
    "server_name": "www.visa.cn", // 伪装域名
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
默认HTTP代理端口为8088，设置好本机代理即可使用

### 推荐伪装站
端口统一为 443；可以为任意网站，可以设置成一些官方政府网站或已备案网站，可能处在白名单列表中，不会影响安全性，用户数据不会发送到伪装站。

用户可以自己去设置喜欢的网站，用 chrome 的开发者工具，查看它所使用的是不是 HTTPS，再去查查 IP 属地。选择困难症？去看看 [这里](https://alexa.chinaz.com/)看看

```bash
maven.apache.org # 泛播 Fastly
www.g2.com # cloudflare
www.legco.gov.hk # 香港
www.gov.hk # 香港
www.hangseng.com # 东京
weibo.cn # 大陆
...
```


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
