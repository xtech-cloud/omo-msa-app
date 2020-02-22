# 简介

应用管理服务


# 编译

进入omo-msa-app目录，执行go build


# 部署

设置环境变量

```
export GIN_MODE=release
export APP_HTTP_ADDR=:80
```

- SQLite

设置以下环境变量

```
export APP_DATABASE_DRIVER=sqlite
export APP_SQLITE_FILEPATH=<db文件所在路径>
```

- Mysql

设置以下环境变量

```
export APP_DATABASE_DRIVER=mysql
export APP_MYSQL_ADDR=127.0.0.1:3306
export APP_MYSQL_USER=<mysyql用户>
export APP_MYSQL_PASSWORD=<mysql密码>
export APP_MYSQL_DATABASE=omo
```

设置完环境变量后启动omo-msa-app


# HTTP API

## app

### `/app/create`

**简要描述:**

- 创建应用
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|appname|是  |string |应用名|

**范例**
> curl -v -H "Content-Type:application/json" -X POST -d '{"appname":"test"}' 127.0.0.1/app/create

**返回示例**

```json
{
    "code":0, 
    "message":"",
    "data":{
        "appkey":"ee488cd5bf1363902a1bcb0359464e48",
        "appsecret":"42f7f2011c1a7a3c7d9ab26332ef99bb",
    }
}
```

**返回参数说明**

|参数名|类型|说明|
|:-----  |:-----|:-----|
| appkey | string | 应用识别码|
| appsecret | string | 应用密钥 |

### `/app/query`

**简要描述:**

- 查询应用
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|appname|是  |string |应用名|

**范例**
> curl -v -H "Content-Type:application/json" -X POST -d '{"appname":"test"}' 127.0.0.1/app/query

**返回示例**

```json
{
    "code":0,
    "message":"",
    "data":{
        "AppName":"test2",
        "AppKey":"bf34b04d9698ca9c95f5def56d960827",
        "AppSecret":"ebe8a84a8708b6cef93d2449bd650910",
        "PrivateKey":"-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAlL7SgYxOYxREy40esuu0jgd2CoDyMRvFRBWpR5he2VQcy8o8\nMSKAfDzH26FVZc9z3VafhD5GVlIO13gBK/336vnN/VBHPQ36F8HGWM7YekQ9+N5Q\ncT4mWqUQsWUatbQFi2C86OpmLGCL3jq2i7S6mxNLJbe7E+dg+cy0SagIlXlbnNXr\ncHnuJoVXBiwhmig1PzHgTRvSjXKXAIjQGWkiCdiyzYx5bif7gNW/ib7PJruXfFLI\n4lF2veaylYcbC/+BY42YNoAWMsGBWnCBLCCLVyKCt698bVftOELUtZOy9+TcZU71\n+pOwPf+eeqhK4pAXDGG1pLz1hBhmVbVlg7Rk0wIDAQABAoIBABYD+Fw8TA3WHiiS\nhoys3lh3Oj1rwG0MUzI0ko2KO9+m12xCTo5nMOUyidI0GtOq1NdZztpf7UExfAjg\nNiwwttUMjDSGAUVEIFQL1jOmyduu5g1DulxIepzH+aSH9mAWeQucEdnXd6/xykHm\nJsaexU/WlzTJ8OKNSIkwhy6vtDWLOwQm+yldYu0TNZJ0jIs1hVjjLsbf/4h6Hk3y\nYIUa1T5NJpOmXULSfgxADNLFFprLNBqFeHarq8u+g38TOLzYC1eF8bsLPWe1Wzoa\nPWCtKcNOvMPPSEjv85KOjvF7Q8OnmEzKhwtgSu8AlxMgVG9lAc8XrWA9S2xSs20l\nJMM3IyECgYEAwoA/moXrevn/SsrQURjdvSjpI1e1C+6Dhu9ivwfMQqdpX5WD23TY\n4f3EWXN8u0Nz5aVHdd7O3bLllZjcUnJay4jFKqNZyjWaGxNmJZDHNth8unlr4fki\nEc3Z9H33FlCQG2/cuq+P7ujV7g1wTAM1yy/IGhUrjyCnme0CE4ZkAK8CgYEAw8br\na9OiSfAVO3thwK0m++BuVC6LqSe0gVazCgJtn0aAgmLeeMhwHoG+7DAtWau0PyYS\nh5UVYF5yaBvX3Zs0ze9wqzMp38954B25nOA3lv1QlnQ1I15aDZlvsreALnhe464d\nJiEfP0rvd4iTiKUKqsOKmtL0bkNl03xtLGXu/x0CgYBcHV0CE7aocUnE5DSwk7RA\nZ+WyRVGLKxTDjRAZJNpKHvs6t5bREo+8x/B75MQH9DQpaJNlcXZLbPRqWxDNQzdY\n+ZdXUDGwIJ6xgAh6dgzDHthDgEnlpZXLFNDKh/XDbbgyJlJFX+ws27yll1u9xC9v\n4VtFbw1IJdD6h1LaaGVoJQKBgQCAbEYV7zev5Kso36CZ8Xt3EhuNYRMAHSmNBkBf\nuoQKTQcTgKOK+4CAon+JE3lMLxQHsIPLKIJjOtE1db4+ggc7Z2uzAdbgF4tM9nLB\nc1tD0ltAtm39C3FrJlFdHH4a/Z7RH2/DiUqkDBXVhWOx6QF8TtTnBqaMhe2Pszky\nPJNwCQKBgDTUcA7sBouf3f23CVY0q073NnT/eDgPrhIe1oc9/SWLnteVRYqc6QCL\nOsXFUzdTLbVskyhuDSaebgIjdVlSY+voxPAMpVyQAXZGXoemKH+oIQpMf8/QBBdi\nlK/ypitZJ9jRXgrXWhie4etiXp+f6UAbTFNQU3W3ZANUtjz9WjT0\n-----END RSA PRIVATE KEY-----\n",
        "PublicKey":"-----BEGIN RSA PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlL7SgYxOYxREy40esuu0\njgd2CoDyMRvFRBWpR5he2VQcy8o8MSKAfDzH26FVZc9z3VafhD5GVlIO13gBK/33\n6vnN/VBHPQ36F8HGWM7YekQ9+N5QcT4mWqUQsWUatbQFi2C86OpmLGCL3jq2i7S6\nmxNLJbe7E+dg+cy0SagIlXlbnNXrcHnuJoVXBiwhmig1PzHgTRvSjXKXAIjQGWki\nCdiyzYx5bif7gNW/ib7PJruXfFLI4lF2veaylYcbC/+BY42YNoAWMsGBWnCBLCCL\nVyKCt698bVftOELUtZOy9+TcZU71+pOwPf+eeqhK4pAXDGG1pLz1hBhmVbVlg7Rk\n0wIDAQAB\n-----END RSA PUBLIC KEY-----\n",
        "Profile":"",
        "CreatedAt": "2019-01-01 13:13:13.22333 +0000 UTC"
    }
}
```

**返回参数说明**

|参数名|类型|说明|
|:-----  |:-----|:-----|
| AppName| string | 应用名|
| AppKey| string | 应用识别码|
| AppSecret| string | 应用密钥 |
| PublicKey| string | 公钥|
| PrivateKey| string | 私钥|
| Proafile| string | 简介|
| CreatedAt| string | 应用创建时间|

### `/app/list`

**简要描述:**

- 列出应用
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
| appname | 是 | string |  应用名 |


**返回示例**

```json
{
    "code":0,
    "message":"",
    "data":{
        "Applications":["test", "test2"]
    }
}
```

**返回参数说明**

|参数名|类型|说明|
|:-----  |:-----|:-----|
| Applications| string[] | 应用名列表|


### `/app/secret/reset`

**简要描述:**

- 重置secret
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|appname|是  |string |应用名|


**返回示例**

```json
{
    "code":0,
    "message":"",
    "data":{
    }
}
```

**返回参数说明**

|参数名|类型|说明|
|:-----  |:-----|:-----|
| | | |

### `/app/rsa/reset`

**简要描述:**

- 重置RSA密钥对
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|appname|是  |string |应用名|


**返回示例**

```json
{
    "code":0,
    "message":"",
    "data":{
    }
}
```

**返回参数说明**

|参数名|类型|说明|
|:-----  |:-----|:-----|
| | | |

#

### `/app/profile/modify`

**简要描述:**

- 更改简介
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|appname | 是| string | 应用名|
|profile| | string | 简介|


**返回示例**

```json
{
    "code":0,
    "message":"",
    "data":{
    }
}
```
