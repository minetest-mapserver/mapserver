
export default [{
	input: 'main.js',
	output: {
		file :'bundle.js',
		format: 'umd',
		sourcemap: true,
		compact: true
	}
},{
	input: 'stats.js',
	output: {
		file :'bundle-stats.js',
		format: 'umd',
		sourcemap: true,
		compact: true
	}
}];
