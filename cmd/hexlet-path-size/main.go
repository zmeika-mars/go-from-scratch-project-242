package main

import ( 
    "github.com/urfave/cli/v3"
    "os"
    "context"
    "log"
	"fmt"
	"strings"
	"path/filepath"
)

func main() {
   cmd := &cli.Command{
       Name: "hexlet-path-size",
       Usage: "print size of a file or directory",
       ArgsUsage: "<path>",

	   Flags: []cli.Flag{
		   &cli.BoolFlag{
			   Name: "human",
			   Aliases: []string{"H"},
			   Usage: "human-readable sizes (auto-select unit)",
		   },
		   &cli.BoolFlag{
			   Name: "all",
			   Aliases: []string{"a"},
			   Usage: "include hidden files and directories",
		   },
		   &cli.BoolFlag{
			   Name: "recursive",
			   Aliases: []string{"r"},
			   Usage: "recursive size of directories",
		   },
	   },

       Action: func(ctx context.Context, cmd *cli.Command) error {
		   path := cmd.Args().Get(0)
         
		   // Получаем значение флага
		   all := cmd.Bool("all")
		   human := cmd.Bool("human")
		   recursive := cmd.Bool("recursive")

		   //Получаем размеры файла или директории
		   rawSize, err := GetPathSize(path, all, recursive)
		   if err != nil {
			   return err
		   }

		   size := FormatSize(rawSize, human)

		   fmt.Printf("%s\t%s\n", size, path)

		   return nil
	   },  
   }

   if err := cmd.Run(context.Background(), os.Args); err != nil {
       log.Fatal(err)
   }
}

func GetPathSize(path string, all bool, recursive bool) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

    // Переменная для хранения общей суммы размеров 
	var rawSize int64
	
	// Если это файл, проверяем скрытый он или нет
	if !info.IsDir() {
		if strings.HasPrefix(info.Name(), ".") && !all {
			return 0, nil
		}

		return info.Size(), nil
	}	
	// Если директория получить ее содержимое
	files, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, file := range files {
		//Если файл скрытый и флаг --all не указан, пропускаем  
		if strings.HasPrefix(file.Name(), ".") && !all {
			continue
		}

		//Если текущий элемент является директорией
		if file.IsDir() {
			// Если передан флаг '-r' проверяем файлы внутри директорий 
			if recursive {
				size, err := GetPathSize(
					filepath.Join(path, file.Name()),
					all,
					recursive,
				)
				if err != nil {
					return 0, err
				}
				//Добавляем размер директорий к общей сумму 
				rawSize += size
			}
			//После обработки директорий, переходим к следующему элементу
			continue
		}


		info, err := file.Info()
		if err != nil {
			return 0, err
		}

		// Добаввялем размер файла к общей сумме
		rawSize += info.Size()
	}

	return rawSize, nil

}


func FormatSize(size int64, human bool) string {
	// Если не передан флаг human, возращать размер в байтах
	if !human {
		return fmt.Sprintf("%dB", size)
	}
	// Список единиц измерения
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

	value := float64(size)
	// Индекс текущей единицы измерения
	unit := 0

	 // Проверяем, что не достигли последней единицы измерения
	for value >= 1024 && unit < len(units)-1 {
		// Делим на 1024 и переходим к следующей единице измерения, пока значение еще больше или равно 1024 
		value = value / 1024
		unit++
	}
	
	// Если размер меньше 1024 байт, вывести размер в байтах
	if unit == 0 {
		return fmt.Sprintf("%dB", size)
	}

	return fmt.Sprintf("%.1f%s", value, units[unit])
}
