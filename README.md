# EasyGinLite
一个简易的golang框架，支持前后端分离、内容渲染前端、作为API，加入redis、mysql、es等常用的库吧
### 文件目录：
+ config（配置文件夹）
  + conf.ini（数据库配置在这里）
  + config.go（初始化配置）
+ dist（前端文件）
  + any files（杂七杂八的文件或者文件夹）
+ middleware（中间件文件夹）
  + cors.go（中间件）
+ router（路由文件夹）
  + router.go（路由配置）
+ utils（工具文件夹）
  + global.go（全局mysql、redis的实例变量）
  + mysql.go（mysql的定义）
  + redis.go（redis的定义）
+ version（Html或者Api的文件夹）
  + v1（版本，可自定义）
    + demo.go（调用全局redis、mysql的例子） 
    + index.go （API、前后端分离的例子）
    + indexHtml.go（前端直接渲染的例子）
+ main.go（执行文件）

### 本人也是个新手，如有好的建议可PR ，我来学习
对于新手而言，综合我从node 、php 、python转入golang的实际情况，我觉得这个框架还是比较友好的




