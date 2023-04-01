### 简介

Nacos 未授权利用工具

Nacos <= 2.2.0 时存在JWT签名密码硬编码问题，可通过默认密钥自行签发JWT，从而绕过认证

另外 Nacos 在 1.4.1 之后新增通过认证头来鉴权，在开启鉴权模式下，可通过添加认证头 `serverIdentity: security` 绕过鉴权

该工具通过自行签发JWT实现绕过鉴权，签名密钥可自定义，如果未提供密钥则使用默认密钥生成JWT，当JWT生成失败时会尝试使用认证头方式绕过鉴权

目前支持的功能如下

- 添加用户并给予权限
- 查看所有 namespace
- 查看指定 namespace 中的配置文件
- 导出配置文件

### 使用帮助

```
Usage:
  nacos [command]

Get or export config from nacos command
  export      Export config from nacos
  get         Get config from nacos
  list        List all namespaces except public

Add user to nacos command
  adduser     Add user

Additional Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help            help for nacos
  -k, --key string      JWT sign key
      --proxy string    socks or http proxy, eg: socks5:127.0.0.1:8080
  -t, --target string   nacos address, eg: http://127.0.0.1:8848

Use "nacos [command] --help" for more information about a command.
```

子命令使用详情请使用 `nacos [command] --help` 获取，需要说明的是

- `export` 未指定 namespace 时会默认导出所的配置文件
- `get` 默认读取 public 中的配置文件
- `adduser` 在未提供用户名或密码时会采用随机值

### 存在的问题

- 由于默认密钥长度问题，对默认密钥 base64 解码时会出现 error，尝试使用了 `StdEncoding` 和 `RawStdEncoding` 两种方式，都存在问题，虽然也能获取到正确的字节数组，但会返回 error 终究会让人心存芥蒂
- 由于本人水平实在有限，不知道怎样才能在解析完参数后再对 client 做初始化，所以只能在每个子命令执行业务逻辑前都先执行一遍初始化 client 的函数，略显呆板

如果对以上问题有好的解决方式，，还请提 issue 或 PR 😘