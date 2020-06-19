# CVTElottery
CVTE实训项目：抽奖轮子

[API接口 - 开发文档](https://www.showdoc.cc/CVTElottery)



# 使用框架/技术

[singo - gin脚手架](https://github.com/Gourouting/singo)

[GORM - golang的ORM技术](http://gorm.book.jasperxu.com/)

[air - gin热启动插件](https://github.com/cosmtrek/air)

[showdoc - 自动生成api文档](https://www.showdoc.cc/page/741656402509783)



# 启动项目 

### 1、配置.env文件

```
MYSQL_DSN="root:123456@/CVTE?charset=utf8&parseTime=True&loc=Local"
REDIS_ADDR="127.0.0.1:6379"
REDIS_PW=""
REDIS_DB=""
SESSION_SECRET="setOnProducation"
GIN_MODE="debug"
LOG_LEVEL="debug"
```



### 2、使用air启动

命令行输入`air` 即可



### 3、生成api文档

在git中执行 `./showdoc_api.sh`