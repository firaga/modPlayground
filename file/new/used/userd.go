package used

type Inter interface {
	isInterface()
}

type A struct {
	One int
}

func (*A) isInterface() {

}

func NewA() Inter {
	return &A{}
}
