factory pattern

场景: 区分同方法的多实现, 解耦代码
作用: 分离具体方法实现, 在代码更新时,只要添加新的interface实现, 并更新下GetMethodImplement
中的代码即可.

tips: 可见, 在添加NewCashPM 后, 有了新的Pay实现,在GetPaymentMethod方法中, 更新代码, 即可使用
新实现.脱离了具体Pay方法实现的捆绑.

实现: 通过interface 来实现传入参数的泛化, 也就是将具有通性的操作抽象为一个通用的接口(interface)
      将这些接口实现汇聚到一个Map里, 在一个GetMethod方法中传入index来获取对应的method
      将方法实现的耦合放在了GetMethod中.
      抽象 --> 整合 --> 使用时抽取对应对象
