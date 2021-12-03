./build.sh

read -p "Press enter to run ./aoc -h: "
echo
./aoc -h
echo

read -p "Press enter to run ./aoc mkday -h: "
echo
./aoc mkday -h
echo

read -p "Press enter to run ./aoc mkday 20 2020: "
echo
./aoc mkday 20 2020
echo

read -p "Press enter to run ./aoc submit -h: "
echo
./aoc submit -h
echo

read -p "Press enter to run ./aoc submit 1 input.txt 1 2015: "
echo
./aoc submit 1 input.txt 1 2015
echo

read -p "Press enter to run ./aoc stats -h: "
echo
./aoc stats -h
echo

echo "These only test programs that terminate."
echo "Remember to also test"
echo "./aoc mkday"
echo "and"
echo "./aoc stats"
echo

read -p "Press enter to end tests: "