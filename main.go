package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Result struct {
	XMLName  xml.Name `xml:"Result" json:"-"`
	U12      string   `xml:"U12" json:"U12"`
	P        string   `xml:"P" json:"P"`
	PD       string   `xml:"PD" json:"PD"`
	F        string   `xml:"F" json:"F"`
	FP       string   `xml:"FP" json:"FP"`
	FPT      string   `xml:"FPT" json:"FPT"`
	UB       string   `xml:"UB" json:"UB"`
	U23      string   `xml:"U23" json:"U23"`
	U31      string   `xml:"U31" json:"U31"`
	V1       string   `xml:"V1" json:"V1"`
	V2       string   `xml:"V2" json:"V2"`
	V3       string   `xml:"V3" json:"V3"`
	I1       string   `xml:"I1" json:"I1"`
	I2       string   `xml:"I2" json:"I2"`
	I3       string   `xml:"I3" json:"I3"`
	IN       string   `xml:"IN" json:"IN"`
	Q        string   `xml:"Q" json:"Q"`
	S        string   `xml:"S" json:"S"`
	CPP      string   `xml:"CPP" json:"CPP"`
	CPT      string   `xml:"CPT" json:"CPT"`
	CQP      string   `xml:"CQP" json:"CQP"`
	CQT      string   `xml:"CQT" json:"CQT"`
	MOT1     string   `xml:"MOT1" json:"MOT1"`
	MOT2     string   `xml:"MOT2" json:"MOT2"`
	MOT3     string   `xml:"MOT3" json:"MOT3"`
	MOT4     string   `xml:"MOT4" json:"MOT4"`
	MOT5     string   `xml:"MOT5" json:"MOT5"`
	IB       string   `xml:"IB" json:"IB"`
	PH       string   `xml:"PH" json:"PH"`
	TH       string   `xml:"TH" json:"TH"`
	WTH      string   `xml:"WTH" json:"WTH"`
	NF       string   `xml:"NF" json:"NF"`
	RPM      string   `xml:"RPM" json:"RPM"`
	CTP      string   `xml:"CTP" json:"CTP"`
	CTT      string   `xml:"CTT" json:"CTT"`
	WTB      string   `xml:"WTB" json:"WTB"`
	OVERRIDE string   `xml:"OVERRIDE" json:"OVERRIDE"`
}

type Data struct {
	XMLName    xml.Name `xml:"data" json:"-"`
	PersonList []Person `xml:"person" json:"people"`
}

type Person struct {
	XMLName   xml.Name `xml:"person" json:"-"`
	Firstname string   `xml:"firstname" json:"firstname"`
	Lastname  string   `xml:"lastname" json:"lastname"`
	Address   *Address `xml:"address" json:"address,omitempty"`
}

type Address struct {
	City  string `xml:"city" json:"city,omitempty"`
	State string `xml:"state" json:"state,omitempty"`
}

func main() {
	//rawXmlData := `
	rawXmlData := `<?xml version="1.0" encoding="UTF-8" ?>
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
	var data Result
	xml.Unmarshal([]byte(rawXmlData), &data)
	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))
}
