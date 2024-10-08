package model

type Pagination struct {
	PageSize  int    `form:"pageSize" json:"pageSize"`
	Current   int    `form:"current" json:"current"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func (p *Pagination) GetOffset() int {
	return p.PageSize * (p.Current - 1)
}

// 角色查询参数
type ReqRoleQueryParam struct {
	TeamUuid string `json:"team_uuid"`
	Name     string `form:"name"`
	IsActive bool   `form:"is_active"`
	Pagination
}

type ReqUserLogin struct {
	// 用户名或邮箱
	Username string `json:"username" binding:"required"`
	// 密码
	Password string `json:"password" binding:"required"`
}

type ReqUserQueryParam struct {
	Email    string `json:"email"`    // 邮箱
	Phone    string `json:"phone"`    // 手机号
	Nickname string `json:"nickname"` // 昵称
	Sex      int    `json:"sex"`      // 性别
	Username string `json:"username"` // 用户名
	Status   int    `json:"status"`   // 状态
	Uuid     string `json:"uuid"`     // uuid
	Pagination
}

// 菜单查询参数
type ReqMenuQueryParam struct {
	Name string `json:"name"`
	Pagination
}

// 删除用户删除参数
type ReqUserDeleteParam struct {
	Uuid string `json:"uuid" binding:"required"`
}

// 删除菜单参数
type ReqMenuDeleteParam struct {
	Uuid string `json:"uuid" binding:"required"`
}

type ReqApiQueryParam struct {
	Path   string `json:"path"`
	Name   string `json:"name"`
	Method string `json:"method"`
	Status int    `json:"status"`
	Pagination
}

// 查询app的参数
type ReqAppQueryParam struct {
	Name   string `json:"name"`    // 名称
	ApiKey string `json:"api_key"` // api_key
	Status int    `json:"status"`  // 状态
	Pagination
}

// uuid参数
type ReqUuidParam struct {
	Uuid  string   `json:"uuid"`
	Uuids []string `json:"uuids"`
}

// api权限参数
type ReqApiPermissionParam struct {
	AppId string `json:"app_id"`
	Pagination
}

type ReqVerificationCodeParam struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type ReqRegisterParam struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

// 服务查询参数
type ReqServerQueryParam struct {
	Name string `json:"name"`
	Pagination
}

// 团队查询参数
type ReqTeamQueryParam struct {
	Name string `json:"name"`
	Pagination
}

// 创建团队成员参数
type ReqTeamMemberCreateParam struct {
	TeamUUID string `json:"team_uuid"`
	UserUUID string `json:"user_uuid"`
}

// 团队成员查询参数
type ReqTeamMemberQueryParam struct {
	TeamUUID string `json:"team_uuid"`
	Pagination
}

type ReqOpLogQueryParam struct {
	UserName string `form:"user_name"`
	Path     string `form:"path"`
	Method   string `form:"method"`
	Status   int    `form:"status"`
	Pagination
}

type ReqLoginLogQueryParam struct {
	Username string `json:"username"`
	Pagination
}

type ReqAPIQueryParam struct {
	Name   string `json:"name"`
	Module string `json:"module"`
	Status int    `json:"status"`
	Pagination
}

type ReqIdParam struct {
	Id int64 `json:"id"`
}

type ReqPermissionQueryParam struct {
	Name string `json:"name"`
	Pagination
}

type ReqUserPermissionQueryParam struct {
	UserUuid       string `json:"user_uuid"`
	PermissionUuid string `json:"permission_uuid"`
	Pagination
}

type ReqPermissionMenuQueryParam struct {
	PermissionUuid string `json:"permission_uuid"`
	MenuUuid       string `json:"menu_uuid"`
	Pagination
}

type ReqMenuAPIQueryParam struct {
	MenuUUID string `json:"menu_uuid"`
	APIUUID  string `json:"api_uuid"`
	Pagination
}

// 产品查询参数
type ReqProductQueryParam struct {
	Name string `json:"name"`
	Pagination
}

// 产品查询参数
type ReqProductSKUQueryParam struct {
	Name string `json:"name"`
	Sku  string `json:"sku"`
	Pagination
}

type ReqProductDeleteParam struct {
	UUids []string `json:"uuids"`
}
