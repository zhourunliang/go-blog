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
CREATE TABLE entries (
    id INT AUTO_INCREMENT,
    title TEXT,
    content TEXT,
    created DATETIME,
    primary key (id)
);
```
+ 运行程序
```bash
go run main.go
```
+ 用浏览器访问http://127.0.0.1:8080

