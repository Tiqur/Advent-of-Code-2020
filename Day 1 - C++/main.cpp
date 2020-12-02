// C++
#include <iostream>
#include <fstream>
#include <vector>


int part1(std::vector<int> input)
{
    for (int i = 0; i < input.size(); i++)
        for (int j = i; j < input.size(); j++)
            if (input[i] + input[j] == 2020) return input[i] * input[j];
    return 0;
}

int part2(std::vector<int> input)
{
    for (int i = 0; i < input.size(); i++)
        for (int j = i; j < input.size(); j++)
            for (int k = j; k < input.size(); k++)
                if (input[i] + input[j] + input[k] == 2020) return input[i] * input[j] * input[k];
    return 0;
}


int main() {
    std::ifstream input;
    std::vector<int> expenseList;
    input.open("input.txt");
    int a;

    while(input >> a)
        expenseList.emplace_back(a);

    std::cout << part1(expenseList) << std::endl;
    std::cout << part2(expenseList) << std::endl;

    return 0;
}
