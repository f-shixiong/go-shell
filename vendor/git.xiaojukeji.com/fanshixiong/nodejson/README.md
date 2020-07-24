golang 解析动态json

## example

```
packag main

import (
        "fmt"
        "git.xiaojukeji.com/fsx/nodejson"
)

func main() {
        data := []byte(`{"name":"fsx","dt":{"year":15},"dts":[{"11":22,"22":22},{"11":22,"22":33}]}`)
        n, _ := nodejson.UnmarshalToNode(data)
        year := n.Get("dt.year")
        fmt.Println(year.Int())
}
```

## function
>Get

获取一个节点的数据


>Set

update一个节点的数据
>Del

删除一个节点的数据
>GetKeySet

获取节点下子key


## TODO
#### 支持array
#### 性能测试



