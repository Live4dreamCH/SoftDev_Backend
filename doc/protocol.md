# 前后端网络接口

## 简介

网络协议使用HTTP，协议内容使用JSON。

## 已知缺陷

1. 此文档中部分次要功能暂不实现，保证项目按时完成
<!-- 2. 此文档中暂未实现“活动收集何时停止”，这使得活动在数据库中永远不会终止 -->

## 功能列表

### 注册

注册时客户端（后简称C）发送用户昵称、密码，服务器端（后简称S）返回用户的Uid（账号）。
HTTP路由为POST /SignUp

用户昵称是字符串，不超过10个utf-8字符，允许重名；
密码是字符串，不超过30位，以上长度由客户端验证；
返回的Uid是一个不超过int32范围的正整数。

#### 请求

```json
{
    "UserName": "Tom 李",
    "Psw": "password123456"
}
```

#### 响应

```json
{
    "Uid": 10086
}
```

### 登录

登录时C发送Uid和密码，如果服务器验证失败，则返回一个错误标识；如果成功，则返回一个成功标识和一个SessionID。
HTTP路由为POST /LogIn

SessionID是一个随机字符串，由一个随机int经过BASE64编码后得到。
每次登陆成功后由S端返回，作为“本次登录成功”的凭证。
在此之后所有需要“登录后进行”的操作，都要附上这个字段。
当此用户发生下线，或重新登录成功（不管是本地还是异地）后，原有的SessionID失效。若再使用原有的SessionID进行身份相关的操作，S端会返回“SessionID错误”的字段。此时C端应提醒用户重新登录，获取新的正确SessionID。
总而言之，同一时间，一个Uid最多对应一个可用的SessionID。

#### 请求

```json
{
    "Uid": "10086",
    "Psw": "password123456"
}
```

#### 响应

**失败：**

```json
{
    "Res": "NO",
    "Reason": "Uid/Psw"
}
```

"Reason"项为"Uid"时表示此Uid不存在；为"Psw"时表示Uid存在，但此密码与Uid不对应。

**成功：**

```json
{
    "Res": "NO",
    "SessionID": "kzR__MA4E17ebfjjuic7M0GdSf_j8I-VXjGO0owm1MU="
}
```

### 修改用户名、密码（暂不实现）

### 下线（暂不实现）

下线时C端发送Uid和SessionID。若二者均正确且互相对应，则使此SessionID失效，本次登陆状态结束；否则，返回错误原因。
HTTP路由为POST /LogOut

即使下线发生错误，下线也应顺利完成（因为此错误多为异地登陆造成）。如果C端有日志，可以记录下来。

#### 请求

```json
{
    "Uid": "10086",
    "SessionID": "kzR__MA4E17ebfjjuic7M0GdSf_j8I-VXjGO0owm1MU="
}
```

#### 响应

**失败：**

```json
{
    "Res": "NO",
    "Reason": "Uid/SessionID"
}
```

"Reason"项为"Uid"时表示此Uid不存在；为"SessionID"时表示Uid存在，但此SessionID与此Uid不对应。

**成功：**

```json
{
    "Res": "OK"
}
```

### 活动发起者 创建活动

C端发送活动信息，S端返回活动号。
将活动发起者称为organizer，简称Org；活动参与者称为participant，简称Part；活动称为activity，简称Act；会议号ActID，具有较强随机性。
HTTP路由为POST /CreateAct

#### 请求

```json
{
    "SessionID": "kzR__MA4E17ebfjjuic7M0GdSf_j8I-VXjGO0owm1MU=",
    "ActName": "讨论",
    "Length": 2,
    "Description": "中三-1307，软工小组第一次讨论，讨论完成项目立项报告",
    "OrgPeriods":[
        "2021-5-16 14:00",
        "2021-5-16 19:30"
    ]
}
```

"ActName"不超过10个字符，"Description"不超过50个字符；
"Length"表示活动持续几个“0.5小时”，如2表示持续1小时；
"OrgPeriods"是发起者的可用时间的列表。时间字符串的格式见上例；此时间表示每段可用时间的起点。
如"2021-5-16 14:00"，结合"Length": 2，可知发起者在2021-5-16 14:00到2021-5-16 15:00之间有空。

#### 响应

**失败：**

```json
{
    "Res": "NO",
    "Reason": "SessionID/error"
}
```

"Reason"项为"SessionID"时表示此SessionID无效；其余则为具体错误。

**成功：**

```json
{
    "Res": "OK",
    "ActID": 100236
}
```

### 活动发起者 撤销活动

C端发送活动号，S端返回撤销结果。

HTTP路由为POST /DropAct

#### 请求

```json
{
    "SessionID": "kzR__MA4E17ebfjjuic7M0GdSf_j8I-VXjGO0owm1MU=",
    "ActID": 100236
}
```

#### 响应

**失败：**

```json
{
    "Res": "NO",
    "Reason": "SessionID/ActID/Auth"
}
```

"Reason"项为"SessionID"时表示此SessionID无效；为"ActID"时表示无此活动号；为"Auth"时表示此活动的发起者不是此用户，无权撤销活动。

**成功：**

```json
{
    "Res": "OK"
}
```

### 活动发起者 查询实时投票结果

C端发送活动号，S端返回参与者们的实时投票结果。

HTTP路由为POST /OrgGetAct

#### 请求

```json
{
    "SessionID": "kzR__MA4E17ebfjjuic7M0GdSf_j8I-VXjGO0owm1MU=",
    "ActID": 100236
}
```

#### 响应

**失败：**

```json
{
    "Res": "NO",
    "Reason": "SessionID/ActID/Auth"
}
```

"Reason"项为"SessionID"时表示此SessionID无效；为"ActID"时表示无此活动号；为"Auth"时表示此活动的发起者不是此用户，无权查询实时投票结果。

**成功：**

```json
{
    "Res": "OK",
    "AvlbPeriods":[
        "2021-5-16 14:00",
        "2021-5-16 19:30"
    ]
}
```

### 活动发起者 修改活动（暂不实现）

C端发送活动号、修改后的活动信息，S端返回修改结果。
HTTP路由为POST /UpdateAct

#### 请求

```json
{
    "SessionID": "kzR__MA4E17ebfjjuic7M0GdSf_j8I-VXjGO0owm1MU=",
    "ActID": 100236,
    "ActName": "讨论",
    "Description": "中三-1307，软工小组第二次讨论，讨论完成项目立项报告",
}
```

暂不允许修改活动的持续时间和发起者可用时间（这样可能导致修改前的投票作废）；如果需要修改，请发起者删除、新建一个活动。

#### 响应

**失败：**

```json
{
    "Res": "NO",
    "Reason": "SessionID/ActID/Auth"
}
```

"Reason"项为"SessionID"时表示此SessionID无效；为"ActID"时表示无此活动号；为"Auth"时表示此活动的发起者不是此用户，无权修改活动。

**成功：**

```json
{
    "Res": "OK"
}
```

### 活动参与者 参加活动、提交时间段

### 活动参与者 修改提交的时间段（暂不实现）

### 活动参与者 退出活动（暂不实现）

### 活动参与者 查询活动状态（暂不实现）

### 查询自身所在的活动列表

C端发送SessionID，S端返回此用户所在的活动列表，以及自身在各活动中的角色。

HTTP路由为POST /GetActs

#### 请求

```json
{
    "SessionID": "kzR__MA4E17ebfjjuic7M0GdSf_j8I-VXjGO0owm1MU=",
}
```

#### 响应

**失败：**

```json
{
    "Res": "NO",
    "Reason": "SessionID"
}
```

"Reason"项为"SessionID"时表示此SessionID无效。

**成功：**

```json
{
    "Res": "OK",
    "Acts": [
        {
            "ActID": 100236,
            "ActName": "讨论",
            "OrgName":"Tom",
            "Length": 2,
            "Description": "中三-1307，软工小组第一次讨论，讨论完成项目立项报告",
            "OrgPeriods": [
                "2021-5-16 14:00",
                "2021-5-16 19:30"
            ]
        },
        {
            "ActID": 100623,
            "ActName": "讨论",
            "OrgName":"Jerry",
            "Length": 2,
            "Description": "中三-1307，软工小组第二次讨论，讨论完成项目立项报告",
            "OrgPeriods": [
                "2021-5-16 14:00",
                "2021-5-16 19:30"
            ]
        }
    ]
}```