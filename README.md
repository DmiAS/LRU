# Задание: LRU

Нужно реализовать структуру данных, которая  удовлетворяет концепции LRU. Т.е. это должен быть такой контейнер типа ключ-значение фиксированного размера. Если добавить в заполненный кэш элемент, должен вытесняться элемент к которому обращались(добавляли или получали) давнее всего. Например
```Go
c := NewCache(3) // {}
c.Put(1, “str1”) // {1: “str1”}
c.Put(2, “str2”) // {1: “str1”, 2: “str2”}
c.Put(3, “str3”) // {1: “str1”, 2: “str2”, 3: “str3”}
c.Get(3)
c.Get(2)
c.Get(1)
c.Get(3)
c.Put(4, “str4”) // {1: “str1”, 3: “str2”, 4: “str4”}
```

Ограничения: 
размер кэша: 0 < size < 30000
ключ uint32
значение string

Операции: **NewCache, Put, Get**.

Постарайтесь реализовать методы с минимальной вычислительной сложностью. Оптимизации по памяти тоже приветствуются. 

Задача*: Добавить возможность задавать размер кэша в байтах потребляемой кэшом памяти, при условии, что все строки имеют длину 128.

Задача**: Добавить к записям TTL(фиксированный для всех записей, например 30с от момента создания, ttl задается при создании кэша), (get на записи с истекшим ttl должны возвращать not found)

# Описание реализации

Кэш создается функцией `func NewCache(cap ICap, ttl float64) *Cache`, которая принимает тип, реализующий интерфейс *ICap* с функцией *GetSize() uint16*, чтобы была возможность задать размер кэша через количество элементов и через потребляемый размер, вторым параметром указывается *ttl* в секундах

Для задания размер кэша количеством элементов используется тип *ElemCap*:
```Go
cacheCap := capacity.NewElemCap(3) // кэш размером на 3 элемента
c := cache.NewCache(cacheCap, 1) // ttl = 1 секунде
```

Для задания размер кэша размером потребляемой памяти используется тип *BytesCap*:
```Go
cacheCap := capacity.BytesCap(132) // размер строки 128 байт + размер ключа 4 байта - т.е. вместимость 1 элемент
c := cache.NewCache(cacheCap, 1) // ttl = 1 секунде
```

Для запуска: `make all` или `make`
