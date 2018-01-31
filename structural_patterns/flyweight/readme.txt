flyweight pattern

info: flyweight 意为轻量, 减少内存和CPU资源的消耗(重复访问相同资源时)

      节约创建resource时消耗的CPU可以在创建对象后放到Cache中, 首次访问将resource
      添加到cache中,重复访问时直接返回cache value的拷贝, 不再重新创建.
      
      拷贝的CPU资源也不想耗?把cache中存放的value改为对象(resource)的引用
      多个client 访问相同资源时, 返回相同资源的ref, 这样每个资源在内存中没有多余的存在

      结构: 多个ref 指向同一个resource 内存空间
      
      如果内存空间有限, 那么将cache设置为固定大小优先队列:Queue(FIFO), 权重关系访问次数
      这样既保证了cache命中, 又限制了cache大小.

note: 用pointer(or reference) 来替代value copy, 结合cache, 同时减少了CPU 和MEM消耗
