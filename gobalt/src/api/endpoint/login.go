package endpoint

type SaltCredsReturn struct {
	Token string `json:"token"`
}

type SaltAPIListReturn struct {
	Return []SaltCredsReturn `json:"return"`
}

type SaltAPILocalReturn struct {
	Local []SaltAPIListReturn
}

type SaltKeyReturn struct {
	Local           []string `json:"local"`
	MinionsRejected []string `json:"minions_rejected"`
	MinionsDenied   []string `json:"minions_denied"`
	MinionsPre      []string `json:"minions_pre"`
	Minions         []string `json:"minions"`
}

type SaltAPIKeyReturn struct {
	Return SaltKeyReturn `json:"return"`
}
