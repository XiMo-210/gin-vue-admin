package admin

type RouterGroup struct {
	StudentInfoRouter
	WxUserRouter
	TaskRouter
	OrganizationRouter
	DepartmentRouter
	CommentScoreRouter
	InformationRouter
}
