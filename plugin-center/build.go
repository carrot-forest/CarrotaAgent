package plugincenter

type PluginCenter struct {
	meta struct {
		IP   string
		Port string
		API  struct {
			BasePath          string
			MessageReportPath string
		}
	}
}

func NewPluginCenter() *PluginCenter {
	pluginCenter := &PluginCenter{}
	return pluginCenter
}

func (pc *PluginCenter) WithAddress(ip string, port string) *PluginCenter {
	pc.meta.IP = ip
	pc.meta.Port = port
	return pc
}

func (pc *PluginCenter) WithMessageReport(messageReportPath string) *PluginCenter {
	pc.meta.API.MessageReportPath = messageReportPath
	return pc
}

func (pc *PluginCenter) WithBasePath(basePath string) *PluginCenter {
	pc.meta.API.BasePath = basePath
	return pc
}

func (pc *PluginCenter) Build() {

}
