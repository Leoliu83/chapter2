#### 如何在go中获取admin权限
1. go get github.com/akavel/rsrc
2. 把nac.manifest 文件拷贝到当前windows项目根目录
3. rsrc -manifest nac.manifest -o nac.syso
4. go build

nac.mainfest的内容为：
UAC管理员权限
```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
<assemblyIdentity
    version="9.0.0.0"
    processorArchitecture="x86"
    name="myapp.exe"
    type="win32"
/>
<description>My App</description>
<trustInfo xmlns="urn:schemas-microsoft-com:asm.v3">
    <security>
        <requestedPrivileges>
            <requestedExecutionLevel level="requireAdministrator" uiAccess="false"/>
        </requestedPrivileges>
    </security>
</trustInfo>
</assembly>
```

生成syso文件，将生成的main.syso文件拷贝到main.go同级目录
```shell
rsrc -manifest main.exe.manifest -ico rc.ico -o rsrc.syso
```


编译生成main.exe  
```shell
go build -o main.exe
```

如果是GUI程序，则需要隐藏命令行窗口
```shell
go build -ldflags="-H windowsgui"
```

"-w"是裁剪gdb调试信息，这样生成的exe体积会小一些。文件会从7.6M精简到6.2M。
```shell
go build -o main.exe -ldflags="-w"
```

其中-w为去掉调试信息（无法使用gdb调试），-s为去掉符号表（暂未清楚具体作用）。文件大小变为5.71M。

‘-s’ 相当于strip掉符号表， 但是以后就没办法在gdb里查看行号和文件了。
‘-w’ flag to the linker to omit the debug information 告知连接器放弃所有debug信息  
```shell
go build -o main.exe -ldflags="-w -s"
```