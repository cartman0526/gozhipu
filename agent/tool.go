package agent

func NewGetOSINFO() Tool {
	return Tool{
		Type: "function",
		Function: Function{
			Name:        "get_os_info",
			Description: "获取指定服务器的系统信息指标，如cpu、memory、disk",
			Parameters: Parameters{
				Type: "object",
				Properties: map[string]Parameter{
					"ostype": {
						Description: "系统信息指标的类型，如cpu、memory、disk",
						Type:        "string",
					},
				},
				Required: []string{"ostype"},
			},
		},
	}
}

func GetOSINFO(ostype string) map[string]interface{} {
	switch ostype {
	case "cpu":
		return map[string]interface{}{
			"cpu": "当前您的linux服务器的cpu负载是23.3%,cpu温度是85度,建议关机降温",
		}
	case "memory":
		return map[string]interface{}{
			"memory": "当前您的linux服务器的内存占用94%，内存剩余6%",
		}
	case "disk":
		return map[string]interface{}{
			"disk": "当前服务器的磁盘性能是正常的",
		}
	default:
		return map[string]interface{}{
			"error": "type参数错误",
		}
	}
}
