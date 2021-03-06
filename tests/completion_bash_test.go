package tests

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestTopCompletion(t *testing.T) {
	testCompletion(t, nil, []string{"12to8", ""}, []string{"completion", "delete", "list", "new", "release"})
}

func TestCompletionCompletion(t *testing.T) {
	testCompletion(t, nil, []string{"12to8", "completion", ""}, []string{"bash"})
}

// TestAddPerformanceCompletion ensures that we can add a performance quickly (one letter + tab for each subcommand)
func TestAddPerformanceCompletion(t *testing.T) {
	testCompletion(t, nil, []string{"12to8", "n"}, []string{"new"})
	testCompletion(t, nil, []string{"12to8", "new", "p"}, []string{"performance"})
}

func TestReleaseTimesheetCompletion(t *testing.T) {
	testCompletion(t, nil, []string{"12to8", "r"}, []string{"release"})
	testCompletion(t, nil, []string{"12to8", "release", ""}, []string{"timesheet"})
}

func TestNewTimesheetCompletion(t *testing.T) {
	testCompletion(t, nil, []string{"12to8", "n"}, []string{"new"})
	testCompletion(t, nil, []string{"12to8", "new", "t"}, []string{"timesheet"})

	now := time.Now()
	year, month, _ := now.Date()
	var prevMonth int
	var prevYear int
	var nextMonth int
	var nextYear int
	if int(month) == 1 {
		prevMonth = 12
		prevYear = year - 1
	} else {
		prevMonth = int(month) - 1
		prevYear = year
	}
	if int(month) == 12 {
		nextMonth = 1
		nextYear = year + 1
	} else {
		nextMonth = int(month) + 1
		nextYear = year
	}
	possibleTimesheets := []string{}
	possibleTimesheets = append(possibleTimesheets, fmt.Sprintf("%02d/%d", prevMonth, prevYear))
	possibleTimesheets = append(possibleTimesheets, fmt.Sprintf("%02d/%d", month, year))
	possibleTimesheets = append(possibleTimesheets, fmt.Sprintf("%02d/%d", nextMonth, nextYear))
	testCompletion(t, nil, []string{"12to8", "new", "timesheet", ""}, possibleTimesheets)
}

func TestNewPerformanceContractCompletion(t *testing.T) {
	c := &dockerID{}
	c.start925r(t, "basic_projects")
	defer c.stop925r(t)
	testCompletion(t, c, []string{"12to8", "new", "pe"}, []string{"performance"})
	testCompletion(t, c, []string{"12to8", "new", "performance", "-c", ""}, []string{`"Go Consultancy [Python & Co]"`, `"Internal Stuff (c) [Golang Tech]"`})
	testCompletion(t, c, []string{"12to8", "new", "performance", "-c", `"`}, []string{"Go Consultancy [Python & Co]", "Internal Stuff (c) [Golang Tech]"})
	testCompletion(t, c, []string{"12to8", "new", "performance", "-c", `'`}, []string{"Go Consultancy [Python & Co]", "Internal Stuff (c) [Golang Tech]"})
}

func TestNewPerformanceRateCompletion(t *testing.T) {
	c := &dockerID{}
	c.start925r(t, "basic_projects")
	defer c.stop925r(t)
	testCompletion(t, c, []string{"12to8", "new", "performance", "-m", ""}, []string{"1.00", "2.00"})
	testCompletion(t, c, []string{"12to8", "new", "performance", "-m", "1"}, []string{"1.00"})
}

func TestNewPerformanceStandbyCompletion(t *testing.T) {
	testCompletion(t, nil, []string{"12to8", "new", "performance", "-t", ""}, []string{"activity", "standby"})
	testCompletion(t, nil, []string{"12to8", "new", "performance", "-t", "s"}, []string{"standby"})
}

func TestNewPerformanceCompletion(t *testing.T) {
	var dates []string
	for i := -5; i <= 5; i++ {
		d := time.Now().AddDate(0, 0, i)
		dates = append(dates, fmt.Sprintf("%02d/%02d/%d", d.Day(), int(d.Month()), d.Year()))
	}
	testCompletion(t, nil, []string{"12to8", "new", "performance", ""}, dates)
	testCompletion(t, nil, []string{"12to8", "new", "performance", "1", ""}, []string{"8.0"})
}

func completionBashCode(cli []string) string {
	flatcli := strings.Replace(strings.Join(cli, " "), `"`, `\"`, -1)
	flatcli = strings.Replace(flatcli, "'", `\'`, -1)
	words := len(cli) - 1
	return fmt.Sprintf(`
export PATH=..:$PATH
. /usr/share/bash-completion/bash_completion
. <(12to8 completion bash)
COMP_WORDS=(%s)
COMP_CWORD=%d
COMP_LINE="%s"
COMP_POINT=${#COMP_LINE}
_xfunc 12to8 __start_12to8
printf '%%s\n' "${COMPREPLY[@]}"
`, flatcli, words, flatcli)
}

func testCompletion(t *testing.T, c *dockerID, cli []string, expected []string) {
	var expectedOut bytes.Buffer
	for _, expectedLine := range expected {
		expectedOut.WriteString(fmt.Sprintf("%s\n", expectedLine))
	}
	tc := &CmdTestCase{
		Name:     "Autocomplete",
		Cmd:      "bash",
		Args:     []string{"-c", completionBashCode(cli)},
		OutLines: len(expected),
		OutText:  expectedOut.String(),
	}
	if c != nil {
		tc.Env = append(RunAsUser, c.EndpointEnv())
	}
	tc.Run(t)
}
