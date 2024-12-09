package fun

import (
	"fmt"
	"math/rand"
	"strings"
)

// UserAgentRandom generates a random DESKTOP browser user-agent on every requests
func UserAgentRandom() string {
	return uaGens[rand.Intn(len(uaGens))]()
}

// UserAgentRandomMobile generates a random MOBILE browser user-agent on every requests
func UserAgentRandomMobile() string {
	return uaGensMobile[rand.Intn(len(uaGensMobile))]()

}

var uaGens = []func() string{
	genFirefoxUA,
	genChromeUA,
	genEdgeUA,
}

var uaGensMobile = []func() string{
	genMobilePixel7UA,
	genMobilePixel6UA,
	genMobilePixel5UA,
	genMobileNexus10UA,
}

// RandomUserAgent generates a random DESKTOP browser user-agent on every requests

var ffVersions = []float32{
	// NOTE: Only version released after Jun 1, 2022 will be listed.
	// Data source: https://en.wikipedia.org/wiki/Firefox_version_history

	// Firefox 109 was released on January 17, 2023
	109.0,
	110.0,
	111.0,
	112.0,
	113.0,
	113.0,
	113.0,
	113.0,
	113.0,
	114.0,
	115.0,
	116.0,
	117.0,
	118.0,
	119.0,
	120.0,
	121.0,

	// Firefox 122 was released on January 23, 2024
	122.0,
	123.0,
	124.0,
	125.0,
	126.0,
	127.0,
	128.0,
}

var chromeVersions = []string{
	// NOTE: Only version released after Jun 1, 2022 will be listed.
	// Data source: https://chromereleases.googleblog.com/search/label/Stable%20updates

	// https://chromereleases.googleblog.com/2023/01/stable-channel-update-for-desktop.html
	"109.0.5414.74",
	"109.0.5414.75",
	"109.0.5414.87",

	// https://chromereleases.googleblog.com/2023/01/stable-channel-update-for-desktop_24.html
	"109.0.5414.119",
	"109.0.5414.120",

	// https://chromereleases.googleblog.com/2023/02/stable-channel-update-for-desktop.html
	"110.0.5481.77",
	"110.0.5481.78",

	// https://chromereleases.googleblog.com/2023/02/stable-channel-desktop-update.html
	"110.0.5481.96",
	"110.0.5481.97",

	// https://chromereleases.googleblog.com/2023/02/stable-channel-desktop-update_14.html
	"110.0.5481.100",

	// https://chromereleases.googleblog.com/2023/02/stable-channel-desktop-update_16.html
	"110.0.5481.104",

	// https://chromereleases.googleblog.com/2023/02/stable-channel-desktop-update_22.html
	"110.0.5481.177",
	"110.0.5481.178",

	// https://chromereleases.googleblog.com/2023/02/stable-channel-desktop-update_97.html
	"109.0.5414.129",

	// https://chromereleases.googleblog.com/2023/03/stable-channel-update-for-desktop.html
	"111.0.5563.64",
	"111.0.5563.65",

	// https://chromereleases.googleblog.com/2023/03/stable-channel-update-for-desktop_21.html
	"111.0.5563.110",
	"111.0.5563.111",

	// https://chromereleases.googleblog.com/2023/03/stable-channel-update-for-desktop_27.html
	"111.0.5563.146",
	"111.0.5563.147",

	// https://chromereleases.googleblog.com/2023/04/stable-channel-update-for-desktop.html
	"112.0.5615.49",
	"112.0.5615.50",

	// https://chromereleases.googleblog.com/2023/04/stable-channel-update-for-desktop_12.html
	"112.0.5615.86",
	"112.0.5615.87",

	// https://chromereleases.googleblog.com/2023/04/stable-channel-update-for-desktop_14.html
	"112.0.5615.121",

	// https://chromereleases.googleblog.com/2023/04/stable-channel-update-for-desktop_18.html
	"112.0.5615.137",
	"112.0.5615.138",
	"112.0.5615.165",

	// https://chromereleases.googleblog.com/2023/05/stable-channel-update-for-desktop.html
	"113.0.5672.63",
	"113.0.5672.64",

	// https://chromereleases.googleblog.com/2023/05/stable-channel-update-for-desktop_8.html
	"113.0.5672.92",
	"113.0.5672.93",

	// https://chromereleases.googleblog.com/2023/06/stable-channel-update-for-desktop_26.html
	"114.0.5735.198",
	"114.0.5735.199",

	// https://chromereleases.googleblog.com/2023/07/stable-channel-update-for-desktop_25.html
	"115.0.5790.110",

	// https://chromereleases.googleblog.com/2023/08/stable-channel-update-for-desktop_29.html
	"116.0.5845.140",
	"116.0.5845.141",

	// https://chromereleases.googleblog.com/2023/09/extended-stable-channel-update-for.html
	"116.0.5845.228",

	// https://chromereleases.googleblog.com/2023/10/stable-channel-update-for-desktop_31.html
	"119.0.6045.105",
	"119.0.6045.106",

	// https://chromereleases.googleblog.com/2023/11/extended-stable-channel-update-for_28.html
	"118.0.5993.159",

	// https://chromereleases.googleblog.com/2023/12/stable-channel-update-for-desktop_20.html
	"120.0.6099.129",
	"120.0.6099.130",
}

var edgeVersions = []string{
	// NOTE: Only version released after Jun 1, 2022 will be listed.
	// Data source: https://learn.microsoft.com/en-us/deployedge/microsoft-edge-release-schedule

	// 2023
	"109.0.0.0,109.0.1518.49",
	"110.0.0.0,110.0.1587.41",
	"111.0.0.0,111.0.1661.41",
	"112.0.0.0,112.0.1722.34",
	"113.0.0.0,113.0.1774.3",
	"114.0.0.0,114.0.1823.37",
	"115.0.0.0,115.0.1901.183",
	"116.0.0.0,116.0.1938.54",
	"117.0.0.0,117.0.2045.31",
	"118.0.0.0,118.0.2088.46",
	"119.0.0.0,119.0.2151.44",
	"120.0.0.0,120.0.2210.61",

	// 2024
	"121.0.0.0,121.0.2277.83",
	"122.0.0.0,122.0.2365.52",
	"123.0.0.0,123.0.2420.53",
	"125.0.0.0,124.0.2478.51",
	"125.0.0.0,125.0.2535.51",
	"126.0.0.0,126.0.2592.56",
	"127.0.0.0,127.0.2651.74",
	"128.0.0.0,128.0.2739.42",
}

var pixel7AndroidVersions = []string{
	// Data source:
	// - https://developer.android.com/about/versions
	// - https://source.android.com/docs/setup/about/build-numbers#source-code-tags-and-builds
	"13",
}

var pixel6AndroidVersions = []string{
	// Data source:
	// - https://developer.android.com/about/versions
	// - https://source.android.com/docs/setup/about/build-numbers#source-code-tags-and-builds
	"12",
	"13",
}

var pixel5AndroidVersions = []string{
	// Data source:
	// - https://developer.android.com/about/versions
	// - https://source.android.com/docs/setup/about/build-numbers#source-code-tags-and-builds
	"11",
	"12",
	"13",
}

var nexus10AndroidVersions = []string{
	// Data source:
	// - https://developer.android.com/about/versions
	// - https://source.android.com/docs/setup/about/build-numbers#source-code-tags-and-builds
	"4.4.2",
	"4.4.4",
	"5.0",
	"5.0.1",
	"5.0.2",
	"5.1",
	"5.1.1",
}

var nexus10Builds = []string{
	// Data source: https://source.android.com/docs/setup/about/build-numbers#source-code-tags-and-builds

	"LMY49M", // android-5.1.1_r38 (Lollipop)
	"LMY49J", // android-5.1.1_r37 (Lollipop)
	"LMY49I", // android-5.1.1_r36 (Lollipop)
	"LMY49H", // android-5.1.1_r35 (Lollipop)
	"LMY49G", // android-5.1.1_r34 (Lollipop)
	"LMY49F", // android-5.1.1_r33 (Lollipop)
	"LMY48Z", // android-5.1.1_r30 (Lollipop)
	"LMY48X", // android-5.1.1_r25 (Lollipop)
	"LMY48T", // android-5.1.1_r19 (Lollipop)
	"LMY48M", // android-5.1.1_r14 (Lollipop)
	"LMY48I", // android-5.1.1_r9 (Lollipop)
	"LMY47V", // android-5.1.1_r1 (Lollipop)
	"LMY47D", // android-5.1.0_r1 (Lollipop)
	"LRX22G", // android-5.0.2_r1 (Lollipop)
	"LRX22C", // android-5.0.1_r1 (Lollipop)
	"LRX21P", // android-5.0.0_r4.0.1 (Lollipop)
	"KTU84P", // android-4.4.4_r1 (KitKat)
	"KTU84L", // android-4.4.3_r1 (KitKat)
	"KOT49H", // android-4.4.2_r1 (KitKat)
	"KOT49E", // android-4.4.1_r1 (KitKat)
	"KRT16S", // android-4.4_r1.2 (KitKat)
	"JWR66Y", // android-4.3_r1.1 (Jelly Bean)
	"JWR66V", // android-4.3_r1 (Jelly Bean)
	"JWR66N", // android-4.3_r0.9.1 (Jelly Bean)
	"JDQ39 ", // android-4.2.2_r1 (Jelly Bean)
	"JOP40F", // android-4.2.1_r1.1 (Jelly Bean)
	"JOP40D", // android-4.2.1_r1 (Jelly Bean)
	"JOP40C", // android-4.2_r1 (Jelly Bean)
}

var osStrings = []string{
	// MacOS - Catalina
	"Macintosh; Intel Mac OS X 10_15",
	"Macintosh; Intel Mac OS X 10_15_1",
	"Macintosh; Intel Mac OS X 10_15_2",
	"Macintosh; Intel Mac OS X 10_15_3",
	"Macintosh; Intel Mac OS X 10_15_4",
	"Macintosh; Intel Mac OS X 10_15_5",
	"Macintosh; Intel Mac OS X 10_15_6",
	"Macintosh; Intel Mac OS X 10_15_7",

	// MacOS - Big Sur
	"Macintosh; Intel Mac OS X 11_0",
	"Macintosh; Intel Mac OS X 11_0_1",
	"Macintosh; Intel Mac OS X 11_1",
	"Macintosh; Intel Mac OS X 11_2",
	"Macintosh; Intel Mac OS X 11_2_1",
	"Macintosh; Intel Mac OS X 11_2_2",
	"Macintosh; Intel Mac OS X 11_2_3",
	"Macintosh; Intel Mac OS X 11_3",
	"Macintosh; Intel Mac OS X 11_3_1",
	"Macintosh; Intel Mac OS X 11_4",
	"Macintosh; Intel Mac OS X 11_5",
	"Macintosh; Intel Mac OS X 11_5_1",
	"Macintosh; Intel Mac OS X 11_5_2",
	"Macintosh; Intel Mac OS X 11_6",
	"Macintosh; Intel Mac OS X 11_6_1",
	"Macintosh; Intel Mac OS X 11_6_2",
	"Macintosh; Intel Mac OS X 11_6_3",
	"Macintosh; Intel Mac OS X 11_6_4",
	"Macintosh; Intel Mac OS X 11_6_5",
	"Macintosh; Intel Mac OS X 11_6_6",
	"Macintosh; Intel Mac OS X 11_6_7",
	"Macintosh; Intel Mac OS X 11_6_8",
	"Macintosh; Intel Mac OS X 11_7",
	"Macintosh; Intel Mac OS X 11_7_1",
	"Macintosh; Intel Mac OS X 11_7_2",
	"Macintosh; Intel Mac OS X 11_7_3",
	"Macintosh; Intel Mac OS X 11_7_4",
	"Macintosh; Intel Mac OS X 11_7_5",
	"Macintosh; Intel Mac OS X 11_7_6",

	// MacOS - Monterey
	"Macintosh; Intel Mac OS X 12_0",
	"Macintosh; Intel Mac OS X 12_0_1",
	"Macintosh; Intel Mac OS X 12_1",
	"Macintosh; Intel Mac OS X 12_2",
	"Macintosh; Intel Mac OS X 12_2_1",
	"Macintosh; Intel Mac OS X 12_3",
	"Macintosh; Intel Mac OS X 12_3_1",
	"Macintosh; Intel Mac OS X 12_4",
	"Macintosh; Intel Mac OS X 12_5",
	"Macintosh; Intel Mac OS X 12_5_1",
	"Macintosh; Intel Mac OS X 12_6",
	"Macintosh; Intel Mac OS X 12_6_1",
	"Macintosh; Intel Mac OS X 12_6_2",
	"Macintosh; Intel Mac OS X 12_6_3",
	"Macintosh; Intel Mac OS X 12_6_4",
	"Macintosh; Intel Mac OS X 12_6_5",

	// MacOS - Ventura
	"Macintosh; Intel Mac OS X 13_0",
	"Macintosh; Intel Mac OS X 13_0_1",
	"Macintosh; Intel Mac OS X 13_1",
	"Macintosh; Intel Mac OS X 13_2",
	"Macintosh; Intel Mac OS X 13_2_1",
	"Macintosh; Intel Mac OS X 13_3",
	"Macintosh; Intel Mac OS X 13_3_1",

	// Windows
	"Windows NT 10.0; Win64; x64",
	"Windows NT 5.1",
	"Windows NT 6.1; WOW64",
	"Windows NT 6.1; Win64; x64",

	// Linux
	"X11; Linux x86_64",
}

// Generates Firefox Browser User-Agent (Desktop)
//
// -> "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:87.0) Gecko/20100101 Firefox/87.0"
func genFirefoxUA() string {
	version := ffVersions[rand.Intn(len(ffVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s; rv:%.1f) Gecko/20100101 Firefox/%.1f", os, version, version)
}

// Generates Chrome Browser User-Agent (Desktop)
//
// -> "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36"
func genChromeUA() string {
	version := chromeVersions[rand.Intn(len(chromeVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", os, version)
}

// Generates Microsoft Edge User-Agent (Desktop)
//
// -> "User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36 Edg/90.0.818.39"
func genEdgeUA() string {
	version := edgeVersions[rand.Intn(len(edgeVersions))]
	chromeVersion := strings.Split(version, ",")[0]
	edgeVersion := strings.Split(version, ",")[1]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 Edg/%s", os, chromeVersion, edgeVersion)
}

// Generates Pixel 7 Browser User-Agent (Mobile)
//
// -> Mozilla/5.0 (Linux; Android 13; Pixel 7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36
func genMobilePixel7UA() string {
	android := pixel7AndroidVersions[rand.Intn(len(pixel7AndroidVersions))]
	chrome := chromeVersions[rand.Intn(len(chromeVersions))]
	return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; Pixel 7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", android, chrome)
}

// Generates Pixel 6 Browser User-Agent (Mobile)
//
// -> "Mozilla/5.0 (Linux; Android 13; Pixel 6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36"
func genMobilePixel6UA() string {
	android := pixel6AndroidVersions[rand.Intn(len(pixel6AndroidVersions))]
	chrome := chromeVersions[rand.Intn(len(chromeVersions))]
	return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; Pixel 6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", android, chrome)
}

// Generates Pixel 5 Browser User-Agent (Mobile)
//
// -> "Mozilla/5.0 (Linux; Android 13; Pixel 5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36"
func genMobilePixel5UA() string {
	android := pixel5AndroidVersions[rand.Intn(len(pixel5AndroidVersions))]
	chrome := chromeVersions[rand.Intn(len(chromeVersions))]
	return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; Pixel 5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", android, chrome)
}

// Generates Nexus 10 Browser User-Agent (Mobile)
//
// -> "Mozilla/5.0 (Linux; Android 5.1.1; Nexus 10 Build/LMY48T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.91 Safari/537.36"
func genMobileNexus10UA() string {
	build := nexus10Builds[rand.Intn(len(nexus10Builds))]
	android := nexus10AndroidVersions[rand.Intn(len(nexus10AndroidVersions))]
	chrome := chromeVersions[rand.Intn(len(chromeVersions))]
	return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; Nexus 10 Build/%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", android, build, chrome)
}
