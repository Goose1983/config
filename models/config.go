package models

type Configuration struct {
	Server           Server
	App              Application
	Logger           Logger
	Site             Site
	DSN              DSN
	API              API
	Globals          Globals
	Secret           Secret
	Vmware           Vmware
	SSL              SSL
	Jenkins          Jenkins
	Ldap             Ldap
	EmailSendingInfo EmailSendingInfo
	Workers          Workers
}

type Workers struct {
	SyncStackActive               bool
	SyncUserActive                bool
	VMwareNightHandleActive       bool
	VMwareExpireTimeCheckerActive bool
}

type SSL struct {
	Cer string
	Key string
}

type EmailSendingInfo struct {
	Host        string
	UrmsAddress string
	ItAddress   string
}

type Server struct {
	Environment   string
	Host          string
	Port          string
	Schema        string
	HTTPTransport HTTPTransport
	Frontend      Frontend
	Version       string
}

type Application struct {
	PathToAvatars string
}
type HTTPTransport struct {
	DialerTimout          int
	DialerKeepAlive       int
	ExpectContinueTimeout int
	IdleConnTimeout       int
	MaxIdleConns          int
	MaxIdleConnsPerHost   int
}

type Frontend struct {
	Host   string
	Port   string
	Schema string
}

type Logger struct {
	AuditFile  string
	LogFile    string
	MaxSizeMb  int
	MaxBackups int
	MaxAgeDays int
	Compress   bool
	Prefix     string
}

type Site struct {
	Author      string
	Copyright   string
	Description string
	Title       string
}

type DSN struct {
	Host string
	Port string
	User string
	Pass string
	Base string
}

type API struct {
	Arty           string
	Asterisk       string
	Cloud          Cloud
	Docstore       string
	Exchange       string
	Gitlab         string
	Jenkins        string
	Jira           string
	Ldap           string
	Openshift      string
	OSHv3Prometeus string
	OSHv4          string
	OSHv4Oauth     string
	OSHv4Prometeus string
	Sonar          string
	Vault          string
	Vmware         string
	Wiki           string
}

type Cloud struct {
	Share  string
	Webdav string
}

type Globals struct {
	GroupAdmin             string
	GroupManagers          string
	GroupGitlabAuto        string
	GroupGitlabOwners      string
	GroupSonarAdmins       string
	GroupSonarProjectAdmin string
	GroupViewer            string
	GroupWiKiAdmins        string
	Roles                  string
}

type Ldap struct {
	Domain string
	Path   string
}

type Secret struct {
	Basic       string
	Cloud       string
	Domain      string
	Gitlab      string
	Login       string
	Password    string
	OSHv3Bearer string
	OSHv4Bearer string
}

type Vmware struct {
	ArchiveHostID          string
	Sessions               int
	Timer                  string
	Workers                int
	ArchivationWindowStart string
	ArchivationWindowEnd   string
	DevDomain              string
	DevLogin               string
	DevPassword            string
	DevAPI                 string
	Stand                  string
}

type Jenkins struct {
	AuthType string
	Token    string
}
