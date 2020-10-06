# Golang `Timezone` Time-Zone based research

## References (Golang Timezone)

* [datetime - How to get timezone from country - Stack Overflow](https://stackoverflow.com/questions/51346476/how-to-get-timezone-from-country)
* [go - How to get equivalent time in another timezone - Stack Overflow](https://stackoverflow.com/questions/42940065/how-to-get-equivalent-time-in-another-timezone)
* [timezone - Get a list of valid time zones in Go - Stack Overflow](https://stackoverflow.com/questions/40120056/get-a-list-of-valid-time-zones-in-go)
* [go - How to get the current timestamp in other timezones in Golang? - Stack Overflow](https://stackoverflow.com/questions/27991671/how-to-get-the-current-timestamp-in-other-timezones-in-golang)
* [Golang : Display list of time zones with GMT](https://www.socketloop.com/tutorials/golang-display-list-of-timezones-with-gmt)
* [Timezone sample code that works](https://play.golang.org/p/GenIhfkGLQw)

## Code Samples (Golang Timezone)

https://github.com/aukgit/pkg-testing/blob/master/cmd/timezone/timezone.go

```go
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp/syntax"
	"runtime"
	"strings"
	"time"
	"unicode/utf8"
	// "golang.org/x/sys/windows/registry"
)

// zoneDirs adapted from https://golang.org/src/time/zoneinfo_unix.go

// https://golang.org/doc/install/source#environment
// list of available GOOS as of 10th Feb 2017
// android, darwin, dragonfly, freebsd, linux, netbsd, openbsd, plan9, solaris,windows

var zoneDirs = map[string]string{
	"android":   "/system/usr/share/zoneinfo/",
	"darwin":    "/usr/share/zoneinfo/",
	"dragonfly": "/usr/share/zoneinfo/",
	"freebsd":   "/usr/share/zoneinfo/",
	"linux":     "/usr/share/zoneinfo/",
	"netbsd":    "/usr/share/zoneinfo/",
	"openbsd":   "/usr/share/zoneinfo/",
	// "plan9":"/adm/timezone/", -- no way to test this platform
	"solaris": "/usr/share/lib/zoneinfo/",
	"windows": `SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones\`,
}

var zoneDir string

var timeZones []string

// InSlice ... check if an element is inside a slice
func InSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// ReadTZFile ... read timezone file and append into timeZones slice
func ReadTZFile(path string) {
	files, _ := ioutil.ReadDir(zoneDir + path)
	for _, f := range files {
		if f.Name() != strings.ToUpper(f.Name()[:1])+f.Name()[1:] {
			continue
		}
		if f.IsDir() {
			ReadTZFile(path + "/" + f.Name())
		} else {
			tz := (path + "/" + f.Name())[1:]
			// check if tz is already in timeZones slice
			// append if not
			if !InSlice(tz, timeZones) { // need a more efficient method...

				// convert string to rune
				tzRune, _ := utf8.DecodeRuneInString(tz[:1])

				if syntax.IsWordChar(tzRune) { // filter out entry that does not start with A-Za-z such as +VERSION
					timeZones = append(timeZones, tz)
				}
			}
		}
	}

}

func ListTimeZones() {
	if runtime.GOOS == "nacl" || runtime.GOOS == "" {
		fmt.Println("Unsupported platform")
		os.Exit(0)
	}

	// detect OS
	fmt.Println("Time zones available for : ", runtime.GOOS)
	fmt.Println("------------------------")

	fmt.Println("Retrieving time zones from : ", zoneDirs[runtime.GOOS])

	if runtime.GOOS != "windows" {
		for _, zoneDir = range zoneDirs {
			ReadTZFile("")
		}
	} else { // let's handle Windows
		// if you're building this on darwin/linux
		// chances are you will encounter
		// undefined: registry in registry.OpenKey error message
		// uncomment below if compiling on Windows platform

		// k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones`, registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE)

		// if err != nil {
		// fmt.Println(err)
		// }
		// defer k.Close()

		// names, err := k.ReadSubKeyNames(-1)
		// if err != nil {
		// fmt.Println(err)
		// }

		// fmt.Println("Number of timezones : ", len(names))
		// for i := 0; i <= len(names)-1; i++ {
		// check if tz is already in timeZones slice
		// append if not
		// if !InSlice(names[i], timeZones) { // need a more efficient method...
		//  timeZones = append(timeZones, names[i])
		// }
		// }

		// UPDATE : Reading from registry is not reliable
		// better to parse output result by "tzutil /g" command
		// REMEMBER : There is no time difference between Coordinated Universal Time and Greenwich Mean Time ....
		cmd := exec.Command("tzutil", "/l")

		data, err := cmd.Output()

		if err != nil {
			panic(err)
		}

		fmt.Println("UTC is the same as GMT")
		fmt.Println("There is no time difference between Coordinated Universal Time and Greenwich Mean Time ....")
		GMTed := bytes.Replace(data, []byte("UTC"), []byte("GMT"), -1)

		fmt.Println(string(GMTed))

	}

	now := time.Now()

	for _, v := range timeZones {

		if runtime.GOOS != "windows" {

			location, err := time.LoadLocation(v)
			if err != nil {
				fmt.Println(err)
			}

			// extract the GMT
			t := now.In(location)
			t1 := fmt.Sprintf("%s", t.Format(time.RFC822Z))
			tArray := strings.Fields(t1)
			gmtTime := strings.Join(tArray[4:], "")
			hours := gmtTime[0:3]
			minutes := gmtTime[3:]

			gmt := "GMT" + fmt.Sprintf("%s:%s", hours, minutes)
			fmt.Println(gmt + " " + v)

		} else {
			fmt.Println(v)
		}

	}
	fmt.Println("Total timezone ids : ", len(timeZones))
}

func main() {
	ListTimeZones()
}
```

### Sample outputs from Windows machine

```
GOROOT=C:\Go #gosetup
GOPATH=D:\github\go-workspace #gosetup
C:\Go\bin\go.exe build -o C:\Users\Administrator\AppData\Local\Temp\___go_build_github_com_aukgit_pkgtesting_cmd_timezone.exe github.com/aukgit/pkgtesting/cmd/timezone #gosetup
C:\Users\Administrator\AppData\Local\Temp\___go_build_github_com_aukgit_pkgtesting_cmd_timezone.exe #gosetup
Time zones available for :  windows
------------------------
Retrieving time zones from :  SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones\
UTC is the same as GMT
There is no time difference between Coordinated Universal Time and Greenwich Mean Time ....
(GMT-12:00) International Date Line West 
Dateline Standard Time

(GMT-11:00) Coordinated Universal Time-11 
GMT-11

(GMT-10:00) Aleutian Islands 
Aleutian Standard Time

(GMT-10:00) Hawaii 
Hawaiian Standard Time

(GMT-09:30) Marquesas Islands 
Marquesas Standard Time

(GMT-09:00) Alaska 
Alaskan Standard Time

(GMT-09:00) Coordinated Universal Time-09 
GMT-09

(GMT-07:00) Yukon 
Yukon Standard Time

(GMT-08:00) Baja California 
Pacific Standard Time (Mexico)

(GMT-08:00) Coordinated Universal Time-08 
GMT-08

(GMT-08:00) Pacific Time (US & Canada) 
Pacific Standard Time

(GMT-07:00) Arizona 
US Mountain Standard Time

(GMT-07:00) Chihuahua, La Paz, Mazatlan 
Mountain Standard Time (Mexico)

(GMT-07:00) Mountain Time (US & Canada) 
Mountain Standard Time

(GMT-06:00) Central America 
Central America Standard Time

(GMT-06:00) Central Time (US & Canada) 
Central Standard Time

(GMT-06:00) Easter Island 
Easter Island Standard Time

(GMT-06:00) Guadalajara, Mexico City, Monterrey 
Central Standard Time (Mexico)

(GMT-06:00) Saskatchewan 
Canada Central Standard Time

(GMT-05:00) Bogota, Lima, Quito, Rio Branco 
SA Pacific Standard Time

(GMT-05:00) Chetumal 
Eastern Standard Time (Mexico)

(GMT-05:00) Eastern Time (US & Canada) 
Eastern Standard Time

(GMT-05:00) Haiti 
Haiti Standard Time

(GMT-05:00) Havana 
Cuba Standard Time

(GMT-05:00) Indiana (East) 
US Eastern Standard Time

(GMT-05:00) Turks and Caicos 
Turks And Caicos Standard Time

(GMT-04:00) Asuncion 
Paraguay Standard Time

(GMT-04:00) Atlantic Time (Canada) 
Atlantic Standard Time

(GMT-04:00) Caracas 
Venezuela Standard Time

(GMT-04:00) Cuiaba 
Central Brazilian Standard Time

(GMT-04:00) Georgetown, La Paz, Manaus, San Juan 
SA Western Standard Time

(GMT-04:00) Santiago 
Pacific SA Standard Time

(GMT-03:30) Newfoundland 
Newfoundland Standard Time

(GMT-03:00) Araguaina 
Tocantins Standard Time

(GMT-03:00) Brasilia 
E. South America Standard Time

(GMT-03:00) Cayenne, Fortaleza 
SA Eastern Standard Time

(GMT-03:00) City of Buenos Aires 
Argentina Standard Time

(GMT-03:00) Greenland 
Greenland Standard Time

(GMT-03:00) Montevideo 
Montevideo Standard Time

(GMT-03:00) Punta Arenas 
Magallanes Standard Time

(GMT-03:00) Saint Pierre and Miquelon 
Saint Pierre Standard Time

(GMT-03:00) Salvador 
Bahia Standard Time

(GMT-02:00) Coordinated Universal Time-02 
GMT-02

(GMT-01:00) Azores 
Azores Standard Time

(GMT-01:00) Cabo Verde Is. 
Cape Verde Standard Time

(GMT) Coordinated Universal Time 
GMT

(GMT+00:00) Dublin, Edinburgh, Lisbon, London 
GMT Standard Time

(GMT+00:00) Monrovia, Reykjavik 
Greenwich Standard Time

(GMT+00:00) Sao Tome 
Sao Tome Standard Time

(GMT+01:00) Casablanca 
Morocco Standard Time

(GMT+01:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna 
W. Europe Standard Time

(GMT+01:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague 
Central Europe Standard Time

(GMT+01:00) Brussels, Copenhagen, Madrid, Paris 
Romance Standard Time

(GMT+01:00) Sarajevo, Skopje, Warsaw, Zagreb 
Central European Standard Time

(GMT+01:00) West Central Africa 
W. Central Africa Standard Time

(GMT+02:00) Amman 
Jordan Standard Time

(GMT+02:00) Athens, Bucharest 
GTB Standard Time

(GMT+02:00) Beirut 
Middle East Standard Time

(GMT+02:00) Cairo 
Egypt Standard Time

(GMT+02:00) Chisinau 
E. Europe Standard Time

(GMT+02:00) Damascus 
Syria Standard Time

(GMT+02:00) Gaza, Hebron 
West Bank Standard Time

(GMT+02:00) Harare, Pretoria 
South Africa Standard Time

(GMT+02:00) Helsinki, Kyiv, Riga, Sofia, Tallinn, Vilnius 
FLE Standard Time

(GMT+02:00) Jerusalem 
Israel Standard Time

(GMT+02:00) Kaliningrad 
Kaliningrad Standard Time

(GMT+02:00) Khartoum 
Sudan Standard Time

(GMT+02:00) Tripoli 
Libya Standard Time

(GMT+02:00) Windhoek 
Namibia Standard Time

(GMT+03:00) Baghdad 
Arabic Standard Time

(GMT+03:00) Istanbul 
Turkey Standard Time

(GMT+03:00) Kuwait, Riyadh 
Arab Standard Time

(GMT+03:00) Minsk 
Belarus Standard Time

(GMT+03:00) Moscow, St. Petersburg 
Russian Standard Time

(GMT+03:00) Nairobi 
E. Africa Standard Time

(GMT+03:30) Tehran 
Iran Standard Time

(GMT+04:00) Abu Dhabi, Muscat 
Arabian Standard Time

(GMT+04:00) Astrakhan, Ulyanovsk 
Astrakhan Standard Time

(GMT+04:00) Baku 
Azerbaijan Standard Time

(GMT+04:00) Izhevsk, Samara 
Russia Time Zone 3

(GMT+04:00) Port Louis 
Mauritius Standard Time

(GMT+04:00) Saratov 
Saratov Standard Time

(GMT+04:00) Tbilisi 
Georgian Standard Time

(GMT+04:00) Volgograd 
Volgograd Standard Time

(GMT+04:00) Yerevan 
Caucasus Standard Time

(GMT+04:30) Kabul 
Afghanistan Standard Time

(GMT+05:00) Ashgabat, Tashkent 
West Asia Standard Time

(GMT+05:00) Ekaterinburg 
Ekaterinburg Standard Time

(GMT+05:00) Islamabad, Karachi 
Pakistan Standard Time

(GMT+05:00) Qyzylorda 
Qyzylorda Standard Time

(GMT+05:30) Chennai, Kolkata, Mumbai, New Delhi 
India Standard Time

(GMT+05:30) Sri Jayawardenepura 
Sri Lanka Standard Time

(GMT+05:45) Kathmandu 
Nepal Standard Time

(GMT+06:00) Astana 
Central Asia Standard Time

(GMT+06:00) Dhaka 
Bangladesh Standard Time

(GMT+06:00) Omsk 
Omsk Standard Time

(GMT+06:30) Yangon (Rangoon) 
Myanmar Standard Time

(GMT+07:00) Bangkok, Hanoi, Jakarta 
SE Asia Standard Time

(GMT+07:00) Barnaul, Gorno-Altaysk 
Altai Standard Time

(GMT+07:00) Hovd 
W. Mongolia Standard Time

(GMT+07:00) Krasnoyarsk 
North Asia Standard Time

(GMT+07:00) Novosibirsk 
N. Central Asia Standard Time

(GMT+07:00) Tomsk 
Tomsk Standard Time

(GMT+08:00) Beijing, Chongqing, Hong Kong, Urumqi 
China Standard Time

(GMT+08:00) Irkutsk 
North Asia East Standard Time

(GMT+08:00) Kuala Lumpur, Singapore 
Singapore Standard Time

(GMT+08:00) Perth 
W. Australia Standard Time

(GMT+08:00) Taipei 
Taipei Standard Time

(GMT+08:00) Ulaanbaatar 
Ulaanbaatar Standard Time

(GMT+08:45) Eucla 
Aus Central W. Standard Time

(GMT+09:00) Chita 
Transbaikal Standard Time

(GMT+09:00) Osaka, Sapporo, Tokyo 
Tokyo Standard Time

(GMT+09:00) Pyongyang 
North Korea Standard Time

(GMT+09:00) Seoul 
Korea Standard Time

(GMT+09:00) Yakutsk 
Yakutsk Standard Time

(GMT+09:30) Adelaide 
Cen. Australia Standard Time

(GMT+09:30) Darwin 
AUS Central Standard Time

(GMT+10:00) Brisbane 
E. Australia Standard Time

(GMT+10:00) Canberra, Melbourne, Sydney 
AUS Eastern Standard Time

(GMT+10:00) Guam, Port Moresby 
West Pacific Standard Time

(GMT+10:00) Hobart 
Tasmania Standard Time

(GMT+10:00) Vladivostok 
Vladivostok Standard Time

(GMT+10:30) Lord Howe Island 
Lord Howe Standard Time

(GMT+11:00) Bougainville Island 
Bougainville Standard Time

(GMT+11:00) Chokurdakh 
Russia Time Zone 10

(GMT+11:00) Magadan 
Magadan Standard Time

(GMT+11:00) Norfolk Island 
Norfolk Standard Time

(GMT+11:00) Sakhalin 
Sakhalin Standard Time

(GMT+11:00) Solomon Is., New Caledonia 
Central Pacific Standard Time

(GMT+12:00) Anadyr, Petropavlovsk-Kamchatsky 
Russia Time Zone 11

(GMT+12:00) Auckland, Wellington 
New Zealand Standard Time

(GMT+12:00) Coordinated Universal Time+12 
GMT+12

(GMT+12:00) Fiji 
Fiji Standard Time

(GMT+12:45) Chatham Islands 
Chatham Islands Standard Time

(GMT+13:00) Coordinated Universal Time+13 
GMT+13

(GMT+13:00) Nuku'alofa 
Tonga Standard Time

(GMT+13:00) Samoa 
Samoa Standard Time

(GMT+14:00) Kiritimati Island 
Line Islands Standard Time


Total timezone ids :  0

Process finished with exit code 0

```