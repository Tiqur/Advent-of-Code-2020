// Disclaimer, I have never used Dart before
import 'dart:io';

void main() async {
  var input = File("input.txt").readAsString();
  var groups = (await input).split("\n\r");

  print("Part 1: " + part1(groups).toString());
  print("Part 2: " + part2(groups).toString());

}

int part1(groups) {
  var total = 0;
  groups.forEach((ans) {
    Set answers = new Set();

    for (int i = 0; i < ans.length; i++)  
      if (ans.codeUnitAt(i) > 96) answers.add(ans[i]); // validate character

    // for each group, add the amount of unique answers
    total += answers.length;
  });

  return total;
}


int part2(groups) {
    var total = 0;
    groups.forEach((group) {
    // create map to store answers as key and amount as value: {'a': 0}
    Map answers = new Map();
    group.trim().split("\n").forEach((person) {
     person.split("").toSet().forEach((answer) { 
       // if answer doesn't exist in list, create new one and initialize value to 0 
       answers.putIfAbsent(answer, () => 0);
       // if answer does exist, increment it's value
       if (answers.containsKey(answer)) 
        answers[answer]++;
     });
    });

    // for each question that all groups agreed on ( map value equals group size ), increment total
    answers.forEach((key, value) {
      if (value == group.trim().split("\n").length)
        total++;
    });
  });

  return total;
}