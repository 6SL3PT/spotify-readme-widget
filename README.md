# 🎧 Spotify Readme Widget

<div align='center'>

![Spotify](https://spotify-readme-widget-6sl3pts-projects.vercel.app/api/widget)

</div>

Yet another Spotify "Now Playing" widget for your Readme.

This project inspired by [novatorem](https://github.com/novatorem/novatorem). I've been a long-time user of his widget, and I really loved what it offered. Eventually, I wanted something that I could completely control and style my own way. So, I built this widget using **Go**, **Echo**, and **Templ**.

## Getting Started

1. Please follow novatorem [Setup](https://github.com/novatorem/novatorem/blob/main/SetUp.md) until you got the refresh token
2. Register on [Vercel](https://vercel.com/) (if not have account)
3. Fork this repo
4. Add the forked repo as new project on Vercel
5. Add environment variables:
  - `SPOTIFY_CLIENT_ID`
  - `SPOTIFY_CLIENT_SECRET`
  - `SPOTIFY_REFRESH_TOKEN`
6. Deploy

> NOTE: [Protection Bypass for Automation](https://vercel.com/docs/deployment-protection/methods-to-bypass-deployment-protection/protection-bypass-automation) need to be disabled else the request will return 401

## API Usage

Widget available at: `GET <your-vercel-project-endpoint>/api/widget`

On README.md you can add them using `![Spotify](<widget-url>)`

## Customization

You can customize the template in `templates/index.templ`, as for how to use Templ please refer to the [official document](https://templ.guide/)

