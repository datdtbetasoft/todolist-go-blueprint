package constants

type ResponseCode struct {
	Code    string
	Message string
}

var (
	// Authentication & User Management Errors
	InvalidRequestParams       = ResponseCode{"ERR001", "Invalid request parameters"}
	EncryptPasswordFail        = ResponseCode{"ERR002", "Encrypt password fail"}
	RegisterUserFail           = ResponseCode{"ERR003", "Register user fail"}
	UserNotFound               = ResponseCode{"ERR004", "User not found"}
	PasswordIncorrect          = ResponseCode{"ERR005", "The password is incorrect"}
	LoginFail                  = ResponseCode{"ERR006", "Login fail, something wrong"}
	PhoneNumberExists          = ResponseCode{"ERR007", "This phone number is already registered"}
	EmailExists                = ResponseCode{"ERR008", "This email is already registered"}
	TokenInvalid               = ResponseCode{"ERR009", "Invalid token"}
	LogoutFail                 = ResponseCode{"ERR010", "Logout fail, something wrong"}
	TokenExpired               = ResponseCode{"ERR011", "Token expired"}
	EmailDoesNotExists         = ResponseCode{"ERR012", "Email address is not registered"}
	ForgotPasswordFail         = ResponseCode{"ERR013", "Forgot password fail, something wrong"}
	VerifyCodeExpired          = ResponseCode{"ERR014", "Verify code expired"}
	ChangePasswordFail         = ResponseCode{"ERR015", "Change password fail"}
	UpdateProfileFail          = ResponseCode{"ERR022", "Update profile fail"}
	SendMailForgotPasswordFail = ResponseCode{"ERR110", "Send mail verification fail"}
	InvalidEmailFail           = ResponseCode{"ERR110", "Invalid email address format"}
	InvalidPhoneNumberFail     = ResponseCode{"ERR111", "Invalid phone number format"}

	// File Management Errors
	UploadFileFail       = ResponseCode{"ERR016", "Upload file fail"}
	ObjectIdNotFound     = ResponseCode{"ERR017", "Object id not found"}
	ObjectKeyNotFound    = ResponseCode{"ERR018", "Object key not found"}
	DownloadObjectFail   = ResponseCode{"ERR019", "Download object fail"}
	MimeNotFound         = ResponseCode{"ERR020", "Mime not found"}
	UnsupportedMediaType = ResponseCode{"ERR042", "Unsupported media type"}

	// Chat & Room Management Errors
	CreateRoomFail       = ResponseCode{"ERR047", "Create room fail"}
	JoinInRoomFail       = ResponseCode{"ERR048", "Join in room fail"}
	GetRoomFail          = ResponseCode{"ERR049", "Get room fail"}
	SendMessageWSFail    = ResponseCode{"ERR050", "Send message to websocket fail"}
	CreateMessageFail    = ResponseCode{"ERR051", "Create message fail"}
	GetMessageFail       = ResponseCode{"ERR052", "Get message fail"}
	RoomNotFound         = ResponseCode{"ERR053", "Room not found"}
	ConnectWebsocketFail = ResponseCode{"ERR054", "Connect to websocket fail"}
	UpdateRoomFail       = ResponseCode{"ERR070", "Update room status fail"}
	CheckMemberFail      = ResponseCode{"ERR071", "Some thing wrong when check member in room"}
	WrongMember          = ResponseCode{"ERR072", "Member is not in this room"}
	DeleteRoomFail       = ResponseCode{"ERR088", "Delete room fail"}
	SetHiddenRoomFail    = ResponseCode{"ERR089", "Set hidden room fail"}
	DeleteMessageFail    = ResponseCode{"ERR093", "Delete message fail"}

	// Organization Management Errors
	CreateOrganizationFail       = ResponseCode{"ERR060", "Create organization failed"}
	GetUserOrganizationsFail     = ResponseCode{"ERR061", "Get user organizations failed"}
	AccessDeniedToOrganization   = ResponseCode{"ERR062", "Access denied to organization"}
	OrganizationNotFound         = ResponseCode{"ERR063", "Organization not found"}
	OnlyAdminCanUpdate           = ResponseCode{"ERR064", "Only admin can update organization"}
	UpdateOrganizationFail       = ResponseCode{"ERR065", "Update organization failed"}
	OnlyAdminOrSubAdminCanInvite = ResponseCode{"ERR066", "Only admin or sub-admin can invite members"}
	CreateInvitationFail         = ResponseCode{"ERR067", "Create invitation failed"}
	JoinOrganizationFail         = ResponseCode{"ERR068", "Join organization failed"}
	GetUserInfoFail              = ResponseCode{"ERR069", "Failed to get user info"}
	GetOrganizationMembersFail   = ResponseCode{"ERR070", "Get organization members failed"}
	OnlyAdminCanUpdateMemberRole = ResponseCode{"ERR072", "Only admin can update member roles"}
	UpdateMemberRoleFail         = ResponseCode{"ERR073", "Update member role failed"}
	OnlyAdminCanRemoveMembers    = ResponseCode{"ERR074", "Only admin can remove members"}
	CannotRemoveYourself         = ResponseCode{"ERR075", "Cannot remove yourself from organization"}
	RemoveMemberFail             = ResponseCode{"ERR076", "Remove member failed"}
	InviteCodeRequired           = ResponseCode{"ERR400", "Invite code is required"}
	InvalidRole                  = ResponseCode{"ERR401", "Invalid role"}
	NoFieldsToUpdate             = ResponseCode{"ERR402", "No fields to update"}

	// Project Management Errors
	CreateProjectFail           = ResponseCode{"ERR200", "Create project failed"}
	UpdateProjectFail           = ResponseCode{"ERR201", "Update project failed"}
	ProjectNotFound             = ResponseCode{"ERR202", "Project not found"}
	GetProjectsFail             = ResponseCode{"ERR203", "Get projects failed"}
	GetProjectFail              = ResponseCode{"ERR204", "Get project failed"}
	DeleteProjectFail           = ResponseCode{"ERR205", "Delete project failed"}
	AccessDenied                = ResponseCode{"ERR206", "Access denied"}
	AddMemberFail               = ResponseCode{"ERR207", "Add member failed"}
	RemoveProjectMemberFail     = ResponseCode{"ERR208", "Remove project member failed"}
	GetMembersFail              = ResponseCode{"ERR209", "Get members failed"}
	UpdateProjectMemberRoleFail = ResponseCode{"ERR210", "Update project member role failed"}

	// Task Management Errors
	CreateTaskFail           = ResponseCode{"ERR300", "Create task failed"}
	UpdateTaskFail           = ResponseCode{"ERR301", "Update task failed"}
	TaskNotFound             = ResponseCode{"ERR302", "Task not found"}
	GetTasksFail             = ResponseCode{"ERR303", "Get tasks failed"}
	GetTaskFail              = ResponseCode{"ERR304", "Get task failed"}
	DeleteTaskFail           = ResponseCode{"ERR305", "Delete task failed"}
	MoveTaskFail             = ResponseCode{"ERR306", "Move task failed"}
	CreateTaskListFail       = ResponseCode{"ERR307", "Create task list failed"}
	UpdateTaskListFail       = ResponseCode{"ERR308", "Update task list failed"}
	TaskListNotFound         = ResponseCode{"ERR309", "Task list not found"}
	GetTaskListsFail         = ResponseCode{"ERR310", "Get task lists failed"}
	DeleteTaskListFail       = ResponseCode{"ERR311", "Delete task list failed"}
	AddTaskCommentFail       = ResponseCode{"ERR312", "Add task comment failed"}
	GetTaskCommentsFail      = ResponseCode{"ERR313", "Get task comments failed"}
	UpdateTaskCommentFail    = ResponseCode{"ERR314", "Update task comment fail"}
	DeleteTaskCommentFail    = ResponseCode{"ERR315", "Delete task comment failed"}
	AddTaskSubscriberFail    = ResponseCode{"ERR316", "Add task subscriber failed"}
	RemoveTaskSubscriberFail = ResponseCode{"ERR317", "Remove task subscriber failed"}
	GetTaskSubscribersFail   = ResponseCode{"ERR318", "Get task subscribers failed"}
	UploadTaskAttachmentFail = ResponseCode{"ERR319", "Upload task attachment failed"}
	DeleteTaskAttachmentFail = ResponseCode{"ERR320", "Delete task attachment failed"}
	GetTaskAttachmentsFail   = ResponseCode{"ERR321", "Get task attachments failed"}
	SearchTasksFail          = ResponseCode{"ERR322", "Search tasks failed"}
	GetTaskHistoryFail       = ResponseCode{"ERR323", "Get task history failed"}
	DeleteMultipleTasksFail  = ResponseCode{"ERR324", "Delete multiple tasks failed"}
	CommentNotFound          = ResponseCode{"ERR325", "Comment not found"}
	CannotDeleteDefaultList  = ResponseCode{"ERR326", "Cannot delete default list"}
	DeleteCommentFail        = ResponseCode{"ERR327", "Delete comment failed"}

	// Notification Errors
	GetTokenDeviceFail   = ResponseCode{"ERR078", "Get Token Device fail"}
	SetTokenDeviceFail   = ResponseCode{"ERR079", "Set Token Device fail"}
	UnSubscribeTopicFail = ResponseCode{"ERR080", "UnSubscribe Topic FCM fail"}
	SubscribeTopicFail   = ResponseCode{"ERR081", "Subscribe Topic FCM fail"}
	SendMessageFCMFail   = ResponseCode{"ERR083", "Send message to FCM fail"}

	// General Errors
	InternalServerError = ResponseCode{"ERR500", "Internal server error"}
	ConvertJSONFail     = ResponseCode{"ERR082", "Convert Object to JSON fail"}
	EncryptRSAFail      = ResponseCode{"ERR099", "Encrypt data fail"}
	DecryptRSAFail      = ResponseCode{"ERR099", "Decrypt data fail"}
	GenerateKeyPairFail = ResponseCode{"ERR100", "Generate key pair fail"}
	GenerateKeyPemFail  = ResponseCode{"ERR101", "Generate key pem fail"}
	UserCountFail       = ResponseCode{"ERR102", "Count user fail"}
	ReadFileFail        = ResponseCode{"ERR109", "Read file fail"}
	DeleteUserFail      = ResponseCode{"ERR098", "Delete user fail"}
	DeleteS3ObjectFail  = ResponseCode{"ERR097", "Delete s3 object fail"}
)

var (
	// Authentication & User Management Success
	RegisterUserSuccess           = ResponseCode{"INF001", "Register user success"}
	ValidPhoneNumber              = ResponseCode{"INF002", "Valid phone number"}
	LoginSuccess                  = ResponseCode{"INF003", "Login success"}
	LogoutSuccess                 = ResponseCode{"INF004", "Logout success"}
	SendMailForgotPasswordSuccess = ResponseCode{"INF005", "Send mail forgot password success"}
	ChangePasswordSuccess         = ResponseCode{"INF006", "Change password success"}
	GetUserInfoSuccess            = ResponseCode{"INF008", "Get user info success"}
	SendEmailRegisterSuccess      = ResponseCode{"INF009", "Send email register success"}
	UpdateProfileSuccess          = ResponseCode{"INF010", "Update profile success"}
	DeleteAccountSuccess          = ResponseCode{"INF055", "Delete user success"}
	ValidEmail                    = ResponseCode{"INF049", "Valid email"}
	UpdateKeySuccess              = ResponseCode{"INF060", "update key success"}
	EmailVerifiedSuccess          = ResponseCode{"INF061", "Email verified successfully"}

	// File Management Success
	UploadFileSuccess = ResponseCode{"INF007", "Upload file success"}

	// Chat & Room Management Success
	CreateRoomSuccess       = ResponseCode{"INF027", "Create room success"}
	CreateMessageSuccess    = ResponseCode{"INF028", "Create message success"}
	GetMessageSuccess       = ResponseCode{"INF029", "Get message success"}
	GetRoomSuccess          = ResponseCode{"INF030", "Get room success"}
	ConnectWebsocketSuccess = ResponseCode{"INF032", "Connect websocket success"}
	DeleteRoomSuccess       = ResponseCode{"INF047", "Delete room success"}

	// Organization Management Success
	CreateOrganizationSuccess     = ResponseCode{"INF062", "Create organization success"}
	GetUserOrganizationsSuccess   = ResponseCode{"INF063", "Get user organizations success"}
	GetOrganizationSuccess        = ResponseCode{"INF064", "Get organization success"}
	UpdateOrganizationSuccess     = ResponseCode{"INF065", "Update organization success"}
	CreateInvitationSuccess       = ResponseCode{"INF066", "Create invitation success"}
	JoinOrganizationSuccess       = ResponseCode{"INF067", "Join organization success"}
	GetOrganizationMembersSuccess = ResponseCode{"INF068", "Get organization members success"}
	UpdateMemberRoleSuccess       = ResponseCode{"INF069", "Update member role success"}
	RemoveMemberSuccess           = ResponseCode{"INF070", "Remove member success"}

	// Project Management Success
	CreateProjectSuccess           = ResponseCode{"INF200", "Create project success"}
	UpdateProjectSuccess           = ResponseCode{"INF201", "Update project success"}
	GetProjectSuccess              = ResponseCode{"INF202", "Get project success"}
	GetProjectsSuccess             = ResponseCode{"INF203", "Get projects success"}
	DeleteProjectSuccess           = ResponseCode{"INF204", "Delete project success"}
	AddMemberSuccess               = ResponseCode{"INF205", "Add member success"}
	RemoveProjectMemberSuccess     = ResponseCode{"INF206", "Remove project member success"}
	GetMembersSuccess              = ResponseCode{"INF207", "Get members success"}
	UpdateProjectMemberRoleSuccess = ResponseCode{"INF208", "Update member role success"}

	// Task Management Success
	CreateTaskSuccess           = ResponseCode{"INF300", "Create task success"}
	UpdateTaskSuccess           = ResponseCode{"INF301", "Update task success"}
	GetTaskSuccess              = ResponseCode{"INF302", "Get task success"}
	GetTasksSuccess             = ResponseCode{"INF303", "Get tasks success"}
	DeleteTaskSuccess           = ResponseCode{"INF304", "Delete task success"}
	MoveTaskSuccess             = ResponseCode{"INF305", "Move task success"}
	CreateTaskListSuccess       = ResponseCode{"INF306", "Create task list success"}
	UpdateTaskListSuccess       = ResponseCode{"INF307", "Update task list success"}
	GetTaskListSuccess          = ResponseCode{"INF308", "Get task list success"}
	GetTaskListsSuccess         = ResponseCode{"INF309", "Get task lists success"}
	DeleteTaskListSuccess       = ResponseCode{"INF310", "Delete task list success"}
	AddTaskCommentSuccess       = ResponseCode{"INF311", "Add task comment success"}
	GetTaskCommentsSuccess      = ResponseCode{"INF312", "Get task comments success"}
	UpdateTaskCommentSuccess    = ResponseCode{"INF313", "Update task comment success"}
	DeleteTaskCommentSuccess    = ResponseCode{"INF314", "Delete task comment success"}
	AddTaskSubscriberSuccess    = ResponseCode{"INF315", "Add task subscriber success"}
	RemoveTaskSubscriberSuccess = ResponseCode{"INF316", "Remove task subscriber success"}
	GetTaskSubscribersSuccess   = ResponseCode{"INF317", "Get task subscribers success"}
	UploadTaskAttachmentSuccess = ResponseCode{"INF318", "Upload task attachment success"}
	DeleteTaskAttachmentSuccess = ResponseCode{"INF319", "Delete task attachment success"}
	GetTaskAttachmentsSuccess   = ResponseCode{"INF320", "Get task attachments success"}
	SearchTasksSuccess          = ResponseCode{"INF321", "Search tasks success"}
	GetTaskHistorySuccess       = ResponseCode{"INF322", "Get task history success"}
	DeleteMultipleTasksSuccess  = ResponseCode{"INF323", "Delete multiple tasks success"}

	// Notification Success
	SendNotificationSuccess = ResponseCode{"INF031", "Send notification success"}
	RegisterDeviceSuccess   = ResponseCode{"INF052", "Device registration success"}
)
