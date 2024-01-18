# golang-10x-engineer

## 010_gift_并发

### 资源竞争和原子操作

- 多协程并发修改同一块内存，产生资源竞争
- n++不是原子操作，并发执行时会存在脏写
  - 因为n++实际上是三个步骤：取出n，加1，写入n
  - 测试时需要开1000个并发协程才能观察到脏写
- 把n++封装成原子操作，解除资源竞争，避免脏写
  - func atomic.AddInt32(addr *int32, delta int32)(new int32)
  - func atomic.LoadInt32(addr *32) (val int32)