package charset

/*
	字符集探测代理，多个字符集可以通过add方式添加
*/
type CharsetDetectorProxy struct {
	detectors []CharsetDetector
}

/*
	添加字符集探测器
*/
func (cdp *CharsetDetectorProxy) Add(charsetDetector CharsetDetector) {
	cdp.detectors = append(cdp.detectors, charsetDetector)
}
