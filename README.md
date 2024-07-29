.env оставлен намеренно

сколонировать репрозиторий и использовать
```
git clone https://github.com/SKOLIA0/message-processor.git
cd message-processor
docker-compose up --build
```
проверить что работают все 3 контейнера

app
    
kafka
    
zookeeper
```
sudo docker ps
```
иначе повторять пока не запустятся все три
```
sudo docker-compose down
sudo docker-compose up -d
```
проверка
```
sudo docker ps
```

Откройте Postman и нажмите на кнопку "New" -> "HTTP Request".

Выберите метод POST.

Введите URL сервиса 
```
http://82.97.241.37/messages
```
Перейдите на вкладку "Body".

Выберите формат raw.

В выпадающем списке выберите JSON.

Введите тело вашего запроса в формате JSON.
```
{
    "message": "Hello, Messaggio GO!"
}
```
Выберите метод GET.

Введите URL сервиса 
```
http://82.97.241.37/stats
```
