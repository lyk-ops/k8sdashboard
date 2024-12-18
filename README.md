# kubeimooc

imooc k8s 课程学习

开发环境说明：
    1. go语言环境 go1.22.5 windows/amd64
    2. 

## 项目初始化

### web框架的选型
```bash
go get -u github.com/gin-gonic/gin
```
>参考文档：https://github.com/gin-gonic/gin

```bash
go get -u github.com/spf13/viper
```
>参考文档：https://github.com/spf13/viper

### K8s集成
```bash
go get -u k8s.io/client-go
```
>参考文档：https://github.com/kubernetes/client-go
## 项目接口开发
### Pod管理接口开发
- 命名空间列表接口 
- Pod创建
- Pod编辑(更新/升级)
- Pod查看（详情、列表）
  - 展示podrequest数据，用于重新创建
- Pod删除

接口调优
1.pod更新会多出来一个挂载卷
2.更新pod执行的步骤多，存在超时情况
3.Pod列表支持关键字搜索

### NodeScheduling接口开发
- Node列表接口/Node详情列表
- Node标签管理
- Node污点管理
- 查看node上所有的pod

pod管理接口改动
- pod新增容忍参数 tolerations
- Pod选择哪种方式调度 NodeName NodeSelector NodeAffinity

### 应用与配置分离接口
- configMap 新增、删除、修改、查询
- Secret 新增、修改、删除、查询
  - Pod管理接口改动
- 新增ConfigMap和configmapkey
- 新增secret和secretkey
###

### K8s卷管理
PersistentVolume
- 创建
- 删除
- 查看  列表
PersistentVolumeClaim
- 创建
- 删除
- 查看  列表
SotrageClass
- 创建
- 删除
- 查看  列表

pod管理的修改（卷管理）
优化点：
- downward fileRefPath没有显示
- PVC选择PV或者SC只能二选一
- 模糊查询keyword
- PV 显示storage-class字段