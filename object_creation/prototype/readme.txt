prototype pattern

info: 节约构建对象时的性能消耗, 在程序载入时, 创建好prototype OBJ, 通过Clone的方式获取新对象,
      在具体属性上进行单独设置, 互不影响

core: interface, cache, Cloner

note: 理解为静态变量, 新对象都是对该静态变量的copy by value
      构建prototype 对象的消耗和在GetClone() 时通过copy by value 的方式获取新对象的方式
      两者性能节约有多大?
      在test 中, 通过clone方式获取对象反而慢了一些, 但两者基本相同
