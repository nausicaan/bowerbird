package main

import (
	"encoding/json"
	"strings"
)

// Launch the program and execute the appropriate code
func main() {
	json.Unmarshal(apiget("db/lastfix.json"), &history)
	// json.Unmarshal(apiget(jira.LastFix), &history)
	var freebies, developers, deployments []string
	for index, element := range reads {
		json.Unmarshal(apiget(element), &swimlane)
		if swimlane.Total > 0 {
			for _, issue := range swimlane.Issues {
				ticket = issue.Key
				plugin = issue.Fields.Summary
				switch index {
				case 0: // ToDo SwimLane
					if strings.Contains(plugin, "-premium-") {
						flag = "-p"
						premium()
					} else if strings.Contains(plugin, "bcgov-plugin") || strings.Contains(plugin, "bcgov-theme") {
						developers = append(developers, plugin, ticket)
					} else {
						freebies = append(freebies, plugin, ticket)
					}
				case 1: // Testing SwimLane
					upd := issue.Fields.Updated
					passed := (amount(upd[:26] + string(':') + upd[26:])).Hours()
					if passed > benchmark {
						deployments = append(deployments, plugin, ticket)
						issue.Fields.FixVersions[0].Name = convert()
						issue.Fields.Status.Category.Name = "Ready for Deploy"
						// body, _ := json.Marshal(swimlane)
						// execute("-e", "curl", "-D-", "-X", "POST", "-d", string(body), "-H", "Authorization: Bearer "+jira.Token, "-H", "Content-Type: application/json", jira.Base+"issue/"+ticket)
					}
				}
			}
		}
	}

	// changedir()
	// if len(freebies) > 0 {
	// 	flag = "-w"
	// 	wpackagist(freebies)
	// }

	// if len(developers) > 0 {
	// 	flag = "-r"
	// 	inhouse(developers)
	// }

	// if len(deployments) > 0 {
	// 	flag = "-r"
	// 	releases(deployments)
	// }
}
