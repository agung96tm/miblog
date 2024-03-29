package dto

type UserInBlogPost struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type BlogPost struct {
	ID    uint            `json:"id"`
	Title string          `json:"title"`
	Body  string          `json:"body"`
	User  *UserInBlogPost `json:"user"`
}

type BlogPostQueryParams struct {
	SearchFilter
	OrderFilter
	PaginationParams
}

func (q BlogPostQueryParams) SearchFields() []string {
	return []string{"title", "body"}
}

func (q BlogPostQueryParams) OrderFields() []string {
	return []string{"id"}
}

func (q BlogPostQueryParams) DefaultPageSize() int {
	return 10
}

type BlogPostPagination struct {
	List       []*BlogPost `json:"list"`
	Pagination *Pagination `json:"pagination"`
}

type BlogPostCreateRequest struct {
	Title string `json:"title" validate:"required,min=10,max=200"`
	Body  string `json:"body" validate:"required,min=10"`
}

type BlogPostCreateResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type BlogPostUpdateRequest struct {
	Title string `json:"title" validate:"min=10,max=200"`
	Body  string `json:"body" validate:"min=10"`
}

type BlogPostUpdateResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
