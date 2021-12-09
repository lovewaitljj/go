package main

import "fmt"

func main() {

	for {
		var name string
		fmt.Print("姓名:")
		fmt.Scanln(&name)

		var weight float64
		fmt.Print("体重（Kg）:")
		fmt.Scanln(&weight) //取这个变量的编号

		var height float64
		fmt.Print("身高（m）:")
		fmt.Scanln(&height)

		var age int
		fmt.Print("年龄:")
		fmt.Scanln(&age)

		var sex string
		fmt.Print("性别（男/女）:")
		fmt.Scanln(&sex)

		var sexWeight int
		if sex == "男" {
			sexWeight = 1
		} else {
			sexWeight = 0
		}
		//体脂计算器
		var bmi float64 = weight / (height * height)
		var fatRate float64 = (1.23*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight))
		fmt.Println("体脂率是：", fatRate)

		//身材匹配器

		if sex == "男" {
			if age >= 18 && age <= 39 {
				if fatRate >= 0.1 && fatRate <= 0.16 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else {
					fmt.Println("目前是：偏瘦，要多吃多练")
				}
			} else if age >= 18 && age <= 39 {
				if fatRate >= 0.1 && fatRate <= 0.16 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else {
					fmt.Println("目前是：偏瘦，要多吃多练")
				}
			} else if age >= 18 && age <= 39 {
				if fatRate >= 0.1 && fatRate <= 0.16 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else {
					fmt.Println("目前是：偏瘦，要多吃多练")
				}
			} else {
				fmt.Println("我们不看未成年人的体脂率，变化太大了")
				if fatRate >= 0.1 && fatRate <= 0.16 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else {
					fmt.Println("目前是：偏瘦，要多吃多练")
				}
			}
		} else {
			if age >= 18 && age <= 39 {
				if fatRate >= 0.1 && fatRate <= 0.16 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else {
					fmt.Println("目前是：偏瘦，要多吃多练")
				}
			} else if age >= 18 && age <= 39 {
				if fatRate >= 0.1 && fatRate <= 0.16 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else {
					fmt.Println("目前是：偏瘦，要多吃多练")
				}
			} else if age >= 18 && age <= 39 {
				if fatRate >= 0.1 && fatRate <= 0.16 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else {
					fmt.Println("目前是：偏瘦，要多吃多练")
				}
			} else {
				fmt.Println("我们不看未成年人的体脂率，变化太大了")
				if fatRate >= 0.1 && fatRate <= 0.16 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else {
					fmt.Println("目前是：偏瘦，要多吃多练")
				}
			}
		}
		var wheterContinue string
		fmt.Print("是否录入下一个：")
		fmt.Scanln(&wheterContinue)
		if wheterContinue != "y" {
			break
		}
	}
}
