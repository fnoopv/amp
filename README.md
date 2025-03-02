# Asset Management Platform (资产管理平台)

## 代码风格约定

- 字段,方法,复杂的逻辑都要写注释

## 数据库字段约定

- 数据库字段必须显示声明字段名称, 比如`gorm:"column:username"`
- 时间使用毫秒`milli`
- 唯一索引名称使用`uni_`开头, 比如`uni_username`
- 一般索引名称使用`idx_`开头, 比如`idx_username`
