package article

// Article は 仮のリソースです
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles は Articleの集まりです
type Articles []Article
