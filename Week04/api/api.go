package api
/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2020/12/16
 @Author  : bishop
 @Software: GoLand
 @Description:
*/
import (
	"Go-000/Week04/service/logic"
	"context"
	"errors"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	service *logic.Service
}

func NewServer(s *logic.Service) *Server {
	return &Server{service: s}
}

func (srv *Server) Run() error {
	lis, err := net.Listen("tcp", viper.GetString("grpc.port"))
	if err != nil {
		return err
	}
	g, ctx := errgroup.WithContext(context.Background())
	s := grpc.NewServer()
	g.Go(func() error {
		go func() {
			<-ctx.Done()
			s.GracefulStop()
			log.Printf("Shutdown Server")
		}()
		//pb.GetUser(s, srv.service)
		return s.Serve(lis)
	})
	g.Go(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		for {
			select {
			case <-ctx.Done():
				return nil
			case s := <-c:
				log.Printf("get a signal %s", s.String())
				switch s {
				case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
					return errors.New("Close by signal " + s.String())
				case syscall.SIGHUP:
				default:
					return errors.New("undefined signal")
				}
			}
		}
	})
	return g.Wait()
}