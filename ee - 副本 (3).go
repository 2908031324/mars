package main

import (
	"fmt"
)

type user struct {
	id int
	username string
	address string
	school map[string]string
	education []map[string]string
}

//https://www.runoob.com/go/go-structures.html 菜鸟教程

func main(){
	fmt.Printf("%v", user{
		education: []map[string]string{
			{"middle school": "三年"},
			{"university": "四年"},
		},
	})

	var uu user
	uu.id = 1
	uu.username = "中国"
	uu.address  = "广州天河区"
	uu.school   = map[string]string{
		"陆中":"高一至高三",
		"大学":"华南师范大学",
	}
	uu.education = []map[string]string{
		{"middle school": "三年", "awards":"获取全国数学竞赛二等奖"},
		{"university": "四年", "awards":"获取全国物理竞赛二等奖"},
	}

	fmt.Println(uu)

	fmt.Printf("\n\n\n%v\n", uu)

	fmt.Printf("%v", uu.id)
	var struct_pointer *user
	fmt.Println("**************************")
	fmt.Println(struct_pointer)

	ppp(&uu)
	//for _, v1 := range uu{
	//
	//}


	aaa := map[string]string{
		"k1":"name1",
		"k2":"name2",
	}
	for _, vvv := range aaa {
		fmt.Printf("\n-...-%v\n", vvv)
	}

}

func ppp(uu *user){
	fmt.Printf("值：\n", uu.education)

}
