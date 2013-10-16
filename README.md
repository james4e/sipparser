sipparser
=========
[![Build Status](https://drone.io/github.com/danielvargas/sipparser/status.png)](https://drone.io/github.com/danielvargas/sipparser/latest) [![Go Walker](http://gowalker.org/api/v1/badge)](http://gowalker.org/github.com/danielvargas/sipparser)
###About

sipparser is a high performce parser for Session Initiated
Protocol messages based on [sip_parser](https://bitbucket.org/sdr/sip_parser).  It provides a library for use in building 
SIP user agents or any program in go that needs to be able to
parse SIP messages.

###Installation

(the following step assume that you have git and go
installed already)
```shell
git get github.com/danielvargas/sipparser
```

###Usage

The library has an easy to use interface.  

* call sipparser.ParseMsg(msg string)
* you'll get back a *SipMsg struct with the following:
    + State is the last parsing state
    + Error is an os.Error
    + Msg is the raw msg
    + CallingParty is a *CallingParty struct 
    + Body is the body of the message
    + StartLine is the parsed StartLine 
    + Headers is a slice of *Headers and will only contain headers that do not get parsed 
    + Accept is a *Accept struct 
    + AlertInfo is just the string of the Alert-Info hdr
    + Allow is a slice of strings of the methods that are allowed
    + AllowEvents is a slice of strings of the supported event types
    + ContentDisposition is a *ContentDisposition struct 
    + ContentLength is the value of the Content-Length hdr
    + ContentLengthInt is the int value of the ContentLength
    + ContentType is the header value for the Content-Type hdr
    +  From is a *From struct 
    + MaxForwards is the hdr value for the Max-Forwards hdr
    + MaxForwardsInt is the int value of the MaxForwards field
    + Organization is the value for the Organization hdr
    + To is a *From struct
    + Contact is a *From struct ** NOTE ** Contact is not parsed automatically. You have to call *SipMsg.ParseContact() to get this value.
    + ContactVal is the raw value of the contact hdr
    + CallId is the call-id for the message
    + Cseq is a *Cseq struct
    + Rack is a *Rack struct
    + Reason is a *Reason struct
    + Rseq is the value of the RSeq hdr
    + RseqInt is the int value of the Rseq
    + RecordRoute is a slice of *URI's structs
    + Route is a slice of *URI's structs 
    + Via is a slice of *Via structs
    + Require is a slice of the required extensions (string values) from the Require hdr
    + Supported is a slice of the supported extensions (string values) from the Supported hdr
    + Privacy is the value of the Privacy hdr
    + ProxyRequire is a slice of strings from the Proxy-Require hdr
    + RemotePartyIdVal is the value from the Remote-Party-Id hdr
    + RemotePartyId is the RemotePartyId struct ** NOTE ** In order to actually get this value you have to call *SipMsg.ParseRemotePartyId()
    + PAssertedIdVal is the value from the P-Asserted-Identity hdr
    + PAssertedId is the *PAssertedId struct ** NOTE ** In order to actually get this value you have to call *SipMsg.ParsePAssertedId()
    + Unsupported is a slice of the unsupported extensions from the Unsupported hdr
    + UserAgent is the value of the User-Agent hdr
    + Server is the value of the Server hdr
    + Subject is the value of the Subject hdr
    + Warning is a *Warning struct

###SIP grammar from RFC3261
```
   alphanum  =  ALPHA / DIGIT

   reserved    =  ";" / "/" / "?" / ":" / "@" / "&" / "=" / "+ "
                  / "$" / ","
   unreserved  =  alphanum / mark
   mark        =  "-" / "_" / "." / "!" / "~" / "*" / "'"
                  / "(" / ")"
   escaped     =  "%" HEXDIG HEXDIG

   LWS  =  [*WSP CRLF] 1*WSP ; linear whitespace
   SWS  =  [LWS] ; sep whitespace

   HCOLON  =  *( SP / HTAB ) ":" SWS

   TEXT-UTF8-TRIM  =  1*TEXT-UTF8char *(*LWS TEXT-UTF8char)
   TEXT-UTF8char   =  %x21-7E / UTF8-NONASCII
   UTF8-NONASCII   =  %xC0-DF 1UTF8-CONT
                   /  %xE0-EF 2UTF8-CONT
                   /  %xF0-F7 3UTF8-CONT
                   /  %xF8-Fb 4UTF8-CONT
                   /  %xFC-FD 5UTF8-CONT
   UTF8-CONT       =  %x80-BF

   LHEX  =  DIGIT / %x61-66 ;lowercase a-f

   token       =  1*(alphanum / "-" / "." / "!" / "%" / "*"
                  / "_" / "+ " / "`" / "'" / "~" )
   separators  =  "(" / ")" / "<" / ">" / "@" /
                  "," / ";" / ":" / "\" / DQUOTE /
                  "/" / "[" / "]" / "?" / "=" /
                  "{" / "}" / SP / HTAB
   word        =  1*(alphanum / "-" / "." / "!" / "%" / "*" /
                  "_" / "+ " / "`" / "'" / "~" /
                  "(" / ")" / "<" / ">" /
                  ":" / "\" / DQUOTE /
                  "/" / "[" / "]" / "?" /
                  "{" / "}" )

   STAR    =  SWS "*" SWS ; asterisk
   SLASH   =  SWS "/" SWS ; slash
   EQUAL   =  SWS "=" SWS ; equal
   LPAREN  =  SWS "(" SWS ; left parenthesis
   RPAREN  =  SWS ")" SWS ; right parenthesis
   RAQUOT  =  ">" SWS ; right angle quote
   LAQUOT  =  SWS "<"; left angle quote
   COMMA   =  SWS "," SWS ; comma
   SEMI    =  SWS ";" SWS ; semicolon
   COLON   =  SWS ":" SWS ; colon
   LDQUOT  =  SWS DQUOTE; open double quotation mark
   RDQUOT  =  DQUOTE SWS ; close double quotation mark

   comment  =  LPAREN *(ctext / quoted-pair / comment) RPAREN
   ctext    =  %x21-27 / %x2A-5B / %x5D-7E / UTF8-NONASCII
               / LWS

   quoted-string  =  SWS DQUOTE *(qdtext / quoted-pair ) DQUOTE
   qdtext         =  LWS / %x21 / %x23-5B / %x5D-7E
                     / UTF8-NONASCII

   quoted-pair  =  "\" (%x00-09 / %x0B-0C
                   / %x0E-7F)

   SIP-URI          =  "sip:" [ userinfo ] hostport
                       uri-parameters [ headers ]
   SIPS-URI         =  "sips:" [ userinfo ] hostport
                       uri-parameters [ headers ]
   userinfo         =  ( user / telephone-subscriber ) [ ":" password ] "@"
   user             =  1*( unreserved / escaped / user-unreserved )
   user-unreserved  =  "&" / "=" / "+ " / "$" / "," / ";" / "?" / "/"
   password         =  *( unreserved / escaped /
                       "&" / "=" / "+ " / "$" / "," )
   hostport         =  host [ ":" port ]
   host             =  hostname / IPv4address / IPv6reference
   hostname         =  *( domainlabel "." ) toplabel [ "." ]
   domainlabel      =  alphanum
                       / alphanum *( alphanum / "-" ) alphanum
   toplabel         =  ALPHA / ALPHA *( alphanum / "-" ) alphanum
   IPv4address    =  1*3DIGIT "." 1*3DIGIT "." 1*3DIGIT "." 1*3DIGIT
   IPv6reference  =  "[" IPv6address "]"
   IPv6address    =  hexpart [ ":" IPv4address ]
   hexpart        =  hexseq / hexseq "::" [ hexseq ] / "::" [ hexseq ]
   hexseq         =  hex4 *( ":" hex4)
   hex4           =  1*4HEXDIG
   port           =  1*DIGIT

   uri-parameters    =  *( ";" uri-parameter)
   uri-parameter     =  transport-param / user-param / method-param
                        / ttl-param / maddr-param / lr-param / other-param
   transport-param   =  "transport="
                        ( "udp" / "tcp" / "sctp" / "tls"
                        / other-transport)
   other-transport   =  token
   user-param        =  "user=" ( "phone" / "ip" / other-user)
   other-user        =  token
   method-param      =  "method=" Method
   ttl-param         =  "ttl=" ttl
   maddr-param       =  "maddr=" host
   lr-param          =  "lr"
   other-param       =  pname [ "=" pvalue ]
   pname             =  1*paramchar
   pvalue            =  1*paramchar
   paramchar         =  param-unreserved / unreserved / escaped
   param-unreserved  =  "[" / "]" / "/" / ":" / "&" / "+ " / "$"

   headers         =  "?" header *( "&" header )
   header          =  hname "=" hvalue
   hname           =  1*( hnv-unreserved / unreserved / escaped )
   hvalue          =  *( hnv-unreserved / unreserved / escaped )
   hnv-unreserved  =  "[" / "]" / "/" / "?" / ":" / "+ " / "$"

   SIP-message    =  Request / Response
   Request        =  Request-Line
                     *( message-header )
                     CRLF
                     [ message-body ]
   Request-Line   =  Method SP Request-URI SP SIP-Version CRLF
   Request-URI    =  SIP-URI / SIPS-URI / absoluteURI
   absoluteURI    =  scheme ":" ( hier-part / opaque-part )
   hier-part      =  ( net-path / abs-path ) [ "?" query ]
   net-path       =  "//" authority [ abs-path ]
   abs-path       =  "/" path-segments
   opaque-part    =  uric-no-slash *uric
   uric           =  reserved / unreserved / escaped
   uric-no-slash  =  unreserved / escaped / ";" / "?" / ":" / "@"
                     / "&" / "=" / "+ " / "$" / ","
   path-segments  =  segment *( "/" segment )
   segment        =  *pchar *( ";" param )
   param          =  *pchar
   pchar          =  unreserved / escaped /
                     ":" / "@" / "&" / "=" / "+ " / "$" / ","
   scheme         =  ALPHA *( ALPHA / DIGIT / "+ " / "-" / "." )
   authority      =  srvr / reg-name
   srvr           =  [ [ userinfo "@" ] hostport ]
   reg-name       =  1*( unreserved / escaped / "$" / ","
                     / ";" / ":" / "@" / "&" / "=" / "+ " )
   query          =  *uric
   SIP-Version    =  "SIP" "/" 1*DIGIT "." 1*DIGIT

   message-header  =  (Accept
                   /  Accept-Encoding
                   /  Accept-Language
                   /  Alert-Info
                   /  Allow
                   /  Authentication-Info
                   /  Authorization
                   /  Call-ID
                   /  Call-Info
                   /  Contact
                   /  Content-Disposition
                   /  Content-Encoding
                   /  Content-Language
                   /  Content-Length
                   /  Content-Type
                   /  CSeq
                   /  Date
                   /  Error-Info
                   /  Expires
                   /  From
                   /  In-Reply-To
                   /  Max-Forwards
                   /  MIME-Version
                   /  Min-Expires
                   /  Organization
                   /  Priority
                   /  Proxy-Authenticate
                   /  Proxy-Authorization
                   /  Proxy-Require
                   /  Record-Route
                   /  Reply-To
                   /  Require
                   /  Retry-After
                   /  Route
                   /  Server
                   /  Subject
                   /  Supported
                   /  Timestamp
                   /  To
                   /  Unsupported
                   /  User-Agent
                   /  Via
                   /  Warning
                   /  WWW-Authenticate
                   /  extension-header) CRLF

   INVITEm           =  %x49.4E.56.49.54.45 ; INVITE in caps
   ACKm              =  %x41.43.4B ; ACK in caps
   OPTIONSm          =  %x4F.50.54.49.4F.4E.53 ; OPTIONS in caps
   BYEm              =  %x42.59.45 ; BYE in caps
   CANCELm           =  %x43.41.4E.43.45.4C ; CANCEL in caps
   REGISTERm         =  %x52.45.47.49.53.54.45.52 ; REGISTER in caps
   Method            =  INVITEm / ACKm / OPTIONSm / BYEm
                        / CANCELm / REGISTERm
                        / extension-method
   extension-method  =  token
   Response          =  Status-Line
                        *( message-header )
                        CRLF
                        [ message-body ]

   Status-Line     =  SIP-Version SP Status-Code SP Reason-Phrase CRLF
   Status-Code     =  Informational
                  /   Redirection
                  /   Success
                  /   Client-Error
                  /   Server-Error
                  /   Global-Failure
                  /   extension-code
   extension-code  =  3DIGIT
   Reason-Phrase   =  *(reserved / unreserved / escaped
                      / UTF8-NONASCII / UTF8-CONT / SP / HTAB)

   Informational  =  "100"  ;  Trying
                 /   "180"  ;  Ringing
                 /   "181"  ;  Call Is Being Forwarded
                 /   "182"  ;  Queued
                 /   "183"  ;  Session Progress

   Success  =  "200"  ;  OK

   Redirection  =  "300"  ;  Multiple Choices
               /   "301"  ;  Moved Permanently
               /   "302"  ;  Moved Temporarily
               /   "305"  ;  Use Proxy
               /   "380"  ;  Alternative Service

   Client-Error  =  "400"  ;  Bad Request
                /   "401"  ;  Unauthorized
                /   "402"  ;  Payment Required
                /   "403"  ;  Forbidden
                /   "404"  ;  Not Found
                /   "405"  ;  Method Not Allowed
                /   "406"  ;  Not Acceptable
                /   "407"  ;  Proxy Authentication Required
                /   "408"  ;  Request Timeout
                /   "410"  ;  Gone
                /   "413"  ;  Request Entity Too Large
                /   "414"  ;  Request-URI Too Large
                /   "415"  ;  Unsupported Media Type
                /   "416"  ;  Unsupported URI Scheme
                /   "420"  ;  Bad Extension
                /   "421"  ;  Extension Required
                /   "423"  ;  Interval Too Brief
                /   "480"  ;  Temporarily not available
                /   "481"  ;  Call Leg/Transaction Does Not Exist
                /   "482"  ;  Loop Detected
                /   "483"  ;  Too Many Hops
                /   "484"  ;  Address Incomplete
                /   "485"  ;  Ambiguous
                /   "486"  ;  Busy Here
                /   "487"  ;  Request Terminated
                /   "488"  ;  Not Acceptable Here
                /   "491"  ;  Request Pending
                /   "493"  ;  Undecipherable

   Server-Error  =  "500"  ;  Internal Server Error
                /   "501"  ;  Not Implemented
                /   "502"  ;  Bad Gateway
                /   "503"  ;  Service Unavailable
                /   "504"  ;  Server Time-out
                /   "505"  ;  SIP Version not supported
                /   "513"  ;  Message Too Large

   Global-Failure  =  "600"  ;  Busy Everywhere
                  /   "603"  ;  Decline
                  /   "604"  ;  Does not exist anywhere
                  /   "606"  ;  Not Acceptable

   Accept         =  "Accept" HCOLON
                      [ accept-range *(COMMA accept-range) ]
   accept-range   =  media-range *(SEMI accept-param)
   ; Why not SLASH ? //PPe
   media-range    =  ( "*" "/" "*"
                     / ( m-type SLASH "*" )
                     / ( m-type SLASH m-subtype )
                     ) *( SEMI m-parameter )
   accept-param   =  ("q" EQUAL qvalue) / generic-param
   qvalue         =  ( "0" [ "." 0*3DIGIT ] )
                     / ( "1" [ "." 0*3("0") ] )
   generic-param  =  token [ EQUAL gen-value ]
   gen-value      =  token / host / quoted-string

   Accept-Encoding  =  "Accept-Encoding" HCOLON
                        [ encoding *(COMMA encoding) ]
   encoding         =  codings *(SEMI accept-param)
   codings          =  content-coding / "*"
   content-coding   =  token

   Accept-Language  =  "Accept-Language" HCOLON
                        [ language *(COMMA language) ]
   language         =  language-range *(SEMI accept-param)
   language-range   =  ( ( 1*8ALPHA *( "-" 1*8ALPHA ) ) / "*" )

   Alert-Info   =  "Alert-Info" HCOLON alert-param *(COMMA alert-param)
   alert-param  =  LAQUOT absoluteURI RAQUOT *( SEMI generic-param )

   Allow  =  "Allow" HCOLON [Method *(COMMA Method)]

   Authorization     =  "Authorization" HCOLON credentials
   credentials       =  ("Digest" LWS digest-response)
                        / other-response
   digest-response   =  dig-resp *(COMMA dig-resp)
   dig-resp          =  username / realm / nonce / digest-uri
                         / dresponse / algorithm / cnonce
                         / opaque / message-qop
                         / nonce-count / auth-param
   username          =  "username" EQUAL username-value
   username-value    =  quoted-string
   digest-uri        =  "uri" EQUAL LDQUOT digest-uri-value RDQUOT
   digest-uri-value  =  rquest-uri ; Equal to request-uri as specified
                        by HTTP/1.1
   message-qop       =  "qop" EQUAL qop-value
   cnonce            =  "cnonce" EQUAL cnonce-value
   cnonce-value      =  nonce-value
   nonce-count       =  "nc" EQUAL nc-value
   nc-value          =  8LHEX
   dresponse         =  "response" EQUAL request-digest
   request-digest    =  LDQUOT 32LHEX RDQUOT
   auth-param        =  auth-param-name EQUAL
                        ( token / quoted-string )
   auth-param-name   =  token
   other-response    =  auth-scheme LWS auth-param
                        *(COMMA auth-param)
   auth-scheme       =  token

   Authentication-Info  =  "Authentication-Info" HCOLON ainfo
                           *(COMMA ainfo)
   ainfo                =  nextnonce / message-qop
                            / response-auth / cnonce
                            / nonce-count
   nextnonce            =  "nextnonce" EQUAL nonce-value
   response-auth        =  "rspauth" EQUAL response-digest
   response-digest      =  LDQUOT *LHEX RDQUOT

   Call-ID  =  ( "Call-ID" / "i" ) HCOLON callid
   callid   =  word [ "@" word ]

   Call-Info   =  "Call-Info" HCOLON info *(COMMA info)
   info        =  LAQUOT absoluteURI RAQUOT *( SEMI info-param)
   info-param  =  ( "purpose" EQUAL ( "icon" / "info"
                  / "card" / token ) ) / generic-param

   Contact        =  ("Contact" / "m" ) HCOLON
                     ( STAR / (contact-param *(COMMA contact-param)))
   contact-param  =  (name-addr / addr-spec) *(SEMI contact-params)
   name-addr      =  [ display-name ] LAQUOT addr-spec RAQUOT
   addr-spec      =  SIP-URI / SIPS-URI / absoluteURI
   display-name   =  *(token LWS)/ quoted-string

   contact-params     =  c-p-q / c-p-expires
                         / contact-extension
   c-p-q              =  "q" EQUAL qvalue
   c-p-expires        =  "expires" EQUAL delta-seconds
   contact-extension  =  generic-param
   delta-seconds      =  1*DIGIT

   Content-Disposition   =  "Content-Disposition" HCOLON
                            disp-type *( SEMI disp-param )
   disp-type             =  "render" / "session" / "icon" / "alert"
                            / disp-extension-token
   disp-param            =  handling-param / generic-param
   handling-param        =  "handling" EQUAL
                            ( "optional" / "required"
                            / other-handling )
   other-handling        =  token
   disp-extension-token  =  token

   Content-Encoding  =  ( "Content-Encoding" / "e" ) HCOLON
                        content-coding *(COMMA content-coding)

   Content-Language  =  "Content-Language" HCOLON
                        language-tag *(COMMA language-tag)
   language-tag      =  primary-tag *( "-" subtag )
   primary-tag       =  1*8ALPHA
   subtag            =  1*8ALPHA

   Content-Length  =  ( "Content-Length" / "l" ) HCOLON 1*DIGIT
   Content-Type     =  ( "Content-Type" / "c" ) HCOLON media-type
   media-type       =  m-type SLASH m-subtype *(SEMI m-parameter)
   m-type           =  discrete-type / composite-type
   discrete-type    =  "text" / "image" / "audio" / "video"
                       / "application" / extension-token
   composite-type   =  "message" / "multipart" / extension-token
   extension-token  =  ietf-token / x-token
   ietf-token       =  token
   x-token          =  "x-" token
   m-subtype        =  extension-token / iana-token
   iana-token       =  token
   m-parameter      =  m-attribute EQUAL m-value
   m-attribute      =  token
   m-value          =  token / quoted-string

   CSeq  =  "CSeq" HCOLON 1*DIGIT LWS Method

   Date          =  "Date" HCOLON SIP-date
   SIP-date      =  rfc1123-date
   rfc1123-date  =  wkday "," SP date1 SP time SP "GMT"
   date1         =  2DIGIT SP month SP 4DIGIT
                    ; day month year (e.g., 02 Jun 1982)
   time          =  2DIGIT ":" 2DIGIT ":" 2DIGIT
                    ; 00:00:00 - 23:59:59
   wkday         =  "Mon" / "Tue" / "Wed"
                    / "Thu" / "Fri" / "Sat" / "Sun"
   month         =  "Jan" / "Feb" / "Mar" / "Apr"
                    / "May" / "Jun" / "Jul" / "Aug"
                    / "Sep" / "Oct" / "Nov" / "Dec"

   Error-Info  =  "Error-Info" HCOLON error-uri *(COMMA error-uri)
   error-uri   =  LAQUOT absoluteURI RAQUOT *( SEMI generic-param )

   Expires     =  "Expires" HCOLON delta-seconds
   From        =  ( "From" / "f" ) HCOLON from-spec
   from-spec   =  ( name-addr / addr-spec )
                  *( SEMI from-param )
   from-param  =  tag-param / generic-param
   tag-param   =  "tag" EQUAL token

   In-Reply-To  =  "In-Reply-To" HCOLON callid *(COMMA callid)

   Max-Forwards  =  "Max-Forwards" HCOLON 1*DIGIT

   MIME-Version  =  "MIME-Version" HCOLON 1*DIGIT "." 1*DIGIT

   Min-Expires  =  "Min-Expires" HCOLON delta-seconds

   Organization  =  "Organization" HCOLON [TEXT-UTF8-TRIM]

   Priority        =  "Priority" HCOLON priority-value
   priority-value  =  "emergency" / "urgent" / "normal"
                      / "non-urgent" / other-priority
   other-priority  =  token

   Proxy-Authenticate  =  "Proxy-Authenticate" HCOLON challenge
   challenge           =  ("Digest" LWS digest-cln *(COMMA digest-cln))
                          / other-challenge
   other-challenge     =  auth-scheme LWS auth-param
                          *(COMMA auth-param)
   digest-cln          =  realm / domain / nonce
                           / opaque / stale / algorithm
                           / qop-options / auth-param
   realm               =  "realm" EQUAL realm-value
   realm-value         =  quoted-string
   domain              =  "domain" EQUAL LDQUOT URI
                          *( 1*SP URI ) RDQUOT
   URI                 =  absoluteURI / abs-path
   nonce               =  "nonce" EQUAL nonce-value
   nonce-value         =  quoted-string
   opaque              =  "opaque" EQUAL quoted-string
   stale               =  "stale" EQUAL ( "true" / "false" )
   algorithm           =  "algorithm" EQUAL ( "MD5" / "MD5-sess"
                          / token )
   qop-options         =  "qop" EQUAL LDQUOT qop-value
                          *("," qop-value) RDQUOT
   qop-value           =  "auth" / "auth-int" / token

   Proxy-Authorization  =  "Proxy-Authorization" HCOLON credentials

   Proxy-Require  =  "Proxy-Require" HCOLON option-tag
                     *(COMMA option-tag)
   option-tag     =  token

   Record-Route  =  "Record-Route" HCOLON rec-route *(COMMA rec-route)
   rec-route     =  name-addr *( SEMI rr-param )
   rr-param      =  generic-param

   Reply-To      =  "Reply-To" HCOLON rplyto-spec
   rplyto-spec   =  ( name-addr / addr-spec )
                    *( SEMI rplyto-param )
   rplyto-param  =  generic-param
   Require       =  "Require" HCOLON option-tag *(COMMA option-tag)

   Retry-After  =  "Retry-After" HCOLON delta-seconds
                   [ comment ] *( SEMI retry-param )

   retry-param  =  ("duration" EQUAL delta-seconds)
                   / generic-param

   Route        =  "Route" HCOLON route-param *(COMMA route-param)
   route-param  =  name-addr *( SEMI rr-param )

   Server           =  "Server" HCOLON server-val *(LWS server-val)
   server-val       =  product / comment
   product          =  token [SLASH product-version]
   product-version  =  token

   Subject  =  ( "Subject" / "s" ) HCOLON [TEXT-UTF8-TRIM]

   Supported  =  ( "Supported" / "k" ) HCOLON
                 [option-tag *(COMMA option-tag)]

   Timestamp  =  "Timestamp" HCOLON 1*(DIGIT)
                  [ "." *(DIGIT) ] [ LWS delay ]
   delay      =  *(DIGIT) [ "." *(DIGIT) ]

   To        =  ( "To" / "t" ) HCOLON ( name-addr
                / addr-spec ) *( SEMI to-param )
   to-param  =  tag-param / generic-param

   Unsupported  =  "Unsupported" HCOLON option-tag *(COMMA option-tag)
   User-Agent  =  "User-Agent" HCOLON server-val *(LWS server-val)

   Via               =  ( "Via" / "v" ) HCOLON via-parm *(COMMA via-parm)
   via-parm          =  sent-protocol LWS sent-by *( SEMI via-params )
   via-params        =  via-ttl / via-maddr
                        / via-received / via-branch
                        / via-extension
   via-ttl           =  "ttl" EQUAL ttl
   via-maddr         =  "maddr" EQUAL host
   via-received      =  "received" EQUAL (IPv4address / IPv6address)
   via-branch        =  "branch" EQUAL token
   via-extension     =  generic-param
   sent-protocol     =  protocol-name SLASH protocol-version
                        SLASH transport
   protocol-name     =  "SIP" / token
   protocol-version  =  token
   transport         =  "UDP" / "TCP" / "TLS" / "SCTP"
                        / other-transport
   sent-by           =  host [ COLON port ]
   ttl               =  1*3DIGIT ; 0 to 255

   Warning        =  "Warning" HCOLON warning-value *(COMMA warning-value)
   warning-value  =  warn-code SP warn-agent SP warn-text
   warn-code      =  3DIGIT
   warn-agent     =  hostport / pseudonym
                     ;  the name or pseudonym of the server adding
                     ;  the Warning header, for use in debugging
   warn-text      =  quoted-string
   pseudonym      =  token

   WWW-Authenticate  =  "WWW-Authenticate" HCOLON challenge

   extension-header  =  header-name HCOLON header-value
   header-name       =  token
   header-value      =  *(TEXT-UTF8char / UTF8-CONT / LWS)
   message-body  =  *OCTET
```
###Note: cleanWs improvement

```
This is a top 20 sample I took before modifying the cleanWs function in 
sipparser/utils.go.

The initial version of cleanWs looped through each char in the string
and compared characters.  This version occupied 80%+  of the time in
the bench program.  The new version uses the strings package to trim the 
space at the end of the line then loop through the returned slice of strings.
The net result is a 7X improvement in speed of the bench program and a reduction
of the time spent to just over 20%.   

before the modification:
Total: 1802 samples
     203  11.3%  11.3%     1098  60.9% runtime.mallocgc
     128   7.1%  18.4%      347  19.3% sweep
     112   6.2%  24.6%      129   7.2% MCentral_Alloc
     110   6.1%  30.7%     1494  82.9% sipparser.cleanWs
      92   5.1%  35.8%      434  24.1% runtime.MCache_Alloc
      89   4.9%  40.7%      598  33.2% concatstring
      85   4.7%  45.4%       85   4.7% runtime.mcpy
      80   4.4%  49.9%       80   4.4% runtime.memclr
      79   4.4%  54.3%      229  12.7% runtime.MCache_Free
      75   4.2%  58.4%     1124  62.4% runtime.gostringsize
      70   3.9%  62.3%      120   6.7% MCentral_Free
      60   3.3%  65.6%      756  42.0% runtime.intstring
      53   2.9%  68.6%       53   2.9% runtime.markallocated
      50   2.8%  71.4%       50   2.8% runtime.stringiter
      50   2.8%  74.1%       50   2.8% scanblock
      44   2.4%  76.6%       44   2.4% runtime.SizeToClass
      41   2.3%  78.9%       41   2.3% runtime.MHeap_Lookup
      40   2.2%  81.1%       40   2.2% runtime.markspan
      33   1.8%  82.9%       33   1.8% runtime.MSpanList_IsEmpty
      30   1.7%  84.6%      341  18.9% runtime.MCentral_AllocList

after the modification:
Total: 254 samples
      18   7.1%   7.1%       18   7.1% runtime.stringiter
      14   5.5%  12.6%       38  15.0% sipparser.getCrlf
      13   5.1%  17.7%       14   5.5% sipparser.getHdrFunc
      11   4.3%  22.0%       11   4.3% MCentral_Alloc
       9   3.5%  25.6%       78  30.7% runtime.mallocgc
       9   3.5%  29.1%        9   3.5% strings.Count
       9   3.5%  32.7%       22   8.7% sweep
       8   3.1%  35.8%       38  15.0% runtime.MCache_Alloc
       6   2.4%  38.2%       12   4.7% MHeap_AllocLocked
       6   2.4%  40.6%       14   5.5% runtime.MCache_Free
       6   2.4%  42.9%       30  11.8% runtime.makeslice
       6   2.4%  45.3%       60  23.6% sipparser.cleanWs
       6   2.4%  47.6%      184  72.4% sipparser.getHeaders
       5   2.0%  49.6%        6   2.4% MCentral_Free
       5   2.0%  51.6%       19   7.5% runtime.growslice
       5   2.0%  53.5%        5   2.0% scanblock
       5   2.0%  55.5%        9   3.5% sipparser.getBracks
       5   2.0%  57.5%       33  13.0% sipparser.parseUriHost
       5   2.0%  59.4%       30  11.8% strings.Map
       4   1.6%  61.0%       34  13.4% makeslice1
```

###License

The project is governed by a BSD license that can be found in the LICENSE.txt file.