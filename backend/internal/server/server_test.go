package server

import (
	"context"
	"testing"
	"tracker/pb"
)

func TestSyncHealthData(t *testing.T) {
	// 1. Инициализируем наш сервер
	srv := NewTrackerServer()

	// 2. Описываем тест-кейсы (табличный подход)
	tests := []struct {
		name    string
		request *pb.SyncHealthDataRequest
		wantSuccess bool
	}{
		{
			name: "Успешный кейс со всеми данными",
			request: &pb.SyncHealthDataRequest{
				DeviceId: "iphone_15_pro",
				Date:     "2026-06-24",
				Sleep: &pb.SleepData{
					TotalHours: 7.5,
				},
				Activity: &pb.ActivityData{
					Steps:          10500,
					ActiveCalories: 450,
					WorkoutDone:    true,
				},
			},
			wantSuccess: true,
		},
		{
			name: "Успешный кейс без данных об активности (только сон)",
			request: &pb.SyncHealthDataRequest{
				DeviceId: "samsung_s24",
				Date:     "2026-06-24",
				Sleep: &pb.SleepData{
					TotalHours: 6.0,
				},
				Activity: nil, // Проверяем, что сервер не упадет в панику, если активности нет
			},
			wantSuccess: true,
		},
		{
			name: "Успешный кейс с абсолютно пустым запросом",
			request: &pb.SyncHealthDataRequest{},
			wantSuccess: true,
		},
	}

	// 3. Запускаем перебор тест-кейсов
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Вызываем тестируемый метод напрямую
			resp, err := srv.SyncHealthData(context.Background(), tt.request)

			// Проверяем, что метод не вернул системную ошибку Go (err)
			if err != nil {
				t.Fatalf("Метод вернул непредвиденную ошибку: %v", err)
			}

			// Проверяем бизнес-логику: флаг успеха в ответе Protobuf
			if resp.GetSuccess() != tt.wantSuccess {
				t.Errorf("SyncHealthData() success = %v, want %v", resp.GetSuccess(), tt.wantSuccess)
			}
		},
	)
}
}