package main

import "fmt"
/*
接口嵌套
接受者为值
接受者为指针
是否能改变对应的值
 */
type open interface {
	open() string
}
type read interface {
	read()string
}


type mio interface {
	open
	read
	write() string
}

type fileinfo struct {
	path string
}


func (f *fileinfo)open()string  {
	return "opened:"+f.path
}
func (f *fileinfo)read()string  {
	return "read:"+f.path
}
func (f fileinfo)write()string  {
	return "write:"+f.path
}
//-----------------------------------------
type animal interface {
	miao()string
	addAge(n int)
	getAge() int
	intro() string
}

type cat struct {
	name string
	age int

}

func (m cat)miao()string  {
	return "miao"
}
func (m cat )addAge(n int)  {
	m.age=m.age+n
}
func (m cat)getAge()int  {
	return m.age
}

func (m cat)intro()string{
	return fmt.Sprintf("name=%s,age=%d",m.name,m.age)
}



func main()  {


	var io1 mio=&fileinfo{"/usr/local"} //有指针接受者必须&
	fmt.Println(io1.read())
	fmt.Println(io1.write())

	fmt.Println("----------------------\n")

	var cat1 animal=cat{"小花",10}
	fmt.Println(cat1)
	cat1.addAge(2)
	fmt.Println(cat1.intro()) //name=小花,age=10   age想被改变必须接收者为指针





}


