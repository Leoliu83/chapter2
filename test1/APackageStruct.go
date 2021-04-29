package test1

/*
	依照规范，工作空间由src，bin，pkg三个目录组成，通常需要将空间路径添加到GOPATH环境变量列表中，以便相关的工具能够正常工作
	在工作空间中，包括子包在内的所有源码文件都保存在src目录下，至于pin和pkg两个目录，主要影响go install/get 命令，它们会将编译
	结果安装到这两个目录下，以实现增量编译。
	目录结构如下：
	workspace/
		|
		+ —— src/  <-- 源码
		|     |
		|     + —— server/ <--  main 必须在main包下，不知为何这里是server？
		|     |       |
		|     |       + —— main.go
		|     |
		|     + —— service/
		|             |
		|             + —— user.go
		|
		+ —— bin/     <-- 可执行文件安装路径，不会创建额外的子目录
		|     |
		|     + —— server.exe
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

/*
	包名通常用单数形式，源码必须使用utf8格式，否则编译出错
	包路径必须唯一，包名可重复
	特殊含义包名：
	main         : 可执行入口，入口函数（main.main）
	all    		 : 标准库以及GOPATH中所能找到的所有包
	std,cmd      : 标准库及工具链
	documentation: 存储文档信息，无法导入
*/

/*
	权限：
	所有成员在包内均可访问，无论是否在同一源码中。单只有首字母大写的成员可以被导出，在包外可见
	该规则适用于：全局变量，全局常量，类型，结构字段，函数，方法等。
	但是可以通过指针转换等方式绕开
*/

/*
	vendor 专门用来存放第三方包，第三方依赖包在go.mod中定义，可以直接通过vscode工具直接生成vendor目录
	同名包的导入规则：
	当目录如下显示：
	   src/
		|
		+ —— server/
		        |
		        + —— vendor/
		        |      |
		        |      + —— p/  <-- 这个称为 p1
		        |      |
		        |      + —— x/
		        |           |
		        |           + —— test.go
		        |           |
		        |           + —— vendor/
		        |                   |
		        |                   + —— p/  <-- 这个称为 p2
		        |
		        + —— main.go
	目录中有两个名为 p/ 的包，在main.go 和 test.go 中分别导入p时，各自对应谁？
	规则是：从当前源文件的目录开始逐级向上构造vendor全路径，直到发现匹配的目标为止。匹配失败，则依旧搜索GOPATH
	对于main.go 而言，当前目录为server/，可构造出的路径是 src/server/vendor/p，也就是p1
	而对test.go 而言，当前目录为x/ 最先构造出的路径为 src/server/vendor/x/vendor/p，也就是p2
	在go1.5以下，若要使用vendor机制，必须设置环境变量GO15VENDOREXPERIMENT=1, >=1.5则默认开启
	使用go get 下载第三方包时，依旧使用GOPATH第一个工作空间，而非vendor。
*/
