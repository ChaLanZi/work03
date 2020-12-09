# work03

基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。

实验结果 control +c 打印：
^Chttp服务内部goroutine退出
关闭服务来源：通过信号关闭http服务
Process finished with exit code 0

