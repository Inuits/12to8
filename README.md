# 12to8

A client for [9to5r](https://github.com/kalmanolah/925r)

## Install

```
$ go get github.com/Inuits/12to8
```

## Usage

### Lists
List users:

```
12to8 list users
```

List timesheets:

```
12to8 list timesheets
```

List companies:

```
12to8 list companies
```


### Actions

Create new timesheet:

```
12to8 new timesheet
12to8 new timesheet 9
12to8 new timesheet 9/2017
```

Release timesheet:
```
12to8 release timesheet
12to8 release timesheet -f
12to8 release timesheet 9
12to8 release timesheet 9/2017
```

