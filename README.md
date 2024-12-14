<h1 align="center">Global Entry Slot Notifier</h1>

<p align="center">
   <img src="./globalentry.png" width="80" alt="Global Entry"/><img src="./globalentry.png" width="80" alt="Global Entry"/><img src="./globalentry.png" width="80" alt="Global Entry"/>
</p>

Global Entry Slot Notifier is a command-line tool that checks for available Global Entry appointment slots at specified locations 
and sends notifications via native system notification or to [ntfy app](https://ntfy.sh/), available on Android, iPhone, and Mac.

## Features

- Periodically checks for available Global Entry appointment slots.
- Sends notifications via system notification **or** ntfy app (using [ntfy.sh](https://ntfy.sh/)).
- Configurable via command-line flags.

## Install

1. You can either build the binary using `go` (make sure you have go installed) with the `make` command

```bash
 make all
```

2. Or Download the latest binary from the [releases](https://github.com/arun0009/global-entry-slot-notifier/releases) page 
   as `global-entry-slot-notifier` (with `.exe` extension for windows) 
   
   Note: You may need to grant permission to run the binary.

### Usage

```bash
 ./global-entry-slot-notifier -l <location_id> -n <notifier_type> [-t <ntfy_topic>] [-i <duration>]
```

### Flags
* `-l`, `--location` (required): Specify the [location ID](https://github.com/arun0009/global-entry-slot-notifier?tab=readme-ov-file#pick-your-location-id-from-below-to-use-in-flag-above) for the Global Entry appointment (see below location ids).
* `-n`, `--notifier` (required): Specify the notifier type (app or system).
* `-t`, `--topic`    (required if notifier is app): Specify the ntfy.sh [topic](https://docs.ntfy.sh/) to send notifications to.
* `-i`, `--interval` (optional): Specify the interval (in seconds, e.g. 30s) at which to check for available appointments. Default is 60s.

### Examples

1. System Notification

```bash
 ./global-entry-slot-notifier -l 5446 -n system -i 90s
```

2. App Notification (first create your [topic on ntfy app](https://docs.ntfy.sh/))

```bash
 ./global-entry-slot-notifier -l 5446 -n app -t my-ntfy-topic
```

##### Pick your location id from below to use in flag (above)

| ID    | Enrollment Center Name                                                                                                                |
|-------|---------------------------------------------------------------------------------------------------------------------------------------|
| 5001  | Hidalgo Enrollment Center                                                                                                             |
| 5002  | Otay Mesa Enrollment Center                                                                                                           |
| 5003  | Brownsville Enrollment Center                                                                                                         |
| 5004  | Laredo Enrollment Center                                                                                                              |
| 5005  | El Paso Enrollment Center                                                                                                             |
| 5006  | Calexico Enrollment Center                                                                                                            |
| 5007  | Nogales, AZ                                                                                                                           |
| 5021  | Champlain NEXUS and FAST                                                                                                              |
| 5023  | Detroit Enrollment Center Global Entry                                                                                                |
| 5024  | Port Huron Enrollment Center                                                                                                          |
| 5060  | Warroad Enrollment Center                                                                                                             |
| 5080  | Sault Ste Marie Enrollment Center                                                                                                     |
| 5101  | Houlton POE/Woodstock                                                                                                                 |
| 5120  | Sweetgrass NEXUS and FAST Enrollment Center                                                                                           |
| 5140  | JFK International Global Entry EC                                                                                                     |
| 5141  | Houston Intercontinental Global Entry EC                                                                                              |
| 5142  | Washington Dulles International Global Entry EC                                                                                     |
| 5161  | Niagara Falls Enrollment Center                                                                                                       |
| 5180  | Los Angeles International Global Entry EC                                                                                             |
| 5181  | Miami International Airport                                                                                                           |
| 5182  | Atlanta International Global Entry EC                                                                                                 |
| 5183  | Chicago O'Hare International Global Entry EC                                                                                        |
| 5223  | Derby Line Enrollment Center                                                                                                          |
| 5300  | Dallas-Fort Worth International Airport Global Entry                                                                                  |
| 5320  | Detroit Metro Airport                                                                                                                 |
| 5340  | Honolulu Enrollment Center                                                                                                            |
| 5360  | Las Vegas Enrollment Center                                                                                                           |
| 5380  | Orlando International Airport                                                                                                         |
| 5400  | San Juan Global Entry Enrollment Center                                                                                               |
| 5420  | Seatac International Airport Global Entry EC                                                                                          |
| 5441  | Boston-Logan Global Entry Enrollment Center                                                                                             |
| 5443  | Fort Lauderdale Global Entry Enrollment Center                                                                                        |
| 5444  | Newark Liberty Intl Airport                                                                                                           |
| 5445  | Philadelphia International Airport                                                                                                    |
| 5446  | San Francisco Global Entry Enrollment Center                                                                                          |
| 5447  | Sanford Global Entry Enrollment Center                                                                                                |
| 5460  | San Luis Enrollment Center                                                                                                            |
| 5500  | Calais Enrollment Center                                                                                                              |
| 5520  | Lansdowne (Thousand Islands Bridge)                                                                                                   |
| 6480  | U.S. Custom House - Bowling Green                                                                                                     |
| 6840  | Minneapolis - St. Paul Global Entry EC                                                                                              |
| 6920  | Doha International Airport                                                                                                            |
| 6940  | Denver International Airport                                                                                                          |
| 7160  | Phoenix Sky Harbor Global Entry Enrollment Center                                                                                     |
| 7520  | San Antonio International Airport                                                                                                     |
| 7540  | Anchorage Enrollment Center                                                                                                           |
| 7600  | Salt Lake City International Airport                                                                                                |
| 7680  | Cincinnati Enrollment Center                                                                                                          |
| 7740  | Milwaukee Enrollment Center                                                                                                           |
| 7820  | Austin-Bergstrom International Airport                                                                                                  |
| 7940  | Baltimore Washington Thurgood Marshall  Airport                                                                                       |
| 7960  | Portland, OR Enrollment Center                                                                                                        |
| 8020  | Tampa Enrollment Center                                                                                                               |
| 8040  | Albuquerque Enrollment Center                                                                                                         |
| 8100  | Douglas Enrollment Center                                                                                                             |
| 8120  | Washington, DC Enrollment Center                                                                                                      |
| 9040  | Singapore (Singapore, U.S. Embassy)                                                                                                 |
| 9101  | Grand Portage                                                                                                                         |
| 9140  | Guam International Airport                                                                                                            |
| 9180  | Cleveland U.S. Customs and border protection                                                                                          |
| 9200  | Pittsburgh International Airport                                                                                                      |
| 9240  | Tucson Enrollment Center                                                                                                              |
| 9260  | West Palm Beach Enrollment Center                                                                                                       |
| 9740  | New Orleans Enrollment Center                                                                                                         |
| 10260 | Nashville Enrollment Center                                                                                                           |
| 11001 | Rockford-Chicago International Airport                                                                                                |
| 11002 | Peoria international airport                                                                                                          |
| 11841 | Port Clinton, Ohio Enrollment Center                                                                                                  |
| 11981 | Chicago Field Office Enrollment Center                                                                                                |
| 12021 | St. Louis Enrollment Center                                                                                                           |
| 12161 | Boise Enrollment Center                                                                                                               |
| 12781 | Kansas City Enrollment Center                                                                                                         |
| 13321 | Blaine Global Entry Enrollment Center                                                                                                 |
| 13621 | Memphis Intl Airport Global Enrollment Center                                                                                           |
| 14181 | International Falls Global Entry Enrollment Center                                                                                    |
| 14321 | Charlotte-Douglas International Airport                                                                                               |
| 14381 | Fairbanks Enrollment Center                                                                                                           |
| 14481 | Gulfport-Biloxi Global Entry Enrollment Center                                                                                        |
| 14681 | Bradley International Airport Enrollment Center                                                                                         |
| 14981 | Richmond, VA Enrollment Center                                                                                                        |
| 15221 | Pembina Global Entry Enrollment Center                                                                                                |
| 16226 | Eagle Pass                                                                                                                            |
| 16242 | Dayton Enrollment Center                                                                                                              |
| 16248 | Treasure Coast International Airport                                                                                                |
| 16251 | Sweetgrass Global Entry Enrollment Center                                                                                               |
| 16271 | Tri-cities Enrollment Center                                                                                                          |
| 16277 | Huntsville Global Entry Enrollment Center                                                                                             |
| 16278 | South Bend Enrollment Center                                                                                                          |
| 16282 | Mobile Regional Airport Enrollment Center                                                                                               |
| 16460 | Del Rio Enrollment Center                                                                                                             |
| 16461 | Des Moines GE Enrollment Center                                                                                                       |
| 16463 | Springfield – Branson National Airport, MO                                                                                          |
| 16467 | Omaha, NE Enrollment Center                                                                                                           |
| 16475 | SEAFO - Bozeman Airport                                                                                                               |
| 16488 | SEAFO - Missoula International Airport                                                                                                |
| 16519 | Champlain Global Entry                                                                                                                |
| 16547 | San Diego International Airport EC                                                                                                    |
| 16555 | Norfolk EC                                                                                                                            |
| 16586 | Atlantic City  Airport Global Entry Mobile Event                                                                                      |
| 16610 | BAL-FO Wilmington Delaware Enrollment Center                                                                                            |
| 16611 | BAL-FO Harrisburg Enrollment Center                                                                                                     |
| 16657 | Chicago Mobile Event                                                                                                                  |
| 16661 | PortMiami                                                                                                                             |
| 16662 | Fort Lauderdale International Airport (Terminal 1)                                                                                  |
| 16672 | Moline-Quad Cities International Airport                                                                                              |
| 16674 | Ontario Intl Airport GE (California)                                                                                                |
| 16683 | Fargo Satellite Enrollment Center                                                                                                     |
| 16693 | Special Event-Ontario Intl Airport(California)2024                                                                                  |
| 16699 | Port Huron Enrollment Center - RI                                                                                                     |
| 16705 | CFO - Louisville Intl Airport 2024                                                                                                  |
| 16709 | New Jersey Metro Area Mobile EC                                                                                                       |
| 16711 | Niagara Falls Enrollment Center GE Only                                                                                               |
| 16717 | San Juan Seaport                                                                                                                      |
| 16723 | Shreveport Regional Airport                                                                                                           |
| 16734 | Blue Grass Airport 2024                                                                                                               |
| 16748 | New York Metro Area Mobile EC                                                                                                         |
| 16755 | ERIE-Tom Ridge Field 2024                                                                                                             |
| 16759 | Alexandria Bay, NY - U.S. Port of Entry - GE ONLY                                                                                     |
| 16771 | Ogdensburg Enrollment Center - GE Only                                                                                                |
| 16781 | Chicago FO/ Wichita Airport                                                                                                            |
| 16793 | Houston Field Office GEEC                                                                                                             |
| 16809 | St. Thomas Cyril E. King Airport Enrollment Center                                                                                    |
| 16837 | Warwick RI Enrollment Center                                                                                                          |


Note: If you don't find your location above, please look at the updated [list](https://ttp.cbp.dhs.gov/schedulerapi/locations/?temporary=false&inviteOnly=false&operational=true&serviceName=Global%20Entry
)

### Advanced

You can download the binary on [raspberry pi](https://www.raspberrypi.com/) or on cloud e.g. [alwaysdata free tier](https://www.alwaysdata.com/en/) 
and run it as background process to notify you via [ntfy.sh](https://ntfy.sh/))

```bash
 curl -L https://github.com/arun0009/global-entry-slot-notifier/releases/download/v1.0/global-entry-slot-notifier_1.0_linux_amd64 -o global-entry-slot-notifier
 ./global-entry-slot-notifier -l 5446 -n app -t my-ntfy-topic &
```

### License
This project is licensed under the MIT License. See the LICENSE file for details.

### Contributing
Contributions are welcome! Please open an issue or submit a pull request for any changes.

### Contact
For questions or feedback, please contact arun@gopalpuri.com