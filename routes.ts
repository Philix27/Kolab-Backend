class ApiRoutes { 
static AnnouncementGetAll_GET = "/api/v1/announcement/";
static AnnouncementGetOne_GET = "/api/v1/announcement/id";
static AuthGetCredentials_GET = "/api/v1/auth/get_auth_credentials";
static Chat_GET = "/api/v1/chat/";
static AuthViewMembers_GET = "/api/v1/organization/members";
static TaskEpicGetOne_GET = "/api/v1/task_epic/id";
static TaskEpicGet_GET = "/api/v1/task_epic/";
static TaskStoryGetOne_GET = "/api/v1/task_story/id";
static TaskStoryGetAll_GET = "/api/v1/task_story/";
static Task_GetOne_GET = "/api/v1/task/id";
static Task_GetAll_GET = "/api/v1/task/";
static UserGet_GET = "/api/v1/user/";
static NoteGet_GET = "/api/v1/notes/";
static AnnouncementGetAll_HEAD = "/api/v1/announcement/";
static AnnouncementGetOne_HEAD = "/api/v1/announcement/id";
static AuthGetCredentials_HEAD = "/api/v1/auth/get_auth_credentials";
static Chat_HEAD = "/api/v1/chat/";
static AuthViewMembers_HEAD = "/api/v1/organization/members";
static TaskEpicGetOne_HEAD = "/api/v1/task_epic/id";
static TaskEpicGet_HEAD = "/api/v1/task_epic/";
static TaskStoryGetOne_HEAD = "/api/v1/task_story/id";
static TaskStoryGetAll_HEAD = "/api/v1/task_story/";
static Task_GetOne_HEAD = "/api/v1/task/id";
static Task_GetAll_HEAD = "/api/v1/task/";
static UserGet_HEAD = "/api/v1/user/";
static NoteGet_HEAD = "/api/v1/notes/";
static AnnouncementCreate_POST = "/api/v1/announcement/";
static AuthCreateUser_POST = "/api/v1/auth/create_user";
static AuthSendEmailOtp_POST = "/api/v1/auth/send_email_otp";
static AuthVerifyEmailOtp_POST = "/api/v1/auth/verify_email_otp";
static AuthLogin_POST = "/api/v1/auth/login";
static AuthResetOtp_POST = "/api/v1/auth/resent_otp";
static AuthLogout_POST = "/api/v1/auth/logout";
static AuthCreateOrh_POST = "/api/v1/organization/";
static AuthInviteMember_POST = "/api/v1/organization/invite";
static TaskEpicCreate_POST = "/api/v1/task_epic/";
static TaskStoryCreate_POST = "/api/v1/task_story/";
static Task_Create_POST = "/api/v1/task/";
static TeamCreate_POST = "/api/v1/team/create_team";
static TeamGetAll_POST = "/api/v1/team/get_teams";
static TeamGetMembers_POST = "/api/v1/team/get_team_members";
static TeamAdd_POST = "/api/v1/team/add_member";
static TeamRemove_POST = "/api/v1/team/remove_member";
static TeamDelete_POST = "/api/v1/team/delete_team";
static AuthUpdateInfo_PUT = "/api/v1/organization/update_info";
static TaskEpicUpdate_PUT = "/api/v1/task_epic/";
static TaskStoryUpdate_PUT = "/api/v1/task_story/";
static Task_Update_PUT = "/api/v1/task/";
static AnnouncementDelete_DELETE = "/api/v1/announcement/";
static AuthDelete_DELETE = "/api/v1/organization/";
static TaskEpicDelete_DELETE = "/api/v1/task_epic/";
static TaskStoryDelete_DELETE = "/api/v1/task_story/";
static Task_Delete_DELETE = "/api/v1/task/";
static AnnouncementPatch_PATCH = "/api/v1/announcement/:id/expired";
}