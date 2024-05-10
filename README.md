# OP URL TRACKER

Telegram bot that sends a message if a specified className is found in the HTML retrieved by a provided URL, every request is done in its own goroutine.


To use the bot, you must configure some variables in a config file that is going to be created when the executable is running  
Also you can create the config file by yourself, it has to be named "config" and placed in the same directory as your executable  
If you want to use another file name, there is a variable named configPath at the start of the main function that you can edit.

- **URLS**: URLs to be scanned, must be separated by a comma and without spaces (url1,url2,url3)
- **CLASSNAMES**: Classname to be searched in every URL, must be separated by a comma and without spaces (class1,class2,class3)
- **TOKEN**: Telegram bot token
- **CHAT_ID**: Telegram user chat id, is possible to get it with this bot: https://telegram.me/get_id_bot
- **MESSAGE**: Message that will be sent if the className is found in the URL's HTML  

**URLs** and **CLASSNAMES** have to be same length, because class_i is going to be searched in url_i.  
Example:  
**URLS**=url1,url2,url3  
**CLASSNAMES**=class1,class2,class3  
class1 is going to be searched in url1, class2 in url2, and so on.