package duckType

type Shaper interface {
	Area() float32
}

type TopologicalGenus interface {
	Rank() int
}
type Square struct {
	Side float32
}

func (sq *Square)Area()float32  {
	return sq.Side*sq.Side
}
func (sq *Square)Rank()int  {
	return 1
}

type Rectangle struct {
	Length,Width float32
}

func (r Rectangle)Area()float32  {
	return r.Length*r.Width
}

func (r Rectangle)Rank()int  {
	return 2

}


