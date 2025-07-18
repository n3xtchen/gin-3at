package entity

type Category struct {
	ID          int    // 业务唯一标识
	Name        string // 分类名称
	Description string // 分类描述
	// ParentID    int    // 父分类ID，0表示顶级分类
	// Parent      *Category // 父分类	Products []Product // 该分类下的商品
	Products []Product // 该分类下的商品
}
