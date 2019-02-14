package lib

func BAdd(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l + r.(int)
	default:
		Error("anyway it happend,type= %#v \n", l)
	}
	return nil
}
