package os

import (
	"bytes"
	"context"
	"fmt"

	// "github.com/go-ping/ping"
	// "golang.org/x/net/icmp"
	"log"
	// "net"
	"os/exec"
	"runtime"

	// "strings"
	// "gostudy/vendor/golang.org/x/text/encoding/simplifiedchinese"
	"time"
)

/*
	初始化，设置日志格式
	Linux can use go-expect implements automatic operation.
	https://github.com/Netflix/go-expect
	https://github.com/google/goexpect
*/
func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println(runtime.GOOS)
	log.Println(runtime.GOARCH)
	initInstancesUsers()
}

var insts map[string]map[string]string

type Sqlplus struct {
	envs     map[string]string
	instname string
	tnsname  string
	username string
	password string
	scripts  []string
	cmd      *exec.Cmd
}

/*
	初始化实例，用户名，密码放入map
*/
func initInstancesUsers() {
	var instname string
	insts = make(map[string]map[string]string)

	// fata
	users := map[string]string{}
	instname = "fata"
	insts[instname] = users
	log.Printf("%s: %p", instname, insts[instname])
	users["HSFA"] = "Hsfa_1357zc"
	users["REPORT_QUERY"] = "report_1357zc"
	users["FACX"] = "H_tai123"
	users["HS_RUN"] = "Hsfb_1591zc"
	users["HS_TQ"] = "Hsfb_1591zc"
	users["HS_HIS"] = "Hsfb_1591zc"
	users["HSLIQPOWER"] = "hsliqpower"
	users["CC_EASTFAX"] = "cc_eastfax"
	users["HSFA_I9"] = "Hsfa_i9_1357zc"
	users["HTAM_COM"] = "HTAM_COM"
	users["DC_QRY"] = "Dc_qry123!"
	users["HEX_QRY"] = "Hex#qry123!"

	// trade
	users = map[string]string{}
	instname = "trade"
	insts[instname] = users
	log.Printf("%s: %p", instname, insts[instname])
	users["HSFA_QUERY"] = "hsfa_1357zc"
	users["CREDIT_QUERY"] = "credit_1357zc"
	users["HTAMCX"] = "H_tai123"
	users["TRADE_ZC"] = "Htrade_1357zc"
	users["DC_QRY"] = "Dc_qry123!"
	users["HEX_QRY"] = "Hex#qry123!"

	// 财汇
	users = map[string]string{}
	instname = "caihui"
	insts[instname] = users
	log.Printf("%s: %p", instname, insts[instname])
	users["FINCHINA"] = "finchina"
	users["DC_QRY"] = "Dc_qry123!"

	// 信评
	users = map[string]string{}
	instname = "credit"
	insts[instname] = users
	log.Printf("%s: %p", instname, insts[instname])
	users["MYCREDIT"] = "mycredit_1357zc"

	// 携宁
	users = map[string]string{}
	insts["sinitek"] = users
	log.Printf("sinitek: %p", insts["credit"])
	users["AIMS"] = "htzc_aims"
	users["CREDIT_QUERY"] = "credit_1357zc"
	users["PROJECT"] = "PROJECT"

	// 反洗钱
	users = map[string]string{}
	instname = "aml"
	insts[instname] = users
	log.Printf("%s: %p", instname, insts[instname])
	users["HSAML"] = "hsaml_1357zc"

	// TA4
	users = map[string]string{}
	instname = "ta4db"
	insts[instname] = users
	log.Printf("%s: %p", instname, insts[instname])
	users["HTTA4"] = "Hsta_1591zc"
	users["TA4_QUERY"] = "htta4321"
	users["DS_ZC1"] = "Hds_2580zc1"
	users["DS_ZC2"] = "Hds_2580zc1"
	users["DC_QRY"] = "Dc_qry123!"

	// 数据中心
	users = map[string]string{}
	instname = "dcdb"
	insts[instname] = users
	log.Printf("%s: %p", instname, insts[instname])
	users["DD_AODS"] = "DHdu_aatoadis123!"
	users["DD_SODS"] = "DHdu_astoadis123!"
	users["DD_DW"] = "DHdu_adtwa1i23!"
	users["DD_ETL"] = "DHdu#aettali123!"
	users["DD_PORTAL"] = "DHdu#aptoarital123!"
	users["DD_DM"] = "Wrd@er@3!Tr3"
	users["DD_ZW"] = "DHdu_aztwa1i23!"
	users["DD_FR_DC_VIEW_1"] = "dHdu_aftraidcview1123!"
	users["DD_FRREPORT"] = "Dd_frreport123!"
	users["DD_FK_01"] = "DHdu_aftka#i01123!"
	users["DD_LHTZ_01"] = "Dd_lhtz_01123!"
	users["DD_TYJX"] = "DHdu_attyajix123!"
	users["DD_SRC_ETL"] = "DHdu#astraci#etl123!"

	// 衡泰风控
	users = map[string]string{}
	instname = "xrisk"
	insts[instname] = users
	log.Printf("%s: %p", instname, insts[instname])
	users["XRISK_MD"] = "Huatai123"
	users["XRISK"] = "Huatai123"

	// 数据中心应用
	users = map[string]string{}
	instname = "dcappdb"
	insts[instname] = users
	log.Printf("%s: %p", instname, insts[instname])
	users["DD_FOF_APP"] = "DHdu#aftoafi#app123!"
	users["DD_DATA_SERVICE"] = "DHdu#adtaatia#srv123!"
	users["DD_FUNDNV"] = "DHdu#aftuanidnv123!"
	users["DD_CF"] = "DHdu#actfa1i23!"
	users["DD_UC"] = "DHdu#autca1i23!"

	// 客户服务系统
	users = map[string]string{}
	instname = "dcdbcss"
	insts[instname] = users
	log.Printf("%s: %p", instname, insts[instname])
	users["SRC"] = "SHruca1t2a3i_"
	users["KETTLE"] = "KHeutattlaei123_"
	users["CRM"] = "CHruma1t2a3i_"
	users["STAT"] = "SHtuaatt1a2i3_"
	users["PUB_SYS"] = "PHuubastyasi123_"
	users["LIVEBOS"] = "LHiuvaetbaois123_"
	users["CSS_QRY"] = "Cssqry123!"

}

type SqlplusConnEnv struct {
	port     int    // 4 byte or 8 byte
	username string // 8 byte
	password string // 8 byte
	url      string // 8 byte
	tnsname  string // 8 byte
	// sqltexts []string // 8 byte
}

/*
	设置数据库url
	xxx.xxx.xxx.xxx:port/srvname
	xxx.xxx.xxx.xxx:port:insname
*/
func (env *SqlplusConnEnv) SetUrl(url string) {
	// info := strings.Split(url, ":")
	// var ipaddr string
	// var port int
	// var name string
	// switch len(info) {
	// case 2:
	// 	ipaddr = info[0]
	// 	port = info[1]
	// case 3:

	// default:

	// }

}

func (env *SqlplusConnEnv) ParseUrl(ipaddr string, srvname string, insname string) {

}

func Ping(ipaddr string) {
	// 这里不需要conn ，因为这里都是以ip作为参数，不需要再通过conn获取地址
	// 如果使用host ，则需要使用 conn,err 作为变量申明，并通过 conn.RemoteAddr() 获取地址
	// _, err := net.DialTimeout("ipv4:icmp", ipaddr, 50*time.Millisecond)
	// if err != nil {
	// 	log.Fatalf("Ping [%s] error: %+v", ipaddr, err)
	// }
	// ping.New("").Resolve()
}

/*
	设置sqlplus输出格式，只输出结果
*/
func sqlplusEnvPureResult() string {
	var str bytes.Buffer
	// 如果sql出现了error，则退出
	str.WriteString("whenever sqlerror exit 9\n")
	str.WriteString("set heading off\n")
	str.WriteString("set feedback off\n")
	str.WriteString("set timing off\n")
	str.WriteString("set pagesize 0\n")
	str.WriteString("set linesize 32767\n")
	str.WriteString("set arraysize 1000\n")
	// log.Println(str.String())
	return str.String()
}

func cmdENV(sid string) []string {
	return []string{
		fmt.Sprintf("NLS_LANG=%s", "AMERICAN_AMERICA.ZHS16GBK"),
		fmt.Sprintf("ORACLE_SID=%s", sid),
	}

}

func RunSqlWithSqlplus(sqltext string, sqlplusEnv string, sid string, username string) {
	rtcode := 0
	// 设置程序超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "sqlplus", "-s", "/nolog")
	cmd.Env = append(cmd.Env, cmdENV("")...)
	var stdin, stdout, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	cmd.Stdin = &stdin
	stdin.WriteString(sqlplusEnv)
	stdin.WriteString(initConnectSql(username, "Oracle123!1"))
	stdin.WriteString(sqltext)
	cmd.Start()
	cmd.Wait()
	log.Println("   [ERROR]: ", stderr.String())
	log.Println("  [OUTPUT]: ", stdout.String())
	// 退出码
	rtcode = cmd.ProcessState.ExitCode()
	log.Println("[EXITCODE]: ", rtcode)
	if rtcode != 0 {
		log.Fatal("Sqlplus exit with errors!")
	}

}

/*
	校验命令是否存在
	@param command  命令字符串，例如"sqlplus"
	@return bool 	命令有效:true,命令无效 false
*/
func IsValidCommand(command string) bool {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("where", command)
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("which", command)
	}
	output, err := cmd.CombinedOutput()
	rtcode := cmd.ProcessState.ExitCode()
	if rtcode != 0 {
		log.Fatalf("Check command error! Detail: [Errors:%s \n Output: %s]", err, output)
		return false
	}
	log.Printf("Output: %s", output)
	return true
}

/*
	由于windows的cmd默认是GBK所以，golang要显示需要转换成utf8
*/
func transGbkToUtf8(str string) {
	// simplifiedchinese.GB18030.NewDecoder().Bytes()
}

/*
	设置是否使用 slient mode 执行
*/
func (sqlplus *Sqlplus) setMod(ctx context.Context, silent bool) {
	if silent {
		sqlplus.cmd = exec.CommandContext(ctx, "sqlplus", "-s", "/nolog")
	} else {
		sqlplus.cmd = exec.CommandContext(ctx, "sqlplus", "/nolog")
	}
}

func RunSqlWithSqlplusOnlyResult(sqltext string, sid string, username string) {
	RunSqlWithSqlplus(sqltext, sqlplusEnvPureResult(), sid, username)
}

func initConnectSql(username string, password string) string {
	// log.Println(fmt.Sprintf("connect %s/%s \n", username, password))
	return fmt.Sprintf("connect %s/%s \n", username, password)
}
