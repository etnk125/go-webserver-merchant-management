# 1st GO web server 

## Run
```
go mod tidy
go run main.go
```

## Details  
- [x] 1. Create a web server that listens on port 8080.
- [x] 2. Project should have structure [controller,service,repo].
- [ ] 3. Project should authenticate user.
- [ ] 4. Project should have tests.
- [ ] 5. Run as batch job.
- [ ] 6. Separate config file.
- [ ] 7. Write file

## API route in table

| Method | Route                                | Body [JSON]        | Auth  | Description                    | status |
| ------ | ------------------------------------ | ------------------ | ----- | ------------------------------ | ------ |
| GET    | `/`                                  |                    | False | Health check                   | ✅      |
| GET    | `/health`                            |                    | False | Health check                   | ✅      |
| POST   | `/merchant`                          | name, bank_account | False | Register merchant              | ✅      |
| GET    | `/merchant/:merchant_id`             | name, bank_account | True  | Get merchant info              | 🟨      |
| PUT    | `/merchant/:merchant_id`             | name, bank_account | True  | Update merchant info           | 🟨      |
| POST   | `/merchant/:merchant_id/product`     | name, price        | True  | Add product to that merchant   | ⛔      |
| GET    | `/merchant/:merchant_id/product/all` |                    | True  | Get products of that merchant  | ⛔      |
| POST   | `/product/:product_id`               | quantity           | False | Buy product from that merchant | ⛔      |
| GET    | `/sell-report/:date`                 |                    | True  | Get sell report                | ⛔      |




<!-- | `/product/:product_id`   | True  |                    | GET    | Get product                   | -->





