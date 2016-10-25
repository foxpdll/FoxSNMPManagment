package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/base64"
    "time"
    "flag"
    _ "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/alouca/gosnmp"
    gosnmp2 "github.com/soniah/gosnmp"
    "strings"
    "strconv"
)

var (
    addr = flag.String("addr", ":81", "http service address")
    ifType = []string{"other", "regular1822", "hdh1822", "ddnX25", "rfc877x25", "ethernetCsmacd", "iso88023Csmacd", "iso88024TokenBus", "iso88025TokenRing", "iso88026Man", "starLan", "proteon10Mbit", "proteon80Mbit", "hyperchannel", "fddi", "lapb", "sdlc", "ds1", "e1", "basicISDN", "primaryISDN", "propPointToPointSerial", "ppp", "softwareLoopback", "eon", "ethernet3Mbit", "nsip", "slip", "ultra", "ds3", "sip", "frameRelay", "rs232", "para", "arcnet", "arcnetPlus", "atm", "miox25", "sonet", "x25ple", "iso88022llc", "localTalk", "smdsDxi", "frameRelayService", "v35", "hssi", "hippi", "modem", "aal5", "sonetPath", "sonetVT", "smdsIcip", "propVirtual", "propMultiplexor", "ieee80212", "fibreChannel", "hippiInterface", "frameRelayInterconnect", "aflane8023", "aflane8025", "cctEmul", "fastEther", "isdn", "v11", "v36", "g703at64k", "g703at2mb", "qllc", "fastEtherFX", "channel", "ieee80211", "ibm370parChan", "escon", "dlsw", "isdns", "isdnu", "lapd", "ipSwitch", "rsrb", "atmLogical", "ds0", "ds0Bundle", "bsc", "async", "cnr", "iso88025Dtr", "eplrs", "arap", "propCnls", "hostPad", "termPad", "frameRelayMPI", "x213", "adsl", "radsl", "sdsl", "vdsl", "iso88025CRFPInt", "myrinet", "voiceEM", "voiceFXO", "voiceFXS", "voiceEncap", "voiceOverIp", "atmDxi", "atmFuni", "atmIma", "pppMultilinkBundle", "ipOverCdlc", "ipOverClaw", "stackToStack", "virtualIpAddress", "mpc", "ipOverAtm", "iso88025Fiber", "tdlc", "gigabitEthernet", "hdlc", "lapf", "v37", "x25mlp", "x25huntGroup", "trasnpHdlc", "interleave", "fast", "ip", "docsCableMaclayer", "docsCableDownstream", "docsCableUpstream", "a12MppSwitch", "tunnel", "coffee", "ces", "atmSubInterface", "l2vlan", "l3ipvlan", "l3ipxvlan", "digitalPowerline", "mediaMailOverIp", "dtm", "dcn", "ipForward", "msdsl", "ieee1394", "if-gsn", "dvbRccMacLayer", "dvbRccDownstream", "dvbRccUpstream", "atmVirtual", "mplsTunnel", "srp", "voiceOverAtm", "voiceOverFrameRelay", "idsl", "compositeLink", "ss7SigLink", "propWirelessP2P", "frForward", "rfc1483", "usb", "ieee8023adLag", "bgppolicyaccounting", "frf16MfrBundle", "h323Gatekeeper", "h323Proxy", "mpls", "mfSigLink", "hdsl2", "shdsl", "ds1FDL", "pos", "dvbAsiIn", "dvbAsiOut", "plc", "nfas", "tr008", "gr303RDT", "gr303IDT", "isup", "propDocsWirelessMaclayer", "propDocsWirelessDownstream", "propDocsWirelessUpstream", "hiperlan2", "propBWAp2Mp", "sonetOverheadChannel", "digitalWrapperOverheadChannel", "aal2", "radioMAC", "atmRadio", "imt", "mvl", "reachDSL", "frDlciEndPt", "atmVciEndPt", "opticalChannel", "opticalTransport", "propAtm", "voiceOverCable", "infiniband", "teLink", "q2931", "virtualTg", "sipTg", "sipSig", "docsCableUpstreamChannel", "econet", "pon155", "pon622", "bridge", "linegroup", "voiceEMFGD", "voiceFGDEANA", "voiceDID", "mpegTransport", "sixToFour", "gtp", "pdnEtherLoop1", "pdnEtherLoop2", "opticalChannelGroup", "homepna", "gfp", "ciscoISLvlan", "actelisMetaLOOP", "fcipLink", "rpr", "qam", "lmp", "cblVectaStar", "docsCableMCmtsDownstream", "adsl2", "macSecControlledIF", "macSecUncontrolledIF", "aviciOpticalEther", "atmbond"}
    ifStatus = []string{"up","down","testing","unknown","dormant","notPresent","lowerLayerDown"}
    cdStatus = []string{"ok","open","short","open-short","crosstalk","unknown","count","no-cable","other"}
)

func checkIfaces(w http.ResponseWriter, r *http.Request){
    ip:=r.URL.Path[13:]
    fmt.Fprintf(w,"<html><head></head><body>")
    for i := range snmp_coms_r{
	s, err := gosnmp.NewGoSNMP(ip, snmp_coms_r[i], gosnmp.Version2c, 1)
	checkErr(err)
	resp,err := s.Get(".1.3.6.1.2.1.17.1.1.0")
	if err == nil{
	    _ = resp
//	    umac := fmt.Sprintf("%X",resp.Variables[0].Value)
	    fmt.Fprintf(w,"<script>function changeAlias(e,num,txt){var code = (e.keyCode ? e.keyCode : e.which); if (code == 13) {window.open('/swSetAlias/%s/'+num+'/'+txt.value);}};</script>",ip)

	    fmt.Fprintf(w,"<table border=1 width='100%'><tr><td width='50%' height=300>")
	    fmt.Fprintf(w,"<div style='width:100%%; height:100%%;'><object style='width:100%%; height:100%%;' type='text/html' data='/swMainData/%s'></object></div>",ip)
	    fmt.Fprintf(w,"</td><td width='50%' height=300>")
	    fmt.Fprintf(w,"<div style='width:100%%; height:100%%;'><object style='width:100%%; height:100%%;' type='text/html' data='/ospfIpNeib/%s'></object></div>",ip)
	    fmt.Fprintf(w,"</td></tr><tr><td>")
	    fmt.Fprintf(w,"<div style='width:100%%; height:100%%;'><object style='width:100%%; height:100%%;' type='text/html' data='/swArpTable/%s'></object></div>",ip)
	    fmt.Fprintf(w,"</td></tr></table>")

		fmt.Fprintf(w,"<a href='#' onclick='window.open(\"/swCheckAllCable/%s/\");'>ShowAllCableStatus</a>",ip)

	    fmt.Fprintf(w,"<table border=1 width='100%'>")
	    fmt.Fprintf(w,"<tr><th>ifId</th><th>ifAdminStatus</th><th>ifOperStatus</th><th>ifName</th><th>ifAlias</th><th>ifSpeed</th><th>ifType</th><th>ifMAC</th><th>ifInErr</th><th>Mcast</th><th>LLDP&FDB Table</th></tr>")
	    ifadmstat, err := s.Walk(".1.3.6.1.2.1.2.2.1.7")
	    checkErr(err)
	    ifoperstat, err := s.Walk(".1.3.6.1.2.1.2.2.1.8")
	    checkErr(err)
	    ifname, err := s.Walk(".1.3.6.1.2.1.2.2.1.2")
	    checkErr(err)
	    ifalias, err := s.Walk("1.3.6.1.2.1.31.1.1.1.18")
	    checkErr(err)
	    ifmac, err := s.Walk(".1.3.6.1.2.1.2.2.1.6")
	    checkErr(err)
	    ifspeed, err := s.Walk("1.3.6.1.2.1.2.2.1.5")
	    checkErr(err)
	    iftype, err := s.Walk("1.3.6.1.2.1.2.2.1.3")
	    checkErr(err)
	    ifinerr, err := s.Walk("1.3.6.1.2.1.2.2.1.14")
	    checkErr(err)
//	    mcast14, _ := s.Walk(".1.3.6.1.4.1.171.12.73.2.1.5.1.2")
//	    checkErr(err)
	    mcast15, _ := s.Walk(".1.3.6.1.4.1.171.12.73.2.1.2.1.4")
	    checkErr(err)
	    mcast16, _ := s.Walk(".1.3.6.1.4.1.171.11.113.1.5.2.7.13.1.4")
	    checkErr(err)
	    lldpchasid, _ := s.Walk(".1.0.8802.1.1.2.1.4.1.1.5")
	    checkErr(err)
	    lldpporttype, _ := s.Walk(".1.0.8802.1.1.2.1.4.1.1.6")
	    checkErr(err)
	    lldpport, _ := s.Walk(".1.0.8802.1.1.2.1.4.1.1.7")
	    checkErr(err)
	    lldpportdescr, _ := s.Walk(".1.0.8802.1.1.2.1.4.1.1.8")
	    checkErr(err)
	    lldpsysname, _ := s.Walk(".1.0.8802.1.1.2.1.4.1.1.9")
	    checkErr(err)
	    lldppeerip, _ := s.Walk(".1.0.8802.1.1.2.1.4.2.1.3")
	    checkErr(err)
	    arptable, _ := s.Walk(".1.3.6.1.2.1.4.22.1.2")
	    checkErr(err)
	    fdbtable, _ := s.Walk(".1.3.6.1.2.1.17.7.1.2.2.1.2")
	    checkErr(err)
////	    db, err := sql.Open("mysql", mysql_user+":"+mysql_password+"@tcp("+mysql_host+":"+mysql_port+")/"+mysql_db+"?charset=utf8")
////	    checkErr(err)
////	    stmt, err := db.Prepare("update sw_if set if_admin_status=?, if_oper_status=? where umac=CONV(?,16,10) and if_id=?")
////	    checkErr(err)
	    color:="green"
	    for i := range ifadmstat {
		ifadmstat[i].Name=strings.Replace(ifadmstat[i].Name, ".1.3.6.1.2.1.2.2.1.7.", "", 1)
		ifoperstat[i].Name=strings.Replace(ifoperstat[i].Name, ".1.3.6.1.2.1.2.2.1.8.", "", 1)
		if ifoperstat[i].Value != 1 {
		    color="#aa5555"
		}else {
		    color="#55aa55"
		}
		if ifType[iftype[i].Value.(int)-1]=="ethernetCsmacd" && ifspeed[i].Value.(int)/1000000 == 10 {
		    color="#55aaaa"
		}
		if ifStatus[ifadmstat[i].Value.(int)-1]!="up" {
		    color="#ff3333"
		}
		fmt.Fprintf(w,"<tr bgcolor=%s>",color)
//		fmt.Fprintf(w,"<td>%s</td>",umac)
		fmt.Fprintf(w,"<td>%s</td>",ifadmstat[i].Name)
		fmt.Fprintf(w,"<td>%s<br>",ifStatus[ifadmstat[i].Value.(int)-1])
		fmt.Fprintf(w,"<a href='#' style='color: yellow;' onclick='window.open(\"/swPortUp/%s/%s\");'>Вкл </a>",ip,ifadmstat[i].Name)
		fmt.Fprintf(w,"<a href='#' style='color: yellow;' onclick='window.open(\"/swPortDown/%s/%s\");'>Выкл</a>",ip,ifadmstat[i].Name)
		fmt.Fprintf(w,"</td>")
//		fmt.Fprintf(w,"<td>%s</td>",ifStatus[ifoperstat[i].Value.(int)-1])
		fmt.Fprintf(w,"<td><a href='#' onclick='window.open(\"/swCheckCable/%s/%s\");'>%s</a></td>",ip,ifadmstat[i].Name,ifStatus[ifoperstat[i].Value.(int)-1])
		fmt.Fprintf(w,"<td>%s</td>",ifname[i].Value)
		fmt.Fprintf(w,"<td><textarea id='txtArea%s' onkeypress='changeAlias(event,%s,this);'>",ifadmstat[i].Name,ifadmstat[i].Name)
		fmt.Fprintf(w,"%s</textarea></td>",ifalias[i].Value)
		fmt.Fprintf(w,"<td>%d</td>",ifspeed[i].Value.(int)/1000000)
		fmt.Fprintf(w,"<td>%s</td>",ifType[iftype[i].Value.(int)-1])
		fmt.Fprintf(w,"<td>%s</td>",fmt.Sprintf("%X",ifmac[i].Value))
		fmt.Fprintf(w,"<td>%d</td>",ifinerr[i].Value)
//		fmt.Fprintf(w,"<td>")
//		for j := range mcast14 {
//		    boid := strings.Split(mcast14[j].Name,".")
//		    if boid[20] == ifadmstat[i].Name {
//			fmt.Fprintf(w,"%s<br>",mcast14[j].Value)
//		    }
//		}
//		fmt.Fprintf(w,"</td>")
		fmt.Fprintf(w,"<td>")

		for j := range mcast16 {

		    cmip:=strings.Split(mcast16[j].Name,".")[18]
		    cmip=cmip+"."+strings.Split(mcast16[j].Name,".")[19]
		    cmip=cmip+"."+strings.Split(mcast16[j].Name,".")[20]
		    cmip=cmip+"."+strings.Split(mcast16[j].Name,".")[21]
		    tport:=strings.Split(mcast16[j].Name,".")[22]
		    if fmt.Sprintf("%d",i+1)==tport {
			fmt.Fprintf(w,"%s(%s)<br>",cmip,mcast16[j].Value)
		    }
		}

		for j := range mcast15 {
		    b0:=fmt.Sprintf("%X",mcast15[j].Value)[0:8]
		    b1:=fmt.Sprintf("%X",mcast15[j].Value)[8:16]
		    bi0,_ := strconv.ParseInt(b0, 16, 64)
		    bi1,_ := strconv.ParseInt(b1, 16, 64)
		    bs:=fmt.Sprintf("%32.32b%32.32b<br>",bi0,bi1)
		    cmip:=strings.Split(mcast15[j].Name,".")[20]
		    cmip=cmip+"."+strings.Split(mcast15[j].Name,".")[21]
		    cmip=cmip+"."+strings.Split(mcast15[j].Name,".")[22]
		    cmip=cmip+"."+strings.Split(mcast15[j].Name,".")[23]
		    if bs[i:i+1] == "1" {
			fmt.Fprintf(w,"%s<br>",cmip)
		    }
		    //fmt.Fprintf(w,"%s",bs)
		    //boid := strings.Split(mcast15[j].Name,".")
		    //if boid[20] == ifadmstat[i].Name {
		    //}
		}
		fmt.Fprintf(w,"</td>")
		fmt.Fprintf(w,"<td>\n")
		for j := range lldpchasid {
		    if fmt.Sprintf("%d",i+1) == strings.Split(lldpchasid[j].Name,".")[13] {
			fmt.Fprintf(w,"<table border=1><tr><td>ChassisID</td><td>%s</td><tr>",fmt.Sprintf("%X",lldpchasid[j].Value))
			if lldpporttype[j].Value == 3 {
			    fmt.Fprintf(w,"<tr><td>PortID</td><td>%s</td></tr>",fmt.Sprintf("%X",lldpport[j].Value))
			}else{
			    fmt.Fprintf(w,"<tr><td>PortID</td><td>%s</td></tr>",lldpport[j].Value)
			}
			fmt.Fprintf(w,"<tr><td>PortDescr</td><td>%s</td></tr>",lldpportdescr[j].Value)
			fmt.Fprintf(w,"<tr><td>SystemName</td><td>%s</td></tr>",lldpsysname[j].Value)
//			fmt.Fprintf(w,"%i",lldppeerip[0].Name)
			for k := range lldppeerip {
				if strings.Split(lldppeerip[k].Name,".")[13] == fmt.Sprintf("%d",i+1){
				lldppeeripc:=strings.Split(lldppeerip[k].Name,".")[17]+"."+strings.Split(lldppeerip[k].Name,".")[18]+"."+strings.Split(lldppeerip[k].Name,".")[19]+"."+strings.Split(lldppeerip[k].Name,".")[20]
				fmt.Fprintf(w,"<tr><td>PeerIP</td><td><a href=%s>%s</a></td></tr>",lldppeeripc,lldppeeripc)
			    }
			}
			fmt.Fprintf(w,"</table>")
			for k := range arptable {
			    ipr:=strings.Replace(arptable[k].Name,".1.3.6.1.2.1.4.22.1.2.","",1)
			    ipr=strings.Replace(ipr,strings.Split(ipr,".")[0]+".","",1)
			    if lldpchasid[j].Value == arptable[k].Value {
				fmt.Fprintf(w,"%s <a href=/checkIfaces/%s>%s</a><br>",fmt.Sprintf("%X",arptable[k].Value),ipr,ipr)
			    }
			}
		    }
		}
		for k := range fdbtable {
		    macr:=strings.Split(strings.Replace(fdbtable[k].Name,".1.3.6.1.2.1.17.7.1.2.2.1.2.","",1),".")
		    if ifadmstat[i].Name == fmt.Sprintf("%d",fdbtable[k].Value) {
			m1,_:=strconv.Atoi(macr[1])
			m2,_:=strconv.Atoi(macr[2])
			m3,_:=strconv.Atoi(macr[3])
			m4,_:=strconv.Atoi(macr[4])
			m5,_:=strconv.Atoi(macr[5])
			m6,_:=strconv.Atoi(macr[6])
			fmt.Fprintf(w,"Vlan:%s_Mac:",macr[0])
			macrs:=fmt.Sprintf("%2X%2X%2X%2X%2X%2X",m1,m2,m3,m4,m5,m6)
			macrs=strings.Replace(macrs," ","0",16)
			fmt.Fprintf(w,"%s<br>",macrs)
		    }
		}
		fmt.Fprintf(w,"</td>")
		fmt.Fprintf(w,"</tr>\n")
////		_, err = stmt.Exec(ifadmstat[i].Value,ifoperstat[i].Value,umac, ifadmstat[i].Name)
////		checkErr(err)
	    }
	    fmt.Fprintf(w,"</table><br>")
	    fmt.Fprintf(w,"</body></html>")
	    return
	}
    }
    fmt.Fprintf(w,"<h1>Хост не отвечает!</h1></body></html>")

}

func setSnmpStr(ip string, oid string, val string){
    for i := range snmp_coms_w {
	gosnmp2.Default.Target = ip
	gosnmp2.Default.Community = snmp_coms_w[i]
	err := gosnmp2.Default.Connect()
	if err != nil {
	    log.Fatalf("Connect() err: %v", err)
	}
	defer gosnmp2.Default.Conn.Close()
	var pdus []gosnmp2.SnmpPDU
	pdus = append(pdus, gosnmp2.SnmpPDU{oid, 4, val, nil})
	gosnmp2.Default.Set(pdus)
	gosnmp2.Default.Conn.Close()
    }
}

func setSnmpInt(ip string, oid string, val int){
    for i := range snmp_coms_w {
	gosnmp2.Default.Target = ip
	gosnmp2.Default.Community = snmp_coms_w[i]
	err := gosnmp2.Default.Connect()
	if err != nil {
	    log.Fatalf("Connect() err: %v", err)
	}
	defer gosnmp2.Default.Conn.Close()
	var pdus []gosnmp2.SnmpPDU
	pdus = append(pdus, gosnmp2.SnmpPDU{oid, 2, val, nil})
	gosnmp2.Default.Set(pdus)
	gosnmp2.Default.Conn.Close()
    }
}

func swChec1Cable(w http.ResponseWriter,ip string, port string){
    setSnmpInt(ip,".1.3.6.1.4.1.171.12.58.1.1.1.12."+port,1)
    for i := range snmp_coms_r{
		s, err := gosnmp.NewGoSNMP(ip, snmp_coms_r[i], gosnmp.Version2c, 1)
		checkErr(err)
		resp, err := s.Get(".1.3.6.1.4.1.171.12.58.1.1.1.12."+port)
		if err == nil{
//	    	fmt.Fprintf(w,"%d",resp.Variables[0].Value)
			fmt.Fprintf(w,"<tr><td>"+port+"</td>")
			resp,_ = s.Get(".1.3.6.1.4.1.171.12.58.1.1.1.3."+port)
			fmt.Fprintf(w,"<td>%d</td>",resp.Variables[0].Value)
			resp,_ = s.Get(".1.3.6.1.4.1.171.12.58.1.1.1.4."+port)
			fmt.Fprintf(w,"<td>%s</td>",cdStatus[resp.Variables[0].Value.(int)])
			resp,_ = s.Get(".1.3.6.1.4.1.171.12.58.1.1.1.5."+port)
			fmt.Fprintf(w,"<td>%s</td>",cdStatus[resp.Variables[0].Value.(int)])
			resp,_ = s.Get(".1.3.6.1.4.1.171.12.58.1.1.1.6."+port)
			fmt.Fprintf(w,"<td>%s</td>",cdStatus[resp.Variables[0].Value.(int)])
			resp,_ = s.Get(".1.3.6.1.4.1.171.12.58.1.1.1.7."+port)
			fmt.Fprintf(w,"<td>%s</td>",cdStatus[resp.Variables[0].Value.(int)])
			resp,_ = s.Get(".1.3.6.1.4.1.171.12.58.1.1.1.8."+port)
			fmt.Fprintf(w,"<td>%d</td>",resp.Variables[0].Value)
			resp,_ = s.Get(".1.3.6.1.4.1.171.12.58.1.1.1.9."+port)
			fmt.Fprintf(w,"<td>%d</td>",resp.Variables[0].Value)
			resp,_ = s.Get(".1.3.6.1.4.1.171.12.58.1.1.1.10."+port)
			fmt.Fprintf(w,"<td>%d</td>",resp.Variables[0].Value)
			resp,_ = s.Get(".1.3.6.1.4.1.171.12.58.1.1.1.11."+port)
			fmt.Fprintf(w,"<td>%d</td></tr>",resp.Variables[0].Value)
			return
		}
    }
}

func swCheckCable(w http.ResponseWriter, r *http.Request) {
    ip:=strings.Split(r.URL.Path[14:],"/")[0]
    port:=strings.Split(r.URL.Path[14:],"/")[1]
    fmt.Fprintf(w,"<table border=1><th>Port</th><th>Status</th><th>1</th><th>2</th><th>3</th><th>4</th><th>1</th><th>2</th><th>3</th><th>4</th>")
	swChec1Cable(w,ip,port)
    fmt.Fprintf(w,"</table>")
	return
}

func swCheckAllCable(w http.ResponseWriter, r *http.Request) {
    ip:=strings.Split(r.URL.Path[17:],"/")[0]
    fmt.Fprintf(w,"<table border=1><th>Port</th><th>Status</th><th>1</th><th>2</th><th>3</th><th>4</th><th>1</th><th>2</th><th>3</th><th>4</th>")
	swChec1Cable(w,ip,"1")
	swChec1Cable(w,ip,"2")
	swChec1Cable(w,ip,"3")
	swChec1Cable(w,ip,"4")
	swChec1Cable(w,ip,"5")
	swChec1Cable(w,ip,"6")
	swChec1Cable(w,ip,"7")
	swChec1Cable(w,ip,"8")
	swChec1Cable(w,ip,"9")
	swChec1Cable(w,ip,"10")
	swChec1Cable(w,ip,"11")
	swChec1Cable(w,ip,"12")
	swChec1Cable(w,ip,"13")
	swChec1Cable(w,ip,"14")
	swChec1Cable(w,ip,"15")
	swChec1Cable(w,ip,"16")
	swChec1Cable(w,ip,"17")
	swChec1Cable(w,ip,"18")
	swChec1Cable(w,ip,"19")
	swChec1Cable(w,ip,"20")
	swChec1Cable(w,ip,"21")
	swChec1Cable(w,ip,"22")
	swChec1Cable(w,ip,"23")
	swChec1Cable(w,ip,"24")
    fmt.Fprintf(w,"</table>")
	return
}

func swSetAlias(w http.ResponseWriter, r *http.Request) {
    ip:=strings.Split(r.URL.Path[12:],"/")[0]
    port:=".1.3.6.1.2.1.31.1.1.1.18."+strings.Split(r.URL.Path[12:],"/")[1]
    txt:=strings.Split(r.URL.Path[12:],"/")[2]
    fmt.Fprintf(w,"<html><body onload='window.close();'></body></html>")
    setSnmpStr(ip,port,txt)
}

func swPortUp(w http.ResponseWriter, r *http.Request) {
    ip:=strings.Split(r.URL.Path[10:],"/")[0]
    port:="1.3.6.1.2.1.2.2.1.7."+strings.Split(r.URL.Path[10:],"/")[1]
    fmt.Fprintf(w,"<html><body onload='window.close();'></body></html>")
    setSnmpInt(ip,port,1)
}

func swPortDown(w http.ResponseWriter, r *http.Request) {
    ip:=strings.Split(r.URL.Path[12:],"/")[0]
    port:="1.3.6.1.2.1.2.2.1.7."+strings.Split(r.URL.Path[12:],"/")[1]
    fmt.Fprintf(w,"<html><body onload='window.close();'></body></html>")
    setSnmpInt(ip,port,2)
}

//func swFdbTable(w http.ResponseWriter, r *http.Request) {
//}

func swArpTable(w http.ResponseWriter, r *http.Request) {
    ip:=r.URL.Path[12:]
    for i := range snmp_coms_r{
	s, err := gosnmp.NewGoSNMP(ip, snmp_coms_r[i], gosnmp.Version2c, 1)
	checkErr(err)
	_ , err = s.Get(".1.3.6.1.2.1.17.1.1.0")
	if err == nil{
	    arptable, _ := s.Walk(".1.3.6.1.2.1.4.22.1.2")
	    checkErr(err)
	    fmt.Fprintf(w,"<table>")
	    for k := range arptable {
		ipr:=strings.Replace(arptable[k].Name,".1.3.6.1.2.1.4.22.1.2.","",1)
		ipr=strings.Replace(ipr,strings.Split(ipr,".")[0]+".","",1)
		fmt.Fprintf(w,"<tr><td>%s</td><td><a href=/checkIfaces/%s>%s</a></td></tr>",fmt.Sprintf("%X",arptable[k].Value),ipr,ipr)
	    }
	    fmt.Fprintf(w,"</table>")
	    return
	}
    }
}

func swMainData(w http.ResponseWriter, r *http.Request) {
    ip:=r.URL.Path[12:]
    for i := range snmp_coms_r{
	s, err := gosnmp.NewGoSNMP(ip, snmp_coms_r[i], gosnmp.Version2c, 1)
	checkErr(err)
	resp , err := s.Get(".1.3.6.1.2.1.17.1.1.0")
	if err == nil{
	    umac := fmt.Sprintf("%X",resp.Variables[0].Value)
	    //fmt.Fprintf(w,"<html><head></head><body>%s<br>\n",snmp_coms_r[i])
	    fmt.Fprintf(w,"<table border=1><tr><td>MAC Address</td><td>%s</td></tr>\n",umac)
	    resp, _ = s.Get(".1.3.6.1.2.1.1.1.0")
	    fmt.Fprintf(w,"<tr><td>Device Type</td><td>%s</td></tr>\n",resp.Variables[0].Value.(string))
	    resp, _ = s.Get(".1.3.6.1.2.1.16.19.2.0")
	    fmt.Fprintf(w,"<tr><td>Firmware Version</td><td>%s</td></tr>\n",resp.Variables[0].Value.(string))
	    resp, _ = s.Get(".1.3.6.1.2.1.16.19.3.0")
	    fmt.Fprintf(w,"<tr><td>Hardware Version</td><td>%s</td></tr>\n",resp.Variables[0].Value.(string))
	    resp, _ = s.Get(".1.3.6.1.2.1.1.3.0")
	    fmt.Fprintf(w,"<tr><td>System Uptime</td><td>%d</td></tr>\n",resp.Variables[0].Value)
	    resp, _ = s.Get(".1.3.6.1.2.1.1.4.0")
	    fmt.Fprintf(w,"<tr><td>System Contact</td><td>%s</td></tr>\n",resp.Variables[0].Value.(string))
	    resp, _ = s.Get(".1.3.6.1.2.1.1.5.0")
	    fmt.Fprintf(w,"<tr><td>System Name</td><td>%s</td></tr>\n",resp.Variables[0].Value.(string))
	    resp, _ = s.Get(".1.3.6.1.2.1.1.6.0")
	    fmt.Fprintf(w,"<tr><td>System Location</td><td>%s</td></tr>\n",resp.Variables[0].Value.(string))
	    resp, _ = s.Get(".1.3.6.1.4.1.171.12.15.2.3.1.21.0")

	    stpt := time.Unix(int64(resp.Variables[0].Value.(int)/100),0)
	    fmt.Fprintf(w,"<tr><td>STP Last Change</td><td>%d %d:%d:%d</td></tr>\n",stpt.Day()-1,stpt.Hour(),stpt.Minute(),stpt.Second())

	    fmt.Fprintf(w,"</table>")
	    return
	}
    }

}

func ospfIpNeib(w http.ResponseWriter, r *http.Request) {
    ip:=r.URL.Path[12:]
    for i := range snmp_coms_r{
	s, err := gosnmp.NewGoSNMP(ip, snmp_coms_r[i], gosnmp.Version2c, 1)
	checkErr(err)
	_ , err = s.Get(".1.3.6.1.2.1.17.1.1.0")
	if err == nil{
	    fmt.Fprintf(w,"<table border=1>")
	    ospfIpNeib, err := s.Walk(".1.3.6.1.2.1.14.10.1.1")
	    checkErr(err)
	    for i := range ospfIpNeib {
		fmt.Fprintf(w,"<tr><td><a href=/checkIfaces/%s>%s</a></td></tr>",ospfIpNeib[i].Value,ospfIpNeib[i].Value)
	    }
	    fmt.Fprintf(w,"</table>")
	    return
	}
    }
}

func rd(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w,"<html><head><meta http-equiv=\"refresh\" content=\"0; url=/checkIfaces/")
    fmt.Fprintf(w,"%s",strings.Replace(strings.Replace(r.URL.Path[4:],"http:","",1),"/","",1))
    fmt.Fprintf(w,"\"/></head><body></body></html>")
}

type handler func(w http.ResponseWriter, r *http.Request)
func BasicAuth(pass handler) handler {
    return func(w http.ResponseWriter, r *http.Request) {
	_,a:=r.Header["Authorization"]
	if !a {
	    w.Header().Set("WWW-Authenticate", "Basic realm=\"Password protected area\"")
	    w.WriteHeader(401)
	    fmt.Fprintf(w,"<?xml version=\"1.0\" encoding=\"iso-8859-1\"?>\n")
	    fmt.Fprintf(w,"<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\"\n")
	    fmt.Fprintf(w,"\"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\">\n")
	    fmt.Fprintf(w,"<html xmlns=\"http://www.w3.org/1999/xhtml\" xml:lang=\"en\" lang=\"en\">\n")
	    fmt.Fprintf(w,"<head><title>401 - Unauthorized</title></head><body><h1>401 - Unauthorized</h1></body></html>")
	    return 
	}
	auth := strings.SplitN(r.Header["Authorization"][0], " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
	    http.Error(w, "bad syntax", http.StatusBadRequest)
	    return
	}
	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 || !Validate(pair[0], pair[1]) {
	    http.Error(w, "authorization failed", http.StatusUnauthorized)
	    return
	}
	pass(w, r)
    }
}

func Validate(username, password string) bool {
    if username == http_user && password == http_password {
        return true
    }
    return false
}

func mainPage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><head></head><body>\n")
    fmt.Fprintf(w, "<textarea id='goip' onkeypress='var code = (event.keyCode ? event.keyCode : event.which); if (code == 13) {window.open(\"/checkIfaces/\"+this.value);}'>127.0.0.1</textarea><br>\n")
//    fmt.Fprintf(w, "<a href=/checkIfaces/217.107.196.33>217.107.196.33</a><br>\n")
//    fmt.Fprintf(w, "<a href=/checkIfaces/10.39.3.50>10.39.3.50</a><br>\n")
//    fmt.Fprintf(w, "<a href=/checkIfaces/10.45.111.108>10.45.111.108</a><br>\n")
    fmt.Fprintf(w, "</body></html>\n")
//    db, err := sql.Open("mysql", mysql_user+":"+mysql_password+"@tcp("+mysql_host+":"+mysql_port+")/"+mysql_db+"?charset=utf8")
//    checkErr(err)
//    rows, err := db.Query("SELECT * FROM sw_ip")
//    checkErr(err)
//    for rows.Next() {
//	var umac int
//	var ip string
//	var if_id int
//	err = rows.Scan(&umac, &ip, &if_id)
//	checkErr(err)
//	fmt.Fprintf(w, "%d %s %d\n", umac,ip,if_id)
//	fmt.Print(umac)
//	fmt.Print(ip)
//	fmt.Println(if_id)
//    }
//    s, err := gosnmp.NewGoSNMP(r.URL.Path[1:], "public", gosnmp.Version2c, 5)
//    checkErr(err)
//    resp, err := s.Walk(".1.3.6.1.2.1.2.2.1.2")
//    checkErr(err)
//    for i := range resp {
//	resp[i].Name=strings.Replace(resp[i].Name, ".1.3.6.1.2.1.2.2.1.2.", "", 1)
//	fmt.Println(fmt.Sprintf("%s '%s'",resp[i].Name,resp[i].Value))
//	fmt.Fprintf(w, "%s '%s'\n", resp[i].Name,resp[i].Value)
//    }
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

////var mysql_host="217.107.196.60"
////var mysql_port="3306"
////var mysql_user="foxpdll"
////var mysql_password="qwe123"
////var mysql_db="foxnms"
var snmp_coms_r=[...]string{"public", "readzt"}
var snmp_coms_w=[...]string{"private", "wrzt"}
var http_user="tech"
var http_password="tech123"



func main() {
    flag.Parse()
    fmt.Println("Address:",*addr)
    http.HandleFunc("/checkIfaces/", checkIfaces)
    http.HandleFunc("/ospfIpNeib/", ospfIpNeib)
    http.HandleFunc("/swMainData/", swMainData)
    http.HandleFunc("/swArpTable/", swArpTable)
    http.HandleFunc("/swPortUp/", swPortUp)
    http.HandleFunc("/swPortDown/", swPortDown)
    http.HandleFunc("/swSetAlias/", swSetAlias)
    http.HandleFunc("/swCheckCable/", swCheckCable)
    http.HandleFunc("/swCheckAllCable/", swCheckAllCable)
    http.HandleFunc("/rd/", rd)
    http.HandleFunc("/",BasicAuth(mainPage))
//    http.HandleFunc("/",mainPage)
    http.ListenAndServe(*addr, nil)
}
