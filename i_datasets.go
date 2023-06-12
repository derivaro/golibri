package golibri

type varenvs struct {
	Keys      []string   `yaml:"keys"`
	Databases []database `yaml:"databases"`
}

type database struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
	Typ  string `yaml:"typ"`
}

type Row struct {
	FI []string
	Id string
}

type Datset struct {
	Name      string
	ColsCount int
	RowsCount int
	Cols      *Row
	Typs      *Row
	Gtyps     *Row
	Rows      *[]Row
	C         map[string]int
	Ptyps     *Row
}

func (DS *Datset) SnowColsConv() {

	var j int
	for j = 0; j < DS.ColsCount; j++ {

		DS.Gtyps.FI[j] = "number"
		DS.Ptyps.FI[j] = "numeric"
		if DS.Typs.FI[j] == "VARCHAR" {
			DS.Gtyps.FI[j] = "string"
			DS.Ptyps.FI[j] = "text"
		}
		if DS.Typs.FI[j] == "TEXT" {
			DS.Gtyps.FI[j] = "string"
			DS.Ptyps.FI[j] = "text"
		}
		if DS.Typs.FI[j] == "DATE" {
			DS.Gtyps.FI[j] = "timestamp"
			DS.Ptyps.FI[j] = "timestamp"
		}
		if DS.Typs.FI[j] == "DATETIME" {
			DS.Gtyps.FI[j] = "timestamp"
			DS.Ptyps.FI[j] = "timestamp"
		}
		if DS.Typs.FI[j] == "TIMESTAMP" {
			DS.Gtyps.FI[j] = "timestamp"
			DS.Ptyps.FI[j] = "timestamp"
		}
		if DS.Typs.FI[j] == "TIMESTAMP_TZ" {
			DS.Gtyps.FI[j] = "timestamp"
			DS.Ptyps.FI[j] = "timestamp"
		}
		if DS.Typs.FI[j] == "TIMESTAMP_NTZ" {
			DS.Gtyps.FI[j] = "timestamp"
			DS.Ptyps.FI[j] = "timestamp"
		}
		if DS.Typs.FI[j] == "BOOLEAN" {
			DS.Gtyps.FI[j] = "boolean"
			DS.Ptyps.FI[j] = "boolean"
		}
	}

}
