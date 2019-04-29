package main

type Empty interface {}
type semaphore chan Empty

func (s semaphore)P(n int)  {
	e:=new(Empty)
	for i:=0;i<n ;i++  {
		s<-e
	}
}
func (s semaphore)V(n int)  {
	for i:=0;i<n ;i++  {
		<-s
	}

}
func (s semaphore)Lock()  {
	s.P(1)
}
func (s semaphore)Unlock(){
	s.V(1)
}
func (s semaphore)Wait(n int)  {
	s.P(n)
}
func (s semaphore)Signal()  {
	s.P(1);
}
func main()  {
		
}
