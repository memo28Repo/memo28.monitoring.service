package monitoringErrors

import "gorm.io/gorm"

// FrontendErrorReport 表示前端错误报告的 Gorm 模型。
type FrontendErrorReport struct {
	gorm.Model                // 添加默认的 gorm.Model，包含 ID、CreatedAt、UpdatedAt、DeletedAt 字段
	UserId            string  `gorm:"column:user_id;" json:"userId"` // 用户id 用于筛选错误id时发挥作用
	BrowserLanguage   string  `gorm:"column:browser_language;" json:"browserLanguage"`
	Type              string  `gorm:"column:type;" json:"type"` // window.error | unhandledrejection
	IsMobile          bool    `gorm:"column:is_mobile;" json:"isMobile"`
	IsTouchDevice     bool    `gorm:"column:is_touch_device;" json:"isTouchDevice"`
	DeviceOrientation string  `gorm:"column:device_orientation;" json:"deviceOrientation"`
	ScreenWidth       int     `gorm:"column:screen_width;" json:"screenWidth"`
	ScreenHeight      int     `gorm:"column:screen_height;" json:"screenHeight"`
	UserAgent         string  `gorm:"column:user_agent;" json:"userAgent"`
	Message           string  `gorm:"column:message;" json:"message"`
	Stack             string  `gorm:"column:stack;" json:"stack"`
	Colno             int     `gorm:"column:colno;" json:"colno"`
	Lineno            int     `gorm:"column:lineno;" json:"lineno"`
	Kind              string  `gorm:"column:kind;" json:"kind"`
	DownLink          float64 `gorm:"column:downLink;" json:"downLink"`
	Rtt               int64   `gorm:"column:rtt;" json:"rtt"`
	SaveData          bool    `gorm:"column:saveData;" json:"saveData"`
	EffectiveType     string  `gorm:"column:effectiveType;" json:"effectiveType"`

	// 性能监控字段

	Name                      string  `json:"name" gorm:"column:name"`
	EntryType                 string  `json:"entryType" gorm:"column:entry_type"`
	StartTime                 float64 `json:"startTime" gorm:"column:start_time"`
	Duration                  float64 `json:"duration" gorm:"column:duration"`
	InitiatorType             string  `json:"initiatorType" gorm:"column:initiator_type"`
	JsHeapSizeLimit           int64   `json:"jsHeapSizeLimit" gorm:"column:js_heap_size_limit"`
	TotalJSHeapSize           int64   `json:"totalJSHeapSize" gorm:"column:total_js_heap_size"`
	UsedJSHeapSize            int64   `json:"usedJSHeapSize" gorm:"column:used_js_heap_size"`
	DeliveryType              string  `json:"deliveryType" gorm:"column:delivery_type"`
	NextHopProtocol           string  `json:"nextHopProtocol" gorm:"column:next_hop_protocol"`
	RenderBlockingStatus      string  `json:"renderBlockingStatus" gorm:"column:render_blocking_status"`
	WorkerStart               int64   `json:"workerStart" gorm:"column:worker_start"`
	RedirectStart             int64   `json:"redirectStart" gorm:"column:redirect_start"`
	RedirectEnd               int64   `json:"redirectEnd" gorm:"column:redirect_end"`
	FetchStart                float64 `json:"fetchStart" gorm:"column:fetch_start"`
	DomainLookupStart         float64 `json:"domainLookupStart" gorm:"column:domain_lookup_start"`
	DomainLookupEnd           float64 `json:"domainLookupEnd" gorm:"column:domain_lookup_end"`
	ConnectStart              float64 `json:"connectStart" gorm:"column:connect_start"`
	SecureConnectionStart     float64 `json:"secureConnectionStart" gorm:"column:secure_connection_start"`
	ConnectEnd                float64 `json:"connectEnd" gorm:"column:connect_end"`
	RequestStart              float64 `json:"requestStart" gorm:"column:request_start"`
	ResponseStart             float64 `json:"responseStart" gorm:"column:response_start"`
	FirstInterimResponseStart float64 `json:"firstInterimResponseStart" gorm:"column:first_interim_response_start"`
	ResponseEnd               float64 `json:"responseEnd" gorm:"column:response_end"`
	TransferSize              float64 `json:"transferSize" gorm:"column:transfer_size"`
	EncodedBodySize           int64   `json:"encodedBodySize" gorm:"column:encoded_body_size"`
	DecodedBodySize           int64   `json:"decodedBodySize" gorm:"column:decoded_body_size"`
	ResponseStatus            int     `json:"responseStatus" gorm:"column:response_status"`
	ServerTiming              []byte  `json:"serverTiming" gorm:"column:server_timing"`
}

// Attribution 表示性能监控中的性能长任务的属性。
type Attribution struct {
	Name          string  `json:"name" gorm:"column:name"`
	EntryType     string  `json:"entryType" gorm:"column:entry_type"`
	StartTime     float64 `json:"startTime" gorm:"column:start_time"`
	Duration      float64 `json:"duration" gorm:"column:duration"`
	ContainerType string  `json:"containerType" gorm:"column:container_type"`
	ContainerSrc  string  `json:"containerSrc" gorm:"column:container_src"`
	ContainerId   string  `json:"containerId" gorm:"column:container_id"`
	ContainerName string  `json:"containerName" gorm:"column:container_name"`
}

// TableName 指定模型对应的数据库表名。
func (FrontendErrorReport) TableName() string {
	return "frontend_error_reports"
}
