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
	EvtDst	= App.Flag("event", "Destination for New Relic events").Default("BundApplicationEvent").String()
	Debug   = App.Flag("debug", "Enable debug mode.").Default("false").Bool()
	CDebug  = App.Flag("core-debug", "Enable core debug mode.").Default("false").Bool()
	Color   = App.Flag("color", "--color : Enable colors on terminal --no-color : Disable colors .").Default("true").Bool()
	VBanner = App.Flag("banner", "Display [ MBUND ] banner .").Default("false").Bool()
	Timeout = App.Flag("timeout", "Timeout for common NRBUND operations").Default("5s").Duration()
	Etcd				= App.Flag("etcd", "ETCD endpoint location").Default("127.0.0.1:2379").Strings()
	Gnats   		= App.Flag("gnats", "GNATS endpoint location").Default("0.0.0.0:4222").String()
	ShowResult 	= App.Flag("displayresult", "Display result of [ MBUND ] expression evaluation").Default("false").Bool()
	JPool   = App.Flag("jobpool", "Pool size for job runner").Default("500").Int()
	JCon   	= App.Flag("jobconcurrency", "Concurrency for job runner").Default("1").Int()
	Args    = App.Flag("args", "String of arguments passed to a script").String()


	Version = App.Command("version", "Display information about [ MBUND ]")
	VTable  = Version.Flag("table", "Display [ MBUND ] inner information .").Default("true").Bool()

	Shell      	= App.Command("shell", "Run [ MBUND ] in interactive shell")
	ShowSResult = Shell.Flag("result", "Display result of expressions evaluated in [ MBUND ] shell").Default("false").Short('r').Bool()
	SExpr 			= Shell.Arg("expression", "[ MBUND ] expression passed to shell.").String()

	Run        	= App.Command("run", "Run NRBUND in non-interactive mode")
	Scripts    	= Run.Arg("Scripts", "[ MBUND ] code to load").Strings()
	ShowRResult = Run.Flag("result", "Display result of scripts execution as it returned by [ MBUND ]").Default("false").Short('r').Bool()

	Eval 				= App.Command("eval", "Evaluate a [ MBUND ] expression")
	EStdin  		= Eval.Flag("stdin", "Read [ MBUND ] expression from STDIN .").Default("false").Bool()
	Expr 				= Eval.Arg("expression", "[ MBUND ] expression.").String()
	ShowEResult = Eval.Flag("result", "Display result of [ MBUND ] expression evaluation").Default("false").Short('r').Bool()

	Agitator   	= App.Command("agitator", "Run [ MBUND ] Agitator")
	UploadConf  = App.Flag("updateconf", "Update etcd configuration from local Agitator configuration").Default("false").Bool()
	AConf 			= Agitator.Flag("conf", "Configuration file for Agitator scheduler.").Required().Strings()

	Agent   		= App.Command("agent", "Run [ MBUND ] Agent")

	Config   		= App.Command("config", "Upload configuration to ETCD")
	SConf       = Config.Flag("conf", "BUND script that will set the context to be uploaded to ETCD").Strings()
	CShow 			= Config.Flag("show", "Display configuration stored in ETCD").Default("false").Bool()

	Submit   		= App.Command("submit", "Schedule NRBUND script to be executed")
	SArgs       = Submit.Flag("arg", "Pass positional argument to the script").Strings()
	SReturn 		= Submit.Flag("return", "HJSON instructions for sending data to New Relic").String()
	SScript 		= Submit.Arg("script", "BUND URL to the script, submitted to NRBUND for execution").Default("-").String()


	Sync   			= App.Command("sync", "Send NRBUND SYNC event")

	Take   			= App.Command("take", "Take a single scheduled NRBUND script and execute it")

	Watch   		= App.Command("watch", "Watch for NRBUND event on message bus and print them to Stdout")

	Stop    		= App.Command("stop", "Send 'STOP' signal to a NRBUND bus")

)
