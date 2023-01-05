# idgener

用于生成随机不(难以)重复id的函数.

目前支持4种算法

+ uuid4
+ sonyflake
+ snowflake
+ ulid

## 用法

+ 使用默认生成器--调用`func Next(idgen_enum IDGENAlgorithm) (string, error)`

+ 创建生成器--调用`func IDGenNameToIDGen(idgen_enum IDGENAlgorithm) (IDGenInterface, error)`
