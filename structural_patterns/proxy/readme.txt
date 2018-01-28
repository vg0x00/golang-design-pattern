proxy pattern

info: proxy pattern需要和原interface 接口相同, 使用者应看不出调用接口上有什么不同(透明),
      由于Proxy这一层的加入, 可以添加诸如cache, authentication之类的中间操作, 同时, 也
      可以对原接口进行掩盖(删减)
      
sample: 接口A用于访问数据库, 通过Proxy提供一个Cache, 两者接口相同, 如果cache命中, 那么返回
        cache中的数据, 否则, 从数据库中获取.
        维护cache size 不超过设定值.

unit test: 组合test, 在一个Testxxxx function 中嵌套多个测试实例
           t.Run("title for this sub-test", func(t *testing.T) {
               // testing code here
               // ...
           })
