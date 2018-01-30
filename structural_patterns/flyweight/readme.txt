flyweight pattern

info: flyweight 意为轻量, 减少内存和CPU资源的消耗(在重复访问时)
      节约CPU可以在创建对象后放到Cache中, 那么在访问时返回cache value的拷贝
      
      拷贝的CPU资源也不想耗?把cache中存放的Key改为对象的引用
      多个client 访问相同资源时, 返回相同资源的ref, 这样每个资源在内存中没有多余的存在
      
      如果内存空间有限, 那么将cache设置为固定大小附带权重的Queue, 权重就是访问次数
      这样既保证了cache命中, 又限制了cache大小.

note: 用pointer(or reference) 来替代value copy, 结合cache, 同时减少了CPU 和MEM消耗
