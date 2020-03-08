package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

func dup1() {
	counts := make(map[string]int)
	inputs := bufio.NewScanner(os.Stdin)
	for inputs.Scan() {
		counts[inputs.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	fmt.Println(files)
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIdx = 0
	blackIdx = 1
)

func giff() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}

		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}

	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIdx)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func fetch_url() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

func fetch_url_parallel() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // create a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func web_server() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.path=%q\n", r.URL.Path)
}

func test_iota() {
	type WeekDay int
	const (
		Sunday WeekDay = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	type Flags uint
	const (
		flag1 Flags = 1 << iota
		flag2
		flag3
		flag4
		flag5
	)

	fmt.Println(flag1)
	fmt.Println(flag2)
	fmt.Println(flag3)
	fmt.Println(flag4)
	fmt.Println(flag5)

	fmt.Printf("%b\n", flag1)
	fmt.Printf("%b\n", flag2)
	fmt.Printf("%b\n", flag3)
	fmt.Printf("%b\n", flag4)
	fmt.Printf("%b\n", flag5)

	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)

	const (
		pi = math.Pi
	)

	pp := math.Pi
	var fpp float32 = math.Pi
	var lfpp float64 = math.Pi
	fmt.Println(pi)
	fmt.Println(pp)
	fmt.Println(fpp)
	fmt.Println(lfpp)
}

func test_array() {
	q := [...]int{1, 2, 3}
	for i, v := range q {
		fmt.Println(i, v)
	}
	fmt.Println("%T\n", q)
	q = [3]int{2, 3, 4}
	fmt.Println("%T\n", q)

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{EUR: "RR", GBP: "GG", RMB: "RR"}
	fmt.Println(symbol[RMB], len(symbol))

	c1 := sha256.Sum256([]byte("xx"))
	c2 := sha256.Sum256([]byte("xX"))
	fmt.Printf("%x\n%x\n%T\n", c1, c2, c1)
	zero(&c1)
	fmt.Printf("%x\n%x\n%T\n", c1, c2, c1)

}
func zero(ptr *[32]uint8) {
	*ptr = [32]uint8{}
}

func test_slice() {
	var r []rune
	for _, v := range "hello world" {
		r = append(r, v)
	}

	fmt.Printf("%q\n", r)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func test_append() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func test_remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func test_maop() {
	ages := make(map[string]int)
	ages["aa"] = 100
	ages["bb"] = 99

	fmt.Println(ages)
}

type tree struct {
	val         int
	left, right *tree
}

func test_sortTree(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}

	sorted := appendValues(values[:0], root)
	fmt.Println(sorted)
}

func add(t *tree, v int) *tree {
	if t == nil {
		t = new(tree)
		t.val = v
		return t
	}

	if v < t.val {
		t.left = add(t.left, v)
	} else {
		t.right = add(t.right, v)
	}
	return t
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.val)
		values = appendValues(values, t.right)
	}
	return values
}

func main() {
	a := []int{3, 4, 1, 8, 9, 0, 5, 6, 5, 4}
	test_sortTree(a)

	sort.Ints(a)
	fmt.Println(a)

	// test_maop()

	// z := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(test_remove(z, 3))

	// test_append()

	// test_slice()

	// test_array()

	// fetch_url()

	// fetch_url_parallel()

	// web_server()

	// test_iota()

	// var s, sep string
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)

	// var s, sep string
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	// fmt.Println(s)

	// giff()
}
