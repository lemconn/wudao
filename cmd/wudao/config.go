package main

type config struct {
	Debug bool `envconfig:"WUDAO_DEBUG"`

	Server struct {
		Port    string `envconfig:"WUDAO_HTTP_BIND" default:":3000"`
		Proto   string `envconfig:"WUDAO_HTTP_PROTO"`
		Host    string `envconfig:"WUDAO_HTTP_HOST"`
		PidFile string `envconfig:"WUDAO_HTTP_PIDFILE" default:"/var/run/wudao.pid"`
	}

	Gateway struct {
		Kind   string `envconfig:"WUDAO_GATEWAY_KIND" default:"nginx"`
		Prefix string `envconfig:"WUDAO_GATEWAY_PREFIX"`
	}

	Etcd struct {
		Endpoints []string `envconfig:"WUDAO_ETCD_ENDPOINTS"`
		Username  string   `envconfig:"WUDAO_ETCD_USERNAME"`
		Password  string   `envconfig:"WUDAO_ETCD_PASSWORD"`
		Timeout   int      `envconfig:"WUDAO_ETCD_TIMEOUT"`
	}
}
