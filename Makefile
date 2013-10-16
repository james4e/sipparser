include $(GOROOT)/src/Make.inc
 
TARG=sipparser
GOFILES=accept.go authorization.go constants.go contentdisposition.go cseq.go from.go params.go parser.go passertedid.go rack.go reason.go remotepartyid.go startline.go uri.go utils.go via.go warning.go
 
include $(GOROOT)/src/Make.pkg 
