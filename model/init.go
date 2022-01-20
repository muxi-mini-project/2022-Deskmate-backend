//初始化和链接数据库
package model

import (
    //导入MYSQL数据库驱动，这里使用的是GORM库封装的MYSQL驱动导入包，实际上大家看源码就知道，这里等价于导入github.com/go-sql-driver/mysql
    //这里导入包使用了 _ 前缀代表仅仅是导入这个包，但是我们在代码里面不会直接使用。
    _ "github.com/jinzhu/gorm/dialects/mysql"//这个我是go get导入的
    //导入gorm
    "github.com/jinzhu/gorm"
)

// DB 全局变量
var DB *gorm.DB
