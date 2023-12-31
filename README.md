## Утилита для копирования файлов
Утилита копирования файлов (упрощенный аналог `dd`).

Принимает следующие аргументы:
* путь к исходному файлу (`-from`);
* путь к копии (`-to`);
* отступ в источнике (`-offset`), по умолчанию - 0;
* количество копируемых байт (`-limit`), по умолчанию - 0 (весь файл из `-from`).

Особенности:
* offset больше, чем размер файла - невалидная ситуация;
* limit больше, чем размер файла - валидная ситуация, копируется исходный файл до его EOF;
* программа может НЕ обрабатывать файлы, у которых неизвестна длина (например, /dev/urandom);
