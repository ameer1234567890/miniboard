package models

// Conf - web gui config
type Conf struct {
	Host     string
	Port     string
	Theme    string
	Color    string
	Icon     string
	ConfPath string
	YamlPath string
	ColorOn  string
	ColorOff string
	BtnWidth string
	NodePath string
}

// Host - panel element
type Host struct {
	Name  string `yaml:"name"`
	Addr  string `yaml:"addr"`
	Port  string `yaml:"port"`
	URL   string `yaml:"url"`
	Icon  string `yaml:"icon"`
	State bool   `yaml:"state"`
}

// Panel - tab element
type Panel struct {
	Name  string       `yaml:"name"`
	Scan  bool         `yaml:"scan"`
	Hosts map[int]Host `yaml:"hosts"`
}

// Tab - board element
type Tab struct {
	Name   string         `yaml:"name"`
	Panels map[int]string `yaml:"panels"`
}

// Links - all links
type Links struct {
	Tabs   map[int]Tab      `yaml:"tabs"`
	Panels map[string]Panel `yaml:"panels"`
}

// GuiData - web gui data
type GuiData struct {
	Config     Conf
	Themes     []string
	Links      Links
	Panels     map[int]Panel
	CurrentTab string
	TabEdit    int
}
