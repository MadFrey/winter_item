# winter_item

## 版本更新

>v1.01->v1.10 22.03.15
>
> 添加第三方认认证，可以使用github账号进行登录，注册（绑定GitHub账号功能还没完善）
> 
> 添加邮箱验证码修改密码（邮箱验证码登录注册功能还未完善）
> 
> 优化代码，增加jwt中间件（部分功能已经改成jwt中间件，还有部分没有更改）
> 
> 修正了部分错误码 
>
> 1.01-22.03.03
>
> 优化了代码，增加对用户名和密码的检查和对密码的加密，
>
> 新增错误码规范

### 2月19日上传

### **项目部分**

>1.完成了全部的基础功能
>
>2.接口参数与作业文档基本相同，但有个别不同在下方注意事项列出
>
>3.项目部分功能注释后面补上

### 注意事项

>1.在请求头中要放入token和refreshToken，用一个空格隔开
>
>2.因为项目中有文档名为model，因此评论等部分需要传model参数的地方表单key都改为了Model（首字母大写）
>
>3.帖子和1级评论的表单key为file（忘记改了）
>
>4.在完成注册第一次登录后，先更改用户信息，最好都填一下（或者数据库的默认值不要设置为空（NULL））
>
>（可能有遗漏，待后续补充）
