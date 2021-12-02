package models

type Empty struct{}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

//Error ...
type Error struct {
	Error interface{} `json:"error"`
}

//ValidationError ...
type ValidationError struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	UserMessage string `json:"user_message"`
}

type ResponseOK struct {
	Message interface{} `json:"message"`
}

type Response struct {
	ID string `json:"id"`
}

// Find query ...
type FindQueryModel struct {
	Page     int64  `json:"page,string"`
	Search   string `json:"search"`
	Active   bool   `json:"active"`
	Inactive bool   `josn:"inactive"`
	Limit    int64  `json:"limit,string"`
	Sort     string `json:"sort" example:"name|asc"`
	Lang     string `json:"lang"`
}

type GetRequest struct {
	Id   string `json:"id"`
	Slug string `json:"slug"`
	Lang string `json:"lang"`
}

type AuthorizationModel struct {
	Token string `header:"Authorization"`
}

type UserInfo struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

type Meta struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}
