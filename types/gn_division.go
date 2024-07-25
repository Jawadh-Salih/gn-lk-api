package types

type GNDivision struct {
	LifeCode      string `json:"life_code"`
	GnCode        string `json:"gn_code"`
	NameSinhala   string `json:"name_sinhala"`
	NameTamil     string `json:"name_tamil"`
	NameEnglish   string `json:"name_english"`
	MpaCode       string `json:"mpa_code"`
	Province      string `json:"province"`
	District      string `json:"district"`
	DivisionalSec string `json:"divisional_sec"`
}
