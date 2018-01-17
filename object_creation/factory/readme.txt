factory pattern

场景: 区分同方法的多实现, 解耦代码
作用: 分离具体方法实现, 在代码更新时,只要添加新的interface实现, 并更新下GetMethodImplement
中的代码即可.

tips: 可见, 在添加NewCashPM 后, 有了新的Pay实现,在GetPaymentMethod方法中, 更新代码, 即可使用
新实现.脱离了具体Pay方法实现的捆绑.
