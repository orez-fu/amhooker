package models

type Alert struct {
	Annotations  map[string]interface{} `json:"annotations"`
	EndsAt       string                 `json:"endsAt"`
	GeneratorURL string                 `json:"generatorURL"`
	Labels       map[string]interface{} `json:"labels"`
	StartsAt     string                 `json:"startsAt"`
	Status       string                 `json:"status"`
}

type AlertBody struct {
	Alerts            []Alert                `json:"alerts"`
	CommonLabels      map[string]interface{} `json:"commonLabels"`
	CommonAnnotations map[string]interface{} `json:"commonAnnotations"`
	ExternalURL       string                 `json:"externalURL"`
	GroupKey          string                 `json:"groupKey"`
	GroupLabels       map[string]interface{} `json:"groupLabels"`
	Receiver          string                 `json:"receiver"`
	Status            string                 `json:"status"`
	Version           string                 `json:"version"`
}

type AlertInTransform struct {
	Firing   []*Alert `json:"firing"`
	Resolved []*Alert `json:"resolved"`
}

type AlertBodyTransform struct {
	Alerts            AlertInTransform       `json:"alerts"`
	CommonLabels      map[string]interface{} `json:"commonLabels"`
	CommonAnnotations map[string]interface{} `json:"commonAnnotations"`
	ExternalURL       string                 `json:"externalURL"`
	GroupKey          string                 `json:"groupKey"`
	GroupLabels       map[string]interface{} `json:"groupLabels"`
	Receiver          string                 `json:"receiver"`
	Status            string                 `json:"status"`
	Version           string
}

/**
 * Subdivided by 1024
 */
const (
	Kb = iota
	Mb
	Gb
	Tb
	Pb
	Eb
	Zb
	Yb
	Information_Size_MAX
)

/**
 * Subdivided by 10000
 */
const (
	K = iota
	M
	G
	T
	P
	E
	Z
	Y
	Scale_Size_MAX
)
