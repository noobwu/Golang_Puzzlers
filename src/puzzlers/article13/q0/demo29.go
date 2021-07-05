package main

import "fmt"

// 示例1。
// AnimalCategory 代表动物分类学中的基本分类法。
type AnimalCategory struct {
	kingdom string // 界。
	phylum  string // 门。
	class   string // 纲。
	order   string // 目。
	family  string // 科。
	genus   string // 属。
	species string // 种。
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s",
		ac.kingdom, ac.phylum, ac.class, ac.order,
		ac.family, ac.genus, ac.species)
}

// 示例2。
type Animal struct {
	scientificName string // 学名。
	/*
		这个问题的典型回答是：字段声明AnimalCategory代表了Animal类型的一个嵌入字段。
		Go 语言规范规定，如果一个字段的声明中只有字段的类型名而没有字段的名称，那么它就是一个嵌入字段，也可以被称为匿名字段。
		我们可以通过此类型变量的名称后跟“.”，再后跟嵌入字段类型的方式引用到该字段。也就是说，嵌入字段的类型既是类型也是名称。
	*/
	AnimalCategory // 动物基本分类。
}

// 该方法会"屏蔽"掉嵌入字段中的同名方法。
func (a Animal) String() string {
	return fmt.Sprintf("%s (category: %s)",
		a.scientificName, a.AnimalCategory)
}

// 示例3。
/*
多层嵌入的问题(嵌入字段本身也有嵌入字段的情况)
*/
type Cat struct {
	name string
	Animal
}

/*
	最后的最后，如果处于同一个层级的多个嵌入字段拥有同名的字段或方法，
	那么从被嵌入类型的值那里，选择此名称的时候就会引发一个编译错误，因为编译器无法确定被选择的成员到底是哪一个。
*/
// 该方法会"屏蔽"掉嵌入字段中的同名方法。
func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.Animal.AnimalCategory, cat.name)
}

func main() {
	// 示例1。
	category := AnimalCategory{species: "cat"} //实例化一个对象
	fmt.Printf("The animal category: %s\n", category)

	// 示例2。
	animal := Animal{
		scientificName: "American Shorthair",
		AnimalCategory: category,
	}
	fmt.Printf("The animal: %s\n", animal)

	// 示例3。
	cat := Cat{
		name:   "little pig",
		Animal: animal,
	}
	fmt.Printf("The cat: %s\n", cat)
}
