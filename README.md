## go-shell
### how to install
install go
clone https://github.com/f-shixiong/go-shell.git
cd go-shell && go build
### how to use
go-shell [your file]


### 支持内容
>##### 新增功能
>> - [x] 类型强转
>> - [x] shell语法糖
>> - [ ] check关键字 from go2

>##### 运算符
>> - [x] +-*/
>> - [x] all
>>##### func
>> - [x] 自定义func
>> - [x] 结构体func
>> - [x] 递归func
>> - [x] return单个值
>> - [x] return多个值
>> - [ ] 跨域get set ?
>##### stmt关键字
>> - [x] for
>> - [x] switch todo
>> - [x] if
>> - [x] append
>> - [ ] delete
>> - [x] range
>> - [x] make
>> - [x] new
>> - [x] println
>> - [x] var
>> - [x] const
>> - [x] break
>> - [ ] defer
>> - [ ] goto
>##### 数据类型
>> - [x] int,string
>> - [x] map,list
>> - [x] char,bool,int8-64,uint8-64,[]byte
>> - [ ] struct !difficult
>##### import
>> - [x] 非golang官方包
>> - [ ] 官方带internal package

>##### repl
>> - [ ] 支持repl

>##### goruntime
>> - [ ] 支持go
>> - [ ] 支持select

>##### 断言
>> - [ ] 支持断言

### 更新日志
> - 2.10-just can run,支持了基本+-*/
> - 2.12-支持自定义func,支持复杂数据类型,支持for循环,支持了println,支持了new、make、range
> - 2.13-支持append,2.13-支持自定义结构体TODO 非map实现，支持多层结构体
> - 2.14-支持指针
> - 2.15-支持return，支持import，支持struct定义func
> - 2.16-支持动态import
> - 2.17-支持了shell,支持return 多个值，什么都没改自己支持了const，基本靠命,支持+=等
> - 2.18-什么都没改自己支持了const，基本靠命,支持+=等,支持了所有数据结构，所有的运算符
> - 2.19-支持了break关键字
