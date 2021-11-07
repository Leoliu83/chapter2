package charset

import "io"

// // noCopy may be embedded into structs which must not be copied after the first use.
// type noCopy struct{}

// // Lock is a no-op used by -copylocks checker from `go vet`.
// func (*noCopy) Lock() {}

type CharsetDetector interface {
	ReadSeekerCharsetDetector
}

type ReadSeekerCharsetDetector interface {
	/*
		读取指定size的数据，用以探测字符集
	*/
	DetectReadSeekerWithSize(r io.ReadSeeker, size int)
}
