/*
	依照规范，工作空间由src，bin，pkg三个目录组成，通常需要将空间路径添加到GOPATH环境变量列表中，以便相关的工具能够正常工作
	在工作空间中，包括子包在内的所有源码文件都保存在src目录下，至于pin和pkg两个目录，主要影响go install/get 命令，它们会将编译
	结果安装到这两个目录下，以实现增量编译。
	目录结构如下：
	workspace/
		|
		+ —— src/  <-- 源码
		|     |
		|     + —— server/
		|     |       |
		|     |       + —— main.go
		|     |
		|     + —— service/
		|             |
		|             + —— user.go
		|
		+ —— bin/     <-- 可执行文件安装路径，不会创建额外的子目录
		|     |
		|     + —— server
		|
		+ —— pkg/   <-- 包安装路径，按操作系统和平台分离
		      |
			  + —— linux_amd64/
			            |
						+ —— service.a

*/

/*
	环境变量：
	GOPATH: 编译器等相关工具按照GOPATH设置的路径搜索目标，也就是说在导入目标库时，排在列表前的路径比当前路径优先级更高，go get会将下载的第三方包存放在列表的第一个存放路径内
	GOROOT: 指示的是工具链和标准库的位置，也就是golang的安装路径
	GOBIN: 强制替代工作空间的bin目录，作为 go install 的目标保存路径
	对于每个项目设置不同的环境变量，可以运用类似 python 的 virtual environment，一般可以在激活某个项目时，自动设置相关环境变量
*/
/*
	导入包：
	在使用标准库或者第三方包之前，必须使用import导入，参数是工作空间中以src为起始的绝对路径。编译器从标准库开始搜索，依次搜索GOPATH中的各个工作路径
	导入包的方式主要可以归纳成以下几种：
	import "abc/def" 默认方式: def.Func()
	import x "abc/def" 别名方式: x.Func()
	import . "abc/def" 简便方式: Func()
	import _ "abc/def" 初始化方式，不能调用，仅仅用作初始化目标包
	import "../def" 相对路径方式: def.Func()
	注意：在设置了GOPATH的工作空间中，相对路径会导致编译失败。查看GOPATH命令：go env GOPATH
*/