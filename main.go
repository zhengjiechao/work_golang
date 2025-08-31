package main

import (
	"work_golang/task4/models"
	"work_golang/task4/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	// 任务一

	// // 1.只出现一次的数字
	// slice := []int{1, 1, 2, 5, 2, 4, 4}
	// result := task1.SingleNumber(slice)
	// fmt.Println(result)

	// // 2.回文数
	// result := task1.IsPalindrome(25552)
	// fmt.Println(result)

	// // 3.有效的括号
	// result := task1.IsValid("({[]})")
	// fmt.Println(result)

	// // 4.最长公共前缀
	// strArr := []string{"hhfvdes", "hhfadsf"}
	// result := task1.LongestCommonPrefix(strArr)
	// fmt.Println(result)

	// 5.加一
	// intArr := []int{5, 9, 8, 9}
	// result := task1.PlusOne(intArr)
	// fmt.Println(result)

	// 6.删除有序数组中的重复项
	// intArr := []int{5, 9, 8, 9}
	// result := task1.RemoveDuplicates(intArr)
	// fmt.Println(result)

	// // 7.合并区间
	// arr := [][]int{[]int{5, 6}, []int{1, 2}, []int{2, 4}, []int{5, 5}, []int{5, 5}, []int{3, 3}}
	// task1.Merge(arr)

	// // 8.两数之和
	// intArr := []int{5, 9, 8, 9}
	// result := task1.TwoSum(intArr, 14)
	// fmt.Println(result)

	// 任务二

	// // 1.1
	// num := 100
	// task2.Pointer1(&num)
	// fmt.Println(num)

	// // 1.2
	// arr := []int{1, 2, 3}
	// task2.Pointer2(&arr)
	// fmt.Println(arr)

	// // 2.1
	// task2.Goroutine1()

	// 2.2
	// arr := []func(){
	// 	func() {
	// 		time.Sleep(time.Second * 1)
	// 	},
	// 	func() {
	// 		time.Sleep(time.Second * 2)
	// 	},
	// 	func() {
	// 		time.Sleep(time.Millisecond * 500)
	// 	},
	// }
	// task2.Goroutine2(arr)

	// // 3.1
	// r := task2.Rectangle{}
	// r.Area()
	// r.Perimeter()

	// c := task2.Circle{}
	// c.Area()
	// c.Perimeter()

	// // 3.2
	// e := task2.Employee{
	// 	EmployeeID: 1,
	// 	Person: task2.Person{
	// 		Name: "zhagnsan",
	// 		Age:  55,
	// 	},
	// }
	// e.PrintInfo()

	// // 4.1
	// task2.Channel1()

	// // 4.2
	// task2.Channel2()

	// // 5.1
	// task2.Lock1()

	// // 5.2
	// task2.Lock2()

	// 任务三
	// 1.2转账
	// task3.Trans()

	// 2.1
	// employeeArr := task3.Query1()
	// fmt.Println(employeeArr)

	// employee := task3.Query2()
	// fmt.Println(employee)

	// 2.2
	// booArr := task3.QueryBook()
	// fmt.Println(booArr)

	// 3.1
	// task3.CreateTable()

	// 3.2
	// users := task3.QueryUsers()
	// fmt.Println(users)
	// post := task3.QueryPost()
	// fmt.Println(post)

	// 3.3
	// post := task3.Post{
	// 	Title:  "语文",
	// 	UserId: 1,
	// 	Status: "无评论",
	// }
	// task3.DB.Create(&post)

	// comment := task3.Comment{
	// 	Id:     2,
	// 	PostId: 2,
	// }
	// task3.DB.Delete(&comment)

	// task4
	models.CreateTable()

	r := gin.Default()
	routers.UserRouterInit(r)
	routers.PostRouterInit(r)
	routers.CommentRouterInit(r)

	r.Run(":9000")

}
