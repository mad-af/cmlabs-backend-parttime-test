let n = 1;
const last = 100 ;
for ( ; n <= last ; n++ ) {
	let print = `${n}. `;
    const mThree = multipleCheks (n, 3);
    const mFive = multipleCheks (n, 5);
  
    if (mThree) {
      print += 'Fizz';
    }
    if (mFive) {
      print += 'Buzz';
    }
  	console.log(print);
}

function multipleCheks (n, x) {
  return n % x === 0;
}