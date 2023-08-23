
## tips
### return语句后面没有显式地指定返回值
```go
// GetDailyAlgorithmRecord 根据id获取DailyAlgorithmRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (DARService *DailyAlgorithmRecordService) GetDailyAlgorithmRecord(id uint) (DAR DailyAlgorithm.DailyAlgorithmRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&DAR).Error
	return
}
```
在GetDailyAlgorithmRecord方法中，return语句后面没有显式地指定返回值。这是因为在函数定义中已经定义了函数的返回类型。

在这种情况下，函数会自动将函数内部定义的局部变量作为返回值返回。在GetDailyAlgorithmRecord方法中，DAR和err变量是在函数内部定义的，并且它们的类型与函数的返回类型匹配。

因此，当执行return语句时，函数会自动将DAR和err变量的值作为返回值返回，无需显式指定。

这种方式使得代码更加简洁，同时也符合Go语言的函数返回值规范。调用方可以使用多重赋值的方式接收返回值。