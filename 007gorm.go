package main

/**
  gorm sqlite3
  参考： http://gorm.book.jasperxu.com
*/
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "root:abc123@tcp(localhost:3306)/go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("链接数据库异常")
	}
	defer db.Close()

	// 表名去掉s
	db.SingularTable(true)

	// 自动迁移模式
	db.AutoMigrate(&Product{})

	// 启用Logger，显示详细日志
	db.LogMode(true)

	// 设置连接池
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	/**

	//insert
	db.Create(&Product{Code: "L1212", Price: 110})

	//读取
	var product Product
	db.First(&product, 1)
	db.First(&product, "code=?", "L1212")

	//更新
	db.Model(&product).Update("Price", 220)

	//删除
	db.Delete(&product)

	// 造数据
	for i := 0; i < 50; i++ {
		db.Create(&Product{Code: string(i), Price: uint(i)})
	}



	// list
	rows, err := db.Model(&Product{}).Where("price > ?", 45).Rows()
	defer rows.Close()

	for rows.Next() {
		var product Product
		db.ScanRows(rows, &product)
		fmt.Printf("id=%d,code=%s,price=%d\n", product.ID, product.Code, product.Price)
	}

	// raw sql
	db.Exec("update product set price=price + 100 where price > 45")

	// raw批量更新
	db.Exec("update product set price=price + 100 where price > ?",46)

	*/

	// 切片迭代
	var productList []Product
	db.Select([]string{"id", "code", "price"}).Find(&productList)
	for _, product := range productList {
		fmt.Printf("id=%d,price=%d\n", product.ID, product.Price)
	}

}
