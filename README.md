
项目经验不是很丰富，欢迎大家提意见

如果有好的项目实践欢迎大家滴滴我呀

+ 参考了Mikael大佬的[go-zero-looklook](https://github.com/Mikaelemmmm/go-zero-looklook)项目！很好的学习项目！

# 主要技术

+ go-zero
+ mysql —— gorm
+ redis
+ kafka
+ asynq

# 微服务功能实现

## 用户服务

+ 用户注册
+ 用户登录
+ 用户推出
+ 邮箱验证码发送
+ 绑定邮箱
+ 用户基本信息展示
+ 修改邮箱
+ 修改用户基本信息
+ 修改密码
+ 上传头像
+ 获取账户金额
+ 添加收获地址
+ 用户收货地址列表
+ 更新收货地址
+ 删除收货地址

## 商品服务

### 通用商品模块

+ 首页商品轮播图
+ 首页分类名称列表
+ 首页商品推荐列表
+ 商品搜索
+ 商品分类列表
+ 商品详情信息

### 商家商品模块

+ 上传商品
+ 上架商品
+ 修改商品信息
+ 下架商品
+ 删除商品
+ 添加秒杀商品
+ 删除秒杀商品

### 用户商品模块

+ 添加收藏商品
+ 收藏商品列表
+ 删除收藏商品

### 秒杀商品模块

+ 秒杀商品列表
+ 秒杀商品详情

## 购物车服务

+ 添加商品到购物车
+ 购物车商品列表
+ 删除购物车商品
+ 修改购物车商品的信息

## 订单服务

+ 创建普通商品订单
+ 创建秒杀商品订单
+ 订单列表
+ 订单详情
+ 删除订单信息

## 支付服务

因为是个人项目，没法开通第三方支付途径，所以只实现了用户钱包余额支付功能

+ 订单付款（用户钱包）
