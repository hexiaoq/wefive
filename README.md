# 政府部门接口文档

## 基础url地址: http://121.37.152.118:7777/

### 1. 根据部门Id获取业务 businesses/get
method:
```
POST
```
发送请求(deptId 发送时必须为字符串类型):
```
{
    "deptId": "4"
}
```
返回数据：
```
{
    "msg": "获取业务成功！",
    "code": 200,
    "data": [
        {
            "bus_id": 3,
            "dept_id": 4,
            "bus_name": "身份证办理",
            "description": "携带有关证件进行办理", 
            "requirement": "中国公民"
        },
        {
            "bus_id": 4,
            "dept_id": 4,
            "bus_name": "办理房产证",
            "description": "由工作人员取票排队等候准备好相关材料", 
            "requirement": "需要房产证明"
        }
    ]
}
```

### 2. 根据业务Id获取材料 material/get
method:
```
POST
```
发送请求(busId 发送时必须为字符串类型):
```
{
    "busId": "4"
}
```
返回数据：
```
{
    "msg": "获取业务材料成功！",
    "code": 200,
    "data": [
        {
            "bus_id": "3",
            "description": "本人一寸免冠照片",
            "material_id": "12",
            "material_name": "证明材料",
            "photo_url": ""
        },
        {
            "bus_id": "3",
            "description": "需要在有效期内的房产证明",
            "material_id": "13",
            "material_name": "房产证",
            "photo_url": ""
        }
    ]
}
```

### 3. 获取热门业务 businesses/getHot
method:
```
GET
```
返回数据：
```
{
    "msg": "获取业务材料成功！",
    "code": 200,
    "data": [
        
       
    ]
}
```
