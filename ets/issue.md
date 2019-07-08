# issue

## 创建数据库关联

    1. 数据库连接: connect--"root:tester@tcp(127.0.0.1:3306)/lottery?charset=utf8"
    2. cd $GOPATH/src/github.com/go-xorm/cmd/xorm
    3. xorm reverse mysql "root:tester@tcp(127.0.0.1:3306)/lottery?charset=utf8" templates/goxorm 在models目录下生成的数据模型文件
