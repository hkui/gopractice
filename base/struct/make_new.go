package main

type m1 map[string]string
type s1 struct{
	 a int
	 b string
}

func main()  {
	ss1:=new(s1)
	(*ss1).a=10
	(*ss1).b="bbb"

	//ss2:=make(s1)  //编译错误

	mm1:=make(m1)
	mm1["a"]="ajax"


}
/**
试图 make() 一个结构体变量，会引发一个编译错误，这还不是太糟糕，
但是 new() 一个映射并试图使用数据填充它，将会引发运行时错误！
因为 new(m) 返回的是一个指向 nil 的指针，它尚未被分配内存
所以在使用 map 时要特别谨慎
 */
