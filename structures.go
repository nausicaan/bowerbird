package main

// Atlassian builds a list of jira tokens and api addresses
type Atlassian struct {
	Base    string `json:"base"`
	Token   string `json:"token"`
	Done    string `json:"done"`
	ToDo    string `json:"todo"`
	Deploy  string `json:"deploy"`
	Testing string `json:"testing"`
	LastFix string `json:"lastfix"`
}

// Desso contains the JSON parameters for a new Jira ticket
type Desso struct {
	Total  int64 `json:"total"`
	Issues []struct {
		Key    string `json:"key,omitempty"`
		Fields struct {
			Issuetype struct {
				ID string `json:"id,omitempty"`
			} `json:"issuetype,omitempty"`
			Status struct {
				Category struct {
					ID   int64  `json:"id,omitempty"`
					Name string `json:"name,omitempty"`
				} `json:"statusCategory,omitempty"`
			} `json:"status,omitempty"`
			Project struct {
				ID  string `json:"id,omitempty"`
				Key string `json:"key,omitempty"`
			} `json:"project,omitempty"`
			Labels      []string     `json:"labels,omitempty"`
			Updated     string       `json:"updated,omitempty"`
			Summary     string       `json:"summary,omitempty"`
			FixVersions []FixVersion `json:"fixVersions,omitempty"`
		} `json:"fields,omitempty"`
	} `json:"issues,omitempty"`
}

// FixVersion holds the information about a release
type FixVersion struct {
	Name        string `json:"name,omitempty"`
	Released    bool   `json:"released,omitempty"`
	ReleaseDate string `json:"releaseDate,omitempty"`
}

// Satis structure captures the contents of the composer.json file for typical premium plugins
type Satis struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
}

// Event structure captures the contents of the composer.json file for Events Calendar related items
type Event struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
	Require struct {
		EventsCalendar string `json:"wpackagist-plugin/the-events-calendar"`
	} `json:"require"`
}

const (
	bv        string  = "3.0"
	benchmark float64 = 168.00
	upbranch  string  = "update/"
	relbranch string  = "release/"
	halt      string  = "program halted "
	zero      string  = "Insufficient arguments supplied -"
)

var (
	event    Event
	satis    Satis
	history  Desso
	swimlane Desso
	jira     Atlassian
	flag     string
	plugin   string
	release  string
	ticket   string
	folder   []string
	number   []string
	queries  = []string{jira.ToDo, jira.Testing, jira.Deploy}
	reads    = []string{"db/todo.json", "db/inprogress.json", "db/empty.json"}
)
