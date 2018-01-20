creational pattern

tips: 创建对象需要很多步骤, 那么对于多种此种对象的场景, 先构建一个空壳, 在需要对象时,设置builder属性
      再获取对象, 执行construct构建builder中的Obj, 然后从Builder中获取OBJ

core: interface, setBuilder, construct

note: 最终还是通过Builder获取对象, 构建一个空壳, 设置具体的Builder, 执行共同的构建流程, 再从
      Builder中获取对象
      分离了细分Builder的具体实现, 将耦合放在了空壳子中
      
      [?] 将builder 设置到抽象出来的空壳中, 对象改变时, 同样Builder也改变, 这些都需要
      对应更改, 此模式并未改变此现状.
