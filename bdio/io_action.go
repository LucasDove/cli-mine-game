package bdio

type InputReceiver interface {
	Input() (x, y int32, err error)
}

type OutputReceiver interface {
	Output(bvalue [][]int8)
}