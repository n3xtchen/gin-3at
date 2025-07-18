package entity

type Product struct {
	ID          int      // 业务唯一标识
	Name        string   // 商品名称
	Description string   // 商品描述
	Price       float64  // 商品价格
	ImageURL    string   // 商品图片
	Stock       int      // 商品库存
	CategoryID  int      // 商品分类ID
	Category    Category // 商品分类
}
