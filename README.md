## go-shell
### how to install
install go
clone https://github.com/f-shixiong/go-shell.git
cd go-shell && go build
### how to use
go-shell [your file]


### 支持内容
>##### 新增功能
```
√  类型强转
√  shell语法糖
× check关键字 from go2
```
>##### 运算符
```
√ + - * / & % ^ && := = ...  
```

>##### func
```
√ 自定义func,结构体func,递归func,return单个值,return多个值,跨域get set 
```
>##### stmt关键字
```
√ for switch if append delete range make new println var const break
```
```
× defer goto [...] go select
```
>##### 数据类型
```
× int,string,map,list,char,bool,int8-64,uint8-64,[]byte
```
```
× struct !difficult
```
>##### import
```
√ 支持了非golang官方包
×  官网引用internal的package
```
>##### 断言
```
× 支持隐式强转，不需要断言
```

### 更新日志
```
19.02.10-just can run,支持了基本+-*/
19.02.12-支持自定义func,支持复杂数据类型,支持for循环,支持了println,支持了new、make、range
19.02.13-支持append,2.13-支持自定义结构体TODO 非map实现，支持多层结构体
19.02.14-支持指针
19.02.15-支持return，支持import，支持struct定义func
19.02.16-支持动态import
19.02.17-支持了shell,支持return 多个值，什么都没改自己支持了const，基本靠命,支持+=等
19.02.18-什么都没改自己支持了const，基本靠命,支持+=等,支持了所有数据结构，所有的运算符
19.02.19-支持了break关键字,支持delete
```
