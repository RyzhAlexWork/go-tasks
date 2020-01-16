# Set
Реализовано множество (set) для intов через создание структуры с одним значением типа map.

# Методы
Create() *Set - создаёт пустое множество, возвращает его.

Add(s *Set, value int) bool - добавляет элемент. Возвращает true при успешном добавлении, иначе false.

Print(s *Set) - выводит все элементы множества.

Has(s *Set, value int) bool - проверяет, есть ли заданный элемент в множестве. Если есть, возвращает true, иначе false.

Clear(s *Set) - очищает множество от всех элементов.

Delete(s *Set, value int) bool - удаляет заданный элемент из множества. В случае удачи возвращает true, иначе false.

Lem(s *Set) int - возвращает количество элементов множества.