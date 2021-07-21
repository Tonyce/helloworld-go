const data = [
    {a: [1,2,3]},
    4,
    [5,6,7],
    {b: {
        c: [8,9],
    }},
    [10, [11,12,13]],
    [14],
    15
]
[
   1,  2,  3,  4,  5,  6,
   7,  8,  9, 10, 11, 12,
  13, 14, 15
]


function trans(data, array) {
    if (Object.prototype.toString.call(data) == '[object Object]') {
        for (const key in data) {
            const element = data[key];
            trans(element, array)
        }
    } else if (Object.prototype.toString.call(data) == '[object Array]') {
        for (let index = 0; index < data.length; index++) {
            const element = data[index];
            trans(element, array)
        }
    } else {
        array.push(data)
    }
}

const a = []
trans(data, a)
console.log(a)