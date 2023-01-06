[![Build Status](https://api.travis-ci.com/sportorg/pysport.svg?branch=develop)](https://travis-ci.com/sportorg/pysport)
[![Python 3.8](https://img.shields.io/badge/python-v3.8-blue.svg?logo=pythonlang)](https://www.python.org/downloads/)
[![License: GPL v3](https://img.shields.io/badge/license-GPLv3-blue.svg)](https://github.com/sportorg/pysport/blob/develop/LICENSE)
[![Orienteering](https://img.shields.io/badge/sport-Alpine%20Ski-green)](https://github.com/sportorg)
[![Sportorg version](https://img.shields.io/github/v/release/sportorg/pysport)](https://github.com/sportorg/pysport)
[![Orienteering](https://img.shields.io/github/stars/sportorg/pysport?style=social)](https://github.com/sportorg/pysport)

![Bibprintout sportorg](img/logo.png)

# PySport v1.6.1 (SportOrg Fork)

Especially reconfigured for alpine ski competition.
- 

- [ ] Splits GUI removed
- [ ] No start time
- [ ] First run and second run introduced
- [ ] Modified Telegram functionality (Ctrl+T fast shortcut), help legend
- [ ] Modified Reports (Start List and preliminary results)
- [ ] Fixed ini file


```commandline
pip install poetry
poetry install
poetry install -E win  # for Windows
```

Run

Add `DEBUG=True` to `.env` file or `cp .env.example .env`

```commandline
poetry run python SportOrg.pyw
```
# Setup Alpine Ski mode
- [ ] Set two decimals in time control settings
![Mainwindow sportorg](img/Time_settings.png)
 
- [ ] Setup Alpine Ski mode in competition settings
![Bibprintout sportorg](img/Special.png)

- [ ] Refactored preliminary results protocol
![Bibprintout sportorg](img/Protocole.png)

- [ ] Now you-re able to input result in alpine ski fashion
  ![Bibprintout sportorg](img/Results_enter.png)

## build

### cx_Freeze

`python setup_.py build`

## HiDPI issue

`To fix small fonts on HiDPI displays run following in command line`

```commandline
setx QT_AUTO_SCREEN_SCALE_FACTOR "1"
setx QT_AUTO_SCREEN_SCALE_FACTOR "1" /M
setx QT_FONT_DPI "196"
```

## Roadmap

- [ ] Publish to pypi
- [ ] Finish proprietary timing system module (LoRa to USB)
