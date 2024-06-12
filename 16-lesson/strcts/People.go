package strcts

type People struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Nation     string `json:"nation"`
	Field      string `json:"field"`
	ParentName string `json:"parent_name"`
	City       string `json:"city"`
}
