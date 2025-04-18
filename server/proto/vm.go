package proto

type IP struct {
	Name    string `json:"name"`
	Addr    string `json:"addr"`
	Version string `json:"version"`
	Type    string `json:"type"`
}

type GetInfoRsp struct {
	IPs         []IP   `json:"ips"`
	Mdns        string `json:"mdns"`
	Image       string `json:"image"`
	Application string `json:"application"`
	DeviceKey   string `json:"deviceKey"`
}

type GetHardwareRsp struct {
	Version string `json:"version"`
}

type SetGpioReq struct {
	Type     string `validate:"required"`  // reset / power
	Duration uint   `validate:"omitempty"` // press time (unit: milliseconds)
}

type GetGpioRsp struct {
	PWR bool `json:"pwr"` // power led
	HDD bool `json:"hdd"` // hdd led
}

type SetScreenReq struct {
	Type  string `validate:"required"` // resolution / fps / quality
	Value int    `validate:"number"`   // value
}

type GetScriptsRsp struct {
	Files []string `json:"files"`
}

type UploadScriptRsp struct {
	File string `json:"file"`
}

type RunScriptReq struct {
	Name string `validate:"required"`
	Type string `validate:"required"` // foreground | background
}

type RunScriptRsp struct {
	Log string `json:"log"`
}

type DeleteScriptReq struct {
	Name string `validate:"required"`
}

type GetVirtualDeviceRsp struct {
	Network bool `json:"network"`
	Disk    bool `json:"disk"`
}

type UpdateVirtualDeviceReq struct {
	Device string `validate:"required"`
}

type UpdateVirtualDeviceRsp struct {
	On bool `json:"on"`
}

type SetMemoryLimitReq struct {
	Enabled bool  `json:"enabled"`
	Limit   int64 `json:"limit"`
}

type GetMemoryLimitRsp struct {
	Enabled bool  `json:"enabled"`
	Limit   int64 `json:"limit"`
}

type SetOledReq struct {
	Sleep int `json:"sleep"`
}

type GetOLEDRsp struct {
	Exist bool `json:"exist"`
	Sleep int  `json:"sleep"`
}

type GetSSHStateRsp struct {
	Enabled bool `json:"enabled"`
}

type GetSwapStateRsp struct {
	Enabled bool `json:"enabled"`
}

type GetMouseJigglerStateRsp struct {
	Enabled bool `json:"enabled"`
}

type GetMdnsStateRsp struct {
	Enabled bool `json:"enabled"`
}

type SetHostnameReq struct {
	Hostname string `json:"hostname"`
}

type GetHostnameRsp struct {
	Hostname string `json:"hostname"`
}

type SetWebTitleReq struct {
	WebTitle string `json:"webTitle"`
}

type GetWebTitleRsp struct {
	WebTitle string `json:"webTitle"`
}
