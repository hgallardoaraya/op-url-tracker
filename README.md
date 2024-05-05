# OP URL TRACKER

Telegram bot that sends a message if a specified className is found in the HTML retrieved by a provided URL.

To use the bot, you must configure some variables in a config file that is going to be created when the executable is running  
Also you can create the config file by yourself, it has to be named "config" and placed in the same directory as your executable  
If you want to use another file name, there is a variable named configPath at the start of the main function that you can edit.

- **URL**: URL to be scanned
- **CLASSNAME**: Classname to be searched
- **TOKEN**: Telegram bot token
- **CHAT_ID**: Telegram user chat id, is possible to get it with this bot: https://telegram.me/get_id_bot
- **MESSAGE**: Message that will be sent if the className is found in the URL's HTML  
