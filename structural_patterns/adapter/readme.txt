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
