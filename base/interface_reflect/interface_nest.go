package main
//接口嵌套
type open interface {
	open()
}
type read interface {
	read()
}

type write interface {
	write()
}
type close interface {
	close()
}



type io interface {
	open
	read
	write
	close
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
func (f *fileinfo)write()string  {
	return "write:"+f.path
}
func (f *fileinfo)close()string  {
	return "closed:"+f.path
}



func main()  {
	
}


