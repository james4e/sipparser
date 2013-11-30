# sipparser [![Build Status](https://drone.io/github.com/danielvargas/sipparser/status.png)]

    import "github.com/danielvargas/sipparser"


### Index

* [Constants](#constants)
* [Variables](#variables)
* [type Accept](#type-accept)
* [type AcceptParam](#type-acceptparam)
* [type Authorization](#type-authorization)
    + [func (*Authorization) GetParam](#func---authorization--getparam)
* [type CallingPartyInfo](#type-callingpartyinfo)
* [type ContentDisposition](#type-contentdisposition)
* [type Cseq](#type-cseq)
* [type From](#type-from)
* [type Header](#type-header)
    + [func (*Header) String](#func---header--string)
* [type PAssertedId](#type-passertedid)
* [type Param](#type-param)
* [type Rack](#type-rack)
* [type Reason](#type-reason)
* [type RemotePartyId](#type-remotepartyid)
* [type SipMsg](#type-sipmsg)
    + [func (*SipMsg) GetCallingParty](#func---sipmsg--getcallingparty)
    + [func (*SipMsg) GetRURIParamBool](#func---sipmsg--getruriparambool)
    + [func (*SipMsg) GetRURIParamVal](#func---sipmsg--getruriparamval)
    + [func (*SipMsg) ParseContact](#func---sipmsg--parsecontact)
    + [func (*SipMsg) ParsePAssertedId](#func---sipmsg--parsepassertedid)
    + [func (*SipMsg) ParseRemotePartyId](#func---sipmsg--parseremotepartyid)
* [type StartLine](#type-startline)
* [type URI](#type-uri)
    + [func (*URI) Parse](#func---uri--parse)
* [type Via](#type-via)
    + [func (*Via) AddReceived](#func---via--addreceived)
* [type Warning](#type-warning)


#### Constants
```go
const (
	// SIP request or response
	SIP_REQUEST  = "REQUEST"
	SIP_RESPONSE = "RESPONSE"
	// SIP Methods
	SIP_METHOD_INVITE    = "INVITE"
	SIP_METHOD_ACK       = "ACK"
	SIP_METHOD_OPTIONS   = "OPTIONS"
	SIP_METHOD_BYE       = "BYE"
	SIP_METHOD_CANCEL    = "CANCEL"
	SIP_METHOD_REGISTER  = "REGISTER"
	SIP_METHOD_INFO      = "INFO"
	SIP_METHOD_PRACK     = "PRACK"
	SIP_METHOD_SUBSCRIBE = "SUBSCRIBE"
	SIP_METHOD_NOTIFY    = "NOTIFY"
	SIP_METHOD_UPDATE    = "UPDATE"
	SIP_METHOD_MESSAGE   = "MESSAGE"
	SIP_METHOD_REFER     = "REFER"
	SIP_METHOD_PUBLISH   = "PUBLISH"
	// SIP Headers
	SIP_HDR_ACCEPT                        = "accept"          // RFC3261
	SIP_HDR_ACCEPT_CONTACT                = "accept-contact"  // RFC3841
	SIP_HDR_ACCEPT_CONTACT_CMP            = "a"               // RFC3841
	SIP_HDR_ACCEPT_ENCODING               = "accept-encoding" //
	SIP_HDR_ACCEPT_LANGUAGE               = "accept-language"
	SIP_HDR_ACCEPT_RESOURCE_PRIORITY      = "accept-resource-priority" // RFC4412
	SIP_HDR_ALERT_INFO                    = "alert-info"
	SIP_HDR_ALLOW                         = "allow"
	SIP_HDR_ALLOW_EVENTS                  = "allow-events"
	SIP_HDR_ALLOW_EVENTS_CMP              = "u"
	SIP_HDR_ANSWER_MODE                   = "answer-mode"
	SIP_HDR_AUTHENTICATION_INFO           = "authentication-info"
	SIP_HDR_AUTHORIZATION                 = "authorization"
	SIP_HDR_CALL_ID                       = "call-id"
	SIP_HDR_CALL_ID_CMP                   = "i"
	SIP_HDR_CALL_INFO                     = "call-info"
	SIP_HDR_CONTACT                       = "contact"
	SIP_HDR_CONTACT_CMP                   = "m"
	SIP_HDR_CONTENT_DISPOSITION           = "content-disposition"
	SIP_HDR_CONTENT_ENCODING              = "content-encoding"
	SIP_HDR_CONTENT_ENCODING_CMP          = "e"
	SIP_HDR_CONTENT_LANGUAGE              = "content-language"
	SIP_HDR_CONTENT_LENGTH                = "content-length"
	SIP_HDR_CONTENT_LENGTH_CMP            = "l"
	SIP_HDR_CONTENT_TYPE                  = "content-type"
	SIP_HDR_CONTENT_TYPE_CMP              = "c"
	SIP_HDR_CSEQ                          = "cseq"
	SIP_HDR_DATE                          = "date"
	SIP_HDR_ERROR_INFO                    = "error-info"
	SIP_HDR_EVENT                         = "event"
	SIP_HDR_EXPIRES                       = "expires"
	SIP_HDR_FLOW_TIMER                    = "flow-timer"
	SIP_HDR_FROM                          = "from"
	SIP_HDR_FROM_CMP                      = "f"
	SIP_HDR_HISTORY_INFO                  = "history-info"  // from RFC 4244
	SIP_HDR_IDENTITY                      = "identity"      // RFC 4474
	SIP_HDR_IDENTITY_CMP                  = "y"             // RFC 4474
	SIP_HDR_IDENTITY_INFO                 = "identity-info" // RFC 4474
	SIP_HDR_IDENTITY_INFO_CMP             = "n"             // RFC 4474
	SIP_HDR_IN_REPLY_TO                   = "in-reply-to"
	SIP_HDR_JOIN                          = "join" // RFC 3911
	SIP_HDR_MAX_FORWARDS                  = "max-forwards"
	SIP_HDR_MIME_VERSION                  = "mime-version"
	SIP_HDR_MIN_EXPIRES                   = "min-expires"
	SIP_HDR_MIN_SE                        = "min-se" // RFC4028
	SIP_HDR_ORGANIZATION                  = "organization"
	SIP_HDR_PATH                          = "path"               // RFC3327
	SIP_HDR_PERMISSION_MISSING            = "permission-missing" // RFC5360
	SIP_HDR_PRIORITY                      = "priority"
	SIP_HDR_PRIVACY                       = "privacy"
	SIP_HDR_PRIV_ANSWER_MODE              = "priv-answer-mode" // RFC 5373
	SIP_HDR_PROXY_AUTHENTICATE            = "proxy-authenticate"
	SIP_HDR_PROXY_AutHORIZATION           = "proxy-authorization"
	SIP_HDR_PROXY_REQUIRE                 = "proxy-require"
	SIP_HDR_RACK                          = "rack" // RFC 3262
	SIP_HDR_REASON                        = "reason"
	SIP_HDR_RECORD_ROUTE                  = "record-route"
	SIP_HDR_REFER_SUB                     = "refer-sub"                     // RFC4488
	SIP_HDR_REFER_TO                      = "refer-to"                      // RFC 3515, RFC 4508
	SIP_HDR_REFERRED_BY                   = "referred-by"                   // RFC3892
	SIP_HDR_REFERRED_BY_CMP               = "b"                             // RFC3892
	SIP_HDR_REJECT_CONTACT                = "reject-contact"                // RFC3841
	SIP_HDR_REJECT_CONTACT_CMP            = "j"                             // RFC3841
	SIP_HDR_REMOTE_PARTY_ID               = "remote-party-id"               // DRAFT
	SIP_HDR_REPLACES                      = "replaces"                      // RFC3891
	SIP_HDR_REPLY_TO                      = "reply-to"                      // RFC3261
	SIP_HDR_REQUEST_DISPOSITION           = "request-disposition"           // RFC3841
	SIP_HDR_REQUIRE                       = "require"                       // RFC3261
	SIP_HDR_RESOURCE_PRIORITY             = "resource-priority"             // RFC4412
	SIP_HDR_RETRY_AFTER                   = "retry-after"                   // RFC3261
	SIP_HDR_ROUTE                         = "route"                         // RFC3261
	SIP_HDR_RSEQ                          = "rseq"                          // RFC3262
	SIP_HDR_SECUTIRY_CLIENT               = "security-client"               // RFC3329
	SIP_HDR_SECURITY_SERVER               = "security-server"               // RFC3329
	SIP_HDR_SECURITY_VERIFY               = "security-verify"               // RFC3329
	SIP_HDR_SERVER                        = "server"                        // RFC3261
	SIP_HDR_SERVICE_ROUTE                 = "service-route"                 // RFC3608
	SIP_HDR_SESSION_EXPIRES               = "session-expires"               // RFC4028
	SIP_HDR_SESSION_EXPIRES_CMP           = "x"                             // RFC4028
	SIP_HDR_SIP_ETAG                      = "sip-etag"                      // RFC3903
	SIP_HDR_SIP_IF_MATCH                  = "sip-if-match"                  // RFC3903
	SIP_HDR_SUBJECT                       = "subject"                       // RFC3261
	SIP_HDR_SUBJECT_CMP                   = "s"                             // RFC3261
	SIP_HDR_SUBSCRIPTION_STATE            = "subscription-state"            // RFC3265
	SIP_HDR_SUPPORTED                     = "supported"                     // RFC3261
	SIP_HDR_SUPPORTED_CMP                 = "k"                             // RFC3261
	SIP_HDR_SUPPRESS_IF_MATCH             = "suppress-if-match"             // RFC5839
	SIP_HDR_TARGET_DIALOG                 = "target-dialog"                 // RFC4538
	SIP_HDR_TIMESTAMP                     = "timestamp"                     // RFC3261
	SIP_HDR_TO                            = "to"                            // RFC3261
	SIP_HDR_TO_CMP                        = "t"                             // RFC3261
	SIP_HDR_TRIGGER_CONSENT               = "trigger-consent"               // RFC5360
	SIP_HDR_UNSUPPORTED                   = "unsupported"                   // RFC3261
	SIP_HDR_USER_AGENT                    = "user-agent"                    // RFC3261
	SIP_HDR_VIA                           = "via"                           // RFC3261
	SIP_HDR_VIA_CMP                       = "v"                             // RFC3261
	SIP_HDR_WARNING                       = "warning"                       // RFC3261
	SIP_HDR_WWW_AUTHENTICATE              = "www-authenticate"              // RFC3261
	SIP_HDR_P_ACCESS_NETWORK_INFO         = "p-access-network-info"         // RFC3455
	SIP_HDR_P_ANSWER_STATE                = "p-answer-state"                // RFC3455
	SIP_HDR_P_ASSERTED_IDENTITY           = "p-asserted-identity"           // RFC3325
	SIP_HDR_P_ASSERTED_SERVICE            = "p-asserted-service"            // RFC3455
	SIP_HDR_P_ASSOCIATED_URI              = "p-associated-uri"              // RFC3455
	SIP_HDR_P_CALLED_PARTY_ID             = "p-called-party-id"             // RFC3455
	SIP_HDR_P_CHARGING_FUNCTION_ADDRESSES = "p-charging-function-addresses" // RFC3455
	SIP_HDR_P_CHARGING_VECTOR             = "p-charging-vector"             // RFC3455
	SIP_HDR_P_DCS_BILLING_INFO            = "p-dcs-billing-info"            // RFC5503
	SIP_HDR_P_DCS_LAES                    = "p-dcs-laes"                    // RFC5503
	SIP_HDR_P_DCS_OSPS                    = "p-dcs-osps"                    // RFC5503
	SIP_HDR_P_DCS_REDIRECT                = "p-dcs-redirect"                // RFC5503
	SIP_HDR_P_DCS_TRACE_PARTY_ID          = "p-dcs-trace-party-id"          // RFC5503
	SIP_HDR_P_EARLY_MEDIA                 = "p-early-media"                 // RFC5009
	SIP_HDR_P_MEDIA_AUTHORIZATION         = "p-media-authorization"         // RFC3313
	SIP_HDR_P_PREFERRED_IDENTITY          = "p-preferred-identity"          // RFC3325
	SIP_HDR_P_PREFERRED_SERVICE           = "p-preferred-service"           // RFC6050
	SIP_HDR_P_PROFILE_KEY                 = "p-profile-key"                 // RFC5002
	SIP_HDR_P_USER_DATABASE               = "p-user-database"               // RFC4457
	SIP_HDR_P_VISITED_NETWORK_ID          = "p-visited-network-id"          // RFC3455
)
```
constants just holds a common shared set of constants (i.e. hdr values)

```go
const (
	CR                    = "\r"
	LF                    = "\n"
	CALLING_PARTY_DEFAULT = "default"
	CALLING_PARTY_RPID    = "rpid"
	CALLING_PARTY_PAID    = "paid"
)
```

```go
const (
	SIP_SCHEME  = "sip"
	SIPS_SCHEME = "sips"
	TEL_SCHEME  = "tel"
)
```

#### Variables

```go
var (
	SIP_METHODS = []string{
		SIP_METHOD_INVITE,
		SIP_METHOD_ACK,
		SIP_METHOD_OPTIONS,
		SIP_METHOD_BYE,
		SIP_METHOD_CANCEL,
		SIP_METHOD_REGISTER,
		SIP_METHOD_INFO,
		SIP_METHOD_PRACK,
		SIP_METHOD_SUBSCRIBE,
		SIP_METHOD_NOTIFY,
		SIP_METHOD_UPDATE,
		SIP_METHOD_MESSAGE,
		SIP_METHOD_REFER,
		SIP_METHOD_PUBLISH,
	}
)
```

#### type [Accept](#accept)

```go
type Accept struct {
	// Val is the raw value
	Val string
	// Params is a slice of AcceptParam
	Params []*AcceptParam
}
```

Accept is a struct that holds the following: -- the raw value -- a slice of
parced AcceptParam

#### type [AcceptParam](#acceptparam)

```go
type AcceptParam struct {
	Type string
	Val  string
}
```

AcceptParam is just a key:value pair of params for the accept header

#### type [Authorization](#authorization)

```go
type Authorization struct {
	Val         string   "val"
	Credentials string   "credentials"
	Params      []*Param "params"
}
```


#### func (*Authorization) [GetParam](#getparam)

```go
func (a *Authorization) GetParam(param string) *Param
```

#### type [CallingPartyInfo](#callingpartyinfo)

```go
type CallingPartyInfo struct {
	// Name the name
	Name string
	// Number the number
	Number string
	// Anonymous a bool to see if this should be anonymous or not
	Anonymous bool
}
```

CallingPartyInfo is a struct of calling party information. This is populated
into the *SipMsg.CallingParty field when the method GetCallingParty on the
*SipMsg. See below for details of that method.

#### type [ContentDisposition](#contentdisposition)

```go
type ContentDisposition struct {
	// Val is the raw value
	Val string
	// DispType is the display type
	DispType string
	// Params is a slice of *Param
	Params []*Param
}
```

ContentDisposition is a struct that holds a parsed content-disposition hdr: --
Val is the raw value -- DispType is the display type -- Params is slice of
parameters

#### type [Cseq](#cseq)

```go
type Cseq struct {
	// Val is the raw value
	Val string
	// Method is the SIP method
	Method string
	// Digit is the digit
	Digit string
}
```

Cseq is a struct that holds the values for a cseq header:

    -- Val is the raw string value of the cseq hdr
    -- Method is the SIP method
    -- Digit is the numeric indicator for the method

#### type [From](#from)

```go
type From struct {
	// Error is an error
	Error error
	// Val is the raw value
	Val string
	// Name is the name value
	Name string
	// Tag is the tag value
	Tag string
	// URI is a parsed uri
	URI *URI
	// Params are for any generic params that are part of the header
	Params []*Param
}
```

from holds a parsed header that has a format like: "NAME"
<sip:user@hostinfo>;param=val and is used for the parsing of the from, to, and
contact header in the parser program From holds the following public fields:

#### type [Header](#header)

```go
type Header struct {
	Header string
	Val    string
}
```


#### func (*Header) [String](#string)

```go
func (h *Header) String() string
```

#### type [PAssertedId](#passertedid)

```go
type PAssertedId struct {
	Error  error
	Val    string
	Name   string
	URI    *URI
	Params []*Param
}
```

PAssertedId is a struct that holds: -- Error is just an os.Error -- Val is the
raw value -- Name is the name value from the p-asserted-id hdr -- URI is the
parsed uri from the p-asserted-id hdr -- Params is a slice of the *Params from
the p-asserted-id hdr

#### type [Param](#param)

```go
type Param struct {
	Param string "param"
	Val   string "val"
}
```

Param is just a struct that holds a parameter and a value As an example of this
would be something like user=phone

#### type [Rack](#rack)

```go
type Rack struct {
	Val        string // Val is the raw value
	RseqVal    string // RseqVal is the value of the rseq
	CseqVal    string // CseqVal is the value of the cseq
	CseqMethod string // CseqMethod is the value of the cseq method
}
```


#### type [Reason](#reason)

```go
type Reason struct {
	Val   string
	Proto string
	Cause string
	Text  string
}
```

Reason is a struct that holds a parsed reason hdr Fields are as follows: -- Val
is the raw value -- Proto is the protocol (i.e. SIP) -- Cause is the cause code
(i.e. 41) -- Text is the actual text response

#### type [RemotePartyId](#remotepartyid)

```go
type RemotePartyId struct {
	Error   error
	Val     string // Val is the raw value of the hdr
	Name    string // Name is the name from the header
	URI     *URI
	Party   string
	Screen  string
	Privacy string
	Params  []*Param
}
```


#### type [SipMsg](#sipmsg)

```go
type SipMsg struct {
	State              string
	Error              error
	Msg                string
	CallingParty       *CallingPartyInfo
	Body               string
	StartLine          *StartLine
	Headers            []*Header
	Accept             *Accept
	AlertInfo          string
	Allow              []string
	AllowEvents        []string
	Authorization      *Authorization
	ContentDisposition *ContentDisposition
	ContentLength      string
	ContentLengthInt   int
	ContentType        string
	From               *From
	MaxForwards        string
	MaxForwardsInt     int
	Organization       string
	To                 *From
	Contact            *From
	ContactVal         string
	CallId             string
	Cseq               *Cseq
	Rack               *Rack
	Reason             *Reason
	Rseq               string
	RseqInt            int
	RecordRoute        []*URI
	Route              []*URI
	Via                []*Via
	Require            []string
	Supported          []string
	Privacy            string
	ProxyAuthenticate  *Authorization
	ProxyRequire       []string
	RemotePartyIdVal   string
	RemotePartyId      *RemotePartyId
	PAssertedIdVal     string
	PAssertedId        *PAssertedId
	Unsupported        []string
	UserAgent          string
	Server             string
	Subject            string
	Warning            *Warning
	WWWAuthenticate    *Authorization
}
```


#### func  [ParseMsg](#parsemsg)

```go
func ParseMsg(str string) (s *SipMsg)
```

#### func (*SipMsg) [GetCallingParty](#getcallingparty)

```go
func (s *SipMsg) GetCallingParty(str string) error
```

#### func (*SipMsg) [GetRURIParamBool](#getruriparambool)

```go
func (s *SipMsg) GetRURIParamBool(str string) bool
```

#### func (*SipMsg) [GetRURIParamVal](#getruriparamval)

```go
func (s *SipMsg) GetRURIParamVal(str string) string
```

#### func (*SipMsg) [ParseContact](#parsecontact)

```go
func (s *SipMsg) ParseContact(str string)
```

#### func (*SipMsg) [ParsePAssertedId](#parsepassertedid)

```go
func (s *SipMsg) ParsePAssertedId(str string)
```

#### func (*SipMsg) [ParseRemotePartyId](#parseremotepartyid)

```go
func (s *SipMsg) ParseRemotePartyId(str string)
```

#### type [StartLine](#startline)

```go
type StartLine struct {
	Error    error  "err"
	Val      string "val"      // Val is the raw value
	Type     string "type"     // Type is the type of startline (i.e. request or response)
	Method   string "method"   // Method is the method (if request)
	URI      *URI   "uri"      // URI is the *URI (if request)
	Resp     string "resp"     // Resp is the response code (i.e. 183)
	RespText string "resptext" // RespText is the response text (i.e. "Session Progress")
	Proto    string "proto"    // Proto is the protocol (should be "SIP")
	Version  string "version"  // Version is the version (should be "2.0")
}
```


#### func  [ParseStartLine](#parsestartline)

```go
func ParseStartLine(str string) *StartLine
```

#### type [URI](#uri)

```go
type URI struct {
	Error        error  // error if any
	Scheme       string // scheme .. i.e. tel, sip, sips,etc.
	Raw          string // this is the actual raw uri unparsed
	UserInfo     string // this is everything before the "@"
	User         string // this is the actual called party
	UserPassword string // this is the password (i.e. alice:passwd@bob.com)
	HostInfo     string // this is everything after the @ or the entire uri
	Host         string // the host in the uri
	Port         string // the port
	UriParams    []*Param
	Secure       bool // Indicates SIP-URI or SIPS-URI (true for SIPS-URI)
}
```

URI is a struct that holds an error (hopefully nil), the raw value, and the
parsed uri. Fields are as follows: -- Error is the error (or nil) -- Scheme is
the scheme (i.e. sip) -- Raw is the raw value of the uri -- UserInfo is the
user:password;userparams=foo; -- User is the user (i.e. the phone number) --
UserPassword is the user password -- HostInfo is the host:port combination --
Host is the host -- Port is the port (if any) -- UriParams are the uri's
parameters -- Secure is if the scheme is "sips" -- atPos is just used by the
parser to identify where the

    "@" char is in the .Raw field (or 0 if not present)

#### func  [NewURI](#newuri)

```go
func NewURI(s string) *URI
```
NewURI is a convenience function that creates a *URI for you

#### func  [ParseURI](#parseuri)

```go
func ParseURI(s string) *URI
```
ParseURI is NewURI ... but also parses it

#### func (*URI) [Parse](#parse)

```go
func (u *URI) Parse()
```
Parse parses the .Raw field

#### type [Via](#via)

```go
type Via struct {
	State     string // State is the parser state
	Error     error  // Error is an error
	Via       string // Via is the raw value
	Proto     string // Proto is the protocol (i.e. "SIP")
	Version   string // Version is the version (i.e. "2.0")
	Transport string // Transport is the transport method (i.e. "UDP")
	SentBy    string // SentBy is a host:port combination
	Branch    string // Branch is the branch parameter
	Received  string
	RPort     string
	Params    []*Param
}
```

Via is an important part of the *SipMsg. It is a fundamental basis on which to
build route-sets and do call matching. It has the following structs:

#### func (*Via) [AddReceived](#addreceived)

```go
func (v *Via) AddReceived(s string)
```

#### type [Warning](#warning)

```go
type Warning struct {
	Val   string
	Code  string
	Agent string
	Text  string
}
```
