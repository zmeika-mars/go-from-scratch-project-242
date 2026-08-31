### Hexlet tests and linter status:
[![Actions Status](https://github.com/zmeika-mars/go-from-scratch-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/zmeika-mars/go-from-scratch-project-242/actions)

## Demo "Анализатор размера диска"
[Demo Disk Size Analyzer](https://asciinema.org/a/wpeiBEzn0MGFvFAj)

## Установка

Склонировать репозиторий, собрать программу:

```Shell
go build -o bin/hexlet-path-size ./cmd/hexlet-path-size
```

После сборки, программу можно запустить: 

```Shell
./bin/hexlet-path-size <path>
```

## Пример использования программы 

Получение размера файла:

```Shell
./bin/hexlet-path-size testdata/.env
```

Получения размера директории:

```Shell
./bin/hexlet-path-size testdata/
```

## Флаги

`-h`, `--help`

Получить справку.

#### Использование:

```Shell
./bin/hexlet-path-size --help
```

`-r`, `--recursive`

Рекурсивно вычисляет размер вложенных директорий.

#### Использование:

```Shell
./bin/hexlet-path-size -r testdata/
```

`-a`, `--all`

Учитывать скрытые файлы и директории. Без этого флага скрытые файлы пропускаются.

#### Использование:

```Shell
./bin/hexlet-path-size -a testdata/
```

`-h`, `--human`

Выводит размер в удобном читаемом формате.

```Shell
./bin/hexlet-path-size -h testdata/
```
