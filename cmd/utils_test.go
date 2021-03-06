// Copyright © 2017 Julien Pivotto <roidelapluie@inuits.eu>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"testing"
	"time"
)

func TestGetMonthYearFromArg(t *testing.T) {
	m, y, err := getMonthYearFromArg("10/2017")
	if err != nil {
		t.Fatalf("(wanted) %v != (got) %v", nil, err)
	}
	if m != 10 {
		t.Fatalf("(wanted) %v != (got) %v", 10, m)
	}
	if y != 2017 {
		t.Fatalf("(wanted) %v != (got) %v", 2017, y)
	}
	m, y, err = getMonthYearFromArg("09/2007")
	if err != nil {
		t.Fatalf("(wanted) %v != (got) %v", nil, err)
	}
	if m != 9 {
		t.Fatalf("(wanted) %v != (got) %v", 9, m)
	}
	if y != 2007 {
		t.Fatalf("(wanted) %v != (got) %v", 2007, y)
	}
	m, y, err = getMonthYearFromArg("9")
	if err != nil {
		t.Fatalf("(wanted) %v != (got) %v", nil, err)
	}
	if m != 9 {
		t.Fatalf("(wanted) %v != (got) %v", 9, m)
	}
	currentYear := time.Now().Year()
	if y != currentYear {
		t.Fatalf("(wanted) %v != (got) %v", currentYear, y)
	}
}

func TestGetDaysMonthYearFromArg(t *testing.T) {
	currentYear := time.Now().Year()
	currentMonth := int(time.Now().Month())
	d, m, y, err := getDaysMonthYearFromArg("1/10/2007")
	if err != nil {
		t.Fatalf("(wanted) %v != (got) %v", nil, err)
	}
	if d != 1 {
		t.Fatalf("(wanted) %v != (got) %v", 1, d)
	}
	if m != 10 {
		t.Fatalf("(wanted) %v != (got) %v", 10, m)
	}
	if y != 2007 {
		t.Fatalf("(wanted) %v != (got) %v", 2007, y)
	}
	d, m, y, err = getDaysMonthYearFromArg("1/09")
	if err != nil {
		t.Fatalf("(wanted) %v != (got) %v", nil, err)
	}
	if d != 1 {
		t.Fatalf("(wanted) %v != (got) %v", 1, d)
	}
	if m != 9 {
		t.Fatalf("(wanted) %v != (got) %v", 9, m)
	}
	if y != currentYear {
		t.Fatalf("(wanted) %v != (got) %v", currentYear, y)
	}
	d, m, y, err = getDaysMonthYearFromArg("2")
	if err != nil {
		t.Fatalf("(wanted) %v != (got) %v", nil, err)
	}
	if d != 2 {
		t.Fatalf("(wanted) %v != (got) %v", 2, d)
	}
	if m != currentMonth {
		t.Fatalf("(wanted) %v != (got) %v", currentMonth, m)
	}
	if y != currentYear {
		t.Fatalf("(wanted) %v != (got) %v", currentYear, y)
	}
}
