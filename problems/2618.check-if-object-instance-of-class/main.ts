function checkIfInstanceOf(obj: any, classFunction: any): boolean {
  if (typeof obj === "undefined" || obj === null) {
    return false
  }

  if (!classFunction) {
    return false
  }

  let proto = obj['__proto__']
  while (proto) {
    if (proto === classFunction.prototype) {
      return true
    }

    proto = proto['__proto__']
  }

  return false
};


console.log(checkIfInstanceOf(Error(), Error()))
