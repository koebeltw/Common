package loadConf

import (
	"log"
	"strconv"
	"strings"

	"github.com/zieckey/goini"
)

// // LoadConf blabla
// type LoadConf struct {
// }

// // CreateLoadConf blabla
// func CreateLoadConf() (l *LoadConf) {
// 	return &LoadConf{}
// }

// LoadIP blabla
func LoadIP(addr, content, subContent string) (IP string) {
	// currentDir, err := os.Getwd()
	ini := goini.New()
	err := ini.ParseFile(addr)
	if err != nil {
		// log.Println(currentDir + "/config/" + strings.ToLower(content) + "/config.ini" + "Error")
		return
	}
	IP, ok := ini.SectionGet(content, subContent)
	if !ok {
		log.Println(content, subContent, "ini.SectionGet Error:", ok)
		return
	}

	log.Println("LoadIP:", IP)
	return
}

// LoadIPMap blabla
func LoadIPMap(addr, content, subContent string) (IPMap map[uint16]string) {
	IPMap = make(map[uint16]string, 0)
	// currentDir, err := os.Getwd()
	ini := goini.New()
	err := ini.ParseFile(addr)
	if err != nil {
		// log.Println(currentDir + "/config/" + strings.ToLower(content) + "/config.ini" + "Error")
		return
	}
	v, ok := ini.GetKvmap(content)
	if !ok {
		log.Println("ini.GetKvmap Error:", ok)
		return
	}

	for key, value := range v {
		if strings.Contains(strings.ToLower(key), strings.ToLower(subContent)) == false {
			if strings.Contains(key, subContent) == false {
				continue
			}
			log.Println(subContent, "ToLower Error")
			continue
		}

		// tcp, err := net.ResolveTCPAddr("tcp", value)
		// if err != nil {
		// 	log.Printf("Coneect ResolveTCPAddr Error: %s\n", err.Error())
		// 	return
		// }
		// tcp.IP.String(), tcp.Port

		s := strings.Replace(key, subContent, "", -1)
		if s == "" {
			IPMap[0] = value
			continue
		}

		if ID, err := strconv.Atoi(s); err == nil {
			IPMap[uint16(ID)] = value
		}

	}

	log.Println(content, subContent, "IPMap:", IPMap)
	return
}
