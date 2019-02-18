## go-shell
### how to install
install go
clone https://github.com/f-shixiong/go-shell.git
cd go-shell && go build
### how to use
go-shell [your file]

### 支持项
#### 基本+-*/
> - [x] +-*/
> - [ ] all
#### func
> - [x] 自定义func
> - [x] 结构体func
> - [ ] 递归func
> - [x] return单个值
> - [x] return多个值
#### stmt关键字
> - [x] for
> - [ ] switch
> - [ ] if
> - [x] append
> - [ ]
#### 数据类型
> - [x]int,string
> - [ ]all
> - [ ]struct
#### import
> - [x]非golang官方包
> - [ ]官方带internal package

#### candy
> - [x] 支持@shell@直接执行shell
> - [ ] 支持check关键字

### 支持日志
>> 2.10-支持基本+-*/
>> 2.12-支持自定义func
>> 2.12-支持复杂数据类型
>> 2.12-支持for循环
>> 2.13-支持append
>> 2.13-支持自定义结构体 TODO 非map实现
>> 2.13-支持多层结构体
>> 2.14-支持指针
>> 2.15-支持return
>> 2.15-支持import
>> 2.15-支持struct-func
>> 2.16-支持动态import
>> 2.17-支持shell
>> 2.17-支持return 多个值
