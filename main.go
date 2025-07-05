package main

import (
  "bufio"
  "flag"
  "fmt"
  "math/rand"
  "os"
  "strings"
  "time"
)

var substitutions= map[string][]string { 
  "A": {"\u0100", "\u0102", "\u0104", "\u01CD", "\u01DE", "\uFF21"},
  "a": {"\u0101", "\u0103", "\u0105", "\u01CE", "\u01DF", "\u03B1", "\uFF41"},
  "C": {"\u0106", "\u0108", "\u010A", "\u010C", "\u2102", "\u212D", "\uFF23"},
  "c": {"\u0107", "\u0109", "\u010B", "\u010D", "\uFF43"},
  "D": {"\u010E", "\uFF24"},
  "d": {"\u010F", "\u0111", "\u03B4", "\uFF44"},
  "E": {"\u0112", "\u0114", "\u0116", "\u0118", "\u011A", "\u2107", "\u2130", "\uFF25"},
  "e": {"\u0113", "\u0115", "\u0117", "\u0119", "\u011B", "\u03B5", "\u212E", "\u213F", "\uFF45"},
  "G": {"\u011C", "\u011E", "\u0120", "\u0122", "\u01E4", "\u01E6", "\u0393", "\uFF27"},
  "g": {"\u011D", "\u011F", "\u0121", "\u0123", "\u01E5", "\u01E7", "\u0261", "\u210A", "\uFF47"},
  "H": {"\u0124", "\u0126", "\u210B", "\u210C", "\u210D", "\uFF28"},
  "h": {"\u0125", "\u0127", "\u04BB", "\u210E", "\uFF48"},
  "I": {"\u0128", "\u012A", "\u012C", "\u012E", "\u0130", "\u0197", "\u01CF", "\u2110", "\u2111", "\uFF29"},
  "i": {"\u0129", "\u012B", "\u012D", "\u012F", "\u0131", "\u01D0", "\uFF49"},
  "J": {"\u0134", "\uFF2A"},
  "j": {"\u0135", "\u01F0", "\uFF4A"},
  "K": {"\u0136", "\u01E8", "\u212A", "\uFF2B"},
  "k": {"\u0137", "\u01E9", "\uFF4B"},
  "L": {"\u0139", "\u013B", "\u013D", "\u0141", "\u2112", "\uFF2C"},
  "l": {"\u013A", "\u013C", "\u013E", "\u0142", "\u019A", "\u2113", "\uFF4C"},
  "N": {"\u0143", "\u0145", "\u0147", "\u2115", "\uFF2E"},
  "n": {"\u0144", "\u0146", "\u0148", "\u207F", "\u2229", "\uFF4E"},
  "O": {"\u014C", "\u014E", "\u0150", "\u019F", "\u01A0", "\u01D1", "\u01EA", "\u01EC", "\u03A9", "\uFF2F"},
  "o": {"\u014D", "\u014F", "\u0151", "\u01A1", "\u01D2", "\u01EB", "\u01ED", "\u2134", "\uFF4F"},
  "R": {"\u0154", "\u0156", "\u0158", "\u211B", "\u211C", "\u211D", "\uFF32"},
  "r": {"\u0155", "\u0157", "\u0159", "\uFF52"},
  "S": {"\u015A", "\u015C", "\u015E", "\u03A3", "\uFF33"},
  "s": {"\u015B", "\u015D", "\u015F", "\u03C3", "\uFF53"},
  "T": {"\u0162", "\u0164", "\u0166", "\u01AE", "\u0398", "\uFF34"},
  "t": {"\u0163", "\u0165", "\u0167", "\u028B", "\u03C4", "\uFF54"},
  "U": {"\u0168", "\u016A", "\u016C", "\u016E", "\u0170", "\u0172", "\u01AF", "\u01D3", "\u01D5", "\u01D7", "\u01D9", "\u01DB", "\uFF35"},
  "u": {"\u0169", "\u016B", "\u016D", "\u016F", "\u0171", "\u0173", "\u01B0", "\u01D4", "\u01D6", "\u01D8", "\u01DA", "\u01DC", "\uFF55"},
  "W": {"\u0174", "\uFF37"},
  "w": {"\u0175", "\uFF57"},
  "Y": {"\u0176", "\uFF39"},
  "y": {"\u0177", "\uFF59"},
  "Z": {"\u0179", "\u017B", "\u2124", "\u2128", "\uFF3A"},
  "z": {"\u017A", "\u017C", "\u01B6", "\uFF5A"},
  "b": {"\u0180", "\uFF42"},
  "|": {"\u01C0", "\u2223", "\u2758", "\uFF5C"},
  "!": {"\u01C3", "\uFF01"},
  "'": {"\u02B9", "\u02BC", "\u02C8", "\u201B", "\u2032", "\uFF07"},
  "\"": {"\u02BA", "\u030E", "\u201F", "\u2033", "\u2036", "\uFF02"},
  "^": {"\u02C4", "\u0302", "\u2303", "\uFF3E"},
  "`": {"\u02CB", "\u0300", "\u2035", "\uFF40"},
  "_": {"\u02CD", "\u0331", "\u0332", "\u2584", "\uFF3F"},
  "~": {"\u0303", "\u223C", "\uFF5E"},
  ";": {"\u037E", "\uFF1B"},
  "F": {"\u03A6", "\u2131", "\uFF26"},
  "p": {"\u03C0", "\uFF50"},
  "f": {"\u03C6", "\uFF46"},
  ":": {"\u0589", "\u2236", "\uFF1A"},
  "%": {"\u066A", "\uFF05"},
  " ": {"\u2000", "\u2001", "\u2002", "\u2003", "\u2004", "\u2005", "\u2006", "\u2008", "\u2009", "\u200A", "\u3000"},
  "-": {"\u2010", "\u2011", "\u2012", "\u2015", "\u2212", "\u2500", "\u252C", "\u2534", "\u2550", "\u2564", "\u2565", "\u2566", "\u2567", "\u2568", "\u2569", "\uFF0D"},
  "=": {"\u2017", "\u2261", "\u2264", "\u2265", "\uFF1D"},
  "/": {"\u2044", "\u2215", "\uFF0F"},
  "4": {"\u2074", "\u2084", "\uFF14"},
  "5": {"\u2075", "\u2085", "\uFF15"},
  "6": {"\u2076", "\u2086", "\uFF16"},
  "7": {"\u2077", "\u2087", "\uFF17"},
  "8": {"\u2078", "\u2088", "\u221E", "\uFF18"},
  "0": {"\u2080", "\uFF10"},
  "1": {"\u2081", "\uFF11"},
  "2": {"\u2082", "\uFF12"},
  "3": {"\u2083", "\uFF13"},
  "9": {"\u2089", "\uFF19"},
  "P": {"\u20A7", "\u2118", "\u2119", "\uFF30"},
  "Q": {"\u211A", "\uFF31"},
  "B": {"\u212C", "\uFF22"},
  "M": {"\u2133", "\uFF2D"},
  "\\": {"\u2216", "\uFF3C"},
  "*": {"\u2217", "\uFF0A"},
  "v": {"\u221A", "\uFF56"},
  "(": {"\u2320", "\uFF08"},
  ")": {"\u2321", "\uFF09"},
  "<": {"\u2329", "\u3008", "\uFF1C"},
  ">": {"\u232A", "\u3009", "\uFF1E"},
  "+": {"\u250C", "\u2510", "\u2514", "\u2518", "\u251C", "\u253C", "\u2552", "\u2553", "\u2554", "\u2555", "\u2556", "\u2557", "\u2558", "\u2559", "\u255A", "\u255B", "\u255C", "\u255D", "\u256A", "\u256B", "\u256C", "\uFF0B"},
  "[": {"\u301A", "\uFF3B"},
  "]": {"\u301B", "\uFF3D"},
  "#": {"\uFF03"},
  "$": {"\uFF04"},
  "&": {"\uFF06"},
  ",": {"\uFF0C"},
  ".": {"\uFF0E"},
  "@": {"\uFF20"},
  "V": {"\uFF36"},
  "X": {"\uFF38"},
  "m": {"\uFF4D"},
  "q": {"\uFF51"},
  "x": {"\uFF58"},
  "{": {"\uFF5B"},
  "}": {"\uFF5D"},
}

func encodeText(text string) string {
  rand.Seed(time.Now().UnixNano())
  var result strings.Builder
  for _, char := range text {
    upper := string(char)
    if subs, ok := substitutions[upper]; ok {
      result.WriteString(subs[rand.Intn(len(subs))])
    } else {
      result.WriteRune(char)
    }
  }
  return result.String()
}

func main() {
  textPtr := flag.String("text", "", "Text to encode")
  filePtr := flag.String("file", "", "File path to encode")
  flag.Parse()

  if *textPtr != "" && *filePtr != "" {
    fmt.Println("❌ Use only one argument --text or --file.")
    os.Exit(1)
  }

  if *textPtr != "" {
    encoded := encodeText(*textPtr)
    fmt.Println(encoded)
    return
  }

  if *filePtr != "" {
    file, err := os.Open(*filePtr)
    if err != nil {
      fmt.Printf("Can't open the file: %v\n", err)
      os.Exit(1)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      line := scanner.Text()
      fmt.Println(encodeText(line))
    }

    if err := scanner.Err(); err != nil {
      fmt.Printf("Erreur de lecture : %v\n", err)
    }
    return
  }

  fmt.Println("❗ Usage :")
  fmt.Println("  --text \"Text to encode\"")
  fmt.Println("  --file file_path.txt")
}
