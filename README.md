chronos
===
Chronos   Fast and efficient permission validation  (ABAC/RBAC)

### CasbinV2 的ABAC啥玩意  性能如此之低下

### Storage
- [ ] Redis  (分布式实现基于一致性HASH 类似Cassandra)
- [ ] Stele  (支持Raft AND 一致性HASH)

### 进度
- [ ] 基础实现
- [ ] 兼容Casbin API

### Base Config
ABAC
``` 
[request_definition]
r = sub, obj

[policy_definition]
p = name, obj, sub_rule                     # sub_rule: r.sub.Time<1635936065, r.sub.Count<20

[policy_effect]                             # 暂时不适配 
e = some(where (p.eft == allow)) 

[matchers]
m = eval(p.sub_rule) && r.obj == p.obj  && r.sub.Name == p.name
```
RBAC
``` 
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]                              # 暂时不适配 
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
```

### Casbin watch
![](./README/casbin_watchers.png)
Casbin 问题 
- 1.初始化数据必须保持一致
- 2.无法解决脑裂问题

#### 解决办法
- 方案1：(灵感来自Raft协议)
    1. 日志 ，建立每个Casbin实例的 日志， 每修改或者增加 改日志加一
        - 当发生数据同步是发送方需要携带该机日志
        - 同步方进行日志比对
        - 当日志差别非常大时 选用大多数实例快照进行同步
    2. 快照
        - 对当前存储数据进行快照处理，当有设备出现脑裂后恢复进行覆盖同步