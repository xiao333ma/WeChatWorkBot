package model

type GitLabHookModel struct {
	Commits []GitLabHookModelCommits
	Project GitLabHookModelRepository
	UserName string `json:"user_name"`
	EventName string `json:"event_name"`
	EventType string `json:"event_type"`
	Message string
	Ref string
	ObjectAttributes GitLabHookModelObjectAttributes `json:"object_attributes"`
	User GitLabHookModelAuthor
}

func (m *GitLabHookModel)Name () string  {
	if len(m.EventName) > 0 {
		return m.EventName
	} else {
		return m.EventType
	}
}

/*
git 提交信息
*/
type GitLabHookModelCommits struct {
	Message string
	Timestamp string
	Url string
	Author GitLabHookModelAuthor
}

/*
commit 作者相关
*/
type GitLabHookModelAuthor struct {
	Name string
}

/*
仓库信息
*/
type GitLabHookModelRepository struct {
	Name string
	GitHttpUrl string `json:"git_http_url"`
}

/*
merge_request 相关
*/
type GitLabHookModelObjectAttributes struct {
	Title string
	CreatedAt string `json:"created_at"`
	SourceBranch string `json:"source_branch"`
	TargetBranch string `json:"target_branch"`
	Url string
	Action string
}

var MergeAction = map[string]string{
	"open": "创建",
	"reopen": "重新打开",
	"approved": "批准",
	"merge": "合并",
	"close": "关闭",
}
