package api

/**
 * 宝塔面板站点操作Api For Golang
 * @author Jidoer
 * @link https://www.bt.cn/api-doc.pdf
 */
var BtConfig Config
var SysConfig SystemConfig

type WebSite struct {
	WebName      string `json:"webname"`
	Path         string `json:"path"`
	Type_id      int    `json:"type_id"`
	Type         string `json:"type"`
	Version      string `json:"version"`
	Port         int    `json:"port"`
	Ps           string `json:"ps"`
	Ftp          string `json:"ftp"`
	Ftp_username string `json:"ftp_username"`
	Ftp_password string `json:"ftp_password"`
	Sql          string `json:"sql"`
	Coding       string `json:"coding"`
	DataUser     string `json:"data_user"`
	DataPassword string `json:"data_password"`
}

func NewWebSite() *WebSite {
	return &WebSite{}
}

func init() {
	BtConfig = *NewConfig()
	BtConfig.ServerAddress = "http://0.0.0.0:8080"
	SysConfig = *NewSystemConfig()
}

/**
 * 获取系统基础统计
 */
func GetSystemTotal() string {
	murl := BtConfig.ServerAddress + SysConfig.GetSystemTotal
	p_data := GetKeyData()
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取磁盘分区信息
 */
func GetDiskInfo() string {
	murl := BtConfig.ServerAddress + SysConfig.GetDiskInfo
	p_data := GetKeyData()
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取实时状态信息
 * (CPU、内存、网络、负载)
 */
func GetNetWork() string {
	murl := BtConfig.ServerAddress + SysConfig.GetNetWork
	p_data := GetKeyData()
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 检查是否有安装任务
 */
func GetTaskCount() string {
	murl := BtConfig.ServerAddress + SysConfig.GetTaskCount
	p_data := GetKeyData()
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 检查面板更新
 */
func UpdatePanel(check, force bool) string {
	murl := BtConfig.ServerAddress + SysConfig.UpdatePanel
	p_data := GetKeyData()
	p_data["check"] = check
	p_data["force"] = force
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取网站列表
 * @param string page   当前分页
 * @param string limit  取出的数据行数
 * @param string type   分类标识 -1: 分部分类 0: 默认分类
 * @param string order  排序规则 使用 id 降序：id desc 使用名称升序：name desc
 * @param string tojs   分页 JS 回调,若不传则构造 URI 分页连接
 * @param string search 搜索内容
 */
func Websites(search string, page, limit, type_, order string, tojs bool) string {
	murl := BtConfig.ServerAddress + SysConfig.Websites
	p_data := GetKeyData()
	p_data["page"] = page
	p_data["limit"] = limit
	p_data["type"] = type_
	p_data["order"] = order
	p_data["search"] = search
	p_data["tojs"] = tojs
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取网站FTP列表
 * @param string page   当前分页
 * @param string limit  取出的数据行数
 * @param string type   分类标识 -1: 分部分类 0: 默认分类
 * @param string order  排序规则 使用 id 降序：id desc 使用名称升序：name desc
 * @param string tojs   分页 JS 回调,若不传则构造 URI 分页连接
 * @param string search 搜索内容
 */
func WebFtpList(search string, page, limit, type_, order string, tojs bool) string {
	murl := BtConfig.ServerAddress + SysConfig.WebFtpList
	p_data := GetKeyData()
	p_data["page"] = page
	p_data["limit"] = limit
	p_data["type"] = type_
	p_data["order"] = order
	p_data["search"] = search
	p_data["tojs"] = tojs
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取网站SQL列表
 * @param string page   当前分页
 * @param string limit  取出的数据行数
 * @param string type   分类标识 -1: 分部分类 0: 默认分类
 * @param string order  排序规则 使用 id 降序：id desc 使用名称升序：name desc
 * @param string tojs   分页 JS 回调,若不传则构造 URI 分页连接
 * @param string search 搜索内容
 */
func WebSqlList(search string, page, limit, type_, order string, tojs bool) string {
	murl := BtConfig.ServerAddress + SysConfig.WebSqlList
	p_data := GetKeyData()
	p_data["page"] = page
	p_data["limit"] = limit
	p_data["type"] = type_
	p_data["order"] = order
	p_data["search"] = search
	p_data["tojs"] = tojs
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取所有网站分类
 */
func Webtypes() string {
	murl := BtConfig.ServerAddress + SysConfig.Webtypes
	p_data := GetKeyData()
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取已安装的 PHP 版本列表
 */
func GetPHPVersion() string {
	murl := BtConfig.ServerAddress + SysConfig.GetPHPVersion
	p_data := GetKeyData()
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 修改指定网站的PHP版本
 * @param [type] site 网站名
 * @param [type] php  PHP版本
 */
func SetPHPVersion(site, php string) string {
	murl := BtConfig.ServerAddress + SysConfig.SetPHPVersion
	p_data := GetKeyData()
	p_data["siteName"] = site
	p_data["version"] = php
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取指定网站运行的PHP版本
 * @param [type] site 网站名
 */
func GetSitePHPVersion(site string) string {
	murl := BtConfig.ServerAddress + SysConfig.GetSitePHPVersion
	p_data := GetKeyData()
	p_data["siteName"] = site
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 新增网站
 * @param [type] webname      网站域名 json格式
 * @param [type] path         网站路径
 * @param [type] type_id      网站分类ID
 * @param string type         网站类型
 * @param [type] version      PHP版本
 * @param [type] port         网站端口
 * @param [type] ps           网站备注
 * @param [type] ftp          网站是否开通FTP
 * @param [type] ftp_username FTP用户名
 * @param [type] ftp_password FTP密码
 * @param [type] sql          网站是否开通数据库
 * @param [type] codeing      数据库编码类型 utf8|utf8mb4|gbk|big5
 * @param [type] datauser     数据库账号
 * @param [type] datapassword 数据库密码
 */

func AddSite(webinfo WebSite) string {
	murl := BtConfig.ServerAddress + SysConfig.WebAddSite
	p_data := GetKeyData()
	p_data["webname"] = webinfo.WebName
	p_data["path"] = webinfo.Path
	p_data["type_id"] = webinfo.Type_id
	p_data["type"] = webinfo.Type
	p_data["version"] = webinfo.Version
	p_data["port"] = webinfo.Port
	p_data["ps"] = webinfo.Ps
	p_data["ftp"] = webinfo.Ftp
	p_data["ftp_username"] = webinfo.Ftp_username
	p_data["ftp_password"] = webinfo.Ftp_password
	p_data["sql"] = webinfo.Sql
	p_data["codeing"] = webinfo.Coding
	p_data["datauser"] = webinfo.DataUser
	p_data["datapassword"] = webinfo.DataPassword
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 删除网站
 * @param [type] id       网站ID
 * @param [type] webname  网站名称
 * @param [type] ftp      是否删除关联FTP
 * @param [type] database 是否删除关联数据库
 * @param [type] path     是否删除关联网站根目录
 *
 */
func WebDeleteSite(id, webname, ftp, database, path string) string {
	murl := BtConfig.ServerAddress + SysConfig.WebDeleteSite
	p_data := GetKeyData()
	p_data["id"] = id
	p_data["webname"] = webname
	p_data["ftp"] = ftp
	p_data["database"] = database
	p_data["path"] = path
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 停用站点
 * @param [type] id   网站ID
 * @param [type] name 网站域名
 */
func WebSiteStop(id, name string) string {
	murl := BtConfig.ServerAddress + SysConfig.WebSiteStop
	p_data := GetKeyData()
	p_data["id"] = id
	p_data["name"] = name
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 启用网站
 * @param [type] id   网站ID
 * @param [type] name 网站域名
 */

func WebSiteStart(id, name string) string {
	murl := BtConfig.ServerAddress + SysConfig.WebSiteStart
	p_data := GetKeyData()
	p_data["id"] = id
	p_data["name"] = name
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 设置网站到期时间
 * @param [type] id    网站ID
 * @param [type] edate 网站到期时间 格式：2019-01-01，永久：0000-00-00
 */
func WebSetEdate(id, edate string) string {
	murl := BtConfig.ServerAddress + SysConfig.WebSetEdate
	p_data := GetKeyData()
	p_data["id"] = id
	p_data["edate"] = edate
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 修改网站备注
 * @param [type] id 网站ID
 * @param [type] ps 网站备注
 */

func WebSetPs(id, ps string) string {
	murl := BtConfig.ServerAddress + SysConfig.WebSetPs
	p_data := GetKeyData()
	p_data["id"] = id
	p_data["ps"] = ps
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取网站备份列表
 * @param [type] id    网站ID
 * @param string page  当前分页
 * @param string limit 每页取出的数据行数
 * @param string type_  备份类型 目前固定为0
 * @param string tojs  分页js回调若不传则构造 URI 分页连接 get_site_backup
 */
func WebBackupList(id string, page string, limit string, type_ string, tojs string) string {
	murl := BtConfig.ServerAddress + SysConfig.WebBackupList
	p_data := GetKeyData()
	p_data["p"] = page
	p_data["limit"] = limit
	p_data["type"] = type_
	p_data["tojs"] = tojs
	p_data["search"] = id
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 创建网站备份
 * @param [type] id 网站ID
 */
func WebBackupCreate(id string) string {
	murl := BtConfig.ServerAddress + SysConfig.WebToBackup
	p_data := GetKeyData()
	p_data["id"] = id
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 删除网站备份
 * @param [type] id 网站备份ID
 */
func WebBackupDelete(id string) string {
	murl := BtConfig.ServerAddress + SysConfig.WebDelBackup
	p_data := GetKeyData()
	p_data["id"] = id
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 删除数据库备份
 * @param [type] id 数据库备份ID
 */
func SQLDelBackup(id string) string {
	murl := BtConfig.ServerAddress + SysConfig.SQLDelBackup
	p_data := GetKeyData()
	p_data["id"] = id
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 备份数据库
 * @param [type] id 数据库列表ID
 */
func SQLToBackup(id string) string {
	murl := BtConfig.ServerAddress + SysConfig.SQLToBackup
	p_data := GetKeyData()
	p_data["id"] = id
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取网站域名列表
 * @param [type]  id   网站ID
 * @param boolean list 固定传true
 */
func WebDoaminList(id string, list bool) string {
	murl := BtConfig.ServerAddress + SysConfig.WebDoaminList
	p_data := GetKeyData()
	p_data["id"] = id
	p_data["list"] = list
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 添加域名
 * @param [type] id      网站ID
 * @param [type] webname 网站名称
 * @param [type] domain  要添加的域名:端口 80 端品不必构造端口,多个域名用换行符隔开
 */
func WebAddDomain(id, webname, domain string) string {
	murl := BtConfig.ServerAddress + SysConfig.WebAddDomain
	p_data := GetKeyData()
	p_data["id"] = id
	p_data["webname"] = webname
	p_data["domain"] = domain
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 删除网站域名
 * @param [type] id      网站ID
 * @param [type] webname 网站名
 * @param [type] domain  网站域名
 * @param [type] port    网站域名端口
 */
func WebDelDomain(id, webname, domain string, port int) string {
	murl := BtConfig.ServerAddress + SysConfig.WebDelDomain
	p_data := GetKeyData()
	p_data["id"] = id
	p_data["webname"] = webname
	p_data["domain"] = domain
	p_data["port"] = port
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取可选的预定义伪静态列表
 * @param [type] siteName 网站名
 */

func GetRewriteList(siteName string) string {
	murl := BtConfig.ServerAddress + SysConfig.GetRewriteList
	p_data := GetKeyData()
	p_data["siteName"] = siteName
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取预置伪静态规则内容（文件内容）
 * @param [type] path 规则名
 * @param [type] type 0->获取内置伪静态规则；1->获取当前站点伪静态规则
 */
func GetFileBody(path string, type_ int) string {
	murl := BtConfig.ServerAddress + SysConfig.GetFileBody
	p_data := GetKeyData()
	// := type?'vhost/rewrite':'rewrite/nginx';
	path_dir := "rewrite/nginx"
	if type_ == 1 {
		path_dir = "vhost/rewrite"
	}

	//获取当前站点伪静态规则
	///www/server/panel/vhost/rewrite/user_hvVBT_1.test.com.conf
	//获取内置伪静态规则
	///www/server/panel/rewrite/nginx/EmpireCMS.conf
	//保存伪静态规则到站点
	///www/server/panel/vhost/rewrite/user_hvVBT_1.test.com.conf
	///www/server/panel/rewrite/nginx/typecho.conf
	p_data["path"] = "/www/server/panel/" + path_dir + "/" + path + ".conf"
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 保存伪静态规则内容(保存文件内容)
 * @param [type] path     规则名
 * @param [type] data     规则内容
 * @param string encoding 规则编码强转utf-8
 * @param number type     0->系统默认路径；1->自定义全路径
 */
func SaveFileBody(path, data string, encoding string, type_ int) string {
	murl := BtConfig.ServerAddress + SysConfig.SaveFileBody
	path_dir := "/www/server/panel/vhost/rewrite/" + path + ".conf"
	if type_ != 0 {
		path_dir = path
	}
	p_data := GetKeyData()

	p_data["path"] = path_dir
	p_data["data"] = data
	p_data["encoding"] = encoding
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 设置密码访问网站
 * @param [type] id       网站ID
 * @param [type] username 用户名
 * @param [type] password 密码
 */
func SetHasPwd(id, username, password string) string {
	murl := BtConfig.ServerAddress + SysConfig.SetHasPwd
	p_data := GetKeyData()
	p_data["id"] = id
	p_data["username"] = username
	p_data["password"] = password
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 关闭密码访问网站
 * @param [type] id 网站ID
 */
func CloseHasPwd(id string) string {
	murl := BtConfig.ServerAddress + SysConfig.CloseHasPwd
	p_data := GetKeyData()
	p_data["id"] = id
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取网站日志
 * @param [type] site 网站名
 */
func GetSiteLogs(site string) string {
	murl := BtConfig.ServerAddress + SysConfig.GetSiteLogs
	p_data := GetKeyData()
	p_data["siteName"] = site
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取网站盗链状态及规则信息
 * @param [type] id   网站ID
 * @param [type] site 网站名
 */
func GetSecurity(id, site string) string {
	murl := BtConfig.ServerAddress + SysConfig.GetSecurity
	p_data := GetKeyData()
	p_data["id"] = id
	p_data["name"] = site
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 设置网站盗链状态及规则信息
 * @param [type] id      网站ID
 * @param [type] site    网站名
 * @param [type] fix     URL后缀
 * @param [type] domains 许可域名
 * @param [type] status  状态
 */
func SetSecurity(id, site, fix, domains, status string) string {
	murl := BtConfig.ServerAddress + SysConfig.SetSecurity
	p_data := GetKeyData()
	p_data["id"] = id
	p_data["name"] = site
	p_data["fix"] = fix
	p_data["domains"] = domains
	p_data["status"] = status
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取网站三项配置开关（防跨站、日志、密码访问）
 * @param [type] id   网站ID
 * @param [type] path 网站运行目录
 */
func GetDirUserINI(id, path string) string {
	murl := BtConfig.ServerAddress + SysConfig.GetDirUserINI
	p_data := GetKeyData()
	p_data["id"] = id
	p_data["path"] = path
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 开启强制HTTPS
 * @param [type] site 网站域名（纯域名）
 */
func HttpToHttps(site string) string {
	murl := BtConfig.ServerAddress + SysConfig.HttpToHttps
	p_data := GetKeyData()
	p_data["siteName"] = site
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 关闭强制HTTPS
 * @param [type] site 域名(纯域名)
 */
func CloseToHttps(site string) string {
	murl := BtConfig.ServerAddress + SysConfig.CloseToHttps
	p_data := GetKeyData()
	p_data["siteName"] = site
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 设置SSL域名证书
 * @param [type] type 类型
 * @param [type] site 网站名
 * @param [type] key  证书key
 * @param [type] csr  证书PEM
 */
func SetSSL(type_, site, key, csr string) string {
	murl := BtConfig.ServerAddress + SysConfig.SetSSL
	p_data := GetKeyData()
	p_data["type"] = type_
	p_data["siteName"] = site
	p_data["key"] = key
	p_data["csr"] = csr
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 关闭SSL
 * @param [type] updateOf 修改状态码
 * @param [type] site     域名(纯域名)
 */
func CloseSSLConf(updateOf, site string) string {
	murl := BtConfig.ServerAddress + SysConfig.CloseSSLConf
	p_data := GetKeyData()
	p_data["updateOf"] = updateOf
	p_data["siteName"] = site
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取SSL状态及证书信息
 * @param [type] site 域名（纯域名）
 */
func GetSSL(site string) string {
	murl := BtConfig.ServerAddress + SysConfig.GetSSL
	p_data := GetKeyData()
	p_data["siteName"] = site
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取网站默认文件
 * @param [type] id 网站ID
 */
func WebGetIndex(id string) string {
	murl := BtConfig.ServerAddress + SysConfig.WebGetIndex
	p_data := GetKeyData()
	p_data["id"] = id
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 设置网站默认文件
 * @param [type] id    网站ID
 * @param [type] index 内容
 */
func WebSetIndex(id, index string) string {
	murl := BtConfig.ServerAddress + SysConfig.WebSetIndex
	p_data := GetKeyData()
	p_data["id"] = id
	p_data["index"] = index
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取网站流量限制信息
 * @param [type] id [description]
 */
func GetLimit(id string) string {
	murl := BtConfig.ServerAddress + SysConfig.GetLimitNet
	p_data := GetKeyData()
	p_data["id"] = id
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 设置网站流量限制信息
 * @param [type] id         网站ID
 * @param [type] perserver  并发限制
 * @param [type] perip      单IP限制
 * @param [type] limit_rate 流量限制
 */
func SetLimitNet(id, perserver, perip, limit_rate string) string {
	murl := BtConfig.ServerAddress + SysConfig.SetLimitNet
	p_data := GetKeyData()
	p_data["id"] = id
	p_data["perserver"] = perserver
	p_data["perip"] = perip
	p_data["limit_rate"] = limit_rate
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 关闭网站流量限制
 * @param [type] id 网站ID
 */
func CloseLimitNet(id string) string {
	murl := BtConfig.ServerAddress + SysConfig.CloseLimitNet
	p_data := GetKeyData()
	p_data["id"] = id
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取网站301重定向信息
 * @param [type] site 网站名
 */
func Get301Status(site string) string {
	murl := BtConfig.ServerAddress + SysConfig.Get301Status
	p_data := GetKeyData()
	p_data["siteName"] = site
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 设置网站301重定向信息
 * @param [type] site      网站名
 * @param [type] toDomain  目标Url
 * @param [type] srcDomain 来自Url
 * @param [type] type      类型
 */
func Set301Status(site, toDomain, srcDomain, type_ string) string {
	murl := BtConfig.ServerAddress + SysConfig.Set301Status
	p_data := GetKeyData()
	p_data["siteName"] = site
	p_data["toDomain"] = toDomain
	p_data["srcDomain"] = srcDomain
	p_data["type"] = type_
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取网站反代信息及状态
 * @param [type] site [description]
 */
func GetProxyList(site string) string {
	murl := BtConfig.ServerAddress + SysConfig.GetProxyList
	p_data := GetKeyData()
	p_data["sitename"] = site
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 添加网站反代信息
 * @param [type] cache     是否缓存
 * @param [type] proxyname 代理名称
 * @param [type] cachetime 缓存时长 /小时
 * @param [type] proxydir  代理目录
 * @param [type] proxysite 反代URL
 * @param [type] todomain  目标域名
 * @param [type] advanced  高级功能：开启代理目录
 * @param [type] sitename  网站名
 * @param [type] subfilter 文本替换json格式[{"sub1":"百度","sub2":"白底"},{"sub1":"","sub2":""}]
 * @param [type] type      开启或关闭 0关;1开
 */
func CreateProxy(cache, proxyname, cachetime, proxydir, proxysite, todomain, advanced, sitename, subfilter, type_ string) string {
	murl := BtConfig.ServerAddress + SysConfig.CreateProxy
	p_data := GetKeyData()
	p_data["cache"] = cache
	p_data["proxyname"] = proxyname
	p_data["cachetime"] = cachetime
	p_data["proxydir"] = proxydir
	p_data["proxysite"] = proxysite
	p_data["todomain"] = todomain
	p_data["advanced"] = advanced
	p_data["sitename"] = sitename
	p_data["subfilter"] = subfilter
	p_data["type"] = type_
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 修改网站反代信息
 * @param [type] cache     是否缓存
 * @param [type] proxyname 代理名称
 * @param [type] cachetime 缓存时长 /小时
 * @param [type] proxydir  代理目录
 * @param [type] proxysite 反代URL
 * @param [type] todomain  目标域名
 * @param [type] advanced  高级功能：开启代理目录
 * @param [type] sitename  网站名
 * @param [type] subfilter 文本替换json格式[{"sub1":"百度","sub2":"白底"},{"sub1":"","sub2":""}]
 * @param [type] type      开启或关闭 0关;1开
 */
func ModifyProxy(cache, proxyname, cachetime, proxydir, proxysite, todomain, advanced, sitename, subfilter, type_ string) string {
	murl := BtConfig.ServerAddress + SysConfig.ModifyProxy
	p_data := GetKeyData()
	p_data["cache"] = cache
	p_data["proxyname"] = proxyname
	p_data["cachetime"] = cachetime
	p_data["proxydir"] = proxydir
	p_data["proxysite"] = proxysite
	p_data["todomain"] = todomain
	p_data["advanced"] = advanced
	p_data["sitename"] = sitename
	p_data["subfilter"] = subfilter
	p_data["type"] = type_
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取网站域名绑定二级目录信息
 * @param [type] id 网站ID
 */
func GetDirBinding(id string) string {
	murl := BtConfig.ServerAddress + SysConfig.GetDirBinding
	p_data := GetKeyData()
	p_data["id"] = id
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 设置网站域名绑定二级目录
 * @param [type] id      网站ID
 * @param [type] domain  域名
 * @param [type] dirName 目录
 */
func AddDirBinding(id, domain, dirName string) string {
	murl := BtConfig.ServerAddress + SysConfig.AddDirBinding
	p_data := GetKeyData()
	p_data["id"] = id
	p_data["domain"] = domain
	p_data["dirName"] = dirName
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 删除网站域名绑定二级目录
 * @param [type] dirid 子目录ID
 */
func DelDirBinding(dirid string) string {
	murl := BtConfig.ServerAddress + SysConfig.DelDirBinding
	p_data := GetKeyData()
	p_data["id"] = dirid
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 获取网站子目录绑定伪静态信息
 * @param [type] dirid 子目录绑定ID
 */

func GetDirRewrite(dirid string, type_ int) string {
	murl := BtConfig.ServerAddress + SysConfig.GetDirRewrite
	p_data := GetKeyData()
	p_data["id"] = dirid
	if type_ != 0 {
		p_data["add"] = 1
	}
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 修改FTP账号密码
 * @param [type] id           FTPID
 * @param [type] ftp_username 用户名
 * @param [type] new_password 密码
 */
func SetUserPassword(id, ftp_username, new_password string) string {
	murl := BtConfig.ServerAddress + SysConfig.SetUserPassword
	p_data := GetKeyData()
	p_data["id"] = id
	p_data["ftp_username"] = ftp_username
	p_data["new_password"] = new_password
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 修改SQL账号密码
 * @param [type] id   	   SQLID
 * @param [type] name 	   用户名
 * @param [type] password 密码
 */
func ResDatabasePass(id, name, password string) string {
	murl := BtConfig.ServerAddress + SysConfig.ResDatabasePass
	p_data := GetKeyData()
	p_data["id"] = id
	p_data["name"] = name
	p_data["password"] = password
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 启用/禁用FTP
 * @param [type] id       FTPID
 * @param [type] username 用户名
 * @param [type] status   状态 0->关闭;1->开启
 */
func SetStatus(id, username, status string) string {
	murl := BtConfig.ServerAddress + SysConfig.SetStatus
	p_data := GetKeyData()
	p_data["id"] = id
	p_data["username"] = username
	p_data["status"] = status
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 宝塔一键部署列表
 * @param  string search 搜索关键词
 * @return [type]         [description]
 */
func Deployment(search string) string {
	var murl string
	p_data := GetKeyData()
	if search != "" {
		murl = BtConfig.ServerAddress + SysConfig.deployment + "&search=" + search
	} else {
		murl = BtConfig.ServerAddress + SysConfig.deployment
	}
	p_data["search"] = search
	result := HttpPostCookie(murl, p_data)
	return result
}

/**
 * 宝塔一键部署执行
 * @param [type] dname       部署程序名
 * @param [type] site_name   部署到网站名
 * @param [type] php_version PHP版本
 */
func SetupPackage(dname, site_name, php_version string) string {
	murl := BtConfig.ServerAddress + SysConfig.SetupPackage
	p_data := GetKeyData()
	p_data["dname"] = dname
	p_data["site_name"] = site_name
	p_data["php_version"] = php_version
	result := HttpPostCookie(murl, p_data)
	return result
}
