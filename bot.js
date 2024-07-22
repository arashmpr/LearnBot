const { Telegraf } = require("telegraf")

const API_TOKEN = "6854615778:AAH92L0DhGWvBx_S9T79H1SEMKRYrkkzAJo"
const bot = new Telegraf(API_TOKEN)

bot.start((ctx) => 
    ctx.reply("خوش آمدید")
)
bot.launch()