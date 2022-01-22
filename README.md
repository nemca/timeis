# timeis
Show time by timezone.

## Example
```bash
$ timeis
Timezone       Local time  Local date
----------     ----------  ----------
UTC            14:53:48    22-01-2022
Europe/Moscow  17:53:48    22-01-2022
Europe/Kiev    16:53:48    22-01-2022
```

## Install
```bash
go get -u github.com/nemca/timeis
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
