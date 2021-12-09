package main

import "fmt"

func main() {
	var avgFatRate float64 = 0.0
	var cnn int = 0
	for {
		var name string
		fmt.Print("姓名:")
		fmt.Scanln(&name)

		var weight float64
		fmt.Print("体重（Kg）:")
		fmt.Scanln(&weight)

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

		//体脂率计算器
		var bmi float64 = weight / (height * height)
		var fatRate float64 = (1.23*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
		fmt.Printf("bmi是：%2.2f,体脂率是：%2.2f\n", bmi, fatRate)

		cnn++
		avgFatRate = (avgFatRate*float64(cnn-1) + fatRate) / float64(cnn)
		fmt.Printf("共录入%d人记录，平均体脂为%2.2f\n", cnn, avgFatRate)

		if sex == "男" {
			if age >= 18 && age <= 39 {
				if fatRate >= 0 && fatRate < 0.1 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.1 && fatRate < 0.16 {
					fmt.Println("目前是：标准，好好保持哟")
				} else if fatRate >= 0.16 && fatRate < 0.21 {
					fmt.Println("目前是：偏重，该有些适当的运动啦")
				} else if fatRate >= 0.21 && fatRate < 0.26 {
					fmt.Println("目前是：肥胖，快点减肥")
				} else {
					fmt.Println("目前是：严重肥胖，再不减肥就要去医院了")
				}
			} else if age >= 40 && age <= 59 {
				if fatRate >= 0 && fatRate < 0.11 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.11 && fatRate < 0.17 {
					fmt.Println("目前是：标准，好好保持哟")
				} else if fatRate >= 0.17 && fatRate < 0.22 {
					fmt.Println("目前是：偏重，该有些适当的运动啦")
				} else if fatRate >= 0.22 && fatRate < 0.27 {
					fmt.Println("目前是：肥胖，快点减肥")
				} else {
					fmt.Println("目前是：严重肥胖，再不减肥就要去医院了")
				}
			} else if age >= 60 {
				if fatRate >= 0 && fatRate < 0.13 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.13 && fatRate < 0.19 {
					fmt.Println("目前是：标准，好好保持哟")
				} else if fatRate >= 0.19 && fatRate < 0.24 {
					fmt.Println("目前是：偏重，该有些适当的运动啦")
				} else if fatRate >= 0.24 && fatRate < 0.29 {
					fmt.Println("目前是：肥胖，快点减肥")
				} else {
					fmt.Println("目前是：严重肥胖，再不减肥就要去医院了")
				}
			} else {
				fmt.Println("我们不看未成年人的体脂率，变化太大了")
			}
		} else {
			if age >= 18 && age <= 39 {
				if fatRate >= 0 && fatRate < 0.2 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.2 && fatRate < 0.27 {
					fmt.Println("目前是：标准，好好保持哟")
				} else if fatRate >= 0.27 && fatRate < 0.34 {
					fmt.Println("目前是：偏重，该有些适当的运动啦")
				} else if fatRate >= 0.34 && fatRate < 0.39 {
					fmt.Println("目前是：肥胖，快点减肥")
				} else {
					fmt.Println("目前是：严重肥胖，再不减肥就要去医院了")
				}
			} else if age >= 40 && age <= 59 {
				if fatRate >= 0 && fatRate < 0.21 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.21 && fatRate < 0.28 {
					fmt.Println("目前是：标准，好好保持哟")
				} else if fatRate >= 0.28 && fatRate < 0.35 {
					fmt.Println("目前是：偏重，该有些适当的运动啦")
				} else if fatRate >= 0.35 && fatRate < 0.4 {
					fmt.Println("目前是：肥胖，快点减肥")
				} else {
					fmt.Println("目前是：严重肥胖，再不减肥就要去医院了")
				}
			} else if age >= 60 {
				if fatRate >= 0 && fatRate < 0.22 {
					fmt.Println("目前是：偏瘦，要多吃多练")
				} else if fatRate >= 0.22 && fatRate < 0.29 {
					fmt.Println("目前是：标准，好好保持哟")
				} else if fatRate >= 0.29 && fatRate < 0.36 {
					fmt.Println("目前是：偏重，该有些适当的运动啦")
				} else if fatRate >= 0.36 && fatRate < 0.41 {
					fmt.Println("目前是：肥胖，快点减肥")
				} else {
					fmt.Println("目前是：严重肥胖，再不减肥就要去医院了")
				}
			} else {
				fmt.Println("我们不看未成年人的体脂率，变化太大了")
			}
		}

		var wheterContinue string
		fmt.Printf("y 继续，q 退出\n")
		fmt.Print("是否录入下一个：")
		fmt.Scanln(&wheterContinue)
		if wheterContinue != "y" {
			break
		}
	}
}
