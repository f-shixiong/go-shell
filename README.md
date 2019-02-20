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
- [x] 类型强转
- [x] shell语法糖
- [ ] check关键字 from go2
```
>##### 运算符
```
+ - * / & % ^ && := = ...  
```

>>##### func
```
自定义func,结构体func,递归func,return单个值,return多个值,跨域get set 
```
>##### stmt关键字
```
- [x] for switch if append delete range make new println var const break
```
```
- [ ] defer goto [...] go select
```
>##### 数据类型
```
- [x] int,string,map,list,char,bool,int8-64,uint8-64,[]byte
```
>> 未支持
```
struct !difficult
```
>##### import
```
支持了非golang官方包, todo:官网引用internal的package
```
>##### 断言
```
支持隐式强转，不需要断言
```

### 更新日志
```
- 2.10-just can run,支持了基本+-*/
- 2.12-支持自定义func,支持复杂数据类型,支持for循环,支持了println,支持了new、make、range
- 2.13-支持append,2.13-支持自定义结构体TODO 非map实现，支持多层结构体
- 2.14-支持指针
- 2.15-支持return，支持import，支持struct定义func
- 2.16-支持动态import
- 2.17-支持了shell,支持return 多个值，什么都没改自己支持了const，基本靠命,支持+=等
- 2.18-什么都没改自己支持了const，基本靠命,支持+=等,支持了所有数据结构，所有的运算符
- 2.19-支持了break关键字,支持delete
```
