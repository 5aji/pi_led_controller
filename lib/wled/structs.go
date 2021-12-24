package wled

// colors are represented as an array of 3 bytes RGB.
type WledColor = [3]byte

type WledNightLight struct {
	On               bool  `json:"on"`
	Duration         uint8 `json:"dur"`
	Mode             uint8 `json:"mode"`
	TargetBrightness uint8 `json:"tbri"`
}
type WledSegment struct {
	Id      int  `json:"id"`
	On      bool `json:"on"`
	Start   int  `json:"start"`
	Stop    int  `json:"stop"`
	Length  int  `json:"len"`
	Group   int  `json:"grp"`
	Spacing int  `json:"spc"`
	Offset  int  `json:"of"`

	Colors [3]WledColor `json:"col"`

	EffectId        int   `json:"fx"`
	EffectSpeed     uint8 `json:"sx"`
	EffectIntensity uint8 `json:"ix"`

	PalleteId int  `json:"pal"`
	Selected  bool `json:"sel"`
	Reversed  bool `json:"rev"`
}

type WledState struct {
	On         bool           `json:"on"`
	Brightness uint8          `json:"bri"`
	Transition uint8          `json:"trn"`
	Preset     int            `json:"ps"`
	Nightlight WledNightLight `json:"nl"`
	Segments   []WledSegment  `json:"seg"`
}
