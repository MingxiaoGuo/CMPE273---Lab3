package structure

type KeyValuePair struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

type AllPair struct {
	Pairs []KeyValuePair `json:"pairs"`
}
