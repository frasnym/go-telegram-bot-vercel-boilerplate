# Golang Telegram Bot Boilerplate in Vercel

A simple boilerplate to kickstart your Telegram bot development using Golang and deploy it on Vercel.

## Features

- **Golang**: Built using the Go programming language.
- **Telegram Bot API**: Utilizes the Telegram Bot API for interacting with Telegram.
- **Vercel Deployment**: Easily deploy your bot on Vercel with serverless functions.

## Getting Started

### Prerequisites

- [Go](https://golang.org/) installed on your machine.
- [Vercel CLI](https://vercel.com/download) installed for deploying to Vercel.

### Clone the Repository

```bash
git clone https://github.com/your-username/your-repo.git
cd your-repo
```

### Set Up Telegram Bot

1. Create a new Telegram bot on Telegram and obtain the token.
2. Copy the token to the `.env` file.

```env
TELEGRAM_BOT_TOKEN=your-telegram-bot-token
```

### Local Development

```bash
go run main.go
```

Your bot should now be running locally.

### Deploy to Vercel

1. Log in to Vercel using the `vercel login` command.
2. Deploy the bot to Vercel.

```bash
vercel
```

Follow the prompts, and your bot will be deployed on Vercel.

### Deploy to Vercel One Click
[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2Ffrasnym%2Fgo-telegram-bot-vercel-boilerplate&env=TELEGRAM_BOT_TOKEN,PORT&envDescription=TELEGRAM_BOT_TOKEN%20is%20needed%2C%20you%20can%20ask%20BotFather%20for%20it&envLink=https%3A%2F%2Ft.me%2FBotFather&project-name=go-telegram-bot-vercel-boilerplate&repository-name=go-telegram-bot-vercel-boilerplate)

### Disable Vercel Authentication
From your Vercel dashboard:
1. Select the project that you wish to enable Password Protection for
2. Go to Settings then Deployment Protection
3. Learn more about [Vercel Authentication](https://vercel.com/docs/security/deployment-protection/methods-to-protect-deployments/vercel-authentication)

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
