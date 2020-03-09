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
	"log"
	"net/http"
)

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
	CTP      string   `xml:"CTP" json:"CTP"`
	CTT      string   `xml:"CTT" json:"CTT"`
	WTB      float64  `xml:"WTB" json:"WTB"`
	OVERRIDE float64  `xml:"OVERRIDE" json:"OVERRIDE"`
}

func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}

func main() {
	var rawXmlData string

	StringData := flag.String("data", "0", "a string")
	flag.Parse()

	if xmlBytes, err := getXML("http://localhost:1123/WebTelys?Request=Mesure"); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		//var result myXMLstruct
		fmt.Println(xmlBytes)
	}

	if *StringData == "0" {
		rawXmlData = `<?xml version="1.0" encoding="ISO-8859-15" ?>
		<Result DateTime="6/3/2020 10:05">
		<U12>399.8</U12>
		<P>0.0</P>
		<PD>100</PD>
		<F>52.19</F>
		<FP>0.0</FP>
		<FPT>0</FPT>
		<UB>14.7</UB>
		<U23>403.0</U23>
		<U31>400.7</U31>
		<V1>230.5</V1>
		<V2>231.8</V2>
		<V3>232.4</V3>
		<I1>0.0</I1>
		<I2>0.0</I2>
		<I3>0.0</I3>
		<IN>0.0</IN>
		<Q>0.0</Q>
		<S>0.0</S>
		<CPP>2448</CPP>
		<CPT>2448</CPT>
		<CQP>0</CQP>
		<CQT>0</CQT>
		<MOT1>0</MOT1>
		<MOT2>0</MOT2>
		<MOT3>0</MOT3>
		<MOT4>0</MOT4>
		<MOT5>0</MOT5>
		<IB>0</IB>
		<PH>4.7</PH>
		<TH>----</TH>
		<WTH>67</WTH>
		<NF>100</NF>
		<RPM>1565</RPM>
		<CTP>00:55:25</CTP>
		<CTT>581:30:39</CTT>
		<WTB>0</WTB>
		<OVERRIDE>0</OVERRIDE>
		</Result>`
	} else {
		rawXmlData = *StringData
	}

	var feed Result

	reader := bytes.NewReader([]byte(rawXmlData))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReader
	err := decoder.Decode(&feed)
	if err != nil {
		fmt.Println("decoder error:", err)
	}
	//fmt.Println(feed)

	xml.Unmarshal([]byte(rawXmlData), &feed)
	jsonData, _ := json.Marshal(feed)
	fmt.Println(string(jsonData))
}
