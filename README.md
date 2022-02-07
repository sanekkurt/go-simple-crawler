# go-simple-crawler

To use it, you need to send a **post** request to the **/crawler** method with the json content of the form:
```json
{
    "data":[
        "https://ya.ru",
        "https://github.com/",
        "https://mail.ru/"
    ]
}
```

The configuration file must be passed through the **CONFIG_PATH** environment variable or through the **-c** flag

The already prepared configuration file is located in the **example** folder

