function deleteList(n) {
    if(confirm("정말 삭제하시겠습니까?") == true) {
        location.href='/delete/'+n    
    } else {
        return;
    }
}

function groupDelete(n) {
    if(confirm("정말 삭제하시겠습니까?") == true) {
        location.href='/groupdelete/'+n    
    } else {
        return;
    }
}