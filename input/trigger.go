package input

type Trigger interface {
	RecvInput() (x, y int32, err error)
}
