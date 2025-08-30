package task3

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 题目1：基本CRUD操作
// insert into students(name,age,grade) values('张三',20,'三年级');
// select * from students where age > 18;
// update students set grade = '四年级' where name = '张三';
// delete from students where age < 15;

// 题目2 事务语句
var DB *gorm.DB
var err error

func init() {
	dsn := "root:123456@tcp(localhost:3306)/gin?charset=utf8mb4&parseTime=true&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
	}
}

type Account struct {
	Id      string
	Balance int
}

func (a Account) TableName() string {
	return "accounts"
}

type Transaction struct {
	Id            int
	FromAccountId string
	ToAccountId   string
	Amount        int
}

func (t Transaction) TableName() string {
	return "transactions"
}

// 转账操作
func Trans() {
	// transaction函数
	// DB.Transaction(func(tx *gorm.DB) error {
	// 	accountA := Account{Id: "A"}
	// 	tx.Find(&accountA)
	// 	if accountA.Balance >= 100 {
	// 		// A账户减去100元
	// 		accountA.Balance -= 100
	// 		if err := tx.Save(&accountA).Error; err != nil {
	// 			return err
	// 		}

	// 		// B账户加上100元
	// 		accountB := Account{Id: "B"}
	// 		tx.Find(&accountB)
	// 		accountB.Balance += 100
	// 		if err := tx.Save(accountB).Error; err != nil {
	// 			return err
	// 		}

	// 		// 保存转账记录
	// 		transaction := Transaction{
	// 			FromAccountId: "A",
	// 			ToAccountId:   "B",
	// 			Amount:        100,
	// 		}
	// 		if err := tx.Create(&transaction).Error; err != nil {
	// 			return nil
	// 		}
	// 	}
	// 	return nil
	// })

	// 手动提交
	tx := DB.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	countA := Account{Id: "A"}
	tx.Find(&countA)
	if countA.Balance >= 100 {
		countA.Balance -= 100
		if err := tx.Save(&countA).Error; err != nil {
			tx.Rollback()
		}

		countB := Account{Id: "B"}
		tx.Find(&countB)
		countB.Balance += 100
		if err := tx.Save(&countB).Error; err != nil {
			tx.Rollback()
		}

		transaction := Transaction{
			FromAccountId: "A",
			ToAccountId:   "B",
			Amount:        100,
		}

		if err := tx.Save(&transaction).Error; err != nil {
			tx.Rollback()
		}
	}
	tx.Commit()
}
