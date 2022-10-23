const storageSort = localStorage.getItem("sorted");

if (storageSort !== null) {
	if (storageSort === "true") {
		sort = localStorage.getItem("sort");
		sortTable(localStorage.getItem("sortnum"));
		localStorage.removeItem("sorted");
		localStorage.removeItem("sort");
		localStorage.removeItem("sortnum");
	}
}

let today = new Date();
document.getElementById('last-updated').textContent = today.toLocaleTimeString()
var count = 60;
document.getElementById('counter').textContent = count

setInterval(function(){
	count -= 1;
	document.getElementById('counter').textContent = count
	if(count <= 0) {
		if(sortnum != -1) {
			localStorage.setItem("sorted", "true");
			localStorage.setItem("sort", sort);
			localStorage.setItem("sortnum", sortnum);
		}

 		location.reload();
	}
}, 1000);