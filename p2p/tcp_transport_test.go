package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTCPTransport(t *testing.T) {
	type args struct {
		listenAddr string
	}
	tests := []struct {
		name             string
		args             args
		want             Transport
		wantedAssertFunc func(tr *TCPTransport)
	}{
		{
			name: "TestNewTCPTransport",
			args: args{
				listenAddr: ":8080",
			},
			// want: NewTCPTransport("127.0.0.1:8080"),
			wantedAssertFunc: func(tr *TCPTransport) {
				assert.Equalf(t, tr.listenAddress, ":8080", "Unexpected listenAddress")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTCPTransport(tt.args.listenAddr)

			tt.wantedAssertFunc(got)
		})
	}
}

func TestNewTCPTransportServe(t *testing.T) {

	listenAddr := ":8080"
	tr := NewTCPTransport(listenAddr)

	// Server
	assert.Nil(t, tr.ListenAndAccept())

	select {}

}
