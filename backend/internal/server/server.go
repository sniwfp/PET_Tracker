package server

import (
	"context"
	"log"
	"tracker/pb" // импортируем наш сгенерированный код
)

// Твоя структура сервера, которая будет обрабатывать запросы
type TrackerServer struct {
	// Это обязательная штука для обратной совместимости в gRPC
	pb.UnimplementedTrackerServiceServer 
}

// Создаем конструктор для нашего сервера
func NewTrackerServer() *TrackerServer {
	return &TrackerServer{}
}

// Реализуем метод SyncHealthData, который мы прописали в .proto файле
func (s *TrackerServer) SyncHealthData(ctx context.Context, req *pb.SyncHealthDataRequest) (*pb.SyncHealthDataResponse, error) {
	// Вся магия gRPC уже произошла: в переменную req прилетели готовые данные из мобилки
	log.Printf("=== Получены новые метрики ===")
	log.Printf("ID Устройства: %s", req.GetDeviceId())
	log.Printf("Дата: %s", req.GetDate())
	
	if req.GetSleep() != nil {
		log.Printf("🌙 Сон: %.1f часов", req.GetSleep().GetTotalHours())
	}
	
	if req.GetActivity() != nil {
		log.Printf("🏃 Шаги: %d, Калории: %d", req.GetActivity().GetSteps(), req.GetActivity().GetActiveCalories())
	}

	// Возвращаем успешный ответ клиенту
	return &pb.SyncHealthDataResponse{
		Success:      true,
		ErrorMessage: "",
	}, nil
}