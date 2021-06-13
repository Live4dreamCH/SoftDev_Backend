# 前后端网络接口

## 简介

网络协议使用HTTP，协议内容使用JSON。

<!-- ## 已知缺陷

1. 此文档中部分次要功能暂不实现，保证项目按时完成 -->

## 功能列表

### 注册

注册时客户端（后简称C）发送用户昵称、密码，服务器端（后简称S）返回成功与否。
HTTP路由为POST /SignUp

用户昵称是字符串，不超过10个utf-8字符，不允许重名；
密码是字符串，不超过30位，以上长度由客户端验证。

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
    "Res": "OK/NO"
}
```

"Res": "NO"表示此昵称已被注册。

### 登录

登录时C发送用户昵称以及密码。如果服务器验证失败，则返回一个错误标识；如果成功，则返回一个成功标识和一个SessionID。
HTTP路由为POST /LogIn

SessionID是一个随机字符串，由一个随机int经过BASE64编码后得到。
每次登陆成功后由S端返回，作为“本次登录成功”的凭证。
在此之后所有需要“登录后进行”的操作，都要附上这个字段。
当此用户发生下线，或重新登录成功（不管是本地还是异地）后，原有的SessionID失效。若再使用原有的SessionID进行身份相关的操作，S端会返回“SessionID错误”的字段。此时C端应提醒用户重新登录，获取新的正确SessionID。
总而言之，同一时间，一个用户最多对应一个可用的SessionID。

#### 请求

```json
{
    "UserName": "Tom 李",
    "Psw": "password123456"
}
```

#### 响应

**失败：**

```json
{
    "Res": "NO",
    "Reason": "UserName/Psw"
}
```

"Reason"项为"UserName"时表示昵称不存在；为"Psw"时表示昵称存在，但此密码与昵称不对应。

**成功：**

```json
{
    "Res": "YES",
    "SessionID": "kzR__MA4E17ebfjjuic7M0GdSf_j8I-VXjGO0owm1MU="
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
    ],
    "Votes":[
        12, 56
    ]
}
```

### 活动发起者 敲定活动时间，结束投票

活动发起者在几个时间段中最终选择一个时间段，作为活动的时间；选择之后，此活动不允许新的参与者再进行投票。
活动的几个时间段就是发起者创建活动时的OrgPeriods，可以使用下文中的“查询自身所在的活动列表”/GetActs功能获得。
投票结束后不会立即在服务器端数据库里删除此活动，因为还要向参与者展示最终结果。

C端发送SesionID、活动号和活动发起者最终选择的时间段的起点时间。
S端返回操作是否成功。
HTTP路由为POST /StopAct

#### 请求

```json
{
    "SessionID": "kzR__MA4E17ebfjjuic7M0GdSf_j8I-VXjGO0owm1MU=",
    "ActID": 100236,
    "FinalPeriods": "2021-5-16 14:00"
}
```

#### 响应

**失败：**

```json
{
    "Res": "NO",
    "Reason": "SessionID/ActID/Periods/Auth"
}
```

"Reason"项为"SessionID"时表示此SessionID无效；
为"ActID"时表示无此活动号；
为"Periods"时表示时间格式不对/时间不在发起者规定的范畴内；
为"Auth"时表示此活动的发起者不是此用户。

**成功：**

```json
{
    "Res": "OK"
}
```

### 活动参与者 参加活动、提交时间段

C端发送SesionID、活动号和此用户的空闲时间段的起点时间，S端回复成功或失败。

HTTP路由为POST /PartAct

#### 请求

```json
{
    "SessionID": "kzR__MA4E17ebfjjuic7M0GdSf_j8I-VXjGO0owm1MU=",
    "ActID": 100236,
    "PartPeriods":[
        "2021-5-16 14:00",
        "2021-5-16 19:30"
    ]
}
```

#### 响应

**失败：**

```json
{
    "Res": "NO",
    "Reason": "ActID/Periods/Stopped"
}
```

"Reason"项为"ActID"时表示无此活动号，为"Periods"时表示时间格式不对/时间不在发起者规定的范畴内，为"Stopped"表示此活动已停止投票。

**成功：**

```json
{
    "Res": "OK"
}
```

### 查询自身所在的活动列表

C端发送SessionID，S端返回此用户所在的活动的列表，以及活动投票是否结束。
由于用户昵称不允许重复，C端可以用自身的昵称与活动昵发起者昵称比较，判断自身在活动中的身份。

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
            ],
            "Stopped": 0
        },
        {
            "ActID": 100623,
            "ActName": "讨论",
            "OrgName":"Jerry",
            "Length": 2,
            "Description": "中三-1307，软工小组第二次讨论，讨论完成项目立项报告",
            "OrgPeriods": [
                "2021-5-16 14:00"
            ],
            "Stopped": 1
        }
    ]
}
```

"Stopped"表示此活动是否已停止投票。1为是，0为否。
若"Stopped"为1，则"OrgPeriods"只会返回一个时间段，即为活动发起者最终敲定的时间段。
服务器不会返回已经发生过的活动。
