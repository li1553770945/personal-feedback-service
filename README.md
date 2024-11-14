# sheepim提建议

## 初始化项目
```bash
export GOPROXY=direct
export HTTP_PROXY=http://localhost:7890
export HTTPS_PROXY=http://localhost:7890
kitex -module "github.com/li1553770945/personal-feedback-service" -service personal-feedback-service idl/feedback.thrift
cd biz/infra/container
wire
```
## 配置文件示例

```yml
server:
  listen-address: 0.0.0.0:8891
  service-name: personal-feedback-service

etcd:
  endpoint:
    - 127.0.0.1:2379

open-telemetry:
  endpoint: 127.0.0.1:4317

database:
  username: xxx
  password: xxx
  database: xxx
  address: xxx
  port: xxx

```

## 开发环境

```bash
export ENV=development
```

## 生产环境

```bash
export ENV=production
```