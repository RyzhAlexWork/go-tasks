# IsomorphicStrings
Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

# Решение
По мере итерирования по строке s заполняю хэш-таблицу символами. Если ключ с данным символом отсутствует, то проверяю, есть ли в словаре ключ с данным значением. Если есть, то выдать false, так как получается случай, когда данное значение присвоено другому ключу, а значит строчки не изоморфны.
Если такого значения нет, то добавляю ключ с этим значением в словарь. Если данный ключ присутствует в словаре, то проверяю, находится ли в нём такое же значение, как данное. Если нет, то выдать false. В остальных случаях строчки изоморфны.
