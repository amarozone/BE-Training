package model



type Laptops struct {
	Brand string `bson:"Brand"`
	Model string `bson:"Model"`
	Year  string `bson:"Year"`
	Price string `bson:"Price"`
}