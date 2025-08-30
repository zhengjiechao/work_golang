package task3

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "root:123456@tcp(localhost:3306)/gin?charset=utf8mb4&parseTime=true&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}

// 3.1
type User struct {
	Id        int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"notnull;type:varchar(50);comment:'姓名'"`
	PostCount int    `gorm:"type:int;comment:'文章数量'"`
	Posts     []Post `gorm:"foreignKey:UserId"`
}

func (u User) TableName() string {
	return "users"
}

type Post struct {
	Id       int       `gorm:"primaryKey;autoIncrement"`
	Title    string    `gorm:"notnull;type:varchar(100);comment:'标题'"`
	UserId   int       `gorm:"comment:用户id"`
	Status   string    `gorm:"type:varchar(16);comment:'评论状态'"`
	Comments []Comment `gorm:"foreignKey:PostId"`
}

func (p Post) TableName() string {
	return "posts"
}

type Comment struct {
	Id     int    `gorm:"primaryKey;autoIncrement"`
	Desc   string `gorm:"type:varchar(100);comment:'描述'"`
	PostId int    `gorm:"comment:文章id"`
}

func (c Comment) TableName() string {
	return "comments"
}

func CreateTable() {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
}

// 3.2
func QueryUsers() []User {
	var users []User
	db.Preload("Posts.Comments").Find(&users)
	return users
}

func QueryPost() Post {
	var post Post
	db.Table("posts a").Select("a.*,count(b.id) count").Joins("left join comments b on a.id = b.post_id").Group("a.id").Order("count desc").First(&post)
	return post
}

// 3.3 钩子函数
func (p *Post) AfterCreate(tx *gorm.DB) error {
	fmt.Println("执行钩子函数(文章创建时自动更新用户的文章数量统计字段)----------------")
	u := User{Id: p.UserId}
	tx.Find(&u)
	u.PostCount += 1
	err := tx.Save(&u).Error
	return err
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {
	fmt.Println("执行钩子函数(在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 无评论)----------------")
	var newComments []Comment
	tx.Where("post_id = ?", c.PostId).Find(&newComments)
	if len(newComments) == 0 {
		p := Post{Id: c.PostId}
		tx.Find(&p)
		p.Status = "无评论"
		return tx.Save(p).Error
	}
	return nil
}
