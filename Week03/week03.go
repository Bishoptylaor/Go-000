package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

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
 @Time    : 2020/12/9
 @Author  : bishop
 @Software: GoLand
 @Description:
*/
type Server struct {
	srv *http.Server
	env string
}

func Debug(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello world, i am debugging, need lots of time!")
	time.Sleep(time.Second * 10)
}

func Week03(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello world, i am online!")
}

func NewServer(env string) *Server {
	mux := http.NewServeMux()
	switch env {
	case "Debug":
		mux.Handle("/debug", http.HandlerFunc(Debug))
		return &Server{srv: &http.Server{
			Addr:    "0.0.0.0:8080",
			Handler: mux,
		}, env: env}
	case "Week03":
		mux.Handle("/week03", http.HandlerFunc(Week03))
		return &Server{srv: &http.Server{
			Addr:    "0.0.0.0:8081",
			Handler: mux,
		}, env: env}
	default:
		mux.Handle("/", http.HandlerFunc(Debug))
		return &Server{srv: &http.Server{
			Addr:    "0.0.0.0:6666",
			Handler: mux,
		}, env: env}
	}
}

func (s *Server) Start(ctx context.Context) error {
	fmt.Println("starting ", s.env)
	defer func() {
		if err := recover(); err != nil {
			_ = s.Stop(ctx)
		}
	}()
	go func() {
		<-ctx.Done()
		fmt.Println(s.srv.Addr, "has done")
		if err := s.Stop(ctx); err != nil {
			fmt.Println("Forced stop by system.")
		}
	}()
	return s.srv.ListenAndServe()
}
func (s *Server) Stop(ctx context.Context) error {
	fmt.Println("Stopping :", s.srv.Addr)
	return s.srv.Shutdown(ctx)
}

func main(){
	ctx, cancel := context.WithCancel(context.Background())
	eg, _ := errgroup.WithContext(ctx)
	defer func() {
		fmt.Println("all done")
		cancel()
	}()

	eg.Go(func() error {
		ctxTimeOut, cancel := context.WithTimeout(ctx, time.Second * 3)
		defer cancel()
		return NewServer("Debug").Start(ctxTimeOut)
	})
	eg.Go(func() error {
		return NewServer("Week03").Start(ctx)
	})

	// 监听signal
	eg.Go(func() error {
		s := make(chan os.Signal)
		signal.Notify(s, os.Interrupt)
		for{
			select {
			case <-ctx.Done():
				fmt.Println("signal ctx has done")
				return ctx.Err()
			case sgi := <-s:
				fmt.Println("receiving signal")
				cancel()
				return fmt.Errorf("handling err:%s", sgi)
			}
		}
	})
	log.Printf("This pid is %d", syscall.Getpid())


	eg.Go(func() error{
		for{
			select {
			case <-ctx.Done():
				fmt.Println("background ctx done")
				return ctx.Err()
			default:
				fmt.Println("摸鱼ing")
				time.Sleep(1 * time.Second)
			}
		}
	})

	err := eg.Wait()
	if err != nil {
		log.Println(err)
	}
	// time.Sleep(time.Second * 10)
}