package main

type Diagnosis struct {
	NumberOfBroadcasts             int `json:"number_of_broadcasts"`
	NumberOfUpdates                int `json:"number_of_updates"`
	NumberOfRejectedDueToThreshold int `json:"number_of_rejected_due_to_threshold"`
	NumberOfDuplicates             int `json:"number_of_duplicates"`
	NumberOfPkMatches              int `json:"number_of_pk_matches"`
	NumberOfInternalUpdates        int `json:"number_of_internal_updates"`
	NumberOfPacketsDropped         int `json:"number_of_packets_dropped"`
}

type CtNode struct {
	Diagnosis *Diagnosis `json:"diagnosis"`
}

type RunConfig struct {
	Run int `json:"run"`
	N   int `json:"n"`
	L   int `json:"l"`
	DT  int `json:"dt"`
	TTL int `json:"ttl"`
	T   int `json:"t"`
}