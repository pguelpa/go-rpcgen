package main

import (
	"flag"
	"log"
	"testing"
	"net/url"

	"github.com/kylelemons/go-rpcgen/webrpc"
	"github.com/kylelemons/go-rpcgen/examples/remote/offload"
)

var base = flag.String("base", "http://localhost:9999/", "RPC server base URL")

func TestOffload(t *testing.T) {
	tests := []struct{
		In  string
		Out string
	}{
		{"abcd", "dcba"},
		{"racecar", "racecar"},
		{"", ""},
	}

	protos := []webrpc.Protocol{
		webrpc.JSON,
		webrpc.ProtoBuf,
	}

	flag.Parse()
	go main()

	url, err := url.Parse(*base)
	if err != nil {
		log.Fatalf("url: %s", err)
	}

	do := func(pro webrpc.Protocol, s string) string {
		in := &offload.DataSet{Data: &s}
		out := &offload.ResultSet{}

		off := offload.NewOffloadServiceWebClient(pro, url)
		if err := off.Compute(in, out); err != nil {
			log.Fatalf("compute(%q) - %s", s, err)
		}
		if out.Result == nil {
			t.Errorf("compute(%q) returned no result", s)
		}
		return *out.Result
	}

	for _, test := range tests {
		for _, pro := range protos {
			if got, want := do(pro, test.In), test.Out; got != want {
				t.Errorf("compute[%s](%q) = %q, want %q", pro, test.In, got, want)
			}
		}
	}
}