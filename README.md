## POSTGRES

### 1) CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; - Для создания расширения для автоматического создания UUID v4 при создании таски

### 2) sudo docker exec -it go_app psql -h postgres -U feitan -d post

### 3) Создание таблицы в postgres

```
CREATE TABLE IF NOT EXISTS post (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  title TEXT NOT NULL,
  body TEXT,
  completed BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

