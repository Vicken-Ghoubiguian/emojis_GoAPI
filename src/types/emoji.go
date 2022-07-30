package types

// Defining the type 'Emoji' which define an emoji...
type Emoji struct {
	number          int
	title           string
	unicode         string
	cldr_short_name string
	category Category
	subcategory SubCategory
}
