关于超时时间的公共设置
1. 对内网微服务和其它部门的api,尽量控制在100ms之内,最差不能超过300ms
2. 对于公网暴露的api,需要控制在1s之内

为了对于超时部分进行控制,避免在业务变更迭代过程中导致下游的响应时间逐渐变长,可以
1. 在api的注释中说明该api的接口响应时间的标准,
   如果是grpc的话,可以直接在其对应的proto文件上进行配置
   // lagency SLO: 95th in 100ms, 99th in 150ms

在同进程quota的控制 用户 -> service A -> service B -> service C
1. 现象
    从service A的角度来观察,有1s的超时时间,在进行分配的时候,各个耗时操作,例如:调用service B,调用redis,调用mysql,
    这三个操作都有其本身的超时时间,但同时,在计算整体的超时时间之中,也有一个规划,这样对于service A的redis调用来说,实际的超时时间就是本
    阶段原计划的超时时间和目前整体剩余超时时间的min值,这样可以做到更细粒度的控制整体的超时时间,
2. 解决方法和影响
    而控制方法就是contenx,在整体和每个调用的部分都有一个context,这样,对于请求内部的操作就可以控制在一定的时间范围之内,而不是会导致虽然client已经因为超时策略而放弃请求,但是对于服务本身而言,仍然处在阻塞的状态,进而拉低整个服务的处理速度,导致更大量的堆积甚至进程crash

在多个进程quota的控制 用户 -> service A -> service B -> service C
1. 控制的方法
    对于service A在被请求的时候超时时间为1s,这样如果它本身在开始处理的时候已经使用了300ms,那么,当它调用service B的时候,其实超时时间只剩下了700ms,对于这样的情况,
    我们可以在grpc的metadata中传递它最大的响应时间,因为grpc是基于http 2.0实现的,所以本质上,超时时间是在header中的一个字,在下游服务service B收到请求的时候,利用参数初始化自己的超时时间,对于向service C传递的时候,也使用同样的策略;对于在service B已经超时或者接近超时的请求,可以在service B直接进行time out,这样对整体的超时进行了控制
2. 注意
    对于在进程之间传递的超时时间,我们一般会减去5/10ms,因为网络连接和ping需要时间

双峰分布: 95%的请求耗时在100ms内,5%的请求可能永远不会完成(长超时)
对于监控不要只看mean,可以看耗时分布的统计,比如95th,99th;一般来讲,只要能保证95/99线都在100ms之内,基本就可以认为整体是安全的
设置合理的超时,拒绝超长请求,或者当Server不可用要主动失败.

入口处的nginx也要配置proxy timeout,这样可以避免请求粘连到nginx上,并堆积,最终导致整个系统的入口端失效
nginx超时的整体配置: https://juejin.cn/post/6844903652964958216
nginx超时配置的demo: https://cloud.tencent.com/developer/article/1382040

*超时决定着服务线程耗尽