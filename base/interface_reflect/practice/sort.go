package main

import "fmt"

type sorter interface{
	Len()int
	Swap(i ,j int)
	Less(i,j int)bool

}
type IntArray []int
func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] > p[j] }
func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }



type StringArray []string


func (p StringArray) Len() int           { return len(p) }
func (p StringArray) Less(i, j int) bool { return p[i] > p[j] }
func (p StringArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }


func sort(data sorter)  {
	for i:=0;i<data.Len()-1;i++{
		for j:=i+1;j<data.Len() ;j++  {
			if data.Less(i,j){
				data.Swap(i,j)
			}

		}
	}
}

func main()  {
	data:=[]int{2,1,5,3,6,9,0}
	a := IntArray(data)
	fmt.Println(a)
	sort(a)
	fmt.Println(a);


	datastr:=[]string{"php","ajax","java"}
	strarr:=StringArray(datastr)
	fmt.Println(strarr[0])

	fmt.Println(strarr)
	sort(strarr)
	fmt.Println(strarr)









}

