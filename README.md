## Nepali

[![CI status](https://github.com/opensource-nepal/go-nepali/actions/workflows/nepali-ci.yml/badge.svg?branch=main)](https://github.com/opensource-nepal/go-nepali/actions)
[![codecov](https://codecov.io/gh/opensource-nepal/go-nepali/branch/main/graph/badge.svg?token=PTUHYWCJ4I)](https://codecov.io/gh/opensource-nepal/go-nepali)
[![Release](https://img.shields.io/github/v/release/opensource-nepal/go-nepali?label=Latest)](https://github.com/opensource-nepal/go-nepali/releases)
[![License](https://img.shields.io/github/license/opensource-nepal/go-nepali?label=License)](https://github.com/opensource-nepal/go-nepali/blob/main/LICENSE)


This is a `golang` implementation of the package from [opensource-nepal](https://github.com/opensource-nepal)'s `python` package [nepali](https://github.com/opensource-nepal/py-nepali).

We'll try to include all the features provided in the aforementioned package. We'll be releasing features in parts to support all or most features given by `go`'s `time` package.

### Installing

To install this package use `go get`

```shell
$ go get -u github.com/opensource-nepal/go-nepali
```

Or if there is some error then, use `go install`

```shell
$ go install github.com/opensource-nepal/go-nepali
```

#### Usage

_NOTE: Currently this package is in beta version, so it only includes basic features like date conversion, formatting and parsing. Also, the output is read-only._

In this package, we provide 2 `go` packages, `nepalitime` and `dateConverter`.

1. `nepalitime`: The functionalities provided in `nepalitime` are described below:

   1. To get a `NepaliTime` object:

      ```go
      import "github.com/opensource-nepal/go-nepali/nepalitime"

      npTime, err := nepalitime.Date(2079,  10,  15,  14,  29,  6,  7)
      ```

      The `err` object could containing error if you pass in errorneous values into the `Date` function.
      
      Once you have the `NepaliTime` object, you will be able to access methods, some of which are described below:
      1. `GetEnglishTime()` -> Returns the corresponding English time
      2. `Date()` -> Returns year, month and day as separate return values
      3. `Clock()` -> Returns hour, minute and second as separate return values

   2. If you already have a `time.Time` object then, you can get the corresponding `NepaliTime` object using the function `FromEnglishTime`:

      ```go
      import "github.com/opensource-nepal/go-nepali/nepalitime"

      // just to ensure that the timezone is Asia/Kathmandu
      // else the time will be adjusted accordingly
      enTime := time.Date(2023, 1, 29, 14, 29, 6, 7, nepalitime.GetNepaliLocation())
      nt, err := nepalitime.FromEnglishTime(enTime)
      ```

   3. To parse a date string into a `NepaliTime` object the `Parse` function can be used. This is the Nepali equivalent of the `time.Parse` function of go but instead of using the time parsing format of `Mon Jan 2 15:04:05 -0700 MST 2006` we decided to go with the `%Y/%m/%d` style parsing. We intend on supporting the go style formatting in the upcoming releases. Please see [directives](#date-directives) section to know which directives we support.

      ```go
      import "github.com/opensource-nepal/go-nepali/nepalitime"

      datetimeStr := "2079/10/14"
      format := "%Y/%m/%d"

      npTime, err := nepalitime.Parse(datetimeStr, format)
      ```

   4. To get current Nepali time:

      ```go
      import "github.com/opensource-nepal/go-nepali/nepalitime"

      npTime := nepalitime.Now()
      ```

   5. To get the current English time in Nepal:

      ```go
      import "github.com/opensource-nepal/go-nepali/nepalitime"

      enTime := nepalitime.GetCurrentEnglishTime()
      ```
   
   6. To format the `NepaliTime` object to a string that you want.
      ```go
      
      import (
        "fmt"
        "github.com/opensource-nepal/go-nepali/nepalitime"
      )
      
      npTime, _ := nepalitime.Date(2079,  10,  15,  14,  29,  6,  7, 123)
      fmt.Println(npTime.Format("%Y/%m/%d %H:%M:%S:%f"))
      ```
      _Please see [directives](#date-directives) section to know which directives we support._

2. `dateConverter`: The functionalities provided in `dateConverter` are described below:

   1. This is one of the core functionalities in which an English date(not an object) is converted to Nepali date in parts i.e. year, month, day in an array:

      ```go
      import "github.com/opensource-nepal/go-nepali/nepalitime"

      // npDate is an array of 3 length which contains
      // year, month and day sequential form
      npDate, err := dateConverter.EnglishToNepali(2023, 1, 28)
      ```

   2. To convert a Nepali date(not an object) to English date in parts i.e. year, month, day in an array:

      ```go
      import "github.com/opensource-nepal/go-nepali/nepalitime"

      // enDate is an array of 3 length which contains
      // year, month and day sequential form
      enDate, err := dateConverter.NepaliToEnglish(2087, 8, 10)
      ```

#### Date Directives

| Directive | Meaning                                                  | Example                                  |
|-----------|----------------------------------------------------------|------------------------------------------|
| `%d`      | Day of the month as a zero-padded decimal number.        | 01, 02, …, 31                            |
| `%-d`     | Day of the month as a decimal number.                    | 1, 2, …, 31                              |
| `%m`      | Month as a zero-padded decimal number.                   | 01, 02, …, 12                            |
| `%-m`     | Month as a decimal number.                               | 1, 2, …, 12                              |
| `%y`      | Year without century as a zero-padded decimal number.    | 00, 01, …, 99                            |
| `%-y`     | Year without century as a decimal number.                | 0, 1, …, 99                              |
| `%Y`      | Year with century as a zero-padded decimal number.       | 0001, 0002, …, 2078, 2079, …, 9998, 9999 |
| `%-Y`     | Year with century as a decimal number.                   | 1, 2, …, 2078, 2079, …, 9998, 9999       |
| `%H`      | Hour (24-hour clock) as a zero-padded decimal number.    | 00, 01, …, 23                            |
| `%-H`     | Hour (24-hour clock) as a decimal number.                | 0, 1, …, 23                              |
| `%I`      | Hour (12-hour clock) as a zero-padded decimal number.    | 01, 02, …, 12                            |
| `%-I`     | Hour (12-hour clock) as a decimal number.                | 1, 2, …, 12                              |
| `%p`      | Locale’s equivalent of either AM or PM.                  | AM, PM (en_US);<br><br>am, pm (de_DE)    |
| `%M`      | Minute as a zero-padded decimal number.                  | 00, 01, …, 59                            |
| `%-M`     | Minute as a decimal number.                              | 0, 1, …, 59                              |
| `%S`      | Second as a zero-padded decimal number.                  | 00, 01, …, 59                            |
| `%-S`     | Second as a decimal number.                              | 0, 1, …, 59                              |
| `%f`      | Nanosecond as a decimal number, zero-padded to 6 digits. | 000000, 000001, …, 999999                |
| `%-f`     | Nanosecond as a decimal number.                          | 0, 1, …, 999999                          |
| `%%`      | A literal `'%'` character.                               | %                                        |

## Contribution

We appreciate feedback and contribution to this package. To get started please see our [contribution guide](contributing.md)
