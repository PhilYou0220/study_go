package main

//new() 和 make() 都是 golang 中用来分配内存的函数，本文介绍了 make() 和 new() 的三个主要的区别：
//
//接收参数个数不一样：new() 只接收一个参数，而 make() 可以接收多个参数
//返回类型不一样：new() 返回一个指针，而 make() 返回类型和它接收的第一个参数类型一样
//应用场景不一样：make() 专门用来为 slice、map、chan 这样的引用类型分配内存并作初始化，而 new() 用来为其他类型分配内存

//为何引用类型不能用new()
//如果用 new() 来创建 slice，那么创建的 header 中的 pointer 做0值处理，就会被初始化为 nil，
//而 length 和 capacity 也会被初始化为0，这样显然是不正确的