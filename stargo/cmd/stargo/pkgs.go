package main

import (
	"reflect"
	"go.starlark.net/stargo"
	"go.starlark.net/starlark"

	π0 "fmt"
	π1 "bytes"
	π2 "encoding/json"
	π3 "io/ioutil"
	π4 "net/http"
	π5 "go/token"
	π6 "go/ast"
	π7 "go/parser"
	π8 "go/types"
)

var goPackages = starlark.StringDict{
	"fmt": &starlark.Module{
		Name: "fmt",
		Members: starlark.StringDict{
			"Errorf": stargo.ValueOf(π0.Errorf),
			"Formatter": stargo.TypeOf(reflect.TypeOf(new(π0.Formatter)).Elem()),
			"Fprint": stargo.ValueOf(π0.Fprint),
			"Fprintf": stargo.ValueOf(π0.Fprintf),
			"Fprintln": stargo.ValueOf(π0.Fprintln),
			"Fscan": stargo.ValueOf(π0.Fscan),
			"Fscanf": stargo.ValueOf(π0.Fscanf),
			"Fscanln": stargo.ValueOf(π0.Fscanln),
			"GoStringer": stargo.TypeOf(reflect.TypeOf(new(π0.GoStringer)).Elem()),
			"Print": stargo.ValueOf(π0.Print),
			"Printf": stargo.ValueOf(π0.Printf),
			"Println": stargo.ValueOf(π0.Println),
			"Scan": stargo.ValueOf(π0.Scan),
			"ScanState": stargo.TypeOf(reflect.TypeOf(new(π0.ScanState)).Elem()),
			"Scanf": stargo.ValueOf(π0.Scanf),
			"Scanln": stargo.ValueOf(π0.Scanln),
			"Scanner": stargo.TypeOf(reflect.TypeOf(new(π0.Scanner)).Elem()),
			"Sprint": stargo.ValueOf(π0.Sprint),
			"Sprintf": stargo.ValueOf(π0.Sprintf),
			"Sprintln": stargo.ValueOf(π0.Sprintln),
			"Sscan": stargo.ValueOf(π0.Sscan),
			"Sscanf": stargo.ValueOf(π0.Sscanf),
			"Sscanln": stargo.ValueOf(π0.Sscanln),
			"State": stargo.TypeOf(reflect.TypeOf(new(π0.State)).Elem()),
			"Stringer": stargo.TypeOf(reflect.TypeOf(new(π0.Stringer)).Elem()),
		},
	},
	"bytes": &starlark.Module{
		Name: "bytes",
		Members: starlark.StringDict{
			"Buffer": stargo.TypeOf(reflect.TypeOf(new(π1.Buffer)).Elem()),
			"Compare": stargo.ValueOf(π1.Compare),
			"Contains": stargo.ValueOf(π1.Contains),
			"ContainsAny": stargo.ValueOf(π1.ContainsAny),
			"ContainsRune": stargo.ValueOf(π1.ContainsRune),
			"Count": stargo.ValueOf(π1.Count),
			"Equal": stargo.ValueOf(π1.Equal),
			"EqualFold": stargo.ValueOf(π1.EqualFold),
			"ErrTooLarge": stargo.VarOf(&π1.ErrTooLarge),
			"Fields": stargo.ValueOf(π1.Fields),
			"FieldsFunc": stargo.ValueOf(π1.FieldsFunc),
			"HasPrefix": stargo.ValueOf(π1.HasPrefix),
			"HasSuffix": stargo.ValueOf(π1.HasSuffix),
			"Index": stargo.ValueOf(π1.Index),
			"IndexAny": stargo.ValueOf(π1.IndexAny),
			"IndexByte": stargo.ValueOf(π1.IndexByte),
			"IndexFunc": stargo.ValueOf(π1.IndexFunc),
			"IndexRune": stargo.ValueOf(π1.IndexRune),
			"Join": stargo.ValueOf(π1.Join),
			"LastIndex": stargo.ValueOf(π1.LastIndex),
			"LastIndexAny": stargo.ValueOf(π1.LastIndexAny),
			"LastIndexByte": stargo.ValueOf(π1.LastIndexByte),
			"LastIndexFunc": stargo.ValueOf(π1.LastIndexFunc),
			"Map": stargo.ValueOf(π1.Map),
			"MinRead": stargo.ValueOf(π1.MinRead),
			"NewBuffer": stargo.ValueOf(π1.NewBuffer),
			"NewBufferString": stargo.ValueOf(π1.NewBufferString),
			"NewReader": stargo.ValueOf(π1.NewReader),
			"Reader": stargo.TypeOf(reflect.TypeOf(new(π1.Reader)).Elem()),
			"Repeat": stargo.ValueOf(π1.Repeat),
			"Replace": stargo.ValueOf(π1.Replace),
			"ReplaceAll": stargo.ValueOf(π1.ReplaceAll),
			"Runes": stargo.ValueOf(π1.Runes),
			"Split": stargo.ValueOf(π1.Split),
			"SplitAfter": stargo.ValueOf(π1.SplitAfter),
			"SplitAfterN": stargo.ValueOf(π1.SplitAfterN),
			"SplitN": stargo.ValueOf(π1.SplitN),
			"Title": stargo.ValueOf(π1.Title),
			"ToLower": stargo.ValueOf(π1.ToLower),
			"ToLowerSpecial": stargo.ValueOf(π1.ToLowerSpecial),
			"ToTitle": stargo.ValueOf(π1.ToTitle),
			"ToTitleSpecial": stargo.ValueOf(π1.ToTitleSpecial),
			"ToUpper": stargo.ValueOf(π1.ToUpper),
			"ToUpperSpecial": stargo.ValueOf(π1.ToUpperSpecial),
			"Trim": stargo.ValueOf(π1.Trim),
			"TrimFunc": stargo.ValueOf(π1.TrimFunc),
			"TrimLeft": stargo.ValueOf(π1.TrimLeft),
			"TrimLeftFunc": stargo.ValueOf(π1.TrimLeftFunc),
			"TrimPrefix": stargo.ValueOf(π1.TrimPrefix),
			"TrimRight": stargo.ValueOf(π1.TrimRight),
			"TrimRightFunc": stargo.ValueOf(π1.TrimRightFunc),
			"TrimSpace": stargo.ValueOf(π1.TrimSpace),
			"TrimSuffix": stargo.ValueOf(π1.TrimSuffix),
		},
	},
	"encoding/json": &starlark.Module{
		Name: "encoding/json",
		Members: starlark.StringDict{
			"Compact": stargo.ValueOf(π2.Compact),
			"Decoder": stargo.TypeOf(reflect.TypeOf(new(π2.Decoder)).Elem()),
			"Delim": stargo.TypeOf(reflect.TypeOf(new(π2.Delim)).Elem()),
			"Encoder": stargo.TypeOf(reflect.TypeOf(new(π2.Encoder)).Elem()),
			"HTMLEscape": stargo.ValueOf(π2.HTMLEscape),
			"Indent": stargo.ValueOf(π2.Indent),
			"InvalidUTF8Error": stargo.TypeOf(reflect.TypeOf(new(π2.InvalidUTF8Error)).Elem()),
			"InvalidUnmarshalError": stargo.TypeOf(reflect.TypeOf(new(π2.InvalidUnmarshalError)).Elem()),
			"Marshal": stargo.ValueOf(π2.Marshal),
			"MarshalIndent": stargo.ValueOf(π2.MarshalIndent),
			"Marshaler": stargo.TypeOf(reflect.TypeOf(new(π2.Marshaler)).Elem()),
			"MarshalerError": stargo.TypeOf(reflect.TypeOf(new(π2.MarshalerError)).Elem()),
			"NewDecoder": stargo.ValueOf(π2.NewDecoder),
			"NewEncoder": stargo.ValueOf(π2.NewEncoder),
			"Number": stargo.TypeOf(reflect.TypeOf(new(π2.Number)).Elem()),
			"RawMessage": stargo.TypeOf(reflect.TypeOf(new(π2.RawMessage)).Elem()),
			"SyntaxError": stargo.TypeOf(reflect.TypeOf(new(π2.SyntaxError)).Elem()),
			"Token": stargo.TypeOf(reflect.TypeOf(new(π2.Token)).Elem()),
			"Unmarshal": stargo.ValueOf(π2.Unmarshal),
			"UnmarshalFieldError": stargo.TypeOf(reflect.TypeOf(new(π2.UnmarshalFieldError)).Elem()),
			"UnmarshalTypeError": stargo.TypeOf(reflect.TypeOf(new(π2.UnmarshalTypeError)).Elem()),
			"Unmarshaler": stargo.TypeOf(reflect.TypeOf(new(π2.Unmarshaler)).Elem()),
			"UnsupportedTypeError": stargo.TypeOf(reflect.TypeOf(new(π2.UnsupportedTypeError)).Elem()),
			"UnsupportedValueError": stargo.TypeOf(reflect.TypeOf(new(π2.UnsupportedValueError)).Elem()),
			"Valid": stargo.ValueOf(π2.Valid),
		},
	},
	"io/ioutil": &starlark.Module{
		Name: "io/ioutil",
		Members: starlark.StringDict{
			"Discard": stargo.VarOf(&π3.Discard),
			"NopCloser": stargo.ValueOf(π3.NopCloser),
			"ReadAll": stargo.ValueOf(π3.ReadAll),
			"ReadDir": stargo.ValueOf(π3.ReadDir),
			"ReadFile": stargo.ValueOf(π3.ReadFile),
			"TempDir": stargo.ValueOf(π3.TempDir),
			"TempFile": stargo.ValueOf(π3.TempFile),
			"WriteFile": stargo.ValueOf(π3.WriteFile),
		},
	},
	"net/http": &starlark.Module{
		Name: "net/http",
		Members: starlark.StringDict{
			"CanonicalHeaderKey": stargo.ValueOf(π4.CanonicalHeaderKey),
			"Client": stargo.TypeOf(reflect.TypeOf(new(π4.Client)).Elem()),
			"CloseNotifier": stargo.TypeOf(reflect.TypeOf(new(π4.CloseNotifier)).Elem()),
			"ConnState": stargo.TypeOf(reflect.TypeOf(new(π4.ConnState)).Elem()),
			"Cookie": stargo.TypeOf(reflect.TypeOf(new(π4.Cookie)).Elem()),
			"CookieJar": stargo.TypeOf(reflect.TypeOf(new(π4.CookieJar)).Elem()),
			"DefaultClient": stargo.VarOf(&π4.DefaultClient),
			"DefaultMaxHeaderBytes": stargo.ValueOf(π4.DefaultMaxHeaderBytes),
			"DefaultMaxIdleConnsPerHost": stargo.ValueOf(π4.DefaultMaxIdleConnsPerHost),
			"DefaultServeMux": stargo.VarOf(&π4.DefaultServeMux),
			"DefaultTransport": stargo.VarOf(&π4.DefaultTransport),
			"DetectContentType": stargo.ValueOf(π4.DetectContentType),
			"Dir": stargo.TypeOf(reflect.TypeOf(new(π4.Dir)).Elem()),
			"ErrAbortHandler": stargo.VarOf(&π4.ErrAbortHandler),
			"ErrBodyNotAllowed": stargo.VarOf(&π4.ErrBodyNotAllowed),
			"ErrBodyReadAfterClose": stargo.VarOf(&π4.ErrBodyReadAfterClose),
			"ErrContentLength": stargo.VarOf(&π4.ErrContentLength),
			"ErrHandlerTimeout": stargo.VarOf(&π4.ErrHandlerTimeout),
			"ErrHeaderTooLong": stargo.VarOf(&π4.ErrHeaderTooLong),
			"ErrHijacked": stargo.VarOf(&π4.ErrHijacked),
			"ErrLineTooLong": stargo.VarOf(&π4.ErrLineTooLong),
			"ErrMissingBoundary": stargo.VarOf(&π4.ErrMissingBoundary),
			"ErrMissingContentLength": stargo.VarOf(&π4.ErrMissingContentLength),
			"ErrMissingFile": stargo.VarOf(&π4.ErrMissingFile),
			"ErrNoCookie": stargo.VarOf(&π4.ErrNoCookie),
			"ErrNoLocation": stargo.VarOf(&π4.ErrNoLocation),
			"ErrNotMultipart": stargo.VarOf(&π4.ErrNotMultipart),
			"ErrNotSupported": stargo.VarOf(&π4.ErrNotSupported),
			"ErrServerClosed": stargo.VarOf(&π4.ErrServerClosed),
			"ErrShortBody": stargo.VarOf(&π4.ErrShortBody),
			"ErrSkipAltProtocol": stargo.VarOf(&π4.ErrSkipAltProtocol),
			"ErrUnexpectedTrailer": stargo.VarOf(&π4.ErrUnexpectedTrailer),
			"ErrUseLastResponse": stargo.VarOf(&π4.ErrUseLastResponse),
			"ErrWriteAfterFlush": stargo.VarOf(&π4.ErrWriteAfterFlush),
			"Error": stargo.ValueOf(π4.Error),
			"File": stargo.TypeOf(reflect.TypeOf(new(π4.File)).Elem()),
			"FileServer": stargo.ValueOf(π4.FileServer),
			"FileSystem": stargo.TypeOf(reflect.TypeOf(new(π4.FileSystem)).Elem()),
			"Flusher": stargo.TypeOf(reflect.TypeOf(new(π4.Flusher)).Elem()),
			"Get": stargo.ValueOf(π4.Get),
			"Handle": stargo.ValueOf(π4.Handle),
			"HandleFunc": stargo.ValueOf(π4.HandleFunc),
			"Handler": stargo.TypeOf(reflect.TypeOf(new(π4.Handler)).Elem()),
			"HandlerFunc": stargo.TypeOf(reflect.TypeOf(new(π4.HandlerFunc)).Elem()),
			"Head": stargo.ValueOf(π4.Head),
			"Header": stargo.TypeOf(reflect.TypeOf(new(π4.Header)).Elem()),
			"Hijacker": stargo.TypeOf(reflect.TypeOf(new(π4.Hijacker)).Elem()),
			"ListenAndServe": stargo.ValueOf(π4.ListenAndServe),
			"ListenAndServeTLS": stargo.ValueOf(π4.ListenAndServeTLS),
			"LocalAddrContextKey": stargo.VarOf(&π4.LocalAddrContextKey),
			"MaxBytesReader": stargo.ValueOf(π4.MaxBytesReader),
			"MethodConnect": stargo.ValueOf(π4.MethodConnect),
			"MethodDelete": stargo.ValueOf(π4.MethodDelete),
			"MethodGet": stargo.ValueOf(π4.MethodGet),
			"MethodHead": stargo.ValueOf(π4.MethodHead),
			"MethodOptions": stargo.ValueOf(π4.MethodOptions),
			"MethodPatch": stargo.ValueOf(π4.MethodPatch),
			"MethodPost": stargo.ValueOf(π4.MethodPost),
			"MethodPut": stargo.ValueOf(π4.MethodPut),
			"MethodTrace": stargo.ValueOf(π4.MethodTrace),
			"NewFileTransport": stargo.ValueOf(π4.NewFileTransport),
			"NewRequest": stargo.ValueOf(π4.NewRequest),
			"NewServeMux": stargo.ValueOf(π4.NewServeMux),
			"NoBody": stargo.VarOf(&π4.NoBody),
			"NotFound": stargo.ValueOf(π4.NotFound),
			"NotFoundHandler": stargo.ValueOf(π4.NotFoundHandler),
			"ParseHTTPVersion": stargo.ValueOf(π4.ParseHTTPVersion),
			"ParseTime": stargo.ValueOf(π4.ParseTime),
			"Post": stargo.ValueOf(π4.Post),
			"PostForm": stargo.ValueOf(π4.PostForm),
			"ProtocolError": stargo.TypeOf(reflect.TypeOf(new(π4.ProtocolError)).Elem()),
			"ProxyFromEnvironment": stargo.ValueOf(π4.ProxyFromEnvironment),
			"ProxyURL": stargo.ValueOf(π4.ProxyURL),
			"PushOptions": stargo.TypeOf(reflect.TypeOf(new(π4.PushOptions)).Elem()),
			"Pusher": stargo.TypeOf(reflect.TypeOf(new(π4.Pusher)).Elem()),
			"ReadRequest": stargo.ValueOf(π4.ReadRequest),
			"ReadResponse": stargo.ValueOf(π4.ReadResponse),
			"Redirect": stargo.ValueOf(π4.Redirect),
			"RedirectHandler": stargo.ValueOf(π4.RedirectHandler),
			"Request": stargo.TypeOf(reflect.TypeOf(new(π4.Request)).Elem()),
			"Response": stargo.TypeOf(reflect.TypeOf(new(π4.Response)).Elem()),
			"ResponseWriter": stargo.TypeOf(reflect.TypeOf(new(π4.ResponseWriter)).Elem()),
			"RoundTripper": stargo.TypeOf(reflect.TypeOf(new(π4.RoundTripper)).Elem()),
			"SameSite": stargo.TypeOf(reflect.TypeOf(new(π4.SameSite)).Elem()),
			"SameSiteDefaultMode": stargo.ValueOf(π4.SameSiteDefaultMode),
			"SameSiteLaxMode": stargo.ValueOf(π4.SameSiteLaxMode),
			"SameSiteStrictMode": stargo.ValueOf(π4.SameSiteStrictMode),
			"Serve": stargo.ValueOf(π4.Serve),
			"ServeContent": stargo.ValueOf(π4.ServeContent),
			"ServeFile": stargo.ValueOf(π4.ServeFile),
			"ServeMux": stargo.TypeOf(reflect.TypeOf(new(π4.ServeMux)).Elem()),
			"ServeTLS": stargo.ValueOf(π4.ServeTLS),
			"Server": stargo.TypeOf(reflect.TypeOf(new(π4.Server)).Elem()),
			"ServerContextKey": stargo.VarOf(&π4.ServerContextKey),
			"SetCookie": stargo.ValueOf(π4.SetCookie),
			"StateActive": stargo.ValueOf(π4.StateActive),
			"StateClosed": stargo.ValueOf(π4.StateClosed),
			"StateHijacked": stargo.ValueOf(π4.StateHijacked),
			"StateIdle": stargo.ValueOf(π4.StateIdle),
			"StateNew": stargo.ValueOf(π4.StateNew),
			"StatusAccepted": stargo.ValueOf(π4.StatusAccepted),
			"StatusAlreadyReported": stargo.ValueOf(π4.StatusAlreadyReported),
			"StatusBadGateway": stargo.ValueOf(π4.StatusBadGateway),
			"StatusBadRequest": stargo.ValueOf(π4.StatusBadRequest),
			"StatusConflict": stargo.ValueOf(π4.StatusConflict),
			"StatusContinue": stargo.ValueOf(π4.StatusContinue),
			"StatusCreated": stargo.ValueOf(π4.StatusCreated),
			"StatusExpectationFailed": stargo.ValueOf(π4.StatusExpectationFailed),
			"StatusFailedDependency": stargo.ValueOf(π4.StatusFailedDependency),
			"StatusForbidden": stargo.ValueOf(π4.StatusForbidden),
			"StatusFound": stargo.ValueOf(π4.StatusFound),
			"StatusGatewayTimeout": stargo.ValueOf(π4.StatusGatewayTimeout),
			"StatusGone": stargo.ValueOf(π4.StatusGone),
			"StatusHTTPVersionNotSupported": stargo.ValueOf(π4.StatusHTTPVersionNotSupported),
			"StatusIMUsed": stargo.ValueOf(π4.StatusIMUsed),
			"StatusInsufficientStorage": stargo.ValueOf(π4.StatusInsufficientStorage),
			"StatusInternalServerError": stargo.ValueOf(π4.StatusInternalServerError),
			"StatusLengthRequired": stargo.ValueOf(π4.StatusLengthRequired),
			"StatusLocked": stargo.ValueOf(π4.StatusLocked),
			"StatusLoopDetected": stargo.ValueOf(π4.StatusLoopDetected),
			"StatusMethodNotAllowed": stargo.ValueOf(π4.StatusMethodNotAllowed),
			"StatusMisdirectedRequest": stargo.ValueOf(π4.StatusMisdirectedRequest),
			"StatusMovedPermanently": stargo.ValueOf(π4.StatusMovedPermanently),
			"StatusMultiStatus": stargo.ValueOf(π4.StatusMultiStatus),
			"StatusMultipleChoices": stargo.ValueOf(π4.StatusMultipleChoices),
			"StatusNetworkAuthenticationRequired": stargo.ValueOf(π4.StatusNetworkAuthenticationRequired),
			"StatusNoContent": stargo.ValueOf(π4.StatusNoContent),
			"StatusNonAuthoritativeInfo": stargo.ValueOf(π4.StatusNonAuthoritativeInfo),
			"StatusNotAcceptable": stargo.ValueOf(π4.StatusNotAcceptable),
			"StatusNotExtended": stargo.ValueOf(π4.StatusNotExtended),
			"StatusNotFound": stargo.ValueOf(π4.StatusNotFound),
			"StatusNotImplemented": stargo.ValueOf(π4.StatusNotImplemented),
			"StatusNotModified": stargo.ValueOf(π4.StatusNotModified),
			"StatusOK": stargo.ValueOf(π4.StatusOK),
			"StatusPartialContent": stargo.ValueOf(π4.StatusPartialContent),
			"StatusPaymentRequired": stargo.ValueOf(π4.StatusPaymentRequired),
			"StatusPermanentRedirect": stargo.ValueOf(π4.StatusPermanentRedirect),
			"StatusPreconditionFailed": stargo.ValueOf(π4.StatusPreconditionFailed),
			"StatusPreconditionRequired": stargo.ValueOf(π4.StatusPreconditionRequired),
			"StatusProcessing": stargo.ValueOf(π4.StatusProcessing),
			"StatusProxyAuthRequired": stargo.ValueOf(π4.StatusProxyAuthRequired),
			"StatusRequestEntityTooLarge": stargo.ValueOf(π4.StatusRequestEntityTooLarge),
			"StatusRequestHeaderFieldsTooLarge": stargo.ValueOf(π4.StatusRequestHeaderFieldsTooLarge),
			"StatusRequestTimeout": stargo.ValueOf(π4.StatusRequestTimeout),
			"StatusRequestURITooLong": stargo.ValueOf(π4.StatusRequestURITooLong),
			"StatusRequestedRangeNotSatisfiable": stargo.ValueOf(π4.StatusRequestedRangeNotSatisfiable),
			"StatusResetContent": stargo.ValueOf(π4.StatusResetContent),
			"StatusSeeOther": stargo.ValueOf(π4.StatusSeeOther),
			"StatusServiceUnavailable": stargo.ValueOf(π4.StatusServiceUnavailable),
			"StatusSwitchingProtocols": stargo.ValueOf(π4.StatusSwitchingProtocols),
			"StatusTeapot": stargo.ValueOf(π4.StatusTeapot),
			"StatusTemporaryRedirect": stargo.ValueOf(π4.StatusTemporaryRedirect),
			"StatusText": stargo.ValueOf(π4.StatusText),
			"StatusTooEarly": stargo.ValueOf(π4.StatusTooEarly),
			"StatusTooManyRequests": stargo.ValueOf(π4.StatusTooManyRequests),
			"StatusUnauthorized": stargo.ValueOf(π4.StatusUnauthorized),
			"StatusUnavailableForLegalReasons": stargo.ValueOf(π4.StatusUnavailableForLegalReasons),
			"StatusUnprocessableEntity": stargo.ValueOf(π4.StatusUnprocessableEntity),
			"StatusUnsupportedMediaType": stargo.ValueOf(π4.StatusUnsupportedMediaType),
			"StatusUpgradeRequired": stargo.ValueOf(π4.StatusUpgradeRequired),
			"StatusUseProxy": stargo.ValueOf(π4.StatusUseProxy),
			"StatusVariantAlsoNegotiates": stargo.ValueOf(π4.StatusVariantAlsoNegotiates),
			"StripPrefix": stargo.ValueOf(π4.StripPrefix),
			"TimeFormat": stargo.ValueOf(π4.TimeFormat),
			"TimeoutHandler": stargo.ValueOf(π4.TimeoutHandler),
			"TrailerPrefix": stargo.ValueOf(π4.TrailerPrefix),
			"Transport": stargo.TypeOf(reflect.TypeOf(new(π4.Transport)).Elem()),
		},
	},
	"go/token": &starlark.Module{
		Name: "go/token",
		Members: starlark.StringDict{
			"ADD": stargo.ValueOf(π5.ADD),
			"ADD_ASSIGN": stargo.ValueOf(π5.ADD_ASSIGN),
			"AND": stargo.ValueOf(π5.AND),
			"AND_ASSIGN": stargo.ValueOf(π5.AND_ASSIGN),
			"AND_NOT": stargo.ValueOf(π5.AND_NOT),
			"AND_NOT_ASSIGN": stargo.ValueOf(π5.AND_NOT_ASSIGN),
			"ARROW": stargo.ValueOf(π5.ARROW),
			"ASSIGN": stargo.ValueOf(π5.ASSIGN),
			"BREAK": stargo.ValueOf(π5.BREAK),
			"CASE": stargo.ValueOf(π5.CASE),
			"CHAN": stargo.ValueOf(π5.CHAN),
			"CHAR": stargo.ValueOf(π5.CHAR),
			"COLON": stargo.ValueOf(π5.COLON),
			"COMMA": stargo.ValueOf(π5.COMMA),
			"COMMENT": stargo.ValueOf(π5.COMMENT),
			"CONST": stargo.ValueOf(π5.CONST),
			"CONTINUE": stargo.ValueOf(π5.CONTINUE),
			"DEC": stargo.ValueOf(π5.DEC),
			"DEFAULT": stargo.ValueOf(π5.DEFAULT),
			"DEFER": stargo.ValueOf(π5.DEFER),
			"DEFINE": stargo.ValueOf(π5.DEFINE),
			"ELLIPSIS": stargo.ValueOf(π5.ELLIPSIS),
			"ELSE": stargo.ValueOf(π5.ELSE),
			"EOF": stargo.ValueOf(π5.EOF),
			"EQL": stargo.ValueOf(π5.EQL),
			"FALLTHROUGH": stargo.ValueOf(π5.FALLTHROUGH),
			"FLOAT": stargo.ValueOf(π5.FLOAT),
			"FOR": stargo.ValueOf(π5.FOR),
			"FUNC": stargo.ValueOf(π5.FUNC),
			"File": stargo.TypeOf(reflect.TypeOf(new(π5.File)).Elem()),
			"FileSet": stargo.TypeOf(reflect.TypeOf(new(π5.FileSet)).Elem()),
			"GEQ": stargo.ValueOf(π5.GEQ),
			"GO": stargo.ValueOf(π5.GO),
			"GOTO": stargo.ValueOf(π5.GOTO),
			"GTR": stargo.ValueOf(π5.GTR),
			"HighestPrec": stargo.ValueOf(π5.HighestPrec),
			"IDENT": stargo.ValueOf(π5.IDENT),
			"IF": stargo.ValueOf(π5.IF),
			"ILLEGAL": stargo.ValueOf(π5.ILLEGAL),
			"IMAG": stargo.ValueOf(π5.IMAG),
			"IMPORT": stargo.ValueOf(π5.IMPORT),
			"INC": stargo.ValueOf(π5.INC),
			"INT": stargo.ValueOf(π5.INT),
			"INTERFACE": stargo.ValueOf(π5.INTERFACE),
			"LAND": stargo.ValueOf(π5.LAND),
			"LBRACE": stargo.ValueOf(π5.LBRACE),
			"LBRACK": stargo.ValueOf(π5.LBRACK),
			"LEQ": stargo.ValueOf(π5.LEQ),
			"LOR": stargo.ValueOf(π5.LOR),
			"LPAREN": stargo.ValueOf(π5.LPAREN),
			"LSS": stargo.ValueOf(π5.LSS),
			"Lookup": stargo.ValueOf(π5.Lookup),
			"LowestPrec": stargo.ValueOf(π5.LowestPrec),
			"MAP": stargo.ValueOf(π5.MAP),
			"MUL": stargo.ValueOf(π5.MUL),
			"MUL_ASSIGN": stargo.ValueOf(π5.MUL_ASSIGN),
			"NEQ": stargo.ValueOf(π5.NEQ),
			"NOT": stargo.ValueOf(π5.NOT),
			"NewFileSet": stargo.ValueOf(π5.NewFileSet),
			"NoPos": stargo.ValueOf(π5.NoPos),
			"OR": stargo.ValueOf(π5.OR),
			"OR_ASSIGN": stargo.ValueOf(π5.OR_ASSIGN),
			"PACKAGE": stargo.ValueOf(π5.PACKAGE),
			"PERIOD": stargo.ValueOf(π5.PERIOD),
			"Pos": stargo.TypeOf(reflect.TypeOf(new(π5.Pos)).Elem()),
			"Position": stargo.TypeOf(reflect.TypeOf(new(π5.Position)).Elem()),
			"QUO": stargo.ValueOf(π5.QUO),
			"QUO_ASSIGN": stargo.ValueOf(π5.QUO_ASSIGN),
			"RANGE": stargo.ValueOf(π5.RANGE),
			"RBRACE": stargo.ValueOf(π5.RBRACE),
			"RBRACK": stargo.ValueOf(π5.RBRACK),
			"REM": stargo.ValueOf(π5.REM),
			"REM_ASSIGN": stargo.ValueOf(π5.REM_ASSIGN),
			"RETURN": stargo.ValueOf(π5.RETURN),
			"RPAREN": stargo.ValueOf(π5.RPAREN),
			"SELECT": stargo.ValueOf(π5.SELECT),
			"SEMICOLON": stargo.ValueOf(π5.SEMICOLON),
			"SHL": stargo.ValueOf(π5.SHL),
			"SHL_ASSIGN": stargo.ValueOf(π5.SHL_ASSIGN),
			"SHR": stargo.ValueOf(π5.SHR),
			"SHR_ASSIGN": stargo.ValueOf(π5.SHR_ASSIGN),
			"STRING": stargo.ValueOf(π5.STRING),
			"STRUCT": stargo.ValueOf(π5.STRUCT),
			"SUB": stargo.ValueOf(π5.SUB),
			"SUB_ASSIGN": stargo.ValueOf(π5.SUB_ASSIGN),
			"SWITCH": stargo.ValueOf(π5.SWITCH),
			"TYPE": stargo.ValueOf(π5.TYPE),
			"Token": stargo.TypeOf(reflect.TypeOf(new(π5.Token)).Elem()),
			"UnaryPrec": stargo.ValueOf(π5.UnaryPrec),
			"VAR": stargo.ValueOf(π5.VAR),
			"XOR": stargo.ValueOf(π5.XOR),
			"XOR_ASSIGN": stargo.ValueOf(π5.XOR_ASSIGN),
		},
	},
	"go/ast": &starlark.Module{
		Name: "go/ast",
		Members: starlark.StringDict{
			"ArrayType": stargo.TypeOf(reflect.TypeOf(new(π6.ArrayType)).Elem()),
			"AssignStmt": stargo.TypeOf(reflect.TypeOf(new(π6.AssignStmt)).Elem()),
			"Bad": stargo.ValueOf(π6.Bad),
			"BadDecl": stargo.TypeOf(reflect.TypeOf(new(π6.BadDecl)).Elem()),
			"BadExpr": stargo.TypeOf(reflect.TypeOf(new(π6.BadExpr)).Elem()),
			"BadStmt": stargo.TypeOf(reflect.TypeOf(new(π6.BadStmt)).Elem()),
			"BasicLit": stargo.TypeOf(reflect.TypeOf(new(π6.BasicLit)).Elem()),
			"BinaryExpr": stargo.TypeOf(reflect.TypeOf(new(π6.BinaryExpr)).Elem()),
			"BlockStmt": stargo.TypeOf(reflect.TypeOf(new(π6.BlockStmt)).Elem()),
			"BranchStmt": stargo.TypeOf(reflect.TypeOf(new(π6.BranchStmt)).Elem()),
			"CallExpr": stargo.TypeOf(reflect.TypeOf(new(π6.CallExpr)).Elem()),
			"CaseClause": stargo.TypeOf(reflect.TypeOf(new(π6.CaseClause)).Elem()),
			"ChanDir": stargo.TypeOf(reflect.TypeOf(new(π6.ChanDir)).Elem()),
			"ChanType": stargo.TypeOf(reflect.TypeOf(new(π6.ChanType)).Elem()),
			"CommClause": stargo.TypeOf(reflect.TypeOf(new(π6.CommClause)).Elem()),
			"Comment": stargo.TypeOf(reflect.TypeOf(new(π6.Comment)).Elem()),
			"CommentGroup": stargo.TypeOf(reflect.TypeOf(new(π6.CommentGroup)).Elem()),
			"CommentMap": stargo.TypeOf(reflect.TypeOf(new(π6.CommentMap)).Elem()),
			"CompositeLit": stargo.TypeOf(reflect.TypeOf(new(π6.CompositeLit)).Elem()),
			"Con": stargo.ValueOf(π6.Con),
			"Decl": stargo.TypeOf(reflect.TypeOf(new(π6.Decl)).Elem()),
			"DeclStmt": stargo.TypeOf(reflect.TypeOf(new(π6.DeclStmt)).Elem()),
			"DeferStmt": stargo.TypeOf(reflect.TypeOf(new(π6.DeferStmt)).Elem()),
			"Ellipsis": stargo.TypeOf(reflect.TypeOf(new(π6.Ellipsis)).Elem()),
			"EmptyStmt": stargo.TypeOf(reflect.TypeOf(new(π6.EmptyStmt)).Elem()),
			"Expr": stargo.TypeOf(reflect.TypeOf(new(π6.Expr)).Elem()),
			"ExprStmt": stargo.TypeOf(reflect.TypeOf(new(π6.ExprStmt)).Elem()),
			"Field": stargo.TypeOf(reflect.TypeOf(new(π6.Field)).Elem()),
			"FieldFilter": stargo.TypeOf(reflect.TypeOf(new(π6.FieldFilter)).Elem()),
			"FieldList": stargo.TypeOf(reflect.TypeOf(new(π6.FieldList)).Elem()),
			"File": stargo.TypeOf(reflect.TypeOf(new(π6.File)).Elem()),
			"FileExports": stargo.ValueOf(π6.FileExports),
			"Filter": stargo.TypeOf(reflect.TypeOf(new(π6.Filter)).Elem()),
			"FilterDecl": stargo.ValueOf(π6.FilterDecl),
			"FilterFile": stargo.ValueOf(π6.FilterFile),
			"FilterFuncDuplicates": stargo.ValueOf(π6.FilterFuncDuplicates),
			"FilterImportDuplicates": stargo.ValueOf(π6.FilterImportDuplicates),
			"FilterPackage": stargo.ValueOf(π6.FilterPackage),
			"FilterUnassociatedComments": stargo.ValueOf(π6.FilterUnassociatedComments),
			"ForStmt": stargo.TypeOf(reflect.TypeOf(new(π6.ForStmt)).Elem()),
			"Fprint": stargo.ValueOf(π6.Fprint),
			"Fun": stargo.ValueOf(π6.Fun),
			"FuncDecl": stargo.TypeOf(reflect.TypeOf(new(π6.FuncDecl)).Elem()),
			"FuncLit": stargo.TypeOf(reflect.TypeOf(new(π6.FuncLit)).Elem()),
			"FuncType": stargo.TypeOf(reflect.TypeOf(new(π6.FuncType)).Elem()),
			"GenDecl": stargo.TypeOf(reflect.TypeOf(new(π6.GenDecl)).Elem()),
			"GoStmt": stargo.TypeOf(reflect.TypeOf(new(π6.GoStmt)).Elem()),
			"Ident": stargo.TypeOf(reflect.TypeOf(new(π6.Ident)).Elem()),
			"IfStmt": stargo.TypeOf(reflect.TypeOf(new(π6.IfStmt)).Elem()),
			"ImportSpec": stargo.TypeOf(reflect.TypeOf(new(π6.ImportSpec)).Elem()),
			"Importer": stargo.TypeOf(reflect.TypeOf(new(π6.Importer)).Elem()),
			"IncDecStmt": stargo.TypeOf(reflect.TypeOf(new(π6.IncDecStmt)).Elem()),
			"IndexExpr": stargo.TypeOf(reflect.TypeOf(new(π6.IndexExpr)).Elem()),
			"Inspect": stargo.ValueOf(π6.Inspect),
			"InterfaceType": stargo.TypeOf(reflect.TypeOf(new(π6.InterfaceType)).Elem()),
			"IsExported": stargo.ValueOf(π6.IsExported),
			"KeyValueExpr": stargo.TypeOf(reflect.TypeOf(new(π6.KeyValueExpr)).Elem()),
			"LabeledStmt": stargo.TypeOf(reflect.TypeOf(new(π6.LabeledStmt)).Elem()),
			"Lbl": stargo.ValueOf(π6.Lbl),
			"MapType": stargo.TypeOf(reflect.TypeOf(new(π6.MapType)).Elem()),
			"MergeMode": stargo.TypeOf(reflect.TypeOf(new(π6.MergeMode)).Elem()),
			"MergePackageFiles": stargo.ValueOf(π6.MergePackageFiles),
			"NewCommentMap": stargo.ValueOf(π6.NewCommentMap),
			"NewIdent": stargo.ValueOf(π6.NewIdent),
			"NewObj": stargo.ValueOf(π6.NewObj),
			"NewPackage": stargo.ValueOf(π6.NewPackage),
			"NewScope": stargo.ValueOf(π6.NewScope),
			"Node": stargo.TypeOf(reflect.TypeOf(new(π6.Node)).Elem()),
			"NotNilFilter": stargo.ValueOf(π6.NotNilFilter),
			"ObjKind": stargo.TypeOf(reflect.TypeOf(new(π6.ObjKind)).Elem()),
			"Object": stargo.TypeOf(reflect.TypeOf(new(π6.Object)).Elem()),
			"Package": stargo.TypeOf(reflect.TypeOf(new(π6.Package)).Elem()),
			"PackageExports": stargo.ValueOf(π6.PackageExports),
			"ParenExpr": stargo.TypeOf(reflect.TypeOf(new(π6.ParenExpr)).Elem()),
			"Pkg": stargo.ValueOf(π6.Pkg),
			"Print": stargo.ValueOf(π6.Print),
			"RECV": stargo.ValueOf(π6.RECV),
			"RangeStmt": stargo.TypeOf(reflect.TypeOf(new(π6.RangeStmt)).Elem()),
			"ReturnStmt": stargo.TypeOf(reflect.TypeOf(new(π6.ReturnStmt)).Elem()),
			"SEND": stargo.ValueOf(π6.SEND),
			"Scope": stargo.TypeOf(reflect.TypeOf(new(π6.Scope)).Elem()),
			"SelectStmt": stargo.TypeOf(reflect.TypeOf(new(π6.SelectStmt)).Elem()),
			"SelectorExpr": stargo.TypeOf(reflect.TypeOf(new(π6.SelectorExpr)).Elem()),
			"SendStmt": stargo.TypeOf(reflect.TypeOf(new(π6.SendStmt)).Elem()),
			"SliceExpr": stargo.TypeOf(reflect.TypeOf(new(π6.SliceExpr)).Elem()),
			"SortImports": stargo.ValueOf(π6.SortImports),
			"Spec": stargo.TypeOf(reflect.TypeOf(new(π6.Spec)).Elem()),
			"StarExpr": stargo.TypeOf(reflect.TypeOf(new(π6.StarExpr)).Elem()),
			"Stmt": stargo.TypeOf(reflect.TypeOf(new(π6.Stmt)).Elem()),
			"StructType": stargo.TypeOf(reflect.TypeOf(new(π6.StructType)).Elem()),
			"SwitchStmt": stargo.TypeOf(reflect.TypeOf(new(π6.SwitchStmt)).Elem()),
			"Typ": stargo.ValueOf(π6.Typ),
			"TypeAssertExpr": stargo.TypeOf(reflect.TypeOf(new(π6.TypeAssertExpr)).Elem()),
			"TypeSpec": stargo.TypeOf(reflect.TypeOf(new(π6.TypeSpec)).Elem()),
			"TypeSwitchStmt": stargo.TypeOf(reflect.TypeOf(new(π6.TypeSwitchStmt)).Elem()),
			"UnaryExpr": stargo.TypeOf(reflect.TypeOf(new(π6.UnaryExpr)).Elem()),
			"ValueSpec": stargo.TypeOf(reflect.TypeOf(new(π6.ValueSpec)).Elem()),
			"Var": stargo.ValueOf(π6.Var),
			"Visitor": stargo.TypeOf(reflect.TypeOf(new(π6.Visitor)).Elem()),
			"Walk": stargo.ValueOf(π6.Walk),
		},
	},
	"go/parser": &starlark.Module{
		Name: "go/parser",
		Members: starlark.StringDict{
			"AllErrors": stargo.ValueOf(π7.AllErrors),
			"DeclarationErrors": stargo.ValueOf(π7.DeclarationErrors),
			"ImportsOnly": stargo.ValueOf(π7.ImportsOnly),
			"Mode": stargo.TypeOf(reflect.TypeOf(new(π7.Mode)).Elem()),
			"PackageClauseOnly": stargo.ValueOf(π7.PackageClauseOnly),
			"ParseComments": stargo.ValueOf(π7.ParseComments),
			"ParseDir": stargo.ValueOf(π7.ParseDir),
			"ParseExpr": stargo.ValueOf(π7.ParseExpr),
			"ParseExprFrom": stargo.ValueOf(π7.ParseExprFrom),
			"ParseFile": stargo.ValueOf(π7.ParseFile),
			"SpuriousErrors": stargo.ValueOf(π7.SpuriousErrors),
			"Trace": stargo.ValueOf(π7.Trace),
		},
	},
	"go/types": &starlark.Module{
		Name: "go/types",
		Members: starlark.StringDict{
			"Array": stargo.TypeOf(reflect.TypeOf(new(π8.Array)).Elem()),
			"AssertableTo": stargo.ValueOf(π8.AssertableTo),
			"AssignableTo": stargo.ValueOf(π8.AssignableTo),
			"Basic": stargo.TypeOf(reflect.TypeOf(new(π8.Basic)).Elem()),
			"BasicInfo": stargo.TypeOf(reflect.TypeOf(new(π8.BasicInfo)).Elem()),
			"BasicKind": stargo.TypeOf(reflect.TypeOf(new(π8.BasicKind)).Elem()),
			"Bool": stargo.ValueOf(π8.Bool),
			"Builtin": stargo.TypeOf(reflect.TypeOf(new(π8.Builtin)).Elem()),
			"Byte": stargo.ValueOf(π8.Byte),
			"Chan": stargo.TypeOf(reflect.TypeOf(new(π8.Chan)).Elem()),
			"ChanDir": stargo.TypeOf(reflect.TypeOf(new(π8.ChanDir)).Elem()),
			"Checker": stargo.TypeOf(reflect.TypeOf(new(π8.Checker)).Elem()),
			"Comparable": stargo.ValueOf(π8.Comparable),
			"Complex128": stargo.ValueOf(π8.Complex128),
			"Complex64": stargo.ValueOf(π8.Complex64),
			"Config": stargo.TypeOf(reflect.TypeOf(new(π8.Config)).Elem()),
			"Const": stargo.TypeOf(reflect.TypeOf(new(π8.Const)).Elem()),
			"ConvertibleTo": stargo.ValueOf(π8.ConvertibleTo),
			"DefPredeclaredTestFuncs": stargo.ValueOf(π8.DefPredeclaredTestFuncs),
			"Default": stargo.ValueOf(π8.Default),
			"Error": stargo.TypeOf(reflect.TypeOf(new(π8.Error)).Elem()),
			"Eval": stargo.ValueOf(π8.Eval),
			"ExprString": stargo.ValueOf(π8.ExprString),
			"FieldVal": stargo.ValueOf(π8.FieldVal),
			"Float32": stargo.ValueOf(π8.Float32),
			"Float64": stargo.ValueOf(π8.Float64),
			"Func": stargo.TypeOf(reflect.TypeOf(new(π8.Func)).Elem()),
			"Id": stargo.ValueOf(π8.Id),
			"Identical": stargo.ValueOf(π8.Identical),
			"IdenticalIgnoreTags": stargo.ValueOf(π8.IdenticalIgnoreTags),
			"Implements": stargo.ValueOf(π8.Implements),
			"ImportMode": stargo.TypeOf(reflect.TypeOf(new(π8.ImportMode)).Elem()),
			"Importer": stargo.TypeOf(reflect.TypeOf(new(π8.Importer)).Elem()),
			"ImporterFrom": stargo.TypeOf(reflect.TypeOf(new(π8.ImporterFrom)).Elem()),
			"Info": stargo.TypeOf(reflect.TypeOf(new(π8.Info)).Elem()),
			"Initializer": stargo.TypeOf(reflect.TypeOf(new(π8.Initializer)).Elem()),
			"Int": stargo.ValueOf(π8.Int),
			"Int16": stargo.ValueOf(π8.Int16),
			"Int32": stargo.ValueOf(π8.Int32),
			"Int64": stargo.ValueOf(π8.Int64),
			"Int8": stargo.ValueOf(π8.Int8),
			"Interface": stargo.TypeOf(reflect.TypeOf(new(π8.Interface)).Elem()),
			"Invalid": stargo.ValueOf(π8.Invalid),
			"IsBoolean": stargo.ValueOf(π8.IsBoolean),
			"IsComplex": stargo.ValueOf(π8.IsComplex),
			"IsConstType": stargo.ValueOf(π8.IsConstType),
			"IsFloat": stargo.ValueOf(π8.IsFloat),
			"IsInteger": stargo.ValueOf(π8.IsInteger),
			"IsInterface": stargo.ValueOf(π8.IsInterface),
			"IsNumeric": stargo.ValueOf(π8.IsNumeric),
			"IsOrdered": stargo.ValueOf(π8.IsOrdered),
			"IsString": stargo.ValueOf(π8.IsString),
			"IsUnsigned": stargo.ValueOf(π8.IsUnsigned),
			"IsUntyped": stargo.ValueOf(π8.IsUntyped),
			"Label": stargo.TypeOf(reflect.TypeOf(new(π8.Label)).Elem()),
			"LookupFieldOrMethod": stargo.ValueOf(π8.LookupFieldOrMethod),
			"Map": stargo.TypeOf(reflect.TypeOf(new(π8.Map)).Elem()),
			"MethodExpr": stargo.ValueOf(π8.MethodExpr),
			"MethodSet": stargo.TypeOf(reflect.TypeOf(new(π8.MethodSet)).Elem()),
			"MethodVal": stargo.ValueOf(π8.MethodVal),
			"MissingMethod": stargo.ValueOf(π8.MissingMethod),
			"Named": stargo.TypeOf(reflect.TypeOf(new(π8.Named)).Elem()),
			"NewArray": stargo.ValueOf(π8.NewArray),
			"NewChan": stargo.ValueOf(π8.NewChan),
			"NewChecker": stargo.ValueOf(π8.NewChecker),
			"NewConst": stargo.ValueOf(π8.NewConst),
			"NewField": stargo.ValueOf(π8.NewField),
			"NewFunc": stargo.ValueOf(π8.NewFunc),
			"NewInterface": stargo.ValueOf(π8.NewInterface),
			"NewInterfaceType": stargo.ValueOf(π8.NewInterfaceType),
			"NewLabel": stargo.ValueOf(π8.NewLabel),
			"NewMap": stargo.ValueOf(π8.NewMap),
			"NewMethodSet": stargo.ValueOf(π8.NewMethodSet),
			"NewNamed": stargo.ValueOf(π8.NewNamed),
			"NewPackage": stargo.ValueOf(π8.NewPackage),
			"NewParam": stargo.ValueOf(π8.NewParam),
			"NewPkgName": stargo.ValueOf(π8.NewPkgName),
			"NewPointer": stargo.ValueOf(π8.NewPointer),
			"NewScope": stargo.ValueOf(π8.NewScope),
			"NewSignature": stargo.ValueOf(π8.NewSignature),
			"NewSlice": stargo.ValueOf(π8.NewSlice),
			"NewStruct": stargo.ValueOf(π8.NewStruct),
			"NewTuple": stargo.ValueOf(π8.NewTuple),
			"NewTypeName": stargo.ValueOf(π8.NewTypeName),
			"NewVar": stargo.ValueOf(π8.NewVar),
			"Nil": stargo.TypeOf(reflect.TypeOf(new(π8.Nil)).Elem()),
			"Object": stargo.TypeOf(reflect.TypeOf(new(π8.Object)).Elem()),
			"ObjectString": stargo.ValueOf(π8.ObjectString),
			"Package": stargo.TypeOf(reflect.TypeOf(new(π8.Package)).Elem()),
			"PkgName": stargo.TypeOf(reflect.TypeOf(new(π8.PkgName)).Elem()),
			"Pointer": stargo.TypeOf(reflect.TypeOf(new(π8.Pointer)).Elem()),
			"Qualifier": stargo.TypeOf(reflect.TypeOf(new(π8.Qualifier)).Elem()),
			"RecvOnly": stargo.ValueOf(π8.RecvOnly),
			"RelativeTo": stargo.ValueOf(π8.RelativeTo),
			"Rune": stargo.ValueOf(π8.Rune),
			"Scope": stargo.TypeOf(reflect.TypeOf(new(π8.Scope)).Elem()),
			"Selection": stargo.TypeOf(reflect.TypeOf(new(π8.Selection)).Elem()),
			"SelectionKind": stargo.TypeOf(reflect.TypeOf(new(π8.SelectionKind)).Elem()),
			"SelectionString": stargo.ValueOf(π8.SelectionString),
			"SendOnly": stargo.ValueOf(π8.SendOnly),
			"SendRecv": stargo.ValueOf(π8.SendRecv),
			"Signature": stargo.TypeOf(reflect.TypeOf(new(π8.Signature)).Elem()),
			"Sizes": stargo.TypeOf(reflect.TypeOf(new(π8.Sizes)).Elem()),
			"SizesFor": stargo.ValueOf(π8.SizesFor),
			"Slice": stargo.TypeOf(reflect.TypeOf(new(π8.Slice)).Elem()),
			"StdSizes": stargo.TypeOf(reflect.TypeOf(new(π8.StdSizes)).Elem()),
			"String": stargo.ValueOf(π8.String),
			"Struct": stargo.TypeOf(reflect.TypeOf(new(π8.Struct)).Elem()),
			"Tuple": stargo.TypeOf(reflect.TypeOf(new(π8.Tuple)).Elem()),
			"Typ": stargo.VarOf(&π8.Typ),
			"Type": stargo.TypeOf(reflect.TypeOf(new(π8.Type)).Elem()),
			"TypeAndValue": stargo.TypeOf(reflect.TypeOf(new(π8.TypeAndValue)).Elem()),
			"TypeName": stargo.TypeOf(reflect.TypeOf(new(π8.TypeName)).Elem()),
			"TypeString": stargo.ValueOf(π8.TypeString),
			"Uint": stargo.ValueOf(π8.Uint),
			"Uint16": stargo.ValueOf(π8.Uint16),
			"Uint32": stargo.ValueOf(π8.Uint32),
			"Uint64": stargo.ValueOf(π8.Uint64),
			"Uint8": stargo.ValueOf(π8.Uint8),
			"Uintptr": stargo.ValueOf(π8.Uintptr),
			"Universe": stargo.VarOf(&π8.Universe),
			"Unsafe": stargo.VarOf(&π8.Unsafe),
			"UnsafePointer": stargo.ValueOf(π8.UnsafePointer),
			"UntypedBool": stargo.ValueOf(π8.UntypedBool),
			"UntypedComplex": stargo.ValueOf(π8.UntypedComplex),
			"UntypedFloat": stargo.ValueOf(π8.UntypedFloat),
			"UntypedInt": stargo.ValueOf(π8.UntypedInt),
			"UntypedNil": stargo.ValueOf(π8.UntypedNil),
			"UntypedRune": stargo.ValueOf(π8.UntypedRune),
			"UntypedString": stargo.ValueOf(π8.UntypedString),
			"Var": stargo.TypeOf(reflect.TypeOf(new(π8.Var)).Elem()),
			"WriteExpr": stargo.ValueOf(π8.WriteExpr),
			"WriteSignature": stargo.ValueOf(π8.WriteSignature),
			"WriteType": stargo.ValueOf(π8.WriteType),
		},
	},
}
