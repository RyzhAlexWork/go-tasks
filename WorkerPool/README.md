# WorkerPool
Написать программу worker pool с использованием каналов, входные параметры - количество рабочих. У нас есть рабочие, которые начинают работу параллельно (время работы - рандомное число от 1 до 3 секунд), затем они завершают работу и сдают свои результаты на проверку прорабу, который их принимает.

# Реализация
Структура work хранит в себе i - номер рабочего и result - результат его работы.
Создаётся первый канал, куда помещаются все номера рабочих. Затем создаётся второй канал, который принимает значения из первого
канала и выдаёт каждому рабочему работу. Результаты всех выполненных работ принимаются третьим каналом. Прораб проверяет
работу (считывает значения из третьего канала) и в случае нахождения ошибки - сообщает о ней.
