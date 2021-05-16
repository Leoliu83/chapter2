package internal

type InternalObj struct {
	x int
	y int
}

func GetInternalObj() *InternalObj {
	return new(InternalObj)
}
