## go-shell
### Installation 
```
git clone https://github.com/f-shixiong/go-shell.git 
cd go-shell
go build
```
### Getting Started
```
go-shell -f [your file]
go-shell -f "println(1)"
go-shell 
```


### support
>##### what new
![](https://github.com/f-shixiong/go-shell/blob/master/samples/go-shell_1.gif?raw=true)
>
```
√  类型强转    : example -> a := "1"   b := int(a)   b = 1  
√  candy_shell : use shell like  var a = @echo 'valueA'@ -> then a = "valueA"
×  check关键字 from go2
```

### plan
> - gopiter like jupyter 
> - go-i go-repl
> - shell-tools : support use go in shell 

>##### diff from go-compile
```
√  + - * / & % ^ && := = ...  
√  自定义func,结构体func,递归func,return单个值,return多个值,跨域get set 
√  for switch if append delete range make new println var const break
×  defer goto [...] go select
√  int,string,map,list,char,bool,int8-64,uint8-64,[]byte
×  struct !difficult
√  非golang官方包,支持引用gopath路径
×  官网引用internal的package
×  支持隐式强转，不需要断言
```

### update log
```
19.02.10-just can run,支持了基本+-*/
19.02.12-支持自定义func,支持复杂数据类型,支持for循环,支持了println,支持了new、make、range
19.02.13-支持append,-支持自定义结构体 TODO 需要非map实现，支持了多层结构体
19.02.14-支持指针
19.02.15-支持return，支持import，支持struct定义func
19.02.16-支持动态import
19.02.17-支持了shell,支持return 多个值，什么都没改自己支持了const，基本靠命,支持+=等
19.02.18-什么都没改自己支持了const，基本靠命,支持+=等,支持了所有数据结构，所有的运算符
19.02.19-支持了break关键字,支持delete
19.02.21-使用mod@raojian,gen internal_package解决internal无法被引用问题
19.02.22-支持了shell使用管道
19.02.22-支持并行串行shell
19.02.24-支持了repl
```
