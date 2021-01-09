// Package tgtest provides test Telegram server for end-to-end test.
package tgtest

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"net"

	"go.uber.org/zap"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/clock"
	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/internal/proto"
	"github.com/gotd/td/transport"
)

type Server struct {
	server *transport.Server

	key     *rsa.PrivateKey
	cipher  crypto.Cipher
	handler Handler

	clock  clock.Clock
	ctx    context.Context
	cancel context.CancelFunc

	log   *zap.Logger
	msgID *proto.MessageIDGen

	users *users
}

func (s *Server) Key() *rsa.PublicKey {
	return &s.key.PublicKey
}

func (s *Server) Addr() net.Addr {
	return s.server.Addr()
}

func (s *Server) Serve() error {
	return s.serve()
}

func (s *Server) AddSession(key crypto.AuthKey) {
	s.users.addSession(key)
}

func (s *Server) Start() {
	go func() {
		_ = s.Serve()
	}()
}

func (s *Server) Close() {
	if s.cancel != nil {
		s.cancel()
	}

	_ = s.server.Close()
}

func NewServer(name string, suite Suite, codec func() transport.Codec, h Handler) *Server {
	s := NewUnstartedServer(name, suite, codec)
	s.SetHandler(h)
	s.Start()
	return s
}

func NewUnstartedServer(name string, suite Suite, codec func() transport.Codec) *Server {
	k, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	log := suite.Log.Named(name)

	ctx, cancel := context.WithCancel(suite.Ctx)
	s := &Server{
		server: transport.NewCustomServer(codec, newLocalListener(ctx)),
		key:    k,
		cipher: crypto.NewServerCipher(rand.Reader),
		clock:  clock.System,
		ctx:    ctx,
		cancel: cancel,
		log:    log,
		users:  newUsers(),
		msgID:  proto.NewMessageIDGen(clock.System.Now, 100),
	}
	return s
}

func (s *Server) SetHandler(handler Handler) {
	s.handler = handler
}

func (s *Server) SetHandlerFunc(handler func(s Session, msgID int64, in *bin.Buffer) error) {
	s.handler = HandlerFunc(handler)
}

func (s *Server) serve() error {
	s.log.With(zap.String("addr", s.Addr().String())).Info("Serving")
	defer func() {
		l := s.log
		if err := s.ctx.Err(); err != nil {
			l = s.log.With(zap.Error(err))
		}
		l.Info("Stopping")
	}()

	return s.server.Serve(s.ctx, func(ctx context.Context, conn transport.Conn) error {
		err := s.serveConn(ctx, conn)
		if err != nil {
			s.log.With(zap.Error(err)).Info("Serving handler error")
		}
		return err
	})
}
