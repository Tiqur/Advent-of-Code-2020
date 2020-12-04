// Disclaimer, I have never used C#
using System;
using System.Collections.Generic;
using System.IO;


namespace Day_4
{
    class Program
    {
        static void Main(string[] args)
        {
            List<List<KeyValuePair<string, string>>> formattedPassports = HelperFunc();
            Console.WriteLine("Part 1: " + part1(formattedPassports));
            Console.WriteLine("Part 2: " + part2(formattedPassports));
        }

        static List<List<KeyValuePair<string, string>>> HelperFunc()
        {
            string[] passports = File.ReadAllText("input.txt").Split("\n\r\n");
            var formatted = new List<List<KeyValuePair<string, string>>>() {};
            foreach(string p in passports)
            {
                string[] passportFields = p.Replace("\n", " ").Split(" ");
                var fields = new List<KeyValuePair<string, string>>() {};
                foreach (string pf in passportFields)
                {
                   string[] pair = pf.Split(":");
                   fields.Add(new KeyValuePair<string, string>(pair[0], pair[1]));
                }
                formatted.Add(fields);
            }
            return formatted;
        }


        static bool keysValid(List<KeyValuePair<string, string>> passports)
        {
            bool isValid = false;
                var allKeys = new List<string>(){};
                foreach (KeyValuePair<string, string> fields in passports)
                {
                    allKeys.Add(fields.Key);
                } 

                if (allKeys.Contains("byr") && 
                    allKeys.Contains("iyr") && 
                    allKeys.Contains("eyr") && 
                    allKeys.Contains("hgt") && 
                    allKeys.Contains("hcl") && 
                    allKeys.Contains("ecl") && 
                    allKeys.Contains("pid")) {
                        isValid = true;
                    }
            return isValid;
        }


        static int part1(List<List<KeyValuePair<string, string>>> passports)
        {
            int valid = 0;
            foreach (List<KeyValuePair<string, string>> passport in passports)
            {
                if (keysValid(passport)) valid++;
            }
            return valid;
        }
        

        static bool intBetween(string value, int min, int max)
        {
            return (Int32.Parse(value) < min || Int32.Parse(value) > max);
        }


        static int part2(List<List<KeyValuePair<string, string>>> passports)
        {
            int valid = 0;
            foreach (List<KeyValuePair<string, string>> passport in passports)
            {
                if (keysValid(passport)) {
                    bool isvalid = true;
                    foreach (KeyValuePair<string, string> fields in passport)
                    {
                            string value = fields.Value.Trim();


                            // input validation
                            switch(fields.Key)
                            {
                                case "byr":  // Birth year
                                    if (intBetween(value, 1920, 2002)) isvalid = false;
                                break;
                                case "iyr":  // Issue year
                                    if (intBetween(value, 2010, 2020)) isvalid = false;
                                break;
                                case "eyr":  // Expiration  year
                                    if (intBetween(value, 2020, 2030)) isvalid = false;
                                break;
                                case "hgt":  // Height
                                    string height = value.Substring(0, value.Length-2);
                                    string type = value.Substring(value.Length-2);
                                                                        
                                    if (height.Length == 0 || (type != "cm" && type != "in")) {
                                         isvalid = false;
                                    } else {
                                        if (intBetween(height, (type == "cm" ? 150 : 59), (type == "cm" ? 193 : 76))) isvalid = false;
                                    }

                                break;
                                case "hcl":  // Hair Color
                                    if (!value.StartsWith("#") || value.Length != 7) isvalid = false;
                                break;
                                case "ecl":
                                    if (!(
                                    value == "amb" || 
                                    value == "blu" || 
                                    value == "brn" || 
                                    value == "gry" || 
                                    value == "grn" || 
                                    value == "hzl" || 
                                    value == "oth" )) isvalid = false;
                                break;
                                case "pid":
                                    if (value.Length != 9) isvalid = false;
                                break;
                            }
                            
                    } 
                    if (isvalid) valid++;
                }
            }
            return valid;
        }
    }
}