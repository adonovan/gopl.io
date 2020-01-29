// 7.6. Sorting with sort.Interface 부분이다.

/*
1) 일단 Track 이라는 구조체 타입을 선언하고, 초기값을 먹여서 tracks 를 생성해준다.
2) length(), printTracks() 는 파싱, 출력을 위한 간단한 유틸 함수들이다.

1) byArtist, byYear 와 같은 식으로 []*Track 을 타입으로 선언하고
2) 각 타입의 메쏘드로 Len(), Less(), Swap() 을 구현해주면
3) sort.Sort() 에 파라미터로 넘길 수가 있다.
	-> data 에 Len(), Less(), Swap() 이 구현되어 있어야 함을 알 수 있다.

	// Sort sorts data.
	// It makes one call to data.Len to determine n, and O(n*log(n)) calls to
	// data.Less and data.Swap. The sort is not guaranteed to be stable.
	func Sort(data Interface) {
		n := data.Len()
		quickSort(data, 0, n, maxDepth(n))
	}

	func quickSort(data Interface, a, b, maxDepth int) {
	for b-a > 12 { // Use ShellSort for slices <= 12 elements
		if maxDepth == 0 {
			heapSort(data, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivot(data, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSort(data, a, mlo, maxDepth)
			a = mhi // i.e., quickSort(data, mhi, b)
		} else {
			quickSort(data, mhi, b, maxDepth)
			b = mlo // i.e., quickSort(data, a, mlo)
		}
	}
	if b-a > 1 {
		// Do ShellSort pass with gap 6
		// It could be written in this simplified form cause b-a <= 12
		for i := a + 6; i < b; i++ {
			if data.Less(i, i-6) {
				data.Swap(i, i-6)
			}
		}
		insertionSort(data, a, b)
	}
}

4) customSort 라는 녀석은 Less() 함수,
	정렬의 핵심 기준인 이 녀석을 실제 사용하는 타이밍에 쓸 수 있게 해주는 구현이다.

5) init() 함수에서는 sort 의 함수들을 간단히 보여준다.
	- sort.IntsAreSorted()
	- sort.Reverse()
	- sort.IntSlice()
		a) 이건 함수가 아니다. 타입캐스팅이다. IntSlice() 는 Len(), Less(), Swap() 등의 메쏘드가 있다.
		b) 이 덕분에 sort.Interface 를 파라미터로 가지는 sort.Reverse() 에 넣을 수 있게 된다.

		type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
		Swap(i, j int)

		func Reverse(data Interface) Interface
}

*/

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

//!+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//!-main

//!+printTracks
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

//!-printTracks

//!+artistcode
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-artistcode

//!+yearcode
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-yearcode

func main() {
	fmt.Println("byArtist:")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	fmt.Println("\nReverse(byArtist):")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	fmt.Println("\nbyYear:")
	sort.Sort(byYear(tracks))
	printTracks(tracks)

	fmt.Println("\nCustom:")
	//!+customcall
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	//!-customcall
	printTracks(tracks)
}

/*
//!+artistoutput
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go Ahead    Alicia Keys     As I Am            2007  4m36s
Go          Delilah         From the Roots Up  2012  3m38s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s
Go          Moby            Moby               1992  3m37s
//!-artistoutput

//!+artistrevoutput
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go          Moby            Moby               1992  3m37s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s
Go          Delilah         From the Roots Up  2012  3m38s
Go Ahead    Alicia Keys     As I Am            2007  4m36s
//!-artistrevoutput

//!+yearoutput
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go          Moby            Moby               1992  3m37s
Go Ahead    Alicia Keys     As I Am            2007  4m36s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s
Go          Delilah         From the Roots Up  2012  3m38s
//!-yearoutput

//!+customout
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go          Moby            Moby               1992  3m37s
Go          Delilah         From the Roots Up  2012  3m38s
Go Ahead    Alicia Keys     As I Am            2007  4m36s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s
//!-customout
*/

//!+customcode
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

//!-customcode

func init() {
	//!+ints
	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values)) // "false"
	sort.Ints(values)
	fmt.Println(values)                     // "[1 1 3 4]"
	fmt.Println(sort.IntsAreSorted(values)) // "true"
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)                     // "[4 3 1 1]"
	fmt.Println(sort.IntsAreSorted(values)) // "false"
	//!-ints
}

/*
7.8 대부분의 GUI는 상태가 있는 다계층 정렬 테이블 위셋을 제공한다.
기본 정렬 키는 가장 최근에 클릭된 열의 제목이며, 보조 정렬 키는 두 번째로 최근에 클릭된 열의 제목과 같은 식이다.
이런 테이블에 사용하기 위한 sort.Interface 의 구현을 정의하라. 이 접근 방법과 sort.Stable 을 반복 호출해 정렬하는 방법을 비교하라.

7.9 html/template 패키지 (4.6절)를 이용해 트랙을 HTML 테이블로 표시하는 함수로 printTracks 를 대체하라.
이전 예제에 이 방법을 적용하고 열 제목을 클릭할 때 HTML 요청을 생성해 테이블을 정렬하라.

7.10 sort.Interface 타입은 다른 용도로도 쓸 수 있다.
IsPalindrome(s sort.Intepface) bool 함수를 작성해 시퀀스 s가 회문인지 여부, 다시 말해 순서를 바꿔도 시퀀스가 변경되지 않는지를 보고하라.
인덱스 i와 j 의 원소는 !s.Less(i, j) && !s.Less(j, i)일 때 같다고 간추할 수 있다.

*/
