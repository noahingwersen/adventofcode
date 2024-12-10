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
printf "import (\n\t\"bufio\"\n\t\"fmt\"\n\t\"os\"\n\t\"time\"\n)\n\n" >> $gofile
printf "func check(e error) {\n\tif e != nil {\n\t\tpanic(e)\n\t}\n}\n\n" >> $gofile
printf "const filepath = \"test.txt\"\n\n" >> $gofile
printf "func parseInput() []string {\n\tfile, ferr := os.Open(filepath)\n\tcheck(ferr)\n\n\t" >> $gofile
printf "scanner := bufio.NewScanner(file)\n\tvar lines []string\n\t" >> $gofile
printf "for scanner.Scan() {\n\t\tlines = append(lines, scanner.Text())\n\t}\n\tfile.Close()\n\treturn lines\n}\n\n" >> $gofile
printf "func part1() int {\n\treturn 0\n}\n\n" >> $gofile
printf "func part2() int {\n\treturn 0\n}\n\n" >> $gofile
printf "func main() {\n\tstart := time.Now()\n\tfmt.Printf(\"Part 1: %%v in %%v\\\\n\", part1(), time.Since(start))\n\thalf := time.Now()\n\tfmt.Printf(\"Part 2: %%v in %%v\\\\n\", part2(), time.Since(half))\n}" >> $gofile