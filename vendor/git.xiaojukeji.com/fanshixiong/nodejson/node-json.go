package nodejson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Node struct {
	data interface{}
}

func UnmarshalToNode(data []byte) (Node, error) {
	node := Node{}
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.UseNumber()
	err := dec.Decode(&node.data)
	if err != nil {
		fmt.Println(err)
		return Node{}, err
	}
	return node, nil
}

func (n *Node) Marshal() ([]byte, error) {
	return json.Marshal(n.data)

}

func (n *Node) Get(link string) Node {
	if node, err := n.GetE(link); err != nil {
		return Node{}
	} else {
		return node
	}
}

func (n *Node) GetE(link string) (Node, error) {
	linkarr := strings.Split(link, ".")
	l := len(linkarr)
	var data = n.data
	for i, k := range linkarr {
		if i == l {
			if data == nil {
				return Node{}, fmt.Errorf("can not found %s", k)
			}
			return Node{data: data}, nil
		}
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			return Node{}, fmt.Errorf("node %v is last node", data)
		}
		data = dataMap[k]
	}
	return Node{data: data}, nil
}

func (n *Node) GetKeySet() []string {
	if ret, err := n.GetKeySetE(); err != nil {
		return make([]string, 0)
	} else {
		return ret
	}
}

func (n *Node) GetKeySetE() ([]string, error) {
	dataMap, ok := n.data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("node %v is last node", n.data)
	}
	var keys = make([]string, 0)
	for k, _ := range dataMap {
		keys = append(keys, k)
	}
	return keys, nil
}

func (n *Node) Del(link string) {
	n.DelE(link)
}

func (n *Node) DelE(link string) error {
	linkarr := strings.Split(link, ".")
	l := len(linkarr) - 1
	var data = n.data
	for i, k := range linkarr {
		if i == l {
			dataMap, ok := data.(map[string]interface{})
			if !ok {
				return fmt.Errorf("node %v is last node", k)
			}
			delete(dataMap, k)
			return nil
		}
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			return fmt.Errorf("node %v is last node", data)
		}
		data = dataMap[k]
	}
	return nil
}

func (n *Node) Set(link string, v interface{}) {
	n.SetE(link, v)
}

func (n *Node) SetSubNode(link string, v Node) {
	n.SetE(link, v.data)
}

func (n *Node) SetE(link string, v interface{}) error {
	linkarr := strings.Split(link, ".")
	l := len(linkarr) - 1
	var data = n.data
	if l == 0 {
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			dataMap = make(map[string]interface{}, 0)
			n.data = dataMap
		}
		dataMap[link] = v
		n.data = dataMap
	} else {
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			dataMap = make(map[string]interface{}, 0)
			n.data = dataMap
		}
		next := n.Get(linkarr[0])
		if _, ok2 := next.data.(map[string]interface{}); !ok2 {
			subMp := make(map[string]interface{}, 0)
			n.Set(linkarr[0], subMp)
		}
		next.SetE(strings.Join(linkarr[1:l+1], "."), v)
		n.SetSubNode(linkarr[0], next)
	}
	return nil
}

func (n *Node) IsEmpty() bool {
	return n.data == nil
}

func (n *Node) IsArray() bool {
	_, ok := n.data.([]interface{})
	return ok
}

func (n *Node) IsMap() bool {
	_, ok := n.data.(map[string]interface{})
	return ok
}
