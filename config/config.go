package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Addr 			    string		`yaml:"addr"`
	DSN				    string		`yaml:"dsn"`
	MaxIdleConn		    int			`yaml:"max_idle_conn"`
	TokenExpiredDays    int			`yaml:"token_expired_days"`
	TokenRefreshDays    int			`yaml:"token_refresh_days"`
	CookieUser		    string		`yaml:"cookie_user"`
	CookieToken         string		`yaml:"cookie_token"`
	BlogAnnouncementURL string		`yaml:"blog_announcement_url"`
	BlogLinkCount		int         `yaml:"blog_link_count"`
	BlogCommentCount	int      	`yaml:"blog_comment_count"`
	DebugMode			bool      	`yaml:"debug_mode"`
	MediaFolder         string      `yaml:"media_folder"`
	SendEmailReply      bool        `yaml:"send_email_reply"`
	EmailSMTPHost       string      `yaml:"email_smtp_host"`
	EmailSMTPPort       int         `yaml:"email_smtp_port"`
	EmailSMTPUsername   string      `yaml:"email_smtp_username"`
	EmailSMTPPassword   string      `yaml:"email_smtp_password"`
	EmailSMTPEmail      string      `yaml:"email_admin_email"`
}

const (
	SessionName    = "goldenfly"
	SessionUserKey = "g.id"
)

var config *Config

func Load(path string) error {
	result, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(result, &config)
}

func Get() *Config {
	return config
}
