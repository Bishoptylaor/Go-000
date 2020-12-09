## 一、goroutine

### 1. 基础知识

- 应用程序 vs 进程
- 进程 vs multi 线程
- 进程 = 主线程 + 主线程启动的其他线程 + 其他人启动的更多线程

### 2. goroutines

- go关键字创建goroutine，可以认为main就是作为一个特殊的goroutine来执行的
- 操作系统调度线程在可用处理器上运行，Go运行时调度 goroutines 在绑定到单个操作系统线程的逻辑处理器中运行(P)。即使使用这个单一的逻辑处理器和操作系统线程，也可以调度数十万 goroutine 以惊人的效率和性能并发运行。

### 3. Tips

- *Concurrency is not Parallelism.*
- *Keep yourself busy or do the work yourself*
- *Leave concurrency to the caller*
- *Never start a goroutine without* *knowning* *when it will stop*
  - When will it terminate
  - What could pervent it from terminating?
- https://github.com/da440dil/go-workgroup
- *Incomplete Work*
  - 使用 sync.WaitGroup 来追踪每一个创建的 goroutine
  - ![image-20201208174607881](/Users/bishop/Library/Application Support/typora-user-images/image-20201208174607881.png)
  - ![image-20201208174552675](/Users/bishop/Library/Application Support/typora-user-images/image-20201208174552675.png)

## 二、memory model

### Memory Reordering

<img src="/Users/bishop/Library/Application Support/typora-user-images/image-20201208230349193.png" alt="image-20201208230349193" style="zoom:50%;" />

<img src="/Users/bishop/Library/Application Support/typora-user-images/image-20201208230404856.png" alt="image-20201208230404856" style="zoom:50%;" />

<img src="/Users/bishop/Library/Application Support/typora-user-images/image-20201208230413605.png" alt="image-20201208230413605" style="zoom:50%;" />

- *对于多线程的程序，所有的* *CPU* 都会提供“锁”支持，称之为*barrier*，或者* *fence*。它要求：*barrier* 指令要求所有对内存的操作都必须要“扩散”到 *memory* 之后才能继续执行其他对 *memory* 的操作。因此，我们可以用高级点的 *atomic* compare-and-swap，或者直接用更高级的锁，通常是标准库提供。
- 当多个 goroutine 访问共享变量 v 时，它们必须使用同步事件来建立先行发生这一条件来保证读操作能看到需要的写操作。 
- https://www.jianshu.com/p/5e44168f47a3

## 三、package sync

### Share Memory By Communicating

- Do not communicate by sharing memory; instead, share memory by communicating
- Go 的并发原语 goroutines 和 channels 为构造并发软件提供了一种优雅而独特的方法。Go 没有显式地使用锁来协调对共享数据的访问，而是鼓励使用 chan 在 goroutine 之间传递对数据的引用。这种方法确保在给定的时间只有一个goroutine 可以访问数据。

### Detecting Race Conditions With Go

- data race 是两个或多个 goroutine 访问同一个资源(如变量或数据结构)，并尝试对该资源进行读写而不考虑其他 goroutine。这种类型的代码可以创建您见过的最疯狂和最随机的 bug。通常需要大量的日志记录和运气才能找到这些类型的bug。
- go 工具检查数据竞争 race detector 
  - go build -race  /  go test -race
- 使用Go同步语义：mutex
- 指针包含type和data两个部分

## 四、chan

## 五、context

## 六、引用

