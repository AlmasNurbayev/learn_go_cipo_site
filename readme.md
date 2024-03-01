Переносим функционал cipo_site_server на golang, а именно REST и парсер/загрузку 1С XML 

Что будем использовать:
- chi
- slog
- postrges через pgx / sqlx
- goose для миграций, причем необходимо ставить бинарник (для создания миграций) и пакет (для запуска)
  создание - goose -dir internal/storage/migrations create NAME sql
  применение миграций - go run cmd/migrate/main.go


Конфиги:
- параметры БД в env, все остальное в yaml

Пока без докера и отдельного контейнера с БД, обращаемся к хост-системе