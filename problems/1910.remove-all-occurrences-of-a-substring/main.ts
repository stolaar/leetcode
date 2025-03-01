function removeOccurrences(s: string, part: string): string {
  const indexOf = s.indexOf(part)
  if (indexOf >= 0) {
    if (indexOf === 0) {
      return removeOccurrences(s.slice(part.length), part)
    }
    return removeOccurrences(s.slice(0, indexOf) + s.slice(indexOf + part.length), part)
  }

  return s
};
