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

  <a href="https://coveralls.io/github/LacazeThomas/wakapi-stats?branch=main">
    <img src="https://coveralls.io/repos/github/LacazeThomas/wakapi-stats/badge.svg?branch=main">
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
      <img alt="" width="380" src="https://wakapi-stats.thomaslacaze.fr/api/v1/languages?url=https://stats-code.thomaslacaze.fr/api/v1/users/thomaslacaze/stats/last_30_days">
      <img width="800" height="1" alt="">
    </td>
    <td align="center">
      <img alt="" width="380" src="https://wakapi-stats.thomaslacaze.fr/api/v1/editors?url=https://stats-code.thomaslacaze.fr/api/v1/users/thomaslacaze/stats/last_30_days">
      <img width="800" height="1" alt="">
    </td>
  </tr>
</table>

## 🧐 Features

- **✅ Display or not time spent**

- **✅ 100 % free and open-source**

- **✅ No need CI**

- **✅ Generate most used editors**

- **✅ Generate most used programming languages**

- **✅ Generate most used operating systems**

- **✅ Self-hosted**

## 🔧 Endpoint availables 
### (/stats/editors, /stats/languages, /stats/operatingSystems, /stats/machines, /stats/projects)

### Range possibility today, yesterday, week, month, year, 7_days, last_7_days, 30_days, last_30_days, 12_months, last_12_months, any

Please do not forget to put your wakapi endpoint into public 
```
http://localhost:8080/api/v1/editors?url=https://stats-code.thomaslacaze.fr/api/v1/users/thomaslacaze/stats/30_days
```

With real time spent

```
http://localhost:8080/api/v1/editors?timeSpent=true&url=https://stats-code.thomaslacaze.fr/api/v1/users/thomaslacaze/stats/30_days
```


## ☁️ [Public instance](https://wakapi-stats.thomaslacaze.fr)
```
https://wakapi-stats.thomaslacaze.fr/api/v1/editors?url=https://stats-code.thomaslacaze.fr/api/v1/users/thomaslacaze/stats/30_days
```

## 🛠️ Installation Steps

### 🐳 Option 1: Run from Docker run 
```bash
# Run the container
$ docker run -d \
  -p 8080:8080 \
  -e "ENVIRONMENT=prod" \
  --name wakapi-stats thomaslacaze/wakapi-stats:1.1.0
```

### 🐳 Option 2: Run from Docker-compose
**See [here](https://github.com/LacazeThomas/wakapi-stats/blob/main/docker-compose.yml)** 

### 💻 Option 3: Run from source
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
| `ENVIRONMENT`               | `dev`          | Whether to use development or production settings                  |

4. Run the app

```bash
go build -o wakapi-stats && ./wakapi-stats
```

🌟 You are all set!


## 🙇 Special Thanks

- [Ferdinand Mütsch](https://github.com/muety) for amazing [wakapi](https://github.com/muety/wakapi)
- [Will Charczuk](https://github.com/wcharczuk) for the awesome [go-chart](https://github.com/wcharczuk/go-chart)


## Dockerfile

<a href="https://github.com/LacazeThomas/wakapi-stats/blob/main/Dockerfile">Dockerfile</a>
## License

<a href="https://github.com/LacazeThomas/wakapi-stats/blob/main/LICENSE">MIT</a>

<hr>
<p align="center">
Developed with ❤️ in France 
</p>

