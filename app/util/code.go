package util

/* status code */
const (
	STATUS_COMPLITE int = iota
	STATUS_NOCHECK
	STATUS_IGNORE
	STATUS_DOING
	STATUS_FIXPLAN
)

func GetStatus(code int) string {
	switch code {
	case STATUS_COMPLITE:
		return "完了"
	case STATUS_NOCHECK:
		return "未確認"
	case STATUS_IGNORE:
		return "対応不要"
	case STATUS_DOING:
		return "確認中"
	case STATUS_FIXPLAN:
		return "対応予定"
	}
	return ""
}

/* priority code */
const (
	P_IMPORTANT int = iota
	P_EMERGENCY
	P_OTHER
)

func GetPriority(priority int) string {
	switch priority {
	case P_IMPORTANT:
		return "重要"
	case P_EMERGENCY:
		return "緊急"
	case P_OTHER:
		return "その他"
	}
	return ""
}
