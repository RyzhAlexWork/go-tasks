# ValidParentheses
Imagine you are writing a small compiler for your college project and one of the tasks (or say sub-tasks) for the compiler would be to detect if the parenthesis are in place or not.

The algorithm we will look at in this article can be then used to process all the parenthesis in the program your compiler is compiling and checking if all the parenthesis are in place. This makes checking if a given string of parenthesis is valid or not, an important programming problem.

# Реализация
Рекурсивно вызывается функция поиска необходимого закрывающего символа. Если символ найден, возвращается позиция следующего за ним символа. Если нет, возвращаю 0. Если нактнулся на новый открывающий символ, рекурсивно вызываю функцию.
