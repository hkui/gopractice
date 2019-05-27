package main

func main(){
/*
	var ret= map[string]string{"a":"ajax"} //定义时才能用{} 赋值
	ret["age"]="hk"

	fmt.Println(ret["a"])
*/

	/*
	var m1=map[string](map[int]string){"a":{1:"ajax1",3:"php"},"bb":{10:"java"}}
	fmt.Println(m1)


	//make初始化

	map1:=make(map[string]int)
	map1["p"]=1

	map2:=make(map[string]map[int]string)
	//map2["php"][1]="ok"  //直接这样用会报错
	map2["php"]=make(map[int]string)
	map2["php"][1]="ok"
	fmt.Println(map1,map2)
	*/

	//函数和切片作为值类型
/*
	mapfun:=map[int]func(int)int{
		1: func(n int ) int {
			return n
		},
		2: func( n int) int {
			return 2*n
		},
		5: func(n int) int {
			return 5*n
		},
	}
	fmt.Println(mapfun)
	fmt.Println(mapfun[5](5))

	mapslice:=map[int][]byte{
		1:{1,3,5,7},
		2:{4,6,8,9},
	}
	fmt.Println(mapslice)

*/
	//测试键值对是否存在及删除元素
/*
	m1:=map[string]string{"a":"ajax","c":"cpp"}

	b,exists:=m1["b"]

	fmt.Println(b,exists)

	if _,exista:=m1["a"];exista{
		fmt.Println("a exists")
	}else{
		fmt.Println("a not exists")
	}
	delete(m1,"c")
	fmt.Println(m1)
*/
	//假设我们想获取一个 map 类型的切片，使用两次 make() 函数，第一次分配切片，第二次分配 切片中每个 map

/*
	items:=make([]map[int]int,5)
	for i:=range items{
		items[i]=make(map[int]int,1)
		items[i][1] = 2
	}
	fmt.Println(items)
*/

	/*
	a:=map[string]string{"a":"ajax","c":"cpp","j":"java","b":"bjson"}

	keys:=make([]string,len(a))


	i:=0

	for k,_:=range a{
		keys[i]=k
		i++
	}
	sort.Strings(keys)
	fmt.Println(keys)
	for _,v:=range keys{
		fmt.Println(a[v])
	}*/

























}
