Это простая программа для бекапа файлов в MailRuCloud.


Скомпилировать: go build -o Backuper.exe


Как работает:
1. Нужно создать и заполнить файл config_mailru.json аналогично примеру _example
2. Запускаем Backuper.exe и начинается создание .zip в директории исполняемого файла, отправка в cloud и удаление созданного .zip


* Параметр FilesToBackup пока не работает.