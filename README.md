### Introduce
- 仓库用于学习，
- 采用 Go 语言开发，实现各种数据结构，
- 熟练 Go 语言的基础
- 增加对数据结构的理解
这是本项目的初衷，有好的代码实现欢迎提 Issue 一起学习！

### 开发规范
1. 拉取代码
2. 创建分支(master分支下)
```shell
git branch 数据格式名称
git checkout 数据格式名称
```
3. 开始开发，最后合并到 master

### 创建新文件规范
- 例如想编写一个 队列（queue）的数据结构，那就先创建文件夹(queue)，里面文件名和文件夹名保持一致
- 每一种数据结构编写时，每个功能函数的完成都进行一次 commit，commit 规范是 ` git commit -m"queue -> functionName" `
- 每次完成一种数据结构，需要在 `main.go` 里测试基本功能即可，测试需要规范输出