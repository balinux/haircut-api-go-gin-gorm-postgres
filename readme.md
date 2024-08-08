# deskripsi

api yang dibuat untuk hair style denga go gin gorm dan database mysql

## Struktur folder project

```language
├── config
│   └── database.go
├── controllers
│   └── haircut_controller.go
├── models
│   └── haircut.go
├── repositories
│   └── haircut_repository.go
├── services
│   └── haircut_service.go
├── routes
│   └── routes.go
└── main.go
```

## running podman - build image

`podman build -t haircut:postgres .`

## podman - run container

`podman run -d -p 8888:8888 -it --name haircut-postgres haircut:postgres`

## podman stop container

`podman container stop <ID>`

## container delete all container

`podman rm $(podman ps -a -q)`

## test post data

```bash
curl -X POST http://localhost:8888/api/v1/haircuts \
  -H "Content-Type: application/json" \
  -d '{"name": "Short Trim", "description": "A quick short trim", "price": 15.99,"image": "https://cdnwpseller.g
ramedia.net/wp-content/uploads/2022/06/15085945/gaya_model_rambut_undercut_2021_terbaru_discoonected_undercut-56
1x625.jpg"}'

```
