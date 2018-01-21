composite pattern

info: struct 中嵌套struct或者interface
      前者可利用zero-initialization
      后者需要手动指定指针
      解决逻辑单元的嵌套, 包含关系

note: struct 嵌套可以显式指明fieldName, 也可以省去fieldName
      struct 也可嵌套function
      
      type Swimmer struct {
        Swim()
      }

      指明field是一个参数和返回都为空的func
