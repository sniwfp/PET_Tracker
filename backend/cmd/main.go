package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"tracker/internal/server"
	"tracker/pb"
	"tracker/internal/config"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 1. Указываем порт, который сервер будет слушать
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.ServerPort))
	if err != nil {
		log.Fatalf("Не удалось запустить листенер на порту %d: %v", cfg.ServerPort, err)
	}

	// 2. Создаем экземпляр базового gRPC сервера от Google
	grpcServer := grpc.NewServer()

	// 3. Создаем наш кастомный сервер с логикой
	myTrackerServer := server.NewTrackerServer()

	// 4. Связываем их вместе (регистрируем наш сервис)
	pb.RegisterTrackerServiceServer(grpcServer, myTrackerServer)

	stop := make(chan os.Signal, 1)
	// Направляем сигналы SIGINT (Ctrl+C) и SIGTERM (остановка процесса) в наш канал
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// 2. Запускаем gRPC сервер в отдельной горутине, чтобы он не блокировал главный поток
	go func() {
		log.Println("🚀 gRPC сервер запущен на порту :50051...")
		if err := grpcServer.Serve(listener); err != nil {
			log.Printf("Сервер завершил работу: %v", err)
		}
	}()

	// 3. Код "засыпает" здесь, пока из ОС не прилетит сигнал остановки
	<-stop
	log.Println("⚠️ Получен сигнал остановки. Начинаем Graceful Shutdown...")

	// 4. Даем серверу время завершить текущие запросы, но жестко ограничиваем тайм-аут
	shutdownFinished := make(chan struct{})
	go func() {
		grpcServer.GracefulStop() // Плавная остановка gRPC
		close(shutdownFinished)
	}()

	select {
	case <-shutdownFinished:
		log.Println("✅ Все запросы обработаны. Сервер успешно остановлен.")
	case <-time.After(5 * time.Second): // Если за 5 секунд сервер не дорендерил задачи — тушим принудительно
		log.Println("⏱ Превышен лимит ожидания (5с). Принудительное завершение.")
		grpcServer.Stop()
	}
}