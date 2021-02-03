<h1 align="center">
📊 Wakapi-Stats
</h1>

<h2 align="center">
Generator images from Wakapi
</h2>

<p align="center">

  <img src="https://badges.fw-web.space/github/v/release/lacazethomas/wakapi-stats">

  <a href="https://goreportcard.com/report/github.com/LacazeThomas/wakapi-stats">
    <img src="https://goreportcard.com/badge/github.com/LacazeThomas/wakapi-stats">
  </a>
  
  <img src="https://badges.fw-web.space/github/languages/code-size/lacazethomas/wakapi-stats">

  <img src="https://badges.fw-web.space/github/license/LacazeThomas/wakapi-stats">
  
  <a href="https://drone.thomaslacaze.fr/LacazeThomas/wakapi-stats">
    <img src="https://drone.thomaslacaze.fr/api/badges/LacazeThomas/wakapi-stats/status.svg">
  </a>
</p>

<p align="center">
    <a href="https://github.com/LacazeThomas/wakapi-stats/issues/new/choose">Report Bug</a>
    ·
    <a href="https://github.com/LacazeThomas/wakapi-stats/issues/new/choose">Request Feature</a>
</p>


<table>
  <tr>
    <th colspan="2" align="center">
      📊 Metrics past 30 days
    </th>
  </tr>
  <tr>
    <th>💬 Most used programming languages</th>
    <th>🛠️ Most used editors</th>
  </tr>
  <tr>
    <td align="center">
      <img alt="" width="380" src="https://raw.githubusercontent.com/LacazeThomas/LacazeThoma/main/languages.png">
      <img width="800" height="1" alt="">
    </td>
    <td align="center">
      <img alt="" width="380" src="https://raw.githubusercontent.com/LacazeThomas/LacazeThoma/main/editors.png">
      <img width="800" height="1" alt="">
    </td>
  </tr>
</table>

## 🧐 Features

- **✅ 100 % free and open-source**

- **✅ Work with drone.yml**

- **✅ Generate most used editors**

- **✅ Generate most used programming languages**

- **✅ Generate most used operating systems**

- **✅ Self-hosted**


## 🛠️ Installation Steps

### 🐳 Option 1: Run with Drone.io
**See [here](https://github.com/LacazeThomas/LacazeThomas/blob/main/.drone.yml)** 

### 🐳 Option 2: Run from Docker run 
```bash
# Run the container
$ docker run -d \
  -p 3000:3000 \
  -e "PLUGIN_ENVIRONMENT=prod" \
  -e "PLUGIN_API_URL=FILL_ME" \
  -e "PLUGIN_API_KEY=FILL_ME" \
  -e "PLUGIN_ACCESSTOKEN=FILL_ME" \
  -e "PLUGIN_BRANCH=FILL_ME" \
  -e "PLUGIN_MESSAGE=FILL_ME" \
  -e "PLUGIN_COMMITNAME=FILL_ME" \
  -e "PLUGIN_COMMITEMAIL=FILL_ME" \
  -e "PLUGIN_USERNAME=FILL_ME" \
  -e "PLUGIN_REPO=FILL_ME" \
  --name wakapi-stats thomaslacaze/wakapi-stats:1.0.0
```

### 🐳 Option 3: Run from Docker-compose
**See [here](https://github.com/LacazeThomas/wakapi-stats/blob/main/docker-compose.yml)** 

### 💻 Option 4: Run from source
#### Prerequisites
* Go >= 1.8 (with `$GOPATH` properly set)

1. Clone the repository

```bash
git clone https://github.com/LacazeThomas/wakapi-stats.git
```

2. Change the working directory

```bash
cd wakapi-stats
```

3. Setup environnement variables

| Environment Variable      | Default      | Description                                                         |
|---------------------------|--------------|---------------------------------------------------------------------|
| `PLUGIN_ENVIRONMENT`               | `dev`          | Whether to use development or production settings                  |
| `PLUGIN_API_URL`               | None          | Wakapi URL, exemple : https://wakapi.thomaslacaze.fr                  |
| `PLUGIN_API_KEY`               | None          | Wakapi API Key, same as your using in your editors                  |
| `PLUGIN_ACCESSTOKEN`               | None          | Github Token, generate [here](https://github.com/settings/tokens)                  |
| `PLUGIN_BRANCH`               | `main`          | Github Branch where commit will be publish                  |
| `PLUGIN_MESSAGE`               | `Update images from wakapi-stats`          | Displayed message on the commit                  |
| `PLUGIN_COMMITNAME`               | None          | Name for commit                   |
| `PLUGIN_COMMITEMAIL`               | None          | Email for commit                  |
| `PLUGIN_USERNAME`               | None          | Github username                  |
| `PLUGIN_REPO`               | None          | Github repo where commit will be publish                  |


4. Run the app

```bash
go build -o wakapi-stats && ./wakapi-stats
```

🌟 You are all set!


## 🙇 Special Thanks

- [Ferdinand Mütsch](https://github.com/muety) for amazing [wakapi](https://github.com/muety/wakapi)
- [Will Charczuk](https://github.com/wcharczuk) for the awesome [go-chart](https://github.com/wcharczuk/go-chart)

## License

<a href="https://github.com/LacazeThomas/wakapi-stats/blob/main/LICENSE">MIT</a>

<hr>
<p align="center">
Developed with ❤️ in France 
</p>

