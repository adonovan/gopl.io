package main

import (
	"strings"
	"testing"
)

var testString = "Skip to content Search or jump to…  Pull requests Issues Marketplace Explore   @Limza  Repositories Find a repository… Limza / goserver gopl-study / gopl.io Limza / gopl.io Limza / WINR Your teams Find a team… EpicGames/developers Dashboard  You’ve been added to the EpicGames organization! Here are some quick tips for a first-time organization member.  Use the switch context button in the upper left corner of this page to switch between your personal context (Limza) and organizations you are a member of. After you switch contexts you’ll see an organization-focused dashboard that lists out organization repositories and activities. Discover interesting projects and people to populate your personal news feed. Your news feed helps you keep up with recent activity on repositories you watch and people you follow.   ProTip! The feed shows you events from people you follow and repositories you watch. Subscribe to your news feed © 2019 GitHub, Inc. Blog About Shop Contact GitHub Pricing API Training Status Security Terms Privacy Help See what launched at GitHub Universe Say hello to our new GitHub for mobile experience, get more from GitHub Actions and Packages, and learn how we’re shaping the future of software, together. Sponsors Mona GitHub Sponsors is out of beta We're now generally available in 30 countries, with more coming soon. advanced security Securing software together Introducing new ways to identify and prevent security vulnerabilities across your code base. Explore repositories pierrec/lz4 LZ4 compression and decompression in pure Go   Go  376 go-playground/log Simple, configurable and scalable Structured Logging for Go.   Go  270 labstack/gommon Common packages for Go   Go  317 Explore more → "
var testSlice = strings.Split(testString, " ")

func Join1() {
	var s, sep string
	for _, v := range testSlice {
		s += sep + v
		sep = " "
	}
}

func Join2() {
	strings.Join(testSlice, " ")
}

func Benchmark_join1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join1()
	}
}

func Benchmark_join2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join2()
	}
}
