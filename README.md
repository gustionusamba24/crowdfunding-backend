
# Crowdfunding

## Description:

Crowdfunding is a way of raising money to finance projects or business. This project was inspired by [KitaBisa](https://kitabisa.com/). Feel free to try this API.

## Tech Stack:

* Go
* Gin
* GORM
* Validator
* MySQL
* JWT-Authorization
* 3rd Party APIs (Midtrans)
* REST

## ERD

![ERD Crowdfunding](./misc/erd_crowdfunding.svg)

## List Route

## API Route
### Users
- **`POST` - `api/v1/users`**, Create new user
- **`POST` - `api/v1/sessions`**, Login user
- **`POST` - `api/v1/email_checkers`**, Check user's email wheter it is registered or not
- **`POST` - `api/v1/avatars`**, Upload user's avatar

### Campaigns
- **`GET` - `api/v1/campaigns`**, Get all campaigns
- **`GET` - `api/v1/campaigns/:id`**, Get detail campaign by id
- **`POST` - `api/v1/campaigns`**, Create campaign
- **`PUT` - `api/v1/campaigns/:id`**, Update campaign
- **`POST` - `api/v1/campaigns-images`**, Update campaign's image

### Transactions
- **`GET` - `api/v1/campaigns/:id/transactions`**, Get the transaction data from the campaign based on its id
- **`GET` - `api/v1/transactions`**, Get the transaction data from users who make transactions
- **`POST` - `api/v1/transactions`**, Create transaction




