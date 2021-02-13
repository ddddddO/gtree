const env = process.env

let x = ''
if (typeof env.ENDPOINTX === 'undefined') {
	x = 'non'
} else {
	x = env.ENDPOINTX
}

console.log(x)
