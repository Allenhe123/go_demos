package main

import "fmt"

//import "strings"
//import "os"

//import "regexp"
//import "bytes"

//import "encoding/json"
//import "os"

//import "time"

//import "math/rand"

//import "strconv"

//import "net/url"
//import "strings"

//import "crypto/sha1"

//import b64 "encoding/base64"

//import (
//"bufio"
//	"io"
//	"io/ioutil"
//"os"
//)

//import "strings"
//import "os"
//import "bufio"

import "os"

//import "strings"

//import "flag"

////获取目标字符串t出现的第一个索引位置,没有就返回-1
func Index(vc []string, t string) int {
	for i, v := range vc {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vc []string, t string) bool {
	return Index(vc, t) >= 0
}

func Any(vc []string, f func(string) bool) bool {
	for _, v := range vc {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

//满足条件的元素
func Filter(vs []string, f func(string) bool) []string {
	vcf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vcf = append(vcf, v)
		}
	}
	return vcf
}

//返回一个对原始切片中所有字符串执行函数 f 后的新切片。
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	/*
		var p = fmt.Println

		p("Contains:", strings.Contains("test", "es"))
		p("Count:", strings.Count("test", "t"))
		p("HasPrefix:", strings.HasPrefix("test", "te"))
		p("HasSuffix:", strings.HasSuffix("test", "st"))
		p("Index", strings.Index("test", "s"))
		p("Join:", strings.Join([]string{"a", "b"}, "-"))
		p("Repeat:", strings.Repeat("a", 5))
		p("Replace:", strings.Replace("foo", "o", "0", -1))
		p("Replace:", strings.Replace("foo", "o", "0", 1))
		p("Split:", strings.Split("a-b-c-d-e", "-"))
		p("ToLower:", strings.ToLower("TEST"))
		p("ToUpper:", strings.ToUpper("test"))
		p()
		p("Len:", len("hello"))
		p("Char:", "hello"[2])

	*/

	/////////////////////////////////
	/*
		type point struct {
			a int
			b int
		}

		p := point{1, 2}
		fmt.Printf("%v \n", p) //打印了一个对象

		fmt.Printf("%+v \n", p)   //包括字段名一起打印出
		fmt.Printf("%#v \n", p)   //值的运行源代码片段
		fmt.Printf("%T \n", p)    //值的类型
		fmt.Printf("%t \n", true) //bool值
		fmt.Printf("%d \n", 123)
		fmt.Printf("%b \n", 12) //二进制
		fmt.Printf("%c \n", 21)

		fmt.Printf("%x \n", 456)          //十六进制
		fmt.Printf("%f \n", 889.1)        //浮点数
		fmt.Printf("%e \n", 123400000.0)  //科学计数法的形式表示
		fmt.Printf("%E \n", 123400000.0)  //科学计数法的形式表示
		fmt.Printf("%s \n", "\"string\"") //""内的内容支持转义,不包含最外层双引号
		fmt.Printf("%q \n", "\"string\"") //包含双引号

		fmt.Printf("%x", "hex")                    //base-16编码
		fmt.Printf("%p \n", &p)                    //输出一个指针的值
		fmt.Printf("|%6d|%6.2f| \n", 12, 234.0)    //控制位数,不足的空格代替,默认右对齐
		fmt.Printf("|%-6.2f|%-6.2f\n", 1.2, 123.2) //左对齐-
		//控制字符串的宽度
		fmt.Printf("|%6s|%6s| \n", "hello", "world")

		//格式化返回一个字符串,不输出
		s := fmt.Sprintf(" a %s", "string")
		fmt.Println(s)

		//Fprintf格式化并输出到io.Writers而不是os.Stdout
		fmt.Fprintf(os.Stderr, " an %s \n", "error")
	*/

	///////////////////////////////////////////////////////////////
	/*
		match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
		fmt.Println(match) //true

		//优化的regexp结构
		r, _ := regexp.Compile("p([a-z+]ch)")

		fmt.Println(r.MatchString("peach"))

		//fmt.Println(r.FindString("peach punch")) //查找匹配字符串的

		fmt.Println(r.FindStringIndex("peach punch")) //得到的是匹配内容的起始和结束的下标

		fmt.Println(r.FindAllString("peadch punch pinch", -1)) //返回所有的匹配项

		fmt.Println(r.FindStringSubmatchIndex("peach punch"))
		//返回完全匹配和局部匹配的索引位置

		fmt.Println(r.FindAllStringSubmatchIndex(
			"peach punch pinch", -1))
		//All对应到上面的所有函数

		fmt.Println(r.Match([]byte("peach")))

		r = regexp.MustCompile("p([a-z]+)ch") //Compile的变体
		fmt.Println(r)

		fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) //替换

		in := []byte("a peach")
		out := r.ReplaceAllFunc(in, bytes.ToUpper)
		fmt.Println(string(out))
		//Func变量允许传递匹配内容到一个给定的函数中
	*/

	/////////////////////////////////////////////////////////////////////////////////////
	/*
		//将使用这两个结构体来演示自定义类型的编码和解码。
		type Response1 struct {
			Page   int
			Fruits []string
		}

		type Response2 struct {
			Page   int      `json:"page"`
			Fruits []string `json:"fruits"` //自定义json数据键名
		}

		//基本数据类型-->JSON编码
		//func Marshal(v interface{}) ([]byte, error)
		bolB, err := json.Marshal(true)
		if err != nil {
			fmt.Println("error: ", err)
		} else {
			fmt.Println(string(bolB))
			//os.Stdout.Write(bolB)
		}

		intB, err := json.Marshal(12)
		if err != nil {
			fmt.Println("error: ", err)
		} else {
			fmt.Println(string(intB))
			//os.Stdout.Write(intB)
		}

		slc := []string{"1", "2", "3"}
		slcB, err := json.Marshal(slc)
		if err != nil {
			fmt.Println("error: ", err)
		} else {
			fmt.Println(string(slcB)) //必须转换成string，不然打印的是ASCI码
			//os.Stdout.Write(slcB)
		}

		mapD := map[string]int{"a": 2, "b": 1}
		mapB, _ := json.Marshal(mapD)
		fmt.Println(string(mapB))

		resp1 := &Response1{
			Page: 1,
			Fruits: []string{
				"apple",
				"peach",
				"pear",
			},
		}

		resp1B, _ := json.Marshal(resp1)
		fmt.Println(string(resp1B))

		resp2 := Response2{
			Page:   1,
			Fruits: []string{"apple", "pear", "peach"},
		}
		resp2B, _ := json.Marshal(resp2)
		fmt.Println(string(resp2B))

		//解码JSON数据为Go值的过程
		byt := []byte(`{"num":87.2, "strs":["a","c"]}`)
		var data map[string]interface{} //任意类型的值
		if err := json.Unmarshal(byt, &data); err != nil {
			panic(err)
		}
		fmt.Println(data)

		num := data["num"].(float64)
		fmt.Println(num)

		strs := data["strs"].(interface{})
		fmt.Println(strs)

		//解码json值为自定义的类型
		str := `{"page":1, "fruits":["apple","peach"]}`
		res2 := Response2{}
		json.Unmarshal([]byte(str), &res2)
		fmt.Println(res2)
		fmt.Println(res2.Fruits[0])

		res1 := Response1{}
		json.Unmarshal([]byte(str), &res1)
		fmt.Println(res1)
		fmt.Println(res1.Fruits[0])

		//我们经常使用 byte 和 string 作为使用标准输出时数据和 JSON 表示之间的中间值。我们也可以和os.Stdout 一样，
		//直接将 JSON 编码直接输出至 os.Writer流中，或者作为 HTTP 响应体。
		enc := json.NewEncoder(os.Stdout)
		d := map[string]int{"apple": 2, "peach": 987}
		enc.Encode(d)
	*/
	/////////////////////////////////////////////////////////////////////////////////////////////////////
	/*
		now := time.Now()
		fmt.Println(now) //2016-06-31 15:50:13.793654 +0000 UTC

		then := time.Date(2016, 11, 17, 20, 34, 58, 651387237, time.UTC) //构建一个time

		//取出时间的各个部分
		fmt.Println(then.Year())
		fmt.Println(then.Month())
		fmt.Println(then.Day())
		fmt.Println(then.Hour())
		fmt.Println(then.Minute())
		fmt.Println(then.Second())
		fmt.Println(then.Nanosecond()) //纳秒，1,000 纳秒 = 1微秒 1,000,000 纳秒 = 1毫秒
		fmt.Println(then.Location())

		fmt.Println(then.Weekday())

		//比较
		fmt.Println(then.Before(now))
		fmt.Println(then.After(now))
		fmt.Println(then.Equal(now))

		diff := now.Sub(then)
		fmt.Println(diff) //两个时间点的duration

		fmt.Println(diff.Hours())
		fmt.Println(diff.Minutes())
		fmt.Println(diff.Seconds())
		fmt.Println(diff.Nanoseconds())

		fmt.Println(then.Add(diff))  //时间往后移
		fmt.Println(then.Add(-diff)) //时间往前移
	*/
	/*
		now := time.Now()
		secs := now.Unix()
		nanos := now.UnixNano()
		millis := nanos / 1000000
		fmt.Println(now)
		fmt.Println(secs)
		fmt.Println(millis)
		fmt.Println(nanos)

		//将秒数或者毫秒数转为时间
		fmt.Println(time.Unix(secs, 0))
		fmt.Println(time.Unix(0, nanos))
	*/
	/*
		//Go支持基于描述模板的时间格式化和解析
		t := time.Now()
		fmt.Println(t)
		//按照RFC3339进行格式化
		fmt.Println(t.Format(time.RFC3339))

		//时间解析格式
		t1, e := time.Parse(
			time.RFC3339,
			"2013-12-02T22:09:12+00:00",
		)

		fmt.Println(t1) //2013-12-02 22:09:12 +0000 +0000

		fmt.Println(t.Format("2:12PM"))                   //4:114PM
		fmt.Println(t.Format("Mon Jan _2 12:22:12 2091")) //Sat Nov  4 114:44:114 40911
		p := fmt.Println
		p(t.Format("3:04PM"))    //7:37PM
		p(t.Format("Mon Jan _2 15:04:05 2006"))   //Sat Nov  4 19:37:51 2017
		p(t.Format("2006-01-02T15:04:05.999999-07:00"))
		form := "3 04 PM"
		t2, e := time.Parse(form, "8 41 PM")
		p(t2)
		//Format 和 Parse 使用基于例子的形式来决定日期格式，一般你只要使用 time 包中提供的模式常量就行了，但是你也可以实现自定义模式。模式必须使用时间 Mon Jan 2 15:04:05 MST 2006来指定给定时间/字符串的格式化/解析方式。时间一定要按照如下所示：2006为年，15 为小时，Monday 代表星期几，等等。

		fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
		//对于纯数字表示的时间，你也可以使用标准的格式化字符串来提出出时间值得组成。

		ansic := "Mon Jan _2 15:04:05 2006"
		_, e = time.Parse(ansic, "8:41PM")
		p(e)
		//Parse 函数在输入的时间格式不正确是会返回一个错误。
	*/

	/*
		//rand.Intn(m),随机整数,0 <= n <= m
		fmt.Println(rand.Intn(170))
		fmt.Println(rand.Intn(21))

		//伪随机数生成器
		s1 := rand.NewSource(40)
		r1 := rand.New(s1)

		//如果使用相同的种子生成的随机数生成器，将会产生相同的随机数序列。
		s2 := rand.NewSource(40)
		r2 := rand.New(s2)
		fmt.Println(r1.Intn(100))
		fmt.Println(r2.Intn(100))
		fmt.Println()
	*/

	//数字字符串转为数字
	//f, _ := strconv.ParseFloat("12.3", 64) //64解析的位数
	//fmt.Println(f)
	/*
	   ParseInt
	   ParseUint
	   Atoi //基础的10进制转换函数

	   //如果数字字符串不是数字,会报错
	*/

	//////////////////////////////////////////////////////////////////////////////////////

	/*	//url解析
		s := "postgres://user:pass@host.com:5432/path?k=v#f"
		//我们将解析这个 URL 示例，它包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段。

		u, err := url.Parse(s)
		if err != nil {
			panic(err)
		}
		//解析这个 URL 并确保解析没有出错。
		fmt.Println(u.Scheme)
		fmt.Println(u.User)
		fmt.Println(u.User.Username())
		p, _ := u.User.Password()
		fmt.Println(p)
		//User 包含了所有的认证信息，这里调用 Username和 Password 来获取独立值。

		fmt.Println(u.Host) //host.com:5432
		h := strings.Split(u.Host, ":")
		fmt.Println(h[0]) //host.com
		fmt.Println(h[1]) //5432
		//Host 同时包括主机名和端口信息，如过端口存在的话，使用 strings.Split() 从 Host 中手动提取端口。

		fmt.Println(u.Path)     // /path
		fmt.Println(u.Fragment) // f
		//这里我们提出路径和查询片段信息。

		fmt.Println(u.RawQuery) // k=v
		m, _ := url.ParseQuery(u.RawQuery)
		fmt.Println(m)
		fmt.Println(m["k"][0])
		//要得到字符串中的 k=v 这种格式的查询参数，可以使用 RawQuery 函数。你也可以将查询参数解析为一个map。
		//已解析的查询参数 map 以查询字符串为键，对应值字符串切片为值，所以如何只想得到一个键对应的第一个值，将索引位置设置为 [0] 就行了。
	*/

	/*
		/////////Go 在多个 crypto/* 包中实现了一系列散列函数。
		s := "sha1 this string"
		h := sha1.New()
		//产生一个散列值得方式是 sha1.New()，sha1.Write(bytes)，然后 sha1.Sum([]byte{})。这里我们从一个新的散列开始。

		h.Write([]byte(s))
		//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。

		bs := h.Sum(nil)
		//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来在现有的字符切片追加额外的字节切片：一般不需要。

		fmt.Println(s)
		fmt.Printf("%x\n", bs)
		//SHA1 值经常以 16 进制输出，例如在 git commit 中。使用%x 来将散列结果格式化为 16 进制字符串。
		//可以使用和上面相似的方式来计算其他形式的散列值。例如，计算 MD5 散列，引入 crypto/md5 并使用md5.New()方法。
	*/

	/*
		//base64编码
		data := "abc123!?$*&()'-=@~" //这是将要编解码的字符串。

		sEnc := b64.StdEncoding.EncodeToString([]byte(data))
		fmt.Println(sEnc) //Go 同时支持标准的和 URL 兼容的 base64 格式。编码需要使用 []byte 类型的参数，所以要将字符串转成此类型。

		sDec, _ := b64.StdEncoding.DecodeString(sEnc)
		fmt.Println(string(sDec))
		fmt.Println()
		//解码可能会返回错误，如果不确定输入信息格式是否正确，那么，你就需要进行错误检查了。

		uEnc := b64.URLEncoding.EncodeToString([]byte(data))
		fmt.Println(uEnc)
		uDec, _ := b64.URLEncoding.DecodeString(uEnc)
		fmt.Println(string(uDec))
		//使用 URL 兼容的 base64 格式进行编解码。

		//标准 base64 编码和 URL 兼容 base64 编码的编码字符串存在稍许不同（后缀为 + 和 -），但是两者都可以正确解码为原始字符串。
	*/

	////////////////////readfile//////////////////////////////
	/*	data, err := ioutil.ReadFile("E:\\abc.txt") //读取文件内容到内存中
		check(err)
		fmt.Println(string(data))

		f, err := os.OpenFile("E:\\abc.txt", os.O_APPEND, 0666)
		check(err)
		defer f.Close()
		b1 := make([]byte, 5) //最多读取5个字节
		n1, err := f.Read(b1)
		check(err)
		fmt.Printf("%d bytes: %s\n", n1, string(b1))

		//从一个文件中已知的位置开始进行读取
		o2, err := f.Seek(6, 0)
		check(err)
		b2 := make([]byte, 2)
		n2, err := f.Read(b2)
		check(err)
		fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

		//io包提供了一些帮助进行文件读取的函数,ReadAtLeast 得到一个更更好的实现。
		o3, err := f.Seek(6, 0)
		check(err)
		b3 := make([]byte, 2)
		n3, err := io.ReadAtLeast(f, b3, 2)
		check(err)
		fmt.Printf("%d bytes @ %d: %s \n", n3, o3, string(b3))

		_, err = f.Seek(0, 0)
		check(err)

		//缓冲的读取,性能好
		r4 := bufio.NewReader(f)
		b4, err := r4.Peek(5)
		check(err)
		fmt.Printf("5 bytes:%s \n", string(b4))
	*/
	//////////////////////write file//////////////////
	/*
		//写入一个字符串到文件
		d1 := []byte("hello\ngo\n")
		err := ioutil.WriteFile("E:\\data.txt", d1, 0644)
		check(err)

		//先打开一个文件
		f, err := os.Create("E:\\data2.txt")
		check(err)
		defer f.Close() //最后记得关闭文件

		//写入切片
		d2 := []byte{12, 3, 2, 1, 32, 55}
		n2, err := f.Write(d2)
		check(err)
		fmt.Printf("wrote %d bytes\n", n2)

		n3, err := f.WriteString("writes\n")
		fmt.Printf("wrote %d bytes\n", n3)

		f.Sync() //将缓冲区的数据写入硬盘

		//带有缓冲的写入器
		w := bufio.NewWriter(f)
		n4, err := w.WriteString("buffered\n")
		fmt.Printf("wrote %d bytes\n", n4)
		w.Flush() //确保所有的缓存操作已写入底层写入器
	*/
	/*
		///////////////////////////// 行为过滤器////////////////////////////////////
		//带有缓冲的输入
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			ucl := strings.ToUpper(scanner.Text())
			fmt.Println(ucl)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
		//检查Scan的错误.
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	*/

	/*
		////////////////命令行参数///////////////////////////
		//命令行参数是指定程序运行参数的一个常见方式。例如，go run hello.go，程序 go 使用了 run 和 hello.go两个参数。
		//os.Args 提供原始命令行参数访问功能。注意，切片中的第一个参数是该程序的路径，并且 os.Args[1:]保存所有程序的的参数。
		argsWithProg := os.Args
		argsWithoutProg := os.Args[1:]
		arg := os.Args[3]
		//你可以使用标准的索引位置方式取得单个参数的值。

		fmt.Println(argsWithProg)
		fmt.Println(argsWithoutProg)
		fmt.Println(arg)
	*/

	/*
		//////////////////////////命令行标志(参数)////////////////////////////////////////)
		//命令行标志是命令行程序指定选项的常用方式。例如，在wc -l 中，这个 -l 就是一个命令行标志。
		////Go 提供了一个 flag 包，支持基本的命令行标志解析。我们将用这个包来实现我们的命令行程序示例。

		wordPtr := flag.String("word", "foo", "a string")
		//基本的标记声明仅支持字符串、整数和布尔值选项。这里我们声明一个默认值为 "foo" 的字符串标志 word并带有一个简短的描述。
		//这里的 flag.String 函数返回一个字符串指针（不是一个字符串值），在下面我们会看到是如何使用这个指针的。

		numbPtr := flag.Int("numb", 42, "an int")
		boolPtr := flag.Bool("fork", false, "a bool")
		//使用和声明 word 标志相同的方法来声明 numb 和 fork 标志。

		var svar string
		flag.StringVar(&svar, "svar", "bar", "a string var")
		//用程序中已有的参数来声明一个标志也是可以的。注意在标志声明函数中需要使用该参数的指针。

		flag.Parse()
		//所有标志都声明完成以后，调用 flag.Parse() 来执行命令行解析。

		fmt.Println("word:", *wordPtr)
		fmt.Println("numb:", *numbPtr)
		fmt.Println("fork:", *boolPtr)
		fmt.Println("svar:", svar)
		fmt.Println("tail:", flag.Args())
		//这里我们将仅输出解析的选项以及后面的位置参数。注意，我们需要使用类似 *wordPtr 这样的语法来对指针解引用，从而得到选项的实际值。
	*/

	//环境变量是一个在为 Unix 程序传递配置信息的普遍方式。让我们来看看如何设置，获取并列举环境变量。
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	//使用 os.Setenv 来设置一个键值对。使用 os.Getenv获取一个键对应的值。如果键不存在，将会返回一个空字符串。

	fmt.Println()
	for _, e := range os.Environ() {
		fmt.Println(e)
		//pair := strings.Split(e, "=")
		//fmt.Println(pair[0])
	}
	//使用 os.Environ 来列出所有环境变量键值对。这个函数会返回一个 KEY=value 形式的字符串切片。你可以使用strings.Split 来得到键和值。这里我们打印所有的键。

}
