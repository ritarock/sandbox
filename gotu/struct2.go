package main

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      // 匿名フィールド、struct
	Skills     // 匿名フィールド、自分で定義した型。string slice
	int        // ビルトイン型を匿名フィールドとします。
	speciality string
}

func main() {
	// 学生Jannを初期化します。
	jane := Student{Human: Human{"Jane", 35, 100}, speciality: "Biology"}
	// ここで対応するフィールドにアクセスしてみます。
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)
	// 彼のskill技能フィールドを修正します。
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	// 匿名ビルトイン型のフィールドを修正します。
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
}
