# cert-operator

## 项目由来

早些时候写过一个kubernetes自定义资源CRD的controller,主要是声明一个证书需求,由控制器自动申请一个证书,并保存为一个secret,项目路径[cert-controller](https://github.com/fanfengqiang/cert-controller)

本项目与上述项目功能一致,只不过是用了一个框架对其做了重构.

| 项目             | cert-controller                                              | cert-operator                                                |
| ---------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| 框架             | [sample-apiserver](https://github.com/kubernetes/sample-controller) | [operator-sdk](https://github.com/operator-framework/operator-sdk) |
| 框架维护者       | kubernetes官方社区                                           | coreOS                                                       |
| 证书申请方式     | acme                                                         | acme                                                         |
| 证书申请软件     | [acme.sh](https://github.com/Neilpang/acme.sh)               | [lego](https://github.com/go-acme/lego)                      |
| 证书软件调用方式 | 容器中安装acme.sh软件+go调用系统命令                         | go调用lego库                                                 |

operator-sdk成熟度较高,且在openshift中有大量应用.本项目摒弃了上个项目调用系统命令的低效调用方式,直接采用golang调用lego(也为golang开发)原生库.

## 项目构建

```bash
mkdir -p $GOPATH/src/github.com/fanfengqiang
cd $GOPATH/src/github.com/fanfengqiang
git clone git@github.com:fanfengqiang/cert-operator.git
operator-sdk build fanfengqiang/cert-operator:v1.0
```

## 使用说明

1. 创建自定义资源CRD

   ```bash
   # 创建CRD资源
   kubectl apply -f deploy/crds/certoperator_v1beta1_cert_crd.yaml
   # 创建控制器所需的SA
   kubectl apply -f deploy/service_account.yaml
   # 创建角色和角色绑定
   kubectl apply -f deploy/clusterrole.yaml
   kubectl apply -f deploy/clusterrole_binding.yaml
   # 创建控制器
   kubectl apply -f deploy/operator.yaml
   ```

2. 编写cert资源清单并应用

   ```bash
   # cat deploy/crds/certoperator_v1beta1_cert_cr.yaml
   apiVersion: certoperator.5ik8s.com/v1beta1
   kind: Cert
   metadata:
     name: www.5ik8s.com
   spec:
     secretName: www.5ik8s.com
     domain: www.5ik8s.com
     email: "1415228958@qq.com"
     validityPeriod: 30
     provider: alidns
     envs:
       ALICLOUD_ACCESS_KEY: "LTAIXXXXXXXXXyzLu"
       ALICLOUD_SECRET_KEY: "iH5sCTf4CzXXXXXXX9GLL2AsLW2"
   ```

   ```bash
   # 应用资源清单
   kubectl apply -f deploy/crds/certoperator_v1beta1_cert_cr.yaml
   ```

3. 参数定义

   | 参数                 | 含义                               |
   | :------------------- | :--------------------------------- |
   | .metadata.name       | cert资源的名字                     |
   | .spec.secretName     | 生成的secret的名字                 |
   | .spec.domain         | 生成证书的域名                     |
   | .spec.validityPeriod | secret的有效时长，单位天，范围1~89 |
   | .spec.provider       | 域名托管商的标示                   |
   | .spec.envs           | 域名托管商API的accesskey和secret   |

完整域名托管商的格式，accesskey和secret格式[参见](https://go-acme.github.io/lego/dns/)

## 致谢

本项目参考了如下两个项目

- [memcached-operator](https://github.com/operator-framework/operator-sdk-samples/tree/master/memcached-operator)
- [lego](https://github.com/go-acme/lego)
