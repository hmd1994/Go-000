# 学习笔记

## 作业

我们在数据库操作的时候，比如 `dao` 层中当遇到一个 `sql.ErrNoRows` 的时候，是否应该 `Wrap` 这个 `error`，抛给上层。为什么？应该怎么做请写出代码

## 思路

sql.ErrNoRows代表数据库中不存在相应的记录，当dao层遇到sql.ErrNoRows,不同的业务场景下处理方式不同。

sql属于第三方库，需要wrap相关堆栈信息，将相关堆栈信息返回给上层以后，上层对错误信息进行检查，看是否处理该错误