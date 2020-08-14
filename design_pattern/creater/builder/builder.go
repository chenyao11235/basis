package builder

import (
	"github.com/pkg/errors"
)

/* 构建者模式
就是如果一个类构建比较复杂，比如需要对构建参数做校验，那么就把构建的过程交给另一个类
避免为目标定义过多的set方法
*/

//ResourcePoolConfig 资源池
type ResourcePoolConfig struct {
	name     string
	maxTotal int // 最大资源
	maxIdle  int // 最大空闲
	minIdle  int // 最小空闲
}

//Builder 构建者
type Builder interface {
	setName(string)
	setMaxTotal(int)
	setMaxIdle(int)
	setMinIdle(int)
	Build() (*ResourcePoolConfig, error)
}

//MySQLBuilder mysql构建
type MySQLBuilder struct {
	resourcePoolConfig *ResourcePoolConfig
}

//NewMySQLBuilder 工厂函数
func NewMySQLBuilder(c *ResourcePoolConfig) *MySQLBuilder {
	return &MySQLBuilder{
		resourcePoolConfig: c,
	}
}

func (b *MySQLBuilder) setName(name string) *MySQLBuilder {
	b.resourcePoolConfig.name = name
	return b
}

func (b *MySQLBuilder) setMaxTotal(n int) *MySQLBuilder {
	b.resourcePoolConfig.maxTotal = n
	return b
}

func (b *MySQLBuilder) setMaxIdle(n int) *MySQLBuilder {
	b.resourcePoolConfig.maxIdle = n
	return b
}

func (b *MySQLBuilder) setMinIdle(n int) *MySQLBuilder {
	b.resourcePoolConfig.minIdle = n
	return b
}

//Build 构建
func (b *MySQLBuilder) Build() (*ResourcePoolConfig, error) {
	if b.resourcePoolConfig.name == "" {
		return nil, errors.New("name can not be empty")
	}
	if b.resourcePoolConfig.maxIdle < 0 {
		return nil, errors.New("maxIdle can not less 0")
	}
	if b.resourcePoolConfig.minIdle < 0 {
		return nil, errors.New("maxIdle can not less 0")
	}
	if b.resourcePoolConfig.maxTotal < 0 {
		return nil, errors.New("maxIdle can not less 0")
	}
	if b.resourcePoolConfig.maxIdle < b.resourcePoolConfig.minIdle {
		return nil, errors.New("maxIdle can not less than minIdle")
	}
	return b.resourcePoolConfig, nil
}
