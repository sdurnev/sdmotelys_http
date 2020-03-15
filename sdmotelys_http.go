package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data"
	"io/ioutil"
	"os"
)

const version = "0.0.1"

type Result struct {
	XMLName  xml.Name `xml:"Result" json:"-"`
	U12      float64  `xml:"U12" json:"GENS_L1-L2"`
	P        float64  `xml:"P" json:"P"`
	PD       float64  `xml:"PD" json:"PD"`
	F        float64  `xml:"F" json:"GENS_FREQ"`
	FP       float64  `xml:"FP" json:"FP"`
	FPT      float64  `xml:"FPT" json:"FPT"`
	UB       float64  `xml:"UB" json:"GENS_BAT_U"`
	U23      float64  `xml:"U23" json:"GENS_L2-L3"`
	U31      float64  `xml:"U31" json:"GENS_L3-L1"`
	V1       float64  `xml:"V1" json:"GENS_L1-N"`
	V2       float64  `xml:"V2" json:"GENS_L2-N"`
	V3       float64  `xml:"V3" json:"GENS_L3-N"`
	I1       float64  `xml:"I1" json:"GENS_CUR_A_L1"`
	I2       float64  `xml:"I2" json:"GENS_CUR_A_L2"`
	I3       float64  `xml:"I3" json:"GENS_CUR_A_L3"`
	IN       float64  `xml:"IN" json:"IN"`
	Q        float64  `xml:"Q" json:"Q"`
	S        float64  `xml:"S" json:"S"`
	CPP      float64  `xml:"CPP" json:"GENS_ACT_ENER_P"`
	CPT      float64  `xml:"CPT" json:"GENS_ACT_ENER_T"`
	CQP      float64  `xml:"CQP" json:"GENS_RACT_ENER_P"`
	CQT      float64  `xml:"CQT" json:"GENS_RACT_ENER_T"`
	MOT1     float64  `xml:"MOT1" json:"MOT1"`
	MOT2     float64  `xml:"MOT2" json:"MOT2"`
	MOT3     float64  `xml:"MOT3" json:"MOT3"`
	MOT4     float64  `xml:"MOT4" json:"MOT4"`
	MOT5     float64  `xml:"MOT5" json:"MOT5"`
	IB       float64  `xml:"IB" json:"IB"`
	PH       float64  `xml:"PH" json:"GENS_OIL_P"`
	TH       string   `xml:"TH" json:"GENS_OIL_T"`
	WTH      float64  `xml:"WTH" json:"GENS_TEN_D"`
	NF       float64  `xml:"NF" json:"NF"`
	RPM      float64  `xml:"RPM" json:"GENS_RPM"`
	CTP      string   `xml:"CTP" json:"GENS_WORK_T_P"`
	CTT      string   `xml:"CTT" json:"GENS_WORK_T_D"`
	WTB      float64  `xml:"WTB" json:"WTB"`
	OVERRIDE float64  `xml:"OVERRIDE" json:"OVERRIDE"`
	VERSION  string   `xml:"VERSION" json:"VERSION"`
}

func Parse(st string) float64 {
	var h, m, s int
	n, err := fmt.Sscanf(st, "%d:%d:%d", &h, &m, &s)
	if err != nil || n != 3 {
		return 0
	}
	//	fmt.Println(float64(h))
	//	fmt.Println(float64(m))
	hour := float64(h) + float64(m)/60
	return float64(hour)
}

func main() {

	fileName := flag.String("filename", "", "a string")
	flag.Parse()

	if *fileName == "" {
		fmt.Printf("{\"error\":\"no file\", \"VERSION\": \"%s\"}", version)
		os.Exit(2)
	}
	//	serverParam := fmt.Sprint(*fileName)

	xmlFile, err := os.Open(*fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	//fmt.Println("Successfully Opened users.xml")

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var feed Result

	reader := bytes.NewReader(byteValue)
	//reader := bytes.NewReader([]byte(rawXmlData))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReader
	err = decoder.Decode(&feed)
	if err != nil {
		fmt.Printf("{\"decoder error\": \"%s\", \"VERSION\": \"%s\"} \n", err, version)
		os.Exit(4)
	} else {
		xml.Unmarshal(byteValue, &feed)
		t1 := Parse(feed.CTP)
		t2 := Parse(feed.CTT)
		feed.CTP = fmt.Sprintf("%.2f", t1)
		feed.CTT = fmt.Sprintf("%.2f", t2)
		feed.VERSION = version
		jsonData, _ := json.Marshal(feed)
		fmt.Println(string(jsonData))
	}
}
