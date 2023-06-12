package golibri

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func Yyyymmdd_hhnn_ss() string {
	apilast := time.Now().Format("20060102150405")
	return apilast[:8] + "_" + apilast[8:12] + "_" + apilast[12:14]
}

//  var wg *WaitMapObject

type WaitMapObject struct {
	wg   map[string]int
	mu   sync.Mutex
	cond sync.Cond
}

func WaitMap() *WaitMapObject {
	m := &WaitMapObject{}
	m.wg = make(map[string]int)
	m.cond.L = &m.mu
	return m
}

func (m *WaitMapObject) Wait(name string) {
	m.mu.Lock()
	for m.wg[name] != 0 {
		m.cond.Wait()
	}
	m.mu.Unlock()
}

func (m *WaitMapObject) Done(name string) {
	m.mu.Lock()
	no := m.wg[name] - 1
	if no < 0 {
		panic("")
	}
	m.wg[name] = no
	m.mu.Unlock()
	m.cond.Broadcast()
}

func (m *WaitMapObject) Add(name string, no int) {
	m.mu.Lock()
	m.wg[name] = m.wg[name] + no
	m.mu.Unlock()
}

// --------------------------------------------------------------------------

func Rep(st1 string, st2 string, st3 string) string {
	return strings.Replace(st1, st2, st3, -1)
}

// singlquote to Double Quotes
func Rep2(st string) string {
	return strings.Replace(st, "'", "''", -1)
}

func RFi(stPAth string) string {
	st := ""
	file, err := os.Open(stPAth)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		st = st + scanner.Text() + "\n"
	}
	return strings.TrimSpace(st)
}

func Wfi(stPAth string, content string) {
	ioutil.WriteFile(stPAth, []byte(content), 0644)
}

func NotFi(stpath string) bool {
	if _, err := os.Stat(stpath); err != nil {
		if os.IsNotExist(err) {

			return true
		}
	}
	return false
}

func GetMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}

func GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		return ""
	}
	return hostname

}

// ------------------------

func MapFilesKey(fold string, ext string) map[string]string {
	m := map[string]string{}
	nffi := CountFiles(fold, ext)
	if nffi > 0 {
		m = make(map[string]string, nffi)
	}
	f, err := os.Open(fold)
	if err != nil {
		return m
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return m
	}
	for _, v := range files {
		if !v.IsDir() {
			m[v.Name()] = RFi(fold + "/" + v.Name())
		}
	}
	return m
}

func GetFilesInfo(fold string) []os.FileInfo {
	f, err := os.Open(fold)
	if err != nil {
		return nil
	}
	fso, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return fso
}

func osEnv(key string) string {
	vv := Rep(os.Getenv(key), `"`, ``)

	//to put default value if keyenv is forgotten
	if vv == "" {
		switch key {
		case "UndeclaredKeyNameInEnv":
			return "MyDefaultValue"
		default:
			return ""
		}
	}

	return vv
}
