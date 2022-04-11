# timeis
Show time by timezone.

## Example
Without arguments will be shown the time for current location.
```none
$ timeis
Timezone       Local time  Local date
----------     ----------  ----------
UTC            14:53:48    22-01-2022
Europe/Moscow  17:53:48    22-01-2022
Europe/Kiev    16:53:48    22-01-2022
```

With argument `--utc` (or `-u`) parsing it as an UTC time.
```none
$ timeis --utc 13:00:00
Timezone        Local time  Local date  Delta
----------      ----------  ----------  ----------
UTC             13:00:00    11-04-2022  +0000
Europe/Nicosia  16:00:00    11-04-2022  +0300
Europe/Moscow   16:00:00    11-04-2022  +0300
Asia/Tbilisi    17:00:00    11-04-2022  +0400
```

## Install
```bash
go install github.com/nemca/timeis@latest
```

## Example config
The configuration file is stored in the user's home directory.
It is in yaml format and name is `$HOME/.timeis.yaml`.
```yaml
---
timezones:
  - "Europe/Moscow"
  - "Europe/Kiev"
```
