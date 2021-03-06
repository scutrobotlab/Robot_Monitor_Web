# 后端接口文档

### 一、全局状态码

| status | 说明 |
| ------ | ---- |
|    0   | 正常 |
|    1   | 缺少参数 |
|   11   | 无法打开串口 |
|   12   | 在未打开串口情况下关闭串口 |
|   13   | 无法关闭串口 |
|   21   | 不支持的JSON |
|   22   | 变量操作时串口错误 |
|   23   | 重复添加变量 |
|   24   | 删除未添加的变量 |

### 二、串口

#### 2.1 获取当前已打开的串口信息
##### 请求地址
|  方法 |    URL     |
|-------|-----------|
| `GET` | `/serial` |
##### 请求参数
无
##### 响应结果
| 参数 | 类型    | 说明  |
|------|--------|-------|
| Name | string | 串口名 |
| BaudRate |  int   | 波特率 |
##### 调用示例
请求示例：  
`GET /serial`  
响应示例：  
```
{
    "Name":"COM3",
    "BaudRate":115200
}
```

#### 2.2获取当前可用的串口列表
|  方法 |      URL        |
|-------|----------------|
| `GET` | `/serial/list` |
##### 请求参数
无
##### 响应结果
|  参数 |     类型      |  说明   |
|-------|--------------|---------|
| Ports | array string | 串口列表 |
##### 调用示例
请求示例：  
`GET /serial/list`  
响应示例：  
```
{
    "Ports":[
        "COM3",
        "COM4"
    ]
}
```

#### 2.3 打开串口
|    方法    |      URL        |
|------------|----------------|
| `GET/POST` | `/serial/open` |
##### 请求参数
| 参数 |  类型   |          说明           |
|------|--------|-------------------------|
| port | string |           串口名         |
| baud |  int   | 波特率（可选，默认115200）|
##### 响应结果
|  参数  | 类型 |  说明  |
|--------|-----|--------|
| status | int | 状态码 |
##### 调用示例
请求示例：  
`GET /serial/open?port=COM3`  
响应示例：  
```
{
    "status":0
}
```

#### 2.4 关闭串口
|  方法 |      URL        |
|-------|----------------|
| `GET` | `/serial/close` |
##### 请求参数
无
##### 响应结果
|  参数  | 类型 |  说明  |
|--------|-----|--------|
| status | int | 状态码 |
##### 调用示例
请求示例：  
`GET /serial/close`  
响应示例：  
```
{
    "status":0
}
```

### 三、变量

**如未特殊说明，所有请求参数均为JSON格式**
#### 3.1 获取当前已添加的变量信息
##### 请求地址
|  方法 |    URL     |
|-------|-----------|
| `GET` | `/variable-read/list` |
##### 请求参数
无
##### 响应结果
| 参数 | 类型    | 说明  |
|------|--------|-------|
| Variables | array struct | 变量列表 |
| Variables[].Board |  int   | 板子代号 |
| Variables[].Name |  string   | 变量名 |
| Variables[].Type |  string   | 变量类型 |
| Variables[].Addr |  int   | 变量地址 |
| Variables[].Data |  float   | 变量值 |
##### 调用示例
请求示例：  
`GET /variable-read/list`  
响应示例：  
```
{
    "Variables":[
        {
            "Board":1,
            "Name":"traceme",
            "Type":"float",
            "Addr":536889920,
            "Data":0
        },
        {
            "Board":1,
            "Name":"count",
            "Type":"int",
            "Addr":‭536890180‬,
            "Data":0
        }
    ]
}
```

#### 3.2 获取支持的变量类型
##### 请求地址
|  方法 |    URL     |
|-------|-----------|
| `GET` | `/variable/types` |
##### 请求参数
无
##### 响应结果
| 参数 | 类型    | 说明  |
|------|--------|-------|
| Types | array string | 变量类型列表 |
##### 调用示例
请求示例：  
`GET /variable/types`  
响应示例：  
```
{
    "Types":[
        "double","float","int","int16_t","int32_t","int64_t","int8_t","uint16_t","uint32_t","uint64_t","uint8_t"
    ]
}
```

#### 3.3 添加新的变量
##### 请求地址
|  方法 |    URL     |
|-------|-----------|
| `POST` | `/variable-read/add` |
##### 请求参数
| 参数 | 类型    | 说明  |
|------|--------|-------|
| Board |  int   | 板子代号 |
| Name |  string   | 变量名 |
| Type |  string   | 变量类型 |
| Addr |  int   | 变量地址 |
##### 响应结果
|  参数  | 类型 |  说明  |
|--------|-----|--------|
| status | int | 状态码 |
##### 调用示例
请求示例：  
`POST /variable-read/add`  
```
{
    "Board":1,
    "Name":"traceme",
    "Type":"float",
    "Addr":536889920
}
```
响应示例：  
```
{
    "status":0
}
```

#### 3.4 删除已添加的变量
##### 请求地址
|  方法 |    URL     |
|-------|-----------|
| `POST` | `/variable-read/del` |
##### 请求参数
| 参数 | 类型    | 说明  |
|------|--------|-------|
| Board |  int   | 板子代号 |
| Name |  string   | 变量名 |
| Type |  string   | 变量类型 |
| Addr |  int   | 变量地址 |
##### 响应结果
|  参数  | 类型 |  说明  |
|--------|-----|--------|
| status | int | 状态码 |
##### 调用示例
请求示例：  
`POST /variable-read/del`  
```
{
    "Board":1,
    "Name":"traceme",
    "Type":"float",
    "Addr":536889920
}
```
响应示例：  
```
{
    "status":0
}
```

#### 3.5 修改变量的值
##### 请求地址
|  方法 |    URL     |
|-------|-----------|
| `POST` | `/variable-modi/mod` |
##### 请求参数
| 参数 | 类型    | 说明  |
|------|--------|-------|
| Board |  int   | 板子代号 |
| Name |  string   | 变量名 |
| Type |  string   | 变量类型 |
| Addr |  int   | 变量地址 |
| Data |  float   | 变量值 |
##### 响应结果
|  参数  | 类型 |  说明  |
|--------|-----|--------|
| status | int | 状态码 |
##### 调用示例
请求示例：  
`POST /variable-modi/mod`  
```
{
    "Board":1,
    "Name":"traceme",
    "Type":"float",
    "Addr":536889920,
    "Data":100
}
```
响应示例：  
```
{
    "status":0
}
```
