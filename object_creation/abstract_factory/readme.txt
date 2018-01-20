abstract factory pattern

tips: 基于factory pattern, 但是又添加了一层抽象聚合关系, 工厂模式的嵌套
      再通过get方法来根据传入的factory id 获取不同的factory, 继而再通过对Getter传入不同的
      ObjectId获取不同的对象

core: interface, group and route
      通过两层pointer(typeId)来得到对应的object

note: abstract factory pattern 是对factory pattern 的进一步概括, 双层嵌套, 这表明需要的对象
      之间的共性之余, 还存在着诸如应用场景,环境之类的不同
      例如跨平台的控件, 同样完成一个button操作,但存在Windows, Linux, MacOS 之区别, 所以添加这样
      一层抽象是合理的.                     
