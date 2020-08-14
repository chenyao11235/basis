package prototype

/*原型模式
在golang中通常创建一个对象的方法就是：
	new
	make
而在原型模式中是通过clone现有的对象来创建新的对象(本质也是通过new一个新对象，然后把现有对象的属性copy过去)
适用于 对象的创建性能开销比较大， 比如： 对象保存数据库中10万条数据，每新建一个对象就需要从数据库中查询这10万条数据
这时候clone已存在的对象就再合适不过了
*/

// 从来没用过原型模式，这是gorm中的，看看它是怎么实现的

//Statement 不知名对象

// func (stmt *Statement) clone() *Statement {
// 	newStmt := &Statement{
// 		Table:                stmt.Table,
// 		Model:                stmt.Model,
// 		Dest:                 stmt.Dest,
// 		ReflectValue:         stmt.ReflectValue,
// 		Clauses:              map[string]clause.Clause{},
// 		Distinct:             stmt.Distinct,
// 		Selects:              stmt.Selects,
// 		Omits:                stmt.Omits,
// 		Joins:                map[string][]interface{}{},
// 		Preloads:             map[string][]interface{}{},
// 		ConnPool:             stmt.ConnPool,
// 		Schema:               stmt.Schema,
// 		Context:              stmt.Context,
// 		RaiseErrorOnNotFound: stmt.RaiseErrorOnNotFound,
// 	}

// 	for k, c := range stmt.Clauses {
// 		newStmt.Clauses[k] = c
// 	}

// 	for k, p := range stmt.Preloads {
// 		newStmt.Preloads[k] = p
// 	}

// 	for k, j := range stmt.Joins {
// 		newStmt.Joins[k] = j
// 	}

// 	return newStmt
// }
