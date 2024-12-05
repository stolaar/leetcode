const isPrefixOfWord = (sentence: string, searchWord: string): number => sentence.split(" ").findIndex((word) => word.startsWith(searchWord)) + 1


console.log(isPrefixOfWord("i love eating burger", "burg"))
