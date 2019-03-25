### 基于beego的博客
+ golang的beego框架


### 本地环境配置
+ 安装golang
+ 安装beego相关的库
+ 克隆项目到本地
```bash
git clone https://github.com/zhourunliang/go-blog
```
+ 创建数据表
```sql
CREATE TABLE `blog` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` text,
  `content` text,
  `created` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  `password` varchar(100) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;
```
+ 运行程序
```bash
go run main.go
```
+ 用浏览器访问http://127.0.0.1:8080

