package consts

var (
	DefaultFeiShuUrl = "https://open.feishu.cn"
)

//API Host const， v3
const (
	//获取 app_access_token（企业自建应用）
	ApiAppAccessTokenInternal = "/open-apis/auth/v3/app_access_token/internal/"
	//获取 app_access_token（应用商店应用）
	ApiAppAccessToken = "/open-apis/auth/v3/app_access_token/"
	//获取 tenant_access_token（应用商店应用）
	ApiTenantAccessToken = "/open-apis/auth/v3/tenant_access_token/"
	//获取 tenant_access_token（企业自建应用）
	ApiTenantAccessTokenInternal = "/open-apis/auth/v3/tenant_access_token/internal/"
	//重新推送 app_ticket
	ApiAppTicketResend = "/open-apis/auth/v3/app_ticket/resend/"
	//获取登录用户身份
	ApiOAuth2AccessToken = "/connect/qrconnect/oauth2/access_token/"
	//code2session
	ApiTokenLoginValidate = "/open-apis/mina/v2/tokenLoginValidate"
	//刷新access_token
	ApiRefreshAccessToken = "/open-apis/authen/v1/refresh_access_token"

	//////////////////部门和用户
	//获取通讯录授权范围
	ApiScope = "/open-apis/contact/v1/scope/get"
	//获取通讯录授权范围v2
	ApiScopeV2 = "/open-apis/contact/v2/scope/get"

	//获取部门列表
	ApiDepartmentSimpleList = "/open-apis/contact/v1/department/simple/list"
	//获取部门列表 v2
	ApiDepartmentSimpleListV2 = "/open-apis/contact/v2/department/simple/list"

	//获取部门详情
	ApiDepartmentInfoGet = "/open-apis/contact/v1/department/info/get"
	//批量获取部门详情
	ApiDepartmentInfoBatchGet = "/open-apis/contact/v2/department/detail/batch_get"

	//获取部门用户列表
	ApiDepartmentUserList = "/open-apis/contact/v1/department/user/list"
	//获取部门用户列表v2
	ApiDepartmentUserListV2 = "/open-apis/contact/v2/department/user/list"
	//获取用户详情
	ApiDepartmentUserDetailList = "/open-apis/contact/v1/department/user/detail/list"
	//获取用户详情v2
	ApiDepartmentUserDetailListV2 = "/open-apis/contact/v2/department/user/detail/list"

	//批量获取用户信息
	ApiUserBatchGet = "/open-apis/contact/v1/user/batch_get"
	//批量获取用户信息v2
	ApiUserBatchGetV2 = "/open-apis/contact/v2/user/batch_get"

	//////////////////机器人发送消息
	//机器人发送消息
	ApiRobotSendMessage = "/open-apis/message/v4/send/"
	//机器人批量发送消息
	ApiRobotSendBatchMessage = "/open-apis/message/v4/batch_send/"

	//////////////////角色
	//获取角色列表
	ApiRoleList = "/open-apis/contact/v2/role/list"
	//获取角色成员列表
	ApiRoleMemberList = "/open-apis/contact/v2/role/members"

	/////////////////////////////////////////////////////////
	//创建日历
	ApiCalendarCreate = "/open-apis/calendar/v3/calendars"
	//获取日历
	ApiCalendarGet = "/open-apis/calendar/v3/calendar_list/%s"
	//获取日历列表
	ApiCalendarListGet = "/open-apis/calendar/v3/calendar_list"
	//更新日历
	ApiCalendarUpdate = "/open-apis/calendar/v3/calendars/%s"
	//创建日程
	ApiCalendarEventCreate = "/open-apis/calendar/v3/calendars/%s/events"
	//删除日程
	ApiCalendarEventDelete = "/open-apis/calendar/v3/calendars/%s/events/%s"
	//邀请/移除日程参与者
	ApiCalendarEventAttendeesUpdate = "/open-apis/calendar/v3/calendars/%s/events/%s/attendees"
	//获取访问控制列表
	ApiCalendarAttendeesGet = "/open-apis/calendar/v3/calendars/%s/acl"
	//删除访问空值
	ApiCalendarAttendeesDelete = "/open-apis/calendar/v3/calendars/%s/acl/%s"

	//搜索用户F
	ApiSearchUser = "/open-apis/search/v1/user"

	//检验应用管理员
	ApiIsUserAdmin = "/open-apis/application/v3/is_user_admin"
	//查询应用管理员列表
	ApiAdminUserList = "/open-apis/user/v4/app_admin_user/list"

	////////用户群组
	//获取用户所在的群列表
	ApiUserGroupLIst = "/open-apis/user/v4/group_list"
	//获取群成员列表
	ApiChatMembers = "/open-apis/chat/v4/members"
	//搜索用户所在的群列表
	ApiChatSearch = "/open-apis/chat/v4/search"

	////////群信息和群管理
	//创建群
	ApiCreateChat = "/open-apis/chat/v4/create/"
	//获取群列表
	ApiChatList = "/open-apis/chat/v4/list"
	//获取群信息
	ApiChatInfo = "/open-apis/chat/v4/info"
	//更新群信息
	ApiUpdateChat = "/open-apis/chat/v4/update/"
	//拉用户进群
	ApiAddChatUser = "/open-apis/chat/v4/chatter/add/"
	//移除用户出群
	ApiRemoveChatUser = "/open-apis/chat/v4/chatter/delete/"
	//解散群
	ApiDisbandChat = "/open-apis/chat/v4/disband"
	//拉机器人进群
	ApiAddBot = "/open-apis/bot/v4/add"

	/////////订单
	//查询用户是否在应用开通范围
	ApiCheckUser    = "/open-apis/pay/v1/paid_scope/check_user"
	ApiGetOrderList = "/open-apis/pay/v1/order/list"
	ApiGetOrderInfo = "/open-apis/pay/v1/order/get"

	////////云文档
	//查询文档
	ApiSearchDocs = "/open-apis/suite/docs-api/search/object"
	//获取云文档信息
	ApiGetDocMeta = "/open-apis/doc/v2/meta"
	//获取图片
	ApiGetImage = "/open-apis/image/v4/get"
)

//Other Const
const (
	TestAppId     = "cli_9d5e49aae9ae9101"
	TestAppSecret = "HDzPYfWmf8rmhsF2hHSvmhTffojOYCdI"
	TestTicket    = "5c4d8a9c3676a3be7dfc81d1b529a77835bdaa3d"
)

const (
	AccessRoleReader         = "reader"
	AccessRoleFreeBusyReader = "free_busy_reader"
	AccessRoleWriter         = "writer"
	AccessRoleOwner          = "owner"
)

const (
	ActionInvite = "invite"
	ActionRemove = "remove"
)

func GetConst(c string) string {
	return DefaultFeiShuUrl + c
}
