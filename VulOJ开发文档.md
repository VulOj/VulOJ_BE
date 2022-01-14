# 基本要求



## 业务功能

1. 用户登录注册
2. 用户创建班级，邀请成为班级成员
3. 发帖

- 首页

展示帖子若干条，其中携带组织的名称和作者的信息。点击可跳转

- 该项目中，以组织为单位，组织分为`大创项目`以及 `企业组织`

- 用户可以创建 `组织`，作为 `组织`的 `所有者` 也可以邀请某位用户，使该用户成为该组织的 `成员`

- `组织`的 `所有者` 和 `成员`都可以 `发帖` ，其中所有者可以编辑组织的 `状态`以及组织的 `profile`

  这里的设计可以参照知乎，每个人都可以提出一个问题（如何看待xxx）。对应到本项目即一个组织。组织的创建者需要邀请用户加入其中，被邀请且接受的用户才可以在这个组织下发帖。

  然后创建者可以编辑组织的一些信息（像知乎的提问者可以修改问题描述一样）

# 技术栈环境的要求

- 数据库使用 `MYSQL 5.7` 版本及以上
- `Go` 语言使用 `1.15.0` 版本及以上，后端开发使用 `Gin` 框架。

## 数据库设计

使用 `Go` 的 `Gin` 框架完成对数据库的操作，故只规定go中结构体即可。

- `Auth_User`表

```json
Email                string `gorm:"primary_key"`
Username             string
Password             string
Register_timestamp   int64
Last_login_timestamp int64
Role                 string
```

- `Team`表

```json
gorm.Model
Name          string
Owner_email   string
Description   string
```

- `Invite_Record表`

```jsx
gorm.Model
Team_Id	  		int
MemberInvited 	string
Is_accepted     int
```

- `TeamMember` 表

```json
gorm.Model
Team_Id uint
Email   string
```

- `Blog` 表

```jsx
gorm.Model
Team_Id    uint
Auth_Email string
Title      string
Content    string
```

# Api设计

## 1. 邀请用户（/api/inviting)

沈博宇

- post方法，邀请某个用户加入某个组织

### /invite

post方法，前端传入:

```json
{
	"team_id":"123"
	"email_accept":"456@"
	"description":"blabla"
}
```

后端返回：

```json
{
	"msg":"balabala"
}
```

### /accept

- post方法（/accept)，用户接受邀请

前端输入：

```jsx
{
	"Owner_email":"123@"
	"team_id":"123"
	"email_accept":"456@"
//都从url里面来
}
```

后端返回：

```jsx
{
	"msg":"balabala"
}
```

## 2. 组织模块（/api/team)

谭骏

### post方法，前端传入

```json
{
	"team_name":"team name"
	"description":"team description"
	"mile_stone":"项目进度"
	
}
```

后端返回

```jsx
{
	"msg":"balabala"
}
```

### delete方法，前端传入

```json
{
	"team_id":"team id"
}
```

后端返回

```json
{
	"msg":"balabala"
}
```

### put方法，前端传入

```json
{
	"team_id":"team id"
	"new_team_name":"new team name"
	"mile_stone_update":"项目进度更新"
	"description_update":"team description update"
	/*如果用户未改变对应内容则将原来的内容传入*/
}
```

后端返回

```json
{
	"msg":"balabala"
}
```

### get方法，用来获取team的列表，前端传入

```json
{
	"start_id":"start id"
	"list_size":"list size"

}
```

后端返回

```json
{
	{
		"teams":[
			"TeamName" :          "team name",
			"Username":	      "user name",
			"Description":        "team description",
			"Mile_Stone":	      "mile stone",
		],
        }
}
```

## 3. 发帖模块（/api/blog)

郭玺

### Get方法 获取到组织内的blog

```json
/*前端传入*/
{
		"team_id":"队伍ID",
		"start_id":"开始的文章下标（从0开始计数"
		"list_size":"列表大小"
		"is_descend":"是否倒序"
}
/*返回*/
{
		"msg":"文章列表"
}
```

### Post方法，创建一个文章（必须要自己所属的组织内）

```json
/*前端传入*/
{
		"title":"文章标题"
		"content":"文章内容"
		"team_id":"队伍ID"
}
/*后端返回*/
{
		"msg":"......"
}
```

### Delete方法，删除一个文章

```json
/*只删除自己发布的*/
/*前端传入*/
{
		"blog_id":"文章的ID"
}
/*返回*/
{
		"msg":"..."
}
```

## 4. **认证模块（/api/auth)**

### /sendVerifyCode

post方法

```json
{
     "email":"your email"
}
```

后端返回

```json
如果成功
{
     "msg" :"balbal"
}
如果失败
{
     "msg" :"balbal"
}
```

### /isSessionExpired

post方法：

```json
{
		数据放在headers里面的Cookie里面
}
```

后端返回：

```json
{
		//如果过期就返回true
		"is_expired":false
}
```

### /verifyCodeMatch , 用来验证验证码是否正确

post 方法，前端传入：

```json
{
		"email":"your email",
		"verify_code":"balbal"
}
```

后端返回：

```json
{
	"is_match":true
}
```

### /login  ,认证登录

```json
{
		"email":"your email",
		"password":"balbal"
}
```

后端返回

```json
{
  "msg":"blabla"
}
状态码返回：
{
		200：请求成功，
		400：登陆失败或者邮箱不正确
}
```

### /register:注册账号

前端输入

```json
{
		"email":"your email",
		"password":"balbal",
		"verify_code":"balbal"
}
```

后端返回：

```json
{
  "msg": "blabla"
}
状态码返回：
{
		200：请求成功，
		400：登陆失败或者邮箱不正确
}
```

### /signout

后端返回:

```json
{
    "msg": "blabla"
}
```

### /changePassword

前端输入：

```json
{
		"email":"your email",
		"new_password":"balbal",
		"verify_code":"balbal"
}
```

后端返回:

```json
{
    "msg": "blabla"
}
```

## 5. 个人信息（/api/myself)

### get方法，用来获取我的组织,前端传入

```json
{
	"email":"用户邮箱"
}
```

后端返回

```json
{
	{
		"teams":[
			"TeamName" :          "team name",
			"Description":        "team description",
			"Mile_Stone":	      "mile stone",
		],
        }
}
```

## 6. 首页内容（/api/homepage)

### get方法，用来获取首页内容

## 7.大创项目（/api/project）

### get方法，获取大创项目首页内容

### /college，获取对应学院简介及其下的大创项目

### /hot,获取热门项目

### /latest，获取最新项目

# **8.校园企业（/api/company）**

/companycategories：GET方法获取企业分类

/hotcompanies：GET方法获得企业列表

/newcompanies：GET方法获取最新企业

/companies：GET方法获取指定分类企业列表

/companyinfo：GET方法获取企业简介信息

/companyblogs：GET方法获取企业文章列表

/companyblog：GET方法获取企业文章内容