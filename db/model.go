package db

type BloatedTable struct {
	Database    string  `sql:"current_database"`
	Schema      string  `sql:"schemaname"`
	Table       string  `sql:"tablename"`
	Bloat       float32 `sql:"tbloat"`
	WastedBytes int     `sql:"wastedbytes"`
	Indexes     []*BloatedIndex
}

type BloatedIndex struct {
	Name        string  `sql:"iname"`
	Bloat       float32 `sql:"ibloat"`
	WastedBytes int     `sql:"wastedibytes"`
}

type BloatedTableList []*BloatedTable
