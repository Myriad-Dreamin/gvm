# gvm
Extensible virtual machine.

## 为什么没有做成字节码/使用栈分配变量？
这个库暂时肯定只有我一个人用，而这个VM将实验性地运行在一个高代价的数据库上，也就是说目前的模型不基于内存。如果那样去做，不如自行寻找一个已经成熟的VM拿来用。
