package routes

var Temp []Route

type Route struct {
	Method, Pattern string
	Handler         []interface{}
}

func Get(pattern string, handler ...interface{}) {
	Temp = append(Temp, Route{"GET", pattern, handler})
}
func Post(pattern string, handler ...interface{}) {
	Temp = append(Temp, Route{"POST", pattern, handler})
}
func Put(pattern string, handler ...interface{}) {
	Temp = append(Temp, Route{"PUT", pattern, handler})
}
func Patch(pattern string, handler ...interface{}) {
	Temp = append(Temp, Route{"PATCH", pattern, handler})
}
func Delete(pattern string, handler ...interface{}) {
	Temp = append(Temp, Route{"DELETE", pattern, handler})
}
func Options(pattern string, handler ...interface{}) {
	Temp = append(Temp, Route{"OPTIONS", pattern, handler})
}
func Head(pattern string, handler ...interface{}) {
	Temp = append(Temp, Route{"HEAD", pattern, handler})
}
func Any(pattern string, handler ...interface{}) {
	Temp = append(Temp, Route{"*", pattern, handler})
}
