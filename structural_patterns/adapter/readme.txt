adapter pattern

info: 两个接口不兼容, 但又必须一起工作, 这就可以考虑使用adapter pattern 来衔接两者

tips: Open and Close principle
      在设计之初就要考虑到代码的拓展性, 但在实际进展中, 又要尽可能避免对旧代码(已经测试上线)的修改
      以避免麻烦(这难道是偷懒模式?)                                        
      Open: write code that not only work for that one thing.
      Close: try not to change any old (both yours and others) as we can.

note:在struct 中嵌套old interface, 新接口方法中进行方法执行的路径选择, 如果嵌套了旧结构,
     执行旧方法, 如果没有, 那么执行新接口方法.

     新接口要保持对旧接口的兼容, 那么维护一个或几个field 来承载旧接口中的信息, 再整合到新接口
     实现中.在调用中不应该出现是否为旧接口信息.(透明)

go src:
   net/http

   // The HandlerFunc type is an adapter to allow the use of
   // ordinary functions as HTTP handlers. If f is a function  
   // with the appropriate signature, HandlerFunc(f) is a 
   // Handler that calls f.

   type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
   }
   
   type HandleFunc func(ResponseWriter, *Request)
   
   func (h HandleFunc) ServeHTTP(w ResponseWriter, r *Request) {
        h(w, r)
   }

   func() 也是Golang 中的一等公民, 和struct 类似, 可以实现接口: Handler

   具有 (ResponseWriter, *Request) 签名的方法就可以转化为HandleFunc类型,
   而该类型实现了ServeHTTP 方法, 也就是实现了Handler接口, 传递而言, 自定义特殊签名
   的方法也就实现了Handler接口
   
   HandleFunc(personalFunc), 将特定签名的personalFunc 转换为了Handler
