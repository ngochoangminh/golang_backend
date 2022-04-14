# golang_backend

## Installation
### Install golang-migration
```
curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
apt-get update
apt-get install -y migrate
```
### Install sqlc
```
sudo snap install sqlc
```
## Create database and migration
```
migration_folder
    ├── 001_initial_schema.down.sql
    ├── 001_initial_schema.up.sql
    ├── 002_user.down.sql
    ├── 002_user.up.sql
    ├── 003_session.down.sql
    └── 003_session.up.sql
```
Create table:
```
migrate -path database/migration -database "postgresql://<db_user>:<password>@localhost:5432/db_name" -verbose up
```
Drop table:
```
migrate -path database/migration -database "postgresql://<db_user>:<password>@localhost:5432/db_name" -verbose down

```