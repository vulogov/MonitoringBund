package conf

import (
	"fmt"
	"os"
	"time"
	"github.com/google/uuid"
	"gopkg.in/alecthomas/kingpin.v2"
)

type filelist []string

var Argv [][]string

func (i *filelist) Set(value string) error {
	_, err := os.Stat(value)
	if os.IsNotExist(err) {
		return fmt.Errorf("Script file '%s' not found", value)
	} else {
		*i = append(*i, value)
		return nil
	}
}

func (i *filelist) String() string {
	return ""
}

func (i *filelist) IsCumulative() bool {
	return true
}

func FileList(s kingpin.Settings) (target *[]string) {
	target = new([]string)
	s.SetValue((*filelist)(target))
	return
}

var (
	seed    = time.Now().UTC().UnixNano()
	App     = kingpin.New("NRBUND", fmt.Sprintf("[ MBUND ] Language that is Functional and Stack-based: %v", BVersion))
	Name   	= App.Flag("name", "Define cluster name").Required().String()
	Id      = App.Flag("id", "Unique application ID").Default(uuid.New().String()).String()
	ApplicationId = App.Flag("applicationid", "Unique application ID").String()
	NoSync  = App.Flag("nosync", "Do not SYNC into cluster").Default("false").Bool()
	Debug   = App.Flag("debug", "Enable debug mode.").Default("false").Bool()
	CDebug  = App.Flag("core-debug", "Enable core debug mode.").Default("false").Bool()
	Color   = App.Flag("color", "--color : Enable colors on terminal --no-color : Disable colors .").Default("true").Bool()
	VBanner = App.Flag("banner", "Display [ MBUND ] banner .").Default("false").Bool()
	Timeout = App.Flag("timeout", "Timeout for common NRBUND operations").Default("5s").Duration()
	Etcd				= App.Flag("etcd", "ETCD endpoint location").Default("127.0.0.1:2379").Strings()
	Gnats   		= App.Flag("gnats", "GNATS endpoint location").Default("0.0.0.0:4222").String()
	GnatsC  		= App.Flag("nats-cluster", "GNATS cluster addresses").Strings()
	ShowResult 	= App.Flag("displayresult", "Display result of [ MBUND ] expression evaluation").Default("false").Bool()
	JPool   = App.Flag("jobpool", "Pool size for job runner").Default("500").Int()
	JCon   	= App.Flag("jobconcurrency", "Concurrency for job runner").Default("1").Int()
	Catcher = App.Flag("catcher", "Start metric catching to the local metric storage").Default("false").Bool()

	Retention		= App.Flag("retention", "Stored metrics retrention").Default("8760h").Duration()

	NR_account = App.Flag("newrelic_account", "New Relic Account").Envar("NEWRELIC_ACCOUNT").String()
	NR_api_key = App.Flag("newrelic_api_key", "New Relic API key").Envar("NEWRELIC_API_KEY").String()
	NR_lic_key = App.Flag("newrelic_lic_key", "New Relic License key").Envar("NEWRELIC_LICENSE_KEY").String()

	ZBX_account = App.Flag("zabbix_account", "Zabbix API login name").Envar("ZABBIX_ACCOUNT").String()
	ZBX_pass		= App.Flag("zabbix_password", "Zabbix API password").Envar("ZABBIX_PASSWORD").String()
	ZBX_api			= App.Flag("zabbix_api", "Zabbix API endpoint").Envar("ZABBIX_API").String()
	ZBX_host		= App.Flag("zabbix_host", "Zabbix Server/Proxy host").Envar("ZABBIX_HOST").String()
	ZBX_port		= App.Flag("zabbix_port", "Zabbix Server/Proxy port").Envar("ZABBIX_PORT").String()

	PR_url  		= App.Flag("prometheus_pusher_url", "URL of Prometheus pushgateway").Envar("PROMETHEUS_PUSHER_URL").String()
	PR_api_url  = App.Flag("prometheus_api_url", "URL of Prometheus API gateway").Envar("PROMETHEUS_API_URL").String()

	Args    = App.Flag("args", "String of arguments passed to a script").String()


	Version = App.Command("version", "Display information about [ MBUND ]")
	VTable  = Version.Flag("table", "Display [ MBUND ] inner information .").Default("true").Bool()

	Shell      	= App.Command("shell", "Run [ MBUND ] in interactive shell")
	ShowSResult = Shell.Flag("result", "Display result of expressions evaluated in [ MBUND ] shell").Default("false").Short('r').Bool()
	SExpr 			= Shell.Arg("expression", "[ MBUND ] expression passed to shell.").String()

	Run        	= App.Command("run", "Run MBUND in non-interactive mode")
	Scripts    	= Run.Arg("Scripts", "[ MBUND ] code to load").Strings()
	ShowRResult = Run.Flag("result", "Display result of scripts execution as it returned by [ MBUND ]").Default("false").Short('r').Bool()

	Eval 				= App.Command("eval", "Evaluate a [ MBUND ] expression")
	EStdin  		= Eval.Flag("stdin", "Read [ MBUND ] expression from STDIN .").Default("false").Bool()
	Expr 				= Eval.Arg("expression", "[ MBUND ] expression.").String()
	ShowEResult = Eval.Flag("result", "Display result of [ MBUND ] expression evaluation").Default("false").Short('r').Bool()

	Agitator   	= App.Command("agitator", "Run [ MBUND ] Agitator")
	UploadConf  = App.Flag("updateconf", "Update etcd configuration from local Agitator configuration").Default("false").Bool()
	AConf 			= Agitator.Flag("conf", "Configuration file for Agitator scheduler.").Strings()
	ABConf 			= Agitator.Flag("bund-conf", "BUND configuration for Agitator scheduler.").Strings()

	Agent   		= App.Command("agent", "Run [ MBUND ] Agent")

	Config   		= App.Command("config", "Upload configuration to ETCD")
	CNatsLocal 	= Config.Flag("nats-always-local", "Do not propagate NATS address").Default("false").Bool()
	CIdUpdated 	= Config.Flag("id-not-updated", "Do not propagate ID").Default("true").Bool()
	SConf       = Config.Flag("conf", "BUND script that will set the context to be uploaded to ETCD").Strings()
	CUpdate 		= Config.Flag("update", "Update basic application info").Default("false").Bool()
	CDelete 		= Config.Flag("delete", "Delete basic application info").Default("false").Bool()
	CShow 			= Config.Flag("show", "Display configuration stored in ETCD").Default("false").Bool()

	Submit   		= App.Command("submit", "Schedule NRBUND script to be executed")
	SArgs       = Submit.Flag("arg", "Pass positional argument to the script").Strings()
	SReturn 		= Submit.Flag("return", "HJSON instructions for sending data to New Relic").String()
	SScript 		= Submit.Arg("script", "BUND URL to the script, submitted to NRBUND for execution").Default("-").String()


	Sync   			= App.Command("sync", "Send NRBUND SYNC event")

	Take   			= App.Command("take", "Take a single scheduled NRBUND script and execute it")

	Watch   		= App.Command("watch", "Watch for NRBUND event on message bus and print them to Stdout")
	WTele				= Watch.Flag("telemetry", "Watch for telemetry").Default("false").Bool()

	Stop    		= App.Command("stop", "Send 'STOP' signal to a NRBUND bus")

	NRClient    = App.Command("newrelic_client", "Run MBUND native New Relic client")

	ZBXClient   = App.Command("zabbix_client", "Run MBUND native Zabbix client")

	PRClient   = App.Command("prometheus_client", "Run MBUND native Prometheus client")

	NRQLshell   = App.Command("nrql", "Run MBUND native NRQL shell")

	Telemetry   = App.Command("telemetry", "Submit telemetry to MBUND")
	TType 			= Telemetry.Flag("type", "Telemetry type").Default("metric").String()
	THost 			= Telemetry.Flag("host", "Host name for sent telemetry item").Required().String()
	TKey 				= Telemetry.Flag("key", 	"Telemetry key for sent telemetry item").Required().String()
	TDst 				= Telemetry.Flag("destination", "Destination for sent event item").String()
	TMType 			= Telemetry.Flag("metric-type", "Metric type").Default("gauge").String()
	TLSrv 			= Telemetry.Flag("log-service", "Log service").Default("service").String()
	TLLt  			= Telemetry.Flag("log-type", "Log type").Default("logfile").String()
	TValue			= Telemetry.Flag("value", "Telemetry value").Required().String()
	TArgs 			= Telemetry.Arg("attributes", "Telemetry attributes").StringMap()




)
