facade pattern

info: facade (fə'sa:d) appearance or the face of some building.
      就是对象的外在, 表象, 以此遮盖内部的复杂, 此模式用来控制调用者的访问
      只暴露必要的接口.实际的脏活累活还是没有避免, 只是放到了facade 内部.
      
      就好比厨师要瓶油, 厨师不想也不关心油是怎么买来的(走路?骑车?中间出什么意外了?)
      只需要那瓶油.
      这样就可以整一个OilGetter以及其接口
      GetOil() (Oil, error)
      在内部实现取油的具体过程, 并且附加error handling.
      厨师就只需要执行OilGetter.GetOil(), 结果就是拿到了还是没拿到.遮盖了内部实现.

      这更像是Lib的工作模式, 调用者只要知道如何用就可以继续工作, 而不需要知道具体Lib实现.

note: 那岂不是所有的function 都是facade 模式了?同样符合隐藏实现, 简化调用的套路.
