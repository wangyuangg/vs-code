package model

type student struct {
	Name  string
	score float64
}
//because the struct is private, so we can't use it outside the package.

func NewStudent(name string, score float64) *student {
	return &student{
		Name:  name,
		score: score,
	}
}
func (s *student) ReturnInt() float64{
	return s.score
}