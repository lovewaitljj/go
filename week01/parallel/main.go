package main

import "fmt"

func main() {
	var x1 float64
	fmt.Print("请输入第一条直线上第一个点的x坐标:")
	fmt.Scanln(&x1)
	var y1 float64
	fmt.Print("请输入第一条直线上第一个点的y坐标:")
	fmt.Scanln(&y1)

	var x2 float64
	fmt.Print("请输入第一条直线上第二个点的x坐标:")
	fmt.Scanln(&x2)
	var y2 float64
	fmt.Print("请输入第一条直线上第二个点的y坐标:")
	fmt.Scanln(&y2)

	var x3 float64
	fmt.Print("请输入第二条直线上第一个点的x坐标:")
	fmt.Scanln(&x3)
	var y3 float64
	fmt.Print("请输入第二条直线上第一个点的y坐标:")
	fmt.Scanln(&y3)

	var x4 float64
	fmt.Print("请输入第二条直线上第二个点的x坐标:")
	fmt.Scanln(&x4)
	var y4 float64
	fmt.Print("请输入第二条直线上第二个点的y坐标:")
	fmt.Scanln(&y4)

	// 判断两条线段之间的关系
	if (y3-y4) == 0 || (y1-y2) == 0 {
		fmt.Println("该输入不合法，请重新输入")
	} else {
		k1 := (y1 - y2) / (x1 - x2)
		k2 := (y3 - y4) / (x3 - x4)
		if k1 == k2 {
			if y2 == y4 {
				fmt.Println("这两条直线重合")
			} else {
				fmt.Println("这两条直线相互平行")
			}
		} else {
			fmt.Println("这两条直线相交")
		}
	}
}
