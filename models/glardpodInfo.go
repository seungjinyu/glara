package models

type GlaraPodInfo struct {
	PodName        string
	PodLog         string
	OwnerReference string
}

type GlaraPodInfoList struct {
	InfoList []GlaraPodInfo
}
