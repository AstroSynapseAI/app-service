# ASAI App Service 

Core backend service providing application and AI engine API. 

App service supports rest API for authentication, AI engine prompting, and managing the application data and settings. It also supports web socket connection for streaming AI engine results.

The AI engine is currently powered by openai models, with plans to implement Ollama open source models support soon.

## Development setup

Clone the repository and install dependencies:

`git clone https://github.com/astrosynapseai/app-service`

For lsp and intelisense supported development you do need to install go and all the packages locally but it is not required to run the docker compose.

```bash
cd app-service
go get
go mod tidy
```

For local development, the repository comes with a docker-compose.yaml and a Docker file for running the app service with hot reload support [using go air package](https://github.com/cosmtrek/air).

To setup the environment make sure you have docker and docker-compose installed on your local machine. To intialize docker container with hot reloading the first time you run a composer run:

```bash
docker compose run --rm app air init`
```

After the first run or if you don't want hot reloading, use:

```bash
docker compose up
```
### Web Client

Vue based web client is located in `./web` folder. To install dependencies run:

```bash
cd web
yarn install
```

Web client has CORS allowed to run alongised with the dockerised server. Run `yarn run dev` to run the web clinet on port 5173 as usaul. If you need the front end to be serverd at the same port as the dockerised app service, you can run `yarn run docker` to build the front ned for runing in docker, now the frontend is avaliable on port 8082.

### API keys

To authenticate for openai APIs and APIs the agents use, you need to rename `keys.template.yaml` to `keys.yaml` and input the required API keys.

```yaml
open_api_key:    ""
serpapi_api_key: ""
discord_api_key: ""
```

- For openai API key visit:  https://platform.openai.com/
- For serpapi API key visit: https://serpapi.com/
- For discord API key visit: https://discordgsm.com/guide/how-to-get-a-discord-bot-token

To run with docker the web client needs to be built "for production" with `npm run build`, the build will be saved in `./web/static` folder.

