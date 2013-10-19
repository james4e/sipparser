package sipparser

// Imports from the go standard library
import (
	//"fmt"
	"testing"
)

func TestHeader(t *testing.T) {
	h := Header{"t", "v"}
	if h.String() != "t: v" {
		t.Errorf("Error with header.String() method.  Unexpected res.")
	}
}

// makes sure that the body is what is expected
func TestBody(t *testing.T) {
	s := ParseMsg("fake\r\nheader\r\n\r\nbody ...\r\n\r\n")
	if s.Body != "body ...\r\n\r\n" {
		t.Errorf("[TestBody] Error getting the right body from the string.")
	}
}

// actual msg testing
func TestParseMsg(t *testing.T) {
	m := "SIP/2.0 200 OK\r\nVia: SIP/2.0/UDP 0.0.0.0:5060;branch=z9hG4bK24477ab511325213INV52e94be64e6687e3;received=0.0.0.0\r\nContact: <sip:10003053258853@0.0.0.0:6060>\r\nTo: <sip:10003053258853@0.0.0.0;user=phone;noa=national>;tag=a94c095b773be1dd6e8d668a785a9c843f6f2cc0\r\nFrom: <sip:8173383772@0.0.0.0;user=phone;noa=national>;tag=52e94be6-co2998-INS002\r\nCall-ID: 111118149-3524331107-398662@barinfo.fooinfous.com\r\nCSeq: 299801 INVITE\r\nAccept: application/sdp, application/dtmf-relay, text/plain\r\nX-Nonsense-Hdr: nonsense\r\nAllow: PRACK, INVITE, BYE, REGISTER, ACK, OPTIONS, CANCEL, SUBSCRIBE, NOTIFY, INFO, REFER, UPDATE\r\nContent-Type: application/sdp\r\nServer: Dialogic-SIP/10.5.3.231 IMGDAL0001 0\r\nSupported: 100rel, path, replaces, timer, tdialog\r\nContent-Length: 239\r\n\r\nv=0\r\no=Dialogic_SDP 1452654 0 IN IP4 0.0.0.0\r\ns=Dialogic-SIP\r\nc=IN IP4 4.71.122.135\r\nt=0 0\r\nm=audio 11676 RTP/AVP 0 101\r\na=rtpmap:0 PCMU/8000\r\na=rtpmap:101 telephone-event/8000\r\na=fmtp:101 0-15\r\na=silenceSupp:off - - - -\r\na=ptime:20\r\n\r\n"
	s := ParseMsg(m)
	if s.Error != nil {
		t.Errorf("[TestParseMsg] Error parsing msg. Recevied: " + s.Error.Error())
	}
	if len(s.Body) == 0 {
		t.Errorf("[TestParseMsg] Error parsing msg.  Body should have a length.")
	}
	if len(s.Headers) == 0 {
		t.Errorf("[TestParseMsg] Error parsing msg.  Does not appear to be any headers.")
	}
	if s.Via == nil || len(s.Via) == 0 {
		t.Errorf("[TestParseMsg] Error parsing msg.  Does not appear to be any vias parsed.")
		//fmt.Println("msg:", s.Msg)
		//fmt.Println("body:", s.Body)
		//fmt.Println("via:", s.Via)
		//fmt.Println("crlf:", s.crlf)
	}
	if s.ContentLength != "239" {
		t.Errorf("[TestParseMsg] Error parsing msg.  Content length should be 239.  Received: " + s.ContentLength)
	}
	if len(s.Supported) != 5 {
		t.Errorf("[TestParseMsg] Error parsing msg.  s.Support should have length of 5.")
	}
}

func BenchmarkParseMsg(b *testing.B) {
	testInviteMsg := "INVITE sip:15554440000@X.X.X.X:5060;user=phone SIP/2.0\r\nVia: SIP/2.0/UDP X.X.X.X:5060;branch=z9hG4bK34133a599ll241207INV21d7d0684e84a2d2\r\nMax-Forwards: 35\r\nContact: <sip:X.X.X.X:5060>\r\nTo: <sip:15554440000@X.X.X.X;user=phone;noa=national>\r\nFrom: \"Unavailable\"<sip:X.X.X.X;user=phone;noa=national>;tag=21d7d068-co2149-FOOI003\r\nCall-ID: 1393184968_47390262@domain.com\r\nCSeq: 214901 INVITE\r\nAllow: INVITE,ACK,CANCEL,BYE,REFER,OPTIONS,NOTIFY,SUBSCRIBE,PRACK,INFO\r\nContent-Type: application/sdp\r\nDate: Thu, 29 Sep 2011 16:54:42 GMT\r\nUser-Agent: FAKE-UA-DATA\r\nP-Asserted-Identity: \"Unavailable\"<sip:Restricted@X.X.X.X:5060>\r\nContent-Length: 322\r\n\r\nv=0\r\no=- 567791720 567791720 IN IP4 X.X.X.X\r\ns=FAKE-DATA\r\nc=IN IP4 X.X.X.X\r\nt=0 0\r\nm=audio 17354 RTP/AVP 0 8 86 18 96\r\na=rtpmap:0 PCMU/8000\r\na=rtpmap:8 PCMA/8000\r\na=rtpmap:86 G726-32/8000\r\na=rtpmap:18 G729/8000\r\na=rtpmap:96 telephone-event/8000\r\na=maxptime:20\r\na=fmtp:18 annexb=yes\r\na=fmtp:96 0-15\r\na=sendrecv\r\n"
	testTryingMsg := "SIP/2.0 100 Giving a try\r\nVia: SIP/2.0/UDP X.X.X.X:5060;branch=z9hG4bK24477ab511bbebdbINV4eed095f4e84a4c4\r\nTo: <sip:15554440000@2.2.2.2;user=phone;noa=national>\r\nFrom: \"Unknown\"<sip:5556661000@X.X.X.X;user=phone;noa=national>;tag=4eed095f-co11196-FOS002\r\nCall-ID: ba317db9-655f-122f-52a3-0015c5ee608e\r\nCSeq: 1119601 INVITE\r\nUser-Agent: FAKE-UA-DATA\r\nContent-Length: 0\r\n\r\n"
	testOptionsMsg := "OPTIONS sip:some_vendor@X.X.X.X:5060;transport=udp SIP/2.0\r\nVia: SIP/2.0/UDP X.X.X.X:5060;rport;branch=z9hG4bK-d447b18336e8e27b7de04bdc143c043c-alias.foo-info.net-1\r\nAllow-Events: message-summary, refer, dialog, line-seize, presence, call-info, as-feature-event\r\nMax-Forwards: 70\r\nCall-ID: A1BF42F3@alias.foo-info.net\r\nFrom: <sip:some_vendor@X.X.X.X:5060>;tag=alias.foo-info.net+1+0+fde1e5dc\r\nCSeq: 177329608 OPTIONS\r\nOrganization:\r\nSupported: 100rel, resource-priority\r\nContent-Length: 0\r\nContact: <sip:some_vendor@X.X.X.X:5060>\r\nTo: <sip:some_vendor@X.X.X.X>\r\n\r\n"
	testSessionProgressMsg := "SIP/2.0 183 Session Progress\r\nVia: SIP/2.0/UDP X.X.X.X:5060;branch=z9hG4bK24477ab51139bd4bINV66a300a74e861f52;received=X.X.X.X\r\nCall-ID: 889323860_112911699@X.X.X.X\r\nFrom: <sip:8885551000@X.X.X.X;user=phone;noa=national>;tag=66a300a7-co3440-FOS002\r\nTo: <sip:5554441000@X.X.X.X;user=phone;noa=national>;tag=a94c095b773be1dd6e8d668a785a9c8415431228\r\nContact: <sip:5554441000@X.X.X.X:6060>\r\nCSeq: 344001 INVITE\r\nAllow: OPTIONS, CANCEL, UPDATE\r\nServer: Dialogic-SIP/2.2.2.2 IMGDAL0002 1\r\nContent-Type: application/sdp\r\nContent-Length: 239\r\n\r\nv=0\r\no=Dialogic_SDP 5877551 0 IN IP4 X.X.X.X\r\ns=Dialogic-SIP\r\nc=IN IP4 X.X.X.X\r\nt=0 0\r\nm=audio 11792 RTP/AVP 0 101\r\na=rtpmap:0 PCMU/8000\r\na=rtpmap:101 telephone-event/8000\r\na=fmtp:101 0-15\r\na=silenceSupp:off - - - -\r\na=ptime:20\r\n"
	testAckMsg := "ACK sip:18885551000@X.X.X.X:5060;user=phone SIP/2.0\r\nVia: SIP/2.0/UDP X.X.X.X:5060;branch=z9hG4bK34133a599113fd0c3INV290f15484e861d60\r\nTo: <sip:18885551000@X.X.X.X;user=phone;noa=national>;tag=526fa7f1-co3633-FOS002\r\nFrom: \"FOOBARINFO\"<sip:X.X.X.X;user=phone;noa=national>;tag=290f1548-co3802-FOOI003\r\nCall-ID: 218256626-3526401492-574930@foo.bar.com\r\nCSeq: 380201 ACK\r\nContent-Length: 0\r\n"
	testOKMsg := "SIP/2.0 200 OK\r\nVia: SIP/2.0/UDP X.X.X.X:5060;branch=z9hG4bK24477ab5ll45a3c7BYE558eface4e861ebc\r\nTo: \"8885551000\"<sip:8885551000@X.X.X.X:5060>;tag=56e1439e-co6299-FOOI007\r\nFrom: <sip:15554441000@X.X.X.X:5060>;tag=558f0659-co4149-FOS002\r\nCall-ID: 5d1d1cc6753f94863fad610b011c94c9@foo.bar.com\r\nCSeq: 414901 BYE\r\nUser-Agent: FOO-UA-DATA\r\nContent-Length: 0\r\n\r\n"
	testByeMsg := "BYE sip:5554441000@X.X.X.X:5060 SIP/2.0\r\nVia: SIP/2.0/UDP X.X.X.X:5060;branch=z9hG4bK34133a599mm3b558aBYE60e40fcb4e8872ba\r\nMax-Forwards: 35\r\nTo: <sip:5554441000@X.X.X.X;user=phone;noa=national>;tag=747a9240-co11617-FOOI002\r\nFrom: <sip:18884442000@X.X.X.X;user=phone;isup-oli=00;noa=national>;tag=60e40fcb-co3535-FOOI003\r\nCall-ID: CDEA230-32F333@foo.bar.com\r\nCSeq: 353502 BYE\r\nUser-Agent: FAKE-UA-DATA\r\nContent-Length: 0\r\n\r\n"
	testCancelMsg := "CANCEL sip:5554448000@X.X.X.X:5060;user=phone SIP/2.0\r\nVia: SIP/2.0/UDP X.X.X.X:5060;branch=z9hG4bK24477ab5115bcff3INVdf0e3ed4e8875ec\r\nMax-Forwards: 35\r\nTo: <sip:5554448000@X.X.X.X;user=phone;noa=national>\r\nFrom: <sip:18885552000@X.X.X.X;user=phone;noa=national>;tag=df0e3ed-co5470-FOOI002\r\nCall-ID: 50d49946ed0311e0830100151702a864@X.X.X.X\r\nCSeq: 547001 CANCEL\r\nUser-Agent: FAKE-UA-DATA\r\nContent-Length: 0\r\n\r\n"
	for i:=0;i<10000;i++{
	ParseMsg(testInviteMsg)
	ParseMsg(testTryingMsg)
	ParseMsg(testOptionsMsg)
	ParseMsg(testSessionProgressMsg)
	ParseMsg(testAckMsg)
	ParseMsg(testOKMsg)
	ParseMsg(testByeMsg)
	ParseMsg(testCancelMsg)
	}
}

// testing the GetCallingParty functionality
func TestGetCallingParty(t *testing.T) {
	rpid := "\"UNKNOWN\" <sip:8885551000@0.0.0.0>;party=calling;screen=yes;privacy=off"
	s := &SipMsg{RemotePartyIdVal: rpid}
	err := s.GetCallingParty(CALLING_PARTY_RPID)
	if err != nil {
		t.Errorf("[TestGetCallingParty] Err with GetCallingParty.  rcvd: " + err.Error())
	}
	if s.CallingParty == nil {
		t.Errorf("[TestGetCallingParty] Err calling GetCallingParty.  CallingParty field should not be nil.")
	}
	if s.CallingParty.Name != "UNKNOWN" {
		t.Errorf("[TestGetCallingParty] Err calling GetCallingParty.  Name should be \"UNKNOWN\".")
	}
	if s.CallingParty.Number != "8885551000" {
		t.Errorf("[TestGetCallingParty] Err with GetCallingParty. Number should be \"8885551000\".")
	}
	paid := "<sip:8884441000@0.0.0.0:5060;user=phone>"
	s = &SipMsg{PAssertedIdVal: paid}
	err = s.GetCallingParty(CALLING_PARTY_PAID)
	if err != nil {
		t.Errorf("[TestGetCallingParty] Err with GetCallingParty on paid.  rcvd: " + err.Error())
	}
	if s.CallingParty == nil {
		t.Errorf("[TestGetCallingParty] Err with GetCallingParty on paid.  No CallingPartyInfo.")
	}
	if s.CallingParty.Name != "" {
		t.Errorf("[TestGetCallingParty] Err with GetCallingParty on paid.  Name should be \"\".")
	}
	if s.CallingParty.Number != "8884441000" {
		t.Errorf("[TestGetCallingParty] Err with GetCallingParty on paid.  Number should be \"8884441000\".")
	}
	s = &SipMsg{}
	s.parseFrom("\"5556661000\" <sip:5556661000@0.0.0.0>;tag=ZN21rHN5B7U0K")
	err = s.GetCallingParty("")
	if err != nil {
		t.Errorf("[TestGetCallingParty] Err calling GetCallingParty on default.  rcvd: " + err.Error())
	}
	if s.CallingParty == nil {
		t.Errorf("[TestGetCallingParty] Err calling GetCallingParty on default.  No CallingPartyInfo.")
	}
	if s.CallingParty.Name != "5556661000" {
		t.Errorf("[TestGetCallingParty] Err calling GetCallingParty on default.  Name should be \"5556661000\".")
	}
	if s.CallingParty.Number != "5556661000" {
		t.Errorf("[TestGetCallingParty] Err calling GetCallingParty on default. Number should be \"5556661000\".")
	}
}
