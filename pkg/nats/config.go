package nats

type Config struct {
	Url            string      `koanf:"nats.url"`
	Servers        []string    `koanf:"nats.servers"`
	UserJWT        string      `koanf:"nats.user_jwt"`
	Nkey           string      `koanf:"nats.nkey"`
	User           string      `koanf:"nats.user"`
	Password       string      `koanf:"nats.password"`
	Token          string      `koanf:"nats.token"`
	TokenHandler   interface{} `koanf:"nats.token_handler"`
	Timeout        string      `koanf:"nats.timeout"`
	DrainTimeout   string      `koanf:"nats.drain_timeout"`
	FlusherTimeout string      `koanf:"nats.flusher_timeout"`
	Pedantic       bool        `koanf:"nats.pedantic"`
	Secure         bool        `koanf:"nats.secure"`
	TlsConfig      struct {
		CertFile string `koanf:"nats.tls_config.cert_file"`
		KeyFile  string `koanf:"nats.tls_config.key_file"`
		CaFile   string `koanf:"nats.tls_config.ca_file"`
		Verify   bool   `koanf:"nats.tls_config.verify"`
		Timeout  string `koanf:"nats.tls_config.timeout"`
	} `koanf:"nats.tls_config"`
	TlsCertCB           interface{} `koanf:"nats.tls_cert_cb"`
	RootCAsCB           interface{} `koanf:"nats.root_cas_cb"`
	PingInterval        string      `koanf:"nats.ping_interval"`
	inProcessServer     bool        `koanf:"nats.in_process_server"`
	MaxReconnects       int         `koanf:"nats.max_reconnects"`
	ReconnectWait       string      `koanf:"nats.reconnect_wait"`
	MaxPingsOutstanding int         `koanf:"nats.max_pings_outstanding"`
	AllowReconnect      bool        `koanf:"nats.allow_reconnect"`
	Verbose             bool        `koanf:"nats.verbose"`
	NoRandomize         bool        `koanf:"nats.no_randomize"`
	NoEcho              bool        `koanf:"nats.no_echo"`
	Name                string      `koanf:"nats.name"`
	Compression         bool        `koanf:"nats.compression"`

	/*
		Dialer         interface{} `koanf:"nats.dialer"`
		CustomDialer   interface{} `koanf:"nats.custom_dialer"`
		CustomReconnectDelayCB      interface{} `koanf:"nats.custom_reconnect_delay_cb"`
		ReconnectJitter             string      `koanf:"nats.reconnect_jitter"`
		ReconnectJitterTLS          string      `koanf:"nats.reconnect_jitter_tls"`
		ClosedCB                    interface{} `koanf:"nats.closed_cb"`
		DisconnectedCB              interface{} `koanf:"nats.disconnected_cb"`
		DisconnectedErrCB           interface{} `koanf:"nats.disconnected_err_cb"`
		ConnectedCB                 interface{} `koanf:"nats.connected_cb"`
		ReconnectedCB               interface{} `koanf:"nats.reconnected_cb"`
		DiscoveredServersCB         interface{} `koanf:"nats.discovered_servers_cb"`
		AsyncErrorCB                interface{} `koanf:"nats.async_error_cb"`
		ReconnectBufSize            int         `koanf:"nats.reconnect_buf_size"`
		SubChanLen                  int         `koanf:"nats.sub_chan_len"`
		SignatureCB                 interface{} `koanf:"nats.signature_cb"`
		UseOldRequestStyle          bool        `koanf:"nats.use_old_request_style"`
		NoCallbacksAfterClientClose bool        `koanf:"nats.no_callbacks_after_client_close"`
		LameDuckModeHandler         interface{} `koanf:"nats.lame_duck_mode_handler"`
		RetryOnFailedConnect        bool        `koanf:"nats.retry_on_failed_connect"`
		ProxyPath                   string      `koanf:"nats.proxy_path"`
		InboxPrefix                 string      `koanf:"nats.inbox_prefix"`
		IgnoreAuthErrorAbort        bool        `koanf:"nats.ignore_auth_error_abort"`
		SkipHostLookup              bool        `koanf:"nats.skip_host_lookup"`
	*/
}
