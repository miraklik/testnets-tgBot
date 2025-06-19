# 🧪 Testnets Telegram Bot

Telegram bot for automating interaction with Ethereum-compatible testnets. It supports getting information about testnets, checking balance, distributing tokens and other functions.Telegram bot that allows you to store crypto testnets, and notify about testnet completion date.

## 🛠️ Technology
- Golang 1.24
- [Gin](https://gin-gonic.com/) - фреймворт для написание REST API
- MongoDB
- Docker

## 📦 Installation
```bash
git clone https://github.com/miraklik/testnets-tgBot.git
cd testnets-tgBot
docker-compose up --build -d
```

## ⚙️ Customization
```bash 
BOT_TOKEN=your_telegram_bot_token
MONGO_URI=your_mongodb_uri
JWT_SECRET=your_jwt_secret
```

## 🧑‍💻 Contributing
PR and improvements are welcome! Open an issue or fork a repository.