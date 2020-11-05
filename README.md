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
``` 
[request_definition]
r = sub, obj

[policy_definition]
p = name, obj, sub_rule

[policy_effect]                             # 暂时不适配 原谅我编译原理挂科了
e = some(where (p.eft == allow))

[matchers]
m = eval(p.sub_rule) && r.obj == p.obj  && r.sub == p.name
```