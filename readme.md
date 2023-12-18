### Задача 1: Многопоточный Банковский Аккаунт с `sync.Mutex`

#### Цель
Разработайте многопоточно безопасный банковский аккаунт, где горутины выполняют операции внесения и снятия денег. Используйте `sync.Mutex` для обеспечения безопасности данных при конкурентном доступе.

#### Описание Задачи
- Создайте структуру `BankAccount` с полями `balance` (типа `int`) для баланса и `mutex` (типа `sync.Mutex`) для управления доступом к балансу.
- Реализуйте методы `Deposit(amount int)` и `Withdraw(amount int)`, которые будут использовать `mutex` для безопасного доступа к `balance`.
- В `Deposit`, добавьте указанную сумму к `balance`.
- В `Withdraw`, снимите указанную сумму с `balance`, если это возможно.
- В функции `main`, инициируйте `BankAccount` и запустите несколько горутин, которые будут выполнять операции внесения и снятия денег.
- После выполнения всех операций, выведите итоговый баланс для проверки корректности работы.

### Задача 2: Продвинутое Использование `sync.WaitGroup` и `sync.Mutex`

#### Цель
Разработать многопоточно безопасную систему логирования, используя `sync.WaitGroup` для ожидания завершения записи логов и `sync.Mutex` для синхронизации доступа к общему буферу логов.

#### Описание Задачи
- Создайте структуру `LogBuffer`, содержащую `buffer` (слайс строк для хранения логов) и `mutex` (для управления доступом к буферу).
- Реализуйте метод `WriteLog(message string)` в `LogBuffer`, который безопасно добавляет сообщение лога в `buffer`, используя `mutex` для предотвращения состояний гонки.
- В функции `main`, инициируйте `LogBuffer` и запустите несколько горутин, каждая из которых записывает определенное количество сообщений лога.
- Используйте `sync.WaitGroup` для ожидания завершения всех горутин, записывающих логи.
- После завершения всех горутин, выведите содержимое буфера логов.

### Задача 3: Многопоточный Счетчик Посещений с `sync.Map` и `sync/atomic`

#### Цель
Разработать многопоточно безопасную систему для отслеживания посещений веб-сайтов, используя `sync.Map` для хранения счетчиков и `sync/atomic` для атомарного обновления этих счетчиков.

#### Описание Задачи
- Создайте структуру `WebCounter`, содержащую `visits` (типа `sync.Map`), где ключ — это URL веб-сайта, а значение — атомарный счетчик посещений.
- Реализуйте метод `Increment(url string)`, который атомарно увеличивает счетчик посещений для данного URL.
- Реализуйте метод `GetVisits(url string) int`, возвращающий текущее количество посещений для указанного URL.
- В функции `main`, запустите несколько горутин, имитирующих посещения разных сайтов. Каждая горутина должна вызывать метод `Increment` для случайного URL.
- После завершения всех горутин выведите итоговое количество посещений для каждого сайта.


### Задача 4: Сложная Многопоточность с `Channels` и `select`

#### Цель
Реализуйте систему обработки запросов с использованием каналов и `select`, где запросы распределяются по разным обработчикам.

#### Описание:
- Используйте каналы для передачи запросов и результатов обработки.
- Создайте несколько горутин-обработчиков, читающих из канала запросов.
- Используйте `select` для управления каналами запросов, результатов и таймаутов.

---

### Задача 5: Супер Сложная Многопоточная Система

#### Цель
Разработайте распределенную систему обработки данных с использованием всех перечисленных примитивов синхронизации.

#### Описание:
- Структурируйте систему так, чтобы она включала горутины для сбора данных, их обработки, агрегирования результатов и логирования.
- Используйте `sync.Mutex` или `RWMutex` для защиты общих ресурсов.
- Примените `sync.Map` для хранения промежуточных данных.
- Реализуйте `sync/atomic` для счетчиков и флагов состояния.
- Используйте каналы и `select` для управления потоками данных и синхронизации.
- Обеспечьте корректное завершение всех горутин и обработку ошибок.
