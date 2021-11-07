package charset

/*
	字符集结构体
*/
type Charset struct {
	name  string          // 标准名称
	alias map[string]bool // 别名列表
	_     struct{}        // 保证必须使用属性名来初始化结构体
}

/*
	未找到别名则返回false, 找到则返回可作为参数的通用字符集名称
	@param name 字符集名称字符串
	@return string 标准字符集名称
	        bool   是否能找到符合的字符集名称（golang中，当map返回值是bool，如果未找到元素则返回false）
*/
func (c Charset) CheckName(name string) (string, bool) {
	return c.name, c.alias[name]
}

// A lot of Charset defined
var UNKNOWN Charset = Charset{name: "UNKNOWN"}
var US_ASCII Charset = Charset{name: "US-ASCII"}
var UTF_8 Charset = Charset{name: "UTF-8"}
