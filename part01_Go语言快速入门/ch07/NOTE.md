# 07 | 错误处理：如何通过 error、deferred、panic 等处理错误？

## 思考题

人既会走也会跑

``` go
type WalkRun interface {
    Walk()
    Run()
}
```

现在就可以让结构体 person 实现这个接口了

``` go
func (p *person) Walk() {
    fmt.Printf("%s能走\n", p.name)
}
func (p *person) Run(){
   fmt.Printf("%s能跑\n",p.name)
}
```
