# web-control

Панель управления  для https://github.com/xaosBotTeam/go-bot написанная на GO и vue.js 

## Usage

Смортеть https://github.com/xaosBotTeam/infrastructure

## Parameters

| Название параметра | Описание                       | Обязательный          | Значеине по умолчанию |
|--------------------|--------------------------------|-----------------------|-----------------------|
| GO_BOT_URL         | адрес вышестоящего API xaosBot | <center>  ❌ </center> | 127.0.0.1:5504        |
| POSTGRES_USER      |                                | <center>  ✅ </center> | -                     |
| POSTGRES_PASSWORD  |                                | <center>  ✅ </center> | -                     |
| POSTGRES_DB        |                                | <center>  ✅ </center> | -                     |
| POSTGRES_HOST      |                                | <center>  ✅ </center> | -                     |
| POSTGRES_PORT      |                                | <center>  ✅ </center> | -                     |

## API

С полным описанием методов можно ознакомиться [здесь](https://editor.swagger.io/?url=https://raw.githubusercontent.com/xaosBotTeam/web-control/dev/docs/swagger.json) 
Или после развертывания приложения перейти на /swagger/index.html