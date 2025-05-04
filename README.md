<h1 align="center">Global Entry Slot Notifier</h1>

<p align="center">
   <img src="./globalentry.png" width="80" alt="Global Entry"/><img src="./globalentry.png" width="80" alt="Global Entry"/><img src="./globalentry.png" width="80" alt="Global Entry"/>
</p>

**Note: This tool (global-entry-slot-notifier) is designed to be run locally—you'll need to download and run the binary manually with command-line options.**

**If you're looking for a fully cloud-based, no-download solution, check out [global-entry-appointment](https://github.com/arun0009/global-entry-appointment), which lets you register for notifications directly via a [simple web page](https://arun0009.github.io/global-entry-appointment).**

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
 ./global-entry-slot-notifier -l <location_id> -n <notifier_type> [-t <ntfy_topic>] [-i <duration>] [-b <before_date>]
```

### Flags
* `-l`, `--location` (required): Specify the [location ID](https://github.com/arun0009/global-entry-slot-notifier?tab=readme-ov-file#pick-your-location-id-from-below-to-use-in-flag-above) for the Global Entry appointment (see below location ids).
* `-n`, `--notifier` (required): Specify the notifier type (app or system).
* `-t`, `--topic`    (required if notifier is app): Specify the ntfy.sh [topic](https://docs.ntfy.sh/) to send notifications to.
* `-i`, `--interval` (optional): Specify the interval (in seconds, e.g. 30s) at which to check for available appointments. Default is 60s.
* `-b`, `--before` (optional): Specify a cutoff date (YYYY-MM-DD) to only receive notifications for appointment slots before this date.

### Examples

1. System Notification

```bash
 ./global-entry-slot-notifier -l 5446 -n system -i 90s -b 2025-12-30
```

2. App Notification (first create your [topic on ntfy app](https://docs.ntfy.sh/))

```bash
 ./global-entry-slot-notifier -l 5446 -n app -t my-ntfy-topic -b 2025-12-30
```

##### Pick your location id from below to use in flag (above)

| ID   | Enrollment Center Name                                    |
|------|--------------------------------------------------------|
| 8040 | Albuquerque Enrollment Center - Albuquerque International Sunport 2200 Sunport Blvd SE Albuquerque NM 87106 |
| 16759 | Alexandria Bay, NY - U.S. Port of Entry - GE ONLY - 46735 I-81 Alexandria Bay NY 13607 |
| 7540 | Anchorage Enrollment Center - Ted Stevens International Airport 4600 Postmark Drive, RM NA 207 Anchorage AK 99502 |
| 5182 | Atlanta International Global Entry EC - 2600 Maynard H. Jackson Jr. Int'l Terminal Maynard H. Jackson Jr. Blvd. Atlanta GA 30320 |
| 16586 | Atlantic City Airport Global Entry Mobile Event - 101 Atlantic City International Airport Egg Harbor Township NJ 08234 |
| 7820 | Austin-Bergstrom International Airport - 3600 Presidential Blvd. Austin-Bergstrom International Airport Austin TX 78719 |
| 16611 | BAL-FO Harrisburg Enrollment Center - 1215 Manor Drive Suite 301 Mechanicsburg PA 17055 |
| 16610 | BAL-FO Wilmington Delaware Enrollment Center - 908 Churchmans Road Ext, New New Castle DE 19720 |
| 7940 | Baltimore Washington Thurgood Marshall Airport - Baltimore Washington Thurgood Marshall I Lower Level Door 18- Outer Street sign number 59 Linthicum MD 21240 |
| 13321 | Blaine Global Entry Enrollment Center - 8115 Birch Bay Square St. Suite 104 Blaine WA 98230 |
| 16734 | Blue Grass Airport 2024 - 4000 Terminal Drive Lexington KY 40510 |
| 12161 | Boise Enrollment Center - 4655 S Enterprise Street Boise ID 83705 |
| 5441 | Boston-Logan Global Entry Enrollment Center - Logan International Airport, Terminal E East Boston MA 02128 |
| 14681 | Bradley International Airport Enrollment Center - International Arrivals Building/ Terminal B Bradley Airport Windsor Locks CT 06096 |
| 5003 | Brownsville Enrollment Center - 700 Amelia Earhart Dr, Brownsville, Texas Brownsville South Padre Island International Airpo Brownsville TX 78521 |
| 16705 | CFO - Louisville Intl Airport 2024 - Louisville Muhammad Ali International Airport 700 Administration Drive Louisville KY 40209 |
| 5500 | Calais Enrollment Center - 3 Customs Street Calais ME 04619 |
| 5006 | Calexico Enrollment Center - 1699 East Carr Road Calexico CA 92231 |
| 16519 | Champlain Global Entry - 237 West Service Road Global Entry Only Champlain NY 12919 |
| 5021 | Champlain NEXUS and FAST - 237 West Service Road NEXUS, FAST, and Global Entry Enrollment Center Champlain NY 12919 |
| 14321 | Charlotte-Douglas International Airport - Charlotte-Douglas International Airport 5501 Josh Birmingham Parkway Charlotte NC 28208 |
| 16781 | Chicago FO/ Wichita Airport - 2277 Eisenhower Airport Parkway Wichita KS 67209 |
| 11981 | Chicago Field Office Enrollment Center - 610 S. CANAL STREET 6TH FLOOR CHICAGO IL 60607 |
| 16657 | Chicago Mobile Event - 2700 INTERNATIONAL DRIVE WEST CHICAGO IL 60185 |
| 5183 | Chicago O'Hare International Global Entry EC - 10000 West O'Hare Drive Terminal 5, Lower Level (Arrivals Floor) CHICAGO IL 60666 |
| 7680 | Cincinnati Enrollment Center - 4243 Olympic Blvd. Suite. 210 Erlanger KY 41018 |
| 9180 | Cleveland U.S. Customs and border protection - Customs & Border Protection 6747 Engle Road Middleburg Heights OH 44130 |
| 5300 | Dallas-Fort Worth International Airport Global Entry - DFW International Airport - Terminal D First Floor, Gate D22-23 DFW Airport TX 75261 |
| 16242 | Dayton Enrollment Center - 3800 Wright Drive Vandalia OH 45377 |
| 16460 | Del Rio Enrollment Center - 3140 Spur 239 Del Rio TX 78840 |
| 6940 | Denver International Airport - 8400 Denver International Airport Pena Boulevard Denver CO 80249 |
| 5223 | Derby Line Enrollment Center - 107 I-91 South Derby Line VT 05830 |
| 16461 | Des Moines GE Enrollment Center - 6100 Fleur Drive Des Moines IA 50321 |
| 5023 | Detroit Enrollment Center Global Entry - 2810 W. Fort Street Suite 124 Detroit MI 48216 |
| 5320 | Detroit Metro Airport - Detroit Evans Terminal 601 Rogell Dr., Suite 1271 Detroit MI 48242 |
| 6920 | Doha International Airport - Hamad International Airport Doha |
| 8100 | Douglas Enrollment Center - 1 Pan American Aveue (Cargo Facility) Douglas AZ 85607 |
| 16755 | ERIE-Tom Ridge Field 2024 - 4411 West 12th Street Erie PA 16505 |
| 16226 | Eagle Pass - 160 E. Garrison St. Eagle Pass TX 78852 |
| 5005 | El Paso Enrollment Center - 797 S. Zaragoza Rd. Bldg. A El Paso TX 79907 |
| 14381 | Fairbanks Enrollment Center - 6450 Airport Way - Suite 13 Room 1320A Fairbanks AK 99709 |
| 16683 | Fargo Satellite Enrollment Center - 3803 20th Street North Fargo ND 58102 |
| 5443 | Fort Lauderdale Global Entry Enrollment Center - 1800 Eller Drive Suite 104 Ft Lauderdale FL 33316 |
| 16662 | Fort Lauderdale International Airport (Terminal 1) - 100 Terminal Drive Fort Lauderdale FL 33315 |
| 9101 | Grand Portage - 9403 E Highway 61 Grand Portage MN 55605 |
| 9140 | Guam International Airport - 355 Chalan PasaHeru Suite B 224-B Tamuning GU 96913 |
| 14481 | Gulfport-Biloxi Global Entry Enrollment Center - Gulfport-Biloxi International Airport 14035 Airport Road, 2nd Floor (Main Terminal) Gulfport MS 39503 |
| 5001 | Hidalgo Enrollment Center - Anzalduas International Bridge 5911 S. STEWART ROAD Mission TX 78572 |
| 5340 | Honolulu Enrollment Center - 300 Rodgers Blvd Honolulu HI 96819 |
| 5101 | Houlton POE/Woodstock - 27 Customs Loop Houlton ME 04730 |
| 16793 | Houston Field Office GEEC - 2323 S. Shepherd Drive Houston TX 77019 |
| 5141 | Houston Intercontinental Global Entry EC - 3870 North Terminal Road Terminal E Houston TX 77032 |
| 16277 | Huntsville Global Entry Enrollment Center - Huntsville International Airport 1000 Glenn Hearn Blvd SW (Airport) Huntsville AL 35824 |
| 14181 | International Falls Global Entry Enrollment Center - 3214 2nd Ave E International Falls MN 56649 |
| 5140 | JFK International Global Entry EC - JFK International Airport Terminal 4, First Floor (Arrivals Level) Jamaica NY 11430 |
| 12781 | Kansas City Enrollment Center - 1 Kansas City Boulevard Suite 30 Arrivals Level Kansas City MO 64153 |
| 5520 | Lansdowne (Thousand Islands Bridge) - 860 Highway 137 Lansdowne ON K0E1L0 |
| 5004 | Laredo Enrollment Center - 400 San Edwardo Laredo TX 780443130 |
| 5360 | Las Vegas Enrollment Center - 5757 Wayne Newton Blvd Terminal 3 Las Vegas NV 89119 |
| 5180 | Los Angeles International Global Entry EC - 11099 S LA CIENEGA BLVD SUITE 155 LOS ANGELES CA 90045 |
| 13621 | Memphis Intl Airport Global Enrollment Center - 2491 Winchester Suite 230 Memphis TN 38116 |
| 5181 | Miami International Airport - 2100 NW 42nd Ave Miami International Airport, Conc. "J" Miami FL 33126 |
| 7740 | Milwaukee Enrollment Center - 4915 S Howell Avenue, 2nd Floor Milwaukee WI 53207 |
| 6840 | Minneapolis - St. Paul Global Entry EC - 4300 Glumack Drive St. Paul MN 55111 |
| 16282 | Mobile Regional Airport Enrollment Center - 8400 Airport Blvd Mobile AL 36608 |
| 16672 | Moline-Quad Cities International Airport - 3300 69th avenue Moline IL 61265 |
| 10260 | Nashville Enrollment Center - Airport Terminal Nashville TN 37214 |
| 16709 | New Jersey Metro Area Mobile EC - 1180 1st Street New Windsor NY 12553 |
| 9740 | New Orleans Enrollment Center - 1 Terminal Drive Kenner LA 70062 |
| 16748 | New York Metro Area Mobile EC - ` ` ` NY 11530 |
| 5444 | Newark Liberty Intl Airport - Newark Liberty International Airport Terminal B - Level 1 Entrance Area Newark NJ 07114 |
| 5161 | Niagara Falls Enrollment Center - 2250 Whirlpool St Niagara Falls NY 14305 |
| 16711 | Niagara Falls Enrollment Center GE Only - 2250 Whirlpool St Niagara Falls NY 14305 |
| 5007 | Nogales, AZ - 200 N Mariposa Road, Suite B700 Nogales AZ 85621 |
| 16555 | Norfolk EC - U.S. Customs House 101 E Main St Norfolk VA 23510 |
| 16771 | Ogdensburg Enrollment Center - GE Only - 104 Bridge Approach Road Ogdensburg NY 13669 |
| 16467 | Omaha, NE Enrollment Center - 3737A Orville Plaza Omaha NE 68110 |
| 16674 | Ontario Intl Airport GE (California) - 2222 International Way International Terminal Ontario CA 91761 |
| 5380 | Orlando International Airport - 10200 Jeff Fuqua Blvd. South Orlando, FL 32827 Orlando FL 32827 |
| 5002 | Otay Mesa Enrollment Center - 9725 Via De La Amistad San Diego CA 92154 |
| 15221 | Pembina Global Entry Enrollment Center - 10980 Interstate 29 N Pembina ND 58271 |
| 11002 | Peoria international airport - 6100 W. Everett M. Dirksen Parkway International Terminal Peoria IL 61607 |
| 5445 | Philadelphia International Airport - PHILADELPHIA INTL AIRPORT TERMINAL A WEST, 3RD FLOOR PHILADELPHIA PA 19153 |
| 7160 | Phoenix Sky Harbor Global Entry Enrollment Center - 3400 E Sky Harbor Blvd, Terminal 4 CBP-Global Enrollment Center Phoenix AZ 85034 |
| 9200 | Pittsburgh International Airport - 1000 Airport Boulevard Ticketing Level Pittsburgh PA 15231 |
| 11841 | Port Clinton, Ohio Enrollment Center - 709 S.E. Catawba Road Port Clinton OH 43452 |
| 5024 | Port Huron Enrollment Center - 2321 Pine Grove Avenue Port Huron MI 48060 |
| 16699 | Port Huron Enrollment Center - RI - - FO MI 48226 |
| 16661 | PortMiami - 1435 North Cruise Blvd Cruise Terminal D Miami FL 33132 |
| 7960 | Portland, OR Enrollment Center - 8337 NE Alderwood Rd U.S. Customs and Border Protection Area Port Portland OR 97220 |
| 14981 | Richmond, VA Enrollment Center - 5707 Huntsman Road, Suite 104 Ivor Massey Building Richmond VA 23250 |
| 11001 | Rockford-Chicago International Airport - 50 Airport Drive Chicago Rockford International Airport Rockford IL 61109 |
| 16475 | SEAFO - Bozeman Airport - 550 Wings Way Belgrade MT 59714 |
| 16488 | SEAFO - Missoula International Airport - 5225 U.S. HIGWAY 10W Missoula MT 59808 |
| 7600 | Salt Lake City International Airport - 3920 West Terminal Dr. #TCBP-1-060.2 Salt Lake City UT 84116 |
| 7520 | San Antonio International Airport - 9800 Airport Boulevard, Suite 1101 San Antonio TX 78216 |
| 16547 | San Diego International Airport EC - 3835 North Harbor Dr Terminal 2 West San Diego CA 92101 |
| 5446 | San Francisco Global Entry Enrollment Center - International Arrival Level San Francisco CA 94128 |
| 5400 | San Juan Global Entry Enrollment Center - Luis Muñoz Marin International Airport (SJU) Indoor Patio (La Placita de Aerostar) Carolina PR 00983 |
| 16717 | San Juan Seaport - #1 La Puntilla Street San Juan PR 00901 |
| 5460 | San Luis Enrollment Center - 1375 South Avenue E SLU II Global Enrollment Center San Luis AZ 85349 |
| 5447 | Sanford Global Entry Enrollment Center - 1100 Red Cleveland Blvd Sanford FL 32773 |
| 5080 | Sault Ste Marie Enrollment Center - 900 W Portage Ave 1st Floor Sault Ste. Marie MI 49783 |
| 5420 | Seatac International Airport Global Entry EC - 17801 International Blvd, SeaTac, WA 98158 US Customs and Border Protection SeaTac WA 98158 |
| 16723 | Shreveport Regional Airport - 5103 Hollywood Ave Shreveport LA 71101 |
| 9040 | Singapore (Singapore, U.S. Embassy) - U.S. Embassy 27 Napier Road Singapore 258508 |
| 16278 | South Bend Enrollment Center - South Bend International Airport 4501 Progress Drive South Bend IN 46628 |
| 16693 | Special Event-Ontario Intl Airport(California)2024 - 2222 INTERNATIONAL WAY, INTERNATIONAL TERMINAL ONTARIO CA 91761 |
| 16463 | Springfield – Branson National Airport, MO - 2300 N Airport Blvd Springfield MO 65802 |
| 12021 | St. Louis Enrollment Center - 4349 WOODSON RD #201 ST. LOUIS MO 63134 |
| 16809 | St. Thomas Cyril E. King Airport Enrollment Center - Lindbergh Bay St. Thomas VI 00802 |
| 16251 | Sweetgrass Global Entry Enrollment Center - Nexus Enrollment Center 39825 Interstate 15 Sweetgrass MT 59484 |
| 5120 | Sweetgrass NEXUS and FAST Enrollment Center - Nexus Enrollment Center 39825 Interstate 15 Sweetgrass MT 59484 |
| 8020 | Tampa Enrollment Center - Tampa International Airport 4100 George J Bean Pkwy Tampa FL 33607 |
| 16248 | Treasure Coast International Airport - 2990 Curtis King Blvd RM 122 Fort Pierce FL 34946 |
| 16271 | Tri-cities Enrollment Center - Tri-Cities Airport 2525 TN-75 Blountville TN 37617 |
| 9240 | Tucson Enrollment Center - 7081 S. Plumer Avenue Tucson AZ 85756 |
| 6480 | U.S. Custom House - Bowling Green - 1 BOWLING GREEN NEW YORK NY 10004 |
| 5060 | Warroad Enrollment Center - 41059 Warroad Enrollment Center State Hwy 313 N Warroad MN 56763 |
| 16837 | Warwick RI Enrollment Center - 300 Jefferson Blvd, Suite 104 Warwick RI 02886 |
| 5142 | Washington Dulles International Global Entry EC - 1 Saarinen Circle Main Terminal, Lower Level, near arrivals door 1 Sterling VA 20166 |
| 8120 | Washington, DC Enrollment Center - 1300 Pennsylvania Avenue NW Washington DC 20229 |
| 9260 | West Palm Beach Enrollment Center - West Palm Beach Enrollment Center 1 East 11th Street, Third Floor Riviera Beach FL 33404 |

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