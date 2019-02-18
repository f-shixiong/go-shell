## go-shell
### how to install
install go
clone https://github.com/f-shixiong/go-shell.git
cd go-shell && go build
### how to use
go-shell [your file]

### 支持内容
#### 运算符
> - [x] +-*/
> - [ ] all
#### func
> - [x] 自定义func
> - [x] 结构体func
> - [ ] 递归func
> - [x] return单个值
> - [x] return多个值
> - [ ] 跨域get set ?
#### stmt关键字
> - [x] for
> - [x] switch todo
> - [x] if
> - [x] append
> - [ ] delete
> - [x] range
> - [x] make
> - [x] new
> - [x] println
> - [x] var
> - [x] const
> - [ ] break
> - [ ] defer
> - [ ] goto
#### 数据类型
> - [x] int,string
> - [x] map,list
> - [ ] all
> - [ ] struct !difficult
#### import
> - [x] 非golang官方包
> - [ ] 官方带internal package

#### candy
> - [x] 支持@shell@直接执行shell
> - [ ] 支持@shell with var@执行shell
> - [ ] 支持check关键字

#### repl
> - [ ] 支持repl

#### goruntime
> - [ ] 支持go
> - [ ] 支持select

#### 断言
> - [ ] 支持断言

### 更新日志
> - 2.10-just can run,支持了基本+-*/
> - 2.12-支持自定义func,支持复杂数据类型,支持for循环,支持了println,支持了new、make、range
> - 2.13-支持append,2.13-支持自定义结构体TODO 非map实现，支持多层结构体
> - 2.14-支持指针
> - 2.15-支持return，支持import，支持struct定义func
> - 2.16-支持动态import
> - 2.17-支持了shell,支持return 多个值，什么都没改自己支持了const，程序能run全靠命
