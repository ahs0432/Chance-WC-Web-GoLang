var sort = "asc";
var sortnum = -1;

function sortTable(n) {
	var table, rows, rowstmp, i = 0;
	table = document.getElementById("listTable");
	rows = table.getElementsByTagName("TR");
	rowstmp = [];
	
	if (sortnum == -1) {
		sortnum = n;
	} else if (sortnum != n) {
		sortnum = n;
		sort = "asc";
	} else {
		if(sort == "asc") {
			sort = "desc";
		} else {
			sort = "asc";
		}
	}

	for (i = 1; i < rows.length; i++) {
		rowstmp[i-1] = rows[i];
	}

	var rowsResult = [];
	rowsResult = mergeSort(rowstmp, n);

	rowstmp = [];
	
	for (i = 0; i < rowsResult.length; i++) {
		rowstmp[i] = rowsResult[i].innerHTML;
	}

	for (i = 0; i < rowsResult.length; i++) {
		table.rows[i+1].innerHTML = rowstmp[i];
	}
}

function mergeSort(arr, n) {
	if (arr.length < 2) {
		return arr;
	}

	const middle = Math.floor(arr.length / 2);
	const left = arr.slice(0, middle);
	const right = arr.slice(middle);

	//window.alert("LEFT: " + left.length + ", RIGHT: " + right.length);

	return merge(mergeSort(left, n), mergeSort(right, n), n);

	function merge(left, right, n) {
		const resultArray = [];
		let leftIndex = 0;
		let rightIndex = 0;
	
		while (leftIndex < left.length && rightIndex < right.length) {
			if (n > 2) {
				if (sort == "asc") {
					if (left[leftIndex].getElementsByTagName("TD")[n].innerHTML.toLowerCase() < right[rightIndex].getElementsByTagName("TD")[n].innerHTML.toLowerCase()) {
						resultArray.push(left[leftIndex]);
						leftIndex++;
					} else {
						resultArray.push(right[rightIndex]);
						rightIndex++;
					}
				} else {
					if (left[leftIndex].getElementsByTagName("TD")[n].innerHTML.toLowerCase() > right[rightIndex].getElementsByTagName("TD")[n].innerHTML.toLowerCase()) {
						resultArray.push(left[leftIndex]);
						leftIndex++;
					} else {
						resultArray.push(right[rightIndex]);
						rightIndex++;
					}
				}
				
			} else if (n == 2) {
				if (sort == "asc") {
					if (parseFloat(left[leftIndex].getElementsByTagName("TD")[n].innerHTML.split(' ')[0]) < parseFloat(right[rightIndex].getElementsByTagName("TD")[n].innerHTML.split(' ')[0])) {
						resultArray.push(left[leftIndex]);
						leftIndex++;
					} else {
						resultArray.push(right[rightIndex]);
						rightIndex++;
					}
				} else {
					if (parseFloat(left[leftIndex].getElementsByTagName("TD")[n].innerHTML.split(' ')[0]) > parseFloat(right[rightIndex].getElementsByTagName("TD")[n].innerHTML.split(' ')[0])) {
						resultArray.push(left[leftIndex]);
						leftIndex++;
					} else {
						resultArray.push(right[rightIndex]);
						rightIndex++;
					}
				}
			} else {
				if (sort == "asc") {
					if (left[leftIndex].getElementsByTagName("TD")[n].getElementsByTagName("A")[0].innerHTML.toLowerCase() < right[rightIndex].getElementsByTagName("TD")[n].getElementsByTagName("A")[0].innerHTML.toLowerCase()) {
						resultArray.push(left[leftIndex]);
						leftIndex++;
					} else {
						resultArray.push(right[rightIndex]);
						rightIndex++;
					}
				} else {
					if (left[leftIndex].getElementsByTagName("TD")[n].getElementsByTagName("A")[0].innerHTML.toLowerCase() > right[rightIndex].getElementsByTagName("TD")[n].getElementsByTagName("A")[0].innerHTML.toLowerCase()) {
						resultArray.push(left[leftIndex]);
						leftIndex++;
					} else {
						resultArray.push(right[rightIndex]);
						rightIndex++;
					}
				}
			}
		}

		return resultArray.concat(left.slice(leftIndex), right.slice(rightIndex));
	}
}