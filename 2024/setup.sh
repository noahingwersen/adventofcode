#!/bin/bash
echo Enter a day:
read day

# Make sure input is a number
if ! [[ "$day" =~ ^[0-9]+$ ]] ; 
 then exec >&2; echo "error: day must be a number"; exit 1
fi

directory=Day$day
gofile=day$day.go

if [ -d "$directory" ]; then
    rm -r $directory
fi

mkdir $directory
cd $directory
touch test.txt
touch input.txt
touch $gofile

# File template
printf "package main\n\n" >> $gofile
printf "import (\n\t\"bufio\"\n\t\"fmt\"\n\t\"os\"\n)\n\n" >> $gofile
printf "func check(e error) {\n\tif e != nil {\n\t\tpanic(e)\n\t}\n}\n\n" >> $gofile
printf "const filepath = \"test.txt\"\n\n" >> $gofile
printf "func parseInput() {\n\tfile, ferr := os.Open(filepath)\n\tcheck(ferr)\n\n\tscanner := bufio.NewScanner(file)\n\t" >> $gofile
printf "for scanner.Scan() {\n\t\tline := scanner.Text()\n\n\t}\n\tfile.Close()\n}\n\n" >> $gofile
printf "func part1() int {\n\treturn 0\n}\n\n" >> $gofile
printf "func part2() int {\n\treturn 0\n}\n\n" >> $gofile
printf "func main() {\n\tfmt.Println(part1())\n\tfmt.Println(part2())\n}\n" >> $gofile