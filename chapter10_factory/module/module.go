package module

type Student struct {
	Name  string
	Score float64
}

type student2 struct {
	Name  string
	Score float64
}

// 因为student结构体首字母是小写，因此只能在model使用
// 我们通过工厂模式来解决这个问题

func NewStudent(n string, s float64) *student2 {
	return &student2{
		Name:  n,
		Score: s,
	}
}
