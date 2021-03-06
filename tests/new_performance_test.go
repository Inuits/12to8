package tests

import (
	"fmt"
	"testing"
	"time"
)

// TestNewPerformance tests that we can add a new performance
func TestNewPerformance(t *testing.T) {
	c := &dockerID{}
	c.start925r(t, "basic_projects")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	dayInThisMonth := fmt.Sprintf("5/%d/%d", int(time.Now().Month()), time.Now().Year())
	newTimesheet(t, c)
	(&CmdTestCase{
		Name: "Create performance",
		Env:  userEnv,
		Args: []string{"new", "performance", "-c", "Internal Stuff (c) [Golang Tech]", dayInThisMonth, "3", "fix jenkins"},
	}).Run(t)
	(&CmdTestCase{
		Name:     "List performances",
		Env:      userEnv,
		Args:     []string{"list", "performances", "-C", "day,contract,description,duration,multiplier,type"},
		OutLines: 5,
		OutText: `+-----+--------------------+-------------+------+------+----------+
| DAY |      CONTRACT      | DESCRIPTION |  H   |  X   |   KIND   |
+-----+--------------------+-------------+------+------+----------+
|   5 | Internal Stuff (c) | fix jenkins | 3.00 | 1.00 | activity |
+-----+--------------------+-------------+------+------+----------+
`,
	}).Run(t)
}

func TestNewPerformanceWithDoubleRate(t *testing.T) {
	c := &dockerID{}
	c.start925r(t, "basic_projects")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	dayInThisMonth := fmt.Sprintf("5/%d/%d", int(time.Now().Month()), time.Now().Year())
	newTimesheet(t, c)
	(&CmdTestCase{
		Name: "Create performance",
		Env:  userEnv,
		Args: []string{"new", "performance", "-m", "2.00", "-c", "Internal Stuff (c) [Golang Tech]", dayInThisMonth, "3", "fix jenkins"},
	}).Run(t)
	(&CmdTestCase{
		Name:     "List performances",
		Env:      userEnv,
		Args:     []string{"list", "performances", "-C", "day,contract,description,duration,multiplier,type"},
		OutLines: 5,
		OutText: `+-----+--------------------+-------------+------+------+----------+
| DAY |      CONTRACT      | DESCRIPTION |  H   |  X   |   KIND   |
+-----+--------------------+-------------+------+------+----------+
|   5 | Internal Stuff (c) | fix jenkins | 3.00 | 2.00 | activity |
+-----+--------------------+-------------+------+------+----------+
`,
	}).Run(t)
}

func TestNewPerformanceWithWrongRate(t *testing.T) {
	c := &dockerID{}
	c.start925r(t, "basic_projects")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	dayInThisMonth := fmt.Sprintf("5/%d/%d", int(time.Now().Month()), time.Now().Year())
	newTimesheet(t, c)
	(&CmdTestCase{
		Name:          "Create performance",
		Env:           userEnv,
		Args:          []string{"new", "performance", "-m", "3.00", "-c", "Internal Stuff (c) [Golang Tech]", dayInThisMonth, "3", "fix jenkins"},
		ExpectFailure: true,
		OutLines:      3,
		OutText: `Rate 3.00 not found. Possible rates:
 * 1.00 : Normal Hours
 * 2.00 : Double
`,
	}).Run(t)
}

func TestNewPerformanceFloat(t *testing.T) {
	c := &dockerID{}
	c.start925r(t, "basic_projects")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	dayInThisMonth := fmt.Sprintf("9/%d/%d", int(time.Now().Month()), time.Now().Year())
	newTimesheet(t, c)
	(&CmdTestCase{
		Name: "Create performance",
		Env:  userEnv,
		Args: []string{"new", "performance", "-c", "Internal Stuff (c) [Golang Tech]", dayInThisMonth, "4.47", "implement float"},
	}).Run(t)
	(&CmdTestCase{
		Name:     "List performances",
		Env:      userEnv,
		Args:     []string{"list", "performances", "-C", "day,contract,description,duration,multiplier,type"},
		OutLines: 5,
		OutText: `+-----+--------------------+-----------------+------+------+----------+
| DAY |      CONTRACT      |   DESCRIPTION   |  H   |  X   |   KIND   |
+-----+--------------------+-----------------+------+------+----------+
|   9 | Internal Stuff (c) | implement float | 4.47 | 1.00 | activity |
+-----+--------------------+-----------------+------+------+----------+
`,
	}).Run(t)
}
