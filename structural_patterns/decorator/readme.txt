decorator pattern

decorator 意为装饰, 装饰意味着添加新内容, 而被装饰的对象就是真正的core obj

info: 和proxy类似, 将core obj 内嵌到doll obj中, 在外层的doll obj中进行功能拓展,
      两者在内层core obj和外层的doll obj 都要实现相同的接口(interface)
      decorator: 一般都是调用者手动设定decorator constructor的参数, 也就是将core obj inject
                 到decorator obj中, 继而在decorator obj中实现功能拓展.
      proxy: 一般都是proxy 提前设定好内部需要Proxy的Obj, 调用者直接使用proxy obj, 觉察不出
             与直接调用原Obj之间的区别, 如果进行了功能控制, 减少了接口访问, 但也只能哀叹, 没有
             他想要的拓展功能(当然没有! 给你屏蔽了)

diff:  proxy 和 decorator 这两个模式, 在实现上很类似, 都是将core obj 进行封装, 放置到外层的
       doll obj中, 并且经由外部的doll obj提供API服务.

       主要区别在于两者的意图:
       proxy 在于修剪interface, 暴露给使用者仅需的接口, 控制访问者对内部core obj 的访问范围.
       而decorator是通过这种内嵌原始对象, 通过外层doll obj 进行功能上的拓展.

       两者都是实现内部obj接口, 那么都可以实现proxy 嵌套或者decorator 嵌套. 如果外层接口改了,
       变身为Adaptor 模式
