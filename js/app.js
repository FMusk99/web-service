var name = "Some text here!(innerHTML)"
// setTimeout(() => {
// 	 var header = document.getElementById('header-title');
// 	 var headerStyle = document.getElementById('header-title').style;
// 	 headerStyle.opacity = 0;
// 	(function fade(){
// 		if(headerStyle.opacity == 1 ) 
// 		{
// 				headerStyle.opacity = 0
// 				headerStyle.display="block"
// 				delayOpacity(header,headerStyle)
// 				//console.log('-----end-----')
								
// 		} else {
// 			header.style.display="none"
// 			headerStyle.opacity = 1
// 			setTimeout(fade,500)
// 		}	
// 	})();
// }, 500)


async function delayOpacity(header, headerStyle) {
	header.innerHTML = name	
	headerStyle.display="block"
	for( i = 0; i<=10; i++){
		await delayTime(i,headerStyle).then((data)=> {})
		//console.log('awaiting: ', i)
	}
	
}

function delayTime(i,headerStyle) {
	return new Promise((resolve, reject) =>{
		setTimeout(()=>{
		headerStyle.opacity = i/10
		resolve({numb: i/10});
	},120)
	})
}

//to call sum function from module.js
//console.log(sumNumb(2, 3))