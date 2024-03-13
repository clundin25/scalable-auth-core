package main

/*
#include <stdlib.h>
typedef const char cchar;
*/
import "C"

import (
	"context"
	"fmt"
	"os"
	"strings"
	"unsafe"

	"golang.org/x/oauth2/google"
)

//export CreateAccessToken
func CreateAccessToken(uri *C.cchar, scopes *C.cchar, token *C.char, tokenLen *C.ulong) C.int {
	if uri == nil {
		fmt.Fprintln(os.Stderr, "'uri' cannot be NULL!")
		return 0
	}

	if scopes == nil {
		fmt.Fprintln(os.Stderr, "'scopes' cannot be NULL!")
		return 0
	}

	if tokenLen == nil {
		fmt.Fprintln(os.Stderr, "'token_len' cannot be NULL!")
		return 0
	}

	_ = C.GoString(uri)
	scopesStr := C.GoString(scopes)
	scopesList := strings.Split(scopesStr, ",")

	ctx := context.Background()

	c, err := google.FindDefaultCredentials(ctx, scopesList...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed due to %v.", err)
	}

	accessToken, err := c.TokenSource.Token()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create token due to %v.", err)
		return 0
	}

	tokenBytesLen := C.ulong(len(accessToken.AccessToken))

	if token == nil {

		*tokenLen = tokenBytesLen
		return 1
	}

	if *tokenLen < tokenBytesLen {
		fmt.Fprintf(os.Stderr, "Token buffer is too small")
		return 0
	}

	outBytes := unsafe.Slice(token, tokenBytesLen)
	tokenBytes := unsafe.Slice(C.CString(accessToken.AccessToken), tokenBytesLen)
	copy(outBytes, tokenBytes)

	*tokenLen = tokenBytesLen
	return 1
}

func main() {} // Required for CGo compilation
