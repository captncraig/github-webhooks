package webhooks

import (
	"time"
)

type PayloadBase struct {
	Sender     *User       `json:"sender"`
	Repository *Repository `json:"repository"`
}

type User struct {
	Login             string `json:"login"`
	Id                int64  `json:"id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Repository struct {
	Id               int64     `json:"id"`
	Name             string    `json:"name"`
	FullName         string    `json:"full_name"`
	Owner            *User     `json:"owner"`
	Private          bool      `json:"private"`
	HtmlUrl          string    `json:"html_url"`
	Descriprion      string    `json:"description"`
	Fork             bool      `json:"fork"`
	Url              string    `json:"url"`
	ForksUrl         string    `json:"forks_url"`
	KeysUrl          string    `json:"keys_url"`
	CollaboratorsUrl string    `json:"collaborators_url"`
	TeamsUrl         string    `json:"teams_url"`
	HooksUrl         string    `json:"hooks_url"`
	IssueEventsUrl   string    `json:"issue_events_url"`
	EventsUrl        string    `json:"events_url"`
	AssigneesUrl     string    `json:"assignees_url"`
	BranchesUrl      string    `json:"branches_url"`
	TagsUrl          string    `json:"tags_url"`
	BlobsUrl         string    `json:"blobs_url"`
	GitTagsUrl       string    `json:"git_tags_url"`
	GitRefsUrl       string    `json:"git_refs_url"`
	TreesUrl         string    `json:"trees_url"`
	StatusesUrl      string    `json:"statuses_url"`
	LanguagesUrl     string    `json:"languages_url"`
	StargazersUrl    string    `json:"stargazers_url"`
	ContributorsUrl  string    `json:"contributors_url"`
	SubscribersUrl   string    `json:"subscribers_url"`
	SubscriptionUrl  string    `json:"subscription_url"`
	CommitsUrl       string    `json:"commits_url"`
	GitCommitsUrl    string    `json:"git_commits_url"`
	CommentsUrl      string    `json:"comments_url"`
	IssueCommentUrl  string    `json:"issue_comment_url"`
	ContentsUrl      string    `json:"contents_url"`
	CompareUrl       string    `json:"compare_url"`
	MergesUrl        string    `json:"merges_url"`
	ArchiveUrl       string    `json:"archive_url"`
	DownloadsUrl     string    `json:"downloads_url"`
	IssuesUrl        string    `json:"issues_url"`
	PullsUrl         string    `json:"pulls_url"`
	MilestonesUrl    string    `json:"milestones_url"`
	NotificationsUrl string    `json:"notifications_url"`
	LabelsUrl        string    `json:"labels_url"`
	ReleasesUrl      string    `json:"releases_url"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	PushedAt         time.Time `json:"pushed_at"`
	GitUrl           string    `json:"git_url"`
	SshUrl           string    `json:"ssh_url"`
	CloneUrl         string    `json:"clone_url"`
	SvnUrl           string    `json:"svn_url"`
	Homepage         string    `json:"homepage"`
	Size             int       `json:"size"`
	StargazersCount  int       `json:"stargazers_count"`
	WatchersCount    int       `json:"watchers_count"`
	Language         string    `json:"language"`
	HasIssues        bool      `json:"has_issues"`
	HasDownloads     bool      `json:"has_downloads"`
	HasWiki          bool      `json:"has_wiki"`
	HasPages         bool      `json:"has_pages"`
	ForksCount       int       `json:"forks_count"`
	MirrorUrl        string    `json:"mirror_url"`
	OpenIssuesCount  int       `json:"open_issues_count"`
	Forks            int       `json:"forks"`
	OpenIssues       int       `json:"open_issues"`
	DefaultBranch    string    `json:"default_branch"`
}

type Comment struct {
	HtmlUrl   string    `json:"html_url"`
	Url       string    `json:"url"`
	Id        int64     `json:"id"`
	Body      string    `json:"body"`
	Path      string    `json:"path"`
	Position  int       `json:"position"`
	Line      int       `json:"line"`
	CommitId  string    `json:"commit_id"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Commit struct {
	Id        string     `json:"id"`
	Distinct  bool       `json:"distinct"`
	Message   string     `json:"message"`
	Timestamp time.Time  `json:"timestamp"`
	Url       string     `json:"url"`
	Author    *Committer `json:"author"`
	Committer *Committer `json:"committer"`
	Added     []string   `json:"added"`
	Removed   []string   `json:"removed"`
	Modified  []string   `json:"modified"`
}

type Hook struct {
	Id        int64     `json:"id"`
	Url       string    `json:"url"`
	TestUrl   string    `json:"test_url"`
	PingUrl   string    `json:"ping_url"`
	Name      string    `json:"name"`
	Events    []string  `json:"events"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Config    struct {
		Url         string `json:"url"`
		ContentType string `json:"content_type"`
	} `json:"config"`
}

type Committer struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type CommitCommentEvent struct {
	PayloadBase
	Comment *Comment `json:"comment"`
}

type CreateEvent struct {
	PayloadBase
	RefType      string `json:"ref_type"`
	Ref          string `json:"ref"`
	MasterBranch string `json:"master_branch"`
	Description  string `json:"description"`
}

type DeleteEvent struct {
	PayloadBase
	RefType string `json:"ref_type"`
	Ref     string `json:"ref"`
}

type DeploymentEvent struct {
	PayloadBase
	Deployment struct {
		Url           string      `json:"url"`
		Id            int64       `json:"id"`
		Sha           string      `json:"sha"`
		Ref           string      `json:"ref"`
		Task          string      `json:"task"`
		Payload       interface{} `json:"payload"`
		Environment   string      `json:"environment"`
		Description   string      `json:"description"`
		Creator       *Committer  `json:"creator"`
		CreatedAt     time.Time   `json:"created_at"`
		UpdatedAt     time.Time   `json:"updated_at"`
		StatusesUrl   string      `json:"statuses_url"`
		RepositoryUrl string      `json:"repository_url"`
	} `json:"deployment"`
}

type PingEvent struct {
	PayloadBase
	Zen    string `json:"zen"`
	HookId int64  `json:"hook_id"`
	Hook   *Hook  `json:"hook"`
}

type PushEvent struct {
	PayloadBase
	Ref        string     `json:"ref"`
	Before     string     `json:"before"`
	After      string     `json:"after"`
	Created    bool       `json:"created"`
	Deleted    bool       `json:"deleted"`
	Forced     bool       `json:"forced"`
	BaseRef    string     `json:"base_ref"`
	Compare    string     `json:"compare"`
	Commits    []*Commit  `json:"commits"`
	HeadCommit *Commit    `json:"head_commit"`
	Pusher     *Committer `json:"pusher"`
}
