package main

import "fmt"

func main() {
	// 创建数组
	point1 := [2][2]int{}
	point2 := [2][2]int{}
	// 向二维数组添加元素
	fmt.Print("请输入第一条直线上第一个点的坐标:")
	fmt.Scanln(&point1[0][0])
	fmt.Scanln(&point1[0][1])
	fmt.Print("请输入第一条直线上第二个点的坐标:")
	fmt.Scanln(&point1[1][0])
	fmt.Scanln(&point1[1][1])

	fmt.Print("请输入第二条直线上第一个点的坐标:")
	fmt.Scanln(&point2[0][0])
	fmt.Scanln(&point2[0][1])
	fmt.Print("请输入第二条直线上第二个点的坐标:")
	fmt.Scanln(&point2[1][0])
	fmt.Scanln(&point2[1][1])

	// 判断两条线段之间的关系
	if (point1[1][0]-point1[0][0]) == 0 || (point2[1][0]-point2[0][0]) == 0 {
		fmt.Println("该输入不合法，请重新输入")
	} else {
		k1 := (point1[1][1] - point1[0][1]) / (point1[1][0] - point1[0][0])
		k2 := (point2[1][1] - point2[0][1]) / (point2[1][0] - point2[0][0])
		if k1 == k2 {
			if point2[1][1] == point1[1][1] {
				fmt.Println("这两条直线重合")
			} else {
				fmt.Println("这两条直线相互平行")
			}
		} else {
			fmt.Println("这两条直线相交")
		}
	}
}
