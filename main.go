package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Result struct {
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
	rawXmlData := "<?xml version="1.0" encoding="ISO-8859-15" ?>
	<Result DateTime="5/3/2020 22:33">
	<U12>0.0</U12>
	<P>0.0</P>
	<PD>100</PD>
	<F>0.0</F>
	<FP>0.0</FP>
	<FPT>0</FPT>
	<UB>12.7</UB>
	<U23>0.0</U23>
	<U31>0.0</U31>
	<V1>0.0</V1>
	<V2>0.0</V2>
	<V3>0.0</V3>
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
	<PH>0.5</PH>
	<TH>----</TH>
	<WTH>41</WTH>
	<NF>100</NF>
	<RPM>0</RPM>
	<CTP>00:50:35</CTP>
	<CTT>581:25:49</CTT>
	<WTB>0</WTB>
	<OVERRIDE>0</OVERRIDE>
	</Result>"
	var data Result
	xml.Unmarshal([]byte(rawXmlData), &data)
	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))
}