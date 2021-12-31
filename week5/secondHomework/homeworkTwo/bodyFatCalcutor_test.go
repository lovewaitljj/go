package main

import (
	"fmt"
	"testing"
)

func TestCalculateBMI(t *testing.T) {
	//BMI案例
	//录入正常身高、体重，确保计算结果符合预期；
	bmiTest1, errCase1 := CalculateBMI(1.78, 65)
	fmt.Println("bmiTest1:", bmiTest1, " errCase1:", errCase1)
	if bmiTest1 < 0 || errCase1 != nil {
		t.Fatal()
	}

	//录入 0 身高，返回错误；
	bmiTest2, errCase2 := CalculateBMI(0, 65)
	fmt.Println("bmiTest2:", bmiTest2, " errCase2: ", errCase2)
	if bmiTest2 != 0 || errCase2 == nil {
		t.Fatal()
	}

	//录入负数身高，返回错误；
	bmiTest3, errCase3 := CalculateBMI(-5, 65)
	fmt.Println("bmiTest3:", bmiTest3, " errCase3: ", errCase3)
	if bmiTest3 < 0 || errCase3 == nil {
		t.Fatal()
	}

	//录入0体重，返回错误。
	bmiTest4, errCase4 := CalculateBMI(1.78, 0)
	fmt.Println("bmiTest4:", bmiTest4, " errCase4:", errCase4)
	if bmiTest4 != 0 || errCase4 == nil {
		t.Fatal()
	}

	//录入负数体重，返回错误。
	bmiTest5, errCase5 := CalculateBMI(1.78, -80)
	fmt.Println("bmiTest5:", bmiTest5, " errCase5: ", errCase5)
	if bmiTest5 < 0 || errCase5 == nil {
		t.Fatal()
	}
}

func TestCalculateFatRate(t *testing.T) {
	//体脂率案例
	//录入正常 BMI、年龄、性别，确保计算结果符合预期；

	//录入完整的性别、年龄、身高、体重，确保最终获得的健康建议符合预期。
	fatRateCase1, suggest1, errCase1 := CalculateFatRate(21.3, 20, "男")
	fmt.Println("fatRateCase1:", fatRateCase1, " suggest1:", suggest1, " errCase1:", errCase1)
	if fatRateCase1 < 0 || errCase1 != nil {
		t.Fatal()
	}

	//录入非法 BMI，返回错误；
	_, _, errCase2 := CalculateFatRate(0, 20, "男")
	if errCase2 == nil {
		t.Fatal()
	}

	_, _, errCase3 := CalculateFatRate(-5, 20, "男")
	if errCase3 == nil {
		t.Fatal()
	}

	//录入非法年龄，返回错误；
	_, _, errCase4 := CalculateFatRate(23.5, 0, "男")
	if errCase4 == nil {
		t.Fatal()
	}

	_, _, errCase5 := CalculateFatRate(29, 151, "男")
	if errCase5 == nil {
		t.Fatal()
	}

	//录入非法性别，返回错误；
	_, _, errCase6 := CalculateFatRate(26, 20, "不男不女")
	if errCase6 == nil {
		t.Fatal()
	}

}
